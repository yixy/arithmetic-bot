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

	"github.com/spf13/cobra"
	"github.com/yixy/arithmetic-bot/common/rand"
)

var level *int

// expCmd represents the exp command
var expCmd = &cobra.Command{
	Use:   "exp",
	Short: "generate expirement by level",
	Long:  `generate expirements by level.`,
	Run: func(cmd *cobra.Command, args []string) {
		if *level <= 2 {
			//1 big-non-negative + small-non-negative = r, r belongs [0,10]
			//2 non-negative + non-negative = r, r belongs [0,10]
			for i := 1; i <= 20; i++ {
				a := rand.GetInt(10)
				b := rand.GetInt(10 - a)
				if *level == 1 && b > a {
					tmp := a
					a = b
					b = tmp
				}
				fmt.Printf("%d + %d =  \t\t", a, b)
				if i%*column == 0 {
					fmt.Println()
				}
			}
		} else if *level == 3 {
			//	3 non-negative + non-negative = r, r belongs [0,18]
			for i := 1; i <= 20; i++ {
				a := rand.GetInt(18)
				b := rand.GetInt(18 - a)
				fmt.Printf("%d + %d =  \t\t", a, b)
				if i%*column == 0 {
					fmt.Println()
				}
			}
		} else if *level == 4 {
			//	4 multiples of 10 + multiples of 10 = r, r belongs [0,100]
			for i := 1; i <= 20; i++ {
				a := rand.GetInt(100)
				b := rand.GetInt(100 - a)
				fmt.Printf("%d + %d =  \t\t", a/10*10, b/10*10)
				if i%*column == 0 {
					fmt.Println()
				}
			}
		} else if *level == 5 {
			//5 multiples of 10 + non-negative = r, r belongs [0,100]
			for i := 1; i <= 20; i++ {
				a := rand.GetInt(100)
				b := rand.GetInt(100 - a)
				fmt.Printf("%d + %d =  \t\t", a/10*10, b)
				if i%*column == 0 {
					fmt.Println()
				}
			}
		} else if *level == 6 {
			//	6 non-negative + non-negative = r, r belongs [0,100]`)
			for i := 1; i <= 20; i++ {
				a := rand.GetInt(100)
				b := rand.GetInt(100 - a)
				fmt.Printf("%d + %d =  \t\t", a, b)
				if i%*column == 0 {
					fmt.Println()
				}
			}
		} else {
			fmt.Println("level is undefined.")
		}
	},
}

func init() {
	rootCmd.AddCommand(expCmd)

	level = expCmd.Flags().IntP("level", "l", 1, `Specify the level of formula. The detail of level:
	1 big-non-negative + small-non-negative = r, r belongs [0,10]
	2 non-negative + non-negative = r, r belongs [0,10]
	3 on-negative + non-negative = r, r belongs [0,18]
	4 multiples of 10 + multiples of 10 = r, r belongs [0,100]
	5 multiples of 10 + non-negative = r, r belongs [0,100]
	6 non-negative + non-negative = r, r belongs [0,100]
`)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
