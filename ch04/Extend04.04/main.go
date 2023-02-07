package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{"305": "Sue", "204": "Bob", "631": "Jake", "073": "Tracy"}
}

func getUser(id string) (string, bool) {
	userList := getUsers()
	userName, status := userList[id]
	return userName, status
}

func main() {
	for {
		if len(os.Args) < 2 {
			fmt.Println("No valid ID. Please ensure you already input ID")
			os.Exit(1)
		}
		for i := 1; i < len(os.Args); i++ {
			//fmt.Println(os.Args[i])
			userName, isExisted := getUser(os.Args[i])
			if isExisted {
				fmt.Printf("Hello, %s (%v)\n", userName, os.Args[i])
			} else {
				fmt.Printf("No user id (%v)\n", os.Args[i])
			}
		}
		os.Exit(0)
	}
}
