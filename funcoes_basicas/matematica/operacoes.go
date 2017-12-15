package matematica

func Calculo(funcao func(int, int) int, x int, y int) int {
	return funcao(x, y)
}

func Multiplicador(x int, y int) int {
	return x * y
}

func Soma(x int, y int) int {
	return x + y
}

func Divisao(x int, y int) (resultado int) {
	resultado = x % y
	return
}

func Subtracao(x int, y int) int {
	return x - y
}

func DivisaoComResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
