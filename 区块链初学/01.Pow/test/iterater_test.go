package test

import (
	"fmt"
	"testing"
)

//测试模块，测试迭代器
func TestNewDefaultIterat(t *testing.T) {
	data := make(map[string][]byte)

	data["k1"] = []byte("v1")
	data["k2"] = []byte("v2")
	data["k3"] = []byte("v3")

	iter := NewDefalutIteratro(data)

	if iter.length != 3 {
		t.Fatal()
	}
	for iter.Next() {
		fmt.Println(string(iter.Key()), string(iter.Value()))
	}
}
