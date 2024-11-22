package main

import "fmt"

const EnglishPrefix = "Hello, "

func Hello(s string) string {
	if s == "" {
		return EnglishPrefix + "world"
	}
	return EnglishPrefix + s
}
func main() {
	fmt.Println(Hello("you"))
}
