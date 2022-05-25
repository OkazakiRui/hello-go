package main


func s_var(){
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

func s_const(){
	const n int = 100
	const m = 40
	const (
		b = 20
		c = 30
	)
	println(n, m, b, c)
}

func main() {
	// s_var()
	s_const()
}
