package webapp

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// Enumeration of menu item kinds. Those other than NormalKind are treated
// specially.
const (
	NormalKind MenuItemKind = iota
	CutKind
	CopyKind
	PasteKind
	DeleteKind
	SelectAllKind
)

// MenuItemKind holds the different kinds of menu items.
type MenuItemKind int

func (kind MenuItemKind) title() string {
	switch kind {
	case CutKind:
		return i18n.Text("Cut")
	case CopyKind:
		return i18n.Text("Copy")
	case PasteKind:
		return i18n.Text("Paste")
	case DeleteKind:
		return i18n.Text("Delete")
	case SelectAllKind:
		return i18n.Text("Select All")
	default:
		return i18n.Text("INVALID")
	}
}

func (kind MenuItemKind) keyCode() int {
	switch kind {
	case CutKind:
		return keys.VirtualKeyX
	case CopyKind:
		return keys.VirtualKeyC
	case PasteKind:
		return keys.VirtualKeyV
	case DeleteKind:
		return keys.VirtualKeyBackspace
	case SelectAllKind:
		return keys.VirtualKeyA
	default:
		return 0
	}
}

func (kind MenuItemKind) modifiers() keys.Modifiers {
	if kind == DeleteKind {
		return 0
	}
	return keys.PlatformMenuModifier()
}

// Selector returns the selector associated with this kind.
func (kind MenuItemKind) Selector() string {
	switch kind {
	case CutKind:
		return "cut:"
	case CopyKind:
		return "copy:"
	case PasteKind:
		return "paste:"
	case DeleteKind:
		return "delete:"
	case SelectAllKind:
		return "selectAll:"
	default:
		return "handleMenuItem:"
	}
}
