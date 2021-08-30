package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TreasureHoard struct
type TreasureHoard struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	CoinType    string `json:"coin_type"`
}

//GemsArtPrint func
func GemsArtPrint() {
	DB, err := os.Open("./gems-list.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	DB2, err2 := os.Open("./arts-list.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer DB2.Close()

	scanner := bufio.NewScanner(DB)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	var hoardList []TreasureHoard
	for _, eachline := range txtlines {
		// fmt.Println(eachline)
		var hoard TreasureHoard
		hoard.Kind = "gemstone"
		hoard.CoinType = "gold"

		splitLine := strings.Split(eachline, ":")
		hoard.Name = ExtractWholeString(splitLine[0])
		hoard.Value = ExtractWholeInt(splitLine[0])
		hoard.Description = splitLine[1]
		hoardList = append(hoardList, hoard)
	}

	scanner2 := bufio.NewScanner(DB2)
	scanner2.Split(bufio.ScanLines)
	var txtlines2 []string

	for scanner2.Scan() {
		txtlines2 = append(txtlines2, scanner2.Text())
	}
	// var hoardList HoardList
	for _, eachline2 := range txtlines2 {
		// fmt.Println(eachline2)
		var hoard TreasureHoard
		hoard.Kind = "art-object"
		hoard.CoinType = "gold"
		splitLine := strings.Split(eachline2, ":")
		hoard.Value = ExtractWholeInt(splitLine[0])
		hoard.Name = splitLine[1]
		hoard.Description = strings.TrimLeft(splitLine[2], " ")
		hoardList = append(hoardList, hoard)
	}
	prettyJSON, err := json.MarshalIndent(hoardList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}
