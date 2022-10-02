package main

import "fmt"

var books = []string{}

func main() {
	//新增
	books = append(books, "01")
	fmt.Println(books)

	//刪除
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[0:2], slice[3:5]...)
	fmt.Println(slice)
}
