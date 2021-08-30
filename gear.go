package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Gear struct
type Gear struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Kind     string `json:"kind"`
	Cost     int    `json:"cost"`
	CoinType string `json:"coin_type"`
	Weight   int    `json:"weight"`
	Measure  string `json:"measure"`
	Number   int    `json:"number"`
}

// Packs struct
type Packs struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	CoinType    string `json:"coin_type"`
}

// Tools struct
type Tools struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Kind        string `json:"kind"`
	Cost        int    `json:"cost"`
	CoinType    string `json:"coin_type"`
	Weight      int    `json:"weight"`
	Measure     string `json:"measure"`
	Description string `json:"description"`
}

// Mounts struct
type Mounts struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Cost                    int    `json:"cost"`
	CoinType                string `json:"coin_type"`
	CarryingCapacity        int    `json:"carrying_capacity"`
	CarryingCapacityMeasure string `json:"carrying_capacity_measure"`
	Speed                   int    `json:"speed"`
	SpeedMeasure            string `json:"speed_measure"`
}

//GearPrint func
func GearPrint() {
	DB, err := os.Open("./Adventure-Gear.json")
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
	itemSlice := []string{}
	costSlice := []int{}
	coinSlice := []string{}
	weightSlice := []int{}
	measureSlice := []string{}
	for key, value := range result {
		// fmt.Println(key, value)
		item := value.([]interface{})
		switch key {
		case "Item":
			for _, c := range item {
				cv := fmt.Sprintf("%v", c)
				itemSlice = append(itemSlice, cv)
			}
		case "Cost":
			for _, c := range item {
				cv := fmt.Sprintf("%v", c)
				if cv != "" {
					cost := ExtractWholeInt(cv)
					coin := ExtractWholeString(cv)
					costSlice = append(costSlice, cost)
					coinSlice = append(coinSlice, coin)
				} else {
					costSlice = append(costSlice, 0)
					coinSlice = append(coinSlice, "")
				}

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
			}
		}
	}
	var gearList []Gear
	// sum := 0
	for i := 0; i <= len(itemSlice)-1; i++ {
		if itemSlice[i] == "*Ammunition*" || itemSlice[i] == "*Arcane focus*" || itemSlice[i] == "*Druidic focus*" || itemSlice[i] == "*Holy symbol*" {
			continue
		}
		var gear Gear
		// var name, kind string
		name := itemSlice[i]
		kind := "Other"
		if strings.Contains(itemSlice[i], "(20)") {
			gear.Number = 20
			kind = "Ammunition"
		}
		if strings.Contains(itemSlice[i], "(50)") {
			gear.Number = 50
			kind = "Ammunition"
		}
		if itemSlice[i] == "Crystal" || itemSlice[i] == "Orb" || itemSlice[i] == "Rod" || itemSlice[i] == "Staff" || itemSlice[i] == "Wand" {
			kind = "Arcane Focus"
		}
		if itemSlice[i] == "Sprig of mistletoe" || itemSlice[i] == "Totem" || itemSlice[i] == "Wooden staff" || itemSlice[i] == "Yew wand" {
			kind = "Druidic Focus"
		}
		if itemSlice[i] == "Amulet" || itemSlice[i] == "Emblem" || itemSlice[i] == "Reliquary" {
			kind = "Holy Symbol"
		}
		if strings.Contains(itemSlice[i], "’") {
			name = strings.ReplaceAll(itemSlice[i], "’", "")
		}
		if strings.Contains(itemSlice[i], ",") {
			name = strings.ReplaceAll(itemSlice[i], ",", "")
		}
		if strings.Contains(itemSlice[i], "Pick") {
			name = "pick miners"
		}

		if strings.Contains(itemSlice[i], "Scale") {
			name = "scale merchants"
		}
		gear.Name = strings.ToLower(name)
		gear.Title = name
		gear.Kind = strings.ToLower(kind)
		gear.CoinType = CoinType(coinSlice[i])
		gear.Cost = costSlice[i]
		gear.Weight = weightSlice[i]
		gear.Measure = measureSlice[i]
		gearList = append(gearList, gear)
		// fmt.Println(itemSlice[i], costSlice[i], coinSlice[i], weightSlice[i], measureSlice[i])
		// sum += i
	}
	prettyJSON, err := json.MarshalIndent(gearList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))

}

