package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	fmt.Println(toLowerCase("Ass"))
}
