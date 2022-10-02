package structTest

import "fmt"

func Run5() {
	name := "bill"
	namePointer := &name

	fmt.Println("1", namePointer)  // 0xc00008e1e0
	fmt.Println("2", &namePointer) // 0xc0000ae018
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println("3", namePointer)  // 0xc00008e1e0
	fmt.Println("4", &namePointer) // 0xc0000ae028
}

/*
實際上指標（Pointer）本身和 Slice 一樣，都是屬於 Reference Types 的變數。從下面的例子中可以看到：

使用 &name 取出 name 的 pointer 後，不論是 main 或 printPointer 裡面的 namePointer 都指稱到同一個記憶體位置
但由於 Go 本質上仍然是 Pass by Value，因此在 main 中的 &namePointer 會和 printPointer 的 &namePointer 指向到不同的記憶體位址
也就是說，當我把 Pointer 丟到函式的參數中時，實際上這個 Pointer 也被複製了一份新的（原本的 Pointer 的位址是 0xc0000ae018，複製到函式後是 0xc0000ae028，但這兩個記憶體位址，實際上都對應回 0xc00008e1e0。

ref: https://pjchender.dev/golang/pointers/#%E6%8C%87%E6%A8%99%E6%98%AF-reference-types-%E7%9A%84%E8%AE%8A%E6%95%B8
*/
