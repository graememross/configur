package main

import "github.com/graememross/configur/configur"

func main() {
	config := configur.ParseConfig("tester")
	config.Print()
}
