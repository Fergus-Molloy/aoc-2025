package day3

import (
	_ "embed"
	"fmt"
	"iter"
	"strconv"
	"strings"
	"sync"
)

//go:embed 3.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

type Zip struct {
	idx, value int
}

func Pt1() (int, error) {
	sum := 0
	for l := range strings.SplitSeq(inp, "\n") {
		biggest := int(l[0]) - '0'
		var maxVal int

		for _, r := range l[1:] {
			v := int(r - '0')
			currentValue := (biggest * 10) + v
			if currentValue > maxVal {
				maxVal = currentValue
			}
			if v > biggest {
				biggest = v
			}
		}
		sum += maxVal
	}
	return sum, nil
}

// This would finish eventually...
func Pt2_Slow() (int, error) {
	sum := 0
	ansChan := make(chan int)
	doneChan := make(chan struct{})
	wg := new(sync.WaitGroup)

	go func() {
		wg.Wait()
		close(doneChan)
	}()
	for l := range strings.SplitSeq(inp, "\n") {
		wg.Add(1)
		go maxNRec(12, l, ansChan, wg)
	}

	for {
		select {
		case m := <-ansChan:
			sum += m
		case <-doneChan:
			return sum, nil
		}
	}
}

func maxNRec(n int, l string, ansChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	maxV := 0
	for p := range getProductsRec(n, l) {
		if p > maxV {
			maxV = p
		}
	}

	fmt.Printf("%s: max[%d]\n", l, maxV)
	ansChan <- maxV
}

func getProductsRec(n int, l string) iter.Seq[int] {
	return func(yield func(int) bool) {
		if n == 1 {
			for _, r := range l {
				if !yield(toInt(r)) {
					return
				}
			}
		}

		for i, r := range l {
			for p := range getProductsRec(n-1, l[i+1:]) {
				if !yield(toInt(r)*intPow(n-1) + p) {
					return
				}
			}
		}

	}
}

func Pt2() (int, error) {
	sum := 0

	for l := range strings.SplitSeq(inp, "\n") {
		m, err := maxLuce(12, l)
		if err != nil {
			return 0, err
		}
		sum += m
	}

	return sum, nil
}

func Pt2Parallel() (int, error) {
	sum := 0

	wg := new(sync.WaitGroup)
	for l := range strings.SplitSeq(inp, "\n") {

		wg.Go(func() {
			m, err := maxLuce(12, l)
			if err != nil {
				return
			}
			sum += m
		})
	}
	wg.Wait()

	return sum, nil
}

func maxLuce(n int, l string) (int, error) {
	stack := make([]rune, 0)
	count := len(l) - n

	for _, r := range l {
		for len(stack) > 0 && back(stack) < r && count > 0 {
			stack = pop(stack)
			count -= 1
		}
		stack = append(stack, r)
	}
	return strconv.Atoi(toString(stack[:n]))
}

func pop(l []rune) []rune {
	if len(l) > 0 {
		return l[:len(l)-1]
	} else {
		return []rune{}
	}
}

func back(l []rune) rune {
	if len(l) > 0 {
		return l[len(l)-1]
	} else {
		return '0'
	}
}

func toString(l []rune) string {
	out := ""
	for _, r := range l {
		out = out + string(r)
	}
	return out
}
func toInt[T rune | byte](r T) int {
	return int(r - '0')
}

func intPow(n int) int {
	if n == 0 {
		return 1
	}
	base := 10
	for range n - 1 {
		base *= 10
	}
	return base
}
