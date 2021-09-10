#!/usr/bin/env bash

cat 08\ spellcasting.json |jq '.Spellcasting."Spell Lists"' > output/spell-list.json
#Cantrips (0 Level)
sed -i 's/Cantrips (0 Level)/level0/g' output/spell-list.json
#1st Level
sed -i 's/1st Level/level1/g' output/spell-list.json
#2nd Level
sed -i 's/2nd Level/level2/g' output/spell-list.json
#3rd Level
sed -i 's/3rd Level/level3/g' output/spell-list.json
#4th Level
sed -i 's/4th Level/level4/g' output/spell-list.json
#5th Level
sed -i 's/5th Level/level5/g' output/spell-list.json
#6th Level
sed -i 's/6th Level/level6/g' output/spell-list.json
#7th Level
sed -i 's/7th Level/level7/g' output/spell-list.json
#8th Level
sed -i 's/8th Level/level8/g' output/spell-list.json
#9th Level
sed -i 's/9th Level/level9/g' output/spell-list.json
