package main

import (
	"fmt"
	"monkey/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	name := ""
	if user.Name != "" {
		name = user.Name
	} else {
		name = user.Username
	}

	fmt.Printf("Hello %s! This is the monkey programming language!\n", name)
	fmt.Printf("Fell free to type in commands\n")

	// 接入标准输入输出
	repl.Start(os.Stdin, os.Stdout)
}
