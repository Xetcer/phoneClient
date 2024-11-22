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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an entry",
	Long: `This commands deletes an existing entry from
	the phone book application given a phone number.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty!")
			return
		}

		// Создаем запрос
		URL := "http://" + SERVER + ":" + PORT + "/delete/" + number

		// отправляем запрос
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Проверяем HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status coder:", data.StatusCode)
			return
		}

		// Считываем данные
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(string(responseData))

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	//Это локальный флаг, доступный только для команды delete
	deleteCmd.Flags().StringP("tel", "t", "", "Telephone number to delete")
}
