package main

import (
	"fmt"
)

type Hoge map[string]string

func (h Hoge) Get(key string) string {
	return h[key]
}

func main() {
	h := Hoge{"hoge": "hoge", "fuga": "FUGA"}
	fmt.Println(h["hhh"])
}
