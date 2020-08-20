package main

import (
	"fmt"
	"plugin"
)

func main2() {
	p, err := plugin.Open("silk.so")
	if err != nil {
		panic(err)
	}
	m, err := p.Lookup("TransSilk")
	if err != nil {
		panic(err)
	}
	res := m.(func(string) string)("example/2020_08_20_12_15_38_829.silk")
	fmt.Println(res)
}
