// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"fmt"
	_ "runtime/cgo"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtksourceview-5
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_background_pattern_type_get_type()), F: marshalBackgroundPatternType},
		{T: externglib.Type(C.gtk_source_bracket_match_type_get_type()), F: marshalBracketMatchType},
		{T: externglib.Type(C.gtk_source_change_case_type_get_type()), F: marshalChangeCaseType},
		{T: externglib.Type(C.gtk_source_completion_activation_get_type()), F: marshalCompletionActivation},
		{T: externglib.Type(C.gtk_source_completion_column_get_type()), F: marshalCompletionColumn},
		{T: externglib.Type(C.gtk_source_compression_type_get_type()), F: marshalCompressionType},
		{T: externglib.Type(C.gtk_source_newline_type_get_type()), F: marshalNewlineType},
		{T: externglib.Type(C.gtk_source_smart_home_end_type_get_type()), F: marshalSmartHomeEndType},
		{T: externglib.Type(C.gtk_source_view_gutter_position_get_type()), F: marshalViewGutterPosition},
		{T: externglib.Type(C.gtk_source_sort_flags_get_type()), F: marshalSortFlags},
	})
}

type BackgroundPatternType C.gint

const (
	// SourceBackgroundPatternTypeNone: no pattern.
	SourceBackgroundPatternTypeNone BackgroundPatternType = iota
	// SourceBackgroundPatternTypeGrid: grid pattern.
	SourceBackgroundPatternTypeGrid
)

