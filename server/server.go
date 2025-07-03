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
const COMBAT_MOB_ATTACK_OPERATION string = "combat_mob_attack"
const COMBAT_HIT_OPERATION string = "combat_hit"
const COMBAT_KILL_OPERATION string = "combat_kill"
const COMBAT_DIALOG_OPERATION string = "combat_dialog"
const ALERT_OPERATION string = "alert"
const EVENT_OPERATION string = "event"
const CHAT_OPERATION string = "chat"
const YELL_OPERATION string = "yell"
const TARGET_MOB_OPERATION string = "target_mob"
const TARGET_PLAYER_OPERATION string = "target_player"
const PLAYER_UPDATE_OPERATION string = "player_update"
const GET_PLAYER_OPERATION string = "get_player"

const (
	DISCONNECT serverCommands = iota
	CONNECT
	ACK
	TICK
	QUEUE
	COMBAT_MOB_ATTACK
	COMBAT_HIT
	COMBAT_KILL
	COMBAT_DIALOG
	ACTION
	NOTIFY
	CHAT
	ALERT
	EVENT
	YELL
	TARGET_MOB
	TARGET_PLAYER
	PLAYER_UPDATE
	GET_PLAYER
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
	DISCONNECT:        DISCONNECT_OPERATION,
	CONNECT:           CONNECT_OPERATION,
	ACK:               ACK_OPERATION,
	TICK:              TICK_OPERATION,
	QUEUE:             QUEUE_OPERATION,
	COMBAT_MOB_ATTACK: COMBAT_MOB_ATTACK_OPERATION,
	COMBAT_HIT:        COMBAT_HIT_OPERATION,
	COMBAT_KILL:       COMBAT_KILL_OPERATION,
	COMBAT_DIALOG:     COMBAT_DIALOG_OPERATION,
	ACTION:            ACTION_OPERATION,
	NOTIFY:            NOTIFY_OPERATION,
	CHAT:              CHAT_OPERATION,
	ALERT:             ALERT_OPERATION,
	EVENT:             EVENT_OPERATION,
	YELL:              YELL_OPERATION,
	TARGET_MOB:        TARGET_MOB_OPERATION,
	TARGET_PLAYER:     TARGET_PLAYER_OPERATION,
	PLAYER_UPDATE:     PLAYER_UPDATE_OPERATION,
	GET_PLAYER:        GET_PLAYER_OPERATION,
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
	case COMBAT_MOB_ATTACK:
		return ServerOperations[COMBAT_MOB_ATTACK]
	case COMBAT_HIT:
		return ServerOperations[COMBAT_HIT]
	case COMBAT_KILL:
		return ServerOperations[COMBAT_KILL]
	case COMBAT_DIALOG:
		return ServerOperations[COMBAT_DIALOG]
	case TARGET_MOB:
		return ServerOperations[TARGET_MOB]
	case TARGET_PLAYER:
		return ServerOperations[TARGET_PLAYER]
	case PLAYER_UPDATE:
		return ServerOperations[PLAYER_UPDATE]
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
