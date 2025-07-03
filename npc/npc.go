package npc

import (
	"github.com/dHobbs17/rpgcmdutils/common"
	"math/rand/v2"
)

type npcState int
type npcAction int

type Npc struct {
	Name         string
	NpcType      string
	IsPlayer     bool
	MonsterId    monsterId
	Level        int
	Id           int
	Stats        common.Stats
	QueuedAction *common.Action
	Abilities    []string
	Spells       []string
	Skills       common.Skills
	Location     string
	Dialog       npcDialogs
	State        npcState
	Target       *common.Target
	Actions      []npcAction
	Passive      bool
	DefaultState npcState
	Lootable     bool
	Dead         bool
	InCombat     bool
	PossibleLoot []string
	Loot         []string
}

// STATES
const (
	ATTACKING npcState = iota
	FOLLOWING
	MOVING
	DEFENDING
	HIDING
	STALKING
	FLEEING
	WAITING
	DYING
)

const (
	ATTACKING_STATE = "attacking"
	DEFENDING_STATE = "defending"
	MOVING_STATE    = "moving"
	FOLLOWING_STATE = "following"
	HIDING_STATE    = "hiding"
	STALKING_STATE  = "stalking"
	FLEEING_STATE   = "fleeing"
	WAITING_STATE   = "waiting"
	DYING_STATE     = "dying"
)

var npcStates = map[npcState]string{
	ATTACKING: ATTACKING_STATE,
	DEFENDING: DEFENDING_STATE,
	HIDING:    HIDING_STATE,
	MOVING:    MOVING_STATE,
	FOLLOWING: FOLLOWING_STATE,
	STALKING:  STALKING_STATE,
	FLEEING:   FLEEING_STATE,
	WAITING:   WAITING_STATE,
	DYING:     DYING_STATE,
}

// ACTIONS
const (
	ATTACK npcAction = iota
	FLEE
	DEFEND
	HIDE
	STALK
	WAIT
	FOLLOW
	MOVE
)

const (
	ATTACK_ACTION = "attack"
	DEFEND_ACTION = "defend"
	HIDE_ACTION   = "hide"
	STALK_ACTION  = "stalk"
	FOLLOW_ACTION = "follow"
	WAIT_ACTION   = "wait"
	FLEE_ACTION   = "flee"
	MOVE_ACTION   = "move"
)

var NpcStates = map[npcAction]string{
	ATTACK: ATTACK_ACTION,
	DEFEND: DEFEND_ACTION,
	HIDE:   HIDE_ACTION,
	STALK:  STALK_ACTION,
	FOLLOW: FOLLOW_ACTION,
	WAIT:   WAIT_ACTION,
	FLEE:   FLEE_ACTION,
	MOVE:   MOVE_ACTION,
}

// getters and setters
func (n *Npc) Get() Npc                  { return *n }
func (n *Npc) GetName() string           { return n.Name }
func (n *Npc) GetPossibleLoot() []string { return n.PossibleLoot }
func (n *Npc) GetNpcType() string        { return n.NpcType }
func (n *Npc) GetLevel() int             { return n.Level }
func (n *Npc) IsLootable() bool          { return n.Lootable }
func (n *Npc) IsAlive() bool             { return !n.Dead }
func (n *Npc) IsPassive() bool           { return n.Passive }
func (n *Npc) IsInCombat() bool          { return n.InCombat }
func (n *Npc) SetCombat(c bool)          { n.InCombat = c }

func (n *Npc) GetQueuedAction() *common.Action  { return n.QueuedAction }
func (n *Npc) SetQueuedAction(a *common.Action) { n.QueuedAction = a }
func (n *Npc) ClearQueuedAction()               { n.QueuedAction = nil }

func (n *Npc) GetState() npcState  { return n.State }
func (n *Npc) SetState(s npcState) { n.State = s }
func (n *Npc) ResetState()         { n.State = n.DefaultState }

