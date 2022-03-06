/*
Copyright © 2022 yixy <youzhilane01@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/yixy/arithmetic-bot/common/rand"

	"github.com/spf13/cobra"
)

var scope, num *int

// plusCmd represents the plus command
var plusCmd = &cobra.Command{
	Use:   "plus",
	Short: "addition formula",
	Long:  `generate addition formula`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= *num; i++ {
			a := rand.GetInt(*scope)
			b := rand.GetInt(*scope - a)
			fmt.Printf("%d + %d =  \t\t", a, b)
			if i%*column == 0 {
				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(plusCmd)

	column = plusCmd.Flags().IntP("column", "c", 3, "column of the result.")
	scope = plusCmd.Flags().IntP("scope", "s", 100, "Specify the max value of result. ")
	num = plusCmd.Flags().IntP("num", "n", 50, "Specify the number of formula.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// plusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
