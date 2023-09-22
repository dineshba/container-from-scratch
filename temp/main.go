package main

func main() {
	vec := make([]byte, 0)
	for {
		vec = append(vec, make([]byte, 10_000_000)...)
		println(len(vec) / 10_000_000)
	}
}
