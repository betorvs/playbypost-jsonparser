package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Armor struct
type Armor struct {
	Name              string `json:"name"`
	Title             string `json:"title"`
	Kind              string `json:"kind"`
	Cost              int    `json:"cost"`
	CoinType          string `json:"coin_type"`
	ArmorClass        int    `json:"armor_class"`
	DexterityModifier int    `json:"dexterity_modifier"`
	Stealth           bool   `json:"stealth"`
	Strength          int    `json:"strength"`
	Weight            int    `json:"weight"`
	Measure           string `json:"measure"`
}

// Weapon struct
type Weapon struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	Kind           string `json:"kind"`
	Cost           int    `json:"cost"`
	CoinType       string `json:"coin_type"`
	Damage         string `json:"damage"`
	DamageTwoHands string `json:"damage_two_hands,omitempty"`
	DamageType     string `json:"damage_type"`
	Weight         int    `json:"weight"`
	Measure        string `json:"measure"`
	Properties     string `json:"properties"`
}

//ArmorPrint func
func ArmorPrint() {
	DB, err := os.Open("./Armors-List.json")
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
	var ArmorsList []Armor
	for key, value := range result {
		// fmt.Println(key)
		kind := FixName(key)
		// var table map[string][]string
		table := value.(map[string]interface{})
		// fmt.Println(kind)
		stealthSlice := []string{}
		weightSlice := []int{}
		measureSlice := []string{}
		armorSlice := []string{}
		costSlice := []int{}
		coinSlice := []string{}
		armorClassSlice := []int{}
		maxDexSlice := []int{}
		strSlice := []int{}
		for _, a := range table {
			// fmt.Println(k, a)
			am := a.(map[string]interface{})
			// fmt.Println(armor)
			// var armor Armor
			// fmt.Println(kind)

			for kb, b := range am {
				// fmt.Println(kb)

				// armor.Kind = kind

				item := b.([]interface{})
				// fmt.Println(len(item))

				// fmt.Println(kind)
				switch kb {
				case "Stealth":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						stealthSlice = append(stealthSlice, cv)
						// fmt.Println(cv)
						// armor.Stealth = false
						// if cv == "Disadvantage" {
						// armor.Stealth = true
						// }
					}
				case "Weight":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						weight := ExtractWholeInt(cv)
						measure := ExtractWholeString(cv)
						weightSlice = append(weightSlice, weight)
						measureSlice = append(measureSlice, measure)
						// fmt.Println(weight, measure)
						// armor.Weight = weight
						// armor.Measure = measure
					}
				case "Armor":
					for _, c := range item {
						// fmt.Println(kc, c)
						cv := fmt.Sprintf("%v", c)
						armorSlice = append(armorSlice, cv)
						// fmt.Println(cv)
						// armor.Name = strings.ToLower(cv)
					}
				case "Cost":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						cost := ExtractWholeInt(cv)
						coin := ExtractWholeString(cv)
						costSlice = append(costSlice, cost)
						coinSlice = append(coinSlice, coin)
						// fmt.Println(cost, coin)
						// armor.Cost = cost
						// armor.CoinType = coin
					}
				case "Armor Class (AC)":
					for _, c := range item {
						// fmt.Println(kc, c)
						cv := fmt.Sprintf("%v", c)
						var ac int
						if strings.Contains(cv, "Dex") {
							sv := strings.Split(cv, "+")
							ac = ExtractWholeInt(sv[0])

						} else {
							ac = ExtractWholeInt(cv)
						}
						switch kind {
						case "light-armor", "shield":
							maxDexSlice = append(maxDexSlice, 99)
						case "medium-armor":
							maxDexSlice = append(maxDexSlice, 2)
						case "heavy-armor":
							maxDexSlice = append(maxDexSlice, 0)
						}
						armorClassSlice = append(armorClassSlice, ac)
						// fmt.Println(ac)

						// armor.ArmorClass = ac
					}
				case "Strength":
					for _, c := range item {
						// fmt.Println(kc, c)
						cv := fmt.Sprintf("%v", c)

						// fmt.Println(cv)
						if !strings.Contains(cv, "Str") {
							strSlice = append(strSlice, 0)
						} else {
							// fmt.Println(cv)
							str := ExtractWholeInt(cv)

							strSlice = append(strSlice, str)
							// fmt.Println(0)
						}
						// armor.Strength = str
					}
				}

			}

		}
		// fmt.Println(kind, armorSlice, armorClassSlice, stealthSlice, weightSlice, measureSlice, strSlice)
		sum := 0
		for i := 0; i <= len(armorSlice)-1; i++ {
			// fmt.Println(i)
			// fmt.Println(kind, armorSlice[i], costSlice[i], coinSlice[i], armorClassSlice[i], maxDexSlice[i], stealthSlice[i], weightSlice[i], measureSlice[i], strSlice[i])
			var armor Armor
			armor.Kind = strings.ReplaceAll(kind, " ", "-")
			armor.Name = FixName(armorSlice[i])
			armor.Title = armorSlice[i]
			armor.ArmorClass = armorClassSlice[i]
			armor.DexterityModifier = maxDexSlice[i]
			armor.Cost = costSlice[i]
			armor.CoinType = CoinType(coinSlice[i])
			armor.Stealth = false
			if stealthSlice[i] == "Disadvantage" {
				armor.Stealth = true
			}
			armor.Weight = weightSlice[i]
			armor.Measure = measureSlice[i]
			armor.Strength = strSlice[i]
			ArmorsList = append(ArmorsList, armor)
			sum += i
		}

	}
	prettyJSON, err := json.MarshalIndent(ArmorsList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}

