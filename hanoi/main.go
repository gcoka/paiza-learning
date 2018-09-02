package main

import "fmt"

type Stack struct {
	List []int
}

func NewStack() *Stack {
	s := &Stack{}
	s.List = make([]int, 0)
	return s
}

func (s *Stack) Push(v int) {
	s.List = append(s.List, v)
}

func (s *Stack) CanPop() bool {
	return len(s.List) > 0
}

func (s *Stack) Pop() int {
	if !s.CanPop() {
		panic("cannot Pop")
	}

	v := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return v
}

func hanoi(n, stop int, count *int, from, via, to *Stack) {
	if n == 1 {
		v := from.Pop()
		to.Push(v)
		*count++
	} else {
		hanoi(n-1, stop, count, from, to, via)
		if *count >= stop {
			return
		}
		v := from.Pop()
		to.Push(v)
		*count++
		if *count >= stop {
			return
		}
		hanoi(n-1, stop, count, via, from, to)
	}
	return
}

func writeResult(s *Stack) {
	if s.CanPop() {
		for i, v := range s.List {
			fmt.Print(v)
			if i < len(s.List)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	} else {
		fmt.Println("-")
	}
}

func main() {
	from := NewStack() // A
	via := NewStack()  // B
	to := NewStack()   // C

	var n, stop int
	count := 0

	fmt.Scan(&n)
	fmt.Scan(&stop)

	for i := n; i > 0; i-- {
		from.Push(i)
	}

	hanoi(n, stop, &count, from, via, to)
	writeResult(from)
	writeResult(via)
	writeResult(to)
}
