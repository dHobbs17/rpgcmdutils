package player

import (
	"fmt"
	"log"
	"strings"
)

const MOVE_OPERATION string = "move"
const SAY_OPERATION string = "say"
const HIDE_OPERATION string = "hide"
const DISCONNECT_OPERATION string = "disconnect"
const INVALID_OPERATION string = "invalid"

type PlayerCommands int

type PlayerMessage struct {
	Action string
	Data   string
}

type PlayerError struct{ Err error }

func (e PlayerError) Error() string { return e.Err.Error() }

const (
	MOVE PlayerCommands = iota // Move must be first
	SAY
	HIDE
	DISCONNECT // Disconnect must be last
)

var playerOperations = map[PlayerCommands]string{
	MOVE:       MOVE_OPERATION,
	SAY:        SAY_OPERATION,
	DISCONNECT: DISCONNECT_OPERATION,
	HIDE:       HIDE_OPERATION,
}

func GetPlayerOperations() []string {
	var availablePlayerOperations []string
	for d := MOVE; d <= DISCONNECT; d++ {
		availablePlayerOperations = append(availablePlayerOperations, playerOperations[d])
	}
	return availablePlayerOperations
}

func ValidatePlayerOperation(operation string) (PlayerMessage, PlayerError) {
	log.Println("Received player message: "+operation, len(operation))

	// TODO fix bug with say so it allows spaces
	var command, data string
	if strings.Contains(operation, " ") {
		command = strings.Split(operation, " ")[0]
		data = strings.Split(operation, " ")[1]
		log.Println("Player entered Command: "+command, len(command))
		log.Println("Player entered Data: "+data, len(data))
	} else {
		command = strings.Split(operation, " ")[0]
		data = ""
		log.Println("Player entered Command: "+command, len(command))
		log.Println("Player entered Data: "+data, len(data))
	}

	var mappedCommand = MapPlayerOperations(command)

	if mappedCommand == INVALID_OPERATION {
		return PlayerMessage{}, PlayerError{fmt.Errorf("not a valid player command")}
	} else {
		return PlayerMessage{mappedCommand, data}, PlayerError{nil}
	}
}

// do we need this?
func MapPlayerOperations(s string) string {
	switch s {
	case MOVE_OPERATION:
		return playerOperations[MOVE]
	case SAY_OPERATION:
		return playerOperations[SAY]
	case DISCONNECT_OPERATION:
		return playerOperations[DISCONNECT]
	case HIDE_OPERATION:
		return playerOperations[HIDE]
	default:
		return INVALID_OPERATION
	}
}

// do we need this?
func (s PlayerCommands) String() string {
	switch s {
	case MOVE:
		return playerOperations[MOVE]
	case SAY:
		return playerOperations[SAY]
	case HIDE:
		return playerOperations[HIDE]
	case DISCONNECT:
		return playerOperations[DISCONNECT]
	default:
		return INVALID_OPERATION
	}
}
