package common

type TargetContentMessage struct {
	Action string
	Data   Target
	Args   []string
}

type TargetMessage struct {
	Action string
	Data   string
	Args   []string
}

type Target struct {
	Name       string
	TargetType string
	Id         int
	Action     *Action
	CurrentHp  int
	MaxHp      int
	CurrentSp  int
	MaxSp      int
	IsPlayer   bool
}

type Stats struct {
	CurrentHp    int
	CurrentSp    int
	MaxHp        int
	MaxSp        int
	Morale       int
	Hit          int
	Attack       int
	Dodge        int
	Parry        int
	Block        int
	Intelligence int
	Strength     int
	Dexterity    int
	Reputation   int
}

type Action struct {
	Action string
	Data   string
	Args   []string
}

type Skills struct {
	Destruction int
	Conjuration int
	Illusion    int
	Perception  int
	Deception   int
	Stealth     int
	Swords      int
	Maces       int
	Axes        int
	Ranged      int
	Wands       int
	Block       int
	Survival    int
}
