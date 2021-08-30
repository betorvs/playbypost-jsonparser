package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

//Monster struct
type Monster struct {
	Name                  string `json:"name"`
	Meta                  string `json:"meta"`
	ArmorClass            string `json:"Armor Class"`
	HitPoints             string `json:"Hit Points"`
	Speed                 string `json:"Speed"`
	STR                   string `json:"STR"`
	STRMod                string `json:"STR_mod"`
	DEX                   string `json:"DEX"`
	DEXMod                string `json:"DEX_mod"`
	CON                   string `json:"CON"`
	CONMod                string `json:"CON_mod"`
	INT                   string `json:"INT"`
	INTMod                string `json:"INT_mod"`
	WIS                   string `json:"WIS"`
	WISMod                string `json:"WIS_mod"`
	CHA                   string `json:"CHA"`
	CHAMod                string `json:"CHA_mod"`
	DamageVulnerabilities string `json:"Damage Vulnerabilities"`
	DamageImmunities      string `json:"Damage Immunities"`
	ConditionImmunities   string `json:"Condition Immunities"`
	DamageResistances     string `json:"Damage Resistances"`
	SavingThrows          string `json:"Saving Throws"`
	Skills                string `json:"Skills"`
	Senses                string `json:"Senses"`
	Languages             string `json:"Languages"`
	Challenge             string `json:"Challenge"`
	Traits                string `json:"Traits"`
	Actions               string `json:"Actions"`
	LegendaryActions      string `json:"Legendary Actions"`
	ImgURL                string `json:"img_url"`
}

//ReturnMonsterNPC struct
type ReturnMonsterNPC struct {
	Name                   string                 `json:"name"`
	Size                   string                 `json:"size"`
	Type                   string                 `json:"type"`
	Aligment               string                 `json:"aligment"`
	Senses                 string                 `json:"senses"`
	Darkvision             string                 `json:"darkvision"`
	Blindsight             string                 `json:"blindsight"`
	Tremorsense            string                 `json:"tremorsense"`
	Truesight              string                 `json:"truesight"`
	Languages              []string               `json:"languages"`
	Challenge              float64                `json:"challenge"`
	ArmorClass             int                    `json:"armor_class"`
	HitPoints              int                    `json:"hit_points"`
	XP                     int                    `json:"xp"`
	Actions                []string               `json:"actions"`
	WeaponAttack           []WeaponAttack         `json:"weapon_attack"`
	SpecialAttack          []SpecialAttack        `json:"special_attack"`
	SpellCastAbility       SpellCastAbility       `json:"spellcast_abilty,omitempty"`
	InnateSpellCastAbility InnateSpellCastAbility `json:"innate_spellcast_abilty,omitempty"`
	Ability                map[string]int         `json:"ability"`
	Savings                map[string]int         `json:"savings"`
	Skills                 map[string]int         `json:"skills"`
	DamageVulnerabilities  []string               `json:"damage_vulnerabilities"`
	DamageImmunities       []string               `json:"damage_immunities"`
	ConditionImmunities    []string               `json:"condition_immunities"`
	DamageResistances      []string               `json:"damage_resistances"`
	PassivePerception      int                    `json:"passive_perception"`
	Traits                 []string               `json:"traits"`
	LegendaryActions       []string               `json:"legendary_actions"`
	ImgURL                 string                 `json:"img_url"`
}

// WeaponAttack struct
type WeaponAttack struct {
	Name            string `json:"name"`
	Attack          int    `json:"attack"`
	AverageDamage   int    `json:"average_damage"`
	Damage          string `json:"damage"`
	DamageType      string `json:"damage_type"`
	SavingThrows    string `json:"saving_throws,omitempty"`
	DifficultClass  int    `json:"difficult_class,omitempty"`
	ExtraDamage     string `json:"extra_damage,omitempty"`
	ExtraDamageType string `json:"extra_damage_type,omitempty"`
}

// SpecialAttack struct
type SpecialAttack struct {
	Name           string `json:"name"`
	SavingThrows   string `json:"saving_throws"`
	DifficultClass int    `json:"difficult_class"`
	AverageDamage  int    `json:"average_damage,omitempty"`
	Damage         string `json:"damage,omitempty"`
	DamageType     string `json:"damage_type,omitempty"`
	Content        string `json:"content,omitempty"`
}

