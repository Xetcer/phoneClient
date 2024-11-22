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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints the status of the server.",
	Long: `This command prints information about the
	status of the phone book server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")
		// создаем запрос для получения статуса сервера
		URL := "http://" + SERVER + ":" + PORT + "/status"

		// Отправляем запрос на сервер
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// проверяем код состояния HTTP
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}
		// если все ОК, считываем данные
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		/* Если все в порядке, то мы считываем весь текст ответа сервера,
		который представляет собой байтовый срез, и выводим его на экран в виде строки*/
		fmt.Print(string(responseData))

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
