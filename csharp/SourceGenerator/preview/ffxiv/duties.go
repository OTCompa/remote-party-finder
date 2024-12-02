package ffxiv

type DutyInfo struct {
	Name        LocalisedText
	HighEnd     bool
	ContentKind ContentKind
}

var DUTIES = map[uint32]DutyInfo{
	1: {
		Name: LocalisedText{
              En: "The Thousand Maws of Toto-Rak",
              Ja: "監獄廃墟 トトラクの千獄",
              De: "Tausend Löcher von Toto-Rak",
              Fr: "Les Mille Gueules de Toto-Rak",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	2: {
		Name: LocalisedText{
              En: "The Tam-Tara Deepcroft",
              Ja: "地下霊殿 タムタラの墓所",
              De: "Totenacker Tam-Tara",
              Fr: "L'Hypogée de Tam-Tara",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	3: {
		Name: LocalisedText{
              En: "Copperbell Mines",
              Ja: "封鎖坑道 カッパーベル銅山",
              De: "Kupferglocken-Mine",
              Fr: "Les Mines de Clochecuivre",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	4: {
		Name: LocalisedText{
              En: "Sastasha",
              Ja: "天然要害 サスタシャ浸食洞",
              De: "Sastasha",
              Fr: "Sastasha",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	5: {
		Name: LocalisedText{
              En: "The Aurum Vale",
              Ja: "霧中行軍 オーラムヴェイル",
              De: "Goldklamm",
              Fr: "Le Val d'Aurum",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	6: {
		Name: LocalisedText{
              En: "Haukke Manor",
              Ja: "名門屋敷 ハウケタ御用邸",
              De: "Haukke-Herrenhaus",
              Fr: "Le Manoir des Haukke",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	7: {
		Name: LocalisedText{
              En: "Halatali",
              Ja: "魔獣領域 ハラタリ修練所",
              De: "Halatali",
              Fr: "Halatali",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	8: {
		Name: LocalisedText{
              En: "Brayflox's Longstop",
              Ja: "奪還支援 ブレイフロクスの野営地",
              De: "Brüllvolx' Langrast",
              Fr: "Le Bivouac de Brayflox",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	9: {
		Name: LocalisedText{
              En: "The Sunken Temple of Qarn",
              Ja: "遺跡探索 カルン埋没寺院",
              De: "Versunkener Tempel von Qarn",
              Fr: "Le Temple enseveli de Qarn",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	10: {
		Name: LocalisedText{
              En: "The Wanderer's Palace",
              Ja: "旅神聖域 ワンダラーパレス",
              De: "Palast des Wanderers",
              Fr: "Le Palais du Vagabond",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	11: {
		Name: LocalisedText{
              En: "The Stone Vigil",
              Ja: "城塞攻略 ストーンヴィジル",
              De: "Steinerne Wacht",
              Fr: "Le Vigile de Pierre",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	12: {
		Name: LocalisedText{
              En: "Cutter's Cry",
              Ja: "流砂迷宮 カッターズクライ",
              De: "Sägerschrei",
              Fr: "Le Gouffre hurlant",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	13: {
		Name: LocalisedText{
              En: "Dzemael Darkhold",
              Ja: "掃討作戦 ゼーメル要塞",
              De: "Die Feste Dzemael",
              Fr: "La Forteresse de Dzemael",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	14: {
		Name: LocalisedText{
              En: "Amdapor Keep",
              Ja: "邪教排撃 古城アムダプール",
              De: "Die Ruinen von Amdapor",
              Fr: "Le Château d'Amdapor",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	15: {
		Name: LocalisedText{
              En: "Castrum Meridianum",
              Ja: "外郭攻略 カストルム・メリディアヌム",
              De: "Castrum Meridianum - Außenbereich",
              Fr: "Castrum Meridianum",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	16: {
		Name: LocalisedText{
              En: "The Praetorium",
              Ja: "最終決戦 魔導城プラエトリウム",
              De: "Castrum Meridianum - Praetorium",
              Fr: "Le Praetorium",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	17: {
		Name: LocalisedText{
              En: "Pharos Sirius",
              Ja: "怪鳥巨塔 シリウス大灯台",
              De: "Pharos Sirius",
              Fr: "Le Phare de Sirius",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	18: {
		Name: LocalisedText{
              En: "Copperbell Mines (Hard)",
              Ja: "騒乱坑道 カッパーベル銅山 (Hard)",
              De: "Kupferglocken-Mine (schwer)",
              Fr: "Les Mines de Clochecuivre (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	19: {
		Name: LocalisedText{
              En: "Haukke Manor (Hard)",
              Ja: "妖異屋敷 ハウケタ御用邸 (Hard)",
              De: "Haukke-Herrenhaus (schwer)",
              Fr: "Le Manoir des Haukke (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	20: {
		Name: LocalisedText{
              En: "Brayflox's Longstop (Hard)",
              Ja: "盟友支援 ブレイフロクスの野営地 (Hard)",
              De: "Brüllvolx' Langrast (schwer)",
              Fr: "Le Bivouac de Brayflox (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	21: {
		Name: LocalisedText{
              En: "Halatali (Hard)",
              Ja: "剣闘領域 ハラタリ修練所 (Hard)",
              De: "Halatali (schwer)",
              Fr: "Halatali (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	22: {
		Name: LocalisedText{
              En: "The Lost City of Amdapor",
              Ja: "腐敗遺跡 古アムダプール市街",
              De: "Historisches Amdapor",
              Fr: "Les Vestiges de la cité d'Amdapor",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	23: {
		Name: LocalisedText{
              En: "Hullbreaker Isle",
              Ja: "財宝伝説 ハルブレーカー・アイル",
              De: "Schiffbrecher-Insel",
              Fr: "L'Île de Crèvecarène",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	24: {
		Name: LocalisedText{
              En: "The Tam-Tara Deepcroft (Hard)",
              Ja: "惨劇霊殿 タムタラの墓所 (Hard)",
              De: "Totenacker Tam-Tara (schwer)",
              Fr: "L'Hypogée de Tam-Tara (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	25: {
		Name: LocalisedText{
              En: "The Stone Vigil (Hard)",
              Ja: "城塞奪回 ストーンヴィジル (Hard)",
              De: "Steinerne Wacht (schwer)",
              Fr: "Le Vigile de Pierre (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	26: {
		Name: LocalisedText{
              En: "The Sunken Temple of Qarn (Hard)",
              Ja: "遺跡救援 カルン埋没寺院 (Hard)",
              De: "Versunkener Tempel von Qarn (schwer)",
              Fr: "Le Temple enseveli de Qarn (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	27: {
		Name: LocalisedText{
              En: "Snowcloak",
              Ja: "氷結潜窟 スノークローク大氷壁",
              De: "Das Schneekleid",
              Fr: "Manteneige",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	28: {
		Name: LocalisedText{
              En: "Sastasha (Hard)",
              Ja: "逆襲要害 サスタシャ浸食洞 (Hard)",
              De: "Sastasha (schwer)",
              Fr: "Sastasha (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	29: {
		Name: LocalisedText{
              En: "Amdapor Keep (Hard)",
              Ja: "邪念排撃 古城アムダプール (Hard)",
              De: "Die Ruinen von Amdapor (schwer)",
              Fr: "Le Château d'Amdapor (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	30: {
		Name: LocalisedText{
              En: "The Wanderer's Palace (Hard)",
              Ja: "武装聖域 ワンダラーパレス (Hard)",
              De: "Palast des Wanderers (schwer)",
              Fr: "Le Palais du Vagabond (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	31: {
		Name: LocalisedText{
              En: "The Great Gubal Library",
              Ja: "禁書回収 グブラ幻想図書館",
              De: "Große Gubal-Bibliothek",
              Fr: "La Grande bibliothèque de Gubal",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	32: {
		Name: LocalisedText{
              En: "The Keeper of the Lake",
              Ja: "幻龍残骸 黙約の塔",
              De: "Hüter des Sees",
              Fr: "Le Gardien du lac",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	33: {
		Name: LocalisedText{
              En: "Neverreap",
              Ja: "神域浮島 ネバーリープ",
              De: "Nimmerreich",
              Fr: "Nalloncques",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	34: {
		Name: LocalisedText{
              En: "The Vault",
              Ja: "強硬突入 イシュガルド教皇庁",
              De: "Erzbasilika",
              Fr: "La Voûte",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	35: {
		Name: LocalisedText{
              En: "The Fractal Continuum",
              Ja: "博物戦艦 フラクタル・コンティニアム",
              De: "Die Fraktal-Kontinuum",
              Fr: "Le Continuum fractal",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	36: {
		Name: LocalisedText{
              En: "The Dusk Vigil",
              Ja: "廃砦捜索 ダスクヴィジル",
              De: "Abendrot-Wacht",
              Fr: "Le Vigile du Crépuscule",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	37: {
		Name: LocalisedText{
              En: "Sohm Al",
              Ja: "霊峰踏破 ソーム・アル",
              De: "Sohm Al",
              Fr: "Sohm Al",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	38: {
		Name: LocalisedText{
              En: "The Aetherochemical Research Facility",
              Ja: "蒼天聖戦 魔科学研究所",
              De: "Ätherochemisches Forschungslabor",
              Fr: "Le Laboratoire de magismologie",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	39: {
		Name: LocalisedText{
              En: "The Aery",
              Ja: "邪竜血戦 ドラゴンズエアリー",
              De: "Nest des Drachen",
              Fr: "L'Aire",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	40: {
		Name: LocalisedText{
              En: "Pharos Sirius (Hard)",
              Ja: "制圧巨塔 シリウス大灯台 (Hard)",
              De: "Pharos Sirius (schwer)",
              Fr: "Le Phare de Sirius (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	41: {
		Name: LocalisedText{
              En: "Saint Mocianne's Arboretum",
              Ja: "草木庭園 聖モシャーヌ植物園",
              De: "Sankt Mocianne-Arboretum",
              Fr: "L'Arboretum Sainte-Mocianne",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	42: {
		Name: LocalisedText{
              En: "Basic Training: Enemy Parties",
              Ja: "集団戦訓練をくぐり抜けろ！",
              De: "Einer für alle, alle für einen",
              Fr: "Entraînement: groupes d'ennemis",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	43: {
		Name: LocalisedText{
              En: "Under the Armor",
              Ja: "彷徨う死霊を討て！",
              De: "Bockmanns Gefolge",
              Fr: "Chasse au fantôme fantoche",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	44: {
		Name: LocalisedText{
              En: "Basic Training: Enemy Strongholds",
              Ja: "全関門を突破し、最深部の敵を討て！",
              De: "Sturmkommando",
              Fr: "Entraînement: infiltration en base ennemie",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	45: {
		Name: LocalisedText{
              En: "Hero on the Half Shell",
              Ja: "ギルガメを捕獲せよ！",
              De: "Gil oder Leben",
              Fr: "Reconquête d'une carapace escamotée",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	46: {
		Name: LocalisedText{
              En: "Pulling Poison Posies",
              Ja: "有毒妖花を駆除せよ！",
              De: "Unkraut jäten",
              Fr: "Opération fleurs du mal",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	47: {
		Name: LocalisedText{
              En: "Stinging Back",
              Ja: "無法者「似我蜂団」を撃滅せよ！",
              De: "Ins Wespennest stechen",
              Fr: "Expédition punitive contre les Ventrerouge",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	48: {
		Name: LocalisedText{
              En: "All's Well that Ends in the Well",
              Ja: "夢幻のブラキシオを討て！",
              De: "Briaxio ausschalten",
              Fr: "Briaxio à bras raccourcis",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	49: {
		Name: LocalisedText{
              En: "Flicking Sticks and Taking Names",
              Ja: "爆弾魔ゴブリン軍団を撃滅せよ！",
              De: "Bombige Goblins",
              Fr: "Les Gobelins bombardiers",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	50: {
		Name: LocalisedText{
              En: "More than a Feeler",
              Ja: "汚染源モルボルを討て！",
              De: "Tödliches Rankenspiel",
              Fr: "Sus au morbol pollueur",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	51: {
		Name: LocalisedText{
              En: "Annoy the Void",
              Ja: "坑道に現れた妖異ブソを討て！",
              De: "Gefahr aus dem Nichts",
              Fr: "Buso l'immolateur",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	52: {
		Name: LocalisedText{
              En: "Shadow and Claw",
              Ja: "無敵の眷属を従えし、大型妖異を討て！",
              De: "Kampf gegen Schatten",
              Fr: "Ombres et griffes",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	53: {
		Name: LocalisedText{
              En: "Long Live the Queen",
              Ja: "ボムを率いる「ボムクイーン」を討て！",
              De: "Miss Bombastic",
              Fr: "Longue vie à la Reine",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	54: {
		Name: LocalisedText{
              En: "Ward Up",
              Ja: "不気味な陣形を組む妖異をせん滅せよ！",
              De: "Unzertrennlich",
              Fr: "Quintettes infernaux",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	55: {
		Name: LocalisedText{
              En: "Solemn Trinity",
              Ja: "三つ巴の巨人族を制し、遺物を守れ！",
              De: "Wuchtige Dreifaltigkeit",
              Fr: "Trinité sinistre",
            },
		HighEnd: false,
		ContentKind: ContentKindGuildhests,
	},
	56: {
		Name: LocalisedText{
              En: "The Bowl of Embers",
              Ja: "イフリート討伐戦",
              De: "Das Grab der Lohe",
              Fr: "Le Cratère des tisons",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	57: {
		Name: LocalisedText{
              En: "The Navel",
              Ja: "タイタン討伐戦",
              De: "Der Nabel",
              Fr: "Le Nombril",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	58: {
		Name: LocalisedText{
              En: "The Howling Eye",
              Ja: "ガルーダ討伐戦",
              De: "Das Tosende Auge",
              Fr: "Hurlœil",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	59: {
		Name: LocalisedText{
              En: "The Bowl of Embers (Hard)",
              Ja: "真イフリート討滅戦",
              De: "Götterdämmerung - Ifrit",
              Fr: "Le Cratère des tisons (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	60: {
		Name: LocalisedText{
              En: "The Navel (Hard)",
              Ja: "真タイタン討滅戦",
              De: "Götterdämmerung - Titan",
              Fr: "Le Nombril (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	61: {
		Name: LocalisedText{
              En: "The Howling Eye (Hard)",
              Ja: "真ガルーダ討滅戦",
              De: "Götterdämmerung - Garuda",
              Fr: "Hurlœil (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	63: {
		Name: LocalisedText{
              En: "The Bowl of Embers (Extreme)",
              Ja: "極イフリート討滅戦",
              De: "Zenit der Götter - Ifrit",
              Fr: "Le Cratère des tisons (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	64: {
		Name: LocalisedText{
              En: "The Navel (Extreme)",
              Ja: "極タイタン討滅戦",
              De: "Zenit der Götter - Titan",
              Fr: "Le Nombril (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	65: {
		Name: LocalisedText{
              En: "The Howling Eye (Extreme)",
              Ja: "極ガルーダ討滅戦",
              De: "Zenit der Götter - Garuda",
              Fr: "Hurlœil (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	66: {
		Name: LocalisedText{
              En: "Thornmarch (Hard)",
              Ja: "善王モグル・モグXII世討滅戦",
              De: "Königliche Konfrontation (schwer)",
              Fr: "La Lisière de ronces (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	67: {
		Name: LocalisedText{
              En: "Thornmarch (Extreme)",
              Ja: "極王モグル・モグXII世討滅戦",
              De: "Königliche Konfrontation (extrem)",
              Fr: "La Lisière de ronces (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	68: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Ultima's Bane",
              Ja: "究極幻想 アルテマウェポン破壊作戦",
              De: "Heldenlied von Ultima",
              Fr: "Le fléau d'Ultima",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	69: {
		Name: LocalisedText{
              En: "The Gilded Araya",
              Ja: "アスラ討滅戦",
              De: "Prophetie - Asura",
              Fr: "Le temple doré d'Araya",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	70: {
		Name: LocalisedText{
              En: "Special Event I",
              Ja: "イベント用コンテンツ：1",
              De: "Event-Inhalt 1",
              Fr: "Défi spécial I",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	71: {
		Name: LocalisedText{
              En: "Special Event II",
              Ja: "イベント用コンテンツ：2",
              De: "Event-Inhalt 2",
              Fr: "Défi spécial II",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	72: {
		Name: LocalisedText{
              En: "The Whorleater (Hard)",
              Ja: "真リヴァイアサン討滅戦",
              De: "Götterdämmerung - Leviathan",
              Fr: "Le Briseur de marées (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	73: {
		Name: LocalisedText{
              En: "The Whorleater (Extreme)",
              Ja: "極リヴァイアサン討滅戦",
              De: "Zenit der Götter - Leviathan",
              Fr: "Le Briseur de marées (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	74: {
		Name: LocalisedText{
              En: "A Relic Reborn: the Chimera",
              Ja: "ドルムキマイラ討伐戦",
              De: "Kampf gegen die Dhorme-Chimära",
              Fr: "La chimère dhorme du Coerthas",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	75: {
		Name: LocalisedText{
              En: "A Relic Reborn: the Hydra",
              Ja: "ハイドラ討伐戦",
              De: "Kampf gegen die Hydra",
              Fr: "L'hydre d'Halatali",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	76: {
		Name: LocalisedText{
              En: "Battle on the Big Bridge",
              Ja: "ギルガメッシュ討伐戦",
              De: "Duell auf der großen Brücke",
              Fr: "Affrontement sur le grand pont",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	77: {
		Name: LocalisedText{
              En: "The Striking Tree (Hard)",
              Ja: "真ラムウ討滅戦",
              De: "Götterdämmerung - Ramuh",
              Fr: "L'Arbre du jugement (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	78: {
		Name: LocalisedText{
              En: "The Striking Tree (Extreme)",
              Ja: "極ラムウ討滅戦",
              De: "Zenit der Götter - Ramuh",
              Fr: "L'Arbre du jugement (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	79: {
		Name: LocalisedText{
              En: "The Akh Afah Amphitheatre (Hard)",
              Ja: "真シヴァ討滅戦",
              De: "Götterdämmerung - Shiva",
              Fr: "L'Amphithéâtre d'Akh Afah (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	80: {
		Name: LocalisedText{
              En: "The Akh Afah Amphitheatre (Extreme)",
              Ja: "極シヴァ討滅戦",
              De: "Zenit der Götter - Shiva",
              Fr: "L'Amphithéâtre d'Akh Afah (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	81: {
		Name: LocalisedText{
              En: "The Dragon's Neck",
              Ja: "アマジナ杯闘技会決勝戦",
              De: "Das Drachenhals-Kolosseum",
              Fr: "Le Col du dragon",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	82: {
		Name: LocalisedText{
              En: "Urth's Fount",
              Ja: "闘神オーディン討滅戦",
              De: "Jenseits Urths Quelle",
              Fr: "La Fontaine d'Urth",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	84: {
		Name: LocalisedText{
              En: "The Chrysalis",
              Ja: "アシエン・ナプリアレス討伐戦",
              De: "Chrysalis",
              Fr: "La Chrysalide",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	85: {
		Name: LocalisedText{
              En: "Battle in the Big Keep",
              Ja: "真ギルガメッシュ討滅戦",
              De: "Revanche in den Ruinen",
              Fr: "Revanche au vieux château",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	86: {
		Name: LocalisedText{
              En: "Thok ast Thok (Hard)",
              Ja: "真ラーヴァナ討滅戦",
              De: "Götterdämmerung - Ravana",
              Fr: "Thok ast Thok (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	87: {
		Name: LocalisedText{
              En: "Thok ast Thok (Extreme)",
              Ja: "極ラーヴァナ討滅戦",
              De: "Zenit der Götter - Ravana",
              Fr: "Thok ast Thok (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	88: {
		Name: LocalisedText{
              En: "The Limitless Blue (Hard)",
              Ja: "真ビスマルク討滅戦",
              De: "Götterdämmerung - Bismarck",
              Fr: "L'Immensité bleue (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	89: {
		Name: LocalisedText{
              En: "The Limitless Blue (Extreme)",
              Ja: "極ビスマルク討滅戦",
              De: "Zenit der Götter - Bismarck",
              Fr: "L'Immensité bleue (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	90: {
		Name: LocalisedText{
              En: "The Singularity Reactor",
              Ja: "ナイツ・オブ・ラウンド討滅戦",
              De: "Singularitäts-Reaktor",
              Fr: "Le Réacteur de singularité",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	91: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Thordan's Reign",
              Ja: "蒼天幻想 ナイツ・オブ・ラウンド討滅戦",
              De: "Heldenlied von Thordans Fall",
              Fr: "Le règne de Thordan",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	92: {
		Name: LocalisedText{
              En: "The Labyrinth of the Ancients",
              Ja: "クリスタルタワー：古代の民の迷宮",
              De: "Kristallturm - Das Labyrinth der Alten",
              Fr: "La Tour de Cristal - Dédale antique",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	93: {
		Name: LocalisedText{
              En: "The Binding Coil of Bahamut - Turn 1",
              Ja: "大迷宮バハムート：邂逅編1",
              De: "Verschlungene Schatten 1",
              Fr: "Le Labyrinthe de Bahamut I",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	94: {
		Name: LocalisedText{
              En: "The Binding Coil of Bahamut - Turn 2",
              Ja: "大迷宮バハムート：邂逅編2",
              De: "Verschlungene Schatten 2",
              Fr: "Le Labyrinthe de Bahamut II",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	95: {
		Name: LocalisedText{
              En: "The Binding Coil of Bahamut - Turn 3",
              Ja: "大迷宮バハムート：邂逅編3",
              De: "Verschlungene Schatten 3",
              Fr: "Le Labyrinthe de Bahamut III",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	96: {
		Name: LocalisedText{
              En: "The Binding Coil of Bahamut - Turn 4",
              Ja: "大迷宮バハムート：邂逅編4",
              De: "Verschlungene Schatten 4",
              Fr: "Le Labyrinthe de Bahamut IV",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	97: {
		Name: LocalisedText{
              En: "The Binding Coil of Bahamut - Turn 5",
              Ja: "大迷宮バハムート：邂逅編5",
              De: "Verschlungene Schatten 5",
              Fr: "Le Labyrinthe de Bahamut V",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	98: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut - Turn 1",
              Ja: "大迷宮バハムート：侵攻編1",
              De: "Verschlungene Schatten 2 - 1",
              Fr: "Les Méandres de Bahamut I",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	99: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut - Turn 2",
              Ja: "大迷宮バハムート：侵攻編2",
              De: "Verschlungene Schatten 2 - 2",
              Fr: "Les Méandres de Bahamut II",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	100: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut - Turn 3",
              Ja: "大迷宮バハムート：侵攻編3",
              De: "Verschlungene Schatten 2 - 3",
              Fr: "Les Méandres de Bahamut III",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	101: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut - Turn 4",
              Ja: "大迷宮バハムート：侵攻編4",
              De: "Verschlungene Schatten 2 - 4",
              Fr: "Les Méandres de Bahamut IV",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	102: {
		Name: LocalisedText{
              En: "Syrcus Tower",
              Ja: "クリスタルタワー：シルクスの塔",
              De: "Kristallturm - Der Syrcus-Turm",
              Fr: "La Tour de Cristal - Tour de Syrcus",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	103: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut (Savage) - Turn 1",
              Ja: "大迷宮バハムート零式：侵攻編1",
              De: "Verschlungene Schatten 2 - 1 (episch)",
              Fr: "Les Méandres de Bahamut I (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	104: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut (Savage) - Turn 2",
              Ja: "大迷宮バハムート零式：侵攻編2",
              De: "Verschlungene Schatten 2 - 2 (episch)",
              Fr: "Les Méandres de Bahamut II (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	105: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut (Savage) - Turn 3",
              Ja: "大迷宮バハムート零式：侵攻編3",
              De: "Verschlungene Schatten 2 - 3 (episch)",
              Fr: "Les Méandres de Bahamut III (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	106: {
		Name: LocalisedText{
              En: "The Second Coil of Bahamut (Savage) - Turn 4",
              Ja: "大迷宮バハムート零式：侵攻編4",
              De: "Verschlungene Schatten 2 - 4 (episch)",
              Fr: "Les Méandres de Bahamut IV (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	107: {
		Name: LocalisedText{
              En: "The Final Coil of Bahamut - Turn 1",
              Ja: "大迷宮バハムート：真成編1",
              De: "Verschlungene Schatten 3 - 1",
              Fr: "L'Abîme de Bahamut I",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	108: {
		Name: LocalisedText{
              En: "The Final Coil of Bahamut - Turn 2",
              Ja: "大迷宮バハムート：真成編2",
              De: "Verschlungene Schatten 3 - 2",
              Fr: "L'Abîme de Bahamut II",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	109: {
		Name: LocalisedText{
              En: "The Final Coil of Bahamut - Turn 3",
              Ja: "大迷宮バハムート：真成編3",
              De: "Verschlungene Schatten 3 - 3",
              Fr: "L'Abîme de Bahamut III",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	110: {
		Name: LocalisedText{
              En: "The Final Coil of Bahamut - Turn 4",
              Ja: "大迷宮バハムート：真成編4",
              De: "Verschlungene Schatten 3 - 4",
              Fr: "L'Abîme de Bahamut IV",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	111: {
		Name: LocalisedText{
              En: "The World of Darkness",
              Ja: "クリスタルタワー：闇の世界",
              De: "Die Welt der Dunkelheit",
              Fr: "La Tour de Cristal - Monde des Ténèbres",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	112: {
		Name: LocalisedText{
              En: "Alexander - The Fist of the Father",
              Ja: "機工城アレキサンダー：起動編1",
              De: "Alexander - Faust des Vaters",
              Fr: "Alexander - Le Poing du Père",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	113: {
		Name: LocalisedText{
              En: "Alexander - The Cuff of the Father",
              Ja: "機工城アレキサンダー：起動編2",
              De: "Alexander - Elle des Vaters",
              Fr: "Alexander - Le Poignet du Père",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	114: {
		Name: LocalisedText{
              En: "Alexander - The Arm of the Father",
              Ja: "機工城アレキサンダー：起動編3",
              De: "Alexander - Arm des Vaters",
              Fr: "Alexander - Le Bras du Père",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	115: {
		Name: LocalisedText{
              En: "Alexander - The Burden of the Father",
              Ja: "機工城アレキサンダー：起動編4",
              De: "Alexander - Last des Vaters",
              Fr: "Alexander - Le Fardeau du Père",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	116: {
		Name: LocalisedText{
              En: "Alexander - The Fist of the Father (Savage)",
              Ja: "機工城アレキサンダー零式：起動編1",
              De: "Alexander - Faust des Vaters (episch)",
              Fr: "Alexander - Le Poing du Père (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	117: {
		Name: LocalisedText{
              En: "Alexander - The Cuff of the Father (Savage)",
              Ja: "機工城アレキサンダー零式：起動編2",
              De: "Alexander - Elle des Vaters (episch)",
              Fr: "Alexander - Le Poignet du Père (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	118: {
		Name: LocalisedText{
              En: "Alexander - The Arm of the Father (Savage)",
              Ja: "機工城アレキサンダー零式：起動編3",
              De: "Alexander - Arm des Vaters (episch)",
              Fr: "Alexander - Le Bras du Père (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	119: {
		Name: LocalisedText{
              En: "Alexander - The Burden of the Father (Savage)",
              Ja: "機工城アレキサンダー零式：起動編4",
              De: "Alexander - Last des Vaters (episch)",
              Fr: "Alexander - Le Fardeau du Père (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	120: {
		Name: LocalisedText{
              En: "The Void Ark",
              Ja: "魔航船ヴォイドアーク",
              De: "Die Nichts-Arche",
              Fr: "L'Arche du néant",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	127: {
		Name: LocalisedText{
              En: "The Borderland Ruins (Secure)",
              Ja: "外縁遺跡群 (制圧戦)",
              De: "Äußere Ruinen (Sicherung)",
              Fr: "Les Ruines frontalières (annexion)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	130: {
		Name: LocalisedText{
              En: "Seal Rock (Seize)",
              Ja: "シールロック (争奪戦)",
              De: "Robbenholm (Eroberung)",
              Fr: "Le Rocher des tréfonds (invasion)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	131: {
		Name: LocalisedText{
              En: "The Diadem (Easy)",
              Ja: "雲海探索 ディアデム諸島 (Easy)",
              De: "Das Diadem (leicht)",
              Fr: "Le Diadème (facile)",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	132: {
		Name: LocalisedText{
              En: "The Diadem",
              Ja: "雲海探索 ディアデム諸島",
              De: "Das Diadem",
              Fr: "Le Diadème",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	133: {
		Name: LocalisedText{
              En: "The Diadem (Hard)",
              Ja: "雲海探索 ディアデム諸島 (Hard)",
              De: "Das Diadem (schwer)",
              Fr: "Le Diadème (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	134: {
		Name: LocalisedText{
              En: "Containment Bay S1T7",
              Ja: "魔神セフィロト討滅戦",
              De: "Götterdämmerung - Sephirot",
              Fr: "Unité de contention S1P7",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	135: {
		Name: LocalisedText{
              En: "Containment Bay S1T7 (Extreme)",
              Ja: "極魔神セフィロト討滅戦",
              De: "Zenit der Götter - Sephirot",
              Fr: "Unité de contention S1P7 (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	136: {
		Name: LocalisedText{
              En: "Alexander - The Fist of the Son",
              Ja: "機工城アレキサンダー：律動編1",
              De: "Alexander - Faust des Sohnes",
              Fr: "Alexander - Le Poing du Fils",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	137: {
		Name: LocalisedText{
              En: "Alexander - The Cuff of the Son",
              Ja: "機工城アレキサンダー：律動編2",
              De: "Alexander - Elle des Sohnes",
              Fr: "Alexander - Le Poignet du Fils",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	138: {
		Name: LocalisedText{
              En: "Alexander - The Arm of the Son",
              Ja: "機工城アレキサンダー：律動編3",
              De: "Alexander - Arm des Sohnes",
              Fr: "Alexander - Le Bras du Fils",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	139: {
		Name: LocalisedText{
              En: "Alexander - The Burden of the Son",
              Ja: "機工城アレキサンダー：律動編4",
              De: "Alexander - Last des Sohnes",
              Fr: "Alexander - Le Fardeau du Fils",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	140: {
		Name: LocalisedText{
              En: "The Lost City of Amdapor (Hard)",
              Ja: "神聖遺跡 古アムダプール市街 (Hard)",
              De: "Historisches Amdapor (schwer)",
              Fr: "Les Vestiges de la cité d'Amdapor (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	141: {
		Name: LocalisedText{
              En: "The Antitower",
              Ja: "星海観測 逆さの塔",
              De: "Antiturm",
              Fr: "L'Antitour",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	147: {
		Name: LocalisedText{
              En: "Alexander - The Fist of the Son (Savage)",
              Ja: "機工城アレキサンダー零式：律動編1",
              De: "Alexander - Faust des Sohnes (episch)",
              Fr: "Alexander - Le Poing du Fils (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	148: {
		Name: LocalisedText{
              En: "Alexander - The Cuff of the Son (Savage)",
              Ja: "機工城アレキサンダー零式：律動編2",
              De: "Alexander - Elle des Sohnes (episch)",
              Fr: "Alexander - Le Poignet du Fils (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	149: {
		Name: LocalisedText{
              En: "Alexander - The Arm of the Son (Savage)",
              Ja: "機工城アレキサンダー零式：律動編3",
              De: "Alexander - Arm des Sohnes (episch)",
              Fr: "Alexander - Le Bras du Fils (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	150: {
		Name: LocalisedText{
              En: "Alexander - The Burden of the Son (Savage)",
              Ja: "機工城アレキサンダー零式：律動編4",
              De: "Alexander - Last des Sohnes (episch)",
              Fr: "Alexander - Le Fardeau du Fils (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	151: {
		Name: LocalisedText{
              En: "Avoid Area of Effect Attacks",
              Ja: "範囲攻撃を避けよう！",
              De: "Flächenangriffen ausweichen",
              Fr: "Éviter les attaques à aire d'effet",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	152: {
		Name: LocalisedText{
              En: "Execute a Combo to Increase Enmity",
              Ja: "コンボで敵視を集めよう！",
              De: "Mit Kombos Feindseligkeit auf sich ziehen",
              Fr: "Générer de l'inimitié avec un combo",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	153: {
		Name: LocalisedText{
              En: "Execute a Combo in Battle",
              Ja: "実戦でコンボに挑戦しよう！",
              De: "Kombos im Kampf einsetzen",
              Fr: "Effectuer le combo en combat",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	154: {
		Name: LocalisedText{
              En: "Accrue Enmity from Multiple Targets",
              Ja: "複数の敵から敵視を集めよう！",
              De: "Feindseligkeit mehrerer Gegner auf sich ziehen",
              Fr: "Attirer l'inimitié de plusieurs ennemis sur soi",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	155: {
		Name: LocalisedText{
              En: "Engage Multiple Targets",
              Ja: "実戦で複数の敵と戦ってみよう！",
              De: "Gegen mehrere Gegner auf einmal kämpfen",
              Fr: "Affronter plusieurs ennemis",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	156: {
		Name: LocalisedText{
              En: "Execute a Ranged Attack to Increase Enmity",
              Ja: "遠距離から敵視を集めよう！",
              De: "Aus der Ferne Feindseligkeit auf sich ziehen",
              Fr: "Générer de l'inimitié à distance",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	157: {
		Name: LocalisedText{
              En: "Engage Enemy Reinforcements",
              Ja: "敵の増援に対応しよう！",
              De: "Feindliche Verstärkung aufhalten",
              Fr: "Faire face à des renforts ennemis",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	158: {
		Name: LocalisedText{
              En: "Assist Allies in Defeating a Target",
              Ja: "味方と協力して敵を倒そう！",
              De: "Gegner gemeinsam besiegen",
              Fr: "Vaincre un ennemi en assistant des alliés",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	159: {
		Name: LocalisedText{
              En: "Defeat an Occupied Target",
              Ja: "味方が引きつけている敵を倒そう！",
              De: "Den Gegner eines Verbündeten besiegen",
              Fr: "Vaincre un ennemi occupé par un allié",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	160: {
		Name: LocalisedText{
              En: "Avoid Engaged Targets",
              Ja: "敵の攻撃を避けながら戦おう！",
              De: "Angriffen ausweichen",
              Fr: "Combattre en évitant les attaques ennemies",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	161: {
		Name: LocalisedText{
              En: "Engage Enemy Reinforcements",
              Ja: "敵の増援に対応しよう！",
              De: "Feindliche Verstärkung aufhalten",
              Fr: "Éliminer les renforts ennemis",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	162: {
		Name: LocalisedText{
              En: "Interact with the Battlefield",
              Ja: "ギミックを活用して戦おう！",
              De: "Mit dem Gelände interagieren",
              Fr: "Interagir avec le décor en combat",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	163: {
		Name: LocalisedText{
              En: "Heal an Ally",
              Ja: "味方を回復しよう！",
              De: "Verbündete heilen",
              Fr: "Soigner un allié",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	164: {
		Name: LocalisedText{
              En: "Heal Multiple Allies",
              Ja: "複数の味方を回復しよう！",
              De: "Mehrere Verbündete heilen",
              Fr: "Soigner plusieurs alliés",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	165: {
		Name: LocalisedText{
              En: "Avoid Engaged Targets",
              Ja: "敵の攻撃を避けながら戦おう！",
              De: "Angriffen ausweichen",
              Fr: "Combattre en évitant les attaques ennemies",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	166: {
		Name: LocalisedText{
              En: "Final Exercise",
              Ja: "最終訓練！",
              De: "Letzte Übung",
              Fr: "Exercice final",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	167: {
		Name: LocalisedText{
              En: "A Spectacle for the Ages",
              Ja: "四国合同演習",
              De: "Truppenübung der Eorzäischen Allianz",
              Fr: "La grande manœuvre éorzéenne",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	168: {
		Name: LocalisedText{
              En: "The Weeping City of Mhach",
              Ja: "禁忌都市マハ",
              De: "Die Stadt der Tränen",
              Fr: "La Cité défendue de Mhach",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	169: {
		Name: LocalisedText{
              En: "The Final Steps of Faith",
              Ja: "ニーズヘッグ征竜戦",
              De: "Der letzte Schicksalsweg",
              Fr: "La Dernière avancée de la Foi",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	170: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Nidhogg's Rage",
              Ja: "極ニーズヘッグ征竜戦",
              De: "Das Lied von Nidhoggs letztem Ruf",
              Fr: "L'ire de Nidhogg",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	171: {
		Name: LocalisedText{
              En: "Sohr Khai",
              Ja: "天竜宮殿 ソール・カイ",
              De: "Sohr Khai",
              Fr: "Sohr Khai",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	172: {
		Name: LocalisedText{
              En: "Hullbreaker Isle (Hard)",
              Ja: "黒渦伝説 ハルブレーカー・アイル (Hard)",
              De: "Schiffbrecher-Insel (schwer)",
              Fr: "L'Île de Crèvecarène (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	173: {
		Name: LocalisedText{
              En: "A Bloody Reunion",
              Ja: "レグラ・ヴァン・ヒュドルス追撃戦",
              De: "Blutiges Wiedersehen",
              Fr: "Course-poursuite dans le laboratoire",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	174: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 1-10)",
              Ja: "死者の宮殿 B1～B10",
              De: "Palast der Toten (Ebenen 1-10)",
              Fr: "Le Palais des morts (sous-sols 1-10)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	175: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 11-20)",
              Ja: "死者の宮殿 B11～B20",
              De: "Palast der Toten (Ebenen 11-20)",
              Fr: "Le Palais des morts (sous-sols 11-20)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	176: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 21-30)",
              Ja: "死者の宮殿 B21～B30",
              De: "Palast der Toten (Ebenen 21-30)",
              Fr: "Le Palais des morts (sous-sols 21-30)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	177: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 31-40)",
              Ja: "死者の宮殿 B31～B40",
              De: "Palast der Toten (Ebenen 31-40)",
              Fr: "Le Palais des morts (sous-sols 31-40)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	178: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 41-50)",
              Ja: "死者の宮殿 B41～B50",
              De: "Palast der Toten (Ebenen 41-50)",
              Fr: "Le Palais des morts (sous-sols 41-50)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	179: {
		Name: LocalisedText{
              En: "The Aquapolis",
              Ja: "宝物庫 アクアポリス",
              De: "Aquapolis",
              Fr: "L'Aquapole",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	180: {
		Name: LocalisedText{
              En: "The Fields of Glory (Shatter)",
              Ja: "フィールド・オブ・グローリー (砕氷戦)",
              De: "Feld der Ehre (Zersplitterung)",
              Fr: "Les Champs de la Gloire (brise-glace)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	181: {
		Name: LocalisedText{
              En: "The Haunted Manor",
              Ja: "亡霊屋敷 ホーンテッドマナー",
              De: "Das Geisterschloss",
              Fr: "Le Manoir hanté",
            },
		HighEnd: false,
		ContentKind: ContentKind(22),
	},
	182: {
		Name: LocalisedText{
              En: "Xelphatol",
              Ja: "峻厳渓谷 ゼルファトル",
              De: "Xelphatol",
              Fr: "Xelphatol",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	183: {
		Name: LocalisedText{
              En: "Containment Bay P1T6",
              Ja: "女神ソフィア討滅戦",
              De: "Götterdämmerung - Sophia",
              Fr: "Unité de contention P1P6",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	184: {
		Name: LocalisedText{
              En: "Containment Bay P1T6 (Extreme)",
              Ja: "極女神ソフィア討滅戦",
              De: "Zenit der Götter - Sophia",
              Fr: "Unité de contention P1P6 (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	186: {
		Name: LocalisedText{
              En: "Alexander - The Eyes of the Creator",
              Ja: "機工城アレキサンダー：天動編1",
              De: "Alexander - Augen des Schöpfers",
              Fr: "Alexander - Les Yeux du Créateur",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	187: {
		Name: LocalisedText{
              En: "Alexander - The Breath of the Creator",
              Ja: "機工城アレキサンダー：天動編2",
              De: "Alexander - Atem des Schöpfers",
              Fr: "Alexander - Le Souffle du Créateur",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	188: {
		Name: LocalisedText{
              En: "Alexander - The Heart of the Creator",
              Ja: "機工城アレキサンダー：天動編3",
              De: "Alexander - Herz des Schöpfers",
              Fr: "Alexander - Le Cœur du Créateur",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	189: {
		Name: LocalisedText{
              En: "Alexander - The Soul of the Creator",
              Ja: "機工城アレキサンダー：天動編4",
              De: "Alexander - Seele des Schöpfers",
              Fr: "Alexander - L'Âme du Créateur",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	190: {
		Name: LocalisedText{
              En: "Alexander - The Eyes of the Creator (Savage)",
              Ja: "機工城アレキサンダー零式：天動編1",
              De: "Alexander - Augen des Schöpfers (episch)",
              Fr: "Alexander - Les Yeux du Créateur (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	191: {
		Name: LocalisedText{
              En: "Alexander - The Breath of the Creator (Savage)",
              Ja: "機工城アレキサンダー零式：天動編2",
              De: "Alexander - Atem des Schöpfers (episch)",
              Fr: "Alexander - Le Souffle du Créateur (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	192: {
		Name: LocalisedText{
              En: "Alexander - The Heart of the Creator (Savage)",
              Ja: "機工城アレキサンダー零式：天動編3",
              De: "Alexander - Herz des Schöpfers (episch)",
              Fr: "Alexander - Le Cœur du Créateur (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	193: {
		Name: LocalisedText{
              En: "Alexander - The Soul of the Creator (Savage)",
              Ja: "機工城アレキサンダー零式：天動編4",
              De: "Alexander - Seele des Schöpfers (episch)",
              Fr: "Alexander - L'Âme du Créateur (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	194: {
		Name: LocalisedText{
              En: "One Life for One World",
              Ja: "絡み合う宿命",
              De: "Weltenübergreifendes Schicksal",
              Fr: "Destins entrecroisés",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	195: {
		Name: LocalisedText{
              En: "The Triple Triad Battlehall",
              Ja: "トリプルトライアド：カードバトルルーム",
              De: "Triple Triad: Weltensalon",
              Fr: "Arène Triple Triade",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	196: {
		Name: LocalisedText{
              En: "The Great Gubal Library (Hard)",
              Ja: "稀書回収 グブラ幻想図書館 (Hard)",
              De: "Große Gubal-Bibliothek (schwer)",
              Fr: "La Grande bibliothèque de Gubal (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	197: {
		Name: LocalisedText{
              En: "LoVM: Player Battle (RP)",
              Ja: "LoVM：プレイヤー対戦 (RP変動あり)",
              De: "Kampf der Trabanten: Gegen Spieler (um RP)",
              Fr: "Bataille simple contre un joueur (avec PR)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	198: {
		Name: LocalisedText{
              En: "LoVM: Tournament",
              Ja: "LoVM：大会対戦 (プレイヤー対戦）",
              De: "Kampf der Trabanten: Turnier (gegen Spieler)",
              Fr: "Bataille de tournoi contre des joueurs",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	199: {
		Name: LocalisedText{
              En: "LoVM: Player Battle (Non-RP)",
              Ja: "LoVM：プレイヤー対戦 (RP変動なし)",
              De: "Kampf der Trabanten: Gegen Spieler (ohne RP)",
              Fr: "Bataille simple contre un joueur (sans PR)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	202: {
		Name: LocalisedText{
              En: "The Diadem Hunting Grounds (Easy)",
              Ja: "雲海探索 ディアデム諸島：狩猟限定 (Easy)",
              De: "Das Diadem - Jagdgründe (leicht)",
              Fr: "Le Diadème: terrains de chasse (facile)",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	203: {
		Name: LocalisedText{
              En: "The Diadem Hunting Grounds",
              Ja: "雲海探索 ディアデム諸島：狩猟限定",
              De: "Das Diadem - Jagdgründe",
              Fr: "Le Diadème: terrains de chasse",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	204: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 51-60)",
              Ja: "死者の宮殿 B51～B60",
              De: "Palast der Toten (Ebenen 51 - 60)",
              Fr: "Le Palais des morts (sous-sols 51-60)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	205: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 61-70)",
              Ja: "死者の宮殿 B61～B70",
              De: "Palast der Toten (Ebenen 61 - 70)",
              Fr: "Le Palais des morts (sous-sols 61-70)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	206: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 71-80)",
              Ja: "死者の宮殿 B71～B80",
              De: "Palast der Toten (Ebenen 71 - 80)",
              Fr: "Le Palais des morts (sous-sols 71-80)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	207: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 81-90)",
              Ja: "死者の宮殿 B81～B90",
              De: "Palast der Toten (Ebenen 81 - 90)",
              Fr: "Le Palais des morts (sous-sols 81-90)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	208: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 91-100)",
              Ja: "死者の宮殿 B91～B100",
              De: "Palast der Toten (Ebenen 91 - 100)",
              Fr: "Le Palais des morts (sous-sols 91-100)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	209: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 101-110)",
              Ja: "死者の宮殿 B101～B110",
              De: "Palast der Toten (Ebenen 101 - 110)",
              Fr: "Le Palais des morts (sous-sols 101-110)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	210: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 111-120)",
              Ja: "死者の宮殿 B111～B120",
              De: "Palast der Toten (Ebenen 111 - 120)",
              Fr: "Le Palais des morts (sous-sols 111-120)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	211: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 121-130)",
              Ja: "死者の宮殿 B121～B130",
              De: "Palast der Toten (Ebenen 121 - 130)",
              Fr: "Le Palais des morts (sous-sols 121-130)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	212: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 131-140)",
              Ja: "死者の宮殿 B131～B140",
              De: "Palast der Toten (Ebenen 131 - 140)",
              Fr: "Le Palais des morts (sous-sols 131-140)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	213: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 141-150)",
              Ja: "死者の宮殿 B141～B150",
              De: "Palast der Toten (Ebenen 141 - 150)",
              Fr: "Le Palais des morts (sous-sols 141-150)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	214: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 151-160)",
              Ja: "死者の宮殿 B151～B160",
              De: "Palast der Toten (Ebenen 151 - 160)",
              Fr: "Le Palais des morts (sous-sols 151-160)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	215: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 161-170)",
              Ja: "死者の宮殿 B161～B170",
              De: "Palast der Toten (Ebenen 161 - 170)",
              Fr: "Le Palais des morts (sous-sols 161-170)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	216: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 171-180)",
              Ja: "死者の宮殿 B171～B180",
              De: "Palast der Toten (Ebenen 171 - 180)",
              Fr: "Le Palais des morts (sous-sols 171-180)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	217: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 181-190)",
              Ja: "死者の宮殿 B181～B190",
              De: "Palast der Toten (Ebenen 181 - 190)",
              Fr: "Le Palais des morts (sous-sols 181-190)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	218: {
		Name: LocalisedText{
              En: "The Palace of the Dead (Floors 191-200)",
              Ja: "死者の宮殿 B191～B200",
              De: "Palast der Toten (Ebenen 191 - 200)",
              Fr: "Le Palais des morts (sous-sols 191-200)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	219: {
		Name: LocalisedText{
              En: "Baelsar's Wall",
              Ja: "巨大防壁 バエサルの長城",
              De: "Baelsar-Wall",
              Fr: "La Muraille de Baelsar",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	220: {
		Name: LocalisedText{
              En: "Dun Scaith",
              Ja: "影の国ダン・スカー",
              De: "Dun Scaith",
              Fr: "Dun Scaith",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	221: {
		Name: LocalisedText{
              En: "Sohm Al (Hard)",
              Ja: "霊峰浄化 ソーム・アル (Hard)",
              De: "Sohm Al (schwer)",
              Fr: "Sohm Al (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	222: {
		Name: LocalisedText{
              En: "The Carteneau Flats: Heliodrome",
              Ja: "カルテノー平原遭遇戦",
              De: "Heliodrom",
              Fr: "Rixe à l'Héliodrome",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	223: {
		Name: LocalisedText{
              En: "Containment Bay Z1T9",
              Ja: "鬼神ズルワーン討滅戦",
              De: "Götterdämmerung - Zurvan",
              Fr: "Unité de contention Z1P9",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	224: {
		Name: LocalisedText{
              En: "Containment Bay Z1T9 (Extreme)",
              Ja: "極鬼神ズルワーン討滅戦",
              De: "Zenit der Götter - Zurvan",
              Fr: "Unité de contention Z1P9 (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	225: {
		Name: LocalisedText{
              En: "The Diadem - Trials of the Fury",
              Ja: "雲海探索 ディアデム諸島 (狩猟)",
              De: "Das Diadem - Halones Prüfung",
              Fr: "Le Diadème - Épreuves de Halone",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	234: {
		Name: LocalisedText{
              En: "The Diadem - Trials of the Matron",
              Ja: "雲海探索 ディアデム諸島 (採集)",
              De: "Das Diadem - Nophicas Prüfung",
              Fr: "Le Diadème - Épreuves de Nophica",
            },
		HighEnd: false,
		ContentKind: ContentKind(23),
	},
	235: {
		Name: LocalisedText{
              En: "Shisui of the Violet Tides",
              Ja: "海底宮殿 紫水宮",
              De: "Shisui",
              Fr: "Le Palais aux Marées violettes",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	236: {
		Name: LocalisedText{
              En: "The Temple of the Fist",
              Ja: "壊神修行 星導山寺院",
              De: "Tempel der Faust",
              Fr: "Le Temple du Poing",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	237: {
		Name: LocalisedText{
              En: "It's Probably a Trap",
              Ja: "ギョドウ現る！",
              De: "Ein zweifelhaftes Angebot",
              Fr: "Un drôle de Namazu",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	238: {
		Name: LocalisedText{
              En: "The Sirensong Sea",
              Ja: "漂流海域 セイレーン海",
              De: "Sirenen-See",
              Fr: "La Mer du Chant des sirènes",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	239: {
		Name: LocalisedText{
              En: "The Royal Menagerie",
              Ja: "神龍討滅戦",
              De: "Königliche Menagerie",
              Fr: "La Ménagerie royale",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	240: {
		Name: LocalisedText{
              En: "Bardam's Mettle",
              Ja: "伝統試練 バルダム覇道",
              De: "Bardams Probe",
              Fr: "La Force de Bardam",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	241: {
		Name: LocalisedText{
              En: "Doma Castle",
              Ja: "解放決戦 ドマ城",
              De: "Burg Doma",
              Fr: "Le Château de Doma",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	242: {
		Name: LocalisedText{
              En: "Castrum Abania",
              Ja: "巨砲要塞 カストルム・アバニア",
              De: "Castrum Abania",
              Fr: "Castrum Abania",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	243: {
		Name: LocalisedText{
              En: "The Pool of Tribute",
              Ja: "スサノオ討滅戦",
              De: "Götterdämmerung - Susano",
              Fr: "La Crique aux tributs",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	244: {
		Name: LocalisedText{
              En: "The Pool of Tribute (Extreme)",
              Ja: "極スサノオ討滅戦",
              De: "Zenit der Götter - Susano",
              Fr: "La Crique aux tributs (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	245: {
		Name: LocalisedText{
              En: "With Heart and Steel",
              Ja: "抗う力",
              De: "Die Kraft des Widerstands",
              Fr: "Transmigration démoniaque",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	246: {
		Name: LocalisedText{
              En: "Naadam",
              Ja: "終節の合戦",
              De: "Naadam",
              Fr: "La grande bataille du Naadam",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	247: {
		Name: LocalisedText{
              En: "Ala Mhigo",
              Ja: "紅蓮決戦 アラミゴ",
              De: "Ala Mhigo",
              Fr: "Ala Mhigo",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	248: {
		Name: LocalisedText{
              En: "Blood on the Deck",
              Ja: "海都を震わす人斬りの宴！",
              De: "Mord ist sein Hobby",
              Fr: "La légende de Musosai: l'assassin de Limsa Lominsa",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	249: {
		Name: LocalisedText{
              En: "The Face of True Evil",
              Ja: "極悪人コガラシ",
              De: "Der Inbegriff des Bösen",
              Fr: "L'abominable Kogarashi",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	250: {
		Name: LocalisedText{
              En: "Matsuba Mayhem",
              Ja: "松葉門外の変",
              De: "Vorfall auf dem Matsuba-Platz",
              Fr: "Règlement de compte au square Matsuba",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	251: {
		Name: LocalisedText{
              En: "The Battle on Bekko",
              Ja: "ベッコウ島の決闘",
              De: "Entscheidungsschlacht auf Bekko",
              Fr: "L'affrontement de deux justices",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	252: {
		Name: LocalisedText{
              En: "Deltascape V1.0",
              Ja: "次元の狭間オメガ：デルタ編1",
              De: "Deltametrie 1.0",
              Fr: "Deltastice v1.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	253: {
		Name: LocalisedText{
              En: "Deltascape V2.0",
              Ja: "次元の狭間オメガ：デルタ編2",
              De: "Deltametrie 2.0",
              Fr: "Deltastice v2.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	254: {
		Name: LocalisedText{
              En: "Deltascape V3.0",
              Ja: "次元の狭間オメガ：デルタ編3",
              De: "Deltametrie 3.0",
              Fr: "Deltastice v3.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	255: {
		Name: LocalisedText{
              En: "Deltascape V4.0",
              Ja: "次元の狭間オメガ：デルタ編4",
              De: "Deltametrie 4.0",
              Fr: "Deltastice v4.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	256: {
		Name: LocalisedText{
              En: "Deltascape V1.0 (Savage)",
              Ja: "次元の狭間オメガ零式：デルタ編1",
              De: "Deltametrie 1.0 (episch)",
              Fr: "Deltastice v1.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	257: {
		Name: LocalisedText{
              En: "Deltascape V2.0 (Savage)",
              Ja: "次元の狭間オメガ零式：デルタ編2",
              De: "Deltametrie 2.0 (episch)",
              Fr: "Deltastice v2.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	258: {
		Name: LocalisedText{
              En: "Deltascape V3.0 (Savage)",
              Ja: "次元の狭間オメガ零式：デルタ編3",
              De: "Deltametrie 3.0 (episch)",
              Fr: "Deltastice v3.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	259: {
		Name: LocalisedText{
              En: "Deltascape V4.0 (Savage)",
              Ja: "次元の狭間オメガ零式：デルタ編4",
              De: "Deltametrie 4.0 (episch)",
              Fr: "Deltastice v4.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	260: {
		Name: LocalisedText{
              En: "Curious Gorge Meets His Match",
              Ja: "原初的な彼女",
              De: "Die Urkraft in ihr",
              Fr: "L'épreuve de force",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	261: {
		Name: LocalisedText{
              En: "In Thal's Name",
              Ja: "ウル王杯闘技会の始まり",
              De: "Thal zu Ehren",
              Fr: "Le tournoi commémoratif du sultanat",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	262: {
		Name: LocalisedText{
              En: "Kugane Castle",
              Ja: "悪党成敗 クガネ城",
              De: "Schloss Kugane",
              Fr: "Le Château de Kugane",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	263: {
		Name: LocalisedText{
              En: "Emanation",
              Ja: "ラクシュミ討滅戦",
              De: "Götterdämmerung - Lakshmi",
              Fr: "Émanation",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	264: {
		Name: LocalisedText{
              En: "Emanation (Extreme)",
              Ja: "極ラクシュミ討滅戦",
              De: "Zenit der Götter - Lakshmi",
              Fr: "Émanation (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	265: {
		Name: LocalisedText{
              En: "Our Unsung Heroes",
              Ja: "時をかける願い",
              De: "Ein Wunsch aus alten Zeiten",
              Fr: "L'espoir en héritage",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	266: {
		Name: LocalisedText{
              En: "The Heart of the Problem",
              Ja: "燃えよゴージ！",
              De: "Kriegerische Leidenschaft",
              Fr: "Passion guerrière",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	267: {
		Name: LocalisedText{
              En: "Dark as the Night Sky",
              Ja: "漆黒の巨竜",
              De: "Der tobende Drache",
              Fr: "Aussi sombre que la nuit",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	268: {
		Name: LocalisedText{
              En: "The Lost Canals of Uznair",
              Ja: "宝物庫 ウズネアカナル",
              De: "Kanäle von Uznair",
              Fr: "Les Canaux perdus d'Uznair",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	269: {
		Name: LocalisedText{
              En: "The Resonant",
              Ja: "ウリエンジェの秘策",
              De: "Wege zur Transzendenz",
              Fr: "La ruse d'Urianger",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	270: {
		Name: LocalisedText{
              En: "Raising the Sword",
              Ja: "さらなる剣術の高みへ",
              De: "Die hohe Kunst des Schwertkampfs",
              Fr: "La finale des champions",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	271: {
		Name: LocalisedText{
              En: "The Orphans and the Broken Blade",
              Ja: "あと三度、遥かな憧憬",
              De: "Probe des Meisters",
              Fr: "L'aspiration refoulée",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	272: {
		Name: LocalisedText{
              En: "Our Compromise",
              Ja: "あと一度、君に会えたら",
              De: "Aus der Tiefe des Herzens",
              Fr: "La dernière séparation",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	273: {
		Name: LocalisedText{
              En: "Dragon Sound",
              Ja: "紅の竜騎士",
              De: "Der Rubin-Drachenreiter",
              Fr: "Le Dragon écarlate",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	274: {
		Name: LocalisedText{
              En: "When Clans Collide",
              Ja: "影隠忍法帖",
              De: "Aus dem Verborgenen",
              Fr: "La bataille des clans",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	275: {
		Name: LocalisedText{
              En: "Interdimensional Rift",
              Ja: "次元の狭間：外縁",
              De: "Interdimensionaler Riss",
              Fr: "Fissure interdimensionnelle",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	276: {
		Name: LocalisedText{
              En: "The Hidden Canals of Uznair",
              Ja: "宝物庫 ウズネアカナル深層",
              De: "Vergessene Kanäle von Uznair",
              Fr: "Les Canaux cachés d'Uznair",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	277: {
		Name: LocalisedText{
              En: "Astragalos",
              Ja: "アストラガロス (機工戦)",
              De: "Astragalos",
              Fr: "Astragalos (machinerie)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	278: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Shinryu's Domain",
              Ja: "極神龍討滅戦",
              De: "Heldenlied von Shinryu",
              Fr: "Le domaine de Shinryu",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	279: {
		Name: LocalisedText{
              En: "The Drowned City of Skalla",
              Ja: "水没遺構 スカラ",
              De: "Die versunkene Stadt Skalla",
              Fr: "La Cité engloutie de Skalla",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	280: {
		Name: LocalisedText{
              En: "The Unending Coil of Bahamut (Ultimate)",
              Ja: "絶バハムート討滅戦",
              De: "Endlose Schatten von Bahamut (fatal)",
              Fr: "L'Abîme infini de Bahamut (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	281: {
		Name: LocalisedText{
              En: "The Royal City of Rabanastre",
              Ja: "失われた都 ラバナスタ",
              De: "Rabanastre",
              Fr: "La Cité royale de Rabanastre",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	282: {
		Name: LocalisedText{
              En: "Return of the Bull",
              Ja: "英雄の帰還",
              De: "Verrat der Qalyana",
              Fr: "Retour au bercail",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	283: {
		Name: LocalisedText{
              En: "The Forbidden Land, Eureka Anemos",
              Ja: "禁断の地 エウレカ：アネモス編",
              De: "Eureka Anemos",
              Fr: "Eurêka Anemos",
            },
		HighEnd: false,
		ContentKind: ContentKindEureka,
	},
	284: {
		Name: LocalisedText{
              En: "Hells' Lid",
              Ja: "紅玉火山 獄之蓋",
              De: "Höllenspund",
              Fr: "Le Couvercle des enfers",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	285: {
		Name: LocalisedText{
              En: "The Fractal Continuum (Hard)",
              Ja: "暴走戦艦 フラクタル・コンティニアム (Hard)",
              De: "Die Fraktal-Kontinuum (schwer)",
              Fr: "Le Continuum fractal (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	286: {
		Name: LocalisedText{
              En: "Sigmascape V1.0",
              Ja: "次元の狭間オメガ：シグマ編1",
              De: "Sigmametrie 1.0",
              Fr: "Sigmastice v1.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	287: {
		Name: LocalisedText{
              En: "Sigmascape V2.0",
              Ja: "次元の狭間オメガ：シグマ編2",
              De: "Sigmametrie 2.0",
              Fr: "Sigmastice v2.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	288: {
		Name: LocalisedText{
              En: "Sigmascape V3.0",
              Ja: "次元の狭間オメガ：シグマ編3",
              De: "Sigmametrie 3.0",
              Fr: "Sigmastice v3.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	289: {
		Name: LocalisedText{
              En: "Sigmascape V4.0",
              Ja: "次元の狭間オメガ：シグマ編4",
              De: "Sigmametrie 4.0",
              Fr: "Sigmastice v4.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	290: {
		Name: LocalisedText{
              En: "The Jade Stoa",
              Ja: "白虎征魂戦",
              De: "Seelentanz - Byakko",
              Fr: "La Clairière de Jade",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	291: {
		Name: LocalisedText{
              En: "The Jade Stoa (Extreme)",
              Ja: "極白虎征魂戦",
              De: "Seelensturm - Byakko",
              Fr: "La Clairière de Jade (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	292: {
		Name: LocalisedText{
              En: "Sigmascape V1.0 (Savage)",
              Ja: "次元の狭間オメガ零式：シグマ編1",
              De: "Sigmametrie 1.0 (episch)",
              Fr: "Sigmastice v1.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	293: {
		Name: LocalisedText{
              En: "Sigmascape V2.0 (Savage)",
              Ja: "次元の狭間オメガ零式：シグマ編2",
              De: "Sigmametrie 2.0 (episch)",
              Fr: "Sigmastice v2.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	294: {
		Name: LocalisedText{
              En: "Sigmascape V3.0 (Savage)",
              Ja: "次元の狭間オメガ零式：シグマ編3",
              De: "Sigmametrie 3.0 (episch)",
              Fr: "Sigmastice v3.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	295: {
		Name: LocalisedText{
              En: "Sigmascape V4.0 (Savage)",
              Ja: "次元の狭間オメガ零式：シグマ編4",
              De: "Sigmametrie 4.0 (episch)",
              Fr: "Sigmastice v4.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	473: {
		Name: LocalisedText{
              En: "The Valentione's Ceremony",
              Ja: "ヴァレンティオンセレモニー",
              De: "Valentionzeremonie",
              Fr: "La Cérémonie de la Valention",
            },
		HighEnd: false,
		ContentKind: ContentKind(22),
	},
	474: {
		Name: LocalisedText{
              En: "The Great Hunt",
              Ja: "リオレウス狩猟戦",
              De: "Jagd auf Rathalos",
              Fr: "Chasse au Rathalos",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	475: {
		Name: LocalisedText{
              En: "The Great Hunt (Extreme)",
              Ja: "極リオレウス狩猟戦",
              De: "Jagd auf Rathalos (schwer)",
              Fr: "Chasse au Rathalos (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	481: {
		Name: LocalisedText{
              En: "Chocobo Race: Tutorial",
              Ja: "チョコボレース：チュートリアル",
              De: "Chocobo-Rennen: Übungsbahn",
              Fr: "Course d'appentissage",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	482: {
		Name: LocalisedText{
              En: "Race 1 - Hugging the Inside",
              Ja: "第1節：インコースを狙え",
              De: "1: Immer schön innen",
              Fr: "CP1 - Toujours à la corde!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	483: {
		Name: LocalisedText{
              En: "Race 2 - Keep Away",
              Ja: "第2節：範囲攻撃を避けろ",
              De: "2: Vorsicht, Flächenangriff",
              Fr: "CP2 - Gare aux attaques à aires d'effet!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	484: {
		Name: LocalisedText{
              En: "Race 3 - Inability",
              Ja: "第3節：アビリティに頼るな",
              De: "3: Fähigkeiten sind nicht alles",
              Fr: "CP3 - En avant même sans aptitudes!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	485: {
		Name: LocalisedText{
              En: "Race 4 - Heavy Hooves",
              Ja: "第4節：ヘヴィなレース",
              De: "4: Ein gewichtiges Rennen",
              Fr: "CP4 - Une course très très pesante!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	486: {
		Name: LocalisedText{
              En: "Race 5 - Defending the Rush",
              Ja: "第5節：とんずらを阻止せよ",
              De: "5: Hiergeblieben!",
              Fr: "CP5 - Attention aux départs rapides!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	487: {
		Name: LocalisedText{
              En: "Race 6 - Road Rivals",
              Ja: "第6節：ライバルを叩け",
              De: "6: Rivale auf Trab",
              Fr: "CP6 - Pas de pitié pour les rivaux!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	488: {
		Name: LocalisedText{
              En: "Race 7 - Field of Dreams",
              Ja: "第7節：フィールド・オブ・ドリームズ",
              De: "7: Feld der Träume",
              Fr: "CP7 - Prendre du champ avec les champs!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	489: {
		Name: LocalisedText{
              En: "Race 8 - Playing Both Ends",
              Ja: "第8節：漁夫の利を狙え",
              De: "8: Der lachende Dritte",
              Fr: "CP8 - Laissons les rivaux s'éliminer!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	490: {
		Name: LocalisedText{
              En: "Race 9 - Stamina",
              Ja: "第9節：スタミナの戦い",
              De: "9: Der längere Atem",
              Fr: "CP9 - Une bataille d'endurance!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	491: {
		Name: LocalisedText{
              En: "Race 10 - Cat and Mouse",
              Ja: "第10節：逃げる者と追う者と",
              De: "10: Zwei Temperamente",
              Fr: "CP10 - Course poursuite!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	492: {
		Name: LocalisedText{
              En: "Race 11 - Mad Dash",
              Ja: "第11節：全速で駆け抜けろ",
              De: "11: Volle Kraft voraus",
              Fr: "CP11 - À bride abattue!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	493: {
		Name: LocalisedText{
              En: "Race 12 - Bag of Tricks",
              Ja: "第12節：アビリティを駆使せよ",
              De: "12: Fähigkeiten ohne Wenn und Extra",
              Fr: "CP12 - Une épreuve d'aptitude!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	494: {
		Name: LocalisedText{
              En: "Race 13 - Tag Team",
              Ja: "第13節：連携プレイに勝利せよ",
              De: "13: Ein unschlagbares Gespann",
              Fr: "CP13 - Un travail d'équipe!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	495: {
		Name: LocalisedText{
              En: "Race 14 - Heavier Hooves",
              Ja: "第14節：続ヘヴィなレース",
              De: "14: Ein äußerst gewichtiges Rennen",
              Fr: "CP14 - Une course encore plus pesante!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	496: {
		Name: LocalisedText{
              En: "Race 15 - Ultimatum",
              Ja: "第15節：究極のレース",
              De: "15: Rennen der Superlative",
              Fr: "CP15 - La course des extrêmes!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	497: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	498: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	499: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	500: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	501: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	502: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	503: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	504: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	505: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	506: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	507: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	508: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	509: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	510: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	511: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	512: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	513: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	514: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	515: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	516: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	517: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	518: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	519: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	520: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	521: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	522: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	523: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	524: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	525: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	526: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	527: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	528: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	529: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	530: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	531: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	532: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	533: {
		Name: LocalisedText{
              En: "Chocobo Race: Sagolii Road",
              Ja: "チョコボレース：サゴリーロード",
              De: "Chocobo-Rennen: Sagolii-Straße",
              Fr: "Course de chocobos: Route de Sagolii",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	534: {
		Name: LocalisedText{
              En: "Chocobo Race: Costa del Sol",
              Ja: "チョコボレース：コスタ・デル・ソル",
              De: "Chocobo-Rennen: Sonnenküste",
              Fr: "Course de chocobos: Costa del Sol",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	535: {
		Name: LocalisedText{
              En: "Chocobo Race: Tranquil Paths",
              Ja: "チョコボレース：トランキルパス",
              De: "Chocobo-Rennen: Pfad der Seelenruhe",
              Fr: "Course de chocobos: Sentes tranquilles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	536: {
		Name: LocalisedText{
              En: "The Swallow's Compass",
              Ja: "風水霊殿 ガンエン廟",
              De: "Kompass der Schwalbe",
              Fr: "Le Compas de l'Hirondelle",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	537: {
		Name: LocalisedText{
              En: "Castrum Fluminis",
              Ja: "ツクヨミ討滅戦",
              De: "Götterdämmerung - Tsukuyomi",
              Fr: "Castrum Fluminis",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	538: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Tsukuyomi's Pain",
              Ja: "極ツクヨミ討滅戦",
              De: "Zenit der Götter - Tsukuyomi",
              Fr: "Castrum Fluminis (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	539: {
		Name: LocalisedText{
              En: "The Weapon's Refrain (Ultimate)",
              Ja: "絶アルテマウェポン破壊作戦",
              De: "Heldenlied von Ultima (fatal)",
              Fr: "La Fantasmagorie d'Ultima (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	540: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 1-10)",
              Ja: "アメノミハシラ 1～10層",
              De: "Himmelssäule (Ebenen 1-10)",
              Fr: "Le Pilier des Cieux (étages 1-10)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	541: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 11-20)",
              Ja: "アメノミハシラ 11～20層",
              De: "Himmelssäule (Ebenen 11-20)",
              Fr: "Le Pilier des Cieux (étages 11-20)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	542: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 21-30)",
              Ja: "アメノミハシラ 21～30層",
              De: "Himmelssäule (Ebenen 21-30)",
              Fr: "Le Pilier des Cieux (étages 21-30)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	543: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 31-40)",
              Ja: "アメノミハシラ 31～40層",
              De: "Himmelssäule (Ebenen 31-40)",
              Fr: "Le Pilier des Cieux (étages 31-40)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	544: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 41-50)",
              Ja: "アメノミハシラ 41～50層",
              De: "Himmelssäule (Ebenen 41-50)",
              Fr: "Le Pilier des Cieux (étages 41-50)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	545: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 51-60)",
              Ja: "アメノミハシラ 51～60層",
              De: "Himmelssäule (Ebenen 51-60)",
              Fr: "Le Pilier des Cieux (étages 51-60)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	546: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 61-70)",
              Ja: "アメノミハシラ 61～70層",
              De: "Himmelssäule (Ebenen 61-70)",
              Fr: "Le Pilier des Cieux (étages 61-70)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	547: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 71-80)",
              Ja: "アメノミハシラ 71～80層",
              De: "Himmelssäule (Ebenen 71-80)",
              Fr: "Le Pilier des Cieux (étages 71-80)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	548: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 81-90)",
              Ja: "アメノミハシラ 81～90層",
              De: "Himmelssäule (Ebenen 81-90)",
              Fr: "Le Pilier des Cieux (étages 81-90)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	549: {
		Name: LocalisedText{
              En: "Heaven-on-High  (Floors 91-100)",
              Ja: "アメノミハシラ 91～100層",
              De: "Himmelssäule (Ebenen 91-100)",
              Fr: "Le Pilier des Cieux (étages 91-100)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	550: {
		Name: LocalisedText{
              En: "The Ridorana Lighthouse",
              Ja: "封じられた聖塔 リドルアナ",
              De: "Richtfeuer von Ridorana",
              Fr: "Le Phare de Ridorana",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	552: {
		Name: LocalisedText{
              En: "Stage 1: Tutorial",
              Ja: "第1節：チュートリアル",
              De: "Stufe 1: Einführung",
              Fr: "Bataille 1: Tutoriel",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	553: {
		Name: LocalisedText{
              En: "Stage 2: Hatching a Plan",
              Ja: "第2節：チョコチョコチョコチョコボ",
              De: "Stufe 2: Federn lassen",
              Fr: "Bataille 2: Ô beaux chocobos",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	554: {
		Name: LocalisedText{
              En: "Stage 3: The First Move",
              Ja: "第3節：序盤にありそうな戦い",
              De: "Stufe 3: Angriff der Ungeziefer",
              Fr: "Bataille 3: Des ouvertures",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	555: {
		Name: LocalisedText{
              En: "Stage 4: Little Big Beast",
              Ja: "第4節：小さな巨獣",
              De: "Stufe 4: Riesenbaby-Goobbue",
              Fr: "Bataille 4: Grosses petites bestioles",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	556: {
		Name: LocalisedText{
              En: "Stage 5: Turning Tribes",
              Ja: "第5節：蛮族たちの反乱",
              De: "Stufe 5: Aufstand der Wilden",
              Fr: "Bataille 5: Hommes-bêtes révoltés",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	557: {
		Name: LocalisedText{
              En: "Stage 6: Off the Deepcroft",
              Ja: "第6節：恐怖、タラタラの墓所",
              De: "Stufe 6: Gedankenspiele",
              Fr: "Bataille 6: Hypogée de Tam-Tam",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	558: {
		Name: LocalisedText{
              En: "Stage 7: Rivals",
              Ja: "第7節：ライバルたちの共闘",
              De: "Stufe 7: Vierbeinige Feinde",
              Fr: "Bataille 7: Comme chien et chat",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	559: {
		Name: LocalisedText{
              En: "Stage 8: Always Darkest",
              Ja: "第8節：暁は光と闇とを分かつ",
              De: "Stufe 8: Das Bündchen der Morgenröte",
              Fr: "Bataille 8: Héritiers miniatures",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	560: {
		Name: LocalisedText{
              En: "Stage 9: Mine Your Minions",
              Ja: "第9節：爆破、カパカパベル銅山",
              De: "Stufe 9: Auf den Schleim gegangen",
              Fr: "Bataille 9: Déflagration à Clochecuivre",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	561: {
		Name: LocalisedText{
              En: "Stage 10: Children of Mandragora",
              Ja: "第10節：マンドラ帝国の逆襲",
              De: "Stufe 10: Mandragora-Salat",
              Fr: "Bataille 10: Invasion de mandragores",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	562: {
		Name: LocalisedText{
              En: "Stage 11: The Queen and I",
              Ja: "第11節：女王陛下万歳！",
              De: "Stufe 11: Die trabantäische Allianz",
              Fr: "Bataille 11: Longue vie aux Sultanes!",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	563: {
		Name: LocalisedText{
              En: "Stage 12: Breakout",
              Ja: "第12節：デモンズブロック崩し",
              De: "Stufe 12: Alea iacta est",
              Fr: "Bataille 12: Château d'Amdapeur",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	564: {
		Name: LocalisedText{
              En: "Stage 13: My Name Is Cid",
              Ja: "第13節：筆頭機工師ガーロンド",
              De: "Stufe 13: Luftadmiral Garlond",
              Fr: "Bataille 13: Épreuve des Forges",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	565: {
		Name: LocalisedText{
              En: "Stage 14: Like a Nut",
              Ja: "第14節：ナッツの味にはもう飽きた",
              De: "Stufe 14: Eins auf die Nuss",
              Fr: "Bataille 14: Casse-noisette",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	566: {
		Name: LocalisedText{
              En: "Stage 15: Urth's Spout",
              Ja: "第15節：極小オーディン討滅戦",
              De: "Stufe 15: Der Dunkle Reiter",
              Fr: "Bataille 15: Mini-fontaine d'Urth",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	567: {
		Name: LocalisedText{
              En: "Stage 16: Exodus",
              Ja: "第16節：チョコボ大行進",
              De: "Stufe 16: Chocobo-Mania",
              Fr: "Bataille 16: Parade des chocobos",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	568: {
		Name: LocalisedText{
              En: "Stage 17: Over the Wall",
              Ja: "第17節：融解、スノークローク小氷壁",
              De: "Stufe 17: Jenseits des Schneekleids",
              Fr: "Bataille 17: Bonhommes de Manteneige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	569: {
		Name: LocalisedText{
              En: "Stage 18: The Hunt",
              Ja: "第18節：小さなモブハント",
              De: "Stufe 18: Die Trabanten-Jagd",
              Fr: "Bataille 18: Mini-contrat de chasse",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	570: {
		Name: LocalisedText{
              En: "Stage 19: Battle on the Bitty Bridge",
              Ja: "第19節：極小ギルガメッシュ討滅戦",
              De: "Stufe 19: Gilgameschs Hühnerstall",
              Fr: "Bataille 19: Revanche de Gilgamini",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	571: {
		Name: LocalisedText{
              En: "Stage 20: Guiding Light",
              Ja: "第20節：汝にクリスタルの導きあれ",
              De: "Stufe 20: Trabanten des Lichts",
              Fr: "Bataille 20: Pouvoir de la Lumière",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	572: {
		Name: LocalisedText{
              En: "Stage 21: Wise Words",
              Ja: "第21節：賢人ルイゾワの試練",
              De: "Stufe 21: Die Leveilleurs",
              Fr: "Bataille 21: Louisoix et petits fils",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	573: {
		Name: LocalisedText{
              En: "Stage 22: World of Poor Lighting",
              Ja: "第22節：薄闇の世界",
              De: "Stufe 22: Kleine Welt der Dunkelheit",
              Fr: "Bataille 22: Monde des Quasi-ténèbres",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	574: {
		Name: LocalisedText{
              En: "Stage 23: The Binding Coil",
              Ja: "第23節：大迷惑バハムート邂逅編",
              De: "Stufe 23: In die Verschlungenen Schatten",
              Fr: "Bataille 23: Labyrinthe de Basta-mut",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	575: {
		Name: LocalisedText{
              En: "Stage 24: The Final Coil",
              Ja: "第24節：大迷惑バハムート真成編",
              De: "Stufe 24: Bahamut, der Aufzieh-Primae",
              Fr: "Bataille 24: Abîme de Basta-mut",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	576: {
		Name: LocalisedText{
              En: "LoVM: Master Battle",
              Ja: "LoVM：ミニオンクラス (Easy)",
              De: "Kampf der Trabanten: Trabantenkampf",
              Fr: "Bataille simple contre l'ordinateur (facile)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	577: {
		Name: LocalisedText{
              En: "LoVM: Master Battle (Hard)",
              Ja: "LoVM：真ミニオンクラス (Normal)",
              De: "Kampf der Trabanten: Trabantendämmerung",
              Fr: "Bataille simple contre l'ordinateur (normal)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	578: {
		Name: LocalisedText{
              En: "LoVM: Master Battle (Extreme)",
              Ja: "LoVM：極ミニオンクラス (Hard)",
              De: "Kampf der Trabanten: Zenit der Trabanten",
              Fr: "Bataille simple contre l'ordinateur (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	579: {
		Name: LocalisedText{
              En: "LoVM: Master Tournament",
              Ja: "LoVM：大会対戦 (CPU対戦)",
              De: "Kampf der Trabanten: Turnier (gegen Arenameister)",
              Fr: "Bataille de tournoi contre l'ordinateur",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	581: {
		Name: LocalisedText{
              En: "The Forbidden Land, Eureka Pagos",
              Ja: "禁断の地 エウレカ：パゴス編",
              De: "Eureka Pagos",
              Fr: "Eurêka Pagos",
            },
		HighEnd: false,
		ContentKind: ContentKindEureka,
	},
	582: {
		Name: LocalisedText{
              En: "Emissary of the Dawn",
              Ja: "「暁」の少年",
              De: "Der Knabe der Morgenröte",
              Fr: "Voyage en terre hostile",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	583: {
		Name: LocalisedText{
              En: "The Calamity Retold",
              Ja: "新生祭軍事演習",
              De: "Gedenkschlacht der Eorzäischen Allianz",
              Fr: "Les grandes manœuvres commémoratives",
            },
		HighEnd: false,
		ContentKind: ContentKind(22),
	},
	584: {
		Name: LocalisedText{
              En: "Saint Mocianne's Arboretum (Hard)",
              Ja: "草木汚染 聖モシャーヌ植物園 (Hard)",
              De: "Sankt Mocianne-Arboretum (schwer)",
              Fr: "L'Arboretum Sainte-Mocianne (brutal)",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	585: {
		Name: LocalisedText{
              En: "The Burn",
              Ja: "永久焦土 ザ・バーン",
              De: "Das Kargland",
              Fr: "L'Escarre",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	586: {
		Name: LocalisedText{
              En: "The Shifting Altars of Uznair",
              Ja: "宝物庫 ウズネアカナル祭殿",
              De: "Glücksaltäre von Uznair",
              Fr: "Le Temple sacré d'Uznair",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	587: {
		Name: LocalisedText{
              En: "Alphascape V1.0",
              Ja: "次元の狭間オメガ：アルファ編1",
              De: "Alphametrie 1.0",
              Fr: "Alphastice v1.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	588: {
		Name: LocalisedText{
              En: "Alphascape V2.0",
              Ja: "次元の狭間オメガ：アルファ編2",
              De: "Alphametrie 2.0",
              Fr: "Alphastice v2.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	589: {
		Name: LocalisedText{
              En: "Alphascape V3.0",
              Ja: "次元の狭間オメガ：アルファ編3",
              De: "Alphametrie 3.0",
              Fr: "Alphastice v3.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	590: {
		Name: LocalisedText{
              En: "Alphascape V4.0",
              Ja: "次元の狭間オメガ：アルファ編4",
              De: "Alphametrie 4.0",
              Fr: "Alphastice v4.0",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	591: {
		Name: LocalisedText{
              En: "Alphascape V1.0 (Savage)",
              Ja: "次元の狭間オメガ零式：アルファ編1",
              De: "Alphametrie 1.0 (episch)",
              Fr: "Alphastice v1.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	592: {
		Name: LocalisedText{
              En: "Alphascape V2.0 (Savage)",
              Ja: "次元の狭間オメガ零式：アルファ編2",
              De: "Alphametrie 2.0 (episch)",
              Fr: "Alphastice v2.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	593: {
		Name: LocalisedText{
              En: "Alphascape V3.0 (Savage)",
              Ja: "次元の狭間オメガ零式：アルファ編3",
              De: "Alphametrie 3.0 (episch)",
              Fr: "Alphastice v3.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	594: {
		Name: LocalisedText{
              En: "Alphascape V4.0 (Savage)",
              Ja: "次元の狭間オメガ零式：アルファ編4",
              De: "Alphametrie 4.0 (episch)",
              Fr: "Alphastice v4.0 (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	595: {
		Name: LocalisedText{
              En: "Kugane Ohashi",
              Ja: "真ヨウジンボウ討滅戦",
              De: "Duell auf der Kugane-Brücke",
              Fr: "Le Pont Ohashi",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	596: {
		Name: LocalisedText{
              En: "Hells' Kier",
              Ja: "朱雀征魂戦",
              De: "Seelentanz - Suzaku",
              Fr: "Le Nid des Lamentations",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	597: {
		Name: LocalisedText{
              En: "Hells' Kier (Extreme)",
              Ja: "極朱雀征魂戦",
              De: "Seelensturm - Suzaku",
              Fr: "Le Nid des Lamentations (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	598: {
		Name: LocalisedText{
              En: "The Forbidden Land, Eureka Pyros",
              Ja: "禁断の地 エウレカ：ピューロス編",
              De: "Eureka Pyros",
              Fr: "Eurêka Pyros",
            },
		HighEnd: false,
		ContentKind: ContentKindEureka,
	},
	599: {
		Name: LocalisedText{
              En: "Hidden Gorge",
              Ja: "ヒドゥンゴージ (機工戦)",
              De: "Verborgene Schlucht",
              Fr: "Gorge dérobée (machinerie)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	600: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	601: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	602: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	603: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	604: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	605: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	606: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	607: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	608: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	609: {
		Name: LocalisedText{
              En: "The Will of the Moon",
              Ja: "楔石の虚",
              De: "Der Wille der Mondgöttin",
              Fr: "Ralliement dans la steppe",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	610: {
		Name: LocalisedText{
              En: "All's Well That Starts Well",
              Ja: "デビューマッチ",
              De: "Debüt in der Himmlischen Arena",
              Fr: "Début du spectacle",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	611: {
		Name: LocalisedText{
              En: "The Ghimlyt Dark",
              Ja: "境界戦線 ギムリトダーク",
              De: "Die Ghimlyt-Finsternis",
              Fr: "Les Ténèbres de Ghimlyt",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	612: {
		Name: LocalisedText{
              En: "Much Ado About Pudding",
              Ja: "プリン・アラモード",
              De: "Pudding nach Art des Hauses",
              Fr: "Puddings à la mode",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	613: {
		Name: LocalisedText{
              En: "Waiting for Golem",
              Ja: "最初の岩壁「シパクナー」",
              De: "Zipacna - Aller Anfang ist schwer",
              Fr: "Zipacna, le premier obstacle",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	614: {
		Name: LocalisedText{
              En: "Gentlemen Prefer Swords",
              Ja: "怪力の鉄巨人「クレイオス」",
              De: "Kreios - Der Mann aus Stahl",
              Fr: "Kreios, le destructeur d'acier",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	615: {
		Name: LocalisedText{
              En: "The Threepenny Turtles",
              Ja: "ギルガメブラザーズ",
              De: "Die Bruderschaft der Kröten",
              Fr: "La fratrie des gilkhélones",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	616: {
		Name: LocalisedText{
              En: "Eye Society",
              Ja: "ブラインド・フューリー",
              De: "Den Tod im Auge",
              Fr: "Vengeance aveugle",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	617: {
		Name: LocalisedText{
              En: "A Chorus Slime",
              Ja: "シアーハートアタック",
              De: "Heißkalte Schauer",
              Fr: "Pure attaque cardiaque",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	618: {
		Name: LocalisedText{
              En: "Bomb-edy of Errors",
              Ja: "青い牙、赤い牙",
              De: "Explosiv - Vorhut in Blau und Rot",
              Fr: "Crocs bleus et crocs rouges",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	619: {
		Name: LocalisedText{
              En: "To Kill a Mockingslime",
              Ja: "七色の甘味「ギモーヴ」",
              De: "Guimauve und die sieben Gesichter des Todes",
              Fr: "Guimauve, le goût de l'arc-en-ciel",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	620: {
		Name: LocalisedText{
              En: "A Little Knight Music",
              Ja: "偶像の王者「クロムドゥーブ」",
              De: "Crom Dubh - König der Götzen",
              Fr: "Crom Dubh, roi des idoles",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	621: {
		Name: LocalisedText{
              En: "Some Like It Excruciatingly Hot",
              Ja: "爆ボム・ファーストアタック",
              De: "Explosiv - Die zweite Welle",
              Fr: "Les courtes mèches",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	622: {
		Name: LocalisedText{
              En: "The Plant-om of the Opera",
              Ja: "寄生植物「ヒドノラ」",
              De: "Hydnora - Der Parasit, der Paras sieht",
              Fr: "Hydnora, la plante parasite",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	623: {
		Name: LocalisedText{
              En: "Beauty and a Beast",
              Ja: "紅の死妖姫「カーミラ」",
              De: "Carmilla - Prinzessin des Todes",
              Fr: "La mystérieuse Carmilla",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	624: {
		Name: LocalisedText{
              En: "Blobs in the Woods",
              Ja: "死なばもろともーッ！",
              De: "Der Tod kommt süß gekleidet",
              Fr: "La mort n'a pas d'ami",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	625: {
		Name: LocalisedText{
              En: "The Me Nobody Nodes",
              Ja: "アラグの脅威「闘獣システム」",
              De: "Die Bestie steckt im System",
              Fr: "La sphère bestiale, une menace allagoise",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	626: {
		Name: LocalisedText{
              En: "Sunset Bull-evard",
              Ja: "豪腕の獣王「ティクバラン」",
              De: "Tikbalang - Der rechte Arm des Verderbens",
              Fr: "Tikbalang, bras tout-puissant",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	627: {
		Name: LocalisedText{
              En: "The Sword of Music",
              Ja: "剛柔の鉄巨人「クレイオス改」",
              De: "Kreios - Neues vom Mann aus Stahl",
              Fr: "Kreios plie, mais ne rompt pas",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	628: {
		Name: LocalisedText{
              En: "Midsummer Night's Explosion",
              Ja: "爆破デスマッチ",
              De: "Explosiv - Revanche mit Wumms",
              Fr: "Rencontre explosive",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	629: {
		Name: LocalisedText{
              En: "On a Clear Day You Can Smell Forever",
              Ja: "魅惑の芳香「リフレクティブ・レベッカ」",
              De: "Rebekkha - Verführerische Gerüche",
              Fr: "Miroir, mon beau miroir",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	630: {
		Name: LocalisedText{
              En: "Miss Typhon",
              Ja: "名コンビ「オルトロス＆テュポーン」",
              De: "Ultros & Typhon - Zwei wie Rotz und Wasser",
              Fr: "Duo de choc: Orthros et maître Typhon",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	631: {
		Name: LocalisedText{
              En: "Chimera on a Hot Tin Roof",
              Ja: "憤怒の合成獣「アペデマク」",
              De: "Apademak - Die cholerische Chimära",
              Fr: "La chimèrique colère d'Apademak",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	632: {
		Name: LocalisedText{
              En: "Here Comes the Boom",
              Ja: "爆発の対消滅「グランパボム」",
              De: "Explosiv - Urgroßbomber lässt es krachen",
              Fr: "L'histoire détonante de Papi bombo",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	633: {
		Name: LocalisedText{
              En: "Behemoths and Broomsticks",
              Ja: "魔獣の皇太子「クロンプリンツ・ベヒーモス」",
              De: "Kronprinz-Behemoth - Vom Teufel besessen",
              Fr: "Kronprinz béhémoth, le prince héritier",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	634: {
		Name: LocalisedText{
              En: "Amazing Technicolor Pit Fiends",
              Ja: "異形の人形師「エペロギ」",
              De: "Epilogi - Das Ende kommt immer zuletzt",
              Fr: "Epilogi, l'étrange marionnettiste",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	635: {
		Name: LocalisedText{
              En: "Dirty Rotten Azulmagia",
              Ja: "悪の青魔道士「アポカリョープス」",
              De: "Azulmagia - Der Blaumagier des Bösen",
              Fr: "L'abominable Azulmagia",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	636: {
		Name: LocalisedText{
              En: "The Orbonne Monastery",
              Ja: "楽欲の僧院 オーボンヌ",
              De: "Kloster von Orbonne",
              Fr: "Le Monastère d'Orbonne",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	637: {
		Name: LocalisedText{
              En: "The Wreath of Snakes",
              Ja: "青龍征魂戦",
              De: "Seelentanz - Seiryu",
              Fr: "L'Îlot des Amertumes",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	638: {
		Name: LocalisedText{
              En: "The Wreath of Snakes (Extreme)",
              Ja: "極青龍征魂戦",
              De: "Seelensturm - Seiryu",
              Fr: "L'Îlot des Amertumes (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	639: {
		Name: LocalisedText{
              En: "The Forbidden Land, Eureka Hydatos",
              Ja: "禁断の地 エウレカ：ヒュダトス編",
              De: "Eureka Hydatos",
              Fr: "Eurêka Hydatos",
            },
		HighEnd: false,
		ContentKind: ContentKindEureka,
	},
	640: {
		Name: LocalisedText{
              En: "Air Force One",
              Ja: "出撃！ エアフォースパイロット",
              De: "Luftwaffe, Feuer frei!",
              Fr: "As de l'air",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	641: {
		Name: LocalisedText{
              En: "Air Force One",
              Ja: "出撃！ エアフォースパイロット",
              De: "Luftwaffe, Feuer frei!",
              Fr: "As de l'air",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	642: {
		Name: LocalisedText{
              En: "Air Force One",
              Ja: "出撃！ エアフォースパイロット",
              De: "Luftwaffe, Feuer frei!",
              Fr: "As de l'air",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	643: {
		Name: LocalisedText{
              En: "Novice Mahjong (Full Ranked Match)",
              Ja: "ドマ式麻雀：半荘戦一般卓（段位変動有り）",
              De: "Anfänger-Mahjong (komplette Partie, gewertet)",
              Fr: "Mahjong domien: tous rangs (partie classée)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	644: {
		Name: LocalisedText{
              En: "Advanced Mahjong (Full Ranked Match)",
              Ja: "ドマ式麻雀：半荘戦有段卓（段位変動有り）",
              De: "Fortgeschrittenen-Mahjong (komplette Partie, gewertet)",
              Fr: "Mahjong domien: dan seulement (partie classée)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	645: {
		Name: LocalisedText{
              En: "Four-player Mahjong (Full Match, Kuitan Enabled)",
              Ja: "ドマ式麻雀：半荘戦4人セット卓（クイタン有り）",
              De: "4-Spieler-Mahjong (komplette Partie, Kuitan aktiviert)",
              Fr: "Mahjong domien: 4 joueurs (partie avec kuitan)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	646: {
		Name: LocalisedText{
              En: "Messenger of the Winds",
              Ja: "来訪せし風の御使",
              De: "Durch den Sturm und zurück",
              Fr: "La Messagère du vent",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	648: {
		Name: LocalisedText{
              En: "A Requiem for Heroes",
              Ja: "英雄への鎮魂歌",
              De: "Requiem der Helden",
              Fr: "Un requiem pour les héros",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	649: {
		Name: LocalisedText{
              En: "Dohn Mheg",
              Ja: "水妖幻園 ドォーヌ・メグ",
              De: "Dohn Mheg",
              Fr: "Dohn Mheg",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	650: {
		Name: LocalisedText{
              En: "Four-player Mahjong (Full Match, Kuitan Disabled)",
              Ja: "ドマ式麻雀：半荘戦4人セット卓（クイタン無し）",
              De: "4-Spieler-Mahjong (komplette Partie, Kuitan deaktiviert)",
              Fr: "Mahjong domien: 4 joueurs (partie sans kuitan)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	651: {
		Name: LocalisedText{
              En: "The Qitana Ravel",
              Ja: "古跡探索 キタンナ神影洞",
              De: "Irrungen der Qitari",
              Fr: "L'Enchevêtrement des Qitari",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	652: {
		Name: LocalisedText{
              En: "Amaurot",
              Ja: "終末幻想 アーモロート",
              De: "Amaurot",
              Fr: "Amaurote",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	653: {
		Name: LocalisedText{
              En: "Eden's Gate: Resurrection",
              Ja: "希望の園エデン：覚醒編1",
              De: "Edens Erwachen - Auferstehung",
              Fr: "L'Éveil d'Éden - Résurrection",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	654: {
		Name: LocalisedText{
              En: "Eden's Gate: Resurrection (Savage)",
              Ja: "希望の園エデン零式：覚醒編1",
              De: "Edens Erwachen - Auferstehung (episch)",
              Fr: "L'Éveil d'Éden - Résurrection (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	655: {
		Name: LocalisedText{
              En: "The Twinning",
              Ja: "異界遺構 シルクス・ツイニング",
              De: "Der Kristallzwilling",
              Fr: "La Macle de Syrcus",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	656: {
		Name: LocalisedText{
              En: "Malikah's Well",
              Ja: "爽涼離宮 マリカの大井戸",
              De: "Malikahs Brunnen",
              Fr: "Le Puits de Malikah",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	657: {
		Name: LocalisedText{
              En: "The Dancing Plague",
              Ja: "ティターニア討滅戦",
              De: "Offenbarung - Titania",
              Fr: "La Valse du Monarque",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	658: {
		Name: LocalisedText{
              En: "The Dancing Plague (Extreme)",
              Ja: "極ティターニア討滅戦",
              De: "Letzte Läuterung - Titania",
              Fr: "La Valse du Monarque (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	659: {
		Name: LocalisedText{
              En: "Mt. Gulg",
              Ja: "偽造天界 グルグ火山",
              De: "Der Gulg",
              Fr: "Mont Gulg",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	661: {
		Name: LocalisedText{
              En: "Akadaemia Anyder",
              Ja: "創造機関 アナイダアカデミア",
              De: "Akadaemia Anyder",
              Fr: "Akadaemia Anydre",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	666: {
		Name: LocalisedText{
              En: "The Crown of the Immaculate",
              Ja: "イノセンス討滅戦",
              De: "Offenbarung - Innozenz",
              Fr: "La Couronne de l'Immaculé",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	667: {
		Name: LocalisedText{
              En: "The Crown of the Immaculate (Extreme)",
              Ja: "極イノセンス討滅戦",
              De: "Letzte Läuterung - Innozenz",
              Fr: "La Couronne de l'Immaculé (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	676: {
		Name: LocalisedText{
              En: "Holminster Switch",
              Ja: "殺戮郷村 ホルミンスター",
              De: "Holminster",
              Fr: "Holminster",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	678: {
		Name: LocalisedText{
              En: "The Hardened Heart",
              Ja: "揺れる天秤",
              De: "Ob Mitleid oder Hass",
              Fr: "Naissance d'un bourreau",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	679: {
		Name: LocalisedText{
              En: "The Lost and the Found",
              Ja: "古の大再生魔法",
              De: "Alter Zauber",
              Fr: "Magie ancestrale",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	680: {
		Name: LocalisedText{
              En: "Coming Clean",
              Ja: "廃都ナバスアレン",
              De: "Vater und Bruder",
              Fr: "Sur les rails de Nabaath Areng",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	681: {
		Name: LocalisedText{
              En: "Legend of the Not-so-hidden Temple",
              Ja: "仕掛けと呪いと毒と",
              De: "Der Beichtstuhl von Toupasa dem Älteren",
              Fr: "Le Confessionnal de Toupasa l'ancien",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	682: {
		Name: LocalisedText{
              En: "Eden's Gate: Inundation",
              Ja: "希望の園エデン：覚醒編3",
              De: "Edens Erwachen - Überflutung",
              Fr: "L'Éveil d'Éden - Déluge",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	683: {
		Name: LocalisedText{
              En: "Eden's Gate: Inundation (Savage)",
              Ja: "希望の園エデン零式：覚醒編3",
              De: "Edens Erwachen - Überflutung (episch)",
              Fr: "L'Éveil d'Éden - Déluge (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	684: {
		Name: LocalisedText{
              En: "Eden's Gate: Descent",
              Ja: "希望の園エデン：覚醒編2",
              De: "Edens Erwachen - Niederkunft",
              Fr: "L'Éveil d'Éden - Descente",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	685: {
		Name: LocalisedText{
              En: "Eden's Gate: Descent (Savage)",
              Ja: "希望の園エデン零式：覚醒編2",
              De: "Edens Erwachen - Niederkunft (episch)",
              Fr: "L'Éveil d'Éden - Descente (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	686: {
		Name: LocalisedText{
              En: "Nyelbert's Lament",
              Ja: "ナイルベルトの後悔",
              De: "Ein großes Opfer",
              Fr: "Une cupidité bien généreuse",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	687: {
		Name: LocalisedText{
              En: "The Dying Gasp",
              Ja: "ハーデス討滅戦",
              De: "Offenbarung - Hades",
              Fr: "Le Râle de l'Agonie",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	688: {
		Name: LocalisedText{
              En: "The Dungeons of Lyhe Ghiah",
              Ja: "宝物庫 リェー・ギア・ダンジョン",
              De: "Verliese von Lyhe Ghiah",
              Fr: "Le Donjon hypogéen du Lyhe Ghiah",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	689: {
		Name: LocalisedText{
              En: "Eden's Gate: Sepulture",
              Ja: "希望の園エデン：覚醒編4",
              De: "Edens Erwachen - Beerdigung",
              Fr: "L'Éveil d'Éden - Inhumation",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	690: {
		Name: LocalisedText{
              En: "Eden's Gate: Sepulture (Savage)",
              Ja: "希望の園エデン零式：覚醒編4",
              De: "Edens Erwachen - Beerdigung (episch)",
              Fr: "L'Éveil d'Éden - Inhumation (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	691: {
		Name: LocalisedText{
              En: "The Hunter's Legacy",
              Ja: "勇気の狩人",
              De: "Der Legende auf der Spur",
              Fr: "La chasseuse de légende",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	692: {
		Name: LocalisedText{
              En: "The Grand Cosmos",
              Ja: "魔法宮殿 グラン・コスモス",
              De: "Chateau Cosmea",
              Fr: "Le Cosmos coruscant",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	693: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Hades's Elegy",
              Ja: "極ハーデス討滅戦",
              De: "Letzte Läuterung - Hades",
              Fr: "Le Râle de l'Agonie (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	694: {
		Name: LocalisedText{
              En: "The Epic of Alexander (Ultimate)",
              Ja: "絶アレキサンダー討滅戦",
              De: "Alexander (fatal)",
              Fr: "L'Odyssée d'Alexander (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	695: {
		Name: LocalisedText{
              En: "Papa Mia",
              Ja: "豪腕の父親「フンババ・パパ」",
              De: "Papa Mia",
              Fr: "Papa Humbaba, le paternel aux gros bras",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	696: {
		Name: LocalisedText{
              En: "Lock up Your Snorters",
              Ja: "はないき爆破デスマッチ",
              De: "Explosives Schnaufen",
              Fr: "Fungaaah! Et boum!",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	697: {
		Name: LocalisedText{
              En: "Dangerous When Dead",
              Ja: "至妙の傀儡子「ドゥリン」",
              De: "Begnadeter Puppenspieler",
              Fr: "Durinn, marionnettiste d'outre-tombe",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	698: {
		Name: LocalisedText{
              En: "Red, Fraught, and Blue",
              Ja: "水と炎の歌",
              De: "Die Melodie von Feuer und Wasser",
              Fr: "La mélodie du feu et de l'eau",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	699: {
		Name: LocalisedText{
              En: "The Catch of the Siegfried",
              Ja: "世界一の剣士「ジークフリード」",
              De: "Siegfried, der beste Schwertkämpfer der Welt!",
              Fr: "Siegfried, le plus grand bretteur du monde",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	700: {
		Name: LocalisedText{
              En: "The Copied Factory",
              Ja: "複製サレタ工場廃墟",
              De: "Die kopierte Fabrik",
              Fr: "La réplique de l'usine désaffectée",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	701: {
		Name: LocalisedText{
              En: "Onsal Hakair (Danshig Naadam)",
              Ja: "オンサル・ハカイル (終節戦)",
              De: "Onsal Hakair (Danshig Naadam)",
              Fr: "Onsal Hakair (Danshig Naadam)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	702: {
		Name: LocalisedText{
              En: "Vows of Virtue, Deeds of Cruelty",
              Ja: "白き誓約、黒き密約",
              De: "Der Wolf und der Drachenreiter",
              Fr: "Vœux de vertu, actes de cruauté",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	703: {
		Name: LocalisedText{
              En: "As the Heart Bids",
              Ja: "この心が望むがままに",
              De: "Trubel im Traumland",
              Fr: "À l'écoute de soi",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	705: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	706: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	707: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	708: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	709: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	710: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	711: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	712: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	713: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	714: {
		Name: LocalisedText{
              En: "Anamnesis Anyder",
              Ja: "黒風海底 アニドラス・アナムネーシス",
              De: "Anamnesis Anyder",
              Fr: "Anamnesis Anydre",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	715: {
		Name: LocalisedText{
              En: "Eden's Verse: Fulmination",
              Ja: "希望の園エデン：共鳴編1",
              De: "Edens Resonanz - Entladung",
              Fr: "Les Accords d'Éden - Fulmination",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	716: {
		Name: LocalisedText{
              En: "Eden's Verse: Fulmination (Savage)",
              Ja: "希望の園エデン零式：共鳴編1",
              De: "Edens Resonanz - Entladung (episch)",
              Fr: "Les Accords d'Éden - Fulmination (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	717: {
		Name: LocalisedText{
              En: "Cinder Drift",
              Ja: "ルビーウェポン破壊作戦",
              De: "Rubinfeuer - Entfesselung",
              Fr: "Les Nuées de Brandons",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	718: {
		Name: LocalisedText{
              En: "Cinder Drift (Extreme)",
              Ja: "極ルビーウェポン破壊作戦",
              De: "Rubinfeuer - Trauma",
              Fr: "Les Nuées de Brandons (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	719: {
		Name: LocalisedText{
              En: "Eden's Verse: Furor",
              Ja: "希望の園エデン：共鳴編2",
              De: "Edens Resonanz - Raserei",
              Fr: "Les Accords d'Éden - Fureur",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	720: {
		Name: LocalisedText{
              En: "Eden's Verse: Furor (Savage)",
              Ja: "希望の園エデン零式：共鳴編2",
              De: "Edens Resonanz - Raserei (episch)",
              Fr: "Les Accords d'Éden - Fureur (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	721: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	722: {
		Name: LocalisedText{
              En: "The Diadem",
              Ja: "雲海採集 ディアデム諸島",
              De: "Das Diadem - Erschließung",
              Fr: "Le Diadème",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	723: {
		Name: LocalisedText{
              En: "The Bozja Incident",
              Ja: "シタデル・ボズヤ蒸発事変",
              De: "Der Bozja-Vorfall",
              Fr: "Prélude à la catastrophe",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	724: {
		Name: LocalisedText{
              En: "A Sleep Disturbed",
              Ja: "汝、英雄の眠り妨げるは",
              De: "Von schlafenden Helden",
              Fr: "L'épreuve ronka",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	725: {
		Name: LocalisedText{
              En: "Memoria Misera (Extreme)",
              Ja: "極シタデル・ボズヤ追憶戦",
              De: "Memoria Misera (extrem)",
              Fr: "Memoria Misera (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	726: {
		Name: LocalisedText{
              En: "Eden's Verse: Iconoclasm",
              Ja: "希望の園エデン：共鳴編3",
              De: "Edens Resonanz - Bildersturm",
              Fr: "Les Accords d'Éden - Iconoclasme",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	727: {
		Name: LocalisedText{
              En: "Eden's Verse: Iconoclasm (Savage)",
              Ja: "希望の園エデン零式：共鳴編3",
              De: "Edens Resonanz - Bildersturm (episch)",
              Fr: "Les Accords d'Éden - Iconoclasme (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	728: {
		Name: LocalisedText{
              En: "Eden's Verse: Refulgence",
              Ja: "希望の園エデン：共鳴編4",
              De: "Edens Resonanz - Erstarrung",
              Fr: "Les Accords d'Éden - Éclat",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	729: {
		Name: LocalisedText{
              En: "Eden's Verse: Refulgence (Savage)",
              Ja: "希望の園エデン零式：共鳴編4",
              De: "Edens Resonanz - Erstarrung (episch)",
              Fr: "Les Accords d'Éden - Éclat (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	730: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	731: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	732: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	733: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	734: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	735: {
		Name: LocalisedText{
              En: "The Bozjan Southern Front",
              Ja: "南方ボズヤ戦線",
              De: "Bozja-Südfront",
              Fr: "Front sud de Bozja",
            },
		HighEnd: false,
		ContentKind: ContentKindSavetheQueen,
	},
	736: {
		Name: LocalisedText{
              En: "The Puppets' Bunker",
              Ja: "人形タチノ軍事基地",
              De: "Die Puppenfestung",
              Fr: "La base militaire des Pantins",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	737: {
		Name: LocalisedText{
              En: "The Heroes' Gauntlet",
              Ja: "漆黒決戦 ノルヴラント",
              De: "Schlacht um Norvrandt",
              Fr: "La Traversée de Norvrandt",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	738: {
		Name: LocalisedText{
              En: "The Seat of Sacrifice",
              Ja: "ウォーリア・オブ・ライト討滅戦",
              De: "Krieger des Lichts",
              Fr: "Le Trône du Sacrifice",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	739: {
		Name: LocalisedText{
              En: "The Seat of Sacrifice (Extreme)",
              Ja: "極ウォーリア・オブ・ライト討滅戦",
              De: "Krieger des Lichts (extrem)",
              Fr: "Le Trône du Sacrifice (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	740: {
		Name: LocalisedText{
              En: "Sleep Now in Sapphire",
              Ja: "飛べ！ ウェルリトへ ",
              De: "Luftangriff auf Werlyt",
              Fr: "Sur la mer de saphir",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	741: {
		Name: LocalisedText{
              En: "Sleep Now in Sapphire",
              Ja: "飛べ！ ウェルリトへ ",
              De: "Luftangriff auf Werlyt",
              Fr: "Sur la mer de saphir",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	742: {
		Name: LocalisedText{
              En: "The Diadem",
              Ja: "雲海採集 ディアデム諸島",
              De: "Das Diadem - Erschließung",
              Fr: "Le Diadème",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	743: {
		Name: LocalisedText{
              En: "Faded Memories",
              Ja: "色あせた記憶",
              De: "Verblasste Erinnerungen",
              Fr: "Souvenir périssable",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	745: {
		Name: LocalisedText{
              En: "The Shifting Oubliettes of Lyhe Ghiah",
              Ja: "宝物庫 リェー・ギア・ダンジョン祭殿",
              De: "Das Karussell von Lyhe Ghiah",
              Fr: "Le Jardin secret du Lyhe Ghiah",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	746: {
		Name: LocalisedText{
              En: "Matoya's Relict",
              Ja: "魔術工房 マトーヤのアトリエ",
              De: "Matoyas Atelier",
              Fr: "L'Atelier abandonné de Matoya",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	747: {
		Name: LocalisedText{
              En: "Eden's Promise: Litany",
              Ja: "希望の園エデン：再生編2",
              De: "Edens Verheißung - Litanei",
              Fr: "La Promesse d'Éden - Litanie",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	748: {
		Name: LocalisedText{
              En: "Eden's Promise: Litany (Savage)",
              Ja: "希望の園エデン零式：再生編2",
              De: "Edens Verheißung - Litanei (episch)",
              Fr: "La Promesse d'Éden - Litanie (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	749: {
		Name: LocalisedText{
              En: "Eden's Promise: Umbra",
              Ja: "希望の園エデン：再生編1",
              De: "Edens Verheißung - Umbra",
              Fr: "La Promesse d'Éden - Nuée",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	750: {
		Name: LocalisedText{
              En: "Eden's Promise: Umbra (Savage)",
              Ja: "希望の園エデン零式：再生編1",
              De: "Edens Verheißung - Umbra (episch)",
              Fr: "La Promesse d'Éden - Nuée (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	751: {
		Name: LocalisedText{
              En: "Eden's Promise: Anamorphosis",
              Ja: "希望の園エデン：再生編3",
              De: "Edens Verheißung - Anamorphose",
              Fr: "La Promesse d'Éden - Anamorphose",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	752: {
		Name: LocalisedText{
              En: "Eden's Promise: Anamorphosis (Savage)",
              Ja: "希望の園エデン零式：再生編3",
              De: "Edens Verheißung - Anamorphose (episch)",
              Fr: "La Promesse d'Éden - Anamorphose (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	753: {
		Name: LocalisedText{
              En: "The Diadem",
              Ja: "雲海採集 ディアデム諸島",
              De: "Das Diadem - Erschließung",
              Fr: "Le Diadème",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	754: {
		Name: LocalisedText{
              En: "Anything Gogo's",
              Ja: "黄金頭巾「ものまね士ゴゴ」",
              De: "Gogo der Mime",
              Fr: "Gogo le mime",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	755: {
		Name: LocalisedText{
              En: "Triple Triad Open Tournament",
              Ja: "トリプルトライアド：オフィシャルトーナメント",
              De: "Triple Triad: Manderville-Turnier",
              Fr: "Tournoi officiel de Triple Triade",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	756: {
		Name: LocalisedText{
              En: "Triple Triad Invitational Parlor",
              Ja: "トリプルトライアド：カスタムトーナメントルーム",
              De: "Triple Triad: Privatturnier",
              Fr: "Salle de tournoi libre de Triple Triade",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	758: {
		Name: LocalisedText{
              En: "Eden's Promise: Eternity",
              Ja: "希望の園エデン：再生編4",
              De: "Edens Verheißung - Ewigkeit",
              Fr: "La Promesse d'Éden - Éternité",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	759: {
		Name: LocalisedText{
              En: "Eden's Promise: Eternity (Savage)",
              Ja: "希望の園エデン零式：再生編4",
              De: "Edens Verheißung - Ewigkeit (episch)",
              Fr: "La Promesse d'Éden - Éternité (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	760: {
		Name: LocalisedText{
              En: "Delubrum Reginae",
              Ja: "グンヒルド・ディルーブラム",
              De: "Delubrum Reginae",
              Fr: "Delubrum Reginae",
            },
		HighEnd: false,
		ContentKind: ContentKindSavetheQueen,
	},
	761: {
		Name: LocalisedText{
              En: "Delubrum Reginae (Savage)",
              Ja: "グンヒルド・ディルーブラム零式",
              De: "Delubrum Reginae (episch)",
              Fr: "Delubrum Reginae (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindSavetheQueen,
	},
	762: {
		Name: LocalisedText{
              En: "Castrum Marinum",
              Ja: "エメラルドウェポン破壊作戦",
              De: "Smaragdsturm - Entfesselung",
              Fr: "Castrum Marinum",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	763: {
		Name: LocalisedText{
              En: "Castrum Marinum (Extreme)",
              Ja: "極エメラルドウェポン破壊作戦",
              De: "Smaragdsturm - Trauma",
              Fr: "Castrum Marinum (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	764: {
		Name: LocalisedText{
              En: "The Great Ship Vylbrand",
              Ja: "バイルブランドの船出",
              De: "Gute Winde für Vylbrand",
              Fr: "Un navire nommé Vylbrand",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	765: {
		Name: LocalisedText{
              En: "Fit for a Queen",
              Ja: "ゴッド・セイブ・ザ・クイーン",
              De: "Hinab in die Ruinen",
              Fr: "Que les Dieux gardent la Reine",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	766: {
		Name: LocalisedText{
              En: "Novice Mahjong (Quick Ranked Match)",
              Ja: "ドマ式麻雀：東風戦一般卓（段位変動有り）",
              De: "Anfänger-Mahjong (schnelle Partie, gewertet)",
              Fr: "Mahjong domien: tous rangs (partie rapide classée)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	767: {
		Name: LocalisedText{
              En: "Advanced Mahjong (Quick Ranked Match)",
              Ja: "ドマ式麻雀：東風戦有段卓（段位変動有り）",
              De: "Fortgeschrittenen-Mahjong (schnelle Partie, gewertet)",
              Fr: "Mahjong domien: dan seulement (partie rapide classée)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	768: {
		Name: LocalisedText{
              En: "Four-player Mahjong (Quick Match, Kuitan Enabled)",
              Ja: "ドマ式麻雀：東風戦4人セット卓（クイタン有り）",
              De: "4-Spieler-Mahjong (schnelle Partie, Kuitan aktiviert)",
              Fr: "Mahjong domien: 4 joueurs (partie rapide avec kuitan)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	769: {
		Name: LocalisedText{
              En: "Four-player Mahjong (Quick Match, Kuitan Disabled)",
              Ja: "ドマ式麻雀：東風戦4人セット卓（クイタン無し）",
              De: "4-Spieler-Mahjong (schnelle Partie, Kuitan deaktiviert)",
              Fr: "Mahjong domien: 4 joueurs (partie rapide sans kuitan)",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	770: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	771: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	772: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	773: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	774: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	775: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	777: {
		Name: LocalisedText{
              En: "Paglth'an",
              Ja: "黄金平原 パガルザン",
              De: "Die Goldene Ebene von Paglth'an",
              Fr: "La grande prairie de Paglth'an",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	778: {
		Name: LocalisedText{
              En: "Zadnor",
              Ja: "ザトゥノル高原",
              De: "Zadnor-Hochebene",
              Fr: "Hauts plateaux de Zadnor",
            },
		HighEnd: false,
		ContentKind: ContentKindSavetheQueen,
	},
	779: {
		Name: LocalisedText{
              En: "The Tower at Paradigm's Breach",
              Ja: "希望ノ砲台：「塔」",
              De: "Der Turm, Paradigmenbrecher",
              Fr: "La tour de la Contingence",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	780: {
		Name: LocalisedText{
              En: "Death Unto Dawn",
              Ja: "黎明の死闘",
              De: "Kampf im Morgengrauen",
              Fr: "Aube meurtrière",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	781: {
		Name: LocalisedText{
              En: "The Cloud Deck",
              Ja: "ダイヤウェポン捕獲作戦",
              De: "Diamantblitz - Entfesselung",
              Fr: "Le Tillac des Cirrus",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	782: {
		Name: LocalisedText{
              En: "The Cloud Deck (Extreme)",
              Ja: "極ダイヤウェポン捕獲作戦",
              De: "Diamantblitz - Trauma",
              Fr: "Le Tillac des Cirrus (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	783: {
		Name: LocalisedText{
              En: "The Tower of Zot",
              Ja: "異形楼閣 ゾットの塔",
              De: "Der Turm von Zot",
              Fr: "La tour de Zott",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	784: {
		Name: LocalisedText{
              En: "The Stigma Dreamscape",
              Ja: "電脳夢想 スティグマ・フォー",
              De: "Stigma-Holometrie",
              Fr: "Rêve électrique de Stigma-4",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	785: {
		Name: LocalisedText{
              En: "The Tower of Babil",
              Ja: "魔導神門 バブイルの塔",
              De: "Der Turm von Babil",
              Fr: "La tour de Babil",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	786: {
		Name: LocalisedText{
              En: "The Aitiascope",
              Ja: "星海潜航 アイティオン星晶鏡",
              De: "Das Aitiaskop",
              Fr: "Le Prisme de l'Aitia",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	787: {
		Name: LocalisedText{
              En: "Ktisis Hyperboreia",
              Ja: "創造環境 ヒュペルボレア造物院",
              De: "Ktisis Hyperboreia",
              Fr: "L'Hyperborée",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	788: {
		Name: LocalisedText{
              En: "Dragonsong's Reprise (Ultimate)",
              Ja: "絶竜詩戦争",
              De: "Drachenkrieg (fatal)",
              Fr: "La Guerre du chant des dragons (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	789: {
		Name: LocalisedText{
              En: "Vanaspati",
              Ja: "終末樹海 ヴァナスパティ",
              De: "Vanaspati",
              Fr: "Vanaspati",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	790: {
		Name: LocalisedText{
              En: "The Mothercrystal",
              Ja: "ハイデリン討滅戦",
              De: "Prophetie - Hydaelyn",
              Fr: "Le Cristal-mère",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	791: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Hydaelyn's Call",
              Ja: "極ハイデリン討滅戦",
              De: "Eschatos - Hydaelyn",
              Fr: "Le Cristal-mère (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	792: {
		Name: LocalisedText{
              En: "The Dead Ends",
              Ja: "最終幻想 レムナント",
              De: "Das Sternengrab",
              Fr: "L'Issue aux Impasses",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	793: {
		Name: LocalisedText{
              En: "In from the Cold",
              Ja: "寒夜のこと",
              De: "In fremder Haut",
              Fr: "Le voleur de corps",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	794: {
		Name: LocalisedText{
              En: "Smileton",
              Ja: "楽園都市 スマイルトン",
              De: "Smileton",
              Fr: "Risette-sur-lune",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	795: {
		Name: LocalisedText{
              En: "Worthy of His Back",
              Ja: "前代アゼムの手ほどき",
              De: " Es reimt sich auf Gebell",
              Fr: "Le défi de l'ancienne Azem",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	796: {
		Name: LocalisedText{
              En: "The Final Day",
              Ja: "終焉の戦い",
              De: "Prophetie - Endsängerin",
              Fr: "Le Répons final",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	797: {
		Name: LocalisedText{
              En: "The Phantoms' Feast",
              Ja: "道化饗宴 ホーンテッドフィースト",
              De: "Lustiges Bankett",
              Fr: "Le banquet cauchemardesque",
            },
		HighEnd: false,
		ContentKind: ContentKind(22),
	},
	798: {
		Name: LocalisedText{
              En: "Endwalker",
              Ja: "暁月のフィナーレ",
              De: "Endschreiter",
              Fr: "Arpenteur des finitudes",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	799: {
		Name: LocalisedText{
              En: "To Calmer Seas",
              Ja: "融和への船出",
              De: "Im Hafen des Friedens",
              Fr: "Cap sur la paix",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	800: {
		Name: LocalisedText{
              En: "Asphodelos: The Fourth Circle",
              Ja: "万魔殿パンデモニウム：辺獄編4",
              De: "Asphodelos - Vierter Kreis",
              Fr: "Les Limbes du Pandæmonium - Abîme",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	801: {
		Name: LocalisedText{
              En: "Asphodelos: The Fourth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：辺獄編4",
              De: "Asphodelos - Vierter Kreis (episch)",
              Fr: "Les Limbes du Pandæmonium - Abîme (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	802: {
		Name: LocalisedText{
              En: "The Dark Inside",
              Ja: "ゾディアーク討滅戦",
              De: "Prophetie - Zodiark",
              Fr: "Le Cratère des Martyrs",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	803: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Zodiark's Fall",
              Ja: "極ゾディアーク討滅戦",
              De: "Eschatos - Zodiark",
              Fr: "Le Cratère des Martyrs (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	804: {
		Name: LocalisedText{
              En: "As the Heavens Burn",
              Ja: "拡がる終末",
              De: "Rote Himmel, roter Schnee",
              Fr: "L'arène des neiges",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	805: {
		Name: LocalisedText{
              En: "A Path Unveiled",
              Ja: "開かれた道の先へ",
              De: "Offen für neue Wege",
              Fr: "Des esprits et des hommes",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	806: {
		Name: LocalisedText{
              En: "Asphodelos: The Third Circle",
              Ja: "万魔殿パンデモニウム：辺獄編3",
              De: "Asphodelos - Dritter Kreis",
              Fr: "Les Limbes du Pandæmonium - Fournaise",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	807: {
		Name: LocalisedText{
              En: "Asphodelos: The Third Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：辺獄編3",
              De: "Asphodelos - Dritter Kreis (episch)",
              Fr: "Les Limbes du Pandæmonium - Fournaise (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	808: {
		Name: LocalisedText{
              En: "Asphodelos: The First Circle",
              Ja: "万魔殿パンデモニウム：辺獄編1",
              De: "Asphodelos - Erster Kreis",
              Fr: "Les Limbes du Pandæmonium - Parvis",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	809: {
		Name: LocalisedText{
              En: "Asphodelos: The First Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：辺獄編1",
              De: "Asphodelos - Erster Kreis (episch)",
              Fr: "Les Limbes du Pandæmonium - Parvis (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	810: {
		Name: LocalisedText{
              En: "Asphodelos: The Second Circle",
              Ja: "万魔殿パンデモニウム：辺獄編2",
              De: "Asphodelos - Zweiter Kreis",
              Fr: "Les Limbes du Pandæmonium - Cloaque",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	811: {
		Name: LocalisedText{
              En: "Asphodelos: The Second Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：辺獄編2",
              De: "Asphodelos - Zweiter Kreis (episch)",
              Fr: "Les Limbes du Pandæmonium - Cloaque (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	812: {
		Name: LocalisedText{
              En: "A Frosty Reception",
              Ja: "霜雪を踏みしめて",
              De: "Ein frostiger Empfang",
              Fr: "Un accueil glacial",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	813: {
		Name: LocalisedText{
              En: "Sage's Focus",
              Ja: "賢者の短杖",
              De: "Des Weisen wundersames Werkzeug",
              Fr: "Les armes du sage",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	814: {
		Name: LocalisedText{
              En: "The Harvest Begins",
              Ja: "大鎌の意味",
              De: "Die Bedeutung der Sense",
              Fr: "La vraie puissance de la faux",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	815: {
		Name: LocalisedText{
              En: "Ever March Heavensward",
              Ja: "蒼天を仰ぎ、歩み続ける",
              De: "Der Weg zur Erneuerung",
              Fr: "La voie du renouveau",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	816: {
		Name: LocalisedText{
              En: "The Killing Art",
              Ja: "暗殺道",
              De: "Die Kunst des Tötens",
              Fr: "La voie du néant",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	817: {
		Name: LocalisedText{
              En: "Life Ephemeral, Path Eternal",
              Ja: "人命は儚く、術のみちは永久に",
              De: "Das Leben ist kurz, die Kunst ist lang",
              Fr: "Existences éphémères et savoir éternel",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	818: {
		Name: LocalisedText{
              En: "Laid to Rest",
              Ja: "ドマの弔い",
              De: "Domanisches Begräbnis",
              Fr: "Des adieux domiens",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	819: {
		Name: LocalisedText{
              En: "The Excitatron 6000",
              Ja: "宝物庫 エキサイトロン",
              De: "Euphoratron",
              Fr: "Le Ludodrome",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	820: {
		Name: LocalisedText{
              En: "The Gift of Mercy",
              Ja: "僕たちは還り、君を見送ろう",
              De: "Trauer und Hoffnung",
              Fr: "Acceptation",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	822: {
		Name: LocalisedText{
              En: "The Aetherfont",
              Ja: "星霊間欠 ハーム島",
              De: "Ätherborn",
              Fr: "L'île de Haam",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	823: {
		Name: LocalisedText{
              En: "The Lunar Subterrane",
              Ja: "深淵潜行 月の地下渓谷",
              De: "Monduntergrund",
              Fr: "Le Souterrain lunaire",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	824: {
		Name: LocalisedText{
              En: "Worqor Zormor",
              Ja: "山嶺登頂 ウォーコー・ゾーモー",
              De: "Worqor Zormor",
              Fr: "Worqor Zormor",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	825: {
		Name: LocalisedText{
              En: "Origenics",
              Ja: "魂魄工廠 オリジェニクス",
              De: "Origenik",
              Fr: "L'Origenèse",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	826: {
		Name: LocalisedText{
              En: "Ihuykatumu",
              Ja: "濁流遡上 イフイカ・トゥム",
              De: "Ihuykatumu",
              Fr: "La Remontée de l'Ihuykatumu",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	827: {
		Name: LocalisedText{
              En: "Alexandria",
              Ja: "記憶幻想 アレクサンドリア",
              De: "Alexandria",
              Fr: "Alexandrie",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	828: {
		Name: LocalisedText{
              En: "A Father First",
              Ja: "王として、父として",
              De: "Des Vaters viele Sorgen",
              Fr: "L'Aurarque et le père",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	829: {
		Name: LocalisedText{
              En: "The Skydeep Cenote",
              Ja: "遺産踏査 天深きセノーテ",
              De: "Himmelstiefer Cenote",
              Fr: "Le cénote des Cieux infinis",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	830: {
		Name: LocalisedText{
              En: "The Porta Decumana",
              Ja: "アルテマウェポン破壊作戦",
              De: "Porta Decumana",
              Fr: "Porta Decumana",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	831: {
		Name: LocalisedText{
              En: "Vanguard",
              Ja: "外征前哨 ヴァンガード",
              De: "Der Außenposten",
              Fr: "L'Avant-garde",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	832: {
		Name: LocalisedText{
              En: "Worqor Lar Dor",
              Ja: "ヴァリガルマンダ討滅戦",
              De: "Zel Tajaal - Valigarmanda",
              Fr: "Worqor Lar Dor",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	833: {
		Name: LocalisedText{
              En: "Worqor Lar Dor (Extreme)",
              Ja: "極ヴァリガルマンダ討滅戦",
              De: "Gok Tajaal - Valigarmanda",
              Fr: "Worqor Lar Dor (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	834: {
		Name: LocalisedText{
              En: "Tender Valley",
              Ja: "荒野秘境 サボテンダーバレー",
              De: "Kaktorsenke",
              Fr: "La Vallée des Pampas",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	835: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	836: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	837: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	838: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	839: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	840: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	841: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	842: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	843: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	844: {
		Name: LocalisedText{
              En: "Alzadaal's Legacy",
              Ja: "近東秘宝 アルザダール海底遺跡群",
              De: "Alzadaals Vermächtnis",
              Fr: "Le legs d'Alzadaal",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	845: {
		Name: LocalisedText{
              En: "Forlorn Glory",
              Ja: "ネルウァの帝国",
              De: "Nervas Reich",
              Fr: "Nerva, le patriote",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	846: {
		Name: LocalisedText{
              En: "The Minstrel's Ballad: Endsinger's Aria",
              Ja: "終極の戦い",
              De: "Eschatos - Endsängerin",
              Fr: "Le Répons final (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	847: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	848: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	849: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	850: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	851: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	852: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	853: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	854: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	855: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	856: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	857: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	858: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	859: {
		Name: LocalisedText{
              En: "The Palaistra",
              Ja: "パライストラ",
              De: "Die Palästra",
              Fr: "Le Palestre",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	860: {
		Name: LocalisedText{
              En: "The Volcanic Heart",
              Ja: "ヴォルカニック・ハート",
              De: "Das Herz des Vulkans",
              Fr: "Le Cœur volcanique",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	861: {
		Name: LocalisedText{
              En: "Cloud Nine",
              Ja: "クラウドナイン",
              De: "Wolke Sieben",
              Fr: "Le Petit Nuage",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	862: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Custom Match - The Palaistra)",
              Ja: "クリスタルコンフリクト(パライストラ：カスタムマッチ)",
              De: "Crystalline Conflict: Die Palästra (Schaukampf)",
              Fr: "Crystalline Conflict (partie personnalisée - Le Palestre)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	863: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Custom Match - The Volcanic Heart)",
              Ja: "クリスタルコンフリクト(ヴォルカニック・ハート：カスタムマッチ)",
              De: "Crystalline Conflict: Das Herz des Vulkans (Schaukampf)",
              Fr: "Crystalline Conflict (partie personnalisée - Le Cœur volcanique)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	864: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Custom Match - Cloud Nine)",
              Ja: "クリスタルコンフリクト(クラウドナイン：カスタムマッチ)",
              De: "Crystalline Conflict: Wolke Sieben (Schaukampf)",
              Fr: "Crystalline Conflict (partie personnalisée - Le Petit Nuage)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	865: {
		Name: LocalisedText{
              En: "Cape Westwind",
              Ja: "リットアティン強襲戦",
              De: "Kap Westwind",
              Fr: "Le Cap Vendouest",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	866: {
		Name: LocalisedText{
              En: "Aglaia",
              Ja: "輝ける神域 アグライア",
              De: "Aglaia",
              Fr: "Domaine divin - Aglaé",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	867: {
		Name: LocalisedText{
              En: "Devastation",
              Ja: "アシエン・ラハブレア討伐戦",
              De: "Das Antlitz der Kaltblütigkeit",
              Fr: "Le Praetorium en flammes",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	868: {
		Name: LocalisedText{
              En: "The Sil'dihn Subterrane",
              Ja: "ヴァリアントダンジョン シラディハ水道",
              De: "Die Unterstadt von Sil'dih",
              Fr: "Les canalisations sildiennes - Donjon à embranchements",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	869: {
		Name: LocalisedText{
              En: "The Fell Court of Troia",
              Ja: "異界孤城 トロイアコート",
              De: "Der Schwarze Hof von Troia",
              Fr: "Le Château de Troïa",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	870: {
		Name: LocalisedText{
              En: "Storm's Crown",
              Ja: "バルバリシア討滅戦",
              De: "Prophetie - Barbarizia",
              Fr: "La Toison des tempêtes",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	871: {
		Name: LocalisedText{
              En: "Storm's Crown (Extreme)",
              Ja: "極バルバリシア討滅戦",
              De: "Eschatos - Barbarizia",
              Fr: "La Toison des tempêtes (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	872: {
		Name: LocalisedText{
              En: "Abyssos: The Fifth Circle",
              Ja: "万魔殿パンデモニウム：煉獄編1",
              De: "Abyssos - Fünfter Kreis",
              Fr: "Le Purgatoire du Pandæmonium - Cages",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	873: {
		Name: LocalisedText{
              En: "Abyssos: The Fifth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：煉獄編1",
              De: "Abyssos - Fünfter Kreis (episch)",
              Fr: "Le Purgatoire du Pandæmonium - Cages (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	874: {
		Name: LocalisedText{
              En: "Where Everything Begins",
              Ja: "失われた力",
              De: "Wo alles seinen Anfang nimmt",
              Fr: "Au commencement était Zero",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	876: {
		Name: LocalisedText{
              En: "Abyssos: The Seventh Circle",
              Ja: "万魔殿パンデモニウム：煉獄編3",
              De: "Abyssos - Siebter Kreis",
              Fr: "Le Purgatoire du Pandæmonium - Racines",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	877: {
		Name: LocalisedText{
              En: "Abyssos: The Seventh Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：煉獄編3",
              De: "Abyssos - Siebter Kreis (episch)",
              Fr: "Le Purgatoire du Pandæmonium - Racines (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	878: {
		Name: LocalisedText{
              En: "Another Sil'dihn Subterrane",
              Ja: "アナザーダンジョン 異聞シラディハ水道",
              De: "Die kuriose Unterstadt von Sil'dih",
              Fr: "Les canalisations sildiennes annexes - Donjon alternatif",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	879: {
		Name: LocalisedText{
              En: "Another Sil'dihn Subterrane (Savage)",
              Ja: "アナザーダンジョン 異聞シラディハ水道 零式",
              De: "Die kuriose Unterstadt von Sil'dih (episch)",
              Fr: "Les canalisations sildiennes annexes - Donjon alternatif (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	880: {
		Name: LocalisedText{
              En: "Abyssos: The Sixth Circle",
              Ja: "万魔殿パンデモニウム：煉獄編2",
              De: "Abyssos - Sechster Kreis",
              Fr: "Le Purgatoire du Pandæmonium - Croisements",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	881: {
		Name: LocalisedText{
              En: "Abyssos: The Sixth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：煉獄編2",
              De: "Abyssos - Sechster Kreis (episch)",
              Fr: "Le Purgatoire du Pandæmonium - Croisements (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	883: {
		Name: LocalisedText{
              En: "Abyssos: The Eighth Circle",
              Ja: "万魔殿パンデモニウム：煉獄編4",
              De: "Abyssos - Achter Kreis",
              Fr: "Le Purgatoire du Pandæmonium - Hérédité",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	884: {
		Name: LocalisedText{
              En: "Abyssos: The Eighth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：煉獄編4",
              De: "Abyssos - Achter Kreis (episch)",
              Fr: "Le Purgatoire du Pandæmonium - Hérédité (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	885: {
		Name: LocalisedText{
              En: "The Steps of Faith",
              Ja: "皇都イシュガルド防衛戦",
              De: "Der Schicksalsweg",
              Fr: "Le Siège de la sainte Cité d'Ishgard",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	886: {
		Name: LocalisedText{
              En: "Mount Ordeals",
              Ja: "ルビカンテ討滅戦",
              De: "Prophetie - Rubicante",
              Fr: "Le Mont du Supplice",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	896: {
		Name: LocalisedText{
              En: "Lapis Manalis",
              Ja: "雪山冥洞 ラピス・マナリス",
              De: "Lapis Manalis",
              Fr: "Lapis Manalis",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	897: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 1-10)",
              Ja: "オルト・エウレカ B1～B10",
              De: "Eureka Orthos (Ebenen 1-10)",
              Fr: "Eurêka Orthos (sous-sols 1-10)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	898: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 11-20)",
              Ja: "オルト・エウレカ B11～B20",
              De: "Eureka Orthos (Ebenen 11-20)",
              Fr: "Eurêka Orthos (sous-sols 11-20)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	899: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 21-30)",
              Ja: "オルト・エウレカ B21～B30",
              De: "Eureka Orthos (Ebenen 21-30)",
              Fr: "Eurêka Orthos (sous-sols 21-30)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	900: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 31-40)",
              Ja: "オルト・エウレカ B31～B40",
              De: "Eureka Orthos (Ebenen 31-40)",
              Fr: "Eurêka Orthos (sous-sols 31-40)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	901: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 41-50)",
              Ja: "オルト・エウレカ B41～B50",
              De: "Eureka Orthos (Ebenen 41-50)",
              Fr: "Eurêka Orthos (sous-sols 41-50)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	902: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 51-60)",
              Ja: "オルト・エウレカ B51～B60",
              De: "Eureka Orthos (Ebenen 51-60)",
              Fr: "Eurêka Orthos (sous-sols 51-60)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	903: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 61-70)",
              Ja: "オルト・エウレカ B61～B70",
              De: "Eureka Orthos (Ebenen 61-70)",
              Fr: "Eurêka Orthos (sous-sols 61-70)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	904: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 71-80)",
              Ja: "オルト・エウレカ B71～B80",
              De: "Eureka Orthos (Ebenen 71-80)",
              Fr: "Eurêka Orthos (sous-sols 71-80)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	905: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 81-90)",
              Ja: "オルト・エウレカ B81～B90",
              De: "Eureka Orthos (Ebenen 81-90)",
              Fr: "Eurêka Orthos (sous-sols 81-90)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	906: {
		Name: LocalisedText{
              En: "Eureka Orthos (Floors 91-100)",
              Ja: "オルト・エウレカ B91～B100",
              De: "Eureka Orthos (Ebenen 91-100)",
              Fr: "Eurêka Orthos (sous-sols 91-100)",
            },
		HighEnd: false,
		ContentKind: ContentKindDeepDungeons,
	},
	907: {
		Name: LocalisedText{
              En: "Generational Bonding",
              Ja: "はつらつとした親子",
              De: "Väterliche Liebe",
              Fr: "Tel gentilhomme, tel gentilhomme",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	908: {
		Name: LocalisedText{
              En: "The Omega Protocol (Ultimate)",
              Ja: "絶オメガ検証戦",
              De: "Omega (fatal)",
              Fr: "Le Protocole Oméga (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	909: {
		Name: LocalisedText{
              En: "The Shifting Gymnasion Agonon",
              Ja: "宝物庫 エルピス・ギュムナシオン祭殿",
              De: "Gymnasion Agonon",
              Fr: "Elpis Gymnasion",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	910: {
		Name: LocalisedText{
              En: "An Unforeseen Bargain",
              Ja: "パンひとつと引き換えに",
              De: "Der Wert einer Buuds",
              Fr: "Pour une bouchée de pain",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	911: {
		Name: LocalisedText{
              En: "Euphrosyne",
              Ja: "喜びの神域 エウプロシュネ",
              De: "Euphrosyne",
              Fr: "Domaine divin - Euphrosyne",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	912: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	913: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	914: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	915: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	916: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	917: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	918: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	919: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	920: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	921: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	922: {
		Name: LocalisedText{
              En: "The Clockwork Castletown",
              Ja: "東方絡繰御殿",
              De: "Die Mechanische Menagerie",
              Fr: "Le Traquenard oriental",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	923: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Custom Match - The Clockwork Castletown)",
              Ja: "クリスタルコンフリクト(東方絡繰御殿：カスタムマッチ)",
              De: "Crystalline Conflict: Die Mechanische Menagerie (Schaukampf)",
              Fr: "Crys. Conflict (partie perso. - Le Traquenard oriental)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	924: {
		Name: LocalisedText{
              En: "Mount Ordeals (Extreme)",
              Ja: "極ルビカンテ討滅戦",
              De: "Eschatos - Rubicante",
              Fr: "Le Mont du Supplice (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	925: {
		Name: LocalisedText{
              En: "Fangs of the Viper",
              Ja: "ヴァイパーの使命",
              De: "Einführung in die Viperologie",
              Fr: "Le devoir des rôdeurs vipère",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	927: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	928: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	929: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	930: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	931: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	932: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	933: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	934: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	935: {
		Name: LocalisedText{
              En: "Leap of Faith",
              Ja: "挑戦！ ジャンピングアスレチック",
              De: "Kaktor-Kletterwand",
              Fr: "Haute voltige",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	936: {
		Name: LocalisedText{
              En: "Anabaseios: The Ninth Circle",
              Ja: "万魔殿パンデモニウム：天獄編1",
              De: "Anabaseios - Neunter Kreis",
              Fr: "Le Paradis du Pandæmonium - Métempsycose",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	937: {
		Name: LocalisedText{
              En: "Anabaseios: The Ninth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：天獄編1",
              De: "Anabaseios - Neunter Kreis (episch)",
              Fr: "Le Paradis du Pandæmonium - Métempsycose (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	938: {
		Name: LocalisedText{
              En: "Anabaseios: The Tenth Circle",
              Ja: "万魔殿パンデモニウム：天獄編2",
              De: "Anabaseios - Zehnter Kreis",
              Fr: "Le Paradis du Pandæmonium - Monochrome",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	939: {
		Name: LocalisedText{
              En: "Anabaseios: The Tenth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：天獄編2",
              De: "Anabaseios - Zehnter Kreis (episch)",
              Fr: "Le Paradis du Pandæmonium - Monochrome (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	940: {
		Name: LocalisedText{
              En: "Anabaseios: The Eleventh Circle",
              Ja: "万魔殿パンデモニウム：天獄編3",
              De: "Anabaseios - Elfter Kreis",
              Fr: "Le Paradis du Pandæmonium - Tribunal",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	941: {
		Name: LocalisedText{
              En: "Anabaseios: The Eleventh Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：天獄編3",
              De: "Anabaseios - Elfter Kreis (episch)",
              Fr: "Le Paradis du Pandæmonium - Tribunal (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	942: {
		Name: LocalisedText{
              En: "Anabaseios: The Twelfth Circle",
              Ja: "万魔殿パンデモニウム：天獄編4",
              De: "Anabaseios - Zwölfter Kreis",
              Fr: "Le Paradis du Pandæmonium - Apothéose",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	943: {
		Name: LocalisedText{
              En: "Anabaseios: The Twelfth Circle (Savage)",
              Ja: "万魔殿パンデモニウム零式：天獄編4",
              De: "Anabaseios - Zwölfter Kreis (episch)",
              Fr: "Le Paradis du Pandæmonium - Apothéose (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	944: {
		Name: LocalisedText{
              En: "The Gilded Araya",
              Ja: "アスラ討滅戦",
              De: "Prophetie - Asura",
              Fr: "Le temple doré d'Araya",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	945: {
		Name: LocalisedText{
              En: "Mount Rokkon",
              Ja: "ヴァリアントダンジョン 六根山",
              De: "Der Rokkon",
              Fr: "Le mont Rokkon - Donjon à embranchements",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	946: {
		Name: LocalisedText{
              En: "Another Mount Rokkon",
              Ja: "アナザーダンジョン 異聞六根山",
              De: "Der kuriose Rokkon",
              Fr: "Le mont Rokkon annexe - Donjon alternatif",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	947: {
		Name: LocalisedText{
              En: "Another Mount Rokkon (Savage)",
              Ja: "アナザーダンジョン 異聞六根山 零式",
              De: "Der kuriose Rokkon (episch)",
              Fr: "Le mont Rokkon annexe - Donjon alternatif (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	948: {
		Name: LocalisedText{
              En: "A Golden Opportunity",
              Ja: "黄金闘士「ゴールドル」",
              De: "Goldor der Güldene",
              Fr: "Goldor, le mage doré",
            },
		HighEnd: false,
		ContentKind: ContentKindTheMaskedCarnivale,
	},
	949: {
		Name: LocalisedText{
              En: "The Voidcast Dais",
              Ja: "ゴルベーザ討滅戦",
              De: "Prophetie - Golbez",
              Fr: "La Chaire de l'Exilée",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	950: {
		Name: LocalisedText{
              En: "The Voidcast Dais (Extreme)",
              Ja: "極ゴルベーザ討滅戦",
              De: "Eschatos - Golbez",
              Fr: "La Chaire de l'Exilée (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	952: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	953: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	954: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	955: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	956: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	957: {
		Name: LocalisedText{
              En: "Ocean Fishing",
              Ja: "オーシャンフィッシング",
              De: "Auf großer Fahrt",
              Fr: "Pêche en mer",
            },
		HighEnd: false,
		ContentKind: ContentKindDisciplesoftheLand,
	},
	958: {
		Name: LocalisedText{
              En: "Blunderville",
              Ja: "ブランダーヴィル",
              De: "Stolperville",
              Fr: "Blunderville",
            },
		HighEnd: false,
		ContentKind: ContentKindGoldSaucer,
	},
	959: {
		Name: LocalisedText{
              En: "Memory of Embers",
              Ja: "炎影の旅路",
              De: "Der Pfad der Flammen",
              Fr: "Flammes primordiales",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	960: {
		Name: LocalisedText{
              En: "The Game Is Afoot",
              Ja: "その瞳が見据えるもの",
              De: "Die Jagd ist eröffnet",
              Fr: "Par-delà l'horizon",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	961: {
		Name: LocalisedText{
              En: "Aloalo Island",
              Ja: "ヴァリアントダンジョン アロアロ島",
              De: "Aloalo",
              Fr: "L'île d'Aloalo - Donjon à embranchements",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	962: {
		Name: LocalisedText{
              En: "Thaleia",
              Ja: "華めく神域 タレイア",
              De: "Thaleia",
              Fr: "Domaine divin - Thalie",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	964: {
		Name: LocalisedText{
              En: "The Abyssal Fracture",
              Ja: "ゼロムス討滅戦",
              De: "Prophetie - Zeromus",
              Fr: "La Fracture abyssale",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	965: {
		Name: LocalisedText{
              En: "The Abyssal Fracture (Extreme)",
              Ja: "極ゼロムス討滅戦",
              De: "Eschatos - Zeromus",
              Fr: "La Fracture abyssale (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	966: {
		Name: LocalisedText{
              En: "Somewhere Only She Knows",
              Ja: "旅する画家は幻想を征く",
              De: "Begegnungen und Abschiede",
              Fr: "L'œuvre de la maturité",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	967: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	968: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	969: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	970: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	971: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	972: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	973: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	974: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	975: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	976: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	977: {
		Name: LocalisedText{
              En: "The Red Sands",
              Ja: "レッド・サンズ",
              De: "Die Roten Sande",
              Fr: "Les Sables sanglants",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	978: {
		Name: LocalisedText{
              En: "Crystalline Conflict (Custom Match - The Red Sands)",
              Ja: "クリスタルコンフリクト(レッド・サンズ：カスタムマッチ)",
              De: "Crystalline Conflict: Die Roten Sande (Schaukampf)",
              Fr: "Crystalline Conflict (partie personnalisée - Les Sables sanglants)",
            },
		HighEnd: false,
		ContentKind: ContentKindPvP,
	},
	979: {
		Name: LocalisedText{
              En: "Another Aloalo Island",
              Ja: "アナザーダンジョン 異聞アロアロ島",
              De: "Kurioses Aloalo",
              Fr: "L'île d'Aloalo annexe - Donjon alternatif",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	980: {
		Name: LocalisedText{
              En: "Another Aloalo Island (Savage)",
              Ja: "アナザーダンジョン 異聞アロアロ島 零式",
              De: "Kurioses Aloalo (episch)",
              Fr: "L'île d'Aloalo annexe - Donjon alternatif (sadique)",
            },
		HighEnd: false,
		ContentKind: ContentKindVCDungeonFinder,
	},
	981: {
		Name: LocalisedText{
              En: "The Strayborough Deadwalk",
              Ja: "悪夢遊園 ストレイバロー",
              De: "Mahrlingen-Promenade",
              Fr: "Le Parc-aux-Errants",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	982: {
		Name: LocalisedText{
              En: "The Feat of the Brotherhood",
              Ja: "友の試練",
              De: "Von Schuld und Sühne",
              Fr: "L'épreuve de l'amitié",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	983: {
		Name: LocalisedText{
              En: "Taking a Stand",
              Ja: "笑顔を護るための戦い",
              De: "Mit scharfer Axt und leichtem Lächeln",
              Fr: "Préserver la paix",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	984: {
		Name: LocalisedText{
              En: "The Interphos",
              Ja: "エターナルクイーン討滅戦",
              De: "Zel Tajaal - Ewige Königin",
              Fr: "Interphos",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	985: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M1",
              Ja: "至天の座アルカディア：ライトヘビー級1",
              De: "Arkadion - Halbschwergewicht R1",
              Fr: "Poids mi-lourds CCA - match 1",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	986: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M1 (Savage)",
              Ja: "至天の座アルカディア零式：ライトヘビー級1",
              De: "Arkadion - Halbschwergewicht R1 (episch)",
              Fr: "Poids mi-lourds CCA - match 1 (sadique)",
            },
		HighEnd: true,
		ContentKind: ContentKindRaids,
	},
	987: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M2",
              Ja: "至天の座アルカディア：ライトヘビー級2",
              De: "Arkadion - Halbschwergewicht R2",
              Fr: "Poids mi-lourds CCA - match 2",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	988: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M2 (Savage)",
              Ja: "至天の座アルカディア零式：ライトヘビー級2",
              De: "Arkadion - Halbschwergewicht R2 (episch)",
              Fr: "Poids mi-lourds CCA - match 2 (sadique)",
            },
		HighEnd: true,
		ContentKind: ContentKindRaids,
	},
	989: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M3",
              Ja: "至天の座アルカディア：ライトヘビー級3",
              De: "Arkadion - Halbschwergewicht R3",
              Fr: "Poids mi-lourds CCA - match 3",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	990: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M3 (Savage)",
              Ja: "至天の座アルカディア零式：ライトヘビー級3",
              De: "Arkadion - Halbschwergewicht R3 (episch)",
              Fr: "Poids mi-lourds CCA - match 3 (sadique)",
            },
		HighEnd: true,
		ContentKind: ContentKindRaids,
	},
	991: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M4",
              Ja: "至天の座アルカディア：ライトヘビー級4",
              De: "Arkadion - Halbschwergewicht R4",
              Fr: "Poids mi-lourds CCA - match 4",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	992: {
		Name: LocalisedText{
              En: "AAC Light-heavyweight M4 (Savage)",
              Ja: "至天の座アルカディア零式：ライトヘビー級4",
              De: "Arkadion - Halbschwergewicht R4 (episch)",
              Fr: "Poids mi-lourds CCA - match 4 (sadique)",
            },
		HighEnd: true,
		ContentKind: ContentKindRaids,
	},
	993: {
		Name: LocalisedText{
              En: "Cenote Ja Ja Gural",
              Ja: "宝物庫 セノーテ・ジャジャグラル",
              De: "Cenote Ja Ja Gural",
              Fr: "Cénote Ja Ja Gural",
            },
		HighEnd: false,
		ContentKind: ContentKindTreasureHunt,
	},
	994: {
		Name: LocalisedText{
              En: "An Antidote for Anarchy",
              Ja: "毒と癒やしの頂上決戦",
              De: "Heilkunst gegen Hexerei",
              Fr: "Les deux faces du scorpion",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	995: {
		Name: LocalisedText{
              En: "Everkeep",
              Ja: "ゾラージャ討滅戦",
              De: "Zel Tajaal - Zoraal Ja",
              Fr: "Le Pinacle de l'Éternité",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	996: {
		Name: LocalisedText{
              En: "Everkeep (Extreme)",
              Ja: "極ゾラージャ討滅戦",
              De: "Gok Tajaal - Zoraal Ja",
              Fr: "Le Pinacle de l'Éternité (extrême)",
            },
		HighEnd: false,
		ContentKind: ContentKindTrials,
	},
	997: {
		Name: LocalisedText{
              En: "A Hunter True",
              Ja: "一端の狩人",
              De: "Krönender Abschuss",
              Fr: "Chasseuse à lunettes... chasseuse honnête",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	998: {
		Name: LocalisedText{
              En: "The Protector and the Destroyer",
              Ja: "護る者、壊す者",
              De: "Schützer des Volkes, Schlächter des Volkes",
              Fr: "Protecteurs et destructeurs",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	999: {
		Name: LocalisedText{
              En: "The Mightiest Shield",
              Ja: "打倒！ シャルトフィット盗賊団",
              De: "Der mächtigste Schild",
              Fr: "Mauvais cru pour Chalteaufite",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1000: {
		Name: LocalisedText{
              En: "Mind over Manor",
              Ja: "旅する画家は森都を征く",
              De: "Nächtliche Ruhestörung",
              Fr: "Tapage nocturne",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1001: {
		Name: LocalisedText{
              En: "Heroes and Pretenders",
              Ja: "祭器と作られた虚像",
              De: "Ein Mythos in Trümmern",
              Fr: "Le château de cartes s'effondre",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1002: {
		Name: LocalisedText{
              En: "Vengeance of the Viper",
              Ja: "トラルヴィドラールを狩る者",
              De: "Der Wille der Natur",
              Fr: "Le fléau des Tural vidraal",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1003: {
		Name: LocalisedText{
              En: "Dreams of a New Day",
              Ja: "寝ても覚めても",
              De: "Neue Bande der Freundschaft",
              Fr: "Rêves éveillés",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1006: {
		Name: LocalisedText{
              En: "Futures Rewritten (Ultimate)",
              Ja: "絶もうひとつの未来",
              De: "Eine zweite Zukunft (fatal)",
              Fr: "Avenirs réécrits (fatal)",
            },
		HighEnd: true,
		ContentKind: ContentKindUltimateRaids,
	},
	1007: {
		Name: LocalisedText{
              En: "The Jade Stoa (Unreal)",
              Ja: "幻白虎征魂戦",
              De: "Traumprüfung - Byakko",
              Fr: "La Clairière de Jade (irréel)",
            },
		HighEnd: true,
		ContentKind: ContentKindTrials,
	},
	1008: {
		Name: LocalisedText{
              En: "Yuweyawata Field Station",
              Ja: "廃地討究 ユウェヤーワータ",
              De: "Forschungsstation Yuweyawata",
              Fr: "Le centre de recherche de Yuweyawata",
            },
		HighEnd: false,
		ContentKind: ContentKindDungeons,
	},
	1009: {
		Name: LocalisedText{
              En: "The Warmth of Family",
              Ja: "王の家族",
              De: "Ein Land, eine Familie",
              Fr: "La famille de l'Aurarque",
            },
		HighEnd: false,
		ContentKind: ContentKindQuestBattles,
	},
	1012: {
		Name: LocalisedText{
              En: "React to Attack Markers",
              Ja: "マーカーが出る攻撃に対処しよう！",
              De: "Achte auf Attackenmarkierungen!",
              Fr: "Réagir à des marquages",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	1013: {
		Name: LocalisedText{
              En: "React to Floor Markers",
              Ja: "地面に予兆が出る攻撃に対処しよう！",
              De: "Achte auf Geländemarkierungen!",
              Fr: "Réagir à des indications au sol",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	1014: {
		Name: LocalisedText{
              En: "React to Advanced Visual Indicators",
              Ja: "さまざまな攻撃に対処しよう！",
              De: "Stelle dich allerlei Angriffsmanövern!",
              Fr: "Réagir à des mécaniques variées",
            },
		HighEnd: false,
		ContentKind: ContentKind(20),
	},
	1015: {
		Name: LocalisedText{
              En: "Jeuno: The First Walk",
              Ja: "ジュノ：ザ・ファーストウォーク",
              De: "Jeuno: Die erste Etappe",
              Fr: "Jeuno - La première perambulation",
            },
		HighEnd: false,
		ContentKind: ContentKindRaids,
	},
	1017: {
		Name: LocalisedText{
              En: "The Minstrel’s Ballad: Sphene's Burden",
              Ja: "極エターナルクイーン討滅戦",
              De: "Gok Tajaal - Ewige Königin",
              Fr: "Interphos (extrême)",
            },
		HighEnd: true,
		ContentKind: ContentKindTrials,
	},
}
