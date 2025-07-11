package player

type Class struct {
	name            string
	id              int
	skillModifier   skillModifier
	statModifier    statModifier
	spellModifier   spellModifier
	abilityModifier abilityModifier
	jobLevel        int
	jobXp           int
	abilities       []string
	jobPoints       int
}

type skillModifier struct {
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

type statModifier struct {
	hp           int
	sp           int
	morale       int
	attack       int
	dodge        int
	parry        int
	block        int
	intelligence int
	strength     int
	dexterity    int
}

type abilityModifier struct {
}
type spellModifier struct {
}

// CLASS IMPLEMENTATIONS -- TODO implement these via JSON
var Novice = Class{
	name:     "Peasant",
	jobLevel: 1,
}

var Knight = Class{
	name: "Knight",
	skillModifier: skillModifier{
		swords:   1,
		axes:     1,
		maces:    1,
		block:    1,
		survival: 1,
	},
	statModifier: statModifier{
		hp:       10,
		morale:   1,
		block:    1,
		parry:    1,
		strength: 1,
	}}
