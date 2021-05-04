package main

import "fmt"

func main()  {
	h := -12050
	//fmt.Printf("hours: %v\nmin cacl:%v\nreminder:%v\n", h/12, h - 12*(h/12), h%12)

	fmt.Println(h%(12*60), 12*60)
}
