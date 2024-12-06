# diceware

This program is a fun way to learn how to build executables in Go!

It generates strong passphrases using the Diceware algorithm, a method that
relies on rolling dice to select random words from a word list.

For more information on Diceware, check out
[The Diceware Passphrase Home Page](https://theworld.com/~reinhold/diceware.html).

> **Disclaimer:**
>
> The Diceware method explicitly discourages the use of software-based random
> number generators.
> This implementation, however, utilizes the random number generator provided
> by the Go crypto package, which is a widely-accepted source of
> cryptographic-grade randomness.
> While this approach offers a reasonable level of security, it's important to
> recognize that physical dice remain the most secure method for Diceware
> passphrase generation.

## Usage

### Basic Usage

To generate a new passphrase, simply run the diceware executable. It will print the passphrase to your terminal.

### Additional Options

You can customize your passphrase generation with the following options:

- **available-languages**:
  List all supported languages and exit.
- **digits**:
  Force the inclusion of digits in the passphrase.
- **lang _value_**:
  Specify the language to use for word selection.
- **length _value_**:
  Set the desired number of words in the passphrase.
- **lower**:
  Force all words to be lowercase.
- **symbols**:
  Add delimiting symbols between words.
- **upper**:
  Force all words to be uppercase.

By default, the passphrase will be in Pascal case
(first letter of each word capitalized)
and may include digits without delimiters.

## Copyright

Copyright © 2024 Jose Ignacio Croce Busquets

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program.
If not, see <https://www.gnu.org/licenses/>.

Lista di parole Diceware in Italiano
Copyright (C) 2007, ..., 2019 [Tarin Gamberini](http://www.taringamberini.com)

French wodlist by [Mattieu Weber](http://weber.fi.eu.org/index.shtml.en#projects)
