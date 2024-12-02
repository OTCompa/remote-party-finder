package ffxiv
type World string

const (
	Ravana	World = "Ravana"
	Bismarck	World = "Bismarck"
	Asura	World = "Asura"
	Belias	World = "Belias"
	Pandaemonium	World = "Pandaemonium"
	Shinryu	World = "Shinryu"
	Unicorn	World = "Unicorn"
	Yojimbo	World = "Yojimbo"
	Zeromus	World = "Zeromus"
	Twintania	World = "Twintania"
	Brynhildr	World = "Brynhildr"
	Famfrit	World = "Famfrit"
	Lich	World = "Lich"
	Mateus	World = "Mateus"
	Omega	World = "Omega"
	Jenova	World = "Jenova"
	Zalera	World = "Zalera"
	Zodiark	World = "Zodiark"
	Alexander	World = "Alexander"
	Anima	World = "Anima"
	Carbuncle	World = "Carbuncle"
	Fenrir	World = "Fenrir"
	Hades	World = "Hades"
	Ixion	World = "Ixion"
	Kujata	World = "Kujata"
	Typhon	World = "Typhon"
	Ultima	World = "Ultima"
	Valefor	World = "Valefor"
	Exodus	World = "Exodus"
	Faerie	World = "Faerie"
	Lamia	World = "Lamia"
	Phoenix	World = "Phoenix"
	Siren	World = "Siren"
	Garuda	World = "Garuda"
	Ifrit	World = "Ifrit"
	Ramuh	World = "Ramuh"
	Titan	World = "Titan"
	Diabolos	World = "Diabolos"
	Gilgamesh	World = "Gilgamesh"
	Leviathan	World = "Leviathan"
	Midgardsormr	World = "Midgardsormr"
	Odin	World = "Odin"
	Shiva	World = "Shiva"
	Atomos	World = "Atomos"
	Bahamut	World = "Bahamut"
	Chocobo	World = "Chocobo"
	Moogle	World = "Moogle"
	Tonberry	World = "Tonberry"
	Adamantoise	World = "Adamantoise"
	Coeurl	World = "Coeurl"
	Malboro	World = "Malboro"
	Tiamat	World = "Tiamat"
	Ultros	World = "Ultros"
	Behemoth	World = "Behemoth"
	Cactuar	World = "Cactuar"
	Cerberus	World = "Cerberus"
	Goblin	World = "Goblin"
	Mandragora	World = "Mandragora"
	Louisoix	World = "Louisoix"
	Spriggan	World = "Spriggan"
	Sephirot	World = "Sephirot"
	Sophia	World = "Sophia"
	Zurvan	World = "Zurvan"
	Aegis	World = "Aegis"
	Balmung	World = "Balmung"
	Durandal	World = "Durandal"
	Excalibur	World = "Excalibur"
	Gungnir	World = "Gungnir"
	Hyperion	World = "Hyperion"
	Masamune	World = "Masamune"
	Ragnarok	World = "Ragnarok"
	Ridill	World = "Ridill"
	Sargatanas	World = "Sargatanas"
	Sagittarius	World = "Sagittarius"
	Phantom	World = "Phantom"
	Alpha	World = "Alpha"
	Raiden	World = "Raiden"
	Marilith	World = "Marilith"
	Seraph	World = "Seraph"
	Halicarnassus	World = "Halicarnassus"
	Maduin	World = "Maduin"
	Cuchulainn	World = "Cuchulainn"
	Kraken	World = "Kraken"
	Rafflesia	World = "Rafflesia"
	Golem	World = "Golem"
	Cloudtest01	World = "Cloudtest01"
	Cloudtest02	World = "Cloudtest02"
)

var WORLDS = map[uint32]World{
	21: Ravana,
	22: Bismarck,
	23: Asura,
	24: Belias,
	28: Pandaemonium,
	29: Shinryu,
	30: Unicorn,
	31: Yojimbo,
	32: Zeromus,
	33: Twintania,
	34: Brynhildr,
	35: Famfrit,
	36: Lich,
	37: Mateus,
	39: Omega,
	40: Jenova,
	41: Zalera,
	42: Zodiark,
	43: Alexander,
	44: Anima,
	45: Carbuncle,
	46: Fenrir,
	47: Hades,
	48: Ixion,
	49: Kujata,
	50: Typhon,
	51: Ultima,
	52: Valefor,
	53: Exodus,
	54: Faerie,
	55: Lamia,
	56: Phoenix,
	57: Siren,
	58: Garuda,
	59: Ifrit,
	60: Ramuh,
	61: Titan,
	62: Diabolos,
	63: Gilgamesh,
	64: Leviathan,
	65: Midgardsormr,
	66: Odin,
	67: Shiva,
	68: Atomos,
	69: Bahamut,
	70: Chocobo,
	71: Moogle,
	72: Tonberry,
	73: Adamantoise,
	74: Coeurl,
	75: Malboro,
	76: Tiamat,
	77: Ultros,
	78: Behemoth,
	79: Cactuar,
	80: Cerberus,
	81: Goblin,
	82: Mandragora,
	83: Louisoix,
	85: Spriggan,
	86: Sephirot,
	87: Sophia,
	88: Zurvan,
	90: Aegis,
	91: Balmung,
	92: Durandal,
	93: Excalibur,
	94: Gungnir,
	95: Hyperion,
	96: Masamune,
	97: Ragnarok,
	98: Ridill,
	99: Sargatanas,
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
	3000: Cloudtest01,
	3001: Cloudtest02,
}
