package main

import (
	"fmt"
)

// bitmap we'll use to map out the presence of numbers
type Bitmap []byte

func NewBitmap(max int, filePath string) Bitmap {
	b := make(Bitmap, max, max)
	return b
}

// sets the integer number's present in the bitmap
func (bits Bitmap) Set(i int) {
	index := byte(i/8)
	bit := byte(i%8)
	bits[index] |= 1 << bit
	//return bits[i/8] & i%8 != 0
}

// Not sure here?
func (bits Bitmap) Get(number int) bool {
	return true
}


func main() {
	shit := NewBitmap(100, "shit")
	shit.Set(18-1)
	fmt.Println(shit)
}

