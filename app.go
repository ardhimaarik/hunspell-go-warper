package main

import (
	"fmt"
	"github.com/ardhimaarik/go-hunspell-warpper/hunspell"
)

// Example implementation

func main() {

	// init hunspell library and load dictionary
	predictor := hunspell.Hunspell("files/dictionary/hunspell-en.aff", "files/dictionary/hunspell-en.dic")

	// english word with affix
	word := "plants" // english word with affix
	stem := predictor.Stem(word)
	spell := predictor.Spell(word)
	suggest := predictor.Suggest(word)
	fmt.Println("stem (plants) :", stem)
	fmt.Println("spell (plants) :", spell)
	fmt.Println("suggest (plants) :", suggest)

	fmt.Println("===========================================")

	word = "plaant" // english word with typo
	stem = predictor.Stem(word)
	spell = predictor.Spell(word)
	suggest = predictor.Suggest(word)
	fmt.Println("stem (plaant) :", stem)
	fmt.Println("spell (plaant) :", spell)
	fmt.Println("suggest (plaant) :", suggest)

	// // the output will be like this :

	// stem (plants) : [plant]
	// spell (plants) : true
	// suggest (plants) : [plant plans plats pants planets plaints plant's planes slants plaits planks plan's pl ants pl-ants plan ts]
	// ===========================================
	// stem (plaant) : []
	// spell (plaant) : false
	// suggest (plaant) : [plant plaint pliant plangent plan]
}
