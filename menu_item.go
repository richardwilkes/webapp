package webapp

import (
	"unsafe"

	"github.com/richardwilkes/webapp/keys"
)

// MenuItem represents individual actions that can be issued from a menu.
type MenuItem struct {
	PlatformPtr  unsafe.Pointer
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
	driver.MenuItemInitSeparator(item)
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
	driver.MenuItemInit(item, NormalKind)
	return item
}

// NewSpecialMenuItem creates a new menu item that has special handling.
func NewSpecialMenuItem(kind MenuItemKind) *MenuItem {
	item := &MenuItem{
		title:        kind.title(),
		keyCode:      kind.keyCode(),
		keyModifiers: kind.modifiers(),
	}
	driver.MenuItemInit(item, kind)
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
	return driver.MenuItemSubMenu(item)
}

// Enabled returns true if this item is enabled.
func (item *MenuItem) Enabled() bool {
	return item.enabled
}

// Validate validates the item, returning true if enabled.
func (item *MenuItem) Validate() bool {
	if item.Validator != nil {
		item.enabled = item.Validator()
	} else {
		item.enabled = true
	}
	return item.enabled
}

// Dispose of the menu item.
func (item *MenuItem) Dispose() {
	if sub := item.SubMenu(); sub != nil {
		sub.Dispose()
	}
	driver.MenuItemDispose(item)
}
