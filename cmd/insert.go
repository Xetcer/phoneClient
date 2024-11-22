/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new entries",
	Long: `This command inserts new data to the 
	Phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")
		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty!")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Name is emty")
			return
		}

		surename, _ := cmd.Flags().GetString("surename")
		if surename == "" {
			fmt.Println("Surename is emty")
			return
		}

		URL := "http://" + SERVER + ":" + PORT + "/insert/"
		URL = URL + "/" + name + "/" + surename + "/" + number

		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(responseData))

	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "Name value")
	insertCmd.Flags().StringP("surename", "s", "", "Sureame value")
	insertCmd.Flags().StringP("tel", "t", "", "Telephone value")
}
