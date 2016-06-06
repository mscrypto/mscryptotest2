package main

import (
	"fmt"
	"github.com/mscrypto/mscryptotest2/shape"
)


func main() {

	circle1 := &shape.Circle{
		R : 5,
	}

	square1 := &shape.Square{
		Width : 5,
		Height : 5,
	}

	fmt.Println(shape.ShapeArea(circle1))
	fmt.Println(shape.ShapeArea(square1))

}
