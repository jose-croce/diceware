// Diceware passphrase generator
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
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Absolute min/max number of words for a passphrase
const min_words = 5
const max_words = 16

func main() {
	lang := "en"
	length := min_words
	word_modifier := PascalCase
	log.SetPrefix("diceware: ")
	log.SetFlags(0)

	word_map, err := MakeWordMap(lang)
	if err != nil {
		log.Fatal(err)
	}

	available_languages := flag.Bool(
		"available-languages",
		false,
		"List all supported languages and exit.")
	digits := flag.Bool(
		"digits",
		false,
		"Force the inclusion of digits in the passphrase.")
	flag.Func(
		"lang",
		"Specify the language to use for word selection.",
		func(arg string) error {
			lang = arg
			word_map, err = MakeWordMap(lang)
			return err
		},
	)
	flag.Func(
		"length",
		"Set the desired number of words in the passphrase.",
		func(arg string) error {
			passphrase_length, err := strconv.Atoi(arg)
			if err != nil {
				return err
			}
			if passphrase_length < min_words {
				return errors.New("insufficient length to provide adequate security")
			}
			if passphrase_length > max_words {
				return errors.New("passphrase too long")
			}
			return nil
		},
	)
	lower := flag.Bool(
		"lower",
		false,
		"Force all words to be lowercase.")
	symbols := flag.Bool(
		"symbols",
		false,
		"Add delimiting symbols between words.")
	forced := flag.Bool(
		"upper",
		false,
		"Force all words to be uppercase.")

	flag.Parse()

	if *available_languages {
		fmt.Println("Passphrase languages available:")
		lang_keys := make([]string, 0, len(Languages))
		for key := range Languages {
			lang_keys = append(lang_keys, key)
		}
		sort.Strings(lang_keys)
		for _, key := range lang_keys {
			fmt.Printf("\t%s\t%s\n", key, Languages[key].description)
		}
		return
	}

	if *lower {
		if *forced {
			log.Print(
				"Cannot specify -lower and -upper options simultanewously",
			)
			flag.PrintDefaults()
			os.Exit(1)
		}
		word_modifier = LowerCase
	} else if *forced {
		word_modifier = UpperCase
	}

	fmt.Println(
		Generate(
			word_map,
			length,
			word_modifier,
			*digits,
			*symbols,
		),
	)
}
