package main

func tri(n int) int {
	return n * (n + 1) / 2
}

func sqr(n int) int {
	return n * n
}

func pent(n int) int {
	return n * (3*n - 1) / 2
}

func hex(n int) int {
	return n * (2*n - 1)
}

func hept(n int) int {
	return n * (5*n - 3) / 2
}

func oct(n int) int {
	return n * (3*n - 2)
}

func shapeFunc(sides int) func(int) int {
	funcMap := make(map[int]func(int) int)

	funcMap[3] = tri
	funcMap[4] = sqr
	funcMap[5] = pent
	funcMap[6] = hex
	funcMap[7] = hept
	funcMap[8] = oct

	return funcMap[sides]
}
