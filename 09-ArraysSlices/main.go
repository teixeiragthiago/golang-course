package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [10]string
	array1[0] = "Pos 1"

	fmt.Println(array1)

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // os ... fixam o tamanho do array de acordo com q qtd de itens que eu passar
	fmt.Println(array3)

	slice := []int{10, 2, 1, 313, 54, 7676, 1454}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	slice2 := array2[1:3]

	fmt.Println(slice2)

	array2[1] = "Pos Alterada!"
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
