package thoma

import (
	"github.com/genshinsim/gcsim/pkg/character"
	"github.com/genshinsim/gcsim/pkg/core"
)

func init() {
	core.RegisterCharFunc(core.Thoma, NewChar)
}

type char struct {
	*character.Tmpl
	MaxShield float64
}

func NewChar(s *core.Core, p core.CharacterProfile) (core.Character, error) {
	c := char{}
	t, err := character.NewTemplateChar(s, p)
	if err != nil {
		return nil, err
	}
	c.Tmpl = t
	c.Energy = 80
	c.EnergyMax = 80
	c.Weapon.Class = core.WeaponClassSpear
	c.NormalHitNum = 4
	c.BurstCon = 5
	c.SkillCon = 3
	c.CharZone = core.ZoneInazuma
	c.Base.Element = core.Pyro
	c.MaxShield = (shieldppmax[c.TalentLvlSkill()]*c.MaxHP() + shieldflatmax[c.TalentLvlSkill()])

	return &c, nil
}

func (c *char) Init(index int) {
	c.Tmpl.Init(index)
}

func (c *char) ActionStam(a core.ActionType, p map[string]int) float64 {
	switch a {
	case core.ActionDash:
		return 18
	case core.ActionCharge:
		return 25
	default:
		c.Core.Log.Warnf("%v ActionStam for %v not implemented; Character stam usage may be incorrect", c.Base.Key.String(), a.String())
		return 0
	}

}
