package ffxiv

type ClassJob interface {
	isClassJob()
}

type Class string

func (Class) isClassJob() {}

type Job string

func (Job) isClassJob() {}

type NonCombatJob string

func (NonCombatJob) isClassJob() {}

const (
	Gladiator   Class = "Gladiator"
	Pugilist    Class = "Pugilist"
	Marauder    Class = "Marauder"
	Lancer      Class = "Lancer"
	Archer      Class = "Archer"
	Conjurer    Class = "Conjurer"
	Thaumaturge Class = "Thaumaturge"
	Arcanist    Class = "Arcanist"
	Rogue       Class = "Rogue"
)

const (
	Paladin     Job = "Paladin"
	Monk        Job = "Monk"
	Warrior     Job = "Warrior"
	Dragoon     Job = "Dragoon"
	Bard        Job = "Bard"
	WhiteMage   Job = "WhiteMage"
	BlackMage   Job = "BlackMage"
	Summoner    Job = "Summoner"
	Scholar     Job = "Scholar"
	Ninja       Job = "Ninja"
	Machinist   Job = "Machinist"
	DarkKnight  Job = "DarkKnight"
	Astrologian Job = "Astrologian"
	Samurai     Job = "Samurai"
	RedMage     Job = "RedMage"
	BlueMage    Job = "BlueMage"
	Gunbreaker  Job = "Gunbreaker"
	Dancer      Job = "Dancer"
	Reaper      Job = "Reaper"
	Sage        Job = "Sage"
	Viper       Job = "Viper"
	Pictomancer Job = "Pictomancer"
)

const (
	Carpenter     NonCombatJob = "Carpenter"
	Blacksmith    NonCombatJob = "Blacksmith"
	Armorer       NonCombatJob = "Armorer"
	Goldsmith     NonCombatJob = "Goldsmith"
	Leatherworker NonCombatJob = "Leatherworker"
	Weaver        NonCombatJob = "Weaver"
	Alchemist     NonCombatJob = "Alchemist"
	Culinarian    NonCombatJob = "Culinarian"
	Miner         NonCombatJob = "Miner"
	Botanist      NonCombatJob = "Botanist"
	Fisher        NonCombatJob = "Fisher"
)

var JOBS = map[uint32]ClassJob{
	1:  Gladiator,
	2:  Pugilist,
	3:  Marauder,
	4:  Lancer,
	5:  Archer,
	6:  Conjurer,
	7:  Thaumaturge,
	8:  Carpenter,
	9:  Blacksmith,
	10: Armorer,
	11: Goldsmith,
	12: Leatherworker,
	13: Weaver,
	14: Alchemist,
	15: Culinarian,
	16: Miner,
	17: Botanist,
	18: Fisher,
	19: Paladin,
	20: Monk,
	21: Warrior,
	22: Dragoon,
	23: Bard,
	24: WhiteMage,
	25: BlackMage,
	26: Arcanist,
	27: Summoner,
	28: Scholar,
	29: Rogue,
	30: Ninja,
	31: Machinist,
	32: DarkKnight,
	33: Astrologian,
	34: Samurai,
	35: RedMage,
	36: BlueMage,
	37: Gunbreaker,
	38: Dancer,
	39: Reaper,
	40: Sage,
	41: Viper,
	42: Pictomancer,
}
