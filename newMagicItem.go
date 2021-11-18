package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// Common Constants to be used
const (
	Common         = "common"
	Uncommon       = "uncommon"
	Rare           = "rare"
	VeryRare       = "very rare"
	Legendary      = "legendary"
	ArmorCons      = "armor"
	Potion         = "potions"
	Rings          = "rings"
	Rods           = "rods"
	Staffs         = "staffs"
	Wands          = "wands"
	Weapons        = "weapons"
	WoundrousItems = "woundrous items"
	Heal           = "heal"         // to heal you or allies
	Spell          = "spell"        // magical effect from spell
	SpellAttack    = "spell-attack" // spell to attack enemies
	Strength       = "strenght"
	Dexterity      = "dexterity"
	Constitution   = "constitution"
	Intelligence   = "intelligence"
	Wisdom         = "wisdom"
	Charisma       = "charisma"
	Shield         = "shield"
	Bludgeoning    = "bludgeoning"
)

var (
	resistance       = []string{"acid", "cold", "fire", "force", "lightning", "necrotic", "poison", "psychic", "radiant", "thunder"}
	armorsList       = []string{"hide", "chain-shirt", "scale-mail", "breastplate", "half-plate", "ring-mail", "chain-mail", "splint", "plate", "padded", "leather", "studded-leather"}
	armorsMetalList  = []string{"chain-shirt", "scale-mail", "breastplate", "half-plate", "ring-mail", "chain-mail", "splint", "plate"}
	ammunitionList   = []string{"arrows", "blowgun-needles", "crossbow-bolts", "sling-bullets"}
	bonusList        = []int{1, 2, 3}
	weaponsList      = []string{"club", "dagger", "greatclub", "handaxe", "javelin", "light-hammer", "mace", "quarterstaff", "sickle", "spear", "crossbow light", "dart", "shortbow", "sling", "battleaxe", "flail", "glaive", "greataxe", "greatsword", "halberd", "lance", "longsword", "maul", "morningstar", "pike", "rapier", "scimitar", "shortsword", "trident", "war-pick", "warhammer", "whip", "blowgun", "crossbow-hand", "crossbow-heavy", "longbow", "net"}
	anySword         = []string{"greatsword", "longsword", "shortsword", "scimitar", "rapier"}
	giantNames       = map[string]int{"hill": 21, "stone": 23, "frost": 23, "fire": 25, "cloud": 27, "storm": 29}
	damageMeleeType  = []string{"bludgeoning", "piercing", "slashing"}
	allDamageType    = []string{"acid", "cold", "fire", "force", "lightning", "necrotic", "poison", "psychic", "radiant", "thunder", "bludgeoning", "piercing", "slashing"}
	spellcaster      = []string{"bard", "ranger", "sorcerer", "cleric", "warlock", "wizard", "druid", "paladin", "arcane-trickster", "eldritch-knight"}
	tableA           = []string{"A"}
	tableB           = []string{"B"}
	tableC           = []string{"C"}
	tableD           = []string{"D"}
	tableE           = []string{"E"}
	tableF           = []string{"F"}
	tableG           = []string{"G"}
	tableH           = []string{"H"}
	tableI           = []string{"I"}
	monstersTypeList = []string{"Aberrations", "Beasts", "Celestials", "Elementals", "Fey", "Fiends", "Plants", "Undead"}
)

// MagicItem struct
type MagicItem struct {
	Name                  string        `json:"name"`
	Title                 string        `json:"title"`
	Content               string        `json:"content"`
	Category              string        `json:"category"`
	Rarity                string        `json:"rarity"`
	HoardTable            []string      `json:"hoard_table"`
	AttunementRestriction []string      `json:"attunement_restriction,omitempty"`
	RequiredAttunement    bool          `json:"required_attunement"`
	RolePlay              bool          `json:"roleplay"`
	Forbidden             bool          `json:"forbidden"`
	Shape                 string        `json:"shape,omitempty"`
	Feature               *CoreFeatures `json:"magic_feature,omitempty"`
	Power                 *CorePowers   `json:"power,omitempty"`
	Scroll                *Scroll       `json:"scroll,omitempty"`
}

// CoreFeatures struct
// this kind of magic is always activated when in use
type CoreFeatures struct {
	AttackBonus           int            `json:"attack_bonus,omitempty"`           // in attack only
	DamageBonus           int            `json:"damage_bonus,omitempty"`           // in Damage only
	CombateBonus          int            `json:"combate_bonus,omitempty"`          // on attack and damage
	DamageBonusCritical   int            `json:"damage_bonus_critical,omitempty"`  // add bonus to damage if was a critical
	DamageDiceCritical    string         `json:"damage_dice_critical,omitempty"`   // add dice bonus to damage if was a critical
	DamageType            string         `json:"damage_type,omitempty"`            // type of damage from allDamageType
	SpellBonus            int            `json:"spell_bonus,omitempty"`            // on spell attack rolls and difficult class
	SpellAttackBonus      int            `json:"spell_attack_bonus,omitempty"`     // on spell attack rolls
	ArmorClassBonus       int            `json:"armor_class_bonus,omitempty"`      // AC
	SavingBonus           int            `json:"saving_bonus,omitempty"`           // on saving roll
	ExtraHitPointsLevel   int            `json:"extra_hit_points_level,omitempty"` // extra hit points per level
	AbilityBonus          int            `json:"ability_bonus,omitempty"`          // on ability checks roll
	NewAbility            map[string]int `json:"new_ability,omitempty"`            // Change Ability int
	IncreaseAbility       map[string]int `json:"increase_ability,omitempty"`       // Increase Ability int by
	SkillBonus            map[string]int `json:"skill_bonus,omitempty"`            // map skill and add bonus int in check
	WithProficiency       []string       `json:"with_proficiency,omitempty"`       // gives to user that proficiency
	Disvantages           []string       `json:"disvantages,omitempty"`            // list of Disvantages to add
	Advantages            []string       `json:"advantages,omitempty"`             // list of Disvantages to add
	AutoFail              []string       `json:"auto_fail,omitempty"`              // autoFail in something, like ability, skill, attack, anything, add to your list
	DamageResistance      []string       `json:"damage_resistance,omitempty"`      // damageResistance to add from allDamageType list one or more
	DamageVulnerabilities []string       `json:"damage_vulnerabilities,omitempty"` // DamageVulnerabilities to add from allDamageType list one or more
	DamageImmunities      []string       `json:"damage_immunities,omitempty"`      // DamageImmunities to add from allDamageType list one or more
	ConditionImmunities   []string       `json:"condition_immunities,omitempty"`   // ConditionImmunities to add from allDamageType list one or more
	CancelDisvantage      []string       `json:"cancel_disvantage,omitempty"`      // Removes one kind of Disvantages from your list
	CancelCondition       []string       `json:"cancel_condition,omitempty"`       // Removes one kind of Condition from your list
	EnemyDisvantages      []string       `json:"enemy_disvantages,omitempty"`      // list of Disvantages to your enemies
	Curse                 bool           `json:"curse,omitempty"`                  // add disvantage when activated
	OverrideDamageType    string         `json:"override_damage_type,omitempty"`   // used in weapons to cause only one type of damage instead weapons one
	Regeneration          int            `json:"regeneration,omitempty"`           // recovery HP per turn, after his own action
	ProficiencyBonus      int            `json:"proficiency_bonus,omitempty"`      // Increase Character proficiency bonus in all tests
	SpellImmunity         []string       `json:"spell_immunity,omitempty"`         // just a Spell Immunity list
}

// CorePowers struct
// this kind of magic should be trigger or used
type CorePowers struct {
	Purpose                   string   `json:"purpose,omitempty"`                     // in CoreDnDSystem which rule should be used
	Dice                      string   `json:"dice,omitempty"`                        // Dice to roll
	DamageType                string   `json:"damage_type,omitempty"`                 // type of damage from allDamageType
	DamageDice                string   `json:"damage_dice,omitempty"`                 // if have damage, use this dices
	DifficultClass            int      `json:"difficult_class,omitempty"`             // to enemy use
	SavingThrow               string   `json:"saving_throw,omitempty"`                // which ability to use to check
	SpellName                 string   `json:"spell_name,omitempty"`                  // spell name for use
	SpellList                 []string `json:"spell_list,omitempty"`                  // spell list name to be use as spell name
	SpellFreeList             []string `json:"spell_free_list,omitempty"`             // spell free list name to be use as spell name without charge cost
	SpellLevel                int      `json:"spell_level,omitempty"`                 // spell level to be used if apply for that magic
	Duration                  int      `json:"duration,omitempty"`                    // duration time
	AttackRoll                int      `json:"attack_roll,omitempty"`                 // if need to attack, use this value
	AttackNumber              int      `json:"attack_number,omitempty"`               // number of attacks possible
	ConditionMultiple         int      `json:"condition_multiple,omitempty"`          // Condition to trigger any SavingThrow
	ConditionHitPoints        int      `json:"condition_hit_points,omitempty"`        // Hit Points Condition to trigger any SavingThrow
	ArmorClassBonus           int      `json:"armor_class_bonus,omitempty"`           // AC
	EnemyAttackType           string   `json:"enemy_attack_type,omitempty"`           // Enemy Attack Type used to trigger any defense type
	ReduceDamageDice          string   `json:"reduce_damage_dice,omitempty"`          // If enemy attack type matches, reduce damage received
	ReduceDamageAbilitiy      string   `json:"reduce_damage_abality,omitempty"`       // which ability to use with reduce damage dice
	DamageResistance          []string `json:"damage_resistance,omitempty"`           // damageResistance to add from allDamageType list one or more
	Advantages                []string `json:"advantages,omitempty"`                  // add Advantages if enemy list match
	Disvantages               []string `json:"disvantages,omitempty"`                 // add Disvantages in enemy list temporary
	Condition                 []string `json:"condition,omitempty"`                   // add Condition in enemy list temporary
	CancelCondition           []string `json:"cancel_condition,omitempty"`            // Removes one kind of Condition from your list
	CombateMastery            int      `json:"combate_mastery,omitempty"`             // You can choose where to use your bonus, attack/damage or armor class
	Curse                     bool     `json:"curse,omitempty"`                       // add disvantage when activated
	ChargeType                bool     `json:"charge_type,omitempty"`                 // Any kind of item who needs to have charges, like: staff, wands
	Charges                   int      `json:"charges,omitempty"`                     // number of charges
	DiceCharges               string   `json:"dice_charges,omitempty"`                // random way to generate number of charges, used by hoard generator
	RecoveryDiceCharges       string   `json:"recovery_dice_charges,omitempty"`       // random way to recover charges / per time
	ZeroChargesRoll           bool     `json:"zero_charges_roll,omitempty"`           // if reaches 0 charges, roll a d20, on 1, destroyed
	WeaponPropertyRestriction string   `json:"weapon_property_restriction,omitempty"` // if used weapon doenst have that property, will not work
	DamageBonus               int      `json:"damage_bonus,omitempty"`                // add bonus to damage if power in use
	RequireEnemyType          []string `json:"require_enemy_type,omitempty"`          // used to activate a extra damage if enemy type matches
	CombateBonus              int      `json:"combate_bonus,omitempty"`               // on attack and damage
	DamageBonusCritical       int      `json:"damage_bonus_critical,omitempty"`       // add bonus to damage if was a critical
	ExtraHitPoints            int      `json:"extra_hit_points,omitempty"`            // extra Hit Points
}

