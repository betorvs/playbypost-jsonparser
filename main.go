package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	argsWithProg := os.Args
	err := printJSON(argsWithProg[1])
	if err != nil {
		fmt.Printf("error to print: %v\n", err)
	}

}

func printJSON(value string) error {
	switch value {
	case "monster", "m":
		MonsterPrint()
		return nil
	case "spell", "s":
		SpellPrint()
		return nil
	case "magicitem", "mi", "item", "magic":
		NewMagicItemPrint()
		return nil
	case "newmagicitem":
		NewMagicItemPrint()
		return nil
	case "armor":
		ArmorPrint()
		return nil
	case "weapon":
		WeaponPrint()
		return nil
	case "gear":
		GearPrint()
		return nil
	case "packs":
		GearPackPrint()
		return nil
	case "tools":
		ToolsPrint()
		return nil
	case "mounts":
		MountsPrint()
		return nil
	case "hoard":
		GemsArtPrint()
		return nil
	case "services":
		ServicesPrint()
		return nil
	default:
		fmt.Println("Use: go run main.go [monster|spell|magicitem]")
		err := fmt.Errorf("without args")
		return err
	}

}

// StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ExtractWholeInt func
func ExtractWholeInt(value string) int {
	var re = regexp.MustCompile(`[^0-9]+`)
	temp := re.ReplaceAllString(value, "")
	// fmt.Println(temp)
	numbers, err := strconv.Atoi(strings.TrimSpace(temp))
	if err != nil {
		fmt.Printf("error extractWholeInt convert string to int %v", err)
	}
	return numbers
}

// ExtractWholeString func
func ExtractWholeString(value string) string {
	var re = regexp.MustCompile(`[^a-z]+`)
	return re.ReplaceAllString(value, "")
	// numbers, err := strconv.Atoi(strings.TrimSpace(temp))
	// if err != nil {
	// 	fmt.Printf("error extractWholeInt convert string to int %v", err)
	// }
	// return numbers
}

// CoinType return coin in string
func CoinType(kind string) string {
	switch kind {
	case "platinum", "pp":
		return "platinum"
	case "gold", "gp":
		return "gold"
	case "electrum", "ep":
		return "electrum"
	case "silver", "sp":
		return "silver"

	case "copper", "cp":
		return "copper"

	}
	return kind
}

// FixName returns a string with ToLower and Replace " " with "-"
func FixName(value string) string {
	newValue := strings.ReplaceAll(value, " ", "-")
	return strings.ToLower(newValue)
}
