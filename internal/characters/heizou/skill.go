package heizou

import "github.com/genshinsim/gcsim/pkg/core"

func (c *char) skillHoldDuration(stacks int) int {
	//animation duration only
	//diff is the number of stacks we must charge up to reach the desired state
	diff := stacks - c.decStack
	if diff < 0 {
		diff = c.decStack
	}
	if diff > 4 {
		diff = 4
	}
	//it's .75s per stack
	return 45 * diff
}

const skillChargeStart = 12

func (c *char) Skill(p map[string]int) (int, int) {
	f, a := c.ActionFrames(core.ActionSkill, p)
	stacks := p["hold"]
	dur := c.skillHoldDuration(stacks) //this should max out to 3s

	//queue task to increase stacks every 0.75s up to dur
	for i := 45; i <= dur; i++ {
		c.Core.Tasks.Add(func() {
			c.decStack++
		}, skillChargeStart+i)
	}

	stackToBeAdded := p["hold"] //guarantee that you dont surpass more than 4 stacks with holding
	if stackToBeAdded+c.decStack > 4 {
		diff := 4 - c.decStack
		stackToBeAdded = diff
	}
	//TODO: Verify attack frame

	//queue the attack as a task that goes through at the end of the animation; check for stacks then
	//animation should be skillChargeStart + dur + attack animation length
	c.Core.Tasks.Add(func() {
		ai := core.AttackInfo{
			ActorIndex: c.Index,
			Abil:       "Heartstopper Strike",
			AttackTag:  core.AttackTagElementalArt,
			ICDTag:     core.ICDTagNone,
			ICDGroup:   core.ICDGroupDefault,
			StrikeType: core.StrikeTypeDefault,
			Element:    core.Anemo,
			Durability: 25,
			Mult:       skill[c.TalentLvlSkill()] + float64(c.decStack)*decBonus[c.TalentLvlSkill()],
		}
		if c.decStack == 4 {
			ai.Mult += convicBonus[c.TalentLvlSkill()]
		}
		//a4 + energy
		done := false
		cb := func(a core.AttackCB) {
			if done {
				return
			}
			done = true
			count := 2
			switch c.decStack {
			case 2, 3:
				if c.Core.Rand.Float64() < .5 {
					count++
				}
			case 4:
				count++
			}
			c.QueueParticle("heizou", count, core.Anemo, 150)
			c.AddTask(func() {
				c.decStack = 0
				c.Core.Log.NewEvent(
					"stack removed",
					core.LogCharacterEvent,
					c.Index,
				)
			}, "remove stack", 1)
			c.a4()
		}
		//generate snap
		snap := c.Snapshot(&ai)
		//check for c6, increase crit
		if c.Base.Cons >= 6 {
			c.c6(&snap)
		}
		c.Core.Combat.QueueAttackWithSnap(ai, snap, core.NewDefCircHit(3, false, core.TargettableEnemy), 0, cb)
	}, skillChargeStart+dur+f)

	c.SetCD(core.ActionSkill, eCD)

	return skillChargeStart + dur + f, a
}
