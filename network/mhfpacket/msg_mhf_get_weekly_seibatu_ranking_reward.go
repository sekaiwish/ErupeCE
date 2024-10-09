package mhfpacket

import (
	"errors"

	"erupe-ce/network"
	"erupe-ce/network/clientctx"
	"erupe-ce/utils/byteframe"
)

// MsgMhfGetWeeklySeibatuRankingReward represents the MSG_MHF_GET_WEEKLY_SEIBATU_RANKING_REWARD
type MsgMhfGetWeeklySeibatuRankingReward struct {
	AckHandle uint32
	Unk0      uint32
	Unk1      uint32
	Unk2      uint32
	Unk3      uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetWeeklySeibatuRankingReward) Opcode() network.PacketID {
	return network.MSG_MHF_GET_WEEKLY_SEIBATU_RANKING_REWARD
}

// Parse parses the packet from binary
func (m *MsgMhfGetWeeklySeibatuRankingReward) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint32()
	m.Unk1 = bf.ReadUint32()
	m.Unk2 = bf.ReadUint32()
	m.Unk3 = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetWeeklySeibatuRankingReward) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
