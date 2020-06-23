package database

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_UpdateUserStatsForGame(t *testing.T) {
	tests := []struct {
		name    string
		stats   steamclient.UserStatsForGame
		want    steamclient.UserStatsForGame
		wantErr bool
	}{

		{
			name: "Update all UserStatsForGame (ID: all_columns)",
			stats: steamclient.UserStatsForGame{
				SteamID: "all_columns",
				// GameName: "1",
				Stats: steamclient.GameStats{
					SteamID:                                   "all_columns",
					GILessonBombSitesA:                        "changed-1",
					GILessonBombSitesB:                        "changed-2",
					GILessonCsgoCycleWeaponsKb:                "changed-3",
					GILessonCsgoHostageLeadToHrz:              "changed-4",
					GILessonCsgoInstrExplainBombCarrier:       "changed-5",
					GILessonCsgoInstrExplainBuyarmor:          "changed-6",
					GILessonCsgoInstrExplainBuymenu:           "changed-7",
					GILessonCsgoInstrExplainFollowBomber:      "changed-8",
					GILessonCsgoInstrExplainInspect:           "changed-0",
					GILessonCsgoInstrExplainPickupBomb:        "changed-9",
					GILessonCsgoInstrExplainPlantBomb:         "changed-10",
					GILessonCsgoInstrExplainPreventBombPickup: "changed-11",
					GILessonCsgoInstrExplainReload:            "changed-12",
					GILessonCsgoInstrExplainZoom:              "changed-13",
					GILessonCsgoInstrRescueZone:               "changed-14",
					GILessonDefusePlantedBomb:                 "changed-15",
					GILessonFindPlantedBomb:                   "changed-16",
					GILessonTrExplainPlantBomb:                "changed-17",
					GILessonVersionNumber:                     "changed-18",
					LastMatchContributionScore:                "changed-19",
					LastMatchCtWins:                           "changed-20",
					LastMatchDamage:                           "changed-21",
					LastMatchDeaths:                           "changed-22",
					LastMatchDominations:                      "changed-23",
					LastMatchFavweaponID:                      "changed-25",
					LastMatchFavweaponHits:                    "changed-24",
					LastMatchFavweaponKills:                   "changed-26",
					LastMatchFavweaponShots:                   "changed-27",
					LastMatchGgContributionScore:              "changed-28",
					LastMatchKills:                            "changed-29",
					LastMatchMaxPlayers:                       "changed-30",
					LastMatchMoneySpent:                       "changed-31",
					LastMatchMvps:                             "changed-32",
					LastMatchRevenges:                         "changed-33",
					LastMatchRounds:                           "changed-34",
					LastMatchTWins:                            "changed-35",
					LastMatchWins:                             "changed-36",
					SteamStatMatchwinscomp:                    "changed-37",
					SteamStatSurvivedz:                        "changed-38",
					SteamStatXpearnedgames:                    "changed-39",
					TotalBrokenWindows:                        "changed-40",
					TotalContributionScore:                    "changed-41",
					TotalDamageDone:                           "changed-42",
					TotalDeaths:                               "changed-43",
					TotalDefusedBombs:                         "changed-44",
					TotalDominationOverkills:                  "changed-45",
					TotalDominations:                          "changed-46",
					TotalGgMatchesPlayed:                      "changed-47",
					TotalGgMatchesWon:                         "changed-48",
					TotalGunGameContributionScore:             "changed-49",
					TotalGunGameRoundsPlayed:                  "changed-50",
					TotalGunGameRoundsWon:                     "changed-51",
					TotalHitsAk47:                             "changed-52",
					TotalHitsAug:                              "changed-53",
					TotalHitsAwp:                              "changed-54",
					TotalHitsBizon:                            "changed-55",
					TotalHitsDeagle:                           "changed-56",
					TotalHitsElite:                            "changed-57",
					TotalHitsFamas:                            "changed-58",
					TotalHitsFiveseven:                        "changed-59",
					TotalHitsG3sg1:                            "changed-60",
					TotalHitsGalilar:                          "changed-61",
					TotalHitsGlock:                            "changed-62",
					TotalHitsHkp2000:                          "changed-63",
					TotalHitsM249:                             "changed-64",
					TotalHitsM4a1:                             "changed-65",
					TotalHitsMac10:                            "changed-66",
					TotalHitsMag7:                             "changed-67",
					TotalHitsMp7:                              "changed-68",
					TotalHitsMp9:                              "changed-69",
					TotalHitsNegev:                            "changed-70",
					TotalHitsNova:                             "changed-71",
					TotalHitsP250:                             "changed-72",
					TotalHitsP90:                              "changed-73",
					TotalHitsSg556:                            "changed-76",
					TotalHitsSawedoff:                         "changed-74",
					TotalHitsScar20:                           "changed-75",
					TotalHitsSsg08:                            "changed-77",
					TotalHitsTec9:                             "changed-78",
					TotalHitsUmp45:                            "changed-79",
					TotalHitsXm1014:                           "changed-80",
					TotalKills:                                "changed-81",
					TotalKillsAgainstZoomedSniper:             "changed-82",
					TotalKillsAk47:                            "changed-83",
					TotalKillsAug:                             "changed-84",
					TotalKillsAwp:                             "changed-85",
					TotalKillsBizon:                           "changed-86",
					TotalKillsDeagle:                          "changed-87",
					TotalKillsElite:                           "changed-88",
					TotalKillsEnemyBlinded:                    "changed-89",
					TotalKillsEnemyWeapon:                     "changed-90",
					TotalKillsFamas:                           "changed-91",
					TotalKillsFiveseven:                       "changed-92",
					TotalKillsG3sg1:                           "changed-93",
					TotalKillsGalilar:                         "changed-94",
					TotalKillsGlock:                           "changed-95",
					TotalKillsHeadshot:                        "changed-96",
					TotalKillsHegrenade:                       "changed-97",
					TotalKillsHkp2000:                         "changed-98",
					TotalKillsKnife:                           "changed-99",
					TotalKillsKnifeFight:                      "changed-100",
					TotalKillsM249:                            "changed-101",
					TotalKillsM4a1:                            "changed-102",
					TotalKillsMac10:                           "changed-103",
					TotalKillsMag7:                            "changed-104",
					TotalKillsMolotov:                         "changed-105",
					TotalKillsMp7:                             "changed-106",
					TotalKillsMp9:                             "changed-107",
					TotalKillsNegev:                           "changed-108",
					TotalKillsNova:                            "changed-109",
					TotalKillsP250:                            "changed-110",
					TotalKillsP90:                             "changed-111",
					TotalKillsSawedoff:                        "changed-112",
					TotalKillsScar20:                          "changed-113",
					TotalKillsSg556:                           "changed-114",
					TotalKillsSsg08:                           "changed-115",
					TotalKillsTaser:                           "changed-116",
					TotalKillsTec9:                            "changed-117",
					TotalKillsUmp45:                           "changed-118",
					TotalKillsXm1014:                          "changed-119",
					TotalMatchesPlayed:                        "changed-120",
					TotalMatchesWon:                           "changed-121",
					TotalMatchesWonBaggage:                    "changed-122",
					TotalMatchesWonBank:                       "changed-123",
					TotalMatchesWonLake:                       "changed-124",
					TotalMatchesWonSafehouse:                  "changed-125",
					TotalMatchesWonShoots:                     "changed-126",
					TotalMatchesWonStmarc:                     "changed-127",
					TotalMatchesWonSugarcane:                  "changed-128",
					TotalMatchesWonTrain:                      "changed-129",
					TotalMoneyEarned:                          "changed-130",
					TotalMvps:                                 "changed-131",
					TotalPlantedBombs:                         "changed-132",
					TotalProgressiveMatchesWon:                "changed-133",
					TotalRescuedHostages:                      "changed-134",
					TotalRevenges:                             "changed-135",
					TotalRoundsMapArBaggage:                   "changed-136",
					TotalRoundsMapArMonastery:                 "changed-137",
					TotalRoundsMapArShoots:                    "changed-138",
					TotalRoundsMapCsAssault:                   "changed-139",
					TotalRoundsMapCsItaly:                     "changed-140",
					TotalRoundsMapCsMilitia:                   "changed-141",
					TotalRoundsMapCsOffice:                    "changed-142",
					TotalRoundsMapDeAztec:                     "changed-143",
					TotalRoundsMapDeBank:                      "changed-144",
					TotalRoundsMapDeCbble:                     "changed-145",
					TotalRoundsMapDeDust:                      "changed-146",
					TotalRoundsMapDeDust2:                     "changed-147",
					TotalRoundsMapDeInferno:                   "changed-148",
					TotalRoundsMapDeLake:                      "changed-149",
					TotalRoundsMapDeNuke:                      "changed-150",
					TotalRoundsMapDeSafehouse:                 "changed-151",
					TotalRoundsMapDeShorttrain:                "changed-152",
					TotalRoundsMapDeStmarc:                    "changed-153",
					TotalRoundsMapDeSugarcane:                 "changed-154",
					TotalRoundsMapDeTrain:                     "changed-155",
					TotalRoundsMapDeVertigo:                   "changed-156",
					TotalRoundsPlayed:                         "changed-157",
					TotalShotsAk47:                            "changed-158",
					TotalShotsAug:                             "changed-159",
					TotalShotsAwp:                             "changed-160",
					TotalShotsBizon:                           "changed-161",
					TotalShotsDeagle:                          "changed-162",
					TotalShotsElite:                           "changed-163",
					TotalShotsFamas:                           "changed-164",
					TotalShotsFired:                           "changed-165",
					TotalShotsFiveseven:                       "changed-166",
					TotalShotsG3sg1:                           "changed-167",
					TotalShotsGalilar:                         "changed-168",
					TotalShotsGlock:                           "changed-169",
					TotalShotsHit:                             "changed-170",
					TotalShotsHkp2000:                         "changed-171",
					TotalShotsM249:                            "changed-172",
					TotalShotsM4a1:                            "changed-173",
					TotalShotsMac10:                           "changed-174",
					TotalShotsMag7:                            "changed-175",
					TotalShotsMp7:                             "changed-176",
					TotalShotsMp9:                             "changed-177",
					TotalShotsNegev:                           "changed-178",
					TotalShotsNova:                            "changed-179",
					TotalShotsP250:                            "changed-180",
					TotalShotsP90:                             "changed-181",
					TotalShotsSawedoff:                        "changed-182",
					TotalShotsScar20:                          "changed-183",
					TotalShotsSg556:                           "changed-184",
					TotalShotsSsg08:                           "changed-185",
					TotalShotsTaser:                           "changed-186",
					TotalShotsTec9:                            "changed-187",
					TotalShotsUmp45:                           "changed-188",
					TotalShotsXm1014:                          "changed-189",
					TotalTimePlayed:                           "changed-192",
					TotalTrbombMatchesWon:                     "changed-193",
					TotalTRDefusedBombs:                       "changed-190",
					TotalTRPlantedBombs:                       "changed-191",
					TotalWeaponsDonated:                       "changed-194",
					TotalWins:                                 "changed-195",
					TotalWinsMapArBaggage:                     "changed-196",
					TotalWinsMapArMonastery:                   "changed-197",
					TotalWinsMapArShoots:                      "changed-198",
					TotalWinsMapCsAssault:                     "changed-199",
					TotalWinsMapCsItaly:                       "changed-200",
					TotalWinsMapCsMilitia:                     "changed-201",
					TotalWinsMapCsOffice:                      "changed-202",
					TotalWinsMapDeAztec:                       "changed-203",
					TotalWinsMapDeBank:                        "changed-204",
					TotalWinsMapDeCbble:                       "changed-205",
					TotalWinsMapDeDust:                        "changed-206",
					TotalWinsMapDeDust2:                       "changed-207",
					TotalWinsMapDeHouse:                       "changed-208",
					TotalWinsMapDeInferno:                     "changed-209",
					TotalWinsMapDeLake:                        "changed-210",
					TotalWinsMapDeNuke:                        "changed-211",
					TotalWinsMapDeSafehouse:                   "changed-212",
					TotalWinsMapDeShorttrain:                  "changed-213",
					TotalWinsMapDeStmarc:                      "changed-214",
					TotalWinsMapDeSugarcane:                   "changed-215",
					TotalWinsMapDeTrain:                       "changed-216",
					TotalWinsMapDeVertigo:                     "changed-217",
					TotalWinsPistolround:                      "changed-218",
				},
			},
			want: steamclient.UserStatsForGame{
				SteamID: "all_columns",
				// GameName: "1",
				Extra: steamclient.GameExtras{
					SteamID:      "all_columns",
					TotalKD:      "extra0",
					LastMatchKD:  "extra1",
					HitRatio:     "extra2",
					PlayedHours:  "extra3",
					TotalADR:     "extra4",
					LastMatchADR: "extra5",
				},
				Stats: steamclient.GameStats{
					SteamID:                                   "all_columns",
					GILessonBombSitesA:                        "changed-1",
					GILessonBombSitesB:                        "changed-2",
					GILessonCsgoCycleWeaponsKb:                "changed-3",
					GILessonCsgoHostageLeadToHrz:              "changed-4",
					GILessonCsgoInstrExplainBombCarrier:       "changed-5",
					GILessonCsgoInstrExplainBuyarmor:          "changed-6",
					GILessonCsgoInstrExplainBuymenu:           "changed-7",
					GILessonCsgoInstrExplainFollowBomber:      "changed-8",
					GILessonCsgoInstrExplainInspect:           "changed-0",
					GILessonCsgoInstrExplainPickupBomb:        "changed-9",
					GILessonCsgoInstrExplainPlantBomb:         "changed-10",
					GILessonCsgoInstrExplainPreventBombPickup: "changed-11",
					GILessonCsgoInstrExplainReload:            "changed-12",
					GILessonCsgoInstrExplainZoom:              "changed-13",
					GILessonCsgoInstrRescueZone:               "changed-14",
					GILessonDefusePlantedBomb:                 "changed-15",
					GILessonFindPlantedBomb:                   "changed-16",
					GILessonTrExplainPlantBomb:                "changed-17",
					GILessonVersionNumber:                     "changed-18",
					LastMatchContributionScore:                "changed-19",
					LastMatchCtWins:                           "changed-20",
					LastMatchDamage:                           "changed-21",
					LastMatchDeaths:                           "changed-22",
					LastMatchDominations:                      "changed-23",
					LastMatchFavweaponID:                      "changed-25",
					LastMatchFavweaponHits:                    "changed-24",
					LastMatchFavweaponKills:                   "changed-26",
					LastMatchFavweaponShots:                   "changed-27",
					LastMatchGgContributionScore:              "changed-28",
					LastMatchKills:                            "changed-29",
					LastMatchMaxPlayers:                       "changed-30",
					LastMatchMoneySpent:                       "changed-31",
					LastMatchMvps:                             "changed-32",
					LastMatchRevenges:                         "changed-33",
					LastMatchRounds:                           "changed-34",
					LastMatchTWins:                            "changed-35",
					LastMatchWins:                             "changed-36",
					SteamStatMatchwinscomp:                    "changed-37",
					SteamStatSurvivedz:                        "changed-38",
					SteamStatXpearnedgames:                    "changed-39",
					TotalBrokenWindows:                        "changed-40",
					TotalContributionScore:                    "changed-41",
					TotalDamageDone:                           "changed-42",
					TotalDeaths:                               "changed-43",
					TotalDefusedBombs:                         "changed-44",
					TotalDominationOverkills:                  "changed-45",
					TotalDominations:                          "changed-46",
					TotalGgMatchesPlayed:                      "changed-47",
					TotalGgMatchesWon:                         "changed-48",
					TotalGunGameContributionScore:             "changed-49",
					TotalGunGameRoundsPlayed:                  "changed-50",
					TotalGunGameRoundsWon:                     "changed-51",
					TotalHitsAk47:                             "changed-52",
					TotalHitsAug:                              "changed-53",
					TotalHitsAwp:                              "changed-54",
					TotalHitsBizon:                            "changed-55",
					TotalHitsDeagle:                           "changed-56",
					TotalHitsElite:                            "changed-57",
					TotalHitsFamas:                            "changed-58",
					TotalHitsFiveseven:                        "changed-59",
					TotalHitsG3sg1:                            "changed-60",
					TotalHitsGalilar:                          "changed-61",
					TotalHitsGlock:                            "changed-62",
					TotalHitsHkp2000:                          "changed-63",
					TotalHitsM249:                             "changed-64",
					TotalHitsM4a1:                             "changed-65",
					TotalHitsMac10:                            "changed-66",
					TotalHitsMag7:                             "changed-67",
					TotalHitsMp7:                              "changed-68",
					TotalHitsMp9:                              "changed-69",
					TotalHitsNegev:                            "changed-70",
					TotalHitsNova:                             "changed-71",
					TotalHitsP250:                             "changed-72",
					TotalHitsP90:                              "changed-73",
					TotalHitsSg556:                            "changed-76",
					TotalHitsSawedoff:                         "changed-74",
					TotalHitsScar20:                           "changed-75",
					TotalHitsSsg08:                            "changed-77",
					TotalHitsTec9:                             "changed-78",
					TotalHitsUmp45:                            "changed-79",
					TotalHitsXm1014:                           "changed-80",
					TotalKills:                                "changed-81",
					TotalKillsAgainstZoomedSniper:             "changed-82",
					TotalKillsAk47:                            "changed-83",
					TotalKillsAug:                             "changed-84",
					TotalKillsAwp:                             "changed-85",
					TotalKillsBizon:                           "changed-86",
					TotalKillsDeagle:                          "changed-87",
					TotalKillsElite:                           "changed-88",
					TotalKillsEnemyBlinded:                    "changed-89",
					TotalKillsEnemyWeapon:                     "changed-90",
					TotalKillsFamas:                           "changed-91",
					TotalKillsFiveseven:                       "changed-92",
					TotalKillsG3sg1:                           "changed-93",
					TotalKillsGalilar:                         "changed-94",
					TotalKillsGlock:                           "changed-95",
					TotalKillsHeadshot:                        "changed-96",
					TotalKillsHegrenade:                       "changed-97",
					TotalKillsHkp2000:                         "changed-98",
					TotalKillsKnife:                           "changed-99",
					TotalKillsKnifeFight:                      "changed-100",
					TotalKillsM249:                            "changed-101",
					TotalKillsM4a1:                            "changed-102",
					TotalKillsMac10:                           "changed-103",
					TotalKillsMag7:                            "changed-104",
					TotalKillsMolotov:                         "changed-105",
					TotalKillsMp7:                             "changed-106",
					TotalKillsMp9:                             "changed-107",
					TotalKillsNegev:                           "changed-108",
					TotalKillsNova:                            "changed-109",
					TotalKillsP250:                            "changed-110",
					TotalKillsP90:                             "changed-111",
					TotalKillsSawedoff:                        "changed-112",
					TotalKillsScar20:                          "changed-113",
					TotalKillsSg556:                           "changed-114",
					TotalKillsSsg08:                           "changed-115",
					TotalKillsTaser:                           "changed-116",
					TotalKillsTec9:                            "changed-117",
					TotalKillsUmp45:                           "changed-118",
					TotalKillsXm1014:                          "changed-119",
					TotalMatchesPlayed:                        "changed-120",
					TotalMatchesWon:                           "changed-121",
					TotalMatchesWonBaggage:                    "changed-122",
					TotalMatchesWonBank:                       "changed-123",
					TotalMatchesWonLake:                       "changed-124",
					TotalMatchesWonSafehouse:                  "changed-125",
					TotalMatchesWonShoots:                     "changed-126",
					TotalMatchesWonStmarc:                     "changed-127",
					TotalMatchesWonSugarcane:                  "changed-128",
					TotalMatchesWonTrain:                      "changed-129",
					TotalMoneyEarned:                          "changed-130",
					TotalMvps:                                 "changed-131",
					TotalPlantedBombs:                         "changed-132",
					TotalProgressiveMatchesWon:                "changed-133",
					TotalRescuedHostages:                      "changed-134",
					TotalRevenges:                             "changed-135",
					TotalRoundsMapArBaggage:                   "changed-136",
					TotalRoundsMapArMonastery:                 "changed-137",
					TotalRoundsMapArShoots:                    "changed-138",
					TotalRoundsMapCsAssault:                   "changed-139",
					TotalRoundsMapCsItaly:                     "changed-140",
					TotalRoundsMapCsMilitia:                   "changed-141",
					TotalRoundsMapCsOffice:                    "changed-142",
					TotalRoundsMapDeAztec:                     "changed-143",
					TotalRoundsMapDeBank:                      "changed-144",
					TotalRoundsMapDeCbble:                     "changed-145",
					TotalRoundsMapDeDust:                      "changed-146",
					TotalRoundsMapDeDust2:                     "changed-147",
					TotalRoundsMapDeInferno:                   "changed-148",
					TotalRoundsMapDeLake:                      "changed-149",
					TotalRoundsMapDeNuke:                      "changed-150",
					TotalRoundsMapDeSafehouse:                 "changed-151",
					TotalRoundsMapDeShorttrain:                "changed-152",
					TotalRoundsMapDeStmarc:                    "changed-153",
					TotalRoundsMapDeSugarcane:                 "changed-154",
					TotalRoundsMapDeTrain:                     "changed-155",
					TotalRoundsMapDeVertigo:                   "changed-156",
					TotalRoundsPlayed:                         "changed-157",
					TotalShotsAk47:                            "changed-158",
					TotalShotsAug:                             "changed-159",
					TotalShotsAwp:                             "changed-160",
					TotalShotsBizon:                           "changed-161",
					TotalShotsDeagle:                          "changed-162",
					TotalShotsElite:                           "changed-163",
					TotalShotsFamas:                           "changed-164",
					TotalShotsFired:                           "changed-165",
					TotalShotsFiveseven:                       "changed-166",
					TotalShotsG3sg1:                           "changed-167",
					TotalShotsGalilar:                         "changed-168",
					TotalShotsGlock:                           "changed-169",
					TotalShotsHit:                             "changed-170",
					TotalShotsHkp2000:                         "changed-171",
					TotalShotsM249:                            "changed-172",
					TotalShotsM4a1:                            "changed-173",
					TotalShotsMac10:                           "changed-174",
					TotalShotsMag7:                            "changed-175",
					TotalShotsMp7:                             "changed-176",
					TotalShotsMp9:                             "changed-177",
					TotalShotsNegev:                           "changed-178",
					TotalShotsNova:                            "changed-179",
					TotalShotsP250:                            "changed-180",
					TotalShotsP90:                             "changed-181",
					TotalShotsSawedoff:                        "changed-182",
					TotalShotsScar20:                          "changed-183",
					TotalShotsSg556:                           "changed-184",
					TotalShotsSsg08:                           "changed-185",
					TotalShotsTaser:                           "changed-186",
					TotalShotsTec9:                            "changed-187",
					TotalShotsUmp45:                           "changed-188",
					TotalShotsXm1014:                          "changed-189",
					TotalTimePlayed:                           "changed-192",
					TotalTrbombMatchesWon:                     "changed-193",
					TotalTRDefusedBombs:                       "changed-190",
					TotalTRPlantedBombs:                       "changed-191",
					TotalWeaponsDonated:                       "changed-194",
					TotalWins:                                 "changed-195",
					TotalWinsMapArBaggage:                     "changed-196",
					TotalWinsMapArMonastery:                   "changed-197",
					TotalWinsMapArShoots:                      "changed-198",
					TotalWinsMapCsAssault:                     "changed-199",
					TotalWinsMapCsItaly:                       "changed-200",
					TotalWinsMapCsMilitia:                     "changed-201",
					TotalWinsMapCsOffice:                      "changed-202",
					TotalWinsMapDeAztec:                       "changed-203",
					TotalWinsMapDeBank:                        "changed-204",
					TotalWinsMapDeCbble:                       "changed-205",
					TotalWinsMapDeDust:                        "changed-206",
					TotalWinsMapDeDust2:                       "changed-207",
					TotalWinsMapDeHouse:                       "changed-208",
					TotalWinsMapDeInferno:                     "changed-209",
					TotalWinsMapDeLake:                        "changed-210",
					TotalWinsMapDeNuke:                        "changed-211",
					TotalWinsMapDeSafehouse:                   "changed-212",
					TotalWinsMapDeShorttrain:                  "changed-213",
					TotalWinsMapDeStmarc:                      "changed-214",
					TotalWinsMapDeSugarcane:                   "changed-215",
					TotalWinsMapDeTrain:                       "changed-216",
					TotalWinsMapDeVertigo:                     "changed-217",
					TotalWinsPistolround:                      "changed-218",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()

			if err := db.UpdateUserStatsForGame(tt.stats); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdateUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got, err := db.GetPlayerInfoBySteamID(tt.stats.SteamID); err == nil {
				if diff := cmp.Diff(tt.want, got.UserStatsForGame); diff != "" {
					t.Errorf("DataStorage.UpdateUserStatsForGame() mismatch (-want +got):\n%s", diff)
				}
			} else {
				if !(tt.wantErr == (err != nil)) {
					t.Errorf("DataStorage.UpdateUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			// if !reflect.DeepEqual(tt.stats, tt.want) {
			// 	t.Errorf("DataStorage.UpdateStatsForgame() = %v, want %v", tt.stats, tt.want)
			// }
		})
	}
}
