# Hunspell Golang Wrapper
```It is designed for quick and high quality spell checking and correcting for languages with word-level writing system, including languages with rich morphology, complex word compounding and character encoding.```


## Overview
Here's my attempt at combining the implementation to wrapping hunspell C++ library with Golang CGO.

## the system that i used
- Ubuntu 18.04.5 LTS
- Intel core I5 8th Gen

## Requirement
- `git`
- `make`
- hunspell [`requirement`](http://hunspell.github.io/)
- hunspell `dictionary` [#1](https://github.com/hunspell/hunspell#dictionaries) [#2](https://github.com/client9/gospell#where-can-i-get-english-dictionaries)

## Prepare the requirement
- clone hunspell [repository](https://github.com/hunspell/hunspell) OR [download released version](https://github.com/hunspell/hunspell/files/2573619/hunspell-1.7.0.tar.gz) ([v1.7.0](https://github.com/hunspell/hunspell/releases/tag/v1.7.0)) 
```
$ git clone https://github.com/hunspell/hunspell
# Cloning...
$ cd hunspell
$ git checkout v1.7.0 # I currently test using this version
# Compile it
```
OR 
```
# Download commpressed file
$ wget https://github.com/hunspell/hunspell/files/2573619/hunspell-1.7.0.tar.gz
# Extract file
$ tar -xvzf hunspell-1.7.0.tar.gz
$ cd hunspell-1.7.0
# Compile it
```

- compile hunspell library
```
$ autoreconf -vfi
$ ./configure
$ make
$ cd src/hunspell/.libs/
# find `libhunspell-1.7.a` file

# find header file in <path-to-root-library>/hunspell-1.7.0/src/hunspell
# 1) hunspell.h  2)hunvisapi.h

# copy required file to repository
$ cp <path-to-root-library>/hunspell-1.7.0/src/hunspell/.libs/libhunspell-1.7.a <path-to-root-repository>/hunspell-go-warper/hunspell/lib/libhunspell-1.7.a 
$ cp <path-to-root-library>/hunspell-1.7.0/src/hunspell/hunspell.h <path-to-root-repository>/hunspell-go-warper/hunspell/include/hunspell.h
$ cp <path-to-root-library>/hunspell-1.7.0/src/hunspell/hunvisapi.h <path-to-root-repository>/hunspell-go-warper/hunspell/include/hunvisapi.h
```

## Usage 
- write the program 
```
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
}

```


- run the program sample
```
& go run app.go


## the output will be like this :

# stem (plants) : [plant]
# spell (plants) : true
# suggest (plants) : [plant plans plats pants planets plaints plant's planes slants plaits planks plan's pl ants pl-ants plan ts]
# ===========================================
# stem (plaant) : []
# spell (plaant) : false
# suggest (plaant) : [plant plaint pliant plangent plan]
```




## Refference
- [Hunspell main library](https://github.com/hunspell/hunspell)
- [Huspell go implementation](https://github.com/akhenakh/hunspellgo)