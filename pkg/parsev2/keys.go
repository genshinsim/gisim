package parse

import "github.com/genshinsim/gcsim/pkg/core"

var key = map[string]ItemType{
	".": itemDot,
	//commands
	"chain":       itemChain,
	"wait":        itemWait,
	"reset_limit": itemResetLimit,
	"hurt":        itemHurt,
	"energy":      itemEnergy,
	"active":      itemActive,
	"options":     itemOptions,
	//options
	"debug":    itemDebug,
	"iter":     itemIterations,
	"duration": itemDuration,
	"workers":  itemWorkers,
	"er_calc":  itemERCalc,
	//team keywords
	"add":      itemAdd,
	"char":     itemChar,
	"stats":    itemStats,
	"weapon":   itemWeapon,
	"set":      itemSet,
	"lvl":      itemLvl,
	"refine":   itemRefine,
	"cons":     itemCons,
	"talent":   itemTalent,
	"start_hp": itemStartHP,
	"count":    itemCount,
	"param":    itemParam,
	//flags
	"swap_lock": itemSwapLock,
	"if":        itemIf,
	"swap":      itemSwap,
	"onfield":   itemOnField,
	"limit":     itemLimit,
	"try":       itemTry,
	"timeout":   itemTimeout,
	"char_mod":  itemCharMod,
	"particle":  itemParticle,
	"max":       itemMax,
	//energy/hurt event related
	"interval": itemInterval,
	"amount":   itemAmount,
	"once":     itemOnce,
	"every":    itemEvery,
}

var statKeys = map[string]core.StatType{
	"def%":     core.DEFP,
	"def":      core.DEF,
	"hp":       core.HP,
	"hp%":      core.HPP,
	"atk":      core.ATK,
	"atk%":     core.ATKP,
	"er":       core.ER,
	"em":       core.EM,
	"cr":       core.CR,
	"cd":       core.CD,
	"heal":     core.Heal,
	"pyro%":    core.PyroP,
	"hydro%":   core.HydroP,
	"cryo%":    core.CryoP,
	"electro%": core.ElectroP,
	"anemo%":   core.AnemoP,
	"geo%":     core.GeoP,
	"phys%":    core.PhyP,
	"ele%":     core.EleP,
	"dendro%":  core.DendroP,
	"atkspd%":  core.AtkSpd,
	"dmg%":     core.DmgP,
}

var eleKeys = map[string]core.EleType{
	"electro":  core.Electro,
	"pyro":     core.Pyro,
	"cryo":     core.Cryo,
	"hydro":    core.Hydro,
	"frozen":   core.Frozen,
	"anemo":    core.Anemo,
	"dendro":   core.Dendro,
	"geo":      core.Geo,
	"physical": core.Physical,
}

var actionKeys = map[string]core.ActionType{
	"skill":       core.ActionSkill,
	"burst":       core.ActionBurst,
	"attack":      core.ActionAttack,
	"charge":      core.ActionCharge,
	"high_plunge": core.ActionHighPlunge,
	"low_lunge":   core.ActionLowPlunge,
	"aim":         core.ActionAim,
	"dash":        core.ActionDash,
	"jump":        core.ActionJump,
	// "swap":        core.ActionSwap,
}
