package main

import "fmt"

// generics for 1.18~

// func PrintSlice[T any](s []T) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// 特定の条件のみのジェネリクス(interfaceを実装しているもののみを受け取る)
// type Stringer interface {
// 	String() string
// }
// // func f[T fmt.Stringer](xs []T) []string {
// func f[T Stringer](xs []T) []string {
// 	result := []string{}
// 	for _, x := range xs {
// 		result = append(result, x.String())
// 	}
// 	return result
// }

// type MyInt int

// func (i MyInt) String() string {
// 	return strconv.Itoa(int(i))
// }

// typesets
// type Number interface {
// 	// ~を付与することにより、MyIntなどの方もNumberに含まれるようになる
// 	// int | int32 | int64 | float32 | float64
// 	~int | ~int32 | ~int64 | ~float32 | ~float64
// }
// type MyInt int

// func Max[T Number](x, y T) T {
// 	if x >= y {
// 		return x
// 	}
// 	return y
// }

// vector
// type Vector[T any] []T
// type IntVector = Vector[int]

// struct
// 下記は連動しているので、Aを変えるとBとCも変えないとエラーになる(一種の制約のようなもの)
// type T[A any, B []C, C *A] struct {
// 	a A
// 	b B
// 	c C
// }

// generics set
// comparableは、==, !=, <, <=, >, >=を実装しているもの(要素の比較が可能なもの)
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, x := range xs {
		s[x] = struct{}{}
	}
	return s
}

func main() {
	// PrintSlice[int]([]int{1, 2, 3})
	// PrintSlice([]int{1, 2, 3}) // どちらでもOK
	// PrintSlice([]string{"a", "b", "c"})

	// fmt.Println(f([]MyInt{1, 2, 3, 4}))

	// fmt.Println(Max[int](1, 2)) [int]省略可能
	// fmt.Println(Max[float64](1.1, 2.2)) [float64]省略可能
	// fmt.Println(Max[MyInt](1, 2))

	// var v Vector[int] = Vector[int]{1, 2, 3}
	// fmt.Println(v)
	// var v2 Vector[string] = Vector[string]{"a", "b", "c"}
	// fmt.Println(v2)
	// v3 := IntVector{1, 2, 3}
	// fmt.Println(v3)

	// key valueのデータ構造が作成可能になる
	s := NewSet(1, 2, 3)
	fmt.Println(s)

	// var t T[int, []*int, *int]
	// fmt.Printf("A: %T, B: %T, C: %T\n", t.a, t.b, t.c)

	// any型(中身を見るとinterface{}と同じ)
	// var a any = 1
	// fmt.Println("a: ", a)
}
