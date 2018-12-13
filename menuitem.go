package webapp

// MenuItem holds information about menu items.
type MenuItem struct {
	Owner   *Menu
	Index   int
	ID      int
	Title   string
	SubMenu *Menu
}
