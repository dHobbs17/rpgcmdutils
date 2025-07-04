package player

import (
	"encoding/json"
	"fmt"
	"github.com/dHobbs17/rpgcmdutils/common"
	"log"
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type Player struct {
	Name         string
	QueuedAction *common.Action
	Conn         net.Conn
	Level        int
	Id           int
	Xp           int
	Idle         int
	Dead         bool
	Gold         int
	Lootable     bool
	Connected    bool
	Target       *common.Target
	StatPoints   int
	Class        Class
	Stats        common.Stats
	InCombat     bool
	IsPlayer     bool
	Skills       common.Skills
	Spells       []string
	Location     int
	Encoder      *json.Encoder
	Decoder      *json.Decoder
	Quests       []int
	Inventory    []string
	Equipment    equipment
	Loot         []string
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

type PlayerCommands int
type PlayerEquipment int

type PlayerMessage struct {
	Action string
	Data   string
	Args   []string
}

type PlayerContentMessage struct {
	Action string
	Data   Player
	Args   []string
}

//func (m PlayerContentMessage) MarshalJSON() ([]byte, error) {
//	j, err := json.Marshal(struct {
//		Action string
//		Data   Player
//		Args   []string
//	}{
//		Action: m.Action,
//		Data: Player{
//			Name:         m.Data.name,
//			QueuedAction: m.Data.queuedAction,
//			Conn:         m.Data.conn,
//			Level:        m.Data.level,
//			Id:           m.Data.id,
//			Xp:           m.Data.xp,
//			Idle:         m.Data.idle,
//			Dead:         m.Data.dead,
//			Gold:         m.Data.gold,
//			Lootable:     m.Data.lootable,
//			Connected:    m.Data.connected,
//			Target:       m.Data.target,
//			StatPoints:   m.Data.statPoints,
//			Class:        m.Data.class,
//			Stats:        m.Data.stats,
//			InCombat:     m.Data.inCombat,
//			Skills:       m.Data.skills,
//			Spells:       m.Data.spells,
//			Location:     m.Data.location,
//			Encoder:      m.Data.encoder,
//			Decoder:      m.Data.decoder,
//			Quests:       m.Data.quests,
//			Inventory:    m.Data.inventory,
//			Equipment:    m.Data.equipment,
//			Loot:         m.Data.loot,
//		},
//		Args: m.Args,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return j, nil
//}

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
//func (p *Player) Get() Player { return *p }

// func (p *Player) GetConn() net.Conn  { return p.conn }
// func (p *Player) SetConn(c net.Conn) { p.conn = c }
func (p *Player) ClearConn() {
	p.Conn.Close()
	p.Conn = nil
}

//func (p *Player) GetConnected() bool  { return p.connected }
//func (p *Player) SetConnected(b bool) { p.connected = b }

// func (p *Player) GetIdle() int { return p.idle }
func (p *Player) AdjustIdle(i int) {
	p.Idle += i
	if p.Idle <= 0 {
		p.Idle = 0
	}
}

//func (p *Player) SetIdle(i int)  { p.idle = i }
//func (p *Player) GetLevel() int  { return p.level }
//func (p *Player) GetHit() int    { return p.stats.Hit }
//func (p *Player) GetAttack() int { return p.stats.Attack }
//
//func (p *Player) GetId() int { return p.id }
//
//func (p *Player) GetName() string { return p.name }

//func (p *Player) IsLootable() bool                 { return p.lootable }
//func (p *Player) IsAlive() bool                    { return !p.dead }
//func (p *Player) IsPlayer() bool                   { return true }
//func (p *Player) GetQueuedAction() *common.Action  { return p.queuedAction }
//func (p *Player) SetQueuedAction(a *common.Action) { p.queuedAction = a }
//func (p *Player) ClearQueuedAction()               { p.queuedAction = nil }

// func (p *Player) GetEncoder() *json.Encoder  { return p.encoder }
// func (p *Player) SetEncoder(e *json.Encoder) { p.encoder = e }
func (p *Player) ClearEncoder()   { p.Encoder = nil }
func (p *Player) CalcHit() int    { return p.Stats.Hit }
func (p *Player) CalcDamage() int { return p.Stats.Attack }

// func (p *Player) GetGold() int { return p.stats.CurrentHp }
func (p *Player) AdjustGold(g int) {
	p.Gold += g
	if p.Gold <= 0 {
		p.Gold = 0
	}
}

// TODO Add Loot IDs
func (p *Player) AddToInventory(loot string) {
	p.Inventory = append(p.Inventory, loot)
}

// TODO Implement this
//func (p *Player) dropFromInventory(loot string) {
//	p.inventory = append(p.inventory, loot)
//}

func (p *Player) KillPlayer() {
	// add and drop inventory
	p.Loot = p.Inventory
	p.Inventory = []string{}

	// add and drop gold
	p.Loot = append(p.Loot, strconv.Itoa(p.Gold)+" gold")
	p.Gold = 0

	// add and drop equipment
	inv := reflect.ValueOf(p.Equipment)
	for i := 0; i < inv.NumField(); i++ {
		p.Loot = append(p.Loot, inv.Field(i).String())
	}
	p.Equipment = equipment{}

	// mark dead and lootable
	p.Dead = true
	p.Lootable = true
}

func (p *Player) ResetHp() { p.Stats.CurrentHp = p.Stats.MaxHp }
func (p *Player) AdjustHp(hp int) {
	p.Stats.CurrentHp += hp
	if p.Stats.CurrentHp <= 0 {
		p.Stats.CurrentHp = 0
		p.Dead = true
		p.Lootable = true
	}
}

//func (p *Player) IsInCombat() bool { return p.inCombat }
//func (p *Player) SetCombat(c bool) { p.inCombat = c }

func (p *Player) GetTarget() *common.Target       { return p.Target }
func (p *Player) SetTarget(target *common.Target) { p.Target = target }
func (p *Player) ResetTarget()                    { p.Target = nil }

func (p *Player) GetReputation() int     { return p.Stats.Reputation }
func (p *Player) SetReputation(n int)    { p.Stats.Reputation = n }
func (p *Player) AdjustReputation(n int) { p.Stats.Reputation += n }

func (p *Player) GetStats() common.Stats   { return p.Stats }
func (p *Player) GetSkills() common.Skills { return p.Skills }

func (p *Player) GetSp() int    { return p.Stats.CurrentSp }
func (p *Player) GetMaxSp() int { return p.Stats.MaxSp }
func (p *Player) ResetSp()      { p.Stats.CurrentSp = p.Stats.MaxSp }
func (p *Player) AdjustSp(sp int) {
	p.Stats.CurrentSp += sp
	if p.Stats.CurrentSp <= 0 {
		p.Stats.CurrentSp = 0
	}
}

func (p *Player) ToTarget() *common.Target {
	return &common.Target{
		Name:       p.Name,
		TargetType: p.Class.name,
		Id:         p.Id,
		Action:     p.QueuedAction,
		CurrentHp:  p.Stats.CurrentHp,
		CurrentSp:  p.Stats.CurrentSp,
		MaxHp:      p.Stats.MaxHp,
		MaxSp:      p.Stats.MaxSp,
		IsPlayer:   p.IsPlayer,
	}
}

func NewPlayer(conn net.Conn, name string) Player {
	return Player{Id: rand.Int(), // TODO Check for collisions,
		Stats: common.Stats{
			CurrentHp: 10,
			MaxHp:     10,
			Hit:       1,
			Attack:    1,
		},
		IsPlayer: true,
		Conn:     conn,
		Name:     name,
		Class:    Novice,
		Location: 0,
		Level:    1,
	}
}
