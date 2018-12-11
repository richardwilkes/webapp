package webapp

// MenuItem holds information about menu items.
type MenuItem struct {
	ID      int
	Title   string
	SubMenu *Menu
}
