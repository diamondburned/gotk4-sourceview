// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

//export _gotk4_gtksource5_BufferClass_bracket_matched
func _gotk4_gtksource5_BufferClass_bracket_matched(arg0 *C.GtkSourceBuffer, arg1 *C.GtkTextIter, arg2 C.GtkSourceBracketMatchType) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferOverrides](instance0)
	if overrides.BracketMatched == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferOverrides.BracketMatched, got none")
	}

	var _iter *gtk.TextIter     // out
	var _state BracketMatchType // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_state = BracketMatchType(arg2)

	overrides.BracketMatched(_iter, _state)
}

//export _gotk4_gtksource5_GutterRendererClass_activate
func _gotk4_gtksource5_GutterRendererClass_activate(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkTextIter, arg2 *C.GdkRectangle, arg3 C.guint, arg4 C.GdkModifierType, arg5 C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.Activate == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.Activate, got none")
	}

	var _iter *gtk.TextIter     // out
	var _area *gdk.Rectangle    // out
	var _button uint            // out
	var _state gdk.ModifierType // out
	var _nPresses int           // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_button = uint(arg3)
	_state = gdk.ModifierType(arg4)
	_nPresses = int(arg5)

	overrides.Activate(_iter, _area, _button, _state, _nPresses)
}

//export _gotk4_gtksource5_GutterRendererClass_begin
func _gotk4_gtksource5_GutterRendererClass_begin(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkSourceGutterLines) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.Begin == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.Begin, got none")
	}

	var _lines *GutterLines // out

	_lines = wrapGutterLines(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.Begin(_lines)
}

//export _gotk4_gtksource5_GutterRendererClass_change_buffer
func _gotk4_gtksource5_GutterRendererClass_change_buffer(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkSourceBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.ChangeBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.ChangeBuffer, got none")
	}

	var _oldBuffer *Buffer // out

	if arg1 != nil {
		_oldBuffer = wrapBuffer(coreglib.Take(unsafe.Pointer(arg1)))
	}

	overrides.ChangeBuffer(_oldBuffer)
}

//export _gotk4_gtksource5_GutterRendererClass_change_view
func _gotk4_gtksource5_GutterRendererClass_change_view(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkSourceView) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.ChangeView == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.ChangeView, got none")
	}

	var _oldView *View // out

	if arg1 != nil {
		_oldView = wrapView(coreglib.Take(unsafe.Pointer(arg1)))
	}

	overrides.ChangeView(_oldView)
}

//export _gotk4_gtksource5_GutterRendererClass_end
func _gotk4_gtksource5_GutterRendererClass_end(arg0 *C.GtkSourceGutterRenderer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.End == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.End, got none")
	}

	overrides.End()
}

//export _gotk4_gtksource5_GutterRendererClass_query_activatable
func _gotk4_gtksource5_GutterRendererClass_query_activatable(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkTextIter, arg2 *C.GdkRectangle) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.QueryActivatable == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.QueryActivatable, got none")
	}

	var _iter *gtk.TextIter  // out
	var _area *gdk.Rectangle // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_area = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.QueryActivatable(_iter, _area)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtksource5_GutterRendererClass_query_data
func _gotk4_gtksource5_GutterRendererClass_query_data(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkSourceGutterLines, arg2 C.guint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.QueryData == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.QueryData, got none")
	}

	var _lines *GutterLines // out
	var _line uint          // out

	_lines = wrapGutterLines(coreglib.Take(unsafe.Pointer(arg1)))
	_line = uint(arg2)

	overrides.QueryData(_lines, _line)
}

//export _gotk4_gtksource5_GutterRendererClass_snapshot_line
func _gotk4_gtksource5_GutterRendererClass_snapshot_line(arg0 *C.GtkSourceGutterRenderer, arg1 *C.GtkSnapshot, arg2 *C.GtkSourceGutterLines, arg3 C.guint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GutterRendererOverrides](instance0)
	if overrides.SnapshotLine == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GutterRendererOverrides.SnapshotLine, got none")
	}

	var _snapshot *gtk.Snapshot // out
	var _lines *GutterLines     // out
	var _line uint              // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_snapshot = &gtk.Snapshot{
			Snapshot: gdk.Snapshot{
				Object: obj,
			},
		}
	}
	_lines = wrapGutterLines(coreglib.Take(unsafe.Pointer(arg2)))
	_line = uint(arg3)

	overrides.SnapshotLine(_snapshot, _lines, _line)
}

//export _gotk4_gtksource5_GutterRenderer_ConnectQueryData
func _gotk4_gtksource5_GutterRenderer_ConnectQueryData(arg0 C.gpointer, arg1 C.GObject, arg2 C.guint, arg3 C.guintptr) {
	var f func(object *coreglib.Object, p0 uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object *coreglib.Object, p0 uint))
	}

	var _object *coreglib.Object // out
	var _p0 uint                 // out

	_object = coreglib.Take(unsafe.Pointer(&arg1))
	_p0 = uint(arg2)

	f(_object, _p0)
}

//export _gotk4_gtksource5_StyleSchemePreview_ConnectActivate
func _gotk4_gtksource5_StyleSchemePreview_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtksource5_ViewClass_line_mark_activated
func _gotk4_gtksource5_ViewClass_line_mark_activated(arg0 *C.GtkSourceView, arg1 *C.GtkTextIter, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ViewOverrides](instance0)
	if overrides.LineMarkActivated == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ViewOverrides.LineMarkActivated, got none")
	}

	var _iter *gtk.TextIter     // out
	var _button uint            // out
	var _state gdk.ModifierType // out
	var _nPresses int           // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_button = uint(arg2)
	_state = gdk.ModifierType(arg3)
	_nPresses = int(arg4)

	overrides.LineMarkActivated(_iter, _button, _state, _nPresses)
}

//export _gotk4_gtksource5_ViewClass_move_lines
func _gotk4_gtksource5_ViewClass_move_lines(arg0 *C.GtkSourceView, arg1 C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ViewOverrides](instance0)
	if overrides.MoveLines == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ViewOverrides.MoveLines, got none")
	}

	var _down bool // out

	if arg1 != 0 {
		_down = true
	}

	overrides.MoveLines(_down)
}

//export _gotk4_gtksource5_ViewClass_move_words
func _gotk4_gtksource5_ViewClass_move_words(arg0 *C.GtkSourceView, arg1 C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ViewOverrides](instance0)
	if overrides.MoveWords == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ViewOverrides.MoveWords, got none")
	}

	var _step int // out

	_step = int(arg1)

	overrides.MoveWords(_step)
}

//export _gotk4_gtksource5_ViewClass_push_snippet
func _gotk4_gtksource5_ViewClass_push_snippet(arg0 *C.GtkSourceView, arg1 *C.GtkSourceSnippet, arg2 *C.GtkTextIter) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ViewOverrides](instance0)
	if overrides.PushSnippet == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ViewOverrides.PushSnippet, got none")
	}

	var _snippet *Snippet       // out
	var _location *gtk.TextIter // out

	_snippet = wrapSnippet(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_location = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	}

	overrides.PushSnippet(_snippet, _location)
}

//export _gotk4_gtksource5_ViewClass_show_completion
func _gotk4_gtksource5_ViewClass_show_completion(arg0 *C.GtkSourceView) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ViewOverrides](instance0)
	if overrides.ShowCompletion == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ViewOverrides.ShowCompletion, got none")
	}

	overrides.ShowCompletion()
}