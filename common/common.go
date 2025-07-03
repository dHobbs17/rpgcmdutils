package common

type ContentMessage struct {
	Action string
	Data   NpcPlayer
	Args   []string
}

type NpcPlayer interface {
	GetTarget() *NpcPlayer
	Get() *NpcPlayer
	GetLevel() int
	GetId() int
	GetStats() Stats
	getAction() Action
	GetSkills() Skills
	GetCurrentHp() int
	GetMaxHp() int
	GetCurrentSp() int
	IsPlayer() bool
	GetMaxSp() int
}

func GetTarget(np NpcPlayer) *NpcPlayer {
	return np.GetTarget()
}
func Get(np NpcPlayer) *NpcPlayer {
	return &np
}
func GetType(np NpcPlayer) *NpcPlayer {
	return &np
}
func GetId(np NpcPlayer) int {
	return np.GetId()
}
func GetLevel(np NpcPlayer) int {
	return np.GetLevel()
}
func IsPlayer(np NpcPlayer) bool {
	return np.IsPlayer()
}
func GetAction(np NpcPlayer) Action {
	return np.getAction()
}
func GetSkills(np NpcPlayer) Skills {
	return np.GetSkills()
}
func GetStats(np NpcPlayer) Stats {
	return np.GetStats()
}
func GetCurrentHp(np NpcPlayer) int {
	return np.GetCurrentHp()
}
func GetCurrentSp(np NpcPlayer) int {
	return np.GetCurrentSp()
}
func GetMaxHp(np NpcPlayer) int {
	return np.GetMaxHp()
}
func GetMaxSp(np NpcPlayer) int {
	return np.GetMaxSp()
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
