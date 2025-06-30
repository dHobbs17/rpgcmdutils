package npc

type npcState int
type npcAction int

type Npc struct {
	name      string
	npcType   string
	id        monsterId
	stats     Stats
	abilities []string
	spells    []string
	skills    Skills
	location  string
	dialog    []string
	state     int
	target    string
	actions   []npcAction
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
