package main

import "fmt"

func s_var() {
	var n int = 100
	// 自動推論
	var m = 40
	// var を省略 関数ないのみ
	a := 30
	var (
		b = 20
		c = 30
	)
	println(n, m, a, b, c)
}

func s_const() {
	const n int = 100
	const m = 40
	const (
		// 同じ値の場合は右辺を省略することができる
		a = 20
		b
		c
	)
	println(n, m, a, b, c)
}

func s_iota()  {
	const (
		// 連番を付与する
		a = iota
		b
		c
	)
	println(a, b, c)

	const (
		// 左シフトさせる
		d = 1 << iota
		e
		f
	)
	println(d, e, f)
}

func s_if() {
	if a := -12; a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println(2*a)
	}
}

func s_switch() {
	a := 1
	switch a {
		case 1, 2:
			fmt.Println("a is 1 or 2")
			// break がいらない
		default:
			fmt.Println("defualt")
	}
}

// goの繰り返しはforしかない
func s_for()  {
	for i := 0; i < 100; i++ {
		fmt.Printf("i: %d\n", i)
	}

	j := 0
	for j < 100 {
		j++
		fmt.Printf("j: %d\n", j)
	}

	// 無限ループ
	// for{
	// }

	// range を使用したループ
	for k, l := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("k: %d, l: %d\n", k, l)
	}

	var m int
	for{
		fmt.Printf("m: %d\n", m)
		if m == 10 {
			// break out of loop
			break
		}
		m++
	}

	var n int

	LOOP:
		for{
			fmt.Printf("n: %d\n", n)
			if n == 10 {
				// break out of loop
				break LOOP
			}
			n++
		}
}

func main() {
	// s_var()
	// s_const()
	// s_iota()
	// s_if()
	// s_switch()
	s_for()
}
