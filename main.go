package main

import (
	"fmt"
	"os"
)

func main() {
	if username, password := "musa", "abu"; len(os.Args) <= 1 {
		fmt.Println("Usage: [username] [password]")
	} else if os.Args[1] != username {
		fmt.Println("access denied for username :", os.Args[1])
	} else if os.Args[2] != password {
		fmt.Println("access denied for password:", os.Args[2])
	} else {
		fmt.Println("access granted to:", os.Args[1])
	}
}
