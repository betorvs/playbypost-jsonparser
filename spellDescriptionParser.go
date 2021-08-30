package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type spell struct {
	Name              string   `json:"name"`
	Level             int      `json:"level"`
	Title             string   `json:"title"`
	Subtitle          string   `json:"subtitle"`
	CastingTime       string   `json:"casting_time"`
	Range             string   `json:"range"`
	Components        string   `json:"components"`
	Duration          string   `json:"duration"`
	Description       string   `json:"description"`
	AtHigherLevels    string   `json:"at_higher_levels,omitempty"`
	DamageIncrease    string   `json:"damage_increase,omitempty"`
	DamageDice        string   `json:"damage_dice,omitempty"`
	DamageType        string   `json:"damage_type,omitempty"`
	SavingThrow       string   `json:"saving_throw,omitempty"`
	HealDice          string   `json:"heal_dice,omitempty"`
	HealingIncreases  string   `json:"healing_increase,omitempty"`
	ExtraDice         string   `json:"extra_dice,omitempty"`
	ExtraDiceUsage    []string `json:"extra_dice_usage,omitempty"`
	BonusArmorClass   int      `json:"bonus_armor_class,omitempty"`
	BaseArmorClass    int      `json:"base_armor_class,omitempty"`
	MinimumArmorClass int      `json:"minimum_armor_class,omitempty"`
	AttackRolls       bool     `json:"attack_rolls,omitempty"`
}

//SpellPrint func
func SpellPrint() {
	DB, err := os.Open("./Spell-Descriptions.json")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	// var re = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
	var spellList []spell
	for key, value := range result {
		spell := new(spell)
		name := strings.ToLower(strings.ReplaceAll(key, " ", "-"))
		if name == "temperature" {
			continue
		}
		if name == "wind" {
			continue
		}
		if name == "precipitation" {
			continue
		}
		spell.Name = name
		spell.Title = key
		content := value.(map[string]interface{})
		for _, v := range content {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(v)

				for i := 0; i < s.Len(); i++ {
					value := fmt.Sprintln(s.Index(i))
					if strings.Contains(value, "Casting Time:") {
						cast := strings.Split(value, ":")
						spell.CastingTime = strings.TrimSpace(strings.TrimSuffix(strings.ReplaceAll(cast[1], "**", ""), "\n"))
					}
					if strings.Contains(value, "Range:") {
						rang := strings.Split(value, ":")
						spell.Range = strings.TrimSpace(strings.TrimSuffix(strings.ReplaceAll(rang[1], "**", ""), "\n"))
					}
					if strings.Contains(value, "Components:") {
						comp := strings.Split(value, ":")
						spell.Components = strings.TrimSpace(strings.TrimSuffix(strings.ReplaceAll(comp[1], "**", ""), "\n"))
					}
					if strings.Contains(value, "Duration:") {
						dura := strings.Split(value, ":")
						spell.Duration = strings.TrimSpace(strings.TrimSuffix(strings.ReplaceAll(dura[1], "**", ""), "\n"))
					}
					if validateLevel(value) {
						spell.Subtitle = strings.ToLower(strings.TrimSuffix(strings.ReplaceAll(value, "*", ""), "\n"))
						spell.Level = parseLevel(spell.Subtitle)
					}
					if !strings.Contains(value, "Casting Time:") && !strings.Contains(value, "Range:") && !strings.Contains(value, "Components:") && !strings.Contains(value, "Duration:") && !strings.Contains(value, "-level ") && !strings.Contains(value, " cantrip*") {
						spell.Description += fmt.Sprintf(" %s", strings.ReplaceAll(strings.TrimSuffix(value, "\n"), "***", "*"))
						if strings.Contains(value, "At Higher Levels") {
							spell.AtHigherLevels = strings.ReplaceAll(strings.TrimSuffix(value, "\n"), "***", "")
						}
						if strings.Contains(value, "This spell’s damage increases") {
							spell.DamageIncrease = strings.TrimSuffix(value, "\n")
						}
					}
				}
			}
		}
		spellList = append(spellList, *spell)
	}
	var verifySpellList []spell
	for _, item := range spellList {

		var hasSpecialDamage = regexp.MustCompile(`(?m)(takes (\d+)?d(\d+)([\+\-]\d+)? *.+? damage)`)
		var damageTypeRegex = regexp.MustCompile(`(?m)(acid|cold|fire|force|lightning|necrotic|poison|psychic|radiant|thunder|bludgeoning|piercing|slashing)`)
		var dices = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
		var savingThrowRegex = regexp.MustCompile(`(?m)(must *.+? a *.+? saving throw)`)
		var savingRegex = regexp.MustCompile(`(?m)(strength|dexterity|constitution|intelligence|wisdom|charisma)`)
		var reHealDescription = regexp.MustCompile(`(?m)(hit points equal to (\d+)?d(\d+)([\+\-]\d+)? \+ your spellcasting ability modifier)`)
		var reHealingIncreases = regexp.MustCompile(`(?m)(healing increases by (\d+)?d(\d+)([\+\-]\d+)? for each slot level above *.+?\.)`)
		var reBonusAC = regexp.MustCompile(`(?m)(\+[0-9] bonus to AC)`)
		var reAttack = regexp.MustCompile(`(?m)(ake a *.+? spell attack)`)
		damageDescription := hasSpecialDamage.FindString(item.Description)
		if damageDescription != "" {
			damageType := damageTypeRegex.FindString(strings.ToLower(damageDescription))
			spellItem := new(spell)
			spellItem = &item
			spellItem.DamageDice = dices.FindString(strings.ReplaceAll(damageDescription, " ", ""))
			spellItem.DamageType = damageType
		}
		verifySavingThrow := savingThrowRegex.FindString(item.Description)
		if verifySavingThrow != "" {
			spellItem := new(spell)
			spellItem = &item
			saving := savingRegex.FindString(strings.ToLower(verifySavingThrow))
			spellItem.SavingThrow = saving
		}
		verifyHealMagic := reHealDescription.FindString(item.Description)
		if verifyHealMagic != "" {
			diceHeal := dices.FindString(strings.ReplaceAll(verifyHealMagic, " ", ""))
			spellItem := new(spell)
			spellItem = &item
			spellItem.HealDice = diceHeal
		}
		verifyHealingIncreases := reHealingIncreases.FindString(item.Description)
		if verifyHealingIncreases != "" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.HealingIncreases = verifyHealingIncreases
		}
		verifyAttackRolls := reAttack.FindString(item.Description)
		if verifyAttackRolls != "" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.AttackRolls = true
		}

		if item.Name == "resistance" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.ExtraDice = "1d4"
			spellItem.ExtraDiceUsage = []string{"saving"}
		}
		if item.Name == "bless" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.ExtraDice = "1d4"
			spellItem.ExtraDiceUsage = []string{"attack", "saving"}
		}
		if item.Name == "guidance" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.ExtraDice = "1d4"
			spellItem.ExtraDiceUsage = []string{"ability"}
		}
		verifyArmorClassBonus := reBonusAC.FindString(item.Description)
		if verifyArmorClassBonus != "" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.BonusArmorClass = ExtractWholeInt(verifyArmorClassBonus)
		}
		if item.Name == "mage-armor" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.BaseArmorClass = 13
		}
		if item.Name == "barkskin" {
			spellItem := new(spell)
			spellItem = &item
			spellItem.MinimumArmorClass = 16
		}

		// savingRegex.FindString(strings.ToLower(item.Content))
		verifySpellList = append(verifySpellList, item)
	}
	prettyJSON, err := json.MarshalIndent(verifySpellList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}