//WeaponPrint func
func WeaponPrint() {
	DB, err := os.Open("./Weapons-List.json")
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
	var WeaponsList []Weapon
	for key, value := range result {
		kind := strings.ToLower(key)
		table := value.(map[string]interface{})
		nameSlice := []string{}
		weightSlice := []int{}
		measureSlice := []string{}
		damageSlice := []string{}
		damageTypeSlice := []string{}
		propertiesSlice := []string{}
		costSlice := []int{}
		coinSlice := []string{}
		for _, a := range table {
			// fmt.Println(k, a)
			am := a.(map[string]interface{})
			// fmt.Println(armor)
			// var armor Armor
			// fmt.Println(kind)

			for kb, b := range am {
				// fmt.Println(kb)

				// armor.Kind = kind

				item := b.([]interface{})
				// fmt.Println(len(item))

				// fmt.Println(kind)
				switch kb {
				case "Name":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						nameSlice = append(nameSlice, cv)
					}

				case "Properties":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						propertiesSlice = append(propertiesSlice, cv)
					}
				case "Weight":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						measure := ExtractWholeString(cv)
						measureSlice = append(measureSlice, measure)
						if strings.Contains(cv, "lb") {
							weight := ExtractWholeInt(cv)
							weightSlice = append(weightSlice, weight)
						} else {
							weightSlice = append(weightSlice, 0)
						}
						// fmt.Println(weight, measure)
						// armor.Weight = weight
						// armor.Measure = measure
					}
				case "Damage":
					for _, c := range item {
						// fmt.Println(kc, c)
						cv := fmt.Sprintf("%v", c)
						if strings.Contains(cv, " ") {
							cvs := strings.Split(cv, " ")
							damageSlice = append(damageSlice, cvs[0])
							damageTypeSlice = append(damageTypeSlice, cvs[1])
						} else {
							damageSlice = append(damageSlice, "-")
							damageTypeSlice = append(damageTypeSlice, "-")
						}

						// fmt.Println(cv)
						// armor.Name = strings.ToLower(cv)
					}
				case "Cost":
					for _, c := range item {
						cv := fmt.Sprintf("%v", c)
						cost := ExtractWholeInt(cv)
						coin := CoinType(ExtractWholeString(cv))
						costSlice = append(costSlice, cost)
						coinSlice = append(coinSlice, coin)
						// fmt.Println(cost, coin)
						// armor.Cost = cost
						// armor.CoinType = coin
					}
				}

			}

		}
		// fmt.Println(kind, armorSlice, armorClassSlice, stealthSlice, weightSlice, measureSlice, strSlice)
		sum := 0
		for i := 0; i <= len(nameSlice)-1; i++ {
			// fmt.Println(i)
			// fmt.Println(kind, armorSlice[i], costSlice[i], coinSlice[i], armorClassSlice[i], maxDexSlice[i], stealthSlice[i], weightSlice[i], measureSlice[i], strSlice[i])
			var weapon Weapon
			switch kind {
			case "simple melee weapons", "simple ranged weapons":
				weapon.Kind = "simple-weapon"
			case "martial melee weapons", "martial ranged weapons":
				weapon.Kind = "martial-weapon"
			}
			weapon.Title = nameSlice[i]
			weapon.Name = FixName(nameSlice[i])
			if strings.Contains(nameSlice[i], "Crossbow") {
				tempName := strings.ReplaceAll(nameSlice[i], ",", "")
				weapon.Name = FixName(tempName)
			}
			weapon.Damage = damageSlice[i]
			switch weapon.Name {
			case "spear", "trident", "quarterstaff":
				weapon.DamageTwoHands = "1d8"
			case "battleaxe", "longsword", "warhammer":
				weapon.DamageTwoHands = "1d10"
			}
			weapon.DamageType = damageTypeSlice[i]
			weapon.Properties = propertiesSlice[i]
			weapon.Cost = costSlice[i]
			weapon.CoinType = coinSlice[i]
			weapon.Weight = weightSlice[i]
			weapon.Measure = measureSlice[i]
			WeaponsList = append(WeaponsList, weapon)
			sum += i
		}

	}
	prettyJSON, err := json.MarshalIndent(WeaponsList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))

}
