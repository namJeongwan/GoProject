package main

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]int{}

	m[hash(23)] = 10
	m[hash(259)] = 50

}
