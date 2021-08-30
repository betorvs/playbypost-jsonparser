# database/main

It's a json parser to create 4 files used by playbypost-dnd:
- new-magic-itens-list.json  
- new-spell-description-list.json  
- new-monster-list.json  
- spell-list.json  



## How to create it

### Download srd_5e_monsters.json

https://gist.github.com/tkfu

#### patch it


```
--- srd_5e_monsters.json 2020-05-27 22:33:15.000000000 +0200
+++ srd_5e_monsters.json        2020-06-02 23:25:16.000000000 +0200
@@ -4050,7 +4050,7 @@
     "Languages": "Draconic",
     "Challenge": "3 (700 XP)",
     "Traits": "<p><strong>Amphibious</strong>: The dragon can breathe air and water.</p>",
-    "Actions": "<p><strong>Bite</strong>: <em>Melee Weapon Attack</em>: +6 to hit, reach 5 ft., one target. <em>Hit</em>: 9 (1d10 + 4) piercing damage.</p><p><strong>Breath Weapons (Recharge 5–6)</strong>: The dragon uses one of the following breath weapons.</p><p><strong>Fire Breath</strong>: The dragon exhales fire in a 15-foot cone. Each creature in that area must make a DC 13 Dexterity saving throw, taking 22 (4d10) fire damage on a failed save, or half as much damage on a successful one.</p><p><strong>Weakening Breath</strong>: The dragon exhales gas in a 15-foot cone. Each creature in that area must succeed on a DC 13 Strength saving throw or have disadvantage on Strength-based attack rolls, Strength checks, and Strength saving throws for 1 minute. A creature can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success.</p>",
+    "Actions": "<p><strong>Bite.</strong> <em>Melee Weapon Attack:</em> +6 to hit, reach 5 ft., one target. <em>Hit:</em> 9 (1d10 + 4) piercing damage.</p><p><strong>Breath Weapons (Recharge 5–6):</strong> The dragon uses one of the following breath weapons.</p><p><strong>Fire Breath:</strong> The dragon exhales fire in a 15-foot cone. Each creature in that area must make a DC 13 Dexterity saving throw, taking 22 (4d10) fire damage on a failed save, or half as much damage on a successful one.</p><p><strong>Weakening Breath:</strong> The dragon exhales gas in a 15-foot cone. Each creature in that area must succeed on a DC 13 Strength saving throw or have disadvantage on Strength-based attack rolls, Strength checks, and Strength saving throws for 1 minute. A creature can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success.</p>",
     "img_url": "https://media-waterdeep.cursecdn.com/avatars/thumbnails/7/516/315/315/636285466148376212.jpeg"
   },
   {
```


### Download "08 spellcasting.json" and "10 magic items.json"

https://github.com/BTMorton/dnd-5e-srd/tree/master/json


#### Clean up 

```
$ cat 08\ spellcasting.json |jq '.Spellcasting."Spell Descriptions"' > Spell-Descriptions.json

$ cat 10\ magic\ items.json | jq '."Magic Items"' > Magic-Items.json
and remove first line "content"

$ cat 04\ equipment.json | jq '.Equipment.Armor."Armor List"' > Armors-List.json

$ cat 04\ equipment.json | jq '.Equipment.Weapons."Weapons List"' > Weapons-List.json

$ cat 04\ equipment.json | jq '.Equipment."Adventuring Gear"."Adventuring Gear".table' > Adventure-Gear.json

$ cat 04\ equipment.json | jq '.Equipment."Adventuring Gear"."Equipment Packs"' > Gear-Packs.json

$ cat 04\ equipment.json | jq '.Equipment."Tools"."Tools".content' > Tools.json

$ cat 04\ equipment.json | jq '.Equipment."Mounts and Vehicles"."Mounts and Other Animals"' > Mounts.json
```

#### Patch it to add "Hellish Rebuke"
```
--- Spell-Descriptions.json	2020-05-26 15:15:41.000000000 +0200
+++ Spell-Descriptions.json	2020-06-04 08:49:49.000000000 +0200
@@ -1974,6 +1974,17 @@
       "***At Higher Levels.*** When you cast this spell using a spell slot of 3rd level or higher, the damage increases by 1d8 for each slot level above 2nd."
     ]
   },
+  "Hellish Rebuke": {
+    "content": [
+      "*1st-level evocation*",
+      "**Casting Time:** 1 reaction, which you take in response to being damaged by a creature within 60 feet of you that you can see",
+      "**Range:** 60 feet",
+      "**Components:** V, S",
+      "**Duration:** Instantaneous",
+      "You point your finger, and the creature that damaged you is momentarily surrounded by hellish flames. The creature must make a Dexterity saving throw. It takes 2d10 fire damage on a failed save, or half as much damage on a successful one.",
+      "***At Higher Levels.*** When you cast this spell using a spell slot of 2nd level or higher, the damage increases by 1d10 for each slot level above 1st."
+    ]
+  },
   "Heroes’ Feast": {
     "content": [
       "*6th-level conjuration*",
@@ -3883,7 +3894,7 @@
       "You create a wall of ice on a solid surface within range. You can form it into a hemispherical dome or a sphere with a radius of up to 10 feet, or you can shape a flat surface made up of ten 10-foot-square panels. Each panel must be contiguous with another panel. In any form, the wall is 1 foot thick and lasts for the duration.",
       "If the wall cuts through a creature’s space when it appears, the creature within its area is pushed to one side of the wall and must make a Dexterity saving throw. On a failed save, the creature takes 10d6 cold damage, or half as much damage on a successful save.",
       "The wall is an object that can be damaged and thus breached. It has AC 12 and 30 hit points per 10-foot section, and it is vulnerable to fire damage. Reducing a 10-foot section of wall to 0 hit points destroys it and leaves behind a sheet of frigid air in the space the wall occupied. A creature moving through the sheet of frigid air for the first time on a turn must make a Constitution saving throw. That creature takes 5d6 cold damage on a failed save, or half as much damage on a successful one.",
-      "***At Higher Levels.*** When you cast this spell using a spell slot of 7th level or higher, the damage the wall deals when it appears increases by 2d6, and the damage from passing through the sheet of frigid air increases by 1d6, for each slot level above 6th."
+      "***At Higher Levels.*** When you cast this spell using a spell slot of 7th level or higher, the damage the wall deals when it appears increases by 2d6, and the damage from passing through the sheet of frigid air increases by 1d6 for each slot level above 6th."
     ]
   },
   "Wall of Stone": {
```

