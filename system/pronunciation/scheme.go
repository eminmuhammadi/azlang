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

// Main folder for local soundsx
const FOLDER_TO_SOUNDS string = "./../../sounds/"

type Standard struct {
	Name          string
	URL           string
	LocalFile     string
	Pronunciation string
}

var STANDARD_ID_LIST = map[string]Standard{
	"AZLANG-01": {
		Name: "Open back unrounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/e5/Open_back_unrounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-01.ogg",
		Pronunciation: "ɑ",
	},
	"AZLANG-02": {
		Name: "Voiced bilabial plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/2/2c/Voiced_bilabial_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-02.ogg",
		Pronunciation: "b",
	},
	"AZLANG-03": {
		Name: "Voiced postalveolar affricate",
		// This work has been released into the public domain by its author, Octane at English Wikipedia.
		// This applies worldwide. In some countries this may not be legally possible; if so:
		// Octane grants anyone the right to use this work for any purpose, without any conditions,
		// unless such conditions are required by law.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/e6/Voiced_palato-alveolar_affricate.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-03.ogg",
		Pronunciation: "dʒ",
	},
	"AZLANG-04": {
		Name: "Voiceless postalveolar affricate",
		URL:  "https://upload.wikimedia.org/wikipedia/commons/9/97/Voiceless_palato-alveolar_affricate.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-04.ogg",
		Pronunciation: "tʃ",
	},
	"AZLANG-05": {
		Name: "Voiced alveolar plosives",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/0/01/Voiced_alveolar_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-05.ogg",
		Pronunciation: "d",
	},
	"AZLANG-06": {
		Name: "Close-mid front unrounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/6/6c/Close-mid_front_unrounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-06.ogg",
		Pronunciation: "e",
	},
	"AZLANG-07": {
		Name: "Near-open front unrounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/c/c9/Near-open_front_unrounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-07.ogg",
		Pronunciation: "æ",
	},
	"AZLANG-08": {
		Name: "Voiceless labiodental fricative",
		URL:  "https://upload.wikimedia.org/wikipedia/commons/c/c7/Voiceless_labio-dental_fricative.ogg",
		// This file is made available under the Creative Commons CC0 1.0 Universal Public Domain Dedication.
		LocalFile:     "AZLANG-08.ogg",
		Pronunciation: "f",
	},
	"AZLANG-09": {
		Name: "Voiced palatal plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/1/1d/Voiced_palatal_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-09.ogg",
		Pronunciation: "ɟ",
	},
	"AZLANG-10": {
		Name: "Voiced velar fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/4/47/Voiced_velar_fricative.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-10.ogg",
		Pronunciation: "ɣ",
	},
	"AZLANG-11": {
		Name: "Voiceless glottal fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/d/da/Voiceless_glottal_fricative.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-11.ogg",
		Pronunciation: "h",
	},
	"AZLANG-12": {
		Name: "Voiceless velar fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/0/0f/Voiceless_velar_fricative.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-12.ogg",
		Pronunciation: "x",
	},
	"AZLANG-13": {
		Name: "Close back unrounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/e8/Close_back_unrounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-13.ogg",
		Pronunciation: "ɯ",
	},
	"AZLANG-14": {
		Name: "Close front unrounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/9/91/Close_front_unrounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-14.ogg",
		Pronunciation: "i",
	},
	"AZLANG-15": {
		Name: "Voiced palato-alveolar fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/3/30/Voiced_palato-alveolar_sibilant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-15.ogg",
		Pronunciation: "ʒ",
	},
	"AZLANG-16": {
		Name: "Voiceless velar plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/e3/Voiceless_velar_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-16.ogg",
		Pronunciation: "k",
	},
	"AZLANG-17": {
		Name: "Voiceless palatal plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/5/5d/Voiceless_palatal_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-17.ogg",
		Pronunciation: "c",
	},
	"AZLANG-18": {
		Name: "Voiced velar plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/1/12/Voiced_velar_plosive_02.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-18.ogg",
		Pronunciation: "g",
	},
	"AZLANG-19": {
		Name: "Voiced dental, alveolar and postalveolar lateral approximants",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/b/bc/Alveolar_lateral_approximant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-19.ogg",
		Pronunciation: "l",
	},
	"AZLANG-20": {
		Name: "Voiced bilabial nasal",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/a/a9/Bilabial_nasal.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-20.ogg",
		Pronunciation: "m",
	},
	"AZLANG-21": {
		Name: "Voiced dental, alveolar and postalveolar nasals",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/2/29/Alveolar_nasal.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-21.ogg",
		Pronunciation: "n",
	},
	"AZLANG-22": {
		Name: "Voiced velar nasal",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/3/39/Velar_nasal.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-22.ogg",
		Pronunciation: "ŋ",
	},
	"AZLANG-23": {
		Name: "Close-mid back rounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/8/84/Close-mid_back_rounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-23.ogg",
		Pronunciation: "o",
	},
	"AZLANG-24": {
		Name: "Close-mid front rounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/5/53/Close-mid_front_rounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-24.ogg",
		Pronunciation: "ø",
	},
	"AZLANG-25": {
		Name: "Voiceless bilabial plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/5/51/Voiceless_bilabial_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-25.ogg",
		Pronunciation: "p",
	},
	"AZLANG-26": {
		Name: "Voiced alveolar trills",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/c/ce/Alveolar_trill.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-26.ogg",
		Pronunciation: "r",
	},
	"AZLANG-27": {
		Name: "Voiceless alveolar sibilant",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/a/ac/Voiceless_alveolar_sibilant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-27.ogg",
		Pronunciation: "s",
	},
	"AZLANG-28": {
		Name: "Voiceless postalveolar fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/c/cc/Voiceless_palato-alveolar_sibilant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-28.ogg",
		Pronunciation: "ʃ",
	},
	"AZLANG-29": {
		Name: "Voiceless alveolar plosive",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/0/02/Voiceless_alveolar_plosive.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-29.ogg",
		Pronunciation: "t",
	},
	"AZLANG-30": {
		Name: "Close back rounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/5/5d/Close_back_rounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-30.ogg",
		Pronunciation: "u",
	},
	"AZLANG-31": {
		Name: "Close front rounded vowel",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/ea/Close_front_rounded_vowel.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-31.ogg",
		Pronunciation: "y",
	},
	"AZLANG-32": {
		Name: "Voiced labiodental fricative",
		URL:  "https://upload.wikimedia.org/wikipedia/commons/4/42/Voiced_labio-dental_fricative.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-32.ogg",
		Pronunciation: "v",
	},
	"AZLANG-33": {
		Name: "Palatal approximant",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/e/e8/Palatal_approximant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-33.ogg",
		Pronunciation: "j",
	},
	"AZLANG-34": {
		Name: "Voiced alveolar fricative",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/c/c0/Voiced_alveolar_sibilant.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-34.ogg",
		Pronunciation: "z",
	},
	"AZLANG-35": {
		Name: "Glottal stop",
		// Permission is granted to copy, distribute and/or modify this document
		// under the terms of the GNU Free Documentation License, Version 1.2 or
		// any later version published by the Free Software Foundation; with no
		// Invariant Sections, no Front-Cover Texts, and no Back-Cover Texts. A
		// copy of the license is included in the section entitled GNU Free
		// Documentation License.
		URL: "https://upload.wikimedia.org/wikipedia/commons/4/4d/Glottal_stop.ogg",
		// This file is licensed under the Creative Commons Attribution-Share Alike 3.0 Unported license.
		LocalFile:     "AZLANG-35.ogg",
		Pronunciation: "ʔ",
	},
}
