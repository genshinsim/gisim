// Package combat handles all combat related functionalities including
//   - target tracking
//   - target selection
//   - hitbox collision checking
//   - attack queueing
package combat

import (
	"math/rand"

	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

type CharHandler interface {
	CombatByIndex(int) Character
	ApplyHitlag(char int, factor, dur float64)
}

type Character interface {
	ApplyAttackMods(a *AttackEvent, t Target) []interface{}
}

type Handler struct {
	Opt
	targets     []Target
	enemyIdxMap map[int]int
	enemies     []Target
	gadgets     []Target
	player      Target
	TotalDamage float64
}

type Opt struct {
	Events        event.Eventter
	Team          CharHandler
	Rand          *rand.Rand
	Debug         bool
	Log           glog.Logger
	DamageMode    bool
	DefHalt       bool
	EnableHitlag  bool
	DefaultTarget int //index for default target
}

func New(opt Opt) *Handler {
	h := &Handler{
		Opt:         opt,
		enemyIdxMap: make(map[int]int),
	}
	h.targets = make([]Target, 0, 5)

	return h
}

func (h *Handler) SetTargetPos(i int, x, y float64) bool {
	if i < 0 || i > len(h.targets)-1 {
		return false
	}
	h.targets[i].SetPos(x, y)
	h.Log.NewEvent("target position changed", glog.LogSimEvent, -1).
		Write("index", i).
		Write("x", x).
		Write("y", y)
	return true
}

func (h *Handler) KillTarget(i int) bool {
	// don't kill yourself
	if i < 1 || i > len(h.targets)-1 {
		return false
	}
	h.targets[i].Kill()
	h.Events.Emit(event.OnTargetDied, h.targets[i], &AttackEvent{}) // TODO: it's fine?
	h.Log.NewEvent("target is dead", glog.LogSimEvent, -1).Write("index", i)
	return true
}

func (h *Handler) Tick() {
	//collision check happens before each object ticks (as collision may remove the object)
	for i := 0; i < len(h.targets); i++ {
		if !h.targets[i].Collidable() {
			continue
		}
		for j := 0; j < len(h.targets); j++ {
			if i == j {
				continue
			}
			if !h.targets[i].CollidableWith(h.targets[j].Type()) {
				continue
			}
			//check collision here
			if true {
				h.targets[i].CollidedWith(h.targets[j])
			}
		}
	}
	for _, t := range h.targets {
		t.Tick()
	}
}
