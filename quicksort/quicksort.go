package quickSort

import (
	"fmt"
)

func partition(A []int, p int, r int) int {
	x := A[r-1]
	i := p - 1
	for j:=p;j<r-1;j++ {
		if A[j] <= x {
			i ++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r-1] = A[r-1], A[i+1]
	return i + 1
}

func quickSort(A []int, p int, r int){
	if p < r {
		q := partition(A, p, r)
		fmt.Println(q)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}
