package main

import (
	"fmt"
	"strconv"
)

type item struct {
	id        int
	label     string
	x         int
	y         int
	coord     [2]int
	onSelect  bool
	itemImage int
}

func (i *item) updateItem(x, y, id, img int) {
	i.x = x
	i.y = y
	i.id = id
	i.coord = [2]int{x, y}
	i.itemImage = img
}

func newItem() item {
	return item{
		id:        -1,
		label:     "",
		x:         -1,
		y:         -1,
		itemImage: 0,
		onSelect:  false,
	}
}

type menu struct {
	items      []item
	largeItems []item
	offsetX    int
	offsetY    int
	itemWidth  int
	itemHeight int
	menuWidth  int
	menuHeight int
}

func newMenu() menu {
	return menu{
		items:      []item{},
		largeItems: []item{},
		offsetX:    0,
		offsetY:    0,
		itemWidth:  32,
		itemHeight: 32,
		menuWidth:  320,
		menuHeight: 32,
	}
}

func (m *menu) populateMenu() {
	i := newItem()
	id := -1
	order := []int{12, 5, 4, 4, 11, 4, 4, 4}
	for x := range order {
		id += 1
		img := order[x]
		i.updateItem((x-m.offsetX)/m.itemWidth+2, (m.offsetY)/m.itemHeight, id, img)
		m.items = append(m.items, i)
	}
}

func (m *menu) populateLargeMenu(n int) {
	i := newItem()
	id := 0
	img := n
	i.updateItem(0, (m.offsetY)/m.itemHeight, id, img)
	m.largeItems = append(m.largeItems, i)
}

func (m *menu) checkMenu(in input) {
	for i := range m.items {
		if m.items[i].x == in.menuClick[0] && m.items[i].y == in.menuClick[1] {
			if in.mouseButtonLeft {
				m.items[i].onSelect = true
			}
		}
	}
}

func (m *menu) flagCounter(flagCount int) {
	tiles := numberToTiles(flagCount)
	m.items[2].itemImage = tiles[0]
	m.items[3].itemImage = tiles[1]
}

func numberToTiles(count int) [2]int {
	key := map[int]int{0: 4, 1: 3, 2: 6, 3: 0, 4: 1, 5: 9, 6: 7, 7: 2, 8: 8, 9: 10}
	var tiles [2]int
	if count < 10 {
		tiles[0] = 4
		tiles[1] = key[count]
	} else {
		tiles[0] = 3
		tiles[1] = 4
	}
	return tiles
}

func (m *menu) timerDisplay(gameTime int) {
	tiles := timeToTile(gameTime)
	m.items[5].itemImage = tiles[0]
	m.items[6].itemImage = tiles[1]
	m.items[7].itemImage = tiles[2]
}

func timeToTile(time int) [3]int {
	key := map[int]int{0: 4, 1: 3, 2: 6, 3: 0, 4: 1, 5: 9, 6: 7, 7: 2, 8: 8, 9: 10}
	var tiles [3]int
	digits := intToDigitSlice(time)
	if len(digits) > 3 {
		tiles = [3]int{4, 4, 4}
	} else if len(digits) == 3 {
		tiles = [3]int{key[digits[0]], key[digits[1]], key[digits[2]]}
	} else if len(digits) == 2 {
		tiles = [3]int{4, key[digits[0]], key[digits[1]]}
	} else {
		tiles = [3]int{4, 4, key[time]}
	}
	return tiles
}

func intToDigitSlice(n int) []int {
	s := strconv.Itoa(n)
	digits := make([]int, len(s))
	for i, r := range s {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println("Error converting  rune to int:", err)
			return nil
		}
		digits[i] = digit
	}
	return digits
}
