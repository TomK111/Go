package main

import (
	"fmt"
)

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2

	fmt.Println(average)

	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)

	fmt.Println("sum :", 3+2)        // sum: int
	fmt.Println("sum :", 3+2.5)      // sum: float64
	fmt.Println("dif :", 3-1)        // difference: int
	fmt.Println("dif :", 3-0.5)      // difference: float64
	fmt.Println("prod :", 4*5)       // product: int
	fmt.Println("prod :", 4.5*10)    //product: float64
	fmt.Println("quot: ", 8/2)       // quotient: int
	fmt.Println("quot: ", 10.5/5)    // quotient: float64
	fmt.Println("remainder: ", 10%5) // remainder: int (only works for ints)

	f := 8.0
	fmt.Println("remainder: ", int(f)%3)

	a := 6.0
	fmt.Println("remainder: ", int(a)%5)

	var ratios float64 = 3 / 2
	fmt.Println(ratios)

	//under the hood: solutions
	//ratios = float64(int(3) / int(2))
	//ratios = float64(3) / 2

	/*
			Precedence    Operator
		    5             *  /  %  <<  >>  &  &^
		    4             +  -  |  ^
		    3             ==  !=  <  <=  >  >=
		    2             &&
		    1             ||
	*/
}
