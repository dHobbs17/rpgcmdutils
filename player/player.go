package player

import (
	"encoding/json"
	"fmt"
	"github.com/dHobbs17/rpgcmdutils/server"
	"log"
	"net"
	"strings"
)

type Player struct {
	name         string
	queuedAction *server.ServerMessage
	conn         net.Conn
	level        int
	id           int
	xp           int
	idle         int
	connected    bool
	statPoints   int
	class        Class
	stats        Stats
	skills       Skills
	spells       []string
	location     string
	encoder      *json.Encoder
	decoder      *json.Decoder
	quests       []int
}

type Stats struct {
	currentHp    int
	currentSp    int
	maxHp        int
	maxSp        int
	morale       int
	attack       int
	dodge        int
	parry        int
	block        int
	intelligence int
	strength     int
	dexterity    int
}

type Skills struct {
	destruction int
	conjuration int
	illusion    int
	perception  int
	deception   int
	stealth     int
	swords      int
	maces       int
	axes        int
	ranged      int
	wands       int
	block       int
	survival    int
}
type PlayerCommands int

type PlayerMessage struct {
	Action string
	Data   string
	Args   []string
}

// Player Sends...
const (
	MOVE_OPERATION       string = "move"
	ATTACK_OPERATION     string = "attack"
	GRAB_OPERATION       string = "grab"
	CAST_OPERATION       string = "cast"
	USE_OPERATION        string = "use"
	GUARD_OPERATION      string = "guard"
	SEARCH_OPERATION     string = "search"
	TARGET_OPERATION     string = "target"
	STALK_OPERATION      string = "stalk"
	FOLLOW_OPERATION     string = "follow"
	PARTY_OPERATION      string = "party"
	CLAN_OPERATION       string = "clan"
	LOOT_OPERATION       string = "loot"
	INVITE_OPERATION     string = "invite"
	LEAVE_OPERATION      string = "leave"
	LOOK_OPERATION       string = "look"
	SAY_OPERATION        string = "say"
	HIDE_OPERATION       string = "hide"
	DISCONNECT_OPERATION string = "disconnect"
	INVALID_OPERATION    string = "invalid"
)

const (
	MOVE PlayerCommands = iota // Move must be first
	ATTACK
	CAST
	SEARCH
	USE
	GUARD
	GRAB
	TARGET
	STALK
	FOLLOW
	PARTY
	CLAN
	LOOT
	INVITE
	LEAVE
	LOOK
	SAY
	HIDE
	DISCONNECT // Disconnect must be last
)

var playerOperations = map[PlayerCommands]string{
	MOVE:       MOVE_OPERATION,
	ATTACK:     ATTACK_OPERATION,
	CAST:       CAST_OPERATION,
	USE:        USE_OPERATION,
	GRAB:       GRAB_OPERATION,
	GUARD:      GUARD_OPERATION,
	SEARCH:     SEARCH_OPERATION,
	TARGET:     TARGET_OPERATION,
	STALK:      STALK_OPERATION,
	FOLLOW:     FOLLOW_OPERATION,
	PARTY:      PARTY_OPERATION,
	CLAN:       CLAN_OPERATION,
	LOOK:       LOOK_OPERATION,
	LOOT:       LOOT_OPERATION,
	INVITE:     INVITE_OPERATION,
	LEAVE:      LEAVE_OPERATION,
	SAY:        SAY_OPERATION,
	HIDE:       HIDE_OPERATION,
	DISCONNECT: DISCONNECT_OPERATION,
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

	// TODO Update for multi space commands
	command, data, _ := strings.Cut(operation, " ") // bind 3rd param "Found: bool"
	log.Println("Player entered Command: "+command, len(command))
	log.Println("Player entered Data: "+data, len(data))

	var mappedCommand = MapPlayerOperations(command)

	if mappedCommand == INVALID_OPERATION {
		return PlayerMessage{}, PlayerError{fmt.Errorf("not a valid player command")}
	} else {
		return PlayerMessage{Action: mappedCommand, Data: data}, PlayerError{nil}
	}
}

// do we need this?
func MapPlayerOperations(s string) string {
	switch s {
	case MOVE_OPERATION:
		return playerOperations[MOVE]
	case ATTACK_OPERATION:
		return playerOperations[ATTACK]
	case GRAB_OPERATION:
		return playerOperations[GRAB]
	case USE_OPERATION:
		return playerOperations[USE]
	case CAST_OPERATION:
		return playerOperations[CAST]
	case SEARCH_OPERATION:
		return playerOperations[SEARCH]
	case TARGET_OPERATION:
		return playerOperations[TARGET]
	case GUARD_OPERATION:
		return playerOperations[GUARD]
	case STALK_OPERATION:
		return playerOperations[STALK]
	case FOLLOW_OPERATION:
		return playerOperations[FOLLOW]
	case PARTY_OPERATION:
		return playerOperations[PARTY]
	case CLAN_OPERATION:
		return playerOperations[CLAN]
	case LOOK_OPERATION:
		return playerOperations[LOOK]
	case LOOT_OPERATION:
		return playerOperations[LOOT]
	case INVITE_OPERATION:
		return playerOperations[INVITE]
	case LEAVE_OPERATION:
		return playerOperations[LEAVE]
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
	case ATTACK:
		return playerOperations[ATTACK]
	case SEARCH:
		return playerOperations[SEARCH]
	case GRAB:
		return playerOperations[GRAB]
	case CAST:
		return playerOperations[CAST]
	case GUARD:
		return playerOperations[GUARD]
	case USE:
		return playerOperations[USE]
	case TARGET:
		return playerOperations[TARGET]
	case STALK:
		return playerOperations[STALK]
	case FOLLOW:
		return playerOperations[FOLLOW]
	case PARTY:
		return playerOperations[PARTY]
	case CLAN:
		return playerOperations[CLAN]
	case LOOK:
		return playerOperations[LOOK]
	case LOOT:
		return playerOperations[LOOT]
	case INVITE:
		return playerOperations[INVITE]
	case LEAVE:
		return playerOperations[LEAVE]
	case DISCONNECT:
		return playerOperations[DISCONNECT]
	default:
		return INVALID_OPERATION
	}
}

// Errors
type PlayerError struct{ Err error }

func (e PlayerError) Error() string { return e.Err.Error() }
