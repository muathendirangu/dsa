package main

func reverseInPlace(list []int, start int, end int) {
	for i := start; i <= (start+end)/2 && i < end-1+start; i++ {
		list[i], list[end-(i-start)] = list[end-(i-start)], list[i]
	}
}
