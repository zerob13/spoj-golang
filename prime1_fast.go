package main

import "fmt"

func main() {
	number := 0
	primers := make([]int, 4000)
	numprimer := 0
	hash := make([]bool, 32000)
	hash[0] = true
	hash[1] = true
	for i := 2; i < 32000; i++ {
		if !hash[i] {
			primers[numprimer] = i
			numprimer += 1
			for j := i * i; j < 32000; j += i {
				hash[j] = true
			}
		}
	}

	// for i := 3; i <= 32000; i++ {
	// 	isprime := true
	// 	ca := math.Sqrt(float64(i)) + 1.0
	// 	for j := 0; j < numprimer; j++ {
	// 		if float64(j) >= ca {
	// 			break
	// 		}
	// 		if i%primers[j] == 0 {
	// 			isprime = false
	// 			break
	// 		}
	// 	}
	// 	if isprime {
	// 		primers[numprimer] = i
	// 		numprimer += 1
	// 	}
	// }
	fmt.Scanln(&number)
	for i := 0; i < number; i++ {
		var n int
		var m int
		fmt.Scanln(&m, &n)
		if i != 0 {
			fmt.Println()
		}
		if m < 2 {
			m = 2
		}
		primeset := make([]bool, 100001)
		for j := 0; j < 100001; j++ {
			primeset[j] = true
		}
		for k := 0; k < numprimer; k++ {
			p := primers[k]
			start := 0
			if p >= m {
				start = p * 2
			} else {
				start = m + ((p - m%p) % p)
			}
			for w := start; w <= n; w += p {
				primeset[w-m] = false
			}

		}
		for k := m; k <= n; k++ {
			if primeset[k-m] {
				fmt.Println(k)
			}
		}
	}
	// for number != 42 {
	// 	fmt.Println(number)
	// 	fmt.Scanln(&number)
	// }
}
