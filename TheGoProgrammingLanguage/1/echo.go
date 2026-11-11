package main

import (
	//`:`	"fmt"
	"strings"
)

func manual(args []string) string {
	s, sep := "", ""
	for _, i := range args[1:] {
		s += sep + i
		sep = " "
	}
	return s
}

/*
func experiments() {
	fmt.Printf(args[0])          // 1.1
	for i, v := range args[1:] { // 1.2
		fmt.Printf("%d\t%s\n", i+1, v)
	}
}
*/

func joinfunc(args []string) string {
	return strings.Join(args[1:], " ")
}
