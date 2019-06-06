// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a cryptext",
	Long: `Decrypt a cyptext.. For example:
	pontifex decrypt <-k, -s, -o, -f, TEXT TO DECRYPT>

If you pass a TEXT to decrypt pontifex will ignore the -f option`,
	Run: func(cmd *cobra.Command, args []string) {
		var cypherfileFlag, _ = cmd.Flags().GetString("cypherfile")
		var keystreamFlag, _ = cmd.Flags().GetString("keystream")
		var keystreamfileFlag, _ = cmd.Flags().GetString("keystreamfile")
		var verbose, _ = cmd.Flags().GetBool("verbose")

		var cyphertext = ""
		if len(args) == 0 || len(args[0]) == 0 {
			//Use the -f option
			if verbose {
				fmt.Printf("Reading cyphertext from file: %s\n", cypherfileFlag)
			}
			cyphertext = "Readintextfromfile"
		} else {
			cyphertext = args[0]
		}

		if verbose {
			fmt.Printf("cyphertext: %s\nkeystream: %s\nkeystreamfile: %s\n", cyphertext, keystreamFlag, keystreamfileFlag)
		}

		keyfile := "Default"
		//Read in key
		if len(keystreamfileFlag) > 0 {
			//Use the specified keystream file
			keyfile = keystreamfileFlag
		}

		//Read in cyphertext and keystream
		var keyStream = readKey(keyfile, keystreamFlag, verbose)
		var plaintext = decrypt(cyphertext, keyStream, verbose)
		fmt.Println(plaintext)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringP("cypherfile", "f", "", "The file with the message you want to encrypt")
	decryptCmd.Flags().StringP("keystream", "k", "", "The name of the keystream you want to use in your <keyname>.ptfx file")
	decryptCmd.MarkFlagRequired("keystream")
	decryptCmd.Flags().StringP("keystreamfile", "s", "", "The file with the keystreams you want to use. By default it will look for a default keyfile at $HOME/.pontifex")
	decryptCmd.Flags().StringP("output", "o", "", "The name of the output plain text file")
}

func decrypt(cyphertext string, keystream []string, verbose bool) string {
	if verbose {
		fmt.Printf("Decrypting: %s\n", cyphertext)
	}
	translate(cyphertext, keystream)
	return "PONTIFICUS CRYPTUS TODOS"
}
