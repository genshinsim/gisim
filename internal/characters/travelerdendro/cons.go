package travelerdendro

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func (c *char) c1cb() func(a combat.AttackCB) {
	return func(a combat.AttackCB) {
		if c.skillC1 {
			c.AddEnergy("dmc-c1", 3.5)
			c.skillC1 = false
		}

	}
}

func (c *char) c4() {
	c.burstOverflowingLotuslight += 5
	if c.burstOverflowingLotuslight > 10 {
		c.burstOverflowingLotuslight = 10
	}
	c.Core.Log.NewEvent("dmc-c4", glog.LogCharacterEvent, c.Index)
}

// Gets removed on swap - from Kolbiri
func (c *char) c6Init() {
	c.Core.Events.Subscribe(event.OnCharacterSwap, func(args ...interface{}) bool {
		prev := args[0].(int)
		prevChar := c.Core.Player.ByIndex(prev)
		prevChar.DeleteStatMod("dmc-c6")
		return false
	}, "travelerdendro-c6-remove")
}

func (c *char) c6Buff(delay int) {
	m := make([]float64, attributes.EndStatType)
	// A1/C6 buff ticks every 0.3s and applies for 1s. probably counting from gadget spawn - from Kolbiri
	c.Core.Tasks.Add(func() {
		if c.Core.F <= c.burstExpire { //burst isn't expired
			active := c.Core.Player.ActiveChar()
			m[attributes.DendroP] = 0.12
			if c.burstTransfig != attributes.NoElement {
				m[c.burstTransfig] = 0.12
			}
			active.AddStatMod(character.StatMod{
				Base: modifier.NewBase("dmc-c6", 60),
				Amount: func() ([]float64, bool) {
					return m, true
				},
			})
		}
	}, delay)
}