Patch Tools.json:
```
--- Tools-old.json	2021-01-18 10:05:31.000000000 +0100
+++ Tools.json	2021-01-18 10:49:58.000000000 +0100
@@ -1,11 +1,11 @@
-[
+
   {
     "table": {
       "Item": [
         "*Artisan’s tools*",
         "Alchemist’s supplies",
         "Brewer’s supplies",
-        "Calligrapher&#39;s supplies",
+        "Calligraphers supplies",
         "Carpenter’s tools",
         "Cartographer’s tools",
         "Cobbler’s tools",
@@ -125,15 +125,5 @@
         "*"
       ]
     }
-  },
-  "* See the “Mounts and Vehicles” section.",
-  "***Artisan’s Tools.*** These special tools include the items needed to pursue a craft or trade. The table shows examples of the most common types of tools, each providing items related to a single craft. Proficiency with a set of artisan’s tools lets you add your proficiency bonus to any ability checks you make using the tools in your craft. Each type of artisan’s tools requires a separate proficiency.",
-  "***Disguise Kit.*** This pouch of cosmetics, hair dye, and small props lets you create disguises that change your physical appearance. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to create a visual disguise.",
-  "***Forgery Kit.*** This small box contains a variety of papers and parchments, pens and inks, seals and sealing wax, gold and silver leaf, and other supplies necessary to create convincing forgeries of physical documents. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to create a physical forgery of a document.",
-  "***Gaming Set.*** This item encompasses a wide range of game pieces, including dice and decks of cards (for games such as Three-Dragon Ante). A few common examples appear on the Tools table, but other kinds of gaming sets exist. If you are proficient with a gaming set, you can add your proficiency bonus to ability checks you make to play a game with that set. Each type of gaming set requires a separate proficiency.",
-  "***Herbalism Kit.*** This kit contains a variety of instruments such as clippers, mortar and pestle, and pouches and vials used by herbalists to create remedies and potions. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to identify or apply herbs. Also, proficiency with this kit is required to create antitoxin and potions of healing.",
-  "***Musical Instrument.*** Several of the most common types of musical instruments are shown on the table as examples. If you have proficiency with a given musical instrument, you can add your proficiency bonus to any ability checks you make to play music with the instrument. A bard can use a musical instrument as a spellcasting focus. Each type of musical instrument requires a separate proficiency.",
-  "***Navigator’s Tools.*** This set of instruments is used for navigation at sea. Proficiency with navigator&#39;s tools lets you chart a ship&#39;s course and follow navigation charts. In addition, these tools allow you to add your proficiency bonus to any ability check you make to avoid getting lost at sea.",
-  "***Poisoner’s Kit.*** A poisoner’s kit includes the vials, chemicals, and other equipment necessary for the creation of poisons. Proficiency with this kit lets you add your proficiency bonus to any ability checks you make to craft or use poisons.",
-  "***Thieves’ Tools.*** This set of tools includes a small file, a set of lock picks, a small mirror mounted on a metal handle, a set of narrow-bladed scissors, and a pair of pliers. Proficiency with these tools lets you add your proficiency bonus to any ability checks you make to disarm traps or open locks."
-]
+  }
+
```

### Generate new files

spell-list.json
```bash
$ cat 08\ spellcasting.json |jq '.Spellcasting."Spell Lists"' > ../spell-list.json
```

Others
```bash
$ pwd
.../playbypost-dnd/database/main
go run . monster >../new-monster-list.json
go run . spell >../new-spell-description-list.json
go run . magicitem>../new-magic-itens-list.json
go run . armor>../new-armor-list.json
go run . weapon>../new-weapon-list.json
go run . gear>../new-gear-list.json
go run . packs>../new-packs-list.json
go run . tools>../new-tools-list.json
go run . mounts>../new-mounts-list.json
go run . hoard >../new-treasure-hoard-list.json
go run . services >../new-services-list.json
```
