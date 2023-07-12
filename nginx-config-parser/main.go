package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tufanbarisyildirim/gonginx"
)

func main() {
	filePath := "path/to/nginx/default.conf"

	// Read the Nginx configuration file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading Nginx configuration file: %v", err)
		return
	}

	// Parse the Nginx configuration
	conf, err := gonginx.ParseBytes(content)
	if err != nil {
		fmt.Printf("Error parsing Nginx configuration: %v", err)
		return
	}

	// Access the parsed configuration
	for _, block := range conf.Blocks {
		fmt.Printf("Block: %s %s\n", block.Directive, block.Parameters)

		// Access block directives
		for _, directive := range block.Directives {
			fmt.Printf(" - Directive: %s %s\n", directive.Name, directive.Value)
		}
	}
}
