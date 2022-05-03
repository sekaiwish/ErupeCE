package network

//revive:disable
type PacketID uint16

//go:generate stringer -type=PacketID
const (
	MSG_HEAD PacketID = iota
	MSG_SYS_reserve01
	MSG_SYS_reserve02
	MSG_SYS_reserve03
	MSG_SYS_reserve04
	MSG_SYS_reserve05
	MSG_SYS_reserve06
	MSG_SYS_reserve07
	MSG_SYS_ADD_OBJECT
	MSG_SYS_DEL_OBJECT
	MSG_SYS_DISP_OBJECT
	MSG_SYS_HIDE_OBJECT
	MSG_SYS_reserve0C
	MSG_SYS_reserve0D
	MSG_SYS_reserve0E
	MSG_SYS_EXTEND_THRESHOLD
	MSG_SYS_END
	MSG_SYS_NOP
	MSG_SYS_ACK
	MSG_SYS_TERMINAL_LOG
	MSG_SYS_LOGIN
	MSG_SYS_LOGOUT
	MSG_SYS_SET_STATUS
	MSG_SYS_PING
	MSG_SYS_CAST_BINARY
	MSG_SYS_HIDE_CLIENT
	MSG_SYS_TIME
	MSG_SYS_CASTED_BINARY
	MSG_SYS_GET_FILE
	MSG_SYS_ISSUE_LOGKEY
	MSG_SYS_RECORD_LOG
	MSG_SYS_ECHO
	MSG_SYS_CREATE_STAGE
	MSG_SYS_STAGE_DESTRUCT
	MSG_SYS_ENTER_STAGE
	MSG_SYS_BACK_STAGE
	MSG_SYS_MOVE_STAGE
	MSG_SYS_LEAVE_STAGE
	MSG_SYS_LOCK_STAGE
	MSG_SYS_UNLOCK_STAGE
	MSG_SYS_RESERVE_STAGE
	MSG_SYS_UNRESERVE_STAGE
	MSG_SYS_SET_STAGE_PASS
	MSG_SYS_WAIT_STAGE_BINARY
	MSG_SYS_SET_STAGE_BINARY
	MSG_SYS_GET_STAGE_BINARY
	MSG_SYS_ENUMERATE_CLIENT
	MSG_SYS_ENUMERATE_STAGE
	MSG_SYS_CREATE_MUTEX
	MSG_SYS_CREATE_OPEN_MUTEX
	MSG_SYS_DELETE_MUTEX
	MSG_SYS_OPEN_MUTEX
	MSG_SYS_CLOSE_MUTEX
	MSG_SYS_CREATE_SEMAPHORE
	MSG_SYS_CREATE_ACQUIRE_SEMAPHORE
	MSG_SYS_DELETE_SEMAPHORE
	MSG_SYS_ACQUIRE_SEMAPHORE
	MSG_SYS_RELEASE_SEMAPHORE
	MSG_SYS_LOCK_GLOBAL_SEMA
	MSG_SYS_UNLOCK_GLOBAL_SEMA
	MSG_SYS_CHECK_SEMAPHORE
	MSG_SYS_OPERATE_REGISTER
	MSG_SYS_LOAD_REGISTER
	MSG_SYS_NOTIFY_REGISTER
	MSG_SYS_CREATE_OBJECT
	MSG_SYS_DELETE_OBJECT
	MSG_SYS_POSITION_OBJECT
	MSG_SYS_ROTATE_OBJECT
	MSG_SYS_DUPLICATE_OBJECT
	MSG_SYS_SET_OBJECT_BINARY
	MSG_SYS_GET_OBJECT_BINARY
	MSG_SYS_GET_OBJECT_OWNER
	MSG_SYS_UPDATE_OBJECT_BINARY
	MSG_SYS_CLEANUP_OBJECT
	MSG_SYS_reserve4A
	MSG_SYS_reserve4B
	MSG_SYS_reserve4C
	MSG_SYS_reserve4D
	MSG_SYS_reserve4E
	MSG_SYS_reserve4F
	MSG_SYS_INSERT_USER
	MSG_SYS_DELETE_USER
	MSG_SYS_SET_USER_BINARY
	MSG_SYS_GET_USER_BINARY
	MSG_SYS_NOTIFY_USER_BINARY
	MSG_SYS_reserve55
	MSG_SYS_reserve56
	MSG_SYS_reserve57
	MSG_SYS_UPDATE_RIGHT
	MSG_SYS_AUTH_QUERY
	MSG_SYS_AUTH_DATA
	MSG_SYS_AUTH_TERMINAL
	MSG_SYS_reserve5C
	MSG_SYS_RIGHTS_RELOAD
	MSG_SYS_reserve5E
	MSG_SYS_reserve5F
	MSG_MHF_SAVEDATA
	MSG_MHF_LOADDATA
	MSG_MHF_LIST_MEMBER
	MSG_MHF_OPR_MEMBER
	MSG_MHF_ENUMERATE_DIST_ITEM
	MSG_MHF_APPLY_DIST_ITEM
	MSG_MHF_ACQUIRE_DIST_ITEM
	MSG_MHF_GET_DIST_DESCRIPTION
	MSG_MHF_SEND_MAIL
	MSG_MHF_READ_MAIL
	MSG_MHF_LIST_MAIL
	MSG_MHF_OPRT_MAIL
	MSG_MHF_LOAD_FAVORITE_QUEST
	MSG_MHF_SAVE_FAVORITE_QUEST
	MSG_MHF_REGISTER_EVENT
	MSG_MHF_RELEASE_EVENT
	MSG_MHF_TRANSIT_MESSAGE
	MSG_SYS_reserve71
	MSG_SYS_reserve72
	MSG_SYS_reserve73
	MSG_SYS_reserve74
	MSG_SYS_reserve75
	MSG_SYS_reserve76
	MSG_SYS_reserve77
	MSG_SYS_reserve78
	MSG_SYS_reserve79
	MSG_SYS_reserve7A
	MSG_SYS_reserve7B
	MSG_SYS_reserve7C
	MSG_CA_EXCHANGE_ITEM
	MSG_SYS_reserve7E
	MSG_MHF_PRESENT_BOX
	MSG_MHF_SERVER_COMMAND
	MSG_MHF_SHUT_CLIENT
	MSG_MHF_ANNOUNCE
	MSG_MHF_SET_LOGINWINDOW
	MSG_SYS_TRANS_BINARY
	MSG_SYS_COLLECT_BINARY
	MSG_SYS_GET_STATE
	MSG_SYS_SERIALIZE
	MSG_SYS_ENUMLOBBY
	MSG_SYS_ENUMUSER
	MSG_SYS_INFOKYSERVER
	MSG_MHF_GET_CA_UNIQUE_ID
	MSG_MHF_SET_CA_ACHIEVEMENT
	MSG_MHF_CARAVAN_MY_SCORE
	MSG_MHF_CARAVAN_RANKING
	MSG_MHF_CARAVAN_MY_RANK
	MSG_MHF_CREATE_GUILD
	MSG_MHF_OPERATE_GUILD
	MSG_MHF_OPERATE_GUILD_MEMBER
	MSG_MHF_INFO_GUILD
	MSG_MHF_ENUMERATE_GUILD
	MSG_MHF_UPDATE_GUILD
	MSG_MHF_ARRANGE_GUILD_MEMBER
	MSG_MHF_ENUMERATE_GUILD_MEMBER
	MSG_MHF_ENUMERATE_CAMPAIGN
	MSG_MHF_STATE_CAMPAIGN
	MSG_MHF_APPLY_CAMPAIGN
	MSG_MHF_ENUMERATE_ITEM
	MSG_MHF_ACQUIRE_ITEM
	MSG_MHF_TRANSFER_ITEM
	MSG_MHF_MERCENARY_HUNTDATA
	MSG_MHF_ENTRY_ROOKIE_GUILD
	MSG_MHF_ENUMERATE_QUEST
	MSG_MHF_ENUMERATE_EVENT
	MSG_MHF_ENUMERATE_PRICE
	MSG_MHF_ENUMERATE_RANKING
	MSG_MHF_ENUMERATE_ORDER
	MSG_MHF_ENUMERATE_SHOP
	MSG_MHF_GET_EXTRA_INFO
	MSG_MHF_UPDATE_INTERIOR
	MSG_MHF_ENUMERATE_HOUSE
	MSG_MHF_UPDATE_HOUSE
	MSG_MHF_LOAD_HOUSE
	MSG_MHF_OPERATE_WAREHOUSE
	MSG_MHF_ENUMERATE_WAREHOUSE
	MSG_MHF_UPDATE_WAREHOUSE
	MSG_MHF_ACQUIRE_TITLE
	MSG_MHF_ENUMERATE_TITLE
	MSG_MHF_ENUMERATE_GUILD_ITEM
	MSG_MHF_UPDATE_GUILD_ITEM
	MSG_MHF_ENUMERATE_UNION_ITEM
	MSG_MHF_UPDATE_UNION_ITEM
	MSG_MHF_CREATE_JOINT
	MSG_MHF_OPERATE_JOINT
	MSG_MHF_INFO_JOINT
	MSG_MHF_UPDATE_GUILD_ICON
	MSG_MHF_INFO_FESTA
	MSG_MHF_ENTRY_FESTA
	MSG_MHF_CHARGE_FESTA
	MSG_MHF_ACQUIRE_FESTA
	MSG_MHF_STATE_FESTA_U
	MSG_MHF_STATE_FESTA_G
	MSG_MHF_ENUMERATE_FESTA_MEMBER
	MSG_MHF_VOTE_FESTA
	MSG_MHF_ACQUIRE_CAFE_ITEM
	MSG_MHF_UPDATE_CAFEPOINT
	MSG_MHF_CHECK_DAILY_CAFEPOINT
	MSG_MHF_GET_COG_INFO
	MSG_MHF_CHECK_MONTHLY_ITEM
	MSG_MHF_ACQUIRE_MONTHLY_ITEM
	MSG_MHF_CHECK_WEEKLY_STAMP
	MSG_MHF_EXCHANGE_WEEKLY_STAMP
	MSG_MHF_CREATE_MERCENARY
	MSG_MHF_SAVE_MERCENARY
	MSG_MHF_READ_MERCENARY_W
	MSG_MHF_READ_MERCENARY_M
	MSG_MHF_CONTRACT_MERCENARY
	MSG_MHF_ENUMERATE_MERCENARY_LOG
	MSG_MHF_ENUMERATE_GUACOT
	MSG_MHF_UPDATE_GUACOT
	MSG_MHF_INFO_TOURNAMENT
	MSG_MHF_ENTRY_TOURNAMENT
	MSG_MHF_ENTER_TOURNAMENT_QUEST
	MSG_MHF_ACQUIRE_TOURNAMENT
	MSG_MHF_GET_ACHIEVEMENT
	MSG_MHF_RESET_ACHIEVEMENT
	MSG_MHF_ADD_ACHIEVEMENT
	MSG_MHF_PAYMENT_ACHIEVEMENT
	MSG_MHF_DISPLAYED_ACHIEVEMENT
	MSG_MHF_INFO_SCENARIO_COUNTER
	MSG_MHF_SAVE_SCENARIO_DATA
	MSG_MHF_LOAD_SCENARIO_DATA
	MSG_MHF_GET_BBS_SNS_STATUS
	MSG_MHF_APPLY_BBS_ARTICLE
	MSG_MHF_GET_ETC_POINTS
	MSG_MHF_UPDATE_ETC_POINT
	MSG_MHF_GET_MYHOUSE_INFO
	MSG_MHF_UPDATE_MYHOUSE_INFO
	MSG_MHF_GET_WEEKLY_SCHEDULE
	MSG_MHF_ENUMERATE_INV_GUILD
	MSG_MHF_OPERATION_INV_GUILD
	MSG_MHF_STAMPCARD_STAMP
	MSG_MHF_STAMPCARD_PRIZE
	MSG_MHF_UNRESERVE_SRG
	MSG_MHF_LOAD_PLATE_DATA
	MSG_MHF_SAVE_PLATE_DATA
	MSG_MHF_LOAD_PLATE_BOX
	MSG_MHF_SAVE_PLATE_BOX
	MSG_MHF_READ_GUILDCARD
	MSG_MHF_UPDATE_GUILDCARD
	MSG_MHF_READ_BEAT_LEVEL
	MSG_MHF_UPDATE_BEAT_LEVEL
	MSG_MHF_READ_BEAT_LEVEL_ALL_RANKING
	MSG_MHF_READ_BEAT_LEVEL_MY_RANKING
	MSG_MHF_READ_LAST_WEEK_BEAT_RANKING
	MSG_MHF_ACCEPT_READ_REWARD
	MSG_MHF_GET_ADDITIONAL_BEAT_REWARD
	MSG_MHF_GET_FIXED_SEIBATU_RANKING_TABLE
	MSG_MHF_GET_BBS_USER_STATUS
	MSG_MHF_KICK_EXPORT_FORCE
	MSG_MHF_GET_BREAK_SEIBATU_LEVEL_REWARD
	MSG_MHF_GET_WEEKLY_SEIBATU_RANKING_REWARD
	MSG_MHF_GET_EARTH_STATUS
	MSG_MHF_LOAD_PARTNER
	MSG_MHF_SAVE_PARTNER
	MSG_MHF_GET_GUILD_MISSION_LIST
	MSG_MHF_GET_GUILD_MISSION_RECORD
	MSG_MHF_ADD_GUILD_MISSION_COUNT
	MSG_MHF_SET_GUILD_MISSION_TARGET
	MSG_MHF_CANCEL_GUILD_MISSION_TARGET
	MSG_MHF_LOAD_OTOMO_AIROU
	MSG_MHF_SAVE_OTOMO_AIROU
	MSG_MHF_ENUMERATE_GUILD_TRESURE
	MSG_MHF_ENUMERATE_AIROULIST
	MSG_MHF_REGIST_GUILD_TRESURE
	MSG_MHF_ACQUIRE_GUILD_TRESURE
	MSG_MHF_OPERATE_GUILD_TRESURE_REPORT
	MSG_MHF_GET_GUILD_TRESURE_SOUVENIR
	MSG_MHF_ACQUIRE_GUILD_TRESURE_SOUVENIR
	MSG_MHF_ENUMERATE_FESTA_INTERMEDIATE_PRIZE
	MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE
	MSG_MHF_LOAD_DECO_MYSET
	MSG_MHF_SAVE_DECO_MYSET
	MSG_MHF_reserve010F
	MSG_MHF_LOAD_GUILD_COOKING
	MSG_MHF_REGIST_GUILD_COOKING
	MSG_MHF_LOAD_GUILD_ADVENTURE
	MSG_MHF_REGIST_GUILD_ADVENTURE
	MSG_MHF_ACQUIRE_GUILD_ADVENTURE
	MSG_MHF_CHARGE_GUILD_ADVENTURE
	MSG_MHF_LOAD_LEGEND_DISPATCH
	MSG_MHF_LOAD_HUNTER_NAVI
	MSG_MHF_SAVE_HUNTER_NAVI
	MSG_MHF_REGIST_SPABI_TIME
	MSG_MHF_GET_GUILD_WEEKLY_BONUS_MASTER
	MSG_MHF_GET_GUILD_WEEKLY_BONUS_ACTIVE_COUNT
	MSG_MHF_ADD_GUILD_WEEKLY_BONUS_EXCEPTIONAL_USER
	MSG_MHF_GET_TOWER_INFO
	MSG_MHF_POST_TOWER_INFO
	MSG_MHF_GET_GEM_INFO
	MSG_MHF_POST_GEM_INFO
	MSG_MHF_GET_EARTH_VALUE
	MSG_MHF_DEBUG_POST_VALUE
	MSG_MHF_GET_PAPER_DATA
	MSG_MHF_GET_NOTICE
	MSG_MHF_POST_NOTICE
	MSG_MHF_GET_BOOST_TIME
	MSG_MHF_POST_BOOST_TIME
	MSG_MHF_GET_BOOST_TIME_LIMIT
	MSG_MHF_POST_BOOST_TIME_LIMIT
	MSG_MHF_ENUMERATE_FESTA_PERSONAL_PRIZE
	MSG_MHF_ACQUIRE_FESTA_PERSONAL_PRIZE
	MSG_MHF_GET_RAND_FROM_TABLE
	MSG_MHF_GET_CAFE_DURATION
	MSG_MHF_GET_CAFE_DURATION_BONUS_INFO
	MSG_MHF_RECEIVE_CAFE_DURATION_BONUS
	MSG_MHF_POST_CAFE_DURATION_BONUS_RECEIVED
	MSG_MHF_GET_GACHA_POINT
	MSG_MHF_USE_GACHA_POINT
	MSG_MHF_EXCHANGE_FPOINT_2_ITEM
	MSG_MHF_EXCHANGE_ITEM_2_FPOINT
	MSG_MHF_GET_FPOINT_EXCHANGE_LIST
	MSG_MHF_PLAY_STEPUP_GACHA
	MSG_MHF_RECEIVE_GACHA_ITEM
	MSG_MHF_GET_STEPUP_STATUS
	MSG_MHF_PLAY_FREE_GACHA
	MSG_MHF_GET_TINY_BIN
	MSG_MHF_POST_TINY_BIN
	MSG_MHF_GET_SENYU_DAILY_COUNT
	MSG_MHF_GET_GUILD_TARGET_MEMBER_NUM
	MSG_MHF_GET_BOOST_RIGHT
	MSG_MHF_START_BOOST_TIME
	MSG_MHF_POST_BOOST_TIME_QUEST_RETURN
	MSG_MHF_GET_BOX_GACHA_INFO
	MSG_MHF_PLAY_BOX_GACHA
	MSG_MHF_RESET_BOX_GACHA_INFO
	MSG_MHF_GET_SEIBATTLE
	MSG_MHF_POST_SEIBATTLE
	MSG_MHF_GET_RYOUDAMA
	MSG_MHF_POST_RYOUDAMA
	MSG_MHF_GET_TENROUIRAI
	MSG_MHF_POST_TENROUIRAI
	MSG_MHF_POST_GUILD_SCOUT
	MSG_MHF_CANCEL_GUILD_SCOUT
	MSG_MHF_ANSWER_GUILD_SCOUT
	MSG_MHF_GET_GUILD_SCOUT_LIST
	MSG_MHF_GET_GUILD_MANAGE_RIGHT
	MSG_MHF_SET_GUILD_MANAGE_RIGHT
	MSG_MHF_PLAY_NORMAL_GACHA
	MSG_MHF_GET_DAILY_MISSION_MASTER
	MSG_MHF_GET_DAILY_MISSION_PERSONAL
	MSG_MHF_SET_DAILY_MISSION_PERSONAL
	MSG_MHF_GET_GACHA_PLAY_HISTORY
	MSG_MHF_GET_REJECT_GUILD_SCOUT
	MSG_MHF_SET_REJECT_GUILD_SCOUT
	MSG_MHF_GET_CA_ACHIEVEMENT_HIST
	MSG_MHF_SET_CA_ACHIEVEMENT_HIST
	MSG_MHF_GET_KEEP_LOGIN_BOOST_STATUS
	MSG_MHF_USE_KEEP_LOGIN_BOOST
	MSG_MHF_GET_UD_SCHEDULE
	MSG_MHF_GET_UD_INFO
	MSG_MHF_GET_KIJU_INFO
	MSG_MHF_SET_KIJU
	MSG_MHF_ADD_UD_POINT
	MSG_MHF_GET_UD_MY_POINT
	MSG_MHF_GET_UD_TOTAL_POINT_INFO
	MSG_MHF_GET_UD_BONUS_QUEST_INFO
	MSG_MHF_GET_UD_SELECTED_COLOR_INFO
	MSG_MHF_GET_UD_MONSTER_POINT
	MSG_MHF_GET_UD_DAILY_PRESENT_LIST
	MSG_MHF_GET_UD_NORMA_PRESENT_LIST
	MSG_MHF_GET_UD_RANKING_REWARD_LIST
	MSG_MHF_ACQUIRE_UD_ITEM
	MSG_MHF_GET_REWARD_SONG
	MSG_MHF_USE_REWARD_SONG
	MSG_MHF_ADD_REWARD_SONG_COUNT
	MSG_MHF_GET_UD_RANKING
	MSG_MHF_GET_UD_MY_RANKING
	MSG_MHF_ACQUIRE_MONTHLY_REWARD
	MSG_MHF_GET_UD_GUILD_MAP_INFO
	MSG_MHF_GENERATE_UD_GUILD_MAP
	MSG_MHF_GET_UD_TACTICS_POINT
	MSG_MHF_ADD_UD_TACTICS_POINT
	MSG_MHF_GET_UD_TACTICS_RANKING
	MSG_MHF_GET_UD_TACTICS_REWARD_LIST
	MSG_MHF_GET_UD_TACTICS_LOG
	MSG_MHF_GET_EQUIP_SKIN_HIST
	MSG_MHF_UPDATE_EQUIP_SKIN_HIST
	MSG_MHF_GET_UD_TACTICS_FOLLOWER
	MSG_MHF_SET_UD_TACTICS_FOLLOWER
	MSG_MHF_GET_UD_SHOP_COIN
	MSG_MHF_USE_UD_SHOP_COIN
	MSG_MHF_GET_ENHANCED_MINIDATA
	MSG_MHF_SET_ENHANCED_MINIDATA
	MSG_MHF_SEX_CHANGER
	MSG_MHF_GET_LOBBY_CROWD
	MSG_SYS_reserve180
	MSG_MHF_GUILD_HUNTDATA
	MSG_MHF_ADD_KOURYOU_POINT
	MSG_MHF_GET_KOURYOU_POINT
	MSG_MHF_EXCHANGE_KOURYOU_POINT
	MSG_MHF_GET_UD_TACTICS_BONUS_QUEST
	MSG_MHF_GET_UD_TACTICS_FIRST_QUEST_BONUS
	MSG_MHF_GET_UD_TACTICS_REMAINING_POINT
	MSG_SYS_reserve188
	MSG_MHF_LOAD_PLATE_MYSET
	MSG_MHF_SAVE_PLATE_MYSET
	MSG_SYS_reserve18B
	MSG_MHF_GET_RESTRICTION_EVENT
	MSG_MHF_SET_RESTRICTION_EVENT
	MSG_SYS_reserve18E
	MSG_SYS_reserve18F
	MSG_MHF_GET_TREND_WEAPON
	MSG_MHF_UPDATE_USE_TREND_WEAPON_LOG
	MSG_SYS_reserve192
	MSG_SYS_reserve193
	MSG_SYS_reserve194
	MSG_MHF_SAVE_RENGOKU_DATA
	MSG_MHF_LOAD_RENGOKU_DATA
	MSG_MHF_GET_RENGOKU_BINARY
	MSG_MHF_ENUMERATE_RENGOKU_RANKING
	MSG_MHF_GET_RENGOKU_RANKING_RANK
	MSG_MHF_ACQUIRE_EXCHANGE_SHOP
	MSG_SYS_reserve19B
	MSG_MHF_SAVE_MEZFES_DATA
	MSG_MHF_LOAD_MEZFES_DATA
	MSG_SYS_reserve19E
	MSG_SYS_reserve19F
	MSG_MHF_UPDATE_FORCE_GUILD_RANK
	MSG_MHF_RESET_TITLE
	MSG_MHF_ENUMERATE_GUILD_MESSAGE_BOARD
	MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
	MSG_SYS_reserve204
	MSG_SYS_reserve205
	MSG_SYS_reserve206
	MSG_SYS_reserve207
	MSG_SYS_reserve208
	MSG_SYS_reserve209
	MSG_SYS_reserve20A
	MSG_SYS_reserve20B
	MSG_SYS_reserve20C
	MSG_SYS_reserve20D
	MSG_SYS_reserve20E
	MSG_SYS_reserve20F
)

//revive:enable
