package webapp

import (
	"github.com/richardwilkes/webapp/keys"
)

// MenuItem represents individual actions that can be issued from a menu.
type MenuItem struct {
	platformMenuItem
	Validator    func() bool
	Handler      func()
	title        string
	keyCode      int
	keyModifiers keys.Modifiers
	enabled      bool
}

// NewMenuSeparator creates a new menu separator item.
func NewMenuSeparator() *MenuItem {
	item := &MenuItem{}
	item.platformInitMenuSeparator()
	return item
}

// NewMenuItem creates a new menu item with no key accelerator.
func NewMenuItem(title string) *MenuItem {
	return NewMenuItemWithKey(title, 0)
}

// NewMenuItemWithKey creates a new menu item with a key accelerator using the
// platform default modifiers.
func NewMenuItemWithKey(title string, keyCode int) *MenuItem {
	return NewMenuItemWithKeyAndModifiers(title, keyCode, keys.PlatformMenuModifier())
}

// NewMenuItemWithKeyAndModifiers creates a new menu item.
func NewMenuItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers) *MenuItem {
	item := &MenuItem{
		title:        title,
		keyCode:      keyCode,
		keyModifiers: modifiers,
		enabled:      true,
	}
	item.platformInitMenuItem(NormalKind)
	return item
}

// NewSpecialMenuItem creates a new menu item that has special handling.
func NewSpecialMenuItem(kind MenuItemKind) *MenuItem {
	item := &MenuItem{
		title:        kind.title(),
		keyCode:      kind.keyCode(),
		keyModifiers: kind.modifiers(),
	}
	item.platformInitMenuItem(kind)
	return item
}

// Title returns this item's title.
func (item *MenuItem) Title() string {
	return item.title
}

// KeyCode returns the key code that can be used to trigger this item. A value
// of 0 indicates no key is attached.
func (item *MenuItem) KeyCode() int {
	return item.keyCode
}

// KeyModifiers returns the key modifiers that are required to trigger this
// item.
func (item *MenuItem) KeyModifiers() keys.Modifiers {
	return item.keyModifiers
}

// SubMenu returns a sub-menu attached to this item or nil.
func (item *MenuItem) SubMenu() *Menu {
	return item.platformSubMenu()
}

// Enabled returns true if this item is enabled.
func (item *MenuItem) Enabled() bool {
	return item.enabled
}

// Dispose of the menu item.
func (item *MenuItem) Dispose() {
	if sub := item.SubMenu(); sub != nil {
		sub.Dispose()
	}
	item.platformDispose()
}
