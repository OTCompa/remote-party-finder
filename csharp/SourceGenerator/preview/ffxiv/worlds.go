package ffxiv

type Datacenter int

const (
	Elemental Datacenter = 1
	Gaia      Datacenter = 3
	Mana      Datacenter = 5
	Aether    Datacenter = 2
	Primal    Datacenter = 4
	Chaos     Datacenter = 6
	Light     Datacenter = 7
	Crystal   Datacenter = 8
	Materia   Datacenter = 9
	Meteor    Datacenter = 10
	Dynamis   Datacenter = 11
	Shadow    Datacenter = 12
)

func DatacenterJP() []Datacenter {
	return []Datacenter{Elemental, Gaia, Mana, Meteor}
}

func DatacenterNA() []Datacenter {
	return []Datacenter{Aether, Primal, Crystal, Dynamis}
}

func DatacenterEU() []Datacenter {
	return []Datacenter{Chaos, Light, Shadow}
}

func DatacenterOC() []Datacenter {
	return []Datacenter{Materia}
}

func DatacenterOthers() []Datacenter {
	return []Datacenter{}
}

type World struct {
	Name       string
	Datacenter Datacenter
}

var (
	Ravana        World = World{Name: "Ravana", Datacenter: Materia}
	Bismarck      World = World{Name: "Bismarck", Datacenter: Materia}
	Asura         World = World{Name: "Asura", Datacenter: Mana}
	Belias        World = World{Name: "Belias", Datacenter: Meteor}
	Pandaemonium  World = World{Name: "Pandaemonium", Datacenter: Mana}
	Shinryu       World = World{Name: "Shinryu", Datacenter: Meteor}
	Unicorn       World = World{Name: "Unicorn", Datacenter: Meteor}
	Yojimbo       World = World{Name: "Yojimbo", Datacenter: Meteor}
	Zeromus       World = World{Name: "Zeromus", Datacenter: Meteor}
	Twintania     World = World{Name: "Twintania", Datacenter: Light}
	Brynhildr     World = World{Name: "Brynhildr", Datacenter: Crystal}
	Famfrit       World = World{Name: "Famfrit", Datacenter: Primal}
	Lich          World = World{Name: "Lich", Datacenter: Light}
	Mateus        World = World{Name: "Mateus", Datacenter: Crystal}
	Omega         World = World{Name: "Omega", Datacenter: Chaos}
	Jenova        World = World{Name: "Jenova", Datacenter: Aether}
	Zalera        World = World{Name: "Zalera", Datacenter: Crystal}
	Zodiark       World = World{Name: "Zodiark", Datacenter: Light}
	Alexander     World = World{Name: "Alexander", Datacenter: Gaia}
	Anima         World = World{Name: "Anima", Datacenter: Mana}
	Carbuncle     World = World{Name: "Carbuncle", Datacenter: Elemental}
	Fenrir        World = World{Name: "Fenrir", Datacenter: Gaia}
	Hades         World = World{Name: "Hades", Datacenter: Mana}
	Ixion         World = World{Name: "Ixion", Datacenter: Mana}
	Kujata        World = World{Name: "Kujata", Datacenter: Elemental}
	Typhon        World = World{Name: "Typhon", Datacenter: Elemental}
	Ultima        World = World{Name: "Ultima", Datacenter: Gaia}
	Valefor       World = World{Name: "Valefor", Datacenter: Meteor}
	Exodus        World = World{Name: "Exodus", Datacenter: Primal}
	Faerie        World = World{Name: "Faerie", Datacenter: Aether}
	Lamia         World = World{Name: "Lamia", Datacenter: Primal}
	Phoenix       World = World{Name: "Phoenix", Datacenter: Light}
	Siren         World = World{Name: "Siren", Datacenter: Aether}
	Garuda        World = World{Name: "Garuda", Datacenter: Elemental}
	Ifrit         World = World{Name: "Ifrit", Datacenter: Gaia}
	Ramuh         World = World{Name: "Ramuh", Datacenter: Meteor}
	Titan         World = World{Name: "Titan", Datacenter: Mana}
	Diabolos      World = World{Name: "Diabolos", Datacenter: Crystal}
	Gilgamesh     World = World{Name: "Gilgamesh", Datacenter: Aether}
	Leviathan     World = World{Name: "Leviathan", Datacenter: Primal}
	Midgardsormr  World = World{Name: "Midgardsormr", Datacenter: Aether}
	Odin          World = World{Name: "Odin", Datacenter: Light}
	Shiva         World = World{Name: "Shiva", Datacenter: Light}
	Atomos        World = World{Name: "Atomos", Datacenter: Elemental}
	Bahamut       World = World{Name: "Bahamut", Datacenter: Gaia}
	Chocobo       World = World{Name: "Chocobo", Datacenter: Mana}
	Moogle        World = World{Name: "Moogle", Datacenter: Chaos}
	Tonberry      World = World{Name: "Tonberry", Datacenter: Elemental}
	Adamantoise   World = World{Name: "Adamantoise", Datacenter: Aether}
	Coeurl        World = World{Name: "Coeurl", Datacenter: Crystal}
	Malboro       World = World{Name: "Malboro", Datacenter: Crystal}
	Tiamat        World = World{Name: "Tiamat", Datacenter: Gaia}
	Ultros        World = World{Name: "Ultros", Datacenter: Primal}
	Behemoth      World = World{Name: "Behemoth", Datacenter: Primal}
	Cactuar       World = World{Name: "Cactuar", Datacenter: Aether}
	Cerberus      World = World{Name: "Cerberus", Datacenter: Chaos}
	Goblin        World = World{Name: "Goblin", Datacenter: Crystal}
	Mandragora    World = World{Name: "Mandragora", Datacenter: Meteor}
	Louisoix      World = World{Name: "Louisoix", Datacenter: Chaos}
	Spriggan      World = World{Name: "Spriggan", Datacenter: Chaos}
	Sephirot      World = World{Name: "Sephirot", Datacenter: Materia}
	Sophia        World = World{Name: "Sophia", Datacenter: Materia}
	Zurvan        World = World{Name: "Zurvan", Datacenter: Materia}
	Aegis         World = World{Name: "Aegis", Datacenter: Elemental}
	Balmung       World = World{Name: "Balmung", Datacenter: Crystal}
	Durandal      World = World{Name: "Durandal", Datacenter: Gaia}
	Excalibur     World = World{Name: "Excalibur", Datacenter: Primal}
	Gungnir       World = World{Name: "Gungnir", Datacenter: Elemental}
	Hyperion      World = World{Name: "Hyperion", Datacenter: Primal}
	Masamune      World = World{Name: "Masamune", Datacenter: Mana}
	Ragnarok      World = World{Name: "Ragnarok", Datacenter: Chaos}
	Ridill        World = World{Name: "Ridill", Datacenter: Gaia}
	Sargatanas    World = World{Name: "Sargatanas", Datacenter: Aether}
	Sagittarius   World = World{Name: "Sagittarius", Datacenter: Chaos}
	Phantom       World = World{Name: "Phantom", Datacenter: Chaos}
	Alpha         World = World{Name: "Alpha", Datacenter: Light}
	Raiden        World = World{Name: "Raiden", Datacenter: Light}
	Marilith      World = World{Name: "Marilith", Datacenter: Dynamis}
	Seraph        World = World{Name: "Seraph", Datacenter: Dynamis}
	Halicarnassus World = World{Name: "Halicarnassus", Datacenter: Dynamis}
	Maduin        World = World{Name: "Maduin", Datacenter: Dynamis}
	Cuchulainn    World = World{Name: "Cuchulainn", Datacenter: Dynamis}
	Kraken        World = World{Name: "Kraken", Datacenter: Dynamis}
	Rafflesia     World = World{Name: "Rafflesia", Datacenter: Dynamis}
	Golem         World = World{Name: "Golem", Datacenter: Dynamis}
)

