package main

import "fmt"

type humanPlayer struct {
	name string
}

func (hp *humanPlayer) assignName() {
	var name string
	fmt.Scanln(&name)
	(*hp).name = name
}
