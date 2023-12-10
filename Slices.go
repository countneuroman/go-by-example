package main

import (
	"fmt"
)

func main() {
	s := make([]string, 5)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	fmt.Println("set:", s)
	fmt.Println("get:", s[4])

	s = append(s, "f")
	fmt.Println("apd:", s)

	l := s[2:4]
	fmt.Println("sl1:", l)
}
