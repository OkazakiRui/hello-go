package main

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

func main() {
	// s_var()
	// s_const()
	s_iota()
}
