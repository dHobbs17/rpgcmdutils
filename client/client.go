package client

// Client Sends...
const HEARTBEAT_OPERATION string = "heartbeat"
const CONNECT_OPERATION string = "connect"
const DISCONNECT_OPERATION string = "disconnect"
const INVALID_OPERATION string = "invalid"

const (
	CONNECT clientCommand = iota
	DISCONNECT
	HEARTBEAT
)

type clientCommand int

type ClientError struct{ Err error }

func (e ClientError) Error() string { return e.Err.Error() }

type ClientMessage struct {
	Action string
	Data   string
	Args   []string
}

type ReturnControl struct {
}

var ServerOperations = map[clientCommand]string{
	CONNECT:    CONNECT_OPERATION,
	DISCONNECT: DISCONNECT_OPERATION,
	HEARTBEAT:  HEARTBEAT_OPERATION,
}

func (s clientCommand) String() string {
	switch s {
	case CONNECT:
		return ServerOperations[CONNECT]
	case DISCONNECT:
		return ServerOperations[DISCONNECT]
	case HEARTBEAT:
		return ServerOperations[HEARTBEAT]
	default:
		return INVALID_OPERATION
	}
}
