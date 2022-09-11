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
	"testing"
)

const EXAMPLE_OFFICIAL_LATIN = "Arazı ayırdılar,\nQum ilə doyurdular.\nMən səndən ayrılmazdım,\nZülm ilə ayırdılar."
const EXAMPLE_CYRILLIC = "Аразы ајырдылар,\nГум илә дојурдулар.\nМән сәндән ајрылмаздым,\nЗүлм илә ајырдылар."
const EXAMPLE_OLD_LATIN = "Arazь ajьrdьlar,\nQum ilə dojurdular.\nMən səndən ajrьlmazdьm,\nZylm ilə ajьrdьlar."

// Official Latin alphabet to Cyrillic alphabet
func Test_OfficialLatin_To_Cyrillic(t *testing.T) {
	official_latin := NewAlphabet(OFFICIAL_LATIN)
	cyrillic := NewAlphabet(CYRILLIC)

	IDs := official_latin.WordToStandardIDs(EXAMPLE_OFFICIAL_LATIN)
	transformed_sentence := cyrillic.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_CYRILLIC {
		t.Errorf("Expected %s, got %s", EXAMPLE_CYRILLIC, transformed_sentence)
	}
}

// Official Latin alphabet to Old Latin alphabet
func Test_OfficialLatin_To_OldLatin(t *testing.T) {
	official_latin := NewAlphabet(OFFICIAL_LATIN)
	old_latin := NewAlphabet(OLD_LATIN)

	IDs := official_latin.WordToStandardIDs(EXAMPLE_OFFICIAL_LATIN)
	transformed_sentence := old_latin.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_OLD_LATIN {
		t.Errorf("Expected %s, got %s", EXAMPLE_OLD_LATIN, transformed_sentence)
	}
}

// Cyrillic alphabet to Official Latin alphabet
func Test_Cyrillic_To_OfficialLatin(t *testing.T) {
	cyrillic := NewAlphabet(CYRILLIC)
	official_latin := NewAlphabet(OFFICIAL_LATIN)

	IDs := cyrillic.WordToStandardIDs(EXAMPLE_CYRILLIC)
	transformed_sentence := official_latin.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_OFFICIAL_LATIN {
		t.Errorf("Expected %s, got %s", EXAMPLE_OFFICIAL_LATIN, transformed_sentence)
	}
}

// Cyrillic alphabet to Official Latin alphabet
func Test_Cyrillic_To_OldLatin(t *testing.T) {
	cyrillic := NewAlphabet(CYRILLIC)
	old_latin := NewAlphabet(OLD_LATIN)

	IDs := cyrillic.WordToStandardIDs(EXAMPLE_CYRILLIC)
	transformed_sentence := old_latin.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_OLD_LATIN {
		t.Errorf("Expected %s, got %s", EXAMPLE_OLD_LATIN, transformed_sentence)
	}
}

// Old Latin alphabet to Official Latin alphabet
func Test_OldLatin_To_OfficialLatin(t *testing.T) {
	old_latin := NewAlphabet(OLD_LATIN)
	official_latin := NewAlphabet(OFFICIAL_LATIN)

	IDs := old_latin.WordToStandardIDs(EXAMPLE_OLD_LATIN)
	transformed_sentence := official_latin.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_OFFICIAL_LATIN {
		t.Errorf("Expected %s, got %s", EXAMPLE_OFFICIAL_LATIN, transformed_sentence)
	}
}

// Old Latin alphabet to Official Latin alphabet
func Test_OldLatin_To_Cyrillic(t *testing.T) {
	old_latin := NewAlphabet(OLD_LATIN)
	cyrillic := NewAlphabet(CYRILLIC)

	IDs := old_latin.WordToStandardIDs(EXAMPLE_OLD_LATIN)
	transformed_sentence := cyrillic.StandardIDsToWord(IDs)

	if transformed_sentence != EXAMPLE_CYRILLIC {
		t.Errorf("Expected %s, got %s", EXAMPLE_CYRILLIC, transformed_sentence)
	}
}
