package main

import "fmt"

type stack struct {
	Data [411]byte
	i    int
}

func (s *stack) top() (k byte) {
	if s.i == 0 {
		return 0
	}
	return s.Data[s.i-1]
}

func (s *stack) push(k byte) {
	if s.i+1 > 410 {
		return
	}
	s.Data[s.i] = k
	s.i++
}
func (s *stack) pop() byte {
	s.i--
	return s.Data[s.i]
}

func (s *stack) length() int {
	return s.i
}

// func (s stack) String()  {
// 	var str string
// 	for i := 0; i < s.i; i++ {
// 		str = str + "[" +
// 			strconv.Itoa(i) + ":" + strconv.Itoa(s.Data[i]) + "]"
// 	}
// 	return str
// }
func ishighpro(a string, b string) (ishigh bool) {

	// hopwatch.Display("a", a, "b", b)
	if a == "+" {
		// hopwatch.Display("false1", a)
		return false
	}
	if a == "-" {
		return false
	}
	if a == "*" {
		if b == "^" || b == "*" || b == "/" {
			// hopwatch.Display("false2", a)
			return false
		}
		return true
	}
	if a == "/" {
		if b == "^" || b == "/" {
			// hopwatch.Display("false2", a)
			return false
		}
		return true

	}
	if a == "^" {
		if b == "^" {
			// hopwatch.Display("false3", a)
			return false
		}
		return true
	}
	return true
}
func decode(str []byte) (ans []byte, length int) {
	ans = make([]byte, len(str))
	anslen := 0
	s := new(stack)
	for i := 0; i < len(str); i++ {
		// if str[i] >= byte('a') && str[i] <= byte('z') {
		// 	// ans[anslen] = str[i]
		// 	// anslen++
		// 	fmt.Print(string(str[i]))
		// } else if str[i] == byte(')') {
		// 	// ans[anslen] = s.pop()
		// 	// anslen++
		// 	fmt.Print(string(s.pop()))
		// } else if str[i] != byte('(') {
		// 	s.push(str[i])
		// }
		switch {
		case str[i] >= byte('a') && str[i] <= byte('z'):
			ans[anslen] = str[i]
			anslen++
			break
		case str[i] == byte('('):
			s.push(str[i])
			break
		case str[i] == byte(')'):
			for j := s.pop(); j != byte('(') && s.length() > 0; j = s.pop() {
				ans[anslen] = byte(j)
				anslen++
			}
			break
		default:
			for !ishighpro(string(str[i]), string(s.top())) && s.length() != 0 && s.top() != byte('(') {
				ans[anslen] = s.pop()
				anslen++

			}
			s.push(str[i])
			break
		}
	}
	if s.length() != 0 {
		for j := s.pop(); s.length() > 0; j = s.pop() {

			ans[anslen] = j
			anslen++
		}
	}
	return ans, anslen
}
func main() {
	var number int
	fmt.Scanln(&number)
	for i := 0; i < number; i++ {
		var ex string
		fmt.Scanln(&ex)
		// fmt.Println(ex)
		// ex = "(" + ex + ")"
		ans, length := decode([]byte(ex))
		t := ans[0:length]
		fmt.Println(string(t))
	}
}
