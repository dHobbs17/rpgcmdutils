package server

import (
	"encoding/json"
	"net"
)

const DISCONNECT_OPERATION string = "disconnect"
const CONNECT_OPERATION string = "connect"
const ACK_OPERATION string = "ack"
const INVALID_OPERATION string = "invalid"
const TICK_OPERATION string = "tick"
const QUEUE_OPERATION string = "queue"
const ACTION_OPERATION string = "action"

const SERVER_HOST string = "localhost:666"

type serverCommands int

type ServerMessage struct {
	Username string
	Action   string
	Data     string
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

const (
	DISCONNECT serverCommands = iota
	CONNECT
	ACK
	TICK
	QUEUE
	ACTION
)

var ServerOperations = map[serverCommands]string{
	DISCONNECT: DISCONNECT_OPERATION,
	CONNECT:    CONNECT_OPERATION,
	ACK:        ACK_OPERATION,
	TICK:       TICK_OPERATION,
	QUEUE:      QUEUE_OPERATION,
	ACTION:     ACTION_OPERATION,
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

	default:
		return INVALID_OPERATION
	}
}
