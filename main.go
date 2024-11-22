/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
go run main.go list
go run main.go status
go run main.go search --tel 0800123456
go run main.go search --tel 0800
go run main.go delete --tel 2101112223
go run main.go insert -n Michalis -s Tsoukalos -t 2101112223
*/
package main

import "phoneClient/cmd"

func main() {
	cmd.Execute()
}
