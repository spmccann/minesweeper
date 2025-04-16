package main

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
		i.updateItem((x-m.offsetX)/m.itemWidth+1, (m.offsetY)/m.itemHeight, id, img)
		m.items = append(m.items, i)
	}
}

func (m *menu) populateLargeMenu() {
	i := newItem()
	id := 0
	img := 0
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
