package wanderer

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/core/player/character/profile"
)

func init() {
	core.RegisterCharFunc(keys.Wanderer, NewChar)
}

type char struct {
	*tmpl.Character
	skydwellerPoints      int
	maxSkydwellerPoints   int
	a1Buffs               []attributes.Element
	a1AbsorbCheckLocation combat.AttackPattern
	a4Prob                float64
}

func NewChar(s *core.Core, w *character.CharWrapper, _ profile.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 60
	c.NormalHitNum = normalHitNum
	c.SkillCon = 5
	c.BurstCon = 3

	w.Character = &c

	return nil
}

func (c *char) Init() error {

	c.skydwellerPoints = 0
	c.maxSkydwellerPoints = 100
	c.a4Prob = 0.16

	c.a4Init()

	return nil
}

func (c *char) ActionStam(a action.Action, p map[string]int) float64 {
	if c.StatusIsActive(skillKey) {
		return 0
	}
	return c.Character.ActionStam(a, p)
}

func (c *char) Snapshot(ai *combat.AttackInfo) combat.Snapshot {
	ds := c.Character.Snapshot(ai)

	if ai.AttackTag == combat.AttackTagNormal && c.StatusIsActive(skillKey) {
		c.Core.Log.NewEvent("skill NA mult applied", glog.LogCharacterEvent, c.Index).
			Write("prev", ai.Mult).
			Write("next", skillNABonus[c.TalentLvlSkill()]*ai.Mult).
			Write("char", c.Index)
		ai.Mult = skillNABonus[c.TalentLvlSkill()] * ai.Mult
	}

	if ai.AttackTag == combat.AttackTagExtra && c.StatusIsActive(skillKey) {
		c.Core.Log.NewEvent("skill CA mult applied", glog.LogCharacterEvent, c.Index).
			Write("prev", ai.Mult).
			Write("next", skillCABonus[c.TalentLvlSkill()]*ai.Mult).
			Write("char", c.Index)
		ai.Mult = skillCABonus[c.TalentLvlSkill()] * ai.Mult
	}

	return ds
}

// Overwriting of remaining actions to account for falling state

func (c *char) Walk(p map[string]int) action.ActionInfo {
	delay := c.checkForSkillEnd()

	ai := c.Character.Walk(p)
	ai.Frames = func(action action.Action) int { return delay + ai.Frames(action) }
	ai.AnimationLength = delay + ai.AnimationLength
	ai.CanQueueAfter = delay + ai.CanQueueAfter

	return ai
}

func (c *char) Jump(p map[string]int) action.ActionInfo {
	delay := c.checkForSkillEnd()

	ai := c.Character.Jump(p)
	ai.Frames = func(action action.Action) int { return delay + ai.Frames(action) }
	ai.AnimationLength = delay + ai.AnimationLength
	ai.CanQueueAfter = delay + ai.CanQueueAfter

	return ai
}
