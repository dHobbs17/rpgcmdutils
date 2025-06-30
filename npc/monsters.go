package npc

import "fmt"

type monsterId int
type monsterTypeId int

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
		return Rat, MonsterError{nil}
	}
	return Npc{}, MonsterError{fmt.Errorf("not a valid monster Id")}
}

// Errors
type MonsterError struct{ Err error }

func (e MonsterError) Error() string { return e.Err.Error() }

// MOB IMPLEMENTATIONS TODO implement classes and inheritence
var Rat = Npc{
	name:    RAT_NAME,
	id:      RAT,
	npcType: ANIMAL_NAME,
	passive: true,
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
	actions: []npcAction{ATTACK, FLEE, HIDE, WAIT, MOVE},
}
