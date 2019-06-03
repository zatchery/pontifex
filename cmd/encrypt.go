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

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypt called")
		var plaintextFlag, _ = cmd.Flags().GetString("plaintext")
		var plainfileFlag, _ = cmd.Flags().GetString("plainfile")
		var keystreamFlag, _ = cmd.Flags().GetString("keystream")
		var keystreamfileFlag, _ = cmd.Flags().GetString("keystreamfile")
		var verbose, _ = cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("plaintext: %s,\nplainfile: %s,\nkeystream: %s,\nkeystreamfile: %s\n", plaintextFlag, plainfileFlag, keystreamFlag, keystreamfileFlag)
		}

		//Read in plaintext and keystream
		var cypherText = encrypt(plainfileFlag, keystreamFlag)
		//if output file specifed read to file otherwise
		fmt.Println(cypherText)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringP("plaintext", "p", "", "Type the message you want to encrypt")
	encryptCmd.Flags().StringP("plainfile", "f", "", "The file with the message you want to encrypt")
	encryptCmd.Flags().StringP("keystream", "k", "", "The name of the keystream you want to use in your <keyname>.ptfx file")
	encryptCmd.Flags().StringP("keystreamfile", "s", "", "The file with the keystreams you want to use. By default it will look for a default keyfile at $HOME/.pontifex")
	encryptCmd.Flags().StringP("output", "o", "", "The name of the output cypher text file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func encrypt(plaintext string, key string) string {
	return "HDEDKEAPOI EWSDFKJ DIUWKF"
}
