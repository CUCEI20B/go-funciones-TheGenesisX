package main

func Burbuja(s []int64) {
	n := len(s) - 1

	for x := 0; x <= n; x++ {
		i := 0
		j := 1

		for j <= n {
			if s[i] > s[j] {
				aux := s[i]
				s[i] = s[j]
				s[j] = aux
			}
			i++
			j++
		}
	}
}

func main() {

}
