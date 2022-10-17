package travelerdendro

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
)

var burstFrames [][]int

const burstHitmark = 91

func init() {
	burstFrames = make([][]int, 2)

	// Male
	burstFrames[0] = frames.InitAbilSlice(58)
	burstFrames[0][action.ActionSwap] = 57 // Q -> Swap

	// Female
	burstFrames[1] = frames.InitAbilSlice(58)
	burstFrames[1][action.ActionSwap] = 57 // Q -> Swap
}

func (c *char) Burst(p map[string]int) action.ActionInfo {

	c.SetCD(action.ActionBurst, 1200)
	c.ConsumeEnergy(2)

	burstDur := 12 * 60
	if c.Base.Cons >= 2 {
		burstDur += 3 * 60
	}

	s := c.newLeaLotusLamp(burstDur)

	c.Core.Tasks.Add(func() {
		c.Core.Combat.AddGadget(s)
	}, burstHitmark)

	c.burstOverflowingLotuslight = 0
	// A1 adds a stack per second
	for delay := c.Core.F + burstHitmark; delay < burstDur; delay += 60 {
		c.a1Stack(delay)
	}

	// A1/C6 buff ticks every 0.3s and applies for 1s. probably counting from gadget spawn - Kolbiri
	for delay := c.Core.F + burstHitmark; delay < burstDur; delay += 0.3 * 60 {
		c.a1Buff(delay)
	}

	if c.Base.Cons >= 6 {
		for delay := c.Core.F + burstHitmark; delay < burstDur; delay += 0.3 * 60 {
			c.c6Buff(delay)
		}
	}

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(burstFrames[c.gender]),
		AnimationLength: burstFrames[c.gender][action.InvalidAction],
		CanQueueAfter:   burstFrames[c.gender][action.ActionSwap], // earliest cancel
		State:           action.BurstState,
	}
}
