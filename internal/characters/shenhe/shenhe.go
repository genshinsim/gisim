package shenhe

import (
	"github.com/genshinsim/gcsim/pkg/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/keys"
)

func init() {
	core.RegisterCharFunc(keys.Shenhe, NewChar)
}

type char struct {
	*character.Tmpl
	quillcount   []int
	c4count      int
	c4expiry     int
	eNextRecover int
	eTickSrc     int
	eChargeMax   int
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
	c.NormalHitNum = 5
	c.BurstCon = 5
	c.SkillCon = 3
	c.CharZone = core.ZoneLiyue
	c.c4count = 0
	c.c4expiry = 0
	c.eChargeMax = 1
	if c.Base.Cons >= 1 {
		c.eChargeMax = 2
	}
	if c.Base.Cons >= 4 {
		c.c4()
	}

	return &c, nil
}

func (c *char) Init(index int) {
	c.Tmpl.Init(index)
	// if c.Base.Cons >= 6 {
	// 	c.c6()
	// }
	c.a2()
	c.quillcount = make([]int, len(c.Core.Chars))
}

func (c *char) ActionFrames(a core.ActionType, p map[string]int) (int, int) {
	switch a {
	case core.ActionAttack:
		f := 0
		switch c.NormalCounter {
		case 0:
			f = 12
		case 1:
			f = 38 - 12
		case 2:
			f = 72 - 38
		case 3:
			f = 141 - 72
		case 4:
			f = 167 - 141
		}
		f = int(float64(f) / (1 + c.Stats[core.AtkSpd]))
		return f, f
	case core.ActionSkill:
		return 26, 26
	case core.ActionBurst:
		return 99, 99
	case core.ActionCharge:
		return 78, 78
	default:
		c.Core.Log.Warnf("%v: unknown action (%v), frames invalid", c.Base.Key.String(), a)
		return 0, 0
	}
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

// inspired from barbara c2
// TODO: technically always assumes you are inside shenhe burst
func (c *char) a2() {
	val := make([]float64, core.EndStatType)
	val[core.CryoP] = 0.15
	if c.Base.Cons >= 2 {
		val[core.CryoP] += 0.15
	}
	for i, char := range c.Core.Chars {
		if i == c.Index {
			continue
		}
		char.AddMod(core.CharStatMod{
			Key:    "shenhe-a2",
			Expiry: -1,
			Amount: func(a core.AttackTag) ([]float64, bool) {
				if c.Core.Status.Duration("shenheburst") >= 0 {
					return val, true
				} else {
					return nil, false
				}
			},
		})
	}
}

func (c *char) c4() {
	c.AddPreDamageMod(core.PreDamageMod{
		Expiry: -1,
		Key:    "shenhe-c4",
		Amount: func(atk *core.AttackEvent, t core.Target) ([]float64, bool) {
			val := make([]float64, core.EndStatType)

			if atk.Info.AttackTag != core.AttackTagElementalArt && atk.Info.AttackTag != core.AttackTagElementalArtHold {
				return nil, false
			}
			if c.Core.F >= c.c4expiry {
				return nil, false
			}
			val[core.DmgP] += 0.05 * float64(c.c4count)
			c.c4count = 0
			c.c4expiry = 0
			return val, true
		},
	})
}
