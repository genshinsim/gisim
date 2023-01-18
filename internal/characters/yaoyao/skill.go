package yaoyao

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player"
)

var skillFrames []int

func init() {
	skillFrames = frames.InitAbilSlice(39)
	skillFrames[action.ActionDash] = 14
	skillFrames[action.ActionJump] = 14
	skillFrames[action.ActionSwap] = 38
}

func (c *char) Skill(p map[string]int) action.ActionInfo {

	// yuegui spawns at cd frame
	c.Core.Status.Add("yuegui", 500+13)

	procAI := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Radish (Skill)",
		AttackTag:  combat.AttackTagElementalArt,
		ICDTag:     combat.ICDTagElementalArt,
		ICDGroup:   combat.ICDGroupYaoyaoRadishSkill,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Dendro,
		Durability: 25,
		Mult:       skillRadishDMG[c.TalentLvlSkill()],
	}

	yuegui := c.newYueguiThrow(procAI)
	c.Core.Tasks.Add(func() {
		c.Core.Combat.AddGadget(yuegui)
	}, 13)

	c.SetCDWithDelay(action.ActionSkill, 15*60, 13)

	if c.Base.Cons >= 4 {
		c.c4()
	}

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillFrames[action.ActionDash], // earliest cancel
		State:           action.SkillState,
	}
}

func (c *char) getSkillHealInfo() player.HealInfo {
	heal := skillRadishHealing[0][c.TalentLvlBurst()]*c.MaxHP() + skillRadishHealing[1][c.TalentLvlBurst()]
	return player.HealInfo{
		Caller:  c.Index,
		Target:  c.Core.Player.Active(),
		Message: "Yuegui skill",
		Src:     heal,
	}
}
