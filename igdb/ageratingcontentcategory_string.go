// Code generated by "stringer -type=AgeRatingContentCategory"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AgeRatingContentESRB_AlcoholReference-1]
	_ = x[AgeRatingContentESRB_AnimetedBlood-2]
	_ = x[AgeRatingContentESRB_Blood-3]
	_ = x[AgeRatingContentESRB_BloodAndGore-4]
	_ = x[AgeRatingContentESRB_CartoonViolence-5]
	_ = x[AgeRatingContentESRB_ComicMischief-6]
	_ = x[AgeRatingContentESRB_CrudeHumor-7]
	_ = x[AgeRatingContentESRB_DrugReference-8]
	_ = x[AgeRatingContentESRB_FantasyViolence-9]
	_ = x[AgeRatingContentESRB_IntenseViolence-10]
	_ = x[AgeRatingContentESRB_Language-11]
	_ = x[AgeRatingContentESRB_Lyrics-12]
	_ = x[AgeRatingContentESRB_MatureHumor-13]
	_ = x[AgeRatingContentESRB_Nudity-14]
	_ = x[AgeRatingContentESRB_PartialNudity-15]
	_ = x[AgeRatingContentESRB_RealGambling-16]
	_ = x[AgeRatingContentESRB_SexualContent-17]
	_ = x[AgeRatingContentESRB_SexualThemes-18]
	_ = x[AgeRatingContentESRB_SexualViolence-19]
	_ = x[AgeRatingContentESRB_SimulatedGambling-20]
	_ = x[AgeRatingContentESRB_StrongLanguage-21]
	_ = x[AgeRatingContentESRB_StrongLyrics-22]
	_ = x[AgeRatingContentESRB_StrongSexualContent-23]
	_ = x[AgeRatingContentESRB_SuggestiveThemes-24]
	_ = x[AgeRatingContentESRB_TobaccoReference-25]
	_ = x[AgeRatingContentESRB_UseOfAlcohol-26]
	_ = x[AgeRatingContentESRB_UseOfDrugs-27]
	_ = x[AgeRatingContentESRB_UseOfTobacco-28]
	_ = x[AgeRatingContentESRB_Violence-29]
	_ = x[AgeRatingContentESRB_ViolenceReference-30]
	_ = x[AgeRatingContentESRB_AnimatedViolence-31]
	_ = x[AgeRatingContentESRB_MildLanguage-32]
	_ = x[AgeRatingContentESRB_MildViolence-33]
	_ = x[AgeRatingContentESRB_UseOfDrugsAndAlcohol-34]
	_ = x[AgeRatingContentESRB_DrugAndAlcoholReference-35]
	_ = x[AgeRatingContentESRB_MildSuggestiveTheme-36]
	_ = x[AgeRatingContentESRB_MildCartoonViolence-37]
	_ = x[AgeRatingContentESRB_MildBlood-38]
	_ = x[AgeRatingContentESRB_RealisticBloodAndGore-39]
	_ = x[AgeRatingContentESRB_RealisticViolence-40]
	_ = x[AgeRatingContentESRB_AlcoholAndTobaccoReference-41]
	_ = x[AgeRatingContentESRB_MatureSexualTheme-42]
	_ = x[AgeRatingContentESRB_MildAnimatedViolence-43]
	_ = x[AgeRatingContentESRB_MildSexualTheme-44]
	_ = x[AgeRatingContentESRB_UseOfAlcoholAndTobacco-45]
	_ = x[AgeRatingContentESRB_AnimatedBloodAndGore-46]
	_ = x[AgeRatingContentESRB_MildFantasyViolence-47]
	_ = x[AgeRatingContentESRB_MildLyrics-48]
	_ = x[AgeRatingContentESRB_RealisticBlood-49]
	_ = x[AgeRatingContentPEGI_Violence-50]
	_ = x[AgeRatingContentPEGI_Discrimination-51]
	_ = x[AgeRatingContentPEGI_BadLanguage-52]
	_ = x[AgeRatingContentPEGI_Gambling-53]
	_ = x[AgeRatingContentPEGI_OnlineGameplay-54]
	_ = x[AgeRatingContentPEGI_InGamePurchases-55]
	_ = x[AgeRatingContentCERO_Love-56]
	_ = x[AgeRatingContentCERO_SexualContent-57]
	_ = x[AgeRatingContentCERO_Violence-58]
	_ = x[AgeRatingContentCERO_Horror-59]
	_ = x[AgeRatingContentCERO_DrinkingSmoking-60]
	_ = x[AgeRatingContentCERO_Gambling-61]
	_ = x[AgeRatingContentCERO_Crime-62]
	_ = x[AgeRatingContentCERO_ControlledSubstances-63]
	_ = x[AgeRatingContentCERO_LanguageAndOthers-64]
	_ = x[AgeRatingContentGRAC_Sexuality-65]
	_ = x[AgeRatingContentGRAC_Violence-66]
	_ = x[AgeRatingContentGRAC_FearHorrorThreatening-67]
	_ = x[AgeRatingContentGRAC_Languange-68]
	_ = x[AgeRatingContentGRAC_AlcoholTobaccoDrug-69]
	_ = x[AgeRatingContentGRAC_CrimeAntiSocial-70]
	_ = x[AgeRatingContentGRAC_Gambling-71]
	_ = x[AgeRatingContentCLASS_IND_Violencia-72]
	_ = x[AgeRatingContentCLASS_IND_ViolenciaExtrema-73]
	_ = x[AgeRatingContentCLASS_IND_ConteudoSexual-74]
	_ = x[AgeRatingContentCLASS_IND_Nudez-75]
	_ = x[AgeRatingContentCLASS_IND_Sexo-76]
	_ = x[AgeRatingContentCLASS_IND_SexoExplicito-77]
	_ = x[AgeRatingContentCLASS_IND_Drogas-78]
	_ = x[AgeRatingContentCLASS_IND_DrogasLicitas-79]
	_ = x[AgeRatingContentCLASS_IND_DrogasIlícitas-80]
	_ = x[AgeRatingContentCLASS_IND_LinguagemImpropria-81]
	_ = x[AgeRatingContentCLASS_IND_AtosCriminosos-82]
}

