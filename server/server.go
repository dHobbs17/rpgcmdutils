package server

import (
	"encoding/json"
	"github.com/dHobbs17/rpgcmdutils/player"
	"net"
)

// Server Sends...
const DISCONNECT_OPERATION string = "disconnect"
const CONNECT_OPERATION string = "connect"
const ACK_OPERATION string = "ack"
const INVALID_OPERATION string = "invalid"
const TICK_OPERATION string = "tick"
const QUEUE_OPERATION string = "queue"
const ACTION_OPERATION string = "action"
const NOTIFY_OPERATION string = "notify"
const ALERT_OPERATION string = "alert"
const EVENT_OPERATION string = "event"
const CHAT_OPERATION string = "chat"
const YELL_OPERATION string = "yell"

const GET_PLAYER_OPERATION string = "get_player"
const PLAYER_UPDATE_OPERATION string = "player_update"
const PLAYER_SWING_OPERATION string = "player_swing"
const PLAYER_HIT_OPERATION string = "player_hit"
const PLAYER_MISS_OPERATION string = "player_miss"
const PLAYER_DEAD_OPERATION string = "player_dead"
const PLAYER_TARGET_DEAD_OPERATION string = "player_target_dead"
const PLAYER_DIALOG_OPERATION string = "player_dialog"
const PLAYER_TARGET_OPERATION string = "player_target"

const NPC_TARGET_OPERATION string = "npc_target"
const NPC_SWING_OPERATION string = "npc_swing"
const NPC_HIT_OPERATION string = "npc_hit"
const NPC_MISS_OPERATION string = "npc_miss"
const NPC_DEAD_OPERATION string = "npc_dead"
const NPC_TARGET_DEAD_OPERATION string = "npc_target_dead"
const NPC_DIALOG_OPERATION string = "npc_dialog"

const (
	DISCONNECT serverCommands = iota
	CONNECT
	ACK
	TICK
	QUEUE
	ACTION
	NOTIFY
	CHAT
	ALERT
	EVENT
	YELL
	PLAYER_UPDATE
	GET_PLAYER
	PLAYER_SWING
	PLAYER_HIT
	PLAYER_MISS
	PLAYER_DEAD
	PLAYER_TARGET
	PLAYER_TARGET_DEAD
	PLAYER_DIALOG
	NPC_TARGET
	NPC_TARGET_DEAD
	NPC_SWING
	NPC_HIT
	NPC_MISS
	NPC_DEAD
	NPC_DIALOG
)

// Server info
const server_host string = "localhost"
const server_port string = "666"
const server_protocol string = "tcp"
const server_name string = "Dave's cool server"

type serverCommands int

type Data interface{}
type ServerMessage struct {
	Action string
	Data   Data
	Args   []string
}

type InitializeConnection struct {
	Action  string
	Data    string
	Conn    net.Conn
	Err     error
	Player  player.Player
	Encoder *json.Encoder
}

func (s InitializeConnection) String() string {
	return s.Action + ":" + s.Data
}

type ServerErr struct{ Err error }

func (e ServerErr) Error() string { return e.Err.Error() }

var ServerOperations = map[serverCommands]string{
	DISCONNECT:         DISCONNECT_OPERATION,
	CONNECT:            CONNECT_OPERATION,
	ACK:                ACK_OPERATION,
	TICK:               TICK_OPERATION,
	QUEUE:              QUEUE_OPERATION,
	ACTION:             ACTION_OPERATION,
	NOTIFY:             NOTIFY_OPERATION,
	CHAT:               CHAT_OPERATION,
	ALERT:              ALERT_OPERATION,
	EVENT:              EVENT_OPERATION,
	YELL:               YELL_OPERATION,
	PLAYER_UPDATE:      PLAYER_UPDATE_OPERATION,
	PLAYER_SWING:       PLAYER_SWING_OPERATION,
	PLAYER_HIT:         PLAYER_HIT_OPERATION,
	PLAYER_DEAD:        PLAYER_DEAD_OPERATION,
	PLAYER_DIALOG:      PLAYER_DIALOG_OPERATION,
	PLAYER_MISS:        PLAYER_MISS_OPERATION,
	PLAYER_TARGET:      PLAYER_TARGET_OPERATION,
	NPC_TARGET:         NPC_TARGET_OPERATION,
	NPC_SWING:          NPC_SWING_OPERATION,
	NPC_DIALOG:         NPC_DIALOG_OPERATION,
	NPC_MISS:           NPC_MISS_OPERATION,
	NPC_DEAD:           NPC_DEAD_OPERATION,
	NPC_HIT:            NPC_HIT_OPERATION,
	PLAYER_TARGET_DEAD: PLAYER_TARGET_DEAD_OPERATION,
	NPC_TARGET_DEAD:    NPC_TARGET_DEAD_OPERATION,
	GET_PLAYER:         GET_PLAYER_OPERATION,
}

func (s serverCommands) String() string {
	switch s {
	case DISCONNECT:
		return ServerOperations[DISCONNECT]
	case CONNECT:
		return ServerOperations[CONNECT]
	case ACK:
		return ServerOperations[ACK]
	case TICK:
		return ServerOperations[TICK]
	case QUEUE:
		return ServerOperations[QUEUE]
	case ACTION:
		return ServerOperations[ACTION]
	case PLAYER_UPDATE:
		return ServerOperations[PLAYER_UPDATE]
	case PLAYER_DEAD:
		return ServerOperations[PLAYER_DEAD]
	case PLAYER_MISS:
		return ServerOperations[PLAYER_MISS]
	case PLAYER_HIT:
		return ServerOperations[PLAYER_HIT]
	case PLAYER_SWING:
		return ServerOperations[PLAYER_SWING]
	case PLAYER_DIALOG:
		return ServerOperations[PLAYER_DIALOG]
	case PLAYER_TARGET:
		return ServerOperations[PLAYER_TARGET]
	case PLAYER_TARGET_DEAD:
		return ServerOperations[PLAYER_TARGET_DEAD]
	case NPC_TARGET:
		return ServerOperations[NPC_TARGET]
	case NPC_SWING:
		return ServerOperations[NPC_SWING]
	case NPC_DIALOG:
		return ServerOperations[NPC_DIALOG]
	case NPC_MISS:
		return ServerOperations[NPC_MISS]
	case NPC_DEAD:
		return ServerOperations[NPC_DEAD]
	case NPC_TARGET_DEAD:
		return ServerOperations[NPC_TARGET_DEAD]
	case GET_PLAYER:
		return ServerOperations[GET_PLAYER]
	case NOTIFY:
		return ServerOperations[NOTIFY]
	case ALERT:
		return ServerOperations[ALERT]
	case EVENT:
		return ServerOperations[EVENT]
	case CHAT:
		return ServerOperations[CHAT]
	case YELL:
		return ServerOperations[YELL]
	default:
		return INVALID_OPERATION
	}
}

func GetServerUrl() string {
	return server_host + ":" + server_port
}

func GetServerProtocol() string {
	return server_protocol
}

func GetServername() string {
	return server_name
}
