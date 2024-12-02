package ffxiv

type RouletteInfo struct {
	Name LocalisedText
	PVP  bool
}

var ROULETTES = map[uint32]RouletteInfo{
	1: {
		Name: LocalisedText{
              En: "Duty Roulette: Leveling",
              Ja: "コンテンツルーレット：レべリング",
              De: "Zufallsinhalt: Stufensteigerung",
              Fr: "Mission aléatoire : gain de niveaux",
            },
		PVP: false,
	},
	2: {
		Name: LocalisedText{
              En: "Duty Roulette: High-level Dungeons",
              Ja: "コンテンツルーレット：ハイレベリング",
              De: "Zufallsinhalt: Hohe Stufen",
              Fr: "Mission aléatoire : donjons avancés",
            },
		PVP: false,
	},
	3: {
		Name: LocalisedText{
              En: "Duty Roulette: Main Scenario",
              Ja: "コンテンツルーレット：メインクエスト",
              De: "Zufallsinhalt: Hauptszenario",
              Fr: "Mission aléatoire : épopée",
            },
		PVP: false,
	},
	4: {
		Name: LocalisedText{
              En: "Duty Roulette: Guildhests",
              Ja: "コンテンツルーレット：ギルドオーダー",
              De: "Zufallsinhalt: Gildengeheiß",
              Fr: "Mission aléatoire : opérations de guilde",
            },
		PVP: false,
	},
	5: {
		Name: LocalisedText{
              En: "Duty Roulette: Expert",
              Ja: "コンテンツルーレット：エキスパート",
              De: "Zufallsinhalt: Experte",
              Fr: "Mission aléatoire : expert",
            },
		PVP: false,
	},
	6: {
		Name: LocalisedText{
              En: "Duty Roulette: Trials",
              Ja: "コンテンツルーレット：討伐・討滅戦",
              De: "Zufallsinhalt: Prüfung",
              Fr: "Mission aléatoire : défis",
            },
		PVP: false,
	},
	7: {
		Name: LocalisedText{
              En: "Daily Challenge: Frontline",
              Ja: "デイリーチャレンジ：フロントライン",
              De: "Tagesherausforderung: PvP-Front",
              Fr: "Challenge quotidien : Front",
            },
		PVP: true,
	},
	8: {
		Name: LocalisedText{
              En: "Duty Roulette: Level Cap Dungeons",
              Ja: "コンテンツルーレット：レベルキャップダンジョン",
              De: "Zufallsinhalt: Höchststufe",
              Fr: "Mission aléatoire : donjons supérieurs",
            },
		PVP: false,
	},
	9: {
		Name: LocalisedText{
              En: "Duty Roulette: Mentor",
              Ja: "コンテンツルーレット：メンター",
              De: "Zufallsinhalt: Mentor",
              Fr: "Mission aléatoire : mentor",
            },
		PVP: false,
	},
	15: {
		Name: LocalisedText{
              En: "Duty Roulette: Alliance Raids",
              Ja: "コンテンツルーレット：アライアンスレイド",
              De: "Zufallsinhalt: Allianz-Raid",
              Fr: "Mission aléatoire : raids en alliance",
            },
		PVP: false,
	},
	17: {
		Name: LocalisedText{
              En: "Duty Roulette: Normal Raids",
              Ja: "コンテンツルーレット：ノーマルレイド",
              De: "Zufallsinhalt: Normaler Raid",
              Fr: "Mission aléatoire : raids normaux",
            },
		PVP: false,
	},
	18: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos : Route de Sagolii",
            },
		PVP: false,
	},
	19: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos : Costa del Sol",
            },
		PVP: false,
	},
	20: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos : Sentes tranquilles",
            },
		PVP: false,
	},
	21: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	22: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road (No Rewards)",
              Ja: "チョコボレース：サゴリーロード (報酬なし)",
              De: "Chocobo-Rennen: Sagolii-Straße (keine Belohnung)",
              Fr: "Course de chocobos : Route de Sagolii (sans récompense)",
            },
		PVP: false,
	},
	23: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol (No Rewards)",
              Ja: "チョコボレース：コスタ・デル・ソル (報酬なし)",
              De: "Chocobo-Rennen: Sonnenküste (keine Belohnung)",
              Fr: "Course de chocobos : Costa del Sol (sans récompense)",
            },
		PVP: false,
	},
	24: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths (No Rewards)",
              Ja: "チョコボレース：トランキルパス (報酬なし)",
              De: "Chocobo-Rennen: Pfad der Seelenruhe (keine Belohnung)",
              Fr: "Course de chocobos : Sentes tranquilles (sans récompense)",
            },
		PVP: false,
	},
	25: {
		Name: LocalisedText{
              En: "Chocobo Race: Random (No Rewards)",
              Ja: "チョコボレース：コースルーレット (報酬なし)",
              De: "Chocobo-Rennen: Zufallsstrecke (keine Belohnung)",
              Fr: "Course de chocobos : aléatoire (sans récompense)",
            },
		PVP: false,
	},
	26: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	27: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	28: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	29: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	30: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	31: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	32: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	33: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	34: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	35: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	36: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	37: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	38: {
		Name: LocalisedText{
              En: "Chocobo Race: Random",
              Ja: "チョコボレース：コースルーレット",
              De: "Chocobo-Rennen: Zufallsstrecke",
              Fr: "Course de chocobos : aléatoire",
            },
		PVP: false,
	},
	40: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Casual Match)",
              Ja: "クリスタルコンフリクト(カジュアルマッチ)",
              De: "Crystalline Conflict: Freies Spiel",
              Fr: "Crystalline Conflict (partie non classée)",
            },
		PVP: true,
	},
	41: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Ranked Match)",
              Ja: "クリスタルコンフリクト(ランクマッチ)",
              De: "Crystalline Conflict: Gewertetes Spiel",
              Fr: "Crystalline Conflict (partie classée)",
            },
		PVP: true,
	},
}
