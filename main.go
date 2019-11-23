package main

import "fmt"

// cmd test code
// run go install jvm-in-go will create jvm-in-go.exe in classpath/bin/
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s args: %v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
