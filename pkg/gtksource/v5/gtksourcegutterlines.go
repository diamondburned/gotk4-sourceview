// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// glib.Type values for gtksourcegutterlines.go.
var GTypeGutterLines = externglib.Type(C.gtk_source_gutter_lines_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGutterLines, F: marshalGutterLines},
	})
}

// GutterLinesOverrider contains methods that are overridable.
type GutterLinesOverrider interface {
}

type GutterLines struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*GutterLines)(nil)
)

func classInitGutterLinesser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGutterLines(obj *externglib.Object) *GutterLines {
	return &GutterLines{
		Object: obj,
	}
}

func marshalGutterLines(p uintptr) (interface{}, error) {
	return wrapGutterLines(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddClass adds the class name to line.
//
// name will be converted to a #GQuark as part of this process. A faster version
// of this function is available via gtk_source_gutter_lines_add_qclass() for
// situations where the #GQuark is known ahead of time.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - name class name.
//
func (lines *GutterLines) AddClass(line uint, name string) {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 *C.gchar                // out

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_source_gutter_lines_add_class(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(name)
}

// AddQclass adds the class denoted by qname to line.
//
// You may check if a line has qname by calling
// gtk_source_gutter_lines_has_qclass().
//
// You can remove qname by calling gtk_source_gutter_lines_remove_qclass().
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - qname class name as a #GQuark.
//
func (lines *GutterLines) AddQclass(line uint, qname glib.Quark) {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 C.GQuark                // out

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = C.guint32(qname)

	C.gtk_source_gutter_lines_add_qclass(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(qname)
}

// Buffer gets the TextBuffer that the SourceGutterLines represents.
//
// The function returns the following values:
//
//    - textBuffer: TextBuffer.
//
func (lines *GutterLines) Buffer() *gtk.TextBuffer {
	var _arg0 *C.GtkSourceGutterLines // out
	var _cret *C.GtkTextBuffer        // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))

	_cret = C.gtk_source_gutter_lines_get_buffer(_arg0)
	runtime.KeepAlive(lines)

	var _textBuffer *gtk.TextBuffer // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_textBuffer = &gtk.TextBuffer{
			Object: obj,
		}
	}

	return _textBuffer
}

// First gets the line number (starting from 0) for the first line that is user
// visible.
//
// The function returns the following values:
//
//    - guint: line number starting from 0.
//
func (lines *GutterLines) First() uint {
	var _arg0 *C.GtkSourceGutterLines // out
	var _cret C.guint                 // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))

	_cret = C.gtk_source_gutter_lines_get_first(_arg0)
	runtime.KeepAlive(lines)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IterAtLine gets a TextIter for the current buffer at line.
//
// The function takes the following parameters:
//
//    - line number.
//
// The function returns the following values:
//
//    - iter: location for a TextIter.
//
func (lines *GutterLines) IterAtLine(line uint) *gtk.TextIter {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.GtkTextIter           // in
	var _arg2 C.guint                 // out

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg2 = C.guint(line)

	C.gtk_source_gutter_lines_get_iter_at_line(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)

	var _iter *gtk.TextIter // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Last gets the line number (starting from 0) for the last line that is user
// visible.
//
// The function returns the following values:
//
//    - guint: line number starting from 0.
//
func (lines *GutterLines) Last() uint {
	var _arg0 *C.GtkSourceGutterLines // out
	var _cret C.guint                 // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))

	_cret = C.gtk_source_gutter_lines_get_last(_arg0)
	runtime.KeepAlive(lines)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// LineYrange gets the Y range for a line based on mode.
//
// The value for y is relative to the renderers widget coordinates.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - mode: SourceGutterRendererAlignmentMode.
//
// The function returns the following values:
//
//    - y: location for the Y position in widget coordinates.
//    - height: line height based on mode.
//
func (lines *GutterLines) LineYrange(line uint, mode GutterRendererAlignmentMode) (y int, height int) {
	var _arg0 *C.GtkSourceGutterLines                // out
	var _arg1 C.guint                                // out
	var _arg2 C.GtkSourceGutterRendererAlignmentMode // out
	var _arg3 C.gint                                 // in
	var _arg4 C.gint                                 // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = C.GtkSourceGutterRendererAlignmentMode(mode)

	C.gtk_source_gutter_lines_get_line_yrange(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(mode)

	var _y int      // out
	var _height int // out

	_y = int(_arg3)
	_height = int(_arg4)

	return _y, _height
}

// View gets the TextView that the SourceGutterLines represents.
//
// The function returns the following values:
//
//    - textView: TextView.
//
func (lines *GutterLines) View() *gtk.TextView {
	var _arg0 *C.GtkSourceGutterLines // out
	var _cret *C.GtkTextView          // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))

	_cret = C.gtk_source_gutter_lines_get_view(_arg0)
	runtime.KeepAlive(lines)

	var _textView *gtk.TextView // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_textView = &gtk.TextView{
			Widget: gtk.Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: gtk.Accessible{
					Object: obj,
				},
				Buildable: gtk.Buildable{
					Object: obj,
				},
				ConstraintTarget: gtk.ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Scrollable: gtk.Scrollable{
				Object: obj,
			},
		}
	}

	return _textView
}

// HasClass checks to see if gtk_source_gutter_lines_add_class() was called with
// the name for line.
//
// A faster version of this function is provided via
// gtk_source_gutter_lines_has_qclass() for situations where the quark is known
// ahead of time.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - name class name that may be converted, to a #GQuark.
//
// The function returns the following values:
//
//    - ok: TRUE if line contains name.
//
func (lines *GutterLines) HasClass(line uint, name string) bool {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 *C.gchar                // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_source_gutter_lines_has_class(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasQclass checks to see if gtk_source_gutter_lines_add_qclass() was called
// with the quark denoted by qname for line.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - qname containing the class name.
//
// The function returns the following values:
//
//    - ok: TRUE if line contains qname.
//
func (lines *GutterLines) HasQclass(line uint, qname glib.Quark) bool {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 C.GQuark                // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = C.guint32(qname)

	_cret = C.gtk_source_gutter_lines_has_qclass(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(qname)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsCursor checks to see if line contains the insertion cursor.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//
// The function returns the following values:
//
//    - ok: TRUE if the insertion cursor is on line.
//
func (lines *GutterLines) IsCursor(line uint) bool {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)

	_cret = C.gtk_source_gutter_lines_is_cursor(_arg0, _arg1)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsPrelit checks to see if line is marked as prelit. Generally, this means the
// mouse pointer is over the line within the gutter.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//
// The function returns the following values:
//
//    - ok: TRUE if the line is prelit.
//
func (lines *GutterLines) IsPrelit(line uint) bool {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)

	_cret = C.gtk_source_gutter_lines_is_prelit(_arg0, _arg1)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected checks to see if the view had a selection and if that selection
// overlaps line in some way.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//
// The function returns the following values:
//
//    - ok: TRUE if the line contains a selection.
//
func (lines *GutterLines) IsSelected(line uint) bool {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)

	_cret = C.gtk_source_gutter_lines_is_selected(_arg0, _arg1)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveClass removes the class matching name from line.
//
// A faster version of this function is available via
// gtk_source_gutter_lines_remove_qclass() for situations where the #GQuark is
// known ahead of time.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - name class name.
//
func (lines *GutterLines) RemoveClass(line uint, name string) {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 *C.gchar                // out

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_source_gutter_lines_remove_class(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(name)
}

// RemoveQclass reverses a call to gtk_source_gutter_lines_add_qclass() by
// removing the #GQuark matching qname.
//
// The function takes the following parameters:
//
//    - line number starting from zero.
//    - qname to remove from line.
//
func (lines *GutterLines) RemoveQclass(line uint, qname glib.Quark) {
	var _arg0 *C.GtkSourceGutterLines // out
	var _arg1 C.guint                 // out
	var _arg2 C.GQuark                // out

	_arg0 = (*C.GtkSourceGutterLines)(unsafe.Pointer(externglib.InternObject(lines).Native()))
	_arg1 = C.guint(line)
	_arg2 = C.guint32(qname)

	C.gtk_source_gutter_lines_remove_qclass(_arg0, _arg1, _arg2)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(line)
	runtime.KeepAlive(qname)
}
