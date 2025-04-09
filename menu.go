package main

import "fmt"

type item struct {
	id        int
	x         int
	y         int
	coord     [2]int
	onSelect  func()
	itemImage int
}

func (i *item) updateItem(x, y, id int) {
	i.x = x
	i.y = y
	i.id = id
	i.coord = [2]int{x, y}
	i.onSelect = func() {
		fmt.Println("Restarting...")
	}
}

func newItem() item {
	return item{
		id:        -1,
		x:         -1,
		y:         -1,
		itemImage: 0,
	}
}

type menu struct {
	items      []item
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
		offsetX:    64,
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
	for x := m.offsetX; x <= m.menuWidth; x += m.itemWidth {
		id += 1
		i.updateItem((x-m.offsetX)/m.itemWidth, (m.offsetY)/m.itemHeight, id)
		m.items = append(m.items, i)
	}
}

func (m *menu) checkMenu(in input) {
	for i := range m.items {
		fmt.Println(m.items[i].x, m.items[i].y, in.menuClick)
		if m.items[i].x == in.menuClick[0] && m.items[i].y == in.menuClick[1] {
			if in.mouseButtonLeft {
				fmt.Println("menu item clicked")
				m.items[i].onSelect()
			}
		}
	}
}