// Scroll struct
type Scroll struct {
	Content               string `json:"content,omitempty"`
	DifficultClass        int    `json:"difficult_class,omitempty"`
	SavingsDifficultClass int    `json:"savings_difficult_class,omitempty"`
	Attack                int    `json:"attack,omitempty"`
	WizardDifficult       int    `json:"wizard_difficult,omitempty"`
}

//NewMagicItemPrint func
func NewMagicItemPrint() {
	DB, err := os.Open("./Magic-Items.json")
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
	var listMagicItems []MagicItem
	// fmt.Println("switch name {")
	for key, value := range result {
		// fmt.Println(key)
		// fmt.Printf("	case \"%s\":\n", key)
		switch key {
		case "Defender":
			// fmt.Println(key)
			desc, att := contentInterfaceParser(value)
			// fmt.Println(title, desc, att)
			for _, v := range anySword {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableI
				magic.Category = Weapons
				magic.Rarity = Legendary
				power := CorePowers{
					Purpose:        "attack",
					CombateMastery: 3,
				}
				magic.Power = &power
				listMagicItems = append(listMagicItems, magic)
			}
		case "Potion of Water Breathing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		// case "Staff of the Magi":
		// case "Armor of Invulnerability":
		case "Bead of Force":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Category = Potion
			magic.Rarity = Rare
			power := CorePowers{
				Purpose:        "attack-with-saving-damage",
				DifficultClass: 15,
				DamageDice:     "5d4",
				DamageType:     "force",
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		// case "Efreeti Bottle":
		// case "Portable Hole":
		case "Restorative Ointment":
			// Keoghtom's Ointment
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableB
			magic.Category = Potion
			power := CorePowers{
				Purpose:         Heal,
				Dice:            "2d8+2",
				DiceCharges:     "1d4+1",
				ChargeType:      true,
				CancelCondition: []string{"poisoned"},
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)

		// case "Vorpal Sword":
		case "Cape of the Mountebank":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Cloak of Displacement":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			feature := CoreFeatures{
				EnemyDisvantages: []string{"attack"},
			}
			magic.Feature = &feature
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Feather Falling":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = Rings
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Ring of Shooting Stars":
		// case "Staff of Fire":
		case "Gloves of Swimming and Climbing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			skill := map[string]int{"athletics": 5}
			feature := CoreFeatures{
				SkillBonus: skill,
			}
			magic.Feature = &feature
			listMagicItems = append(listMagicItems, magic)
		case "Necklace of Fireballs":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			power := CorePowers{
				Purpose:        "spell-charge-level",
				SpellName:      "fireball",
				DifficultClass: 15,
				ChargeType:     true,
				DiceCharges:    "1d6+3",
				SpellLevel:     3,
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		case "Dagger of Venom":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Shape = "dagger"
			feature := CoreFeatures{
				CombateBonus: 1,
			}
			magic.Feature = &feature
			power := CorePowers{
				Purpose:        "attack-with-saving-damage-condition",
				DamageDice:     "2d10",
				DamageType:     "poison",
				DifficultClass: 15,
				SavingThrow:    Constitution,
				Condition:      []string{"poisoned"},
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Jumping":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = Rings
			listMagicItems = append(listMagicItems, magic)
		// case "Talisman of Ultimate Evil":
		case "Wand of Fear":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Wands
			magic.Rarity = Rare
			power := CorePowers{
				Purpose:             "spell",
				SpellList:           []string{"command", "cone-of-fear"},
				DifficultClass:      15,
				SavingThrow:         Wisdom,
				ChargeType:          true,
				Charges:             7,
				RecoveryDiceCharges: "1d6+1",
				ZeroChargesRoll:     true,
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		// case "Animated Shield":
		case "Cloak of Protection":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{
				ArmorClassBonus: 1,
				SavingBonus:     1,
			}
			magic.Feature = &feature
			listMagicItems = append(listMagicItems, magic)

		case "Helm of Teleportation":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			power := CorePowers{
				Purpose:             "spell",
				SpellName:           "teleport",
				ChargeType:          true,
				Charges:             3,
				RecoveryDiceCharges: "1d3",
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		// case "Robe of Stars":
		case "Wand of Secrets":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			power := CorePowers{
				Purpose:             "roleplay",
				ChargeType:          true,
				Charges:             3,
				RecoveryDiceCharges: "1d3",
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)

		case "Armor of Resistance", "Ring of Resistance", "Potion of Resistance":
			desc, att := contentInterfaceParser(value)
			for _, v := range resistance {
				var name string
				if strings.HasPrefix(key, "Armor") {
					name = fmt.Sprintf("armor-of-resistance-%s", v)
					desc = "You have resistance to one type of damage while you wear this armor."
				}
				if strings.HasPrefix(key, "Ring") {
					name = fmt.Sprintf("ring-of-resistance-%s", v)
					desc = "You have resistance to one damage type while wearing this ring"
				}
				if strings.HasPrefix(key, "Potion") {
					name = fmt.Sprintf("potion-of-resistance-%s", v)
					desc = "When you drink this potion, you gain resistance to one type of damage for 1 hour."
				}
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				feature := CoreFeatures{
					DamageResistance: []string{v},
				}
				magic.Feature = &feature
				if strings.HasPrefix(key, "Armor") {
					magic.Category = ArmorCons
					magic.Rarity = Rare
					for _, a := range armorsList {
						newName := fmt.Sprintf("%s-%s", a, name)
						magic.Name = newName
						magic.Shape = a
						switch a {
						case "hide":
							magic.HoardTable = tableG

						case "chain-shirt":
							magic.HoardTable = tableG

						case "scale-mail":
							magic.HoardTable = tableG

						case "breastplate":
							magic.HoardTable = tableH

						case "half-plate":
							magic.HoardTable = tableI

						case "ring-mail":
							magic.HoardTable = tableG

						case "chain-mail":
							magic.HoardTable = tableG

						case "splint":
							magic.HoardTable = tableH

						case "plate":
							magic.HoardTable = tableI

						case "padded":
							magic.HoardTable = tableG

						case "leather":
							magic.HoardTable = tableG

						case "studded-leather":
							magic.HoardTable = tableH

						}
					}
				}
				if strings.HasPrefix(key, "Ring") {
					magic.Category = Rings
					magic.Rarity = Rare
					magic.HoardTable = tableG
				}
				if strings.HasPrefix(key, "Potion") {
					magic.Category = Potion
					magic.HoardTable = tableB
				}
				listMagicItems = append(listMagicItems, magic)
			}
		case "Gauntlets of Ogre Power":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			str := map[string]int{Strength: 19}
			feature := CoreFeatures{
				NewAbility: str,
			}
			magic.Feature = &feature
			listMagicItems = append(listMagicItems, magic)
		case "Robe of Eyes":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			feature := CoreFeatures{
				Advantages: []string{"perception"},
			}
			magic.Feature = &feature
		// case "Rod of Absorption":
		case "Shield, +1, +2, or +3":
			desc, att := contentInterfaceParser(value)
			for _, a := range bonusList {
				name := fmt.Sprintf("shield-%v", a)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Weapons
				magic.Shape = Shield
				feature := CoreFeatures{
					ArmorClassBonus: a,
				}
				magic.Feature = &feature
				switch a {
				case 1:
					magic.HoardTable = tableB
				case 2:
					magic.HoardTable = tableC
					magic.Rarity = Rare
				case 3:
					magic.HoardTable = tableD
					magic.Rarity = VeryRare
				}
				listMagicItems = append(listMagicItems, magic)
			}
		case "Flame Tongue":
			desc, att := contentInterfaceParser(value)
			for _, v := range anySword {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableG
				magic.Category = Weapons
				magic.Rarity = Rare
				magic.Shape = v
				power := CorePowers{
					Purpose:    "attack-damage-extra",
					DamageDice: "2d6",
					DamageType: "fire",
				}
				magic.Power = &power
				listMagicItems = append(listMagicItems, magic)
			}
		case "Glamoured Studded Leather":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = ArmorCons
			magic.Rarity = Rare
			magic.Shape = "studded-leather"
			feature := CoreFeatures{
				ArmorClassBonus: 1,
			}
			magic.Feature = &feature
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Clairvoyance":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		case "Ammunition, +1, +2, or +3":
			desc, att := contentInterfaceParser(value)
			for _, v := range ammunitionList {
				for _, a := range bonusList {
					name := fmt.Sprintf("ammunition-%s-%v", v, a)
					magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
					magic.Category = Weapons
					magic.Shape = v
					feature := CoreFeatures{
						AttackBonus: a,
					}
					magic.Feature = &feature
					switch a {
					case 1:
						magic.HoardTable = tableB
					case 2:
						magic.HoardTable = tableC
						magic.Rarity = Rare
					case 3:
						magic.HoardTable = tableD
						magic.Rarity = VeryRare
					}
					listMagicItems = append(listMagicItems, magic)
				}
			}

		case "Bracers of Archery":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			power := CorePowers{
				Purpose:                   "attack-conditional-weapon",
				DamageBonus:               2,
				WeaponPropertyRestriction: "ranged",
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		case "Trident of Fish Command":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Weapons
			magic.Shape = "trident"
			power := CorePowers{
				Purpose:             "spell",
				SpellName:           "dominate-beast",
				ChargeType:          true,
				Charges:             3,
				RecoveryDiceCharges: "1d3",
			}
			magic.Power = &power
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Healing":
			desc, att := contentInterfaceParser(value)
			healingNames := map[string]string{"healing": "2d4+2", "greater-healing": "4d4+4", "superior-healing": "8d4+8", "supreme-healing": "10d4+20"}
			for k, v := range healingNames {
				name := fmt.Sprintf("potion-of-%s", k)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Potion
				power := CorePowers{}
				magic.Power = &power
				magic.Power.Purpose = Heal
				magic.Power.Dice = v
				switch k {
				case "healing":
					magic.HoardTable = tableA
					magic.Rarity = "common"
				case "greater-healing":
					magic.HoardTable = []string{"A", "B"}
					magic.Rarity = "uncommon"
				case "superior-healing":
					magic.HoardTable = tableC
					magic.Rarity = "rare"
				case "supreme-healing":
					magic.HoardTable = []string{"D", "E"}
					magic.Rarity = "very rare"
				}
				listMagicItems = append(listMagicItems, magic)
			}

		// case "Potion of Speed":
		case "Potion of Invisibility":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableD
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		// case "Arrow of Slaying":
		// case "Goggles of Night":
		// case "Talisman of the Sphere":
		case "Giant Slayer":
			desc, att := contentInterfaceParser(value)
			localList := append(anySword, "battleaxe", "greataxe", "handaxe")
			for _, v := range localList {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableG
				magic.Category = Weapons
				magic.Rarity = Rare
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				power := CorePowers{}
				magic.Power = &power
				magic.Feature.CombateBonus = 1
				magic.Power.Purpose = "attack-conditional-enemy"
				magic.Power.RequireEnemyType = []string{"giants"}
				magic.Power.DamageDice = "2d6"
				magic.Power.DifficultClass = 15
				magic.Power.SavingThrow = Strength
				magic.Power.Condition = []string{"prone"}
				listMagicItems = append(listMagicItems, magic)
			}
		// case "Talisman of Pure Good":
		case "Oil of Etherealness":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableD
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		// case "Robe of Scintillating Colors":
		case "Censer of Controlling Air Elementals":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Frost Brand":
		case "Periapt of Proof against Poison":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.ConditionImmunities = []string{"poisoned"}
			magic.Feature.DamageImmunities = []string{"poison"}
			listMagicItems = append(listMagicItems, magic)
		// case "Apparatus of the Crab Levers":
		// case "Bag of Devouring":
		// case "Staff of Striking":
		case "Winged Boots":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Eyes of the Eagle":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"perception"}
			listMagicItems = append(listMagicItems, magic)
		// case "Holy Avenger":
		// case "Sword of Sharpness":
		case "Gem of Brightness":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-with-saving-condition"
			magic.Power.Condition = []string{"blinded"}
			magic.Power.Charges = 50
			magic.Power.ChargeType = true
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Constitution
			listMagicItems = append(listMagicItems, magic)
		case "Necklace of Prayer Beads":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-without-spellcasting"
			magic.Power.DiceCharges = "1d4+2"
			magic.Power.ChargeType = true
			magic.Power.SpellLevel = 2
			magic.Power.SpellList = []string{"bless", "cure-wounds", "lesser-restoration", "greater-restoration", "branding-smite", "planar-ally", "wind-walk"}
			magic.AttunementRestriction = []string{"cleric", "druid", "paladin"}
			listMagicItems = append(listMagicItems, magic)
		case "Periapt of Health":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.ConditionImmunities = []string{"disease"}
		case "Dust of Sneezing and Choking":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-with-saving-condition"
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Constitution
			magic.Power.Condition = []string{"incapacitated"}
		// case "Oathbow":
		// case "Staff of Thunder and Lightning":
		// case "Tome of Understanding":
		case "Bag of Holding":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.Category = WoundrousItems
			magic.HoardTable = []string{"A", "B"}
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Mind Shielding":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.Category = Rings
			magic.HoardTable = tableF
			listMagicItems = append(listMagicItems, magic)
		// case "Horseshoes of a Zephyr":
		// case "Nine Lives Stealer":
		case "Wand of the War Mage, +1, +2, or +3":
			desc, att := contentInterfaceParser(value)
			for _, v := range bonusList {
				name := fmt.Sprintf("%s %v", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
				magic.Category = Wands
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.SpellAttackBonus = v
				magic.AttunementRestriction = spellcaster
				switch v {
				case 1:
					magic.HoardTable = tableF
				case 2:
					magic.HoardTable = tableG
					magic.Rarity = Rare
				case 3:
					magic.HoardTable = tableH
					magic.Rarity = VeryRare
				}
				listMagicItems = append(listMagicItems, magic)
			}

		case "Belt of Giant Strength":
			desc, att := contentInterfaceParser(value)
			for k, v := range giantNames {
				name := fmt.Sprintf("belt-of-giant-strength-%s", k)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Potion
				feature := CoreFeatures{}
				magic.Feature = &feature
				ability := map[string]int{Strength: v}
				magic.Feature.NewAbility = ability
				switch k {
				case "hill":
					magic.Rarity = Rare
					magic.HoardTable = tableG
				case "stone":
					magic.Rarity = VeryRare
					magic.HoardTable = tableH
				case "frost":
					magic.Rarity = VeryRare
					magic.HoardTable = tableH
				case "fire":
					magic.Rarity = VeryRare
					magic.HoardTable = tableH
				case "cloud":
					magic.Rarity = Legendary
					magic.HoardTable = tableI
				case "storm":
					magic.Rarity = Legendary
					magic.HoardTable = tableI
				}
				listMagicItems = append(listMagicItems, magic)
			}
		case "Boots of Striding and Springing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Stone of Good Luck (Luckstone)":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName("Stone of Good Luck"), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.SavingBonus = 1
			magic.Feature.AbilityBonus = 1
			listMagicItems = append(listMagicItems, magic)
		case "Cloak of the Manta Ray":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Stone of Controlling Earth Elementals":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Helm of Brilliance":
		case "Horn of Blasting":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-with-saving-damage-condition"
			magic.Power.DamageDice = "5d6"
			magic.Power.DamageType = "thunder"
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Constitution
			magic.Power.Condition = []string{"deafened"}
			listMagicItems = append(listMagicItems, magic)
		// case "Well of Many Worlds":
		case "Broom of Flying":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Dwarven Thrower":
		case "Staff of the Python":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Shape = "quarterstaff"
			magic.AttunementRestriction = []string{"cleric", "druid", "warlock"}
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Ring of Elemental Command":
		case "Vicious Weapon":
			desc, att := contentInterfaceParser(value)
			for _, v := range weaponsList {
				name := fmt.Sprintf("Vicious %s", v)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Weapons
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.DamageBonusCritical = 7
				magic.HoardTable = tableG
				magic.Rarity = Rare
				listMagicItems = append(listMagicItems, magic)
			}
		case "Dust of Dryness":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Oil of Slipperiness":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Water Walking":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = Rings
			listMagicItems = append(listMagicItems, magic)
		// case "Carpet of Flying":
		case "Ring of the Ram":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Rings
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Charges = 3
			magic.Power.ChargeType = true
			magic.Power.DiceCharges = "1d3"
			magic.Power.Purpose = "attack"
			magic.Power.AttackRoll = 7
			magic.Power.DamageDice = "2d10"
			magic.Power.DamageType = "force"
			listMagicItems = append(listMagicItems, magic)
		case "Robe of Useful Items":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Rod of Lordly Might":
		case "Berserker Axe":
			desc, att := contentInterfaceParser(value)
			list := []string{"battleaxe", "greataxe", "handaxe"}
			for _, v := range list {
				name := fmt.Sprintf("Berserker %s", v)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Weapons
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.CombateBonus = 1
				magic.Feature.ExtraHitPointsLevel = 1
				magic.Feature.Curse = true
				magic.HoardTable = tableG
				magic.Rarity = Rare
				listMagicItems = append(listMagicItems, magic)
			}
		// case "Manual of Quickness of Action":
		case "Ring of X-ray Vision":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = Rings
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Arrow-Catching Shield":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = ArmorCons
			magic.Rarity = Rare
			magic.Shape = "shield"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "defense"
			magic.Power.EnemyAttackType = "ranged"
			magic.Power.ArmorClassBonus = 2
			listMagicItems = append(listMagicItems, magic)
		case "Deck of Illusions":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Mirror of Life Trapping":
		case "Staff of the Woodlands":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Shape = "quarterstaff"
			magic.Rarity = Rare
			magic.AttunementRestriction = []string{"druid"}
			feature := CoreFeatures{}
			magic.Feature = &feature
			power := CorePowers{}
			magic.Power = &power
			magic.Feature.CombateBonus = 2
			magic.Feature.SpellAttackBonus = 2
			magic.Power.Purpose = "spell-without-spellcasting"
			magic.Power.ChargeType = true
			magic.Power.Charges = 10
			magic.Power.DiceCharges = "1d6+2"
			magic.Power.SpellList = []string{"animal-friendship", "awaken", "barkskin", "locate-animals-or-plants", "speak-with-animals", "speak-with-plants", "wall-of-thorns"}
			magic.Power.SpellFreeList = []string{"pass-without-trace"}
			listMagicItems = append(listMagicItems, magic)
		case "Pipes of the Sewers":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Deck of Many Things":
		case "Horseshoes of Speed":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Manual of Golems":
		case "Potion of Giant Strength":
			desc, att := contentInterfaceParser(value)
			for k, v := range giantNames {
				name := fmt.Sprintf("potion-of-giant-strength-%s", k)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = Potion
				feature := CoreFeatures{}
				magic.Feature = &feature
				ability := map[string]int{Strength: v}
				magic.Feature.NewAbility = ability
				switch k {
				case "hill":
					magic.HoardTable = tableB
				case "stone":
					magic.Rarity = Rare
					magic.HoardTable = tableC
				case "frost":
					magic.Rarity = Rare
					magic.HoardTable = tableC
				case "fire":
					magic.Rarity = Rare
					magic.HoardTable = tableC
				case "cloud":
					magic.Rarity = VeryRare
					magic.HoardTable = tableD
				case "storm":
					magic.Rarity = Legendary
					magic.HoardTable = tableE
				}
				listMagicItems = append(listMagicItems, magic)
			}

		case "Wand of Lightning Bolts":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Wands
			magic.Rarity = Rare
			magic.AttunementRestriction = spellcaster
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-charge-level"
			magic.Power.SpellName = "lightning-bolt"
			magic.Power.SpellLevel = 3
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Dexterity
			magic.Power.ChargeType = true
			magic.Power.Charges = 7
			magic.Power.RecoveryDiceCharges = "1d6+1"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Wand of Web":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = Wands
			magic.AttunementRestriction = spellcaster
			listMagicItems = append(listMagicItems, magic)
		case "Boots of Elvenkind":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"stealth"}
			listMagicItems = append(listMagicItems, magic)
		case "Eyes of Charming":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Luck Blade":
		// case "Spellguard Shield":
		// case "Tome of Leadership and Influence":
		case "Wand of Wonder":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = Wands
			magic.AttunementRestriction = spellcaster
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Dust of Disappearance":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Ioun Stone":
			list := []string{"Absorption", "Agility", "Awareness", "Fortitude", "Greater Absorption", "Insight", "Intellect", "Leadership", "Mastery", "Protection", "Regeneration", "Reserve", "Strength", "Sustenance"}
			desc, att := contentInterfaceParser(value)
			for _, v := range list {
				name := fmt.Sprintf("%s %s", key, v)
				switch v {
				case "Absorption":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					listMagicItems = append(listMagicItems, magic)

				case "Agility":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Dexterity: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Awareness":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
					magic.HoardTable = tableG
					magic.Category = WoundrousItems
					magic.Rarity = Rare
					listMagicItems = append(listMagicItems, magic)

				case "Fortitude":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Constitution: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Greater Absorption":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
					magic.HoardTable = tableI
					magic.Category = WoundrousItems
					magic.Rarity = Legendary
					listMagicItems = append(listMagicItems, magic)

				case "Insight":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Wisdom: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Intellect":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Intelligence: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Leadership":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Charisma: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Mastery":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableI
					magic.Category = WoundrousItems
					magic.Rarity = Legendary
					feature := CoreFeatures{}
					magic.Feature = &feature
					magic.Feature.ProficiencyBonus = 1
					listMagicItems = append(listMagicItems, magic)

				case "Protection":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableG
					magic.Category = WoundrousItems
					magic.Rarity = Rare
					feature := CoreFeatures{}
					magic.Feature = &feature
					magic.Feature.ArmorClassBonus = 1
					listMagicItems = append(listMagicItems, magic)

				case "Regeneration":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableI
					magic.Category = WoundrousItems
					magic.Rarity = Legendary
					feature := CoreFeatures{}
					magic.Feature = &feature
					magic.Feature.Regeneration = 15
					listMagicItems = append(listMagicItems, magic)

				case "Reserve":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
					magic.HoardTable = tableG
					magic.Category = WoundrousItems
					magic.Rarity = Rare
					listMagicItems = append(listMagicItems, magic)

				case "Strength":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					magic.HoardTable = tableH
					magic.Category = WoundrousItems
					magic.Rarity = VeryRare
					feature := CoreFeatures{}
					magic.Feature = &feature
					ability := map[string]int{Strength: 2}
					magic.Feature.IncreaseAbility = ability
					listMagicItems = append(listMagicItems, magic)

				case "Sustenance":
					magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
					magic.HoardTable = tableG
					magic.Category = WoundrousItems
					magic.Rarity = Rare
					listMagicItems = append(listMagicItems, magic)

				}
			}
		case "Decanter of Endless Water":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Javelin of Lightning":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Weapons
			magic.Shape = "javelin"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-with-saving-ranged-attack-damage-extra"
			magic.Power.DifficultClass = 13
			magic.Power.SavingThrow = Dexterity
			magic.Power.DamageDice = "4d6"
			magic.Power.DamageType = "lightning"
			listMagicItems = append(listMagicItems, magic)

		// case "Sphere of Annihilation":
		case "Staff of Withering":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.AttunementRestriction = []string{"cleric", "druid", "warlock"}
			magic.Shape = "quarterstaff"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-with-saving-damage-condition"
			magic.Power.ChargeType = true
			magic.Power.Charges = 3
			magic.Power.DiceCharges = "1d3"
			magic.Power.DamageDice = "2d10"
			magic.Power.DamageType = "necrotic"
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Constitution
			magic.Power.Disvantages = []string{Strength, Constitution}
			listMagicItems = append(listMagicItems, magic)

		case "Wand of Magic Detection":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Apparatus of the Crab":
		case "Brazier of Commanding Fire Elementals":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Sun Blade":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.Shape = "longsword"
			feature := CoreFeatures{}
			magic.Feature = &feature
			power := CorePowers{}
			magic.Power = &power
			magic.Feature.CombateBonus = 2
			magic.Feature.OverrideDamageType = "radiant"
			magic.Power.Purpose = "attack-conditional-enemy"
			magic.Power.DamageDice = "1d8"
			magic.Power.DamageType = "radiant"
			magic.Power.RequireEnemyType = []string{"undead"}
			listMagicItems = append(listMagicItems, magic)
		// case "Dragon Scale Mail":
		case "Rod of Rulership":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Animal Influence":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Poison":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = Potion
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "poison"
			magic.Power.DamageDice = "3d6"
			magic.Power.DamageType = "poison"
			magic.Power.DifficultClass = 13
			magic.Power.SavingThrow = Constitution
			magic.Power.Condition = []string{"poisoned"}
			listMagicItems = append(listMagicItems, magic)
		// case "Ring of Three Wishes":
		case "Mace of Terror":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Shape = "mace"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-with-saving-condition"
			magic.Power.Condition = []string{"frightened"}
			magic.Power.ChargeType = true
			magic.Power.Charges = 3
			magic.Power.DiceCharges = "1d3"
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Wisdom
			listMagicItems = append(listMagicItems, magic)
		// case "Marvelous Pigments":
		case "Potion of Climbing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.Rarity = Common
			magic.Category = Potion
			magic.HoardTable = tableA
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"athletics"}
			listMagicItems = append(listMagicItems, magic)
		// case "Ring of Regeneration":
		case "Sword of Wounding":
			desc, att := contentInterfaceParser(value)
			for _, v := range anySword {
				name := fmt.Sprintf("%s of Wounding", v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableG
				magic.Category = Weapons
				magic.Shape = v
				power := CorePowers{}
				magic.Power = &power
				magic.Power.Purpose = "attack-with-savings-damage-continuous"
				magic.Power.DamageDice = "1d4"
				magic.Power.DamageType = "necrotic"
				magic.Power.DifficultClass = 15
				magic.Power.SavingThrow = Constitution
				listMagicItems = append(listMagicItems, magic)
			}
		case "Lantern of Revealing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Medallion of Thoughts":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Elemental Gem":
			desc, att := contentInterfaceParser(value)
			list := []string{"Blue Sapphire", "Emerald", "Red Corundum", "Yellow Diamond"}
			for _, v := range list {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
				magic.HoardTable = tableB
				magic.Category = WoundrousItems
				listMagicItems = append(listMagicItems, magic)
			}

		// case "Iron Flask":
		case "Ring of Swimming":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = Rings
			listMagicItems = append(listMagicItems, magic)
		case "Wand of Binding":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Wands
			magic.Rarity = Rare
			magic.AttunementRestriction = spellcaster
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell"
			magic.Power.SpellList = []string{"hold-monster", "hold-person"}
			magic.Power.DifficultClass = 17
			magic.Power.SavingThrow = Wisdom
			magic.Power.ChargeType = true
			magic.Power.Charges = 7
			magic.Power.RecoveryDiceCharges = "1d6+1"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Wings of Flying":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Mantle of Spell Resistance":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"spell"}
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Protection":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Rings
			magic.Rarity = Rare
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.ArmorClassBonus = 1
			magic.Feature.SavingBonus = 1
			listMagicItems = append(listMagicItems, magic)
		case "Necklace of Adaptation":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Amulet of Health":
		case "Mithral Armor":
			desc, att := contentInterfaceParser(value)
			for _, v := range armorsMetalList {
				name := fmt.Sprintf("Mithral %s", v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableB
				magic.Category = ArmorCons
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.CancelDisvantage = []string{"stealth"}
				listMagicItems = append(listMagicItems, magic)
			}

		// case "Wand of Polymorph":
		case "Circlet of Blasting":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = Spell
			magic.Power.SpellName = "scorching-ray"
			magic.Power.AttackRoll = 5
			// magic.Power.DamageDice = "2d6"
			// magic.Power.DamageType = "fire"
			magic.HoardTable = tableC
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Diminution":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.Category = Potion
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = Spell
			magic.Power.SpellName = "enlarge/reduce"
			magic.HoardTable = tableC
			listMagicItems = append(listMagicItems, magic)
		case "Boots of Levitation":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Tome of Clear Thought":
		case "Gloves of Missile Snaring":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "defense-reduce-damage-roll"
			magic.Power.EnemyAttackType = "ranged"
			magic.Power.ReduceDamageDice = "1d10"
			magic.Power.ReduceDamageAbilitiy = Dexterity
			listMagicItems = append(listMagicItems, magic)
		case "Mace of Disruption":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Shape = "mace"
			magic.Rarity = Rare
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "attack-conditional-enemy-damage-condition"
			magic.Power.RequireEnemyType = []string{"fiend", "undead"}
			magic.Power.DamageDice = "2d6"
			magic.Power.DamageType = "radiant"
			magic.Power.ConditionHitPoints = 25
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Wisdom
			magic.Power.Condition = []string{"frightened"}
			listMagicItems = append(listMagicItems, magic)
		// case "Manual of Bodily Health":
		// case "Universal Solvent":
		case "Wind Fan":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Brooch of Shielding":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.DamageResistance = []string{"force"}
			magic.Feature.SpellImmunity = []string{"magic-missile"}
			listMagicItems = append(listMagicItems, magic)
		case "Eyes of Minute Seeing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"investigation"}
			listMagicItems = append(listMagicItems, magic)
		case "Boots of the Winterlands":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.DamageResistance = []string{"cold"}
			listMagicItems = append(listMagicItems, magic)
		case "Armor of Vulnerability":
			desc, att := contentInterfaceParser(value)
			for _, v := range damageMeleeType {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableG
				magic.Category = ArmorCons
				magic.Shape = "plate"
				magic.Rarity = Rare
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.Curse = true
				magic.Feature.DamageResistance = []string{v}
				listMagicItems = append(listMagicItems, magic)
			}
		case "Bag of Beans":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Rod of Alertness":
		// case "Scarab of Protection":
		case "Staff of Swarming Insects":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.AttunementRestriction = []string{"cleric", "druid", "warlock", "bard", "sorcerer", "warlock", "wizard"}
			magic.Shape = "quarterstaff"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell"
			magic.Power.ChargeType = true
			magic.Power.Charges = 10
			magic.Power.SpellList = []string{"giant-insect", "insect-plague"}
			magic.Power.DiceCharges = "1d6+4"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Hat of Disguise":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Scimitar of Speed":
		case "Staff of Charming":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.AttunementRestriction = []string{"cleric", "druid", "warlock", "bard", "sorcerer", "warlock", "wizard"}
			magic.Shape = "quarterstaff"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell"
			magic.Power.ChargeType = true
			magic.Power.Charges = 10
			magic.Power.SpellList = []string{"charm-person", "command", "comprehend-languages"}
			magic.Power.DiceCharges = "1d8+2"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		// case "Artifacts":
		case "Potion of Animal Friendship":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.Category = Potion
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = Spell
			// missing spell
			magic.Power.SpellName = "animal-friendship"
			magic.HoardTable = tableB
			listMagicItems = append(listMagicItems, magic)
		case "Rope of Entanglement":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Crystal Ball":
		case "Eversmoking Bottle":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Handy Haversack":
			// Heward's Handy Haversack
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Immovable Rod":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Mind Reading":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Boots of Speed":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Cloak of Arachnida":
		case "Pearl of Power":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Pipes of Haunting":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Dimensional Shackles":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Elven Chain":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = ArmorCons
			magic.Rarity = Rare
			magic.Shape = "chain-shirt"
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.WithProficiency = []string{"medium-armor"}
			listMagicItems = append(listMagicItems, magic)
		case "Cloak of the Bat":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"stealth"}
			listMagicItems = append(listMagicItems, magic)
		case "Feather Token":
			// Quaal's Feather Token
			desc, att := contentInterfaceParser(value)
			list := []string{"Anchor", "Bird", "Fan", "Boat", "Tree", "Whip"}
			for _, v := range list {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
				if v == "Whip" {
					magic = bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
					power := CorePowers{}
					magic.Power = &power
					magic.Power.Purpose = "attack"
					magic.Power.AttackRoll = 9
					magic.Power.DamageDice = "1d6+5"
					magic.Power.DamageType = "force"
				}
				magic.HoardTable = tableC
				magic.Rarity = Rare
				magic.Category = WoundrousItems
				listMagicItems = append(listMagicItems, magic)
			}

		case "Folding Boat":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Oil of Sharpness":
		case "Potion of Growth":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.Category = Potion
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = Spell
			magic.Power.SpellName = "enlarge/reduce"
			magic.HoardTable = tableB
			listMagicItems = append(listMagicItems, magic)
		// case "Rope of Climbing":
		case "Amulet of Proof against Detection and Location":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Bag of Tricks":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Staff of Frost":
		case "Helm of Comprehending Languages":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Weapon, +1, +2, or +3":
			desc, att := contentInterfaceParser(value)
			for _, v := range weaponsList {
				for _, a := range bonusList {
					name := fmt.Sprintf("%s-%v", v, a)
					magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
					magic.Category = Weapons
					magic.Shape = v
					feature := CoreFeatures{}
					magic.Feature = &feature
					magic.Feature.CombateBonus = a
					switch a {
					case 1:
						magic.HoardTable = tableF
					case 2:
						magic.HoardTable = tableG
						magic.Rarity = Rare
					case 3:
						magic.HoardTable = tableD
						magic.Rarity = VeryRare
					}
					listMagicItems = append(listMagicItems, magic)
				}
			}

		case "Belt of Dwarvenkind":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			power := CorePowers{}
			magic.Power = &power
			ability := map[string]int{Constitution: 2}
			magic.Feature.IncreaseAbility = ability
			magic.Feature.DamageResistance = []string{"poison"}
			magic.Feature.Advantages = []string{"saving-poison"}
			magic.Power.Purpose = "advantage-conditional"
			magic.Power.RequireEnemyType = []string{"dwarves"}
			magic.Power.Advantages = []string{"persuation"}
			listMagicItems = append(listMagicItems, magic)
		case "Cloak of Elvenkind":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.Advantages = []string{"stealth"}
			listMagicItems = append(listMagicItems, magic)
		case "Headband of Intellect":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			ability := map[string]int{Intelligence: 2}
			magic.Feature.NewAbility = ability
			listMagicItems = append(listMagicItems, magic)
		// case "Ring of Spell Turning":
		// case "Rod of Security":
		// case "Sentient Magic Items":
		case "Adamantine Armor":
			desc, att := contentInterfaceParser(value)
			for _, v := range armorsMetalList {
				name := fmt.Sprintf("Adamantine %s", v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.Category = ArmorCons
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				magic.Feature.DamageImmunities = []string{"criticalhit"}
				switch v {
				case "chain-shirt", "chain-mail", "scale-mail", "ring-mail":
					magic.HoardTable = tableF
				case "breastplate", "splint":
					magic.HoardTable = tableG
				case "half-plate", "plate":
					magic.HoardTable = tableH
				}
				listMagicItems = append(listMagicItems, magic)
			}
		// case "Hammer of Thunderbolts":
		// case "Dancing Sword":
		case "Instant Fortress":
			// 	Daern's Instant Fortress
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Evasion":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		// case "Robe of the Archmagi":
		case "Staff of Healing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.AttunementRestriction = []string{"cleric", "druid", "bard"}
			magic.Shape = "quarterstaff"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-charge-level-heal"
			magic.Power.ChargeType = true
			magic.Power.Charges = 10
			magic.Power.SpellList = []string{"cure-wounds", "lesser-restoratio", "mass-cure-wounds"}
			magic.Power.DiceCharges = "1d6+4"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Armor, +1, +2, or +3":
			desc, att := contentInterfaceParser(value)
			for _, v := range armorsList {
				for _, a := range bonusList {
					name := fmt.Sprintf("%s-%v", v, a)
					magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
					magic.Category = ArmorCons
					magic.Shape = v
					feature := CoreFeatures{}
					magic.Feature = &feature
					magic.Feature.ArmorClassBonus = a
					switch a {
					case 1:
						magic.HoardTable = tableF
					case 2:
						magic.HoardTable = tableG
						magic.Rarity = Rare
					case 3:
						magic.HoardTable = tableD
						magic.Rarity = VeryRare
					}
					listMagicItems = append(listMagicItems, magic)
				}
			}
		// case "Cube of Force Faces":
		case "Efficient Quiver":
			// Quiver of Ehlonna
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Potion of Gaseous Form":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableC
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		case "Shield of Missile Attraction":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = ArmorCons
			magic.Rarity = Rare
			magic.Shape = "shield"
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "defense"
			magic.Power.EnemyAttackType = "ranged"
			magic.Power.DamageResistance = []string{"piercing", "bludgeoning"}
			magic.Power.Curse = true
			listMagicItems = append(listMagicItems, magic)
		case "Cube of Force":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Demon Armor":
		// case "Ring of Djinni Summoning":
		case "Wand of Enemy Detection":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = Wands
			listMagicItems = append(listMagicItems, magic)
		case "Bowl of Commanding Water Elementals":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Category = WoundrousItems
			magic.Rarity = Rare
			listMagicItems = append(listMagicItems, magic)
		// case "Manual of Gainful Exercise":
		// case "Ring of Invisibility":
		// case "Ring of Telekinesis":
		case "Wand of Fireballs":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Wands
			magic.Rarity = Rare
			magic.AttunementRestriction = spellcaster
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-charge-level"
			magic.Power.SpellName = "fireball"
			magic.Power.SpellLevel = 3
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Dexterity
			magic.Power.ChargeType = true
			magic.Power.Charges = 7
			magic.Power.RecoveryDiceCharges = "1d6+1"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		// case "Candle of Invocation":
		case "Dragon Slayer":
			desc, att := contentInterfaceParser(value)
			for _, v := range anySword {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableF
				magic.Category = Weapons
				magic.Shape = v
				feature := CoreFeatures{}
				magic.Feature = &feature
				power := CorePowers{}
				magic.Power = &power
				magic.Feature.CombateBonus = 1
				magic.Power.Purpose = "attack-conditional-enemy"
				magic.Power.DamageDice = "3d6"
				magic.Power.RequireEnemyType = []string{"dragon"}
				listMagicItems = append(listMagicItems, magic)
			}

		case "Ring of Warmth":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Rings
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.DamageResistance = []string{"cold"}
			listMagicItems = append(listMagicItems, magic)
		case "Slippers of Spider Climbing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Horn of Valhalla":
			list := []string{"Brass", "Bronze", "Iron", "Silver"}
			desc, att := contentInterfaceParser(value)
			for _, v := range list {
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
				magic.Category = WoundrousItems
				switch v {
				case "Brass":
					magic.Rarity = Rare
					magic.HoardTable = tableG
				case "Bronze":
					magic.Rarity = VeryRare
					magic.HoardTable = tableH
				case "Iron":
					magic.Rarity = Legendary
					magic.HoardTable = tableI
				case "Silver":
					magic.Rarity = Rare
					magic.HoardTable = tableG
				}
				listMagicItems = append(listMagicItems, magic)
			}
		// case "Iron Bands of Binding":
		// case "Sovereign Glue":
		case "Gem of Seeing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Periapt of Wound Closure":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Wand of Paralysis":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Wands
			magic.Rarity = Rare
			magic.AttunementRestriction = spellcaster
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-saving-condition"
			magic.Power.DifficultClass = 15
			magic.Power.SavingThrow = Constitution
			magic.Power.Condition = []string{"paralyzed"}
			magic.Power.ChargeType = true
			magic.Power.Charges = 7
			magic.Power.RecoveryDiceCharges = "1d6+1"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Helm of Telepathy":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableF
			magic.Category = WoundrousItems
			listMagicItems = append(listMagicItems, magic)
		case "Philter of Love":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableB
			magic.Category = Potion
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Free Action":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = Rings
			listMagicItems = append(listMagicItems, magic)
		case "Ring of Spell Storing":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, true, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = Rings
			listMagicItems = append(listMagicItems, magic)
		case "Spell Scroll":
			desc, att := contentInterfaceParser(value)
			list := []string{"cantrip", "1st", "2nd", "3th", "4th", "5th", "6th", "7th", "8th", "9th"}
			for _, v := range list {
				name := fmt.Sprintf("spell-scroll-%s", v)
				magic := bacisDetailsMagicItem(name, key, desc, att, false, false)
				magic.Category = "scroll"
				scroll := Scroll{}
				magic.Scroll = &scroll
				switch v {
				case "cantrip":
					magic.Category = "scroll"
					magic.Rarity = Common
					magic.HoardTable = tableA
					magic.Scroll.Attack = 5
					magic.Scroll.DifficultClass = 10
					magic.Scroll.SavingsDifficultClass = 13
				case "1st":
					magic.Category = "scroll"
					magic.Rarity = Common
					magic.HoardTable = tableA
					magic.Scroll.Attack = 5
					magic.Scroll.DifficultClass = 11
					magic.Scroll.SavingsDifficultClass = 13
					magic.Scroll.WizardDifficult = 11
				case "2nd":
					magic.Category = "scroll"
					magic.HoardTable = []string{"A", "B"}
					magic.Scroll.Attack = 5
					magic.Scroll.DifficultClass = 12
					magic.Scroll.SavingsDifficultClass = 13
					magic.Scroll.WizardDifficult = 12
				case "3th":
					magic.Category = "scroll"
					magic.HoardTable = tableB
					magic.Scroll.Attack = 7
					magic.Scroll.DifficultClass = 13
					magic.Scroll.SavingsDifficultClass = 15
					magic.Scroll.WizardDifficult = 13
				case "4th":
					magic.Category = "scroll"
					magic.Rarity = Rare
					magic.HoardTable = tableC
					magic.Scroll.Attack = 7
					magic.Scroll.DifficultClass = 14
					magic.Scroll.SavingsDifficultClass = 15
					magic.Scroll.WizardDifficult = 14
				case "5th":
					magic.Category = "scroll"
					magic.Rarity = Rare
					magic.HoardTable = tableC
					magic.Scroll.Attack = 9
					magic.Scroll.DifficultClass = 15
					magic.Scroll.SavingsDifficultClass = 17
					magic.Scroll.WizardDifficult = 15
				case "6th":
					magic.Category = "scroll"
					magic.Rarity = VeryRare
					magic.HoardTable = tableD
					magic.Scroll.Attack = 9
					magic.Scroll.DifficultClass = 16
					magic.Scroll.SavingsDifficultClass = 17
					magic.Scroll.WizardDifficult = 16
				case "7th":
					magic.Category = "scroll"
					magic.Rarity = VeryRare
					magic.HoardTable = tableD
					magic.Scroll.Attack = 10
					magic.Scroll.DifficultClass = 17
					magic.Scroll.SavingsDifficultClass = 18
					magic.Scroll.WizardDifficult = 17
				case "8th":
					magic.Category = "scroll"
					magic.Rarity = VeryRare
					magic.HoardTable = []string{"D", "E"}
					magic.Scroll.Attack = 10
					magic.Scroll.DifficultClass = 18
					magic.Scroll.SavingsDifficultClass = 18
					magic.Scroll.WizardDifficult = 18
				case "9th":
					magic.Category = "scroll"
					magic.Rarity = Legendary
					magic.HoardTable = tableE
					magic.Scroll.Attack = 11
					magic.Scroll.DifficultClass = 19
					magic.Scroll.SavingsDifficultClass = 19
					magic.Scroll.WizardDifficult = 19
				}
				listMagicItems = append(listMagicItems, magic)
			}

		// case "Cubic Gate":
		// case "Plate Armor of Etherealness":
		// case "Amulet of the Planes":
		case "Sword of Life Stealing":
			desc, att := contentInterfaceParser(value)
			for _, v := range anySword {
				name := fmt.Sprintf("%s of Life Stealing", v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, false, false)
				magic.HoardTable = tableF
				magic.Category = Weapons
				magic.Shape = v
				power := CorePowers{}
				magic.Power = &power
				magic.Power.Purpose = "attack-critical-damage-extra"
				magic.Power.DamageBonus = 10
				magic.Power.DamageType = "necrotic"
				magic.Power.ExtraHitPoints = 10
				listMagicItems = append(listMagicItems, magic)
			}
		case "Potion of Heroism":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Category = Potion
			magic.Rarity = Rare
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "potion"
			magic.Power.ExtraHitPoints = 10
			magic.Power.SpellName = "bless"
			listMagicItems = append(listMagicItems, magic)
		case "Bracers of Defense":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			feature := CoreFeatures{}
			magic.Feature = &feature
			magic.Feature.ArmorClassBonus = 2
		// case "Dwarven Plate":
		// case "Staff of Power":
		case "Chime of Opening":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableC
			magic.Rarity = Rare
			magic.Category = WoundrousItems
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "roleplay"
			magic.Power.ChargeType = true
			magic.Power.Charges = 10
			listMagicItems = append(listMagicItems, magic)
		// case "Potion of Flying":
		case "Wand of Magic Missiles":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableF
			magic.Category = Wands
			power := CorePowers{}
			magic.Power = &power
			magic.Power.Purpose = "spell-charge-level"
			magic.Power.SpellName = "magic-missile"
			magic.Power.SpellLevel = 1
			magic.Power.ChargeType = true
			magic.Power.Charges = 7
			magic.Power.RecoveryDiceCharges = "1d6+1"
			magic.Power.ZeroChargesRoll = true
			listMagicItems = append(listMagicItems, magic)
		case "Figurine of Wondrous Power":
			list := []string{"Bronze Griffon", "Ebony Fly", "Golden Lions", "Ivory Goats", "Marble Elephant", "Obsidian Steed", "Onyx Dog", "Serpentine Owl", "Silver Raven"}
			for _, v := range list {
				desc, att := contentInterfaceParser(value)
				name := fmt.Sprintf("%s %s", key, v)
				magic := bacisDetailsMagicItem(FixName(name), name, desc, att, true, false)
				magic.Rarity = Rare
				magic.Category = WoundrousItems
				magic.HoardTable = tableG
				if v == "Silver Raven" {
					magic.Rarity = Uncommon
					magic.HoardTable = tableF
				}
				if v == "Obsidian Steed" {
					magic.Rarity = VeryRare
					magic.HoardTable = tableH
				}
				listMagicItems = append(listMagicItems, magic)
			}

		case "Mace of Smiting":
			desc, att := contentInterfaceParser(value)
			magic := bacisDetailsMagicItem(FixName(key), key, desc, att, false, false)
			magic.HoardTable = tableG
			magic.Category = Weapons
			magic.Rarity = Rare
			magic.Shape = "mace"
			feature := CoreFeatures{}
			magic.Feature = &feature
			power := CorePowers{}
			magic.Power = &power
			magic.Feature.CombateBonus = 1
			magic.Feature.DamageBonusCritical = 7
			magic.Power.Purpose = "attack-conditional-enemy"
			magic.Power.RequireEnemyType = []string{"construct"}
			magic.Power.CombateBonus = 3
			magic.Power.DamageBonusCritical = 14
			listMagicItems = append(listMagicItems, magic)

		}
	}
	// items not related in the list imported
	driftContent := `This small sphere of thick glass weighs 1 pound. If you are within 60 feet of it, you can speak its command word and cause it to emanate the light or daylight spell. Once used, the daylight effect can't be used again until the next dawn. \n	You can speak another command word as an action to make the illuminated globe rise into the air and float no more than 5 feet off the ground. The globe hovers in this way until you or another creature grasps it. If you move more than 60 feet from the hovering globe, it follows you until it is within 60 feet of you. It takes the shortest route to do so. If prevented from moving, the globe sinks gently to the ground and becomes inactive, and its light winks out.`
	driftGlobeItem := "Driftglobe"
	driftGlobe := bacisDetailsMagicItem(FixName(driftGlobeItem), driftGlobeItem, driftContent, false, true, false)
	driftGlobe.Category = WoundrousItems
	driftGlobe.HoardTable = []string{"A", "B"}
	listMagicItems = append(listMagicItems, driftGlobe)

	potionFireContent := `After drinking this potion, you can use a bonus action to exhale fire at a target within 30 feet of you. The target must make a DC 13 Dexterity saving throw, taking 4d6 fire damage on a failed save, or half as much damage on a successful one. The effect ends after you exhale the fire three times or when 1 hour has passed. This potion's orange liquid flickers, and smoke fills the top of the container and wafts out whenever it is opened.`
	potionFireItem := "Potion of Fire Breath"
	potionFire := bacisDetailsMagicItem(FixName(potionFireItem), potionFireItem, potionFireContent, false, false, false)
	potionFire.Category = Potion
	potionFire.HoardTable = tableB
	potionFirepower := CorePowers{}
	potionFire.Power = &potionFirepower
	potionFire.Power.Purpose = SpellAttack
	potionFire.Power.DamageDice = "4d6"
	potionFire.Power.DifficultClass = 13
	potionFire.Power.SavingThrow = Dexterity
	potionFire.Power.DamageType = "fire"
	listMagicItems = append(listMagicItems, potionFire)

	// 	Alchemy Jug
	alchemyContent := `This ceramic jug appears to be able to hold a gallon of liquid and weighs 12 pounds whether full or empty. Sloshing sounds can be heard from within the jug when it is shaken, even if the jug is empty. You can use an action and name one liquid from the table below to cause the jug to produce the chosen liquid. Afterward, you can uncork the jug as an action and pour that liquid out, up to 2 gallons per minute. The maximum amount of liquid the jug can produce depends on the liquid you named. Once the jug starts producing a liquid, it can't produce a different one, or more of one that has reached its maximum, until the next dawn. Liquid:Max Amount. Acid:8 ounces,Basic poison: 0,5 ounce, Beer:4 gallons, Honey: 1 gallon, Mayonnaise: 2 gallons, Oil: 1 quart, Vinegar: 2 gallons, Water, fresh:8 gallons, Water, salt: 12 gallons, Wine: 1 gallon`
	alchemyItem := "Alchemy Jug"
	alchemy := bacisDetailsMagicItem(FixName(alchemyItem), alchemyItem, alchemyContent, false, true, false)
	alchemy.Category = WoundrousItems
	alchemy.HoardTable = tableB
	listMagicItems = append(listMagicItems, alchemy)

	// Cap of Water Breathing
	capContent := `While wearing this cap underwater, you can speak its command word as an action to create a bubble of air around your head. It allows you to breathe normally underwater. This bubble stays with you until you speak the command word again, the cap is removed, or you are no longer underwater.`
	capItem := "Cap of Water Breathing"
	cap := bacisDetailsMagicItem(FixName(capItem), capItem, capContent, false, true, false)
	cap.Category = WoundrousItems
	cap.HoardTable = tableB
	listMagicItems = append(listMagicItems, cap)

	// Goggles of Night
	googlesContent := `While wearing these dark lenses, you have darkvision out to a range of 60 feet. If you already have darkvision, wearing the goggles increases its range by 60 feet.`
	googlesItem := "Goggles of Night"
	googles := bacisDetailsMagicItem(FixName(googlesItem), googlesItem, googlesContent, false, true, false)
	googles.Category = WoundrousItems
	googles.HoardTable = tableB
	listMagicItems = append(listMagicItems, googles)

	// Mariner's Armor
	marinerContent := `While wearing this armor, you have a swimming speed equal to your walking speed. In addition, whenever you start your turn underwater with 0 hit points, the armor causes you to rise 60 feet toward the surface. The armor is decorated with fish and shell motifs.`
	marinerItem := "Mariners Armor"
	for _, v := range armorsList {
		name := fmt.Sprintf("Mariners %s", v)
		magic := bacisDetailsMagicItem(FixName(name), marinerItem, marinerContent, true, false, false)
		magic.HoardTable = tableB
		magic.Category = ArmorCons
		magic.Shape = v
		feature := CoreFeatures{}
		magic.Feature = &feature
		power := CorePowers{}
		magic.Power = &power
		magic.Power.Purpose = "roleplay"
		listMagicItems = append(listMagicItems, magic)
	}

	// Saddle of the Cavalier
	saddleContent := `While in this saddle on a mount, you can't be dismounted against your will if you're conscious, and attack rolls against the mount have disadvantage.`
	saddleItem := "Saddle of the Cavalier"
	saddle := bacisDetailsMagicItem(FixName(saddleItem), saddleItem, saddleContent, false, true, false)
	saddle.Category = WoundrousItems
	saddle.HoardTable = tableB
	listMagicItems = append(listMagicItems, saddle)

	// Elixir of Health
	elixirHealthContent := `When you drink this potion, it cures any disease afflicting you, and it removes the blinded, deafened, paralyzed, and poisoned conditions. The clear red liquid has tiny bubbles of light in it.`
	elixirHealthItem := "Elixir of Health"
	elixirHealth := bacisDetailsMagicItem(FixName(elixirHealthItem), elixirHealthItem, elixirHealthContent, false, false, false)
	elixirHealth.Category = Potion
	elixirHealth.HoardTable = tableC
	elixirHealthFeature := CoreFeatures{}
	elixirHealth.Feature = &elixirHealthFeature
	elixirHealth.Feature.CancelCondition = []string{"blinded", "deafened", "paralyzed", "poisoned"}
	elixirHealth.Rarity = Rare
	listMagicItems = append(listMagicItems, elixirHealth)

	// Scroll of Protection
	scrollProtectionContent := `Using an action to read the scroll encloses you in an invisible barrier that extends from you to form a 5-foot-radius, 10-foot-high cylinder. For 5 minutes, this barrier prevents aberrations from entering or affecting anything within the cylinder. The cylinder moves with you and remains centered on you. However, if you move in such a way that an aberration would be inside the cylinder, the effect ends. A creature can attempt to overcome the barrier by using an action to make a DC 15 Charisma check. On a success, the creature ceases to be affected by the barrier.`
	scrollProtectionItem := "Scroll of Protection"
	for _, v := range monstersTypeList {
		name := fmt.Sprintf("%s from %s", scrollProtectionItem, v)
		magic := bacisDetailsMagicItem(FixName(name), scrollProtectionItem, scrollProtectionContent, false, true, false)
		magic.HoardTable = tableC
		magic.Category = "scrolls"
		listMagicItems = append(listMagicItems, magic)
	}

	// Sending Stones
	sendingStoneContent := `Sending stones come in pairs, with each smooth stone carved to match the other so the pairing is easily recognized. While you touch one stone, you can use an action to cast the sending spell from it. The target is the bearer of the other stone. If no creature bears the other stone, you know that fact as soon as you use the stone and don't cast the spell. Once sending is cast through the stones, they can't be used again until the next dawn. If one of the stones in a pair is destroyed, the other one becomes nonmagical.`
	sendingStoneItem := "Sending Stones"
	sendingStone := bacisDetailsMagicItem(FixName(sendingStoneItem), sendingStoneItem, sendingStoneContent, false, true, false)
	sendingStone.Category = WoundrousItems
	sendingStone.HoardTable = tableC
	listMagicItems = append(listMagicItems, sendingStone)

	// Sentinel Shield
	sentinelShieldContent := `While holding this shield, you have advantage on initiative rolls and Wisdom (Perception) checks. The shield is emblazoned with a symbol of an eye.`
	sentinelShieldItem := "Sentinel Shield"
	sentinelShield := bacisDetailsMagicItem(FixName(sentinelShieldItem), sentinelShieldItem, sentinelShieldContent, false, false, false)
	sentinelShield.Shape = "shield"
	sentinelShield.Category = ArmorCons
	sentinelShield.HoardTable = tableF
	sentinelShieldFeature := CoreFeatures{}
	sentinelShield.Feature = &sentinelShieldFeature
	sentinelShield.Feature.Advantages = []string{"initiative", "perception"}
	listMagicItems = append(listMagicItems, sentinelShield)

	// +1 rod of the pact keeper
	pactKeeperContent := `While holding this rod, you gain a +1 bonus to spell attack rolls and to the saving throw DCs of your warlock spells. In addition, you can regain one warlock spell slot as an action while holding the rod. You can't use this property again until you finish a long rest.`
	pactKeeperItem := "Rod of the Pact Keeper"
	for _, v := range bonusList {
		name := fmt.Sprintf("%s %v", pactKeeperItem, v)
		magic := bacisDetailsMagicItem(FixName(name), pactKeeperItem, pactKeeperContent, true, false, false)
		magic.HoardTable = tableF
		magic.AttunementRestriction = []string{"warlock"}
		magic.Category = Rods
		feature := CoreFeatures{}
		magic.Feature = &feature
		magic.Feature.SpellBonus = v
		listMagicItems = append(listMagicItems, magic)
	}

	// 	Staff of the Adder
	staffAdderContent := `You can use a bonus action to speak this staff's command word and make the head of the staff become that of an animate poisonous snake for 1 minute. By using another bonus action to speak the command word again, you return the staff to its normal inanimate form. You can make a melee attack using the snake head, which has a reach of 5 feet. Your proficiency bonus applies to the attack roll. On a hit, the target takes 1d6 piercing damage and must succeed on a DC 15 Constitution saving throw or take 3d6 poison damage. The snake head can be attacked while it is animate. It has an Armor Class of 15 and 20 hit points. If the head drops to 0 hit points, the staff is destroyed. As long as it's not destroyed, the staff regains all lost hit points when it reverts to its inanimate form. Versatile. This weapon can be used with one or two hands. A damage value in parentheses appears with the property—the damage when the weapon is used with two hands to make a melee attack.`
	staffAdderItem := "Staff of the Adder"
	staffAdder := bacisDetailsMagicItem(FixName(staffAdderItem), staffAdderItem, staffAdderContent, true, false, false)
	staffAdder.Category = Weapons
	staffAdder.AttunementRestriction = []string{"cleric", "druid", "warlock"}
	staffAdder.Shape = "quarterstaff"
	staffAdder.HoardTable = tableF
	staffAdderPower := CorePowers{}
	staffAdder.Power = &staffAdderPower
	staffAdder.Power.Purpose = "attack-with-saving-damage"
	staffAdder.Power.DifficultClass = 15
	staffAdder.Power.SavingThrow = Constitution
	staffAdder.Power.DamageDice = "3d6"
	staffAdder.Power.DamageType = "poison"
	listMagicItems = append(listMagicItems, staffAdder)

	// Sword of Vengeance
	swordVengeanceContent := `You gain a +1 bonus to attack and damage rolls made with this magic weapon. Curse. This sword is cursed and possessed by a vengeful spirit. Becoming attuned to it extends the curse to you. As long as you remain cursed, you are unwilling to part with the sword, keeping it on your person at all times. While attuned to this weapon, you have disadvantage on attack rolls made with weapons other than this one. In addition, while the sword is on your person, you must succeed on a DC 15 Wisdom saving throw whenever you take damage in combat. On a failed save you must attack the creature that damaged you until you drop to 0 hit points or it does, or until you can't reach the creature to make a melee attack against it. You can break the curse in the usual ways. Alternatively, casting banishment on the sword forces the vengeful spirit to leave it. The sword then becomes a +1 weapon with no other properties.`
	swordVengeanceItem := "Sword of Vengeance"
	for _, v := range anySword {
		name := fmt.Sprintf("%s of Vengeance", v)
		magic := bacisDetailsMagicItem(FixName(name), swordVengeanceItem, swordVengeanceContent, true, false, false)
		magic.HoardTable = tableF
		magic.Category = Weapons
		magic.Shape = v
		feature := CoreFeatures{}
		magic.Feature = &feature
		power := CorePowers{}
		magic.Power = &power
		magic.Feature.CombateBonus = 1
		magic.Power.Purpose = "cursed-weapon"
		magic.Power.Curse = true
		magic.Power.Disvantages = []string{"attack"}
		listMagicItems = append(listMagicItems, magic)
	}

	// Weapon of Warning
	weaponWarningItem := "Weapon of Warning"
	weaponWarningContent := `This magic weapon warns you of danger. While the weapon is on your person, you have advantage on initiative rolls. In addition, you and any of your companions within 30 feet of you can't be surprised, except when incapacitated by something other than nonmagical sleep. The weapon magically awakens you and your companions within range if any of you are sleeping naturally when combat begins.`
	for _, v := range weaponsList {
		name := fmt.Sprintf("%s of Warning", v)
		weaponWarning := bacisDetailsMagicItem(FixName(name), weaponWarningItem, weaponWarningContent, true, false, false)
		weaponWarning.Category = Weapons
		weaponWarning.HoardTable = tableF
		weaponWarning.Shape = v
		feature := CoreFeatures{}
		weaponWarning.Feature = &feature
		weaponWarning.Feature.Advantages = []string{"initiative"}
		listMagicItems = append(listMagicItems, weaponWarning)
	}

	// Gloves of Thievery
	glovesThieveryContent := `These gloves are invisible while worn. While wearing them, you gain a +5 bonus to Dexterity (Sleight of Hand) checks and Dexterity checks made to pick locks.`
	glovesThieveryItem := "Gloves of Thievery"
	glovesThievery := bacisDetailsMagicItem(FixName(glovesThieveryItem), glovesThieveryItem, glovesThieveryContent, false, false, false)
	glovesThievery.Category = WoundrousItems
	glovesThievery.HoardTable = tableF
	skillMap := map[string]int{"sleight-of-hand": 5}
	glovesThieveryFeature := CoreFeatures{}
	glovesThievery.Feature = &glovesThieveryFeature
	glovesThievery.Feature.SkillBonus = skillMap
	listMagicItems = append(listMagicItems, glovesThievery)

	// Instrument of the Bards
	instrumentContent := `An instrument of the bards is an exquisite example of its kind, superior to an ordinary instrument in every way. Seven types of these instruments exist, each named after a legendary bard college. A creature that attempts to play the instrument without being attuned to it must succeed on a DC 15 Wisdom saving throw or take 2d4 psychic damage. You can use an action to play the instrument and cast one of its spells. Once the instrument has been used to cast a spell, it can't be used to cast that spell again until the next dawn. The spells use your spellcasting ability and spell save DC. You can play the instrument while casting a spell that causes any of its targets to be charmed on a failed saving throw, thereby imposing disadvantage on the save. This effect applies only if the spell has a somatic or a material component.`
	bardInstruments := []string{"Anstruth Harp", "Canaith Mandolin", "Cli Lyre", "Doss Lute", "Fochlucan Bandore", "Mac-Fuirmidh Cittern", "Ollamh Harp"}
	bardItem := "Instrument of the Bards"
	for _, v := range bardInstruments {
		name := fmt.Sprintf("%s %s", bardItem, v)
		instrumentBard := bacisDetailsMagicItem(FixName(name), name, instrumentContent, true, false, false)
		instrumentBard.Category = WoundrousItems
		// feature := CoreFeatures{}
		// instrumentBard.Feature = &feature
		power := CorePowers{}
		instrumentBard.Power = &power
		instrumentBard.Power.Purpose = "spell-without-spellcasting"
		instrumentBard.AttunementRestriction = []string{"bard"}
		switch v {
		case "Anstruth Harp":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "control-weather", "cure-wounds", "wall-of-thorns"}
			instrumentBard.HoardTable = tableH
			instrumentBard.Rarity = VeryRare
			instrumentBard.Power.SpellLevel = 5
		case "Canaith Mandolin":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "cure-wounds", "dispel-magic", "protection-from-energy"}
			instrumentBard.HoardTable = tableG
			instrumentBard.Rarity = Rare
			instrumentBard.Power.SpellLevel = 3
		case "Cli Lyre":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "stone-shape", "wall-of-fire", "wind-wall"}
			instrumentBard.HoardTable = tableG
			instrumentBard.Rarity = Rare
			instrumentBard.Power.SpellLevel = 3
		case "Doss Lute":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "animal-friendship", "protection-from-energy", "protection-from-poison"}
			instrumentBard.HoardTable = tableF
			instrumentBard.Power.SpellLevel = 1
		case "Fochlucan Bandore":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "entangle", "faerie-fire", "shillelagh", "speak-with-animals"}
			instrumentBard.HoardTable = tableF
			instrumentBard.Power.SpellLevel = 1
		case "Mac-Fuirmidh Cittern":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "barkskin", "cure-wounds", "fog-cloud"}
			instrumentBard.HoardTable = tableF
			instrumentBard.Power.SpellLevel = 1
		case "Ollamh Harp":
			instrumentBard.Power.SpellList = []string{"fly", "invisibility", "leviate", "protection-from-evil-and-good", "confusion", "control-weather", "fire-storm"}
			instrumentBard.HoardTable = tableI
			instrumentBard.Rarity = Legendary
			instrumentBard.Power.SpellLevel = 5
		}
		listMagicItems = append(listMagicItems, instrumentBard)
	}

	// Tentacle Rod
	tentacleRodContent := `Made by the drow, this rod is a magic weapon that ends in three rubbery tentacles. While holding the rod, you can use an action to direct each tentacle to attack a creature you can see within 15 feet of you. Each tentacle makes a melee attack roll with a +9 bonus. On a hit, the tentacle deals 1d6 bludgeoning damage. If you hit a target with all three tentacles, it must make a DC 15 Constitution saving throw. On a failure, the creature's speed is halved, it has disadvantage on Dexterity saving throws, and it can't use reactions for 1 minute. Moreover, on each of its turns, it can take either an action or a bonus action, but not both. At the end of each of its turns, it can repeat the saving throw, ending the effect on itself on a success.`
	tentacleRod := bacisDetailsMagicItem(FixName("Tentacle Rod"), "Tentacle Rod", tentacleRodContent, true, false, false)
	tentacleRod.Category = Rods
	tentacleRod.HoardTable = tableG
	tentacleRod.Rarity = Rare
	tentacleRodPower := CorePowers{}
	tentacleRod.Power = &tentacleRodPower
	tentacleRod.Power.Purpose = "attack-with-saving-continuos-disvantage"
	tentacleRod.Power.DamageType = Bludgeoning
	tentacleRod.Power.DamageDice = "1d6"
	tentacleRod.Power.AttackNumber = 3
	tentacleRod.Power.ConditionMultiple = 3
	tentacleRod.Power.DifficultClass = 15
	tentacleRod.Power.SavingThrow = Constitution
	tentacleRod.Power.Disvantages = []string{Dexterity}
	listMagicItems = append(listMagicItems, tentacleRod)

	prettyJSON, err := json.MarshalIndent(listMagicItems, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s", string(prettyJSON))
	// fmt.Println("}")
}

