package player

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type Player struct {
	name         string
	queuedAction *PlayerMessage
	conn         net.Conn
	level        int
	id           int
	xp           int
	idle         int
	dead         bool
	gold         int
	lootable     bool
	connected    bool
	target       *int
	statPoints   int
	class        Class
	stats        Stats
	inCombat     bool
	skills       Skills
	spells       []string
	location     int
	encoder      *json.Encoder
	decoder      *json.Decoder
	quests       []int
	inventory    []string
	equipment    equipment
	loot         []string
}

type equipment struct {
	helm      string
	armor     string
	leftHand  string
	rightHand string
	legs      string
	boots     string
	bracelet  string
	gloves    string
	ring1     string
	ring2     string
}
type Stats struct {
	currentHp    int
	currentSp    int
	maxHp        int
	maxSp        int
	morale       int
	hit          int
	attack       int
	dodge        int
	parry        int
	block        int
	intelligence int
	strength     int
	dexterity    int
	notoriety    int
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
type PlayerEquipment int

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
	SWAP_OPERATION       string = "swap"
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
	SWAP
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

const (
	HELM PlayerEquipment = iota // Move must be first
	ARMOR
	LEGS
	BOOTS
	LEFT_HAND
	RIGHT_HAND
	GLOVES
	BRACELET
	RING1
	RING2
)

const (
	EQUIPMENT_HELM      = "helm"
	EQUIPMENT_ARMOR     = "armor"
	EQUIPMENT_BOOTS     = "boots"
	EQUIPMENT_GLOVES    = "gloves"
	EQUIPMENT_BRACELET  = "bracelet"
	EQUIPMENT_LEGS      = "legs"
	EQUIPMENT_RING1     = "ring1"
	EQUIPMENT_RING2     = "ring2"
	EQUIPMENT_LEFTHAND  = "lefthand"
	EQUIPMENT_RIGHTHAND = "righthand"
)

var playerEquipment = map[PlayerEquipment]string{
	HELM:       EQUIPMENT_HELM,
	ARMOR:      EQUIPMENT_ARMOR,
	LEGS:       EQUIPMENT_LEGS,
	BOOTS:      EQUIPMENT_BOOTS,
	BRACELET:   EQUIPMENT_BRACELET,
	GLOVES:     EQUIPMENT_GLOVES,
	LEFT_HAND:  EQUIPMENT_LEFTHAND,
	RIGHT_HAND: EQUIPMENT_RIGHTHAND,
	RING1:      EQUIPMENT_RING1,
	RING2:      EQUIPMENT_RING2,
}

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
	SWAP:       SWAP_OPERATION,
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
	case SWAP_OPERATION:
		return playerOperations[SWAP]
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
	case SWAP:
		return playerOperations[SWAP]
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

// getters and setters
func (p *Player) GetConn() net.Conn  { return p.conn }
func (p *Player) SetConn(c net.Conn) { p.conn = c }
func (p *Player) ClearConn() {
	p.conn.Close()
	p.conn = nil
}

func (p *Player) GetConnected() bool  { return p.connected }
func (p *Player) SetConnected(b bool) { p.connected = b }

func (p *Player) GetIdle() int { return p.idle }
func (p *Player) AdjustIdle(i int) {
	p.idle += i
	if p.idle <= 0 {
		p.idle = 0
	}
}
func (p *Player) SetIdle(i int) { p.idle = i }

func (p *Player) GetHit() int    { return p.stats.hit }
func (p *Player) GetAttack() int { return p.stats.attack }

func (p *Player) GetId() int { return p.id }

func (p *Player) GetName() string { return p.name }

func (p *Player) IsLootable() bool { return p.lootable }
func (p *Player) IsAlive() bool    { return !p.dead }

func (p *Player) GetQueuedAction() *PlayerMessage  { return p.queuedAction }
func (p *Player) SetQueuedAction(a *PlayerMessage) { p.queuedAction = a }
func (p *Player) ClearQueuedAction()               { p.queuedAction = nil }

func (p *Player) GetEncoder() *json.Encoder  { return p.encoder }
func (p *Player) SetEncoder(e *json.Encoder) { p.encoder = e }
func (p *Player) ClearEncoder()              { p.encoder = nil }

func (p *Player) GetGold() int { return p.stats.currentHp }
func (p *Player) AdjustGold(g int) {
	p.gold += g
	if p.gold <= 0 {
		p.gold = 0
	}
}

// TODO Add Loot IDs
func (p *Player) AddToInventory(loot string) {
	p.inventory = append(p.inventory, loot)
}

// TODO Implement this
//func (p *Player) dropFromInventory(loot string) {
//	p.inventory = append(p.inventory, loot)
//}

func (p *Player) KillPlayer() {
	// add and drop inventory
	p.loot = p.inventory
	p.inventory = []string{}

	// add and drop gold
	p.loot = append(p.loot, strconv.Itoa(p.gold)+" gold")
	p.gold = 0

	// add and drop equipment
	inv := reflect.ValueOf(p.equipment)
	for i := 0; i < inv.NumField(); i++ {
		p.loot = append(p.loot, inv.Field(i).String())
	}
	p.equipment = equipment{}

	// mark dead and lootable
	p.dead = true
	p.lootable = true
}

func (p *Player) GetHp() int { return p.stats.currentHp }
func (p *Player) ResetHp()   { p.stats.currentHp = p.stats.maxHp }
func (p *Player) AdjustHp(hp int) {
	p.stats.currentHp += hp
	if p.stats.currentHp <= 0 {
		p.stats.currentHp = 0
		p.dead = true
		p.lootable = true
	}
}
func (p *Player) IsInCombat() bool { return p.inCombat }
func (p *Player) SetCombat(c bool) { p.inCombat = c }

func (p *Player) GetTarget() *int          { return p.target }
func (p *Player) SetTarget(targetsId *int) { p.target = targetsId }
func (p *Player) ResetTarget()             { p.target = nil }

func (p *Player) GetNotoriety() int     { return p.stats.notoriety }
func (p *Player) SetNotoriety(n int)    { p.stats.notoriety = n }
func (p *Player) AdjustNotoriety(n int) { p.stats.notoriety += n }

func (p *Player) GetSp() int { return p.stats.currentSp }
func (p *Player) ResetSp()   { p.stats.currentSp = p.stats.maxSp }
func (p *Player) AdjustSp(sp int) {
	p.stats.currentSp += sp
	if p.stats.currentSp <= 0 {
		p.stats.currentSp = 0
	}
}

func NewPlayer(conn net.Conn, name string) Player {
	return Player{id: rand.Int(), // TODO Check for collisions,
		stats: Stats{
			currentHp: 10,
			maxHp:     10,
			hit:       1,
			attack:    1,
		},
		conn:     conn,
		name:     name,
		class:    Novice,
		location: 0,
		level:    1,
	}
}
