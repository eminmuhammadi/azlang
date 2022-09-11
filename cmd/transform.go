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

	"github.com/eminmuhammadi/azlang/system"
	"github.com/eminmuhammadi/azlang/system/alphabet"
	cli "github.com/urfave/cli/v2"
)

func Transform() *cli.Command {
	return &cli.Command{
		Name:    "transform",
		Aliases: []string{"t"},
		Usage:   "Transform different type of writing systems",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input string",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "from",
				Aliases:  []string{"f"},
				Usage:    "Transform from",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "to",
				Aliases:  []string{"t"},
				Usage:    "Transform to",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			input := ctx.String("input")

			var fromAlphabet alphabet.Alphabet
			from := ctx.String("from")

			var toAlphabet alphabet.Alphabet
			to := ctx.String("to")

			switch from {
			case "OFFICIAL_LATIN":
				fromAlphabet = alphabet.OFFICIAL_LATIN
			case "OLD_LATIN":
				fromAlphabet = alphabet.OLD_LATIN
			case "CYRILLIC":
				fromAlphabet = alphabet.CYRILLIC
			default:
				return fmt.Errorf("invalid writing system: %s", from)
			}

			switch to {
			case "OFFICIAL_LATIN":
				toAlphabet = alphabet.OFFICIAL_LATIN
			case "OLD_LATIN":
				toAlphabet = alphabet.OLD_LATIN
			case "CYRILLIC":
				toAlphabet = alphabet.CYRILLIC
			default:
				return fmt.Errorf("invalid writing system: %s", to)
			}

			output := system.Transform(input, fromAlphabet, toAlphabet)

			// Print output
			fmt.Println(output)

			return nil
		},
	}
}
