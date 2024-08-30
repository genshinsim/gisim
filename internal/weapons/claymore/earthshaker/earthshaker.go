package earthshaker

import (
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterWeaponFunc(keys.EarthShaker, NewWeapon)
}

type Weapon struct {
	Index int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

// After a party member triggers a Pyro-related reaction, 
// the equipping character's Elemental Skill DMG is increased by 32% for 8s. 
// This effect can be triggered even when the triggering party member is not on the field.
func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := p.Refine

	amt := 0.12 + float64(r)*0.04

	buffSkill := func(args ...interface{}) bool {
		m := make([]float64, attributes.EndStatType)
		m[attributes.DmgP] = amt
		char.AddAttackMod(character.AttackMod{
			Base: modifier.NewBase("earth-shaker", 8*60),
			Amount: func(atk *combat.AttackEvent, t combat.Target) ([]float64, bool) {
				if atk.Info.AttackTag != attacks.AttackTagElementalArt && atk.Info.AttackTag != attacks.AttackTagElementalArtHold {
					return nil, false
				}
				return m, true
			},
		})
		return false
	}

	c.Events.Subscribe(event.OnOverload, buffSkill, "earth-shaker-overload")
	c.Events.Subscribe(event.OnVaporize, buffSkill, "earth-shaker-vaporize")
	c.Events.Subscribe(event.OnMelt, buffSkill, "earth-shaker-melt")
	c.Events.Subscribe(event.OnSwirlPyro, buffSkill, "earth-shaker-pyro-swirl")
	c.Events.Subscribe(event.OnCrystallizePyro, buffSkill, "earth-shaker-pyro-crystallize")
	c.Events.Subscribe(event.OnBurning, buffSkill, "earth-shaker-burning")
	c.Events.Subscribe(event.OnBurgeon, buffSkill, "earth-shaker-burgeon")

	return w, nil
}