// List struct
type List struct {
	Content []string
}

//GearPackPrint func
func GearPackPrint() {
	DB, err := os.Open("./Gear-Packs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	var result List
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	var packsList []Packs
	for _, v := range result.Content {
		if strings.HasPrefix(v, "The starting equipment") {
			continue
		}
		// fmt.Println(v)
		content := strings.Split(v, ".")
		title := strings.Split(content[0], "(")
		name := strings.ReplaceAll(strings.ReplaceAll(title[0], "*", ""), "’", "")
		cost := ExtractWholeInt(title[1])
		coin := CoinType(ExtractWholeString(title[1]))
		desc := strings.ReplaceAll(content[1], "*", "")
		var packs Packs
		packs.Name = strings.ToLower(strings.TrimSpace(name))
		packs.Title = name
		packs.Cost = cost
		packs.CoinType = coin
		packs.Description = strings.TrimSpace(desc)
		packsList = append(packsList, packs)
		// fmt.Printf("name %s, cost %v, coin %s, desc %s \n", name, cost, coin, desc)

	}
	prettyJSON, err := json.MarshalIndent(packsList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
	// fmt.Println(result.Content)
}

type autoGeneratedTools struct {
	Table struct {
		Item   []string `json:"Item"`
		Cost   []string `json:"Cost"`
		Weight []string `json:"Weight"`
	} `json:"table"`
}

//ToolsPrint func
func ToolsPrint() {
	content := map[string]string{
		"Artisans tools":     "These special tools include the items needed to pursue a craft or trade. The table shows examples of the most common types of tools, each providing items related to a single craft. Proficiency with a set of artisan’s tools lets you add your proficiency bonus to any ability checks you make using the tools in your craft. Each type of artisan’s tools requires a separate proficiency.",
		"Disguise kit":       "This pouch of cosmetics, hair dye, and small props lets you create disguises that change your physical appearance. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to create a visual disguise.",
		"Forgery kit":        "This small box contains a variety of papers and parchments, pens and inks, seals and sealing wax, gold and silver leaf, and other supplies necessary to create convincing forgeries of physical documents. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to create a physical forgery of a document.",
		"Gaming set":         "This item encompasses a wide range of game pieces, including dice and decks of cards (for games such as Three-Dragon Ante). A few common examples appear on the Tools table, but other kinds of gaming sets exist. If you are proficient with a gaming set, you can add your proficiency bonus to ability checks you make to play a game with that set. Each type of gaming set requires a separate proficiency.",
		"Herbalism kit":      "This kit contains a variety of instruments such as clippers, mortar and pestle, and pouches and vials used by herbalists to create remedies and potions. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to identify or apply herbs. Also, proficiency with this kit is required to create antitoxin and potions of healing.",
		"Musical instrument": "Several of the most common types of musical instruments are shown on the table as examples. If you have proficiency with a given musical instrument, you can add your proficiency bonus to any ability checks you make to play music with the instrument. A bard can use a musical instrument as a spellcasting focus. Each type of musical instrument requires a separate proficiency.",
		"Navigators tools":   "This set of instruments is used for navigation at sea. Proficiency with navigator&#39;s tools lets you chart a ship&#39;s course and follow navigation charts. In addition, these tools allow you to add your proficiency bonus to any ability check you make to avoid getting lost at sea.",
		"Poisoners kit":      "A poisoner’s kit includes the vials, chemicals, and other equipment necessary for the creation of poisons. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to craft or use poisons.",
		"Thieves tools":      "This set of tools includes a small file, a set of lock picks, a small mirror mounted on a metal handle, a set of narrow-bladed scissors, and a pair of pliers. Proficiency with these tools lets you add your proficiency bonus to any ability checks you make to disarm traps or open locks.",
	}

	DB, err := os.Open("./Tools.json")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	var result autoGeneratedTools
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	var toolList []Tools
	// sum := 0
	for i := 0; i <= len(result.Table.Item)-1; i++ {
		if strings.Contains(result.Table.Item[i], "*") || strings.Contains(result.Table.Item[i], "Vehicles") {
			continue
		}
		name := strings.ReplaceAll(strings.ReplaceAll(result.Table.Item[i], "*", ""), "’", "")
		cost := ExtractWholeInt(result.Table.Cost[i])
		coin := CoinType(ExtractWholeString(result.Table.Cost[i]))
		var weight int
		var desc, kind string
		measure := "lb"
		if !strings.Contains(result.Table.Item[i], "set") {
			weight = ExtractWholeInt(result.Table.Weight[i])
			measure = ExtractWholeString(result.Table.Weight[i])
		}
		intruments := []string{
			"Bagpipes",
			"Drum",
			"Dulcimer",
			"Flute",
			"Lute",
			"Lyre",
			"Horn",
			"Pan flute",
			"Shawm",
			"Viol",
		}

		if StringInSlice(result.Table.Item[i], intruments) {
			kind = "Musical instrument"
			desc = content["Musical instrument"]
		}
		kind = "Artisans tools"
		desc = content["Artisans tools"]
		if strings.Contains(result.Table.Item[i], "set") {
			kind = "Gaming set"
			desc = content["Gaming set"]
		}
		if strings.Contains(result.Table.Item[i], "kit") {
			kind = name
			desc = content[name]
		}
		if strings.Contains(name, "Navigators tools") || strings.Contains(name, "Thieves tools") {
			kind = name
			desc = content[name]
		}
		var tool Tools
		tool.Name = strings.ToLower(name)
		tool.Title = name
		tool.Kind = strings.ToLower(kind)
		tool.Cost = cost
		tool.CoinType = coin
		tool.Weight = weight
		tool.Measure = measure
		tool.Description = desc
		toolList = append(toolList, tool)
		// fmt.Printf("name %s, cost %v, coin %s, weight %v %s , kind %s, desc %s\n", name, cost, coin, weight, measure, kind, desc)

	}

	prettyJSON, err := json.MarshalIndent(toolList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}

type autoGeneratedMounts struct {
	Item             []string `json:"Item"`
	Cost             []string `json:"Cost"`
	Speed            []string `json:"Speed"`
	CarryingCapacity []string `json:"Carrying Capacity"`
}

//MountsPrint func
func MountsPrint() {
	DB, err := os.Open("./Mounts.json")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	var result autoGeneratedMounts
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	var mountList []Mounts
	for i := 0; i <= len(result.Item)-1; i++ {
		name := strings.ReplaceAll(result.Item[i], ",", "")
		cost := ExtractWholeInt(result.Cost[i])
		coin := ExtractWholeString(result.Cost[i])
		carrying := ExtractWholeInt(result.CarryingCapacity[i])
		carryingmeasure := ExtractWholeString(result.CarryingCapacity[i])
		speed := ExtractWholeInt(result.Speed[i])
		speedmeasure := ExtractWholeString(result.Speed[i])
		// fmt.Printf("name %s, cost %v, coin %s, carrying capacity %v %s , speed %v %s\n", name, cost, coin, carrying, carryingmeasure, speed, speedmeasure)
		var mount Mounts
		mount.Name = strings.ToLower(name)
		mount.Title = name
		mount.Cost = cost
		mount.CoinType = CoinType(coin)
		mount.CarryingCapacity = carrying
		mount.CarryingCapacityMeasure = carryingmeasure
		mount.Speed = speed
		mount.SpeedMeasure = speedmeasure
		mountList = append(mountList, mount)
	}
	prettyJSON, err := json.MarshalIndent(mountList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))

}
