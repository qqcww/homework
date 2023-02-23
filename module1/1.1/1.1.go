package main

import "fmt"

func main() {
	strings := []string{"I", "am", "stupid", "and", "weak"}
	for key, value := range strings {
		if value == "stupid" {
			strings[key] = "smart"
		}
		if value == "weak" {
			strings[key] = "strong"
		}
	}
	fmt.Println(strings)
}
