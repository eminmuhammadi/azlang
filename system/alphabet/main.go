package alphabet

//  Copyright 2022- Emin Muhammadi and contributors
//
//  Licensed under the The GNU Affero General Public License,
//  Version 3.0 (the "License"); you may not use this file except
//  in compliance with the License. You may obtain a copy
//  of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.en.html
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the
//  specific language governing permissions and limitations
//  under the License.

import (
	"strings"
)

// Unsupported letters
const UNSUPPORTED_STANDARD_ID = "AZLANG-00"

// Generate new alphabet
func NewAlphabet(alphabet Alphabet) *Alphabet {
	return &Alphabet{
		Name:        alphabet.Name,
		Description: alphabet.Description,
		Letters:     alphabet.Letters,
	}
}

// Transformation struct
type Transformation struct {
	Standard_ID string
	Letter      string
	IsUppercase bool
}

// Is letter uppercase
func IsUppercase(letter string) bool {
	return letter == strings.ToUpper(letter)
}

// Transfrom word to IDs
func (alphabet *Alphabet) WordToStandardIDs(input string) []Transformation {
	var found bool
	transformations := []Transformation{}

	for _, letter := range input {
		found = false

		for _, alphabetLetter := range alphabet.Letters {
			if string(letter) == alphabetLetter.Uppercase || string(letter) == alphabetLetter.Lowercase {
				transformations = append(transformations, Transformation{
					Standard_ID: alphabetLetter.Standard_ID,
					Letter:      string(letter),
					IsUppercase: IsUppercase(string(letter)),
				})

				found = true
				break
			}
		}

		if !found {
			transformations = append(transformations, Transformation{
				Standard_ID: UNSUPPORTED_STANDARD_ID,
				Letter:      string(letter),
				IsUppercase: false,
			})
		}
	}

	return transformations
}

// Transfrom IDs to word
func (alphabet *Alphabet) StandardIDsToWord(transformations []Transformation) string {
	word := ""

	// Add letters
	i := 0
	for _, item := range transformations {
		for _, alphabetLetter := range alphabet.Letters {
			if item.Standard_ID == alphabetLetter.Standard_ID {
				if item.IsUppercase {
					word += alphabetLetter.Uppercase
				} else {
					word += alphabetLetter.Lowercase
				}

				i++
				break
			}
		}

		if item.Standard_ID == UNSUPPORTED_STANDARD_ID {
			word += item.Letter
		}
	}

	if alphabet.IsRTL {
		// reverse string
		return reverse(word)
	}

	return word
}

// Reverse string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