//SpellCastAbility struct
type SpellCastAbility struct {
	Level          int                 `json:"level,omitempty"`
	DifficultClass int                 `json:"difficult_class,omitempty"`
	Attack         int                 `json:"attack,omitempty"`
	Ability        string              `json:"ability,omitempty"`
	List           string              `json:"list,omitempty"`
	CantripsList   []string            `json:"cantrips_list,omitempty"`
	SlotsPerLevel  map[string]int      `json:"slots_per_level,omitempty"`
	ListPerLevel   map[string][]string `json:"list_per_level,omitempty"`
}

//InnateSpellCastAbility struct
type InnateSpellCastAbility struct {
	DifficultClass  int      `json:"difficult_class,omitempty"`
	Ability         string   `json:"ability,omitempty"`
	List            string   `json:"list,omitempty"`
	AtWillList      []string `json:"at_will_list,omitempty"`
	OnePerDayList   []string `json:"one_per_day_list,omitempty"`
	TwoPerDayList   []string `json:"two_per_day_list,omitempty"`
	ThreePerDayList []string `json:"three_per_day_list,omitempty"`
}

// MonsterPrint func
func MonsterPrint() {
	DB, err := os.Open("./srd_5e_monsters.json")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	var monsters []Monster

	err = json.Unmarshal(byteValue, &monsters)
	if err != nil {
		fmt.Println(err)
	}
	var result []ReturnMonsterNPC
	for _, v := range monsters {
		var partialResult ReturnMonsterNPC
		partialResult.Name = strings.ToLower(strings.ReplaceAll(v.Name, " ", "-"))
		if partialResult.Name == "deep-gnome-(svirfneblin)" {
			partialResult.Name = "deep-gnome-svirfneblin"
		}
		meta := strings.Split(v.Meta, ",")
		sizeType := strings.Split(meta[0], " ")
		partialResult.Size = strings.ToLower(sizeType[0])
		partialResult.Type = sizeType[1]
		partialResult.Aligment = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(meta[1]), " ", "-"))
		partialResult.Senses = v.Senses
		partialResult.Languages = []string{}
		if v.Languages != "" && v.Languages != "--" {
			lang := strings.Split(v.Languages, ",")
			for _, l := range lang {
				partialResult.Languages = append(partialResult.Languages, strings.ToLower(strings.TrimSpace(l)))
			}
		}
		challenges := strings.Split(v.Challenge, "(")
		challenge := strings.TrimSpace(challenges[0])
		if challenge != "1/8" && challenge != "1/4" && challenge != "1/2" {
			partialResult.Challenge = float64(ExtractWholeInt(challenge))
		}
		if challenge == "1/8" {
			partialResult.Challenge = 0.12
		}
		if challenge == "1/4" {
			partialResult.Challenge = 0.25
		}
		if challenge == "1/2" {
			partialResult.Challenge = 0.5
		}
		armor := strings.Split(v.ArmorClass, "(")
		armorClass := ExtractWholeInt(armor[0])
		hit := strings.Split(v.HitPoints, "(")
		hitpoints := ExtractWholeInt(hit[0])
		data := strings.Split(v.Challenge, " ")
		xp := ExtractWholeInt(data[1])
		partialResult.ArmorClass = armorClass
		partialResult.HitPoints = hitpoints
		partialResult.XP = xp
		partialResult.ConditionImmunities = []string{}
		partialResult.DamageImmunities = []string{}
		partialResult.DamageVulnerabilities = []string{}
		partialResult.DamageResistances = []string{}
		partialResult.LegendaryActions = []string{}
		condImu := strings.Split(v.ConditionImmunities, ",")
		for _, v := range condImu {
			if v != "" {
				partialResult.ConditionImmunities = append(partialResult.ConditionImmunities, strings.TrimSpace(strings.ToLower(v)))
			}
		}
		damaImu := strings.Split(v.DamageImmunities, ",")
		for _, v := range damaImu {
			if v != "" && !strings.Contains(v, "Slashing from Nonmagical") {
				partialResult.DamageImmunities = append(partialResult.DamageImmunities, strings.TrimSpace(strings.ToLower(v)))
			}
			if strings.Contains(v, "Slashing from Nonmagical") {
				partialResult.DamageImmunities = append(partialResult.DamageImmunities, "slashing")
			}
		}
		damaVul := strings.Split(v.DamageVulnerabilities, ",")
		for _, v := range damaVul {
			if v != "" {
				partialResult.DamageVulnerabilities = append(partialResult.DamageVulnerabilities, strings.TrimSpace(strings.ToLower(v)))
			}
		}
		damaRes := strings.Split(v.DamageResistances, ",")
		for _, v := range damaRes {
			if v != "" && !strings.Contains(v, "Slashing from Nonmagical") {
				partialResult.DamageResistances = append(partialResult.DamageResistances, strings.TrimSpace(strings.ToLower(v)))
			}
			if strings.Contains(v, "Slashing from Nonmagical") {
				partialResult.DamageResistances = append(partialResult.DamageResistances, "slashing")
			}
		}
		skill := make(map[string]int)
		if v.Skills != "" {
			ss := strings.Split(v.Skills, ",")
			for _, v := range ss {
				sss := strings.Split(v, "+")
				skillName := strings.TrimSpace(strings.ToLower(sss[0]))
				skillValueTemp := sss[1]
				skillValue, err := strconv.Atoi(strings.TrimSpace(skillValueTemp))
				if err != nil {
					fmt.Printf("error convert hitpoints string to int %v", err)
				}
				skill[skillName] = skillValue
			}
		}
		var spellCastAbility SpellCastAbility
		var innateSpellCastAbility InnateSpellCastAbility
		specialAttacks := make([]SpecialAttack, 0)
		partialResult.Skills = skill
		traits := parseAction(v.Traits)
		traitSlice := strings.Split(traits, ";")
		for _, v := range traitSlice {
			if v != "" {
				partialResult.Traits = append(partialResult.Traits, v)
				if strings.Contains(v, "Spellcasting") && !strings.Contains(v, "Innate") {
					spellcast := parseSpellcast(v)
					spellCastAbility = spellcast
				}
				if strings.Contains(v, "Innate Spellcasting") {
					spellcast := parseInnateSpellcast(v)
					innateSpellCastAbility = spellcast
				}
				if strings.Contains(v, "Death Burst") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Death Throes") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Fire Aura") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
			}
		}
		legAct := parseAction(v.LegendaryActions)
		legActSlice := strings.Split(legAct, ";")
		for _, v := range legActSlice {
			if v != "" {
				partialResult.LegendaryActions = append(partialResult.LegendaryActions, v)
			}
		}

		partialResult.ImgURL = v.ImgURL

		weapons := make([]WeaponAttack, 0)

		act := parseAction(v.Actions)
		actions := make([]string, 0)
		actSlice := strings.Split(act, ";")
		for _, v := range actSlice {
			if v != "" {
				actions = append(actions, v)
				if strings.Contains(v, "Weapon Attack") {
					weapon := parseAttack(strings.Split(v, ":"))
					weapons = append(weapons, weapon)
				}
				if strings.Contains(v, "Frightful Presence.") && !strings.Contains(v, "Multiattack.") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Breath") && !strings.Contains(v, "Breath Weapons") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Acid Spray") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Melee Spell Attack") {
					weapon := parseAttack(strings.Split(v, ":"))
					weapons = append(weapons, weapon)
				}
				if strings.Contains(v, "Horrifying") || strings.Contains(v, "Wail") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Tentacles") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Fetid Cloud") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Fey Charm") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Fire Ray") && !strings.Contains(v, "Multiattack.") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Engulf") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Death Glare") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Eye Rays") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Draining Kiss") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
				if strings.Contains(v, "Chilling Gaze") && !strings.Contains(v, "Multiattack.") {
					special := parseSpecialAttack(v)
					specialAttacks = append(specialAttacks, special)
				}
			}
		}
		ability := parseAbility(v)
		savings := parseSavings(v)
		if strings.Contains(v.Senses, "Passive Perception") {
			pas := strings.Split(v.Senses, "Perception")
			passive, err := strconv.Atoi(strings.TrimSpace(pas[1]))
			if err != nil {
				fmt.Printf("error convert xp string to int %v", err)
			}
			partialResult.PassivePerception = passive
		}
		if strings.Contains(v.Senses, "Darkvision") {
			dark := strings.Split(v.Senses, "Darkvision")
			darkvision := strings.Split(dark[1], ",")
			partialResult.Darkvision = strings.TrimSpace(darkvision[0])
		}
		if strings.Contains(v.Senses, "Blindsight") {
			blind := strings.Split(v.Senses, "Blindsight")
			blindsight := strings.Split(blind[1], ",")
			partialResult.Blindsight = strings.TrimSpace(blindsight[0])
		}
		if strings.Contains(v.Senses, "Tremorsense") {
			tremor := strings.Split(v.Senses, "Tremorsense")
			tremorsense := strings.Split(tremor[1], ",")
			partialResult.Tremorsense = strings.TrimSpace(tremorsense[0])
		}
		if strings.Contains(v.Senses, "Truesight") {
			trues := strings.Split(v.Senses, "Truesight")
			truesight := strings.Split(trues[1], ",")
			partialResult.Truesight = strings.TrimSpace(truesight[0])
		}
		partialResult.Actions = actions
		partialResult.WeaponAttack = weapons
		partialResult.SpecialAttack = specialAttacks
		partialResult.Savings = savings
		partialResult.Ability = ability
		partialResult.SpellCastAbility = spellCastAbility
		partialResult.InnateSpellCastAbility = innateSpellCastAbility

		result = append(result, partialResult)
	}
	// var verifiedMonsters []ReturnMonsterNPC
	// for _, item := range result {

	// 	if item.Name == "cockatrice" {
	// 		monster := new(ReturnMonsterNPC)
	// 		monster = &item
	// 		var specialAttack SpecialAttack
	// 		specialAttack.Name = "bite-effect"
	// 		specialAttack.DifficultClass = 11
	// 		specialAttack.SavingThrows = "constitution"
	// 		specialAttack.DamageType = "petrified"
	// 		monster.SpecialAttack = append(monster.SpecialAttack, specialAttack)

	// 	}

	// 	verifiedMonsters = append(verifiedMonsters, item)
	// }
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}

