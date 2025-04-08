package main

type menu struct {
	items    []item
	offset   int
	itemSize int
	menuSize int
}

type item struct {
	id        int
	x         int
	y         int
	coord     [2]int
	itemImage int
}
