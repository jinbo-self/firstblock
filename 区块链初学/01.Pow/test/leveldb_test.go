package test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevedb(t *testing.T) {
	db, err := New("")
	check(err)

	err = db.Put([]byte("k1"), []byte("v1"))
	check(err)
	err = db.Put([]byte("k4"), []byte("v4"))
	check(err)
	err = db.Put([]byte("k2"), []byte("v2"))
	check(err)
	err = db.Put([]byte("k7"), []byte("v7"))
	check(err)
	err = db.Put([]byte("k3"), []byte("v3"))
	check(err)

	v, _ := db.Get([]byte("k1"))
	fmt.Printf("%s\n", v)
	if !bytes.Equal(v, []byte("v1")) {
		t.Fatal()
	}
	err = db.Del([]byte("k1"))
	check(err)

	iter := db.Iterator()
	for iter.Next() {
		fmt.Println(string(iter.Key()), string(iter.Value()))
	}
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
