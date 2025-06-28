package server

import (
	"encoding/json"
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
)

// Server info
const server_host string = "localhost"
const server_port string = "666"
const server_protocol string = "tcp"
const server_name string = "Dave's cool server"

type serverCommands int

type ServerMessage struct {
	Action string
	Data   string
	Args   []string
}
type InitializeConnection struct {
	Action   string
	Data     string
	Conn     net.Conn
	Err      error
	Username string
	Encoder  *json.Encoder
}

func (s ServerMessage) String() string {
	return s.Action + ":" + s.Data
}

func (s InitializeConnection) String() string {
	return s.Action + ":" + s.Data
}

type ServerErr struct{ Err error }

func (e ServerErr) Error() string { return e.Err.Error() }

var ServerOperations = map[serverCommands]string{
	DISCONNECT: DISCONNECT_OPERATION,
	CONNECT:    CONNECT_OPERATION,
	ACK:        ACK_OPERATION,
	TICK:       TICK_OPERATION,
	QUEUE:      QUEUE_OPERATION,
	ACTION:     ACTION_OPERATION,
	NOTIFY:     NOTIFY_OPERATION,
	CHAT:       CHAT_OPERATION,
	ALERT:      ALERT_OPERATION,
	EVENT:      EVENT_OPERATION,
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
	case NOTIFY:
		return ServerOperations[NOTIFY]
	case ALERT:
		return ServerOperations[ALERT]
	case EVENT:
		return ServerOperations[EVENT]
	case CHAT:
		return ServerOperations[CHAT]
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
