package pronunciation

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
	alphabet "github.com/eminmuhammadi/azlang/system/alphabet"
)

// Outputs a pronounciation of the given word or sentence
// TODO: Needs improvement
func Phonetize(input string, originAlphabet alphabet.Alphabet) string {
	pronounciation := ""

	origin := alphabet.NewAlphabet(originAlphabet)
	transformations := origin.WordToStandardIDs(input)

	for _, transformation := range transformations {
		if val, ok := STANDARD_ID_LIST[transformation.Standard_ID]; ok {
			pronounciation += string(val.Pronunciation)
		} else {
			pronounciation += transformation.Letter
		}
	}

	return pronounciation
}
