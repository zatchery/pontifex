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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt called")

		var cyphertextFlag, _ = cmd.Flags().GetString("cyphertext")
		var cypherfileFlag, _ = cmd.Flags().GetString("cypherfile")
		var keystreamFlag, _ = cmd.Flags().GetString("keystream")
		var keystreamfileFlag, _ = cmd.Flags().GetString("keystreamfile")
		var verbose, _ = cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("cyphertext: %s,\ncypherfile: %s,\nkeystream: %s,\nkeystreamfile: %s\n", cyphertextFlag, cypherfileFlag, keystreamFlag, keystreamfileFlag)
		}

		//Either -c or -f

		//Read in key

		//Read in cyphertext and keystream
		var cypherText = decrypt(cypherfileFlag, keystreamFlag)
		fmt.Println(cypherText)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringP("cyphertext", "c", "", "Type the message you want to encrypt")
	decryptCmd.Flags().StringP("cypherfile", "f", "", "The file with the message you want to encrypt")
	decryptCmd.Flags().StringP("keystream", "k", "", "The name of the keystream you want to use in your <keyname>.ptfx file")
	decryptCmd.Flags().StringP("keystreamfile", "s", "", "The file with the keystreams you want to use. By default it will look for a default keyfile at $HOME/.pontifex")
	decryptCmd.Flags().StringP("output", "o", "", "The name of the output plain text file")
}

func decrypt(cyphertext string, keystream string) string {
	return "PONTIFICUS CRYPTUS TODOS"
}
