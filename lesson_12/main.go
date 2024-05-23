package main

import "fmt"

// interface
// 最もポピュラーな使い方、異なる方に共通の性質を持たせるために使われる

// type Stringfy interface {
// 	ToString() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) ToString() string {
// 	return p.Name
// }

// type Car struct {
// 	Number string
// 	Model  string
// }

// func (c Car) ToString() string {
// 	return c.Number
// }

// func main() {
// 	var stringers []Stringfy
// 	stringers = append(stringers, Person{"Taro", 21})
// 	stringers = append(stringers, Car{"1234", "AB-1234"})

// 	for _, s := range stringers {
// 		println(s.ToString())
// 	}
// }

// カスタムエラー
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{"カスタムエラーメッセージ", 1234}
}

func main() {
	err := RaiseError()
	if err != nil {
		fmt.Println(err.Error())
	}
	// エラーの構造体は型アサーションで取り出すことができる
	if myErr, ok := err.(*MyError); ok {
		fmt.Println(myErr.ErrCode)
	}
}
