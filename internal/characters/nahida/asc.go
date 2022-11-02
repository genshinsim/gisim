package nahida

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

const a4BuffKey = "nahida-a4"

// When unleashing Illusory Heart, the Shrine of Maya will gain the following effects:
// The Elemental Mastery of the active character within the field will be increased by 20% of the Elemental Mastery of the party member with the highest Elemental Mastery.
// You can gain a maximum of 200 Elemental Mastery in this manner.
func (c *char) a1(dur int) {
	var max float64
	team := c.Core.Player.Chars()
	for _, char := range team {
		em := char.Stat(attributes.EM)
		if em > max {
			max = em
		}
	}
	max = 0.25 * max

	if max > 250 {
		max = 250
	}

	c.a4Buff[attributes.EM] = max

	for i, char := range team {
		idx := i
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBase(a4BuffKey, dur),
			AffectedStat: attributes.EM,
			Amount: func() ([]float64, bool) {
				return c.a4Buff, c.Core.Player.Active() == idx
			},
		})
	}

}

//Each point of Nahida's Elemental Mastery beyond 200 will grant 0.1% Bonus DMG and 0.03% CRIT Rate to Tri-Karma Purification from All Schemes to Know.
//A maximum of 80% Bonus DMG and 24% CRIT Rate can be granted to Tri-Karma Purification in this manner.

func (c *char) a4(em float64) (dmgBuff, crBuff float64) {
	//TODO: this way of buffing doesn't show up in mods; consider changing this to mods somehow??
	em = em - 200
	dmgBuff = em * 0.001
	if dmgBuff > 0.8 {
		dmgBuff = 0.8
	}
	crBuff = em * 0.0003
	if crBuff > .24 {
		crBuff = .24
	}

	c.Core.Combat.Log.NewEvent("tri-karma-a4-buff", glog.LogCharacterEvent, c.Index).
		Write("em_in_excess_of_200", em).
		Write("cr_buff", crBuff).
		Write("dmg_buff", dmgBuff)

	return

}
