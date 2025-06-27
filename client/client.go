package client

const HEARTBEAT_OPERATION string = "heartbeat"
const CONNECT_OPERATION string = "connect"
const DISCONNECT_OPERATION string = "disconnect"
const INVALID_OPERATION string = "invalid"

type clientCommand int

type ClientError struct{ err error }

func (e ClientError) Error() string { return e.err.Error() }

type ClientMessage struct {
	Username string
	Action   string
	Data     string
}

type ReturnControl struct {
}

const (
	CONNECT clientCommand = iota
	DISCONNECT
	HEARTBEAT
)

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