const _AgeRatingContentCategory_name = "AgeRatingContentESRB_AlcoholReferenceAgeRatingContentESRB_AnimetedBloodAgeRatingContentESRB_BloodAgeRatingContentESRB_BloodAndGoreAgeRatingContentESRB_CartoonViolenceAgeRatingContentESRB_ComicMischiefAgeRatingContentESRB_CrudeHumorAgeRatingContentESRB_DrugReferenceAgeRatingContentESRB_FantasyViolenceAgeRatingContentESRB_IntenseViolenceAgeRatingContentESRB_LanguageAgeRatingContentESRB_LyricsAgeRatingContentESRB_MatureHumorAgeRatingContentESRB_NudityAgeRatingContentESRB_PartialNudityAgeRatingContentESRB_RealGamblingAgeRatingContentESRB_SexualContentAgeRatingContentESRB_SexualThemesAgeRatingContentESRB_SexualViolenceAgeRatingContentESRB_SimulatedGamblingAgeRatingContentESRB_StrongLanguageAgeRatingContentESRB_StrongLyricsAgeRatingContentESRB_StrongSexualContentAgeRatingContentESRB_SuggestiveThemesAgeRatingContentESRB_TobaccoReferenceAgeRatingContentESRB_UseOfAlcoholAgeRatingContentESRB_UseOfDrugsAgeRatingContentESRB_UseOfTobaccoAgeRatingContentESRB_ViolenceAgeRatingContentESRB_ViolenceReferenceAgeRatingContentESRB_AnimatedViolenceAgeRatingContentESRB_MildLanguageAgeRatingContentESRB_MildViolenceAgeRatingContentESRB_UseOfDrugsAndAlcoholAgeRatingContentESRB_DrugAndAlcoholReferenceAgeRatingContentESRB_MildSuggestiveThemeAgeRatingContentESRB_MildCartoonViolenceAgeRatingContentESRB_MildBloodAgeRatingContentESRB_RealisticBloodAndGoreAgeRatingContentESRB_RealisticViolenceAgeRatingContentESRB_AlcoholAndTobaccoReferenceAgeRatingContentESRB_MatureSexualThemeAgeRatingContentESRB_MildAnimatedViolenceAgeRatingContentESRB_MildSexualThemeAgeRatingContentESRB_UseOfAlcoholAndTobaccoAgeRatingContentESRB_AnimatedBloodAndGoreAgeRatingContentESRB_MildFantasyViolenceAgeRatingContentESRB_MildLyricsAgeRatingContentESRB_RealisticBloodAgeRatingContentPEGI_ViolenceAgeRatingContentPEGI_DiscriminationAgeRatingContentPEGI_BadLanguageAgeRatingContentPEGI_GamblingAgeRatingContentPEGI_OnlineGameplayAgeRatingContentPEGI_InGamePurchasesAgeRatingContentCERO_LoveAgeRatingContentCERO_SexualContentAgeRatingContentCERO_ViolenceAgeRatingContentCERO_HorrorAgeRatingContentCERO_DrinkingSmokingAgeRatingContentCERO_GamblingAgeRatingContentCERO_CrimeAgeRatingContentCERO_ControlledSubstancesAgeRatingContentCERO_LanguageAndOthersAgeRatingContentGRAC_SexualityAgeRatingContentGRAC_ViolenceAgeRatingContentGRAC_FearHorrorThreateningAgeRatingContentGRAC_LanguangeAgeRatingContentGRAC_AlcoholTobaccoDrugAgeRatingContentGRAC_CrimeAntiSocialAgeRatingContentGRAC_GamblingAgeRatingContentCLASS_IND_ViolenciaAgeRatingContentCLASS_IND_ViolenciaExtremaAgeRatingContentCLASS_IND_ConteudoSexualAgeRatingContentCLASS_IND_NudezAgeRatingContentCLASS_IND_SexoAgeRatingContentCLASS_IND_SexoExplicitoAgeRatingContentCLASS_IND_DrogasAgeRatingContentCLASS_IND_DrogasLicitasAgeRatingContentCLASS_IND_DrogasIlícitasAgeRatingContentCLASS_IND_LinguagemImpropriaAgeRatingContentCLASS_IND_AtosCriminosos"

var _AgeRatingContentCategory_index = [...]uint16{0, 37, 71, 97, 130, 166, 200, 231, 265, 301, 337, 366, 393, 425, 452, 486, 519, 553, 586, 621, 659, 694, 727, 767, 804, 841, 874, 905, 938, 967, 1005, 1042, 1075, 1108, 1149, 1193, 1233, 1273, 1303, 1345, 1383, 1430, 1468, 1509, 1545, 1588, 1629, 1669, 1700, 1735, 1764, 1799, 1831, 1860, 1895, 1931, 1956, 1990, 2019, 2046, 2082, 2111, 2137, 2178, 2216, 2246, 2275, 2317, 2347, 2386, 2422, 2451, 2486, 2528, 2568, 2599, 2629, 2668, 2700, 2739, 2780, 2824, 2864}

func (i AgeRatingContentCategory) String() string {
	i -= 1
	if i < 0 || i >= AgeRatingContentCategory(len(_AgeRatingContentCategory_index)-1) {
		return "AgeRatingContentCategory(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AgeRatingContentCategory_name[_AgeRatingContentCategory_index[i]:_AgeRatingContentCategory_index[i+1]]
}
