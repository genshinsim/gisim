package verdict

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterWeaponFunc(keys.Verdict, NewWeapon)
}

type Weapon struct {
	Index  int
	stacks int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

const (
	buffKey      = "verdict-skill-dmg"
	dmgWindowKey = "verdict-dmg-window"
)

// Increases ATK by 20/25/30/35/40%. When party members obtain Elemental Shards from Crystallize reactions,
// the equipping character will gain 1 Seal, increasing Elemental Skill DMG by 18/22.5/27/31.5/36%.
// The Seal lasts for 15s, and the equipper may have up to 2 Seals at once.
// All of the equipper's Seals will disappear 0.2s after their Elemental Skill deals DMG.
func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := p.Refine

	// perm buff
	m := make([]float64, attributes.EndStatType)
	m[attributes.ATKP] = 0.15 + float64(r)*0.05
	char.AddStatMod(character.StatMod{
		Base:         modifier.NewBase("verdict-atk", -1),
		AffectedStat: attributes.ATKP,
		Amount: func() ([]float64, bool) {
			return m, true
		},
	})

	gainSealFunc := func(args ...interface{}) bool {
		if !char.StatModIsActive(buffKey) {
			w.stacks = 0
		}
		if w.stacks < 2 {
			w.stacks++
		}
		char.AddStatus(buffKey, 15*60, true)
		return false
	}
	skillDmg := 0.135 + float64(r)*0.045
	addSkillDmgFunc := func(args ...interface{}) bool {
		atk := args[1].(*combat.AttackEvent)
		if atk.Info.ActorIndex != char.Index {
			return false
		}
		if atk.Info.AttackTag != attacks.AttackTagElementalArt && atk.Info.AttackTag != attacks.AttackTagElementalArtHold {
			return false
		}
		// don't do anything if not in buff
		if !char.StatusIsActive(buffKey) {
			return false
		}
		// otherwise if this is first time proccing
		// - set duration for dmg window
		// - reset buff once window is over
		if !char.StatusIsActive(dmgWindowKey) {
			char.AddStatus(dmgWindowKey, 0.2*60, true)
			char.QueueCharTask(func() {
				char.DeleteStatus(buffKey)
				w.stacks = 0
			}, 0.2*60)
		}
		skillDmgAdd := skillDmg * float64(w.stacks)
		atk.Snapshot.Stats[attributes.DmgP] += skillDmgAdd

		c.Log.NewEvent("verdict adding skill dmg", glog.LogPreDamageMod, char.Index).
			Write("skill_dmg_added", skillDmgAdd)
		return false
	}

	c.Events.Subscribe(event.OnCrystallizeCryo, gainSealFunc, fmt.Sprintf("vertict-crystallizecryo-%v", char.Base.Key.String()))
	c.Events.Subscribe(event.OnCrystallizeElectro, gainSealFunc, fmt.Sprintf("vertict-crystallizeelectro-%v", char.Base.Key.String()))
	c.Events.Subscribe(event.OnCrystallizeHydro, gainSealFunc, fmt.Sprintf("vertict-crystallizehydro-%v", char.Base.Key.String()))
	c.Events.Subscribe(event.OnCrystallizePyro, gainSealFunc, fmt.Sprintf("vertict-crystallizepyro-%v", char.Base.Key.String()))

	c.Events.Subscribe(event.OnEnemyHit, addSkillDmgFunc, fmt.Sprintf("verdict-onhit-%v", char.Base.Key.String()))

	return w, nil
}