func parseAction(s string) string {
	var result string
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "strong" {
			result += ";"
		}
		if n.Type == 1 {
			result += fmt.Sprintf("%s ", n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return result
}

func parseAttack(v []string) WeaponAttack {
	var weapon WeaponAttack
	// fmt.Println(v[2])
	name := strings.Split(v[0], ".")
	att := strings.Split(v[1], ",")
	hitString := strings.Split(v[2], "(")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		fmt.Println(err)
	}

	atta := reg.ReplaceAllString(att[0], "")
	var attack int
	if atta != "" {
		attack, err = strconv.Atoi(strings.TrimSpace(atta))
		if err != nil {
			fmt.Printf("error convert string to int %v for %s", err, name[0])
		}
	}

	hitClean := reg.ReplaceAllString(hitString[0], "")
	var hit int
	if hitClean != "" {
		hit, err = strconv.Atoi(strings.TrimSpace(hitClean))
		if err != nil {
			fmt.Printf("error convert string to int %v for %s", err, name[0])
		}
	}
	var diceRegex = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
	if hit != 1 {
		var re1 = regexp.MustCompile(`(?m)\(.*?\)`)
		dices := re1.FindString(v[2])
		// fmt.Println(strings.ReplaceAll(dices, " ", ""))
		damage := diceRegex.FindString(strings.ReplaceAll(dices, " ", ""))
		// fmt.Println(damage)
		weapon.Damage = damage
	}
	var damageType string
	var damageTypeRegex = regexp.MustCompile(`(?m)(acid|cold|fire|force|lightning|necrotic|poison|psychic|radiant|thunder|bludgeoning|piercing|slashing)`)
	damageType = damageTypeRegex.FindString(strings.ToLower(v[2]))
	var plusDamageRegex = regexp.MustCompile(`(?m)(plus [0-9] \((\d+)?d(\d+)([\+\-]\d+)?\) *.+? damage)`)
	extraDamageString := plusDamageRegex.FindString(strings.ToLower(v[2]))
	extraDamage := diceRegex.FindString(strings.ReplaceAll(extraDamageString, " ", ""))
	extraDamateType := damageTypeRegex.FindString(strings.ToLower(extraDamageString))

	var difficultClass int
	var savingThrows string
	var dc = regexp.MustCompile(`(?m)(a DC [0-9]+ *.+? saving throw)`)
	var savingRegex = regexp.MustCompile(`(?m)(strength|dexterity|constitution|intelligence|wisdom|charisma)`)
	tempDC := dc.FindString(v[2])
	if tempDC != "" {
		difficultClass = ExtractWholeInt(tempDC)
		savingThrows = savingRegex.FindString(strings.ToLower(tempDC))
	}
	var conditionRegex = regexp.MustCompile(`(?m)(blinded|charmed|deafened|exhaustion|frightened|grappled|incapacitated|invisible|paralyzed|petrified|poisoned|prone|pestrained|stunned|unconscious)`)
	if strings.Contains(v[2], "DC") {
		tempValue := strings.Split(v[2], "DC")
		if len(tempValue) != 0 {
			tempCondition := conditionRegex.FindString(strings.ToLower(tempValue[1]))
			if tempCondition != "" {
				extraDamateType = strings.ToLower(tempCondition)
			}
		}
	}

	weapon.Name = strings.ToLower(strings.ReplaceAll(name[0], " ", "-"))
	switch weapon.Name {
	case "light-crossbow":
		weapon.Name = "crossbow light"
	case "hand-crossbow", "hand-crossbow-(humanoid-or-hybrid-form-only)":
		weapon.Name = "crossbow hand"
	case "heavy-crossbow":
		weapon.Name = "crossbow heavy"
	}
	if hit > 25 {
		hit = 0
	}
	if weapon.Name == "web-(recharge-5–6)" {
		extraDamateType = ""
		damageType = ""
	}
	weapon.AverageDamage = hit
	weapon.Attack = attack
	weapon.DamageType = damageType
	weapon.ExtraDamage = extraDamage
	if weapon.Name == "web-(recharge-5–6)" {
		extraDamateType = ""
	}
	weapon.ExtraDamageType = extraDamateType
	weapon.DifficultClass = difficultClass
	weapon.SavingThrows = savingThrows
	return weapon
}
func parseSpecialAttack(value string) SpecialAttack {
	var special SpecialAttack
	var dc = regexp.MustCompile(`(?m)(a DC [0-9]+ *.+? saving throw)`)
	var savingRegex = regexp.MustCompile(`(?m)(strength|dexterity|constitution|intelligence|wisdom|charisma)`)
	var hasSpecialDamage = regexp.MustCompile(`(?m)((\d+)?d(\d+)([\+\-]\d+)? *.+? damage)`)
	var dices = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
	var damageTypeRegex = regexp.MustCompile(`(?m)(acid|cold|fire|force|lightning|necrotic|poison|psychic|radiant|thunder|bludgeoning|piercing|slashing)`)
	var conditionRegex = regexp.MustCompile(`(?m)(blinded|charmed|deafened|exhaustion|frightened|grappled|incapacitated|invisible|paralyzed|petrified|poisoned|prone|pestrained|stunned|unconscious)`)
	tempName := strings.Split(value, ".")
	special.Name = strings.ToLower(strings.ReplaceAll(tempName[0], " ", "-"))
	if strings.Contains(special.Name, "(") {
		cleanName := strings.Split(tempName[0], "(")
		special.Name = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(cleanName[0]), " ", "-"))
	}
	if strings.Contains(special.Name, ":") {
		cleanName := strings.Split(special.Name, ":")
		special.Name = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(cleanName[0]), " ", "-"))
	}
	// special.Content = tempName[1]
	tempDC := dc.FindString(value)
	if tempDC != "" {
		special.DifficultClass = ExtractWholeInt(tempDC)
		special.SavingThrows = savingRegex.FindString(strings.ToLower(tempDC))
	}

	tempDice := hasSpecialDamage.FindString(value)
	if tempDice != "" {
		special.Damage = dices.FindString(strings.ReplaceAll(tempDice, " ", ""))
		special.DamageType = damageTypeRegex.FindString(strings.ToLower(tempDice))
	}

	if strings.Contains(value, "DC") {
		tempValue := strings.Split(value, "DC")
		if len(tempValue) != 0 {
			tempCondition := conditionRegex.FindString(strings.ToLower(tempValue[1]))
			if tempCondition != "" {
				special.DamageType = strings.ToLower(tempCondition)
			}
		}
	}

	return special
}

