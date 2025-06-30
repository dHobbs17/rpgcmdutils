package npc

import (
	"fmt"
	"math/rand"
)

type monsterId int
type monsterTypeId int

type monsterDialog int

type npcDialogs struct {
	GREETING []string
	DEATH    []string
	DAMAGE   []string
	ATTACK   []string
	WEAK     []string
	RUN      []string
}

const (
	DIALOG_GREET monsterDialog = iota
	DIALOG_DEATH
	DIALOG_ATTACK
	DIALOG_DAMAGE
	DIALOG_WEAK
	DIALOG_RUN
)

const (
	RAT monsterId = iota
)

const (
	RAT_NAME = "rat"
)

var monsterMap = map[monsterId]string{
	RAT: RAT_NAME,
}

const (
	ANIMAL monsterTypeId = iota
)

const (
	ANIMAL_NAME = "animal"
)

var monsterTypeMap = map[monsterTypeId]string{
	ANIMAL: ANIMAL_NAME,
}

func CreateMonster(id monsterId) (Npc, MonsterError) {
	switch id {
	case RAT:
		var rat = Rat
		rat.generateLoot()
		return rat, MonsterError{nil}
	}
	return Npc{}, MonsterError{fmt.Errorf("not a valid monster Id")}
}

// Errors
type MonsterError struct{ Err error }

func (e MonsterError) Error() string { return e.Err.Error() }

// MOB IMPLEMENTATIONS -- TODO implement these via JSON
var Rat = Npc{
	name:       RAT_NAME,
	monsterId:  RAT,
	instanceId: rand.Int(), // TODO Check for collisions
	npcType:    ANIMAL_NAME,
	passive:    true,
	stats: Stats{
		maxHp:        8,
		morale:       2,
		attack:       2,
		dodge:        5,
		intelligence: 1,
		strength:     1,
		dexterity:    4,
	},
	skills: Skills{
		stealth:  4,
		survival: 4,
	},
	dialog: npcDialogs{
		GREETING: []string{"Rat lets out a piecing Screech", "Rat does a spin in anger"},
		DEATH:    []string{"Rat falls to the ground dead", "Rats head explodes"},
		WEAK:     []string{"Rat seems hurt", "Rats left arm has been severed"},
		RUN:      []string{"Rat attempts to flee", "The Rat has had enough of this"},
		DAMAGE:   []string{"Rat has been damaged", "Rat shrugs off your weak attack"},
		ATTACK:   []string{"Rat flails at you", "Rat jumps at you", "Rat attempts to bite you"},
	},
	possibleLoot: []string{"1 gold", "a half eaten apple", "a severed thumb"},
	actions:      []npcAction{ATTACK, FLEE, HIDE, WAIT, MOVE},
}
