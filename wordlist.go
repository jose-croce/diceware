// word lists for diceware
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
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed data/*
var word_fs embed.FS

// Mapping of words indexed by dice values
type WordMap map[uint32]string

// Language data container
type LangData struct {
	// language description
	description string
	// Name of the file with language dictionary
	file_name string
}

// Error in case of bad word list formatting
type BadFormatError struct {
	language string
	line     string
}

// Error interface compliance
func (e *BadFormatError) Error() string {
	return fmt.Sprintf("invalid format for line '%s' from '%s' language", e.line, e.language)
}

// Error in case of bad dice value for word
type BadDiceValueError struct {
	language   string
	dice_value string
	word       string
	inner      error
}

// Error interface compliance
func (e *BadDiceValueError) Error() string {
	return fmt.Sprintf("bad dice value '%s' for word '%s' from '%s' language", e.dice_value, e.word, e.language)
}

// Error interface compliance
func (e *BadDiceValueError) Unwrap() error {
	return e.inner
}

var Languages = map[string]LangData{
	"en": {
		description: "English",
		file_name:   "data/wordlist.txt",
	},
	"beale": {
		description: "alternative Beale list",
		file_name:   "data/wordlist.beale.txt",
	},
	"cat": {
		description: "Catalan",
		file_name:   "data/wordlist.cat.txt",
	},
	"de": {
		description: "German",
		file_name:   "data/wordlist.de.txt",
	},
	"es": {
		description: "Spanish",
		file_name:   "data/wordlist.es.txt",
	},
	"eu": {
		description: "Euskera",
		file_name:   "data/wordlist.eu.txt",
	},
	"fi": {
		description: "Finnish",
		file_name:   "data/wordlist.fi.txt",
	},
	"fr": {
		description: "French",
		file_name:   "data/wordlist.fr.txt",
	},
	"it": {
		description: "Italian",
		file_name:   "data/wordlist.it.txt",
	},
	"la": {
		description: "Latin",
		file_name:   "data/wordlist.la.txt",
	},
	"pt": {
		description: "Portuguese",
		file_name:   "data/wordlist.pt.txt",
	},
}

func MakeWordMap(language string) (WordMap, error) {
	lang_data, ok := Languages[language]
	if !ok {
		return nil, fmt.Errorf(
			"no word list configured for '%s' language",
			language,
		)
	}

	file, err := word_fs.Open(lang_data.file_name)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to open word list file for %s language: %w",
			lang_data.description,
			err,
		)
	}
	defer file.Close()

	word_list_pattern := regexp.MustCompile(`^([1-6]{5})\s+(\S+)\s*$`)
	word_map := make(WordMap)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := word_list_pattern.FindStringSubmatch(line)
		if len(parsed) != 3 {
			return nil, &BadFormatError{language: language, line: line}
		}
		dice_val, err := strconv.ParseUint(parsed[1], 10, 32)
		if err != nil {
			return nil, &BadDiceValueError{language: language, dice_value: parsed[1], word: parsed[2], inner: err}
		}
		word_map[uint32(dice_val)] = parsed[2]
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading word list file for %s language: %w", lang_data.description, err)
	}

	return word_map, nil
}
