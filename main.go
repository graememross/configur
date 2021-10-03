package main

import "gitlab.rossathome.co.uk/graeme/configur/configur"

func main() {
	config := configur.ParseConfig("tester")
	config.Print()
}
