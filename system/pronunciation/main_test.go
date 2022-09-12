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
	"testing"

	alphabet "github.com/eminmuhammadi/azlang/system/alphabet"
)

const EXAMPLE_OFFICIAL_LATIN = "Salam, Dünya!"
const EXAMPLE_CYRILLIC = "Салам, Дүнја!"
const EXAMPLE_OLD_LATIN = "Salam, Dynja!"
const OUTPUT = "sɑlɑm, dynjɑ!"

// Test pronunciation of the given word in Official Latin
func Test_Official_Latin(t *testing.T) {
	pronunciation := Phonetize(EXAMPLE_OFFICIAL_LATIN, alphabet.OFFICIAL_LATIN)
	if pronunciation != OUTPUT {
		t.Errorf("Expected %s, got %s", OUTPUT, pronunciation)
	}
}

// Test pronunciation of the given word in Old Latin
func Test_Old_Latin(t *testing.T) {
	pronunciation := Phonetize(EXAMPLE_OLD_LATIN, alphabet.OLD_LATIN)
	if pronunciation != OUTPUT {
		t.Errorf("Expected %s, got %s", OUTPUT, pronunciation)
	}
}

// Test pronunciation of the given word in Cyrillic
func Test_Cyrillic(t *testing.T) {
	pronunciation := Phonetize(EXAMPLE_CYRILLIC, alphabet.CYRILLIC)
	if pronunciation != OUTPUT {
		t.Errorf("Expected %s, got %s", OUTPUT, pronunciation)
	}
}
