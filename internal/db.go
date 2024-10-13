package internal

var Characters = map[string]map[string]interface{}{
	"ENDO":   {"name": "円堂", "specialTechniques": []string{"GOD_HAND", "BAKURETSU_PUNCH"}, "tp": 100},
	"GOENJI": {"name": "豪炎寺", "specialTechniques": []string{"FIRE_TORNADO"}, "tp": 90},
	"GENDA":  {"name": "源田", "specialTechniques": []string{"POWER_SHIELD"}, "tp": 80},
	"SAKUMA": {"name": "佐久間", "specialTechniques": []string{"TWIN_BOOST"}, "tp": 85},
}

var CharacterIds = []string{"ENDO", "GOENJI", "GENDA", "SAKUMA"}

var Techniques = map[string]map[string]interface{}{
	"GOD_HAND": {
		"attr":  "DEFENCE",
		"power": 30,
		"cost":  15,
		"name":  "ゴッドハンド",
	},
	"BAKURETSU_PUNCH": {
		"attr":  "DEFENCE",
		"power": 25,
		"cost":  10,
		"name":  "ばくれつパンチ",
	},
	"FIRE_TORNADO": {
		"attr":  "OFFENCE",
		"power": 40,
		"cost":  35,
		"name":  "ファイヤートルネード",
	},
	"POWER_SHIELD": {
		"attr":  "DEFENCE",
		"power": 20,
		"cost":  20,
		"name":  "パワーシールド",
	},
	"TWIN_BOOST": {
		"attr":  "OFFENCE",
		"power": 35,
		"cost":  30,
		"name":  "ツインブースト",
	},
}

func GetPlayerTeamCharacterIds() []string {
	return []string{"ENDO", "GOENJI"}
}

func GetOpponentTeamCharacterIds() []string {
	return []string{"GENDA", "SAKUMA"}
}

func GetCharacter(id string) map[string]interface{} {
	return Characters[id]
}

func GetSpecialTechniques(id string) map[string]interface{} {
	return Techniques[id]
}