func parseSpellcast(value string) SpellCastAbility {
	var spell SpellCastAbility
	var reLevel = regexp.MustCompile(`(?m)( *...?-level spellcaster)`)
	var reAbility = regexp.MustCompile(`(?m)( spellcasting ability is *.+? )`)
	var abilityRegex = regexp.MustCompile(`(?m)(strength|dexterity|constitution|intelligence|wisdom|charisma)`)
	var reDC = regexp.MustCompile(`(?m)(spell save DC [0-9]+)`)
	var reAttack = regexp.MustCompile(`(?m)( \+[0-9]+ to hit with spell)`)
	var reSpellPrepared = regexp.MustCompile(`(?m)( \(at will\): *.+ )`)
	// prepared short list
	var reCantripsPrepared = regexp.MustCompile(`(?m)( \(at will\): *.+? *1st )`)
	var reFirstLevel = regexp.MustCompile(`(?m)(1st level \([0-9] slots\): *.+?  *2nd)`)
	var reSecondLevel = regexp.MustCompile(`(?m)(2nd level \([0-9] slots\): *.+?  *3rd)`)
	var reThirdLevel = regexp.MustCompile(`(?m)(3rd level \([0-9] slots\): *.+?  *4th)`)
	var reFourthLevel = regexp.MustCompile(`(?m)(4th level \([0-9] slots\): *.+?  *5th)`)
	var reFifthLevel = regexp.MustCompile(`(?m)(5th level \([0-9] slots\): *.+?  *6th)`)
	var reSixthLevel = regexp.MustCompile(`(?m)(6th level \([0-9] slots\): *.+?  *7th)`)
	var reSeventhLevel = regexp.MustCompile(`(?m)(7th level \([0-9] slots\): *.+?  *8th)`)
	var reEighthLevel = regexp.MustCompile(`(?m)(8th level \([0-9] slots\): *.+?  *9th)`)
	var reNinthLevel = regexp.MustCompile(`(?m)(9th level \([0-9] slot\): *.+?  *.+?  *.+? )`)
	var reSlots = regexp.MustCompile(`(?m)[0-9] slot`)
	tempLevel := reLevel.FindString(value)
	if tempLevel != "" {
		spell.Level = ExtractWholeInt(tempLevel)
	}
	tempAbility := reAbility.FindString(value)
	if tempAbility != "" {
		spell.Ability = abilityRegex.FindString(strings.ToLower(tempAbility))
	}
	tempDC := reDC.FindString(value)
	if tempDC != "" {
		spell.DifficultClass = ExtractWholeInt(tempDC)
	}
	tempAttack := reAttack.FindString(value)
	if tempAttack != "" {
		spell.Attack = ExtractWholeInt(tempAttack)
	}
	tempList := reSpellPrepared.FindString(value)
	if tempList != "" {
		spell.List = tempList
	}
	spell.CantripsList = []string{}
	spell.ListPerLevel = make(map[string][]string)
	spell.SlotsPerLevel = make(map[string]int)
	tempCantripsList := reCantripsPrepared.FindString(value)
	if tempCantripsList != "" {
		tempSlice := strings.Split(tempCantripsList, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")

		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "1st", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			spell.CantripsList = append(spell.CantripsList, finalSpellName)
		}
	}
	tempFirstLevel := reFirstLevel.FindString(value)
	if tempFirstLevel != "" {
		level := "level1"
		tempSlice := strings.Split(tempFirstLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "2nd", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempFirstLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempSecondLevel := reSecondLevel.FindString(value)
	if tempSecondLevel != "" {
		level := "level2"
		tempSlice := strings.Split(tempSecondLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "3rd", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempSecondLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempThirdLevel := reThirdLevel.FindString(value)
	if tempThirdLevel != "" {
		level := "level3"
		tempSlice := strings.Split(tempThirdLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "4th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempThirdLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempFourthLevel := reFourthLevel.FindString(value)
	if tempFourthLevel != "" {
		level := "level4"
		tempSlice := strings.Split(tempFourthLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "5th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempFourthLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempFifthLevel := reFifthLevel.FindString(value)
	if tempFifthLevel != "" {
		level := "level5"
		tempSlice := strings.Split(tempFifthLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "6th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempFifthLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempSixthLevel := reSixthLevel.FindString(value)
	if tempSixthLevel != "" {
		level := "level6"
		tempSlice := strings.Split(tempSixthLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "7th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempSixthLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempSeventhLevel := reSeventhLevel.FindString(value)
	if tempSeventhLevel != "" {
		level := "level7"
		tempSlice := strings.Split(tempSeventhLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "8th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempSeventhLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempEighthLevel := reEighthLevel.FindString(value)
	if tempEighthLevel != "" {
		level := "level8"
		tempSlice := strings.Split(tempEighthLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "9th", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempEighthLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}
	tempNinthLevel := reNinthLevel.FindString(value)
	if tempNinthLevel != "" {
		level := "level9"
		tempSlice := strings.Split(tempNinthLevel, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")
		levelSlice := []string{}
		for _, v := range tempSliceContent {
			tmpName := strings.ReplaceAll(v, "the", "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			if strings.Contains(finalSpellName, "-*the") {
				finalSpellName = strings.ReplaceAll(finalSpellName, "-*the", "")
			}
			levelSlice = append(levelSlice, finalSpellName)
		}
		slot := reSlots.FindString(tempNinthLevel)
		spell.SlotsPerLevel[level] = ExtractWholeInt(slot)
		spell.ListPerLevel[level] = levelSlice
	}

	return spell

}

func parseInnateSpellcast(value string) InnateSpellCastAbility {
	var spell InnateSpellCastAbility
	var reAbility = regexp.MustCompile(`(?m)( spellcasting ability is *.+? )`)
	var reDC = regexp.MustCompile(`(?m)(spell save DC [0-9]+)`)
	var abilityRegex = regexp.MustCompile(`(?m)(strength|dexterity|constitution|intelligence|wisdom|charisma)`)
	var reInnateSpell = regexp.MustCompile(`(?m)( At will: *.+ )`)
	// exclude if found "following"
	var reMephitShortSpellList = regexp.MustCompile(`(?m)( can innately cast *.? *.+?, )`)
	// Innate Spellcast list
	// var reOnePerDay = regexp.MustCompile(`(?m)(1/day each: *.+?  )`)
	// var reThreePerDay = regexp.MustCompile(`(?m)(3/day each: *.+?  )`)
	var reAtWill = regexp.MustCompile(`(?m)( At will: *.+? *day)`)
	var rePerDay = regexp.MustCompile(`(?m)([0-9]/day)`)
	var reThreePerDay = regexp.MustCompile(`(?m)(3/day each: *.+?  )`)
	var reTwoPerDay = regexp.MustCompile(`(?m)(2/day each: *.+?  )`)
	var reOnePerDay = regexp.MustCompile(`(?m)(1/day each: *.+?  )`)
	tempAbility := reAbility.FindString(value)
	if tempAbility != "" {
		lowAbi := strings.ToLower(tempAbility)
		spell.Ability = lowAbi
		ability := abilityRegex.FindString(lowAbi)
		if ability != "" {
			spell.Ability = ability
		}
	}
	tempDC := reDC.FindString(value)
	if tempDC != "" {
		spell.DifficultClass = ExtractWholeInt(tempDC)
	}
	tempList := reInnateSpell.FindString(value)
	if tempList != "" {
		spell.List = tempList
	}
	tempShortList := reMephitShortSpellList.FindString(value)
	if tempShortList != "" && !strings.Contains(tempShortList, "following") {
		var magic string
		if strings.Contains(tempShortList, "fog cloud") {
			magic = "fog-cloud"
		}
		if strings.Contains(tempShortList, "heat metal") {
			magic = "heat-metal"
		}
		if strings.Contains(tempShortList, "blur") {
			magic = "blur"
		}
		spell.List = magic
	}
	spell.AtWillList = []string{}
	spell.OnePerDayList = []string{}
	spell.TwoPerDayList = []string{}
	spell.ThreePerDayList = []string{}
	tempAtWillList := reAtWill.FindString(value)
	if tempAtWillList != "" {
		tempSlice := strings.Split(tempAtWillList, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")

		for _, v := range tempSliceContent {
			tmpName := rePerDay.ReplaceAllString(v, "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			if strings.Contains(finalSpellName, "-(self-only)") {
				finalSpellName = strings.ReplaceAll(finalSpellName, "-(self-only)", "")
			}
			spell.AtWillList = append(spell.AtWillList, finalSpellName)
		}
	}
	tempThreePerDay := reThreePerDay.FindString(value)
	if tempThreePerDay != "" {
		tempSlice := strings.Split(tempThreePerDay, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")

		for _, v := range tempSliceContent {
			tmpName := rePerDay.ReplaceAllString(v, "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			if strings.Contains(finalSpellName, "-(self-only)") {
				finalSpellName = strings.ReplaceAll(finalSpellName, "-(self-only)", "")
			}
			spell.ThreePerDayList = append(spell.ThreePerDayList, finalSpellName)
		}
	}
	tempTwoPerDay := reTwoPerDay.FindString(value)
	if tempTwoPerDay != "" {
		tempSlice := strings.Split(tempTwoPerDay, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")

		for _, v := range tempSliceContent {
			tmpName := rePerDay.ReplaceAllString(v, "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			if strings.Contains(finalSpellName, "-(self-only)") {
				finalSpellName = strings.ReplaceAll(finalSpellName, "-(self-only)", "")
			}
			spell.TwoPerDayList = append(spell.TwoPerDayList, finalSpellName)
		}
	}
	tempOnePerDay := reOnePerDay.FindString(value)
	if tempOnePerDay != "" {
		tempSlice := strings.Split(tempOnePerDay, ":")
		tempSliceContent := strings.Split(tempSlice[1], ",")

		for _, v := range tempSliceContent {
			tmpName := rePerDay.ReplaceAllString(v, "")
			finalSpellName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(tmpName), " ", "-"))
			if strings.Contains(finalSpellName, "-(self-only)") {
				finalSpellName = strings.ReplaceAll(finalSpellName, "-(self-only)", "")
			}
			spell.OnePerDayList = append(spell.OnePerDayList, finalSpellName)
		}
	}

	return spell
}

func parseSavings(monster Monster) map[string]int {
	savings := make(map[string]int)
	modStr := calcAbilityModifier(ExtractWholeInt(monster.STR))
	modDex := calcAbilityModifier(ExtractWholeInt(monster.DEX))
	modCon := calcAbilityModifier(ExtractWholeInt(monster.CON))
	modInt := calcAbilityModifier(ExtractWholeInt(monster.INT))
	modWis := calcAbilityModifier(ExtractWholeInt(monster.WIS))
	modCar := calcAbilityModifier(ExtractWholeInt(monster.CHA))
	// "strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"
	savings["strength"] = modStr
	savings["dexterity"] = modDex
	savings["constitution"] = modCon
	savings["intelligence"] = modInt
	savings["wisdom"] = modWis
	savings["charisma"] = modCar
	if monster.SavingThrows != "" {
		savingsSlice := strings.Split(monster.SavingThrows, ",")
		for _, v := range savingsSlice {
			h := strings.Split(v, "+")
			value := ExtractWholeInt(strings.TrimSpace(h[1]))
			// value, err := strconv.Atoi()
			// if err != nil {
			// 	fmt.Printf("error convert string to int %v for %s", err, monster.Name)
			// }
			switch strings.TrimSpace(h[0]) {
			case "STR":
				if value > modStr {
					savings["strength"] = value
				}
			case "DEX":
				if value > modDex {
					savings["dexterity"] = value
				}
			case "CON":
				if value > modCon {
					savings["constitution"] = value
				}
			case "INT":
				if value > modInt {
					savings["intelligence"] = value
				}
			case "WIS":
				if value > modWis {
					savings["wisdom"] = value
				}
			case "CHA":
				if value > modCar {
					savings["charisma"] = value
				}

			}
		}
	}
	return savings
}

func parseAbility(monster Monster) map[string]int {
	ability := make(map[string]int)
	modStr := ExtractWholeInt(monster.STR)
	modDex := ExtractWholeInt(monster.DEX)
	modCon := ExtractWholeInt(monster.CON)
	modInt := ExtractWholeInt(monster.INT)
	modWis := ExtractWholeInt(monster.WIS)
	modCar := ExtractWholeInt(monster.CHA)
	ability["strength"] = modStr
	ability["dexterity"] = modDex
	ability["constitution"] = modCon
	ability["intelligence"] = modInt
	ability["wisdom"] = modWis
	ability["charisma"] = modCar
	return ability
}

// calcAbilityModifier func
func calcAbilityModifier(attr int) int {
	result := math.Floor((float64(attr) - 10) / 2)
	return int(result)
}

// func cleanName(value string) string {

// 	strings.ToLower(strings.ReplaceAll(name, " ", "-"))

// }
