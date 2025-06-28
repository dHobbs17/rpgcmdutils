package server

import (
	"encoding/json"
	"net"
)

const DISCONNECT_OPERATION string = "disconnect"
const ACK_OPERATION string = "ack"
const INVALID_OPERATION string = "invalid"
const TICK_OPERATION string = "tick"

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
	ACK
	TICK
)

var ServerOperations = map[serverCommands]string{
	DISCONNECT: DISCONNECT_OPERATION,
	ACK:        ACK_OPERATION,
	TICK:       TICK_OPERATION,
}

func (s serverCommands) String() string {
	switch s {
	case DISCONNECT:
		return ServerOperations[DISCONNECT]
	case ACK:
		return ServerOperations[ACK]
	case TICK:
		return ServerOperations[TICK]
	default:
		return INVALID_OPERATION
	}
}
