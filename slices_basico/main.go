package main

import (
	"fmt"
)

func main() {
	var nums []int
	//cap - capacidade do slice/array
	fmt.Println(nums, len(nums), cap(nums))
	//inicializando um slice
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	//Add item no slice
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "São Paulo"
	cidades[1] = "Londrina"
	cidades[2] = "Porto Alegre"
	cidades[3] = "Rio de Janeiro"
	fmt.Println(cidades, len(cidades), cap(cidades))

	//updade slices
	capitais[1] = "Brasilia"
	fmt.Println(capitais, len(capitais), cap(capitais))

}
