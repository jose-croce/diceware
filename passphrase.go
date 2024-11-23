// Passphrase generator
//
// Copyright © 2024 Jose Ignacio Croce Busquets
// This file is part of diceware.
//
// Diceware is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// Diceware is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with Diceware. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

// Define valid word modifiers
type WordModifier int

const (
	PascalCase WordModifier = iota
	LowerCase
	UpperCase
)

// Dice roll emulation. A value in the range [1-6] is returned
func dice_roll() (uint8, error) {
	faces := big.NewInt(6)
	result, err := rand.Int(rand.Reader, faces)
	if err != nil {
		return 0, err
	}
	return uint8(result.Int64()) + 1, nil
}

// Emulate the throw of 5 dice
func dice_throw() (uint32, error) {
	var accum uint32 = 0
	for i := 0; i < 5; i++ {
		roll, err := dice_roll()
		if err != nil {
			return 0, err
		}
		accum = accum*10 + uint32(roll)
	}
	return accum, nil
}

// Apply word modification
func modify_word(base string, modifier WordModifier) string {
	if modifier == UpperCase {
		return strings.ToUpper(base)
	}
	if modifier == LowerCase {
		return strings.ToLower(base)
	}
	initial := unicode.ToUpper(rune(base[0]))
	return string(initial) + base[1:]
}

// Shuffle elements within a list of strings
func shuffle(l []string) []string {
	var result []string
	for aux := l; len(aux) > 0; {
		var idx int = 0
		if len(aux) > 1 {
			limit := big.NewInt(int64(len(aux)))
			big_idx, err := rand.Int(rand.Reader, limit)
			if err != nil {
				log.Fatal(err)
			}
			idx = int(big_idx.Int64())
		}
		result = append(result, aux[idx])
		if idx == 0 {
			aux = aux[1:]
		} else if idx == len(aux) {
			aux = aux[:idx]
		} else {
			aux = append(aux[:idx], aux[idx+1:]...)
		}
	}
	return result
}

// Pass phrase generator
//
// word_map		Mapping of words to use
// length		Passphase length
// modifier		Modification to apply to each word chosen
// force_digits Force the inclussion of digits within the resulting passphrase
func Generate(word_map WordMap, length int, modifier WordModifier, force_digits bool, delimited bool) string {
	var words []string
	for i := 0; i < length; i++ {
		throw, err := dice_throw()
		if err != nil {
			log.Fatal(err)
		}
		words = append(words, modify_word(word_map[throw], modifier))
	}

	if force_digits {
		digits_pattern := regexp.MustCompile(`\d+`)
		if digits_pattern.Find([]byte(strings.Join(words, ""))) == nil {
			words = shuffle(words)[1:]
			max_num := big.NewInt(10000)
			digits, err := rand.Int(rand.Reader, max_num)
			if err != nil {
				log.Fatal(err)
			}
			words = shuffle(append(words, digits.String()))
		}
	}

	if delimited {
		delimiters := strings.Split(
			"! ? @ # $ % & + - * / = < > ( ) [ ] { } _ . : , ; ' \"",
			" ",
		)
		var delimited []string
		for idx, word := range words {
			delimited = append(delimited, word)
			if idx < len(words)-1 {
				delimiters = shuffle(delimiters)
				delimited = append(delimited, delimiters[0])
			}
		}
		words = delimited
	}

	return strings.Join(words, "")
}