func contentInterfaceParser(value interface{}) (string, bool) {
	var desc string
	var requiredAttunement bool

	content := value.(map[string]interface{})
	for _, v := range content {
		// fmt.Println(reflect.TypeOf(v).Kind())
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				value := fmt.Sprintln(s.Index(i))
				if validateItemType(value) {
					title := strings.ReplaceAll(strings.TrimSuffix(value, "\n"), "*", "")
					if strings.Contains(title, "requires attunement") {
						requiredAttunement = true
					}
				}
				if !validateItemType(value) {
					desc += strings.TrimSuffix(value, "\n")

				}

			}
		}
	}
	return desc, requiredAttunement
}

// instanciate a basic magic item with some basic empty content
func bacisDetailsMagicItem(name, key, desc string, att, roleplay, forbidden bool) MagicItem {
	magic := MagicItem{
		Name:                  strings.ToLower(name),
		Title:                 key,
		Content:               desc,
		Category:              "",
		Rarity:                Uncommon,
		HoardTable:            []string{},
		AttunementRestriction: []string{},
		RequiredAttunement:    att,
		RolePlay:              roleplay,
		Forbidden:             forbidden,
		Shape:                 "",
	}
	return magic
}

func validateItemType(title string) bool {
	if strings.Contains(title, "*Wondrous item") {
		return true
	}
	if strings.Contains(title, "*Weapon") {
		return true
	}
	if strings.Contains(title, "*Staff") {
		return true
	}
	if strings.Contains(title, "*Armor (") {
		return true
	}
	if strings.Contains(title, "*Ring") {
		return true
	}
	if strings.Contains(title, "*Wand") {
		return true
	}
	if strings.Contains(title, "*Potion") {
		return true
	}
	if strings.Contains(title, "*Staff") {
		return true
	}
	if strings.Contains(title, "*Rod") {
		return true
	}
	return false
}
