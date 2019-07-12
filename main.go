package main

import(
	"fmt"
	"io/ioutil"
)

func main() {
	dat,err := ioutil.ReadFile("sudoku")
	if err != nil {
		panic("File not found")
	}

	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			cell := string(dat[j+10*i])
			if cell != "\n" {
				fmt.Print(cell)
			}
			// fmt.Print(i,j)
		}
		fmt.Print("\n")
	}
}
