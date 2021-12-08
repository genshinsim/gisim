package nagamasa

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
)

func init() {
	core.RegisterWeaponFunc("katsuragikiri nagamasa", weapon)
	core.RegisterWeaponFunc("katsuragikirinagamasa", weapon)
}

// Increases Elemental Skill DMG by 6~12%. After Elemental Skill hits an opponent, the character loses 3 Energy but regenerates 3~5 Energy every 2s for the next 6s. This effect can occur once every 10s. Can be triggered even when the character is not on the field.
// Same effect as kitain - code largely copied
func weapon(char core.Character, c *core.Core, r int, param map[string]int) {
	var m [core.EndStatType]float64
	base := 0.045 + float64(r)*0.015
	regen := 2.5 + float64(r)*0.5

	m[core.DmgP] = base

	char.AddMod(core.CharStatMod{
		Expiry: -1,
		Key:    "nagamasa-skill-dmg-buff",
		Amount: func(a core.AttackTag) ([core.EndStatType]float64, bool) {
			if a == core.AttackTagElementalArt || a == core.AttackTagElementalArtHold {
				return m, true
			}
			return nil, false
		},
	})

	icd := 0
	c.Events.Subscribe(core.OnDamage, func(args ...interface{}) bool {
		atk := args[1].(*core.AttackEvent)
		if atk.Info.ActorIndex != char.CharIndex() {
			return false
		}
		if atk.Info.AttackTag != core.AttackTagElementalArt {
			return false
		}
		if icd > c.F {
			return false
		}
		icd = c.F + 600 //once every 10 seconds
		char.AddEnergy(-3)
		for i := 120; i <= 360; i += 120 {
			char.AddTask(func() {
				char.AddEnergy(regen)
			}, "nagamasa-restore", i)
		}
		return false
	}, fmt.Sprintf("nagamasa-%v", char.Name()))
}
