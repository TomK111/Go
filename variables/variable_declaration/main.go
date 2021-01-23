package main

import "fmt"

func main() {
	var _speed int
	var SpeeD int
	var speed int
	fmt.Println(speed)
	fmt.Println(_speed)
	fmt.Println(SpeeD)

	var heat float64
	var off bool
	var brand string

	/*MULTIPLE VARIABLE DECLARATIONS
	var (
		speed int
		heat float64
		off bool
		brand string
	)
	*/

	/* PARALLEL VARIABLE DECLARATION
	If the variable has the same data type*
	func main(){
		var speed, velocity int
		fmt.println(speed, velocity)
	}
	*/

	fmt.Println(heat)         //value is 0
	fmt.Println(off)          //value is 0
	fmt.Println(brand)        //empty value
	fmt.Printf("%q\n", brand) //empty strinfg value

	/*ZERO VALUES CHEATSHEET
	booleans = false
	numerics = 0
	strings = ""
	pointers = nil
	*/

	/*BLANK IDENTIFER
	_ = blank identifier
	This operator is used to ignore values returned by function.
	Go Compiler throws an error whenever it encounters a variable declare but not used.
	*/

}
