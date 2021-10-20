package main

import (
	"fmt"
	"reflect"
) 

type par struct {
	key string, 
	value interface{},
}

func (p *par) Set(key string, value interface{}) error {
	if reflect.TypeOf(key) == int || reflect.TypeOf(key) == float64 {
		return errors.New("invalid key")
	}
	p.key = key
	p.value = value
	return nil
}

func (p *par) Get(key string) interface{} {
	return p.value
}

func main() {
	cache := par{}
	err := cache.Set("userId", 42)
	if err != nil {
		log.Fatal(err)
	}
	userId := cache.Get("userId")

	fmt.Println(userId)

	/*cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)*/
}
