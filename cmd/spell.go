package cmd

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
	"fmt"

	alphabet "github.com/eminmuhammadi/azlang/system/alphabet"
	pronunciation "github.com/eminmuhammadi/azlang/system/pronunciation"
	cli "github.com/urfave/cli/v2"
)

func Phonetize() *cli.Command {
	return &cli.Command{
		Name:    "phonetize",
		Aliases: []string{"p"},
		Usage:   "Writes a pronounciation of the given word or sentence",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input string",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "origin",
				Aliases:  []string{"o"},
				Usage:    "Origin of writing system",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			input := ctx.String("input")

			var originAlphabet alphabet.Alphabet
			origin := ctx.String("origin")

			switch origin {
			case "OFFICIAL_LATIN":
				originAlphabet = alphabet.OFFICIAL_LATIN
			case "OLD_LATIN":
				originAlphabet = alphabet.OLD_LATIN
			case "CYRILLIC":
				originAlphabet = alphabet.CYRILLIC
			default:
				return fmt.Errorf("invalid writing system: %s", origin)
			}

			output := pronunciation.Phonetize(input, originAlphabet)

			fmt.Println(output)

			return nil
		},
	}
}
