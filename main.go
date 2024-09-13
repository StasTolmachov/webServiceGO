package main

import (
	"fmt"
	"os"
)

type TV struct {
	Text string
	Var  interface{}
}

// printNamedValueAndType
func pvtt(nv TV) {
	fmt.Println("__________________________________________________")
	fmt.Printf("text: %s\nvalue: %v\ntype: %T\n", nv.Text, nv.Var, nv.Var)
	fmt.Println("__________________________________________________")
}

// printValueAndType example: value: 5 type: int
func pvt(v interface{}) {
	fmt.Println("__________________________________________________")
	fmt.Printf("value: %v\ntype: %T\n", v, v)
	fmt.Println("__________________________________________________")

}

// cat data.txt | go run main.go
//
//	func unig(input io.Reader, output io.Writer) error {
//		in := bufio.NewScanner(input)
//		var prev string
//		for in.Scan() {
//			txt := in.Text()
//			if txt == prev {
//				continue
//			}
//			if txt < prev {
//				return fmt.Errorf("file not sorted")
//			}
//			prev = txt
//			fmt.Fprintln(output, txt)
//		}
//		return nil
//	}
//
//	func main() {
//		err := unig(os.Stdin, os.Stdout)
//		if err != nil {
//			panic(err.Error())
//		}
//	}
func main() {

	var path string = "testdata"
	sl, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error", err)
	}
	pvt(sl)

	for i, v := range sl {
		if v == "d" {
			os.ReadDir(path + v)
		}
	}

}
