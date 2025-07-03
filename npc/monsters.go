package npc

import (
	"fmt"
	"github.com/dHobbs17/rpgcmdutils/common"
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
	Name:      RAT_NAME,
	MonsterId: RAT,
	Id:        rand.Int(), // TODO Check for collisions
	NpcType:   ANIMAL_NAME,
	IsPlayer:  false,
	Passive:   true,
	Stats: common.Stats{
		MaxHp:        8,
		CurrentHp:    8,
		Morale:       2,
		Attack:       2,
		Dodge:        5,
		Intelligence: 1,
		Strength:     1,
		Dexterity:    4,
	},
	Skills: common.Skills{
		Stealth:  4,
		Survival: 4,
	},
	Dialog: npcDialogs{
		GREETING: []string{"Rat lets out a piecing Screech", "Rat does a spin in anger"},
		DEATH:    []string{"Rat falls to the ground dead", "Rats head explodes"},
		WEAK:     []string{"Rat seems hurt", "Rats left arm has been severed"},
		RUN:      []string{"Rat attempts to flee", "The Rat has had enough of this"},
		DAMAGE:   []string{"Rat has been damaged", "Rat shrugs off your weak attack"},
		ATTACK:   []string{"Rat flails at you", "Rat jumps at you", "Rat attempts to bite you"},
	},
	PossibleLoot: []string{"1 gold", "a half eaten apple", "a severed thumb"},
	Actions:      []npcAction{ATTACK, FLEE, HIDE, WAIT, MOVE},
}
