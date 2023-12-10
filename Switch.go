package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool value")
		case int:
			fmt.Println("Int value")
		default:
			fmt.Printf("I dont know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
