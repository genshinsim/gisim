package xilonen

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
)

func (c *char) Jump(p map[string]int) (action.Info, error) {
	if c.nightsoulState.HasBlessing() {
		c.c6()

		if c.Core.Player.LastAction.Type == action.ActionDash && !c.StatusIsActive(c6key) {
			c.reduceNightsoulPoints(15.0 * c.c1ValMod()) // total 20, 5 from dash, 15 from jump
		}
	}
	return c.Character.Jump(p)
}
