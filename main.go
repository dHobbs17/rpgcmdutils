package rpgcmdutils

import (
	"github.com/dHobbs17/rpgcmdutils/common"
	"github.com/dHobbs17/rpgcmdutils/npc"
	"github.com/dHobbs17/rpgcmdutils/player"
)

type NpcPlayer[T npc.Npc | player.Player] struct {
	lootable bool
	dead     bool
	inCombat bool
	name     string
	level    int
	id       int
	alive    bool
	target   *int
	stats    common.Stats
	action   common.Action
	skills   common.Skills
}

func (np *NpcPlayer[T]) GetTarget() *int {
	return np.target
}

func (np *NpcPlayer[T]) GetId() int {
	return np.id
}

func (np *NpcPlayer[T]) GetLevel() int {
	return np.level
}

func (np *NpcPlayer[T]) GetStats() common.Stats {
	return np.stats
}

func (np *NpcPlayer[T]) GetAction() common.Action {
	return np.action
}

func (np *NpcPlayer[T]) GetSkills() common.Skills {
	return np.skills
}
