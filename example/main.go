package main

import (
	"fmt"

	"github.com/hadihammurabi/spidol"
)

func main() {
	t1 := spidol.Red("red text yellow background").BgYellow().Reset()
	t2 := spidol.Yellow("yellow text red background").BgRed().Reset()
	t3 := spidol.Green("green text white background").BgWhite().Reset()
	t4 := spidol.White("normal text but bold").BgBlack().Bold().Reset()
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

	fmt.Println()
	fmt.Println(t1, t2, t3, t4)
}
