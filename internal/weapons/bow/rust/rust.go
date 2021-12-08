package generic

import (
	"github.com/genshinsim/gcsim/pkg/core"
)

func init() {
	core.RegisterWeaponFunc("rust", weapon)
}

func weapon(char core.Character, c *core.Core, r int, param map[string]int) {
	var m [core.EndStatType]float64
	inc := .3 + float64(r)*0.1
	char.AddMod(core.CharStatMod{
		Key: "rust",
		Amount: func(a core.AttackTag) ([core.EndStatType]float64, bool) {
			if a == core.AttackTagNormal {
				m[core.DmgP] = inc
				return m, true
			}
			if a == core.AttackTagExtra {
				m[core.DmgP] = -0.1
				return m, true
			}
			return m, false
		},
		Expiry: -1,
	})
}