// func extractSavingThrow(s string) string {

// 	if strings.Contains(s, "Strength saving throw") {
// 		return "strength"
// 	}
// 	if strings.Contains(s, "Dexterity saving throw") {
// 		return "dexterity"
// 	}
// 	if strings.Contains(s, "Constitution saving throw") {
// 		return "constitution"
// 	}
// 	if strings.Contains(s, "Wisdom saving throw") {
// 		return "wisdom"
// 	}
// 	if strings.Contains(s, "Charisma saving throw") {
// 		return "charisma"
// 	}
// 	if strings.Contains(s, "Intelligence saving throw") {
// 		return "intelligence"
// 	}
// 	return ""
// }

// func parseDamageType(s string) string {
// 	if strings.Contains(s, "acid damage") {
// 		return "acid"
// 	}
// 	if strings.Contains(s, "bludgeoning damage") {
// 		return "bludgeoning"
// 	}
// 	if strings.Contains(s, "cold damage") {
// 		return "cold"
// 	}
// 	if strings.Contains(s, "fire damage") {
// 		return "fire"
// 	}
// 	if strings.Contains(s, "force damage") {
// 		return "force"
// 	}
// 	if strings.Contains(s, "lightning damage") {
// 		return "lightning"
// 	}
// 	if strings.Contains(s, "necrotic damage") {
// 		return "necrotic"
// 	}
// 	if strings.Contains(s, "piercing damage") {
// 		return "piercing"
// 	}
// 	if strings.Contains(s, "poison damage") {
// 		return "poison"
// 	}
// 	if strings.Contains(s, "psychic damage") {
// 		return "psychic"
// 	}
// 	if strings.Contains(s, "radiant damage") {
// 		return "radiant"
// 	}
// 	if strings.Contains(s, "slashing damage") {
// 		return "slashing"
// 	}
// 	if strings.Contains(s, "thunder damage") {
// 		return "thunder"
// 	}
// 	return ""
// }

func parseLevel(title string) int {

	if strings.Contains(title, "cantrip") {
		return 0
	}
	if strings.Contains(title, "1st-level") {
		return 1
	}
	if strings.Contains(title, "2nd-level") {
		return 2
	}
	if strings.Contains(title, "3rd-level") {
		return 3
	}
	if strings.Contains(title, "4th-level") {
		return 4
	}
	if strings.Contains(title, "5th-level") {
		return 5
	}
	if strings.Contains(title, "6th-level") {
		return 6
	}
	if strings.Contains(title, "7th-level") {
		return 7
	}
	if strings.Contains(title, "8th-level") {
		return 8
	}
	if strings.Contains(title, "9th-level") {
		return 9
	}
	return 0
}

func validateLevel(title string) bool {

	if strings.Contains(title, "cantrip*") {
		return true
	}
	if strings.Contains(title, "*1st-level") {
		return true
	}
	if strings.Contains(title, "*2nd-level") {
		return true
	}
	if strings.Contains(title, "*3rd-level") {
		return true
	}
	if strings.Contains(title, "*4th-level") {
		return true
	}
	if strings.Contains(title, "*5th-level") {
		return true
	}
	if strings.Contains(title, "*6th-level") {
		return true
	}
	if strings.Contains(title, "*7th-level") {
		return true
	}
	if strings.Contains(title, "*8th-level") {
		return true
	}
	if strings.Contains(title, "*9th-level") {
		return true
	}
	return false
}

// func levelNumeric() []string {
// 	return []string{"1st-level", "2nd-level", "3rd-level", "4th-level", "5th-level", "6th-level", "7th-level", "8th-level", "9th-level"}
// }
