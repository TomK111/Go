/*BYTES
1 byte = 8 bites = 2^8 = 256 values

byte is an uint8 and it can represent 0-255 values
*/

package main

import (
	"fmt"
)

func main() {
	var b byte
	b = 0
	b = 255
	fmt.Printf("%08b = %d\n", b, b)

}
