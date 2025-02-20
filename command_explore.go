package main

import "fmt"

func commandExplore(config *config, params ...string) error {
	for _, param := range params {
		fmt.Println(param)
	}
	return nil
}
