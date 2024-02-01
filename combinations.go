package main

func combinationsRange(n, k int, ch chan []int) {
	defer close(ch)
	if n < k {
		return
	}
	if k == 0 {
		ch <- []int{}
		return
	}
	if k == 1 {
		for i := 0; i < n; i++ {
			ch <- []int{i}
		}
		return
	}
	if k == n {
		xs := make([]int, n)
		for i := 0; i < n; i++ {
			xs[i] = i
		}
		ch <- xs
		return
	}

	stack := make([]int, 0, k)
	stack = append(stack, 0)
	for len(stack) >= 1 && stack[0] <= n-k {
		for len(stack) < k {
			stack = append(stack, stack[len(stack)-1]+1)
			continue
		}
		toSend := make([]int, k)
		copy(toSend, stack)
		ch <- toSend
		for (len(stack) > 1) && (stack[len(stack)-1] >= n-len(stack)+1) {
			stack = stack[:len(stack)-1]
		}
		stack[len(stack)-1]++
	}
}

func combinations[T any](xs []T, k int, ch chan []T) {
	defer close(ch)
	indexesChannel := make(chan []int)
	go combinationsRange(len(xs), k, indexesChannel)
	for indexes := range indexesChannel {
		res := make([]T, k)
		for i, idx := range indexes {
			res[i] = xs[idx]
		}
		ch <- res
	}
}
