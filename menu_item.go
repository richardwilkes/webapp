package webapp

// MenuItem holds information about menu items.
type MenuItem struct {
	Tag     int
	Title   string
	SubMenu *Menu
}
