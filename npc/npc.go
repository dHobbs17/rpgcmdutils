package npc

import "math/rand/v2"

type npcState int
type npcAction int

type Npc struct {
	name         string
	npcType      string
	monsterId    monsterId
	level        int
	instanceId   int
	stats        Stats
	queuedAction *NpcAction
	abilities    []string
	spells       []string
	skills       Skills
	location     string
	dialog       npcDialogs
	state        npcState
	target       *int
	actions      []npcAction
	passive      bool
	defaultState npcState
	lootable     bool
	dead         bool
	possibleLoot []string
	loot         []string
}

type NpcAction struct {
	Action string
	Data   string
	Args   []string
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
func (n *Npc) getName() string    { return n.name }
func (n *Npc) getNpcType() string { return n.npcType }
func (n *Npc) getLevel() int      { return n.level }
func (n *Npc) isLootable() bool   { return n.lootable }
func (n *Npc) isAlive() bool      { return !n.dead }

func (n *Npc) getQueuedAction() *NpcAction  { return n.queuedAction }
func (n *Npc) setQueuedAction(a *NpcAction) { n.queuedAction = a }
func (n *Npc) clearQueuedAction()           { n.queuedAction = nil }

func (n *Npc) getState() npcState  { return n.state }
func (n *Npc) setState(s npcState) { n.state = s }
func (n *Npc) resetState()         { n.state = n.defaultState }

func (n *Npc) getHp() int   { return n.stats.currentHp }
func (n *Npc) setHp(hp int) { n.stats.currentHp = hp }

func (n *Npc) adjustHp(hp int) {
	n.stats.currentHp += hp
	if n.stats.currentSp <= 0 {
		n.stats.currentHp = 0
		n.dead = true
		n.lootable = true
	}
}

func (n *Npc) resetHp() { n.stats.currentHp = n.stats.maxHp }

func (n *Npc) getId() monsterId { return n.monsterId }

func (n *Npc) getInstanceId() int { return n.instanceId }

func (n *Npc) getTarget() *int          { return n.target }
func (n *Npc) setTarget(targetsId *int) { n.target = targetsId }
func (n *Npc) resetTarget()             { n.target = nil }

func (n *Npc) getSp() int   { return n.stats.currentSp }
func (n *Npc) setSp(sp int) { n.stats.currentSp = sp }
func (n *Npc) resetSp()     { n.stats.currentSp = n.stats.maxSp }

func (n *Npc) getDialogGreeting() string { return getDialog(DIALOG_GREET, n) }
func (n *Npc) getDialogDeath() string    { return getDialog(DIALOG_DEATH, n) }
func (n *Npc) getDialogAttack() string   { return getDialog(DIALOG_ATTACK, n) }
func (n *Npc) getDialogDamage() string   { return getDialog(DIALOG_DAMAGE, n) }
func (n *Npc) getDialogWeak() string     { return getDialog(DIALOG_WEAK, n) }
func (n *Npc) getDialogRun() string      { return getDialog(DIALOG_RUN, n) }

func (n *Npc) getLoot() []string { return n.loot }
func (n *Npc) generateLoot()     { n.loot = append(n.loot, n.possibleLoot[rand.IntN(len(n.possibleLoot))]) }

func getDialog(dia monsterDialog, n *Npc) string {
	switch dia {
	case DIALOG_GREET:
		return n.dialog.GREETING[rand.IntN(len(n.dialog.GREETING))]
	case DIALOG_DEATH:
		return n.dialog.DEATH[rand.IntN(len(n.dialog.DEATH))]
	case DIALOG_ATTACK:
		return n.dialog.ATTACK[rand.IntN(len(n.dialog.ATTACK))]
	case DIALOG_DAMAGE:
		return n.dialog.DAMAGE[rand.IntN(len(n.dialog.DAMAGE))]
	case DIALOG_WEAK:
		return n.dialog.WEAK[rand.IntN(len(n.dialog.WEAK))]
	case DIALOG_RUN:
		return n.dialog.RUN[rand.IntN(len(n.dialog.RUN))]
	default:
		return n.dialog.GREETING[rand.IntN(len(n.dialog.GREETING))]
	}
}
