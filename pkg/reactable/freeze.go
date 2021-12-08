package reactable

import "github.com/genshinsim/gcsim/pkg/core"

func (r *Reactable) tryFreeze(a *core.AttackEvent) {
	//so if already frozen there are 2 cases:
	// 1. src exists but no other coexisting -> attach
	// 2. src does not exist but opposite coexists -> add to freeze durability
	switch a.Info.Element {
	case core.Hydro:
		//if cryo exists we'll trigger freeze regardless if frozen already coexists
		if r.Durability[core.Cryo] > zeroDur {
			consumed := r.triggerFreeze(r.Durability[core.Cryo], a.Info.Durability)
			r.Durability[core.Cryo] -= consumed
			r.Durability[core.Cryo] = max(r.Durability[core.Cryo], 0)
			//TODO: we're not setting src durability to zero here but should be ok b/c no reaction comes after freeze
			//ec should have been taken care of already
			a.Info.Durability -= consumed
			a.Info.Durability = max(a.Info.Durability, 0)
			return
		}
		//otherwise attach hydro only if frozen exists
		if r.Durability[core.Frozen] < zeroDur {
			return
		}
		//try refill first - this will use up all durability if ok
		r.tryRefill(core.Hydro, &a.Info.Durability)
		//otherwise attach
		r.tryAttach(core.Hydro, &a.Info.Durability)
	case core.Cryo:
		if r.Durability[core.Hydro] > zeroDur {
			consumed := r.triggerFreeze(r.Durability[core.Hydro], a.Info.Durability)
			r.Durability[core.Hydro] -= consumed
			r.Durability[core.Hydro] = max(r.Durability[core.Hydro], 0)
			a.Info.Durability -= consumed
			a.Info.Durability = max(a.Info.Durability, 0)
			return
		}
		//otherwise attach cryo only if frozen exists
		if r.Durability[core.Frozen] < zeroDur {
			return
		}
		//try refill first - this will use up all durability if ok
		r.tryRefill(core.Cryo, &a.Info.Durability)
		//otherwise attach
		r.tryAttach(core.Cryo, &a.Info.Durability)
	default:
		//should be here
		return
	}
	r.core.Events.Emit(core.OnFrozen, r.self, a)
}

func max(a, b core.Durability) core.Durability {
	if a > b {
		return a
	}
	return b
}

func min(a, b core.Durability) core.Durability {
	if a > b {
		return b
	}
	return a
}

//add to freeze durability and return amount of durability consumed
func (r *Reactable) triggerFreeze(a, b core.Durability) core.Durability {
	d := min(a, b)
	//trigger freeze should only addDurability and should not touch decay rate
	r.addDurability(core.Frozen, 2*d)
	return d
}
