// Copyright © 2019 Zachary Morin <me@zatchery.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Take plaintext and ecrypt it using a specified key",
	Long: `Example usages:
	pontifex encrypt --plaintext <Message to encrypt> --keystream <Keystream named in the ~/.pontifex file>
	pontifex encrypt --p <Message to encrypt> --k <Keystream named in the ~/.pontifex file>
	pontifex encrypt --keystreamfile <Path to non default keystream file> -p <Message to encrypt>


Pontifex also known as the Solitaire Cypher (https://en.wikipedia.org/wiki/Solitaire_(cipher))
is a way to use a deck of cards to communicate securely.`,
	Run: func(cmd *cobra.Command, args []string) {
		var plaintextFlag, _ = cmd.Flags().GetString("plaintext")
		var plainfileFlag, _ = cmd.Flags().GetString("plainfile")
		var keystreamFlag, _ = cmd.Flags().GetString("keystream")
		var keystreamfileFlag, _ = cmd.Flags().GetString("keystreamfile")
		var verbose, _ = cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("plaintext: %s,\nplainfile: %s,\nkeystream: %s,\nkeystreamfile: %s\n", plaintextFlag, plainfileFlag, keystreamFlag, keystreamfileFlag)
		}

		var plaintext = ""
		if len(args) == 0 || len(args[0]) == 0 {
			//Use the -f option
			if verbose {
				fmt.Printf("Reading plaintext from file: %s\n", plainfileFlag)
			}
			plaintext = "Readintextfromfile"
		} else {
			plaintext = args[0]
		}

		if verbose {
			fmt.Printf("plaintext: %s\nkeystream: %s\n", plaintext, keystreamFlag)
		}

		keyfile := "Default"
		//Read in key
		if len(keystreamfileFlag) > 0 {
			//Use the specified keystream file
			keyfile = keystreamfileFlag
		}

		//Read in plaintext and keystream
		var keyStream = readKey(keyfile, keystreamFlag, verbose)
		var cypherText, _ = encrypt(plaintext, keyStream, verbose)
		fmt.Printf("%s", cypherText)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringP("plaintext", "p", "", "Type the message you want to encrypt")
	encryptCmd.Flags().StringP("plainfile", "f", "", "The file with the message you want to encrypt")
	encryptCmd.Flags().StringP("keystream", "k", "", "The name of the keystream you want to use in your <keyname>.ptfx file")
	encryptCmd.Flags().StringP("keystreamfile", "s", "", "The file with the keystreams you want to use. By default it will look for a default keyfile at $HOME/.pontifex")
	encryptCmd.Flags().StringP("output", "o", "", "The name of the output cypher text file")
}

func encrypt(plaintext string, keystream []string, verbose bool) (string, []string) {
	if verbose {
		fmt.Printf("Encrypting: %s\n", plaintext)
		fmt.Println("Using Keystream: ", keystream)
	}
	return getCypherText(plaintext, keystream, verbose)
}
