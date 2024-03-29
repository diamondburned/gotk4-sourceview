// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// GType values.
var (
	GTypeSpaceLocationFlags = coreglib.Type(C.gtk_source_space_location_flags_get_type())
	GTypeSpaceTypeFlags     = coreglib.Type(C.gtk_source_space_type_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpaceLocationFlags, F: marshalSpaceLocationFlags},
		coreglib.TypeMarshaler{T: GTypeSpaceTypeFlags, F: marshalSpaceTypeFlags},
	})
}

// SpaceLocationFlags contains flags for white space locations.
//
// If a line contains only white spaces (no text), the white spaces match both
// GTK_SOURCE_SPACE_LOCATION_LEADING and GTK_SOURCE_SPACE_LOCATION_TRAILING.
type SpaceLocationFlags C.guint

const (
	// SourceSpaceLocationNone: no flags.
	SourceSpaceLocationNone SpaceLocationFlags = 0b0
	// SourceSpaceLocationLeading: leading white spaces on a line, i.e.
	// the indentation.
	SourceSpaceLocationLeading SpaceLocationFlags = 0b1
	// SourceSpaceLocationInsideText: white spaces inside a line of text.
	SourceSpaceLocationInsideText SpaceLocationFlags = 0b10
	// SourceSpaceLocationTrailing: trailing white spaces on a line.
	SourceSpaceLocationTrailing SpaceLocationFlags = 0b100
	// SourceSpaceLocationAll: white spaces anywhere.
	SourceSpaceLocationAll SpaceLocationFlags = 0b111
)

func marshalSpaceLocationFlags(p uintptr) (interface{}, error) {
	return SpaceLocationFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SpaceLocationFlags.
func (s SpaceLocationFlags) String() string {
	if s == 0 {
		return "SpaceLocationFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(131)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSpaceLocationNone:
			builder.WriteString("None|")
		case SourceSpaceLocationLeading:
			builder.WriteString("Leading|")
		case SourceSpaceLocationInsideText:
			builder.WriteString("InsideText|")
		case SourceSpaceLocationTrailing:
			builder.WriteString("Trailing|")
		case SourceSpaceLocationAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SpaceLocationFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpaceLocationFlags) Has(other SpaceLocationFlags) bool {
	return (s & other) == other
}

// SpaceTypeFlags contains flags for white space types.
type SpaceTypeFlags C.guint

const (
	// SourceSpaceTypeNone: no flags.
	SourceSpaceTypeNone SpaceTypeFlags = 0b0
	// SourceSpaceTypeSpace: space character.
	SourceSpaceTypeSpace SpaceTypeFlags = 0b1
	// SourceSpaceTypeTab: tab character.
	SourceSpaceTypeTab SpaceTypeFlags = 0b10
	// SourceSpaceTypeNewline: line break character. If the
	// SourceBuffer:implicit-trailing-newline property is TRUE,
	// SourceSpaceDrawer also draws a line break at the end of the buffer.
	SourceSpaceTypeNewline SpaceTypeFlags = 0b100
	// SourceSpaceTypeNbsp: non-breaking space character.
	SourceSpaceTypeNbsp SpaceTypeFlags = 0b1000
	// SourceSpaceTypeAll: all white spaces.
	SourceSpaceTypeAll SpaceTypeFlags = 0b1111
)

func marshalSpaceTypeFlags(p uintptr) (interface{}, error) {
	return SpaceTypeFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SpaceTypeFlags.
func (s SpaceTypeFlags) String() string {
	if s == 0 {
		return "SpaceTypeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(121)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSpaceTypeNone:
			builder.WriteString("None|")
		case SourceSpaceTypeSpace:
			builder.WriteString("Space|")
		case SourceSpaceTypeTab:
			builder.WriteString("Tab|")
		case SourceSpaceTypeNewline:
			builder.WriteString("Newline|")
		case SourceSpaceTypeNbsp:
			builder.WriteString("Nbsp|")
		case SourceSpaceTypeAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SpaceTypeFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SpaceTypeFlags) Has(other SpaceTypeFlags) bool {
	return (s & other) == other
}
