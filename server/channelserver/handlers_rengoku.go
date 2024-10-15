package channelserver

import (
	"erupe-ce/config"
	"erupe-ce/utils/db"
	ps "erupe-ce/utils/pascalstring"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"

	"erupe-ce/network/mhfpacket"
	"erupe-ce/utils/byteframe"

	"go.uber.org/zap"
)

func handleMsgMhfSaveRengokuData(s *Session, p mhfpacket.MHFPacket) {
	// saved every floor on road, holds values such as floors progressed, points etc.
	// can be safely handled by the client
	pkt := p.(*mhfpacket.MsgMhfSaveRengokuData)
	database, err := db.GetDB()
	if err != nil {
		s.Logger.Fatal(fmt.Sprintf("Failed to get database instance: %s", err))
	}
	dumpSaveData(s, pkt.RawDataPayload, "rengoku")
	_, err = database.Exec("UPDATE characters SET rengokudata=$1 WHERE id=$2", pkt.RawDataPayload, s.CharID)
	if err != nil {
		s.Logger.Error("Failed to save rengokudata", zap.Error(err))
		s.DoAckSimpleSucceed(pkt.AckHandle, make([]byte, 4))
		return
	}
	bf := byteframe.NewByteFrameFromBytes(pkt.RawDataPayload)
	bf.Seek(71, 0)
	maxStageMp := bf.ReadUint32()
	maxScoreMp := bf.ReadUint32()
	bf.Seek(4, 1)
	maxStageSp := bf.ReadUint32()
	maxScoreSp := bf.ReadUint32()
	var t int
	err = database.QueryRow("SELECT character_id FROM rengoku_score WHERE character_id=$1", s.CharID).Scan(&t)
	if err != nil {
		database.Exec("INSERT INTO rengoku_score (character_id) VALUES ($1)", s.CharID)
	}
	database.Exec("UPDATE rengoku_score SET max_stages_mp=$1, max_points_mp=$2, max_stages_sp=$3, max_points_sp=$4 WHERE character_id=$5", maxStageMp, maxScoreMp, maxStageSp, maxScoreSp, s.CharID)
	s.DoAckSimpleSucceed(pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfLoadRengokuData(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfLoadRengokuData)
	var data []byte
	database, err := db.GetDB()
	if err != nil {
		s.Logger.Fatal(fmt.Sprintf("Failed to get database instance: %s", err))
	}
	err = database.QueryRow("SELECT rengokudata FROM characters WHERE id = $1", s.CharID).Scan(&data)
	if err != nil {
		s.Logger.Error("Failed to load rengokudata", zap.Error(err))
	}
	if len(data) > 0 {
		s.DoAckBufSucceed(pkt.AckHandle, data)
	} else {
		resp := byteframe.NewByteFrame()
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint16(0)
		resp.WriteUint32(0)
		resp.WriteUint16(0)
		resp.WriteUint16(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0) // an extra 4 bytes were missing based on pcaps

		resp.WriteUint8(3) // Count of next 3
		resp.WriteUint16(0)
		resp.WriteUint16(0)
		resp.WriteUint16(0)

		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)

		resp.WriteUint8(3) // Count of next 3
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)

		resp.WriteUint8(3) // Count of next 3
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)

		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)
		resp.WriteUint32(0)

		s.DoAckBufSucceed(pkt.AckHandle, resp.Data())
	}
}

func handleMsgMhfGetRengokuBinary(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetRengokuBinary)
	// a (massively out of date) version resides in the game's /dat/ folder or up to date can be pulled from packets
	data, err := os.ReadFile(filepath.Join(config.GetConfig().BinPath, "rengoku_data.bin"))
	if err != nil {
		panic(err)
	}
	s.DoAckBufSucceed(pkt.AckHandle, data)
}

const rengokuScoreQuery = `, c.name FROM rengoku_score rs
LEFT JOIN characters c ON c.id = rs.character_id
LEFT JOIN guild_characters gc ON gc.character_id = rs.character_id `

type RengokuScore struct {
	Name  string `db:"name"`
	Score uint32 `db:"score"`
}

func handleMsgMhfEnumerateRengokuRanking(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfEnumerateRengokuRanking)

	guild, _ := GetGuildInfoByCharacterId(s, s.CharID)
	isApplicant, _ := guild.HasApplicationForCharID(s, s.CharID)
	if isApplicant {
		guild = nil
	}

	if pkt.Leaderboard == 2 || pkt.Leaderboard == 3 || pkt.Leaderboard == 6 || pkt.Leaderboard == 7 {
		if guild == nil {
			s.DoAckBufSucceed(pkt.AckHandle, make([]byte, 11))
			return
		}
	}

	var score RengokuScore
	var selfExist bool
	i := uint32(1)
	bf := byteframe.NewByteFrame()
	scoreData := byteframe.NewByteFrame()
	database, err := db.GetDB()
	if err != nil {
		s.Logger.Fatal(fmt.Sprintf("Failed to get database instance: %s", err))
	}
	var rows *sqlx.Rows
	switch pkt.Leaderboard {
	case 0:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_stages_mp AS score %s ORDER BY max_stages_mp DESC", rengokuScoreQuery))
	case 1:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_points_mp AS score %s ORDER BY max_points_mp DESC", rengokuScoreQuery))
	case 2:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_stages_mp AS score %s WHERE guild_id=$1 ORDER BY max_stages_mp DESC", rengokuScoreQuery), guild.ID)
	case 3:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_points_mp AS score %s WHERE guild_id=$1 ORDER BY max_points_mp DESC", rengokuScoreQuery), guild.ID)
	case 4:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_stages_sp AS score %s ORDER BY max_stages_sp DESC", rengokuScoreQuery))
	case 5:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_points_sp AS score %s ORDER BY max_points_sp DESC", rengokuScoreQuery))
	case 6:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_stages_sp AS score %s WHERE guild_id=$1 ORDER BY max_stages_sp DESC", rengokuScoreQuery), guild.ID)
	case 7:
		rows, _ = database.Queryx(fmt.Sprintf("SELECT max_points_sp AS score %s WHERE guild_id=$1 ORDER BY max_points_sp DESC", rengokuScoreQuery), guild.ID)
	}

	for rows.Next() {
		rows.StructScan(&score)
		if score.Name == s.Name {
			bf.WriteUint32(i)
			bf.WriteUint32(score.Score)
			ps.Uint8(bf, s.Name, true)
			ps.Uint8(bf, "", false)
			selfExist = true
		}
		if i > 100 {
			i++
			continue
		}
		scoreData.WriteUint32(i)
		scoreData.WriteUint32(score.Score)
		ps.Uint8(scoreData, score.Name, true)
		ps.Uint8(scoreData, "", false)
		i++
	}

	if !selfExist {
		bf.WriteBytes(make([]byte, 10))
	}
	bf.WriteUint8(uint8(i) - 1)
	bf.WriteBytes(scoreData.Data())
	s.DoAckBufSucceed(pkt.AckHandle, bf.Data())
}

func handleMsgMhfGetRengokuRankingRank(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetRengokuRankingRank)
	// What is this for?
	bf := byteframe.NewByteFrame()
	bf.WriteUint32(0) // Max stage overall MP rank
	bf.WriteUint32(0) // Max RdP overall MP rank
	s.DoAckBufSucceed(pkt.AckHandle, bf.Data())
}
