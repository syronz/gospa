package cli

import "fmt"

type CLI struct{}

func (c *CLI) printUsage() {
	fmt.Println("Usage: ")
	fmt.Println(" init -port PORT - specify port when try to create the toml file")
	fmt.Println(" run - start the webserver")
}
