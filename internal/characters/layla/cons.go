package layla

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

const c4Key = "layla-c4"

// The Shield Absorption of the Curtain of Slumber generated by Nights of Formal Focus is increased by 20%.
// Additionally, when unleashing Nights of Formal Focus, she will generate a shield for any nearby party members
// who are not being protected by a Curtain of Slumber. This shield will have 35% of the absorption of a Curtain of Slumber,
// will last for 12s, and will absorb Cryo DMG with 250% effectiveness.
func (c *char) c1() {
	c.Core.Player.Shields.AddShieldBonusMod("layla-c1", -1, func() (float64, bool) {
		return 0.2, true
	})

}

// When Nights of Formal Focus starts to fire off Shooting Stars, it will grant all nearby party members the Dawn Star effect,
// causing their Normal and Charged Attack DMG to increase based on 5% of Layla's Max HP.
// Dawn Star can last up to 3s and will be removed 0.05s after dealing Normal or Charged Attack DMG.
func (c *char) c4() {
	c.Core.Events.Subscribe(event.OnEnemyHit, func(args ...interface{}) bool {
		ae := args[1].(*combat.AttackEvent)
		if ae.Info.AttackTag != combat.AttackTagNormal && ae.Info.AttackTag != combat.AttackTagExtra {
			return false
		}

		char := c.Core.Player.ByIndex(ae.Info.ActorIndex)
		if !char.StatusIsActive(c4Key) {
			return false
		}

		dmgAdded := 0.05 * c.MaxHP()
		ae.Info.FlatDmg += dmgAdded

		c.QueueCharTask(func() { char.DeleteStatus(c4Key) }, 0.05*60)

		c.Core.Log.NewEvent("layla c4 adding damage", glog.LogPreDamageMod, ae.Info.ActorIndex).
			Write("damage_added", dmgAdded)

		return false
	}, "layla-c4")
}

// Shooting Stars from Nights of Formal Focus deal 40% increased DMG, and Starlight Slugs from Dream of the Star-Stream Shaker deal 40% increased DMG.
// Additionally, the interval between the creation of Night Stars via Nights of Formal Focus is decreased by 20%.
func (c *char) c6() {
	m := make([]float64, attributes.EndStatType)
	m[attributes.DmgP] = 0.4

	c.AddAttackMod(character.AttackMod{
		Base: modifier.NewBase("layla-c6", -1),
		Amount: func(atk *combat.AttackEvent, t combat.Target) ([]float64, bool) {
			if atk.Info.AttackTag != combat.AttackTagElementalBurst && atk.Info.Abil != "Shooting Star" {
				return nil, false
			}
			return m, true
		},
	})
}
