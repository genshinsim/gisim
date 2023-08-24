package lyney

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/player"
	"github.com/genshinsim/gcsim/pkg/core/targets"
)

var skillFrames []int

const (
	skillHitmark = 18
	skillCD      = 15 * 60
	skillCDStart = 15

	particleICDKey = "lyney-particle-icd"
	particleICD    = 0.3 * 60
	particleCount  = 5
)

func init() {
	skillFrames = frames.InitAbilSlice(43) // E -> D
	skillFrames[action.ActionAttack] = 42
	skillFrames[action.ActionAim] = 42
	skillFrames[action.ActionSkill] = 42
	skillFrames[action.ActionBurst] = 42
	skillFrames[action.ActionJump] = 42
	skillFrames[action.ActionWalk] = 42
	skillFrames[action.ActionSwap] = 41
}

func (c *char) Skill(p map[string]int) action.ActionInfo {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Bewildering Lights",
		AttackTag:  attacks.AttackTagElementalArt,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 25,
		Mult:       skill[c.TalentLvlSkill()] + skillBonus[c.TalentLvlSkill()]*float64(c.propSurplusStacks),
	}
	c.Core.Player.Heal(player.HealInfo{
		Caller:  c.Index,
		Target:  c.Core.Player.Active(),
		Message: "Bewildering Lights",
		Src:     0.2 * c.MaxHP() * float64(c.propSurplusStacks),
		Bonus:   c.Stat(attributes.Heal),
	})
	c.propSurplusStacks = 0
	c.Core.Log.NewEvent("Lyney Prop Surplus stacks removed", glog.LogCharacterEvent, c.Index).Write("prop_surplus_stacks", c.propSurplusStacks)

	player := c.Core.Combat.Player()
	skillPos := combat.NewCircleHitOnTarget(geometry.CalcOffsetPoint(player.Pos(), geometry.Point{Y: 5.5}, player.Direction()), nil, 5.5)
	c.QueueCharTask(func() {
		c.Core.QueueAttack(ai, skillPos, 0, 0, c.particleCB)
		hatCount := len(c.hats)
		for i := 0; i < hatCount; i++ {
			c.hats[0].skillExplode()
		}
	}, skillHitmark)

	c.SetCDWithDelay(action.ActionSkill, skillCD, skillCDStart)

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillFrames[action.ActionSwap],
		State:           action.SkillState,
	}
}

func (c *char) particleCB(a combat.AttackCB) {
	if a.Target.Type() != targets.TargettableEnemy {
		return
	}
	if c.StatusIsActive(particleICDKey) {
		return
	}
	c.AddStatus(particleICDKey, particleICD, true)

	c.Core.QueueParticle(c.Base.Key.String(), particleCount, attributes.Pyro, c.ParticleDelay)
}
