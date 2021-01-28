package main

//~Do not use pointer receivers for slices, maps, funcs, and chan. They already carry that with them.
//~You can use pointer receivers for ints, arrays, structs, string, array, float64
func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 15}
		tetris    = game{title: "tetris", price: 8}
	)

	minecraft.discount(.5)

	mobydick.print()
	minecraft.print()
	tetris.print()

	var h huge

	for i := 0; i < 10; i++ {
		h.addr()
	}
}
