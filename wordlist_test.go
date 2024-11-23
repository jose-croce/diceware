// wordlist unit tests
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

import "testing"

func TestMakerWordMap(t *testing.T) {
	word_map, err := MakeWordMap("en")
	if err != nil {
		t.Error(err)
	}
	tests := WordMap{
		11111: "a",
		12345: "apathy",
		66666: "@",
	}
	for dice_value, want := range tests {
		got, ok := word_map[dice_value]
		if !ok {
			t.Errorf("Missing mandatory word for %d", dice_value)
		}
		if got != want {
			t.Errorf("Unexpected word for %d (want='%s', got='%s')", dice_value, want, got)
		}
	}
}