func marshalBackgroundPatternType(p uintptr) (interface{}, error) {
	return BackgroundPatternType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BackgroundPatternType.
func (b BackgroundPatternType) String() string {
	switch b {
	case SourceBackgroundPatternTypeNone:
		return "None"
	case SourceBackgroundPatternTypeGrid:
		return "Grid"
	default:
		return fmt.Sprintf("BackgroundPatternType(%d)", b)
	}
}

type BracketMatchType C.gint

const (
	// SourceBracketMatchNone: there is no bracket to match.
	SourceBracketMatchNone BracketMatchType = iota
	// SourceBracketMatchOutOfRange: matching a bracket failed because the
	// maximum range was reached.
	SourceBracketMatchOutOfRange
	// SourceBracketMatchNotFound: matching bracket was not found.
	SourceBracketMatchNotFound
	// SourceBracketMatchFound: matching bracket was found.
	SourceBracketMatchFound
)

func marshalBracketMatchType(p uintptr) (interface{}, error) {
	return BracketMatchType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BracketMatchType.
func (b BracketMatchType) String() string {
	switch b {
	case SourceBracketMatchNone:
		return "None"
	case SourceBracketMatchOutOfRange:
		return "OutOfRange"
	case SourceBracketMatchNotFound:
		return "NotFound"
	case SourceBracketMatchFound:
		return "Found"
	default:
		return fmt.Sprintf("BracketMatchType(%d)", b)
	}
}

type ChangeCaseType C.gint

const (
	// SourceChangeCaseLower: change case to lowercase.
	SourceChangeCaseLower ChangeCaseType = iota
	// SourceChangeCaseUpper: change case to uppercase.
	SourceChangeCaseUpper
	// SourceChangeCaseToggle: toggle case of each character.
	SourceChangeCaseToggle
	// SourceChangeCaseTitle: capitalize each word.
	SourceChangeCaseTitle
)

func marshalChangeCaseType(p uintptr) (interface{}, error) {
	return ChangeCaseType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ChangeCaseType.
func (c ChangeCaseType) String() string {
	switch c {
	case SourceChangeCaseLower:
		return "Lower"
	case SourceChangeCaseUpper:
		return "Upper"
	case SourceChangeCaseToggle:
		return "Toggle"
	case SourceChangeCaseTitle:
		return "Title"
	default:
		return fmt.Sprintf("ChangeCaseType(%d)", c)
	}
}

type CompletionActivation C.gint

const (
	SourceCompletionActivationNone CompletionActivation = iota
	SourceCompletionActivationInteractive
	SourceCompletionActivationUserRequested
)

func marshalCompletionActivation(p uintptr) (interface{}, error) {
	return CompletionActivation(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CompletionActivation.
func (c CompletionActivation) String() string {
	switch c {
	case SourceCompletionActivationNone:
		return "None"
	case SourceCompletionActivationInteractive:
		return "Interactive"
	case SourceCompletionActivationUserRequested:
		return "UserRequested"
	default:
		return fmt.Sprintf("CompletionActivation(%d)", c)
	}
}

type CompletionColumn C.gint

const (
	SourceCompletionColumnIcon CompletionColumn = iota
	SourceCompletionColumnBefore
	SourceCompletionColumnTypedText
	SourceCompletionColumnAfter
	SourceCompletionColumnComment
	SourceCompletionColumnDetails
)

func marshalCompletionColumn(p uintptr) (interface{}, error) {
	return CompletionColumn(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CompletionColumn.
func (c CompletionColumn) String() string {
	switch c {
	case SourceCompletionColumnIcon:
		return "Icon"
	case SourceCompletionColumnBefore:
		return "Before"
	case SourceCompletionColumnTypedText:
		return "TypedText"
	case SourceCompletionColumnAfter:
		return "After"
	case SourceCompletionColumnComment:
		return "Comment"
	case SourceCompletionColumnDetails:
		return "Details"
	default:
		return fmt.Sprintf("CompletionColumn(%d)", c)
	}
}

type CompressionType C.gint

const (
	// SourceCompressionTypeNone: plain text.
	SourceCompressionTypeNone CompressionType = iota
	// SourceCompressionTypeGzip: gzip compression.
	SourceCompressionTypeGzip
)

func marshalCompressionType(p uintptr) (interface{}, error) {
	return CompressionType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CompressionType.
func (c CompressionType) String() string {
	switch c {
	case SourceCompressionTypeNone:
		return "None"
	case SourceCompressionTypeGzip:
		return "Gzip"
	default:
		return fmt.Sprintf("CompressionType(%d)", c)
	}
}

// The function returns the following values:
//
func FileLoaderErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_source_file_loader_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// The function returns the following values:
//
func FileSaverErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_source_file_saver_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type NewlineType C.gint

const (
	// SourceNewlineTypeLf: line feed, used on UNIX.
	SourceNewlineTypeLf NewlineType = iota
	// SourceNewlineTypeCr: carriage return, used on Mac.
	SourceNewlineTypeCr
	// SourceNewlineTypeCrLf: carriage return followed by a line feed, used on
	// Windows.
	SourceNewlineTypeCrLf
)

func marshalNewlineType(p uintptr) (interface{}, error) {
	return NewlineType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for NewlineType.
func (n NewlineType) String() string {
	switch n {
	case SourceNewlineTypeLf:
		return "Lf"
	case SourceNewlineTypeCr:
		return "Cr"
	case SourceNewlineTypeCrLf:
		return "CrLf"
	default:
		return fmt.Sprintf("NewlineType(%d)", n)
	}
}

type SmartHomeEndType C.gint

const (
	// SourceSmartHomeEndDisabled: smart-home-end disabled.
	SourceSmartHomeEndDisabled SmartHomeEndType = iota
	// SourceSmartHomeEndBefore: move to the first/last non-whitespace character
	// on the first press of the HOME/END keys and to the beginning/end of the
	// line on the second press.
	SourceSmartHomeEndBefore
	// SourceSmartHomeEndAfter: move to the beginning/end of the line on the
	// first press of the HOME/END keys and to the first/last non-whitespace
	// character on the second press.
	SourceSmartHomeEndAfter
	// SourceSmartHomeEndAlways always move to the first/last non-whitespace
	// character when the HOME/END keys are pressed.
	SourceSmartHomeEndAlways
)

func marshalSmartHomeEndType(p uintptr) (interface{}, error) {
	return SmartHomeEndType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SmartHomeEndType.
func (s SmartHomeEndType) String() string {
	switch s {
	case SourceSmartHomeEndDisabled:
		return "Disabled"
	case SourceSmartHomeEndBefore:
		return "Before"
	case SourceSmartHomeEndAfter:
		return "After"
	case SourceSmartHomeEndAlways:
		return "Always"
	default:
		return fmt.Sprintf("SmartHomeEndType(%d)", s)
	}
}

type ViewGutterPosition C.gint

const (
	// SourceViewGutterPositionLines: gutter position of the lines renderer.
	SourceViewGutterPositionLines ViewGutterPosition = -30
	// SourceViewGutterPositionMarks: gutter position of the marks renderer.
	SourceViewGutterPositionMarks ViewGutterPosition = -20
)

func marshalViewGutterPosition(p uintptr) (interface{}, error) {
	return ViewGutterPosition(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ViewGutterPosition.
func (v ViewGutterPosition) String() string {
	switch v {
	case SourceViewGutterPositionLines:
		return "Lines"
	case SourceViewGutterPositionMarks:
		return "Marks"
	default:
		return fmt.Sprintf("ViewGutterPosition(%d)", v)
	}
}

type SortFlags C.guint

const (
	// SourceSortFlagsNone: no flags specified.
	SourceSortFlagsNone SortFlags = 0b0
	// SourceSortFlagsCaseSensitive: case sensitive sort.
	SourceSortFlagsCaseSensitive SortFlags = 0b1
	// SourceSortFlagsReverseOrder: sort in reverse order.
	SourceSortFlagsReverseOrder SortFlags = 0b10
	// SourceSortFlagsRemoveDuplicates: remove duplicates.
	SourceSortFlagsRemoveDuplicates SortFlags = 0b100
)

func marshalSortFlags(p uintptr) (interface{}, error) {
	return SortFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SortFlags.
func (s SortFlags) String() string {
	if s == 0 {
		return "SortFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(108)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SourceSortFlagsNone:
			builder.WriteString("None|")
		case SourceSortFlagsCaseSensitive:
			builder.WriteString("CaseSensitive|")
		case SourceSortFlagsReverseOrder:
			builder.WriteString("ReverseOrder|")
		case SourceSortFlagsRemoveDuplicates:
			builder.WriteString("RemoveDuplicates|")
		default:
			builder.WriteString(fmt.Sprintf("SortFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SortFlags) Has(other SortFlags) bool {
	return (s & other) == other
}