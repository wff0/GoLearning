package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	default:
		fmt.Println("It's before noon")
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an int")
		case bool:
			fmt.Println("I'm a bool")
		case string:
			fmt.Println("I'm a string")
		case byte:
			fmt.Println("I'm a byte")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmi(true)
	whatAmi(1)
	whatAmi("string")
	whatAmi(byte('b'))
	whatAmi('b')
	whatAmi(4.3)
}