func (n *Npc) GetHp() int    { return n.Stats.CurrentHp }
func (n *Npc) GetMaxHp() int { return n.Stats.MaxHp }

func (n *Npc) AdjustHp(hp int) {
	n.Stats.CurrentHp += hp
	if n.Stats.CurrentHp <= 0 {
		n.Stats.CurrentHp = 0
		n.Dead = true
		n.Lootable = true
	}
}

func (n *Npc) ResetHp() { n.Stats.CurrentHp = n.Stats.MaxHp }

func (n *Npc) GetMonsterId() monsterId { return n.MonsterId }

func (n *Npc) GetId() int { return n.Id }

func (n *Npc) GetTarget() *common.Target       { return n.Target }
func (n *Npc) SetTarget(target *common.Target) { n.Target = target }
func (n *Npc) ResetTarget()                    { n.Target = nil }

func (n *Npc) GetSp() int    { return n.Stats.CurrentSp }
func (n *Npc) SetSp(sp int)  { n.Stats.CurrentSp = sp }
func (n *Npc) ResetSp()      { n.Stats.CurrentSp = n.Stats.MaxSp }
func (n *Npc) GetMaxSp() int { return n.Stats.MaxSp }
func (n *Npc) AdjustSp(sp int) {
	n.Stats.CurrentSp += sp
	if n.Stats.CurrentSp <= 0 {
		n.Stats.CurrentSp = 0
	}
}

func (n *Npc) ToTarget() *common.Target {
	return &common.Target{
		Name:       n.Name,
		TargetType: n.NpcType,
		Id:         n.Id,
		Action:     n.QueuedAction,
		CurrentHp:  n.Stats.CurrentHp,
		CurrentSp:  n.Stats.CurrentSp,
		MaxHp:      n.Stats.MaxHp,
		MaxSp:      n.Stats.MaxSp,
		IsPlayer:   n.IsPlayer,
	}
}

func (n *Npc) GetDialogGreeting() string { return getDialog(DIALOG_GREET, n) }
func (n *Npc) GetDialogDeath() string    { return getDialog(DIALOG_DEATH, n) }
func (n *Npc) GetDialogAttack() string   { return getDialog(DIALOG_ATTACK, n) }
func (n *Npc) GetDialogDamage() string   { return getDialog(DIALOG_DAMAGE, n) }
func (n *Npc) GetDialogWeak() string     { return getDialog(DIALOG_WEAK, n) }
func (n *Npc) GetDialogRun() string      { return getDialog(DIALOG_RUN, n) }
func (n *Npc) CalcHit() int              { return n.Stats.Hit }
func (n *Npc) CalcDamage() int           { return n.Stats.Attack }
func (n *Npc) GetLoot() []string         { return n.Loot }

func (n *Npc) generateLoot() {
	n.Loot = append(n.Loot, n.PossibleLoot[rand.IntN(len(n.PossibleLoot)-1)])
}

func getDialog(dia monsterDialog, n *Npc) string {
	switch dia {
	case DIALOG_GREET:
		return n.Dialog.GREETING[rand.IntN(len(n.Dialog.GREETING)-1)]
	case DIALOG_DEATH:
		return n.Dialog.DEATH[rand.IntN(len(n.Dialog.DEATH)-1)]
	case DIALOG_ATTACK:
		return n.Dialog.ATTACK[rand.IntN(len(n.Dialog.ATTACK)-1)]
	case DIALOG_DAMAGE:
		return n.Dialog.DAMAGE[rand.IntN(len(n.Dialog.DAMAGE)-1)]
	case DIALOG_WEAK:
		return n.Dialog.WEAK[rand.IntN(len(n.Dialog.WEAK)-1)]
	case DIALOG_RUN:
		return n.Dialog.RUN[rand.IntN(len(n.Dialog.RUN)-1)]
	default:
		return n.Dialog.GREETING[rand.IntN(len(n.Dialog.GREETING)-1)]
	}
}
