package ffxiv
type ContentKind uint32

const (
	ContentKindDutyRoulette ContentKind = 1 + iota
	ContentKindDungeons
	ContentKindGuildhests
	ContentKindTrials
	ContentKindRaids
	ContentKindPvP
	ContentKindQuestBattles
	ContentKindFATEs
	ContentKindTreasureHunt
	ContentKindLevequests
	ContentKindGrandCompany
	ContentKindCompanions
	ContentKindTribalQuests
	ContentKindOverallCompletion
	ContentKindPlayerCommendation
	ContentKindDisciplesoftheLand
	ContentKindDisciplesoftheHand
	ContentKindRetainerVentures
	ContentKindGoldSaucer
	_ // 20
	ContentKindDeepDungeons
	_ // 22
	_ // 23
	ContentKindWondrousTails
	ContentKindCustomDeliveries
	ContentKindEureka
	ContentKindTheMaskedCarnivale
	ContentKindUltimateRaids
	ContentKindSavetheQueen
	ContentKindVCDungeonFinder
	ContentKindOceanFishing
	ContentKindTripleTriad
	ContentKindTheHunt
	ContentKindFishing
	ContentKindGATE
	ContentKindIslandSanctuary
	ContentKindChaoticAllianceRaid
	ContentKindOther
)

func FromUint32(kind uint32) ContentKind {
	if kind >= 1 && kind < 38 {
		return ContentKind(kind)
	}
	return ContentKindOther
}

func (k ContentKind) AsUint32() uint32 {
	return uint32(k)
}

func (k ContentKind) String() string {
	switch k {
	case ContentKindDutyRoulette:
		return "DutyRoulette"
	case ContentKindDungeons:
		return "Dungeons"
	case ContentKindGuildhests:
		return "Guildhests"
	case ContentKindTrials:
		return "Trials"
	case ContentKindRaids:
		return "Raids"
	case ContentKindPvP:
		return "PvP"
	case ContentKindQuestBattles:
		return "QuestBattles"
	case ContentKindFATEs:
		return "FATEs"
	case ContentKindTreasureHunt:
		return "TreasureHunt"
	case ContentKindLevequests:
		return "Levequests"
	case ContentKindGrandCompany:
		return "GrandCompany"
	case ContentKindCompanions:
		return "Companions"
	case ContentKindTribalQuests:
		return "TribalQuests"
	case ContentKindOverallCompletion:
		return "OverallCompletion"
	case ContentKindPlayerCommendation:
		return "PlayerCommendation"
	case ContentKindDisciplesoftheLand:
		return "DisciplesoftheLand"
	case ContentKindDisciplesoftheHand:
		return "DisciplesoftheHand"
	case ContentKindRetainerVentures:
		return "RetainerVentures"
	case ContentKindGoldSaucer:
		return "GoldSaucer"
	case ContentKindDeepDungeons:
		return "DeepDungeons"
	case ContentKindWondrousTails:
		return "WondrousTails"
	case ContentKindCustomDeliveries:
		return "CustomDeliveries"
	case ContentKindEureka:
		return "Eureka"
	case ContentKindTheMaskedCarnivale:
		return "TheMaskedCarnivale"
	case ContentKindUltimateRaids:
		return "UltimateRaids"
	case ContentKindSavetheQueen:
		return "SavetheQueen"
	case ContentKindVCDungeonFinder:
		return "VCDungeonFinder"
	case ContentKindOceanFishing:
		return "OceanFishing"
	case ContentKindTripleTriad:
		return "TripleTriad"
	case ContentKindTheHunt:
		return "TheHunt"
	case ContentKindFishing:
		return "Fishing"
	case ContentKindGATE:
		return "GATE"
	case ContentKindIslandSanctuary:
		return "IslandSanctuary"
	case ContentKindChaoticAllianceRaid:
		return "ChaoticAllianceRaid"
	default:
		return "Other"
	}
}
