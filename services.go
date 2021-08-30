package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Services struct
type Services struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Cost     int    `json:"cost"`
	CoinType string `json:"coin_type"`
	Unit     string `json:"unit"`
	Source   string `json:"source"`
}

//ServicesPrint func
func ServicesPrint() {
	DB, err := os.Open("./services-list.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	scanner := bufio.NewScanner(DB)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	var serviceList []Services
	for _, eachline := range txtlines {
		splitLine := strings.Split(eachline, ":")
		// fmt.Println(splitLine)
		var service Services
		service.Name = strings.ToLower(splitLine[0])
		service.Title = splitLine[0]
		service.Cost = ExtractWholeInt(splitLine[1])
		service.CoinType = splitLine[2]
		service.Unit = splitLine[3]
		service.Source = splitLine[4]
		serviceList = append(serviceList, service)
	}
	prettyJSON, err := json.MarshalIndent(serviceList, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
}