var WORLDS = map[uint32]World{
	21:  Ravana,
	22:  Bismarck,
	23:  Asura,
	24:  Belias,
	28:  Pandaemonium,
	29:  Shinryu,
	30:  Unicorn,
	31:  Yojimbo,
	32:  Zeromus,
	33:  Twintania,
	34:  Brynhildr,
	35:  Famfrit,
	36:  Lich,
	37:  Mateus,
	39:  Omega,
	40:  Jenova,
	41:  Zalera,
	42:  Zodiark,
	43:  Alexander,
	44:  Anima,
	45:  Carbuncle,
	46:  Fenrir,
	47:  Hades,
	48:  Ixion,
	49:  Kujata,
	50:  Typhon,
	51:  Ultima,
	52:  Valefor,
	53:  Exodus,
	54:  Faerie,
	55:  Lamia,
	56:  Phoenix,
	57:  Siren,
	58:  Garuda,
	59:  Ifrit,
	60:  Ramuh,
	61:  Titan,
	62:  Diabolos,
	63:  Gilgamesh,
	64:  Leviathan,
	65:  Midgardsormr,
	66:  Odin,
	67:  Shiva,
	68:  Atomos,
	69:  Bahamut,
	70:  Chocobo,
	71:  Moogle,
	72:  Tonberry,
	73:  Adamantoise,
	74:  Coeurl,
	75:  Malboro,
	76:  Tiamat,
	77:  Ultros,
	78:  Behemoth,
	79:  Cactuar,
	80:  Cerberus,
	81:  Goblin,
	82:  Mandragora,
	83:  Louisoix,
	85:  Spriggan,
	86:  Sephirot,
	87:  Sophia,
	88:  Zurvan,
	90:  Aegis,
	91:  Balmung,
	92:  Durandal,
	93:  Excalibur,
	94:  Gungnir,
	95:  Hyperion,
	96:  Masamune,
	97:  Ragnarok,
	98:  Ridill,
	99:  Sargatanas,
	400: Sagittarius,
	401: Phantom,
	402: Alpha,
	403: Raiden,
	404: Marilith,
	405: Seraph,
	406: Halicarnassus,
	407: Maduin,
	408: Cuchulainn,
	409: Kraken,
	410: Rafflesia,
	411: Golem,
}
