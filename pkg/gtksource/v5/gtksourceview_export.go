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

//export _gotk4_gtksource5_View_ConnectChangeCase
func _gotk4_gtksource5_View_ConnectChangeCase(arg0 C.gpointer, arg1 C.GtkSourceChangeCaseType, arg2 C.guintptr) {
	var f func(caseType ChangeCaseType)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(caseType ChangeCaseType))
	}

	var _caseType ChangeCaseType // out

	_caseType = ChangeCaseType(arg1)

	f(_caseType)
}

//export _gotk4_gtksource5_View_ConnectChangeNumber
func _gotk4_gtksource5_View_ConnectChangeNumber(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(count int))
	}

	var _count int // out

	_count = int(arg1)

	f(_count)
}

//export _gotk4_gtksource5_View_ConnectJoinLines
func _gotk4_gtksource5_View_ConnectJoinLines(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtksource5_View_ConnectLineMarkActivated
func _gotk4_gtksource5_View_ConnectLineMarkActivated(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.gint, arg5 C.guintptr) {
	var f func(iter *gtk.TextIter, button uint, state gdk.ModifierType, nPresses int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg5))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *gtk.TextIter, button uint, state gdk.ModifierType, nPresses int))
	}

	var _iter *gtk.TextIter     // out
	var _button uint            // out
	var _state gdk.ModifierType // out
	var _nPresses int           // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_button = uint(arg2)
	_state = gdk.ModifierType(arg3)
	_nPresses = int(arg4)

	f(_iter, _button, _state, _nPresses)
}

//export _gotk4_gtksource5_View_ConnectMoveLines
func _gotk4_gtksource5_View_ConnectMoveLines(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(down bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(down bool))
	}

	var _down bool // out

	if arg1 != 0 {
		_down = true
	}

	f(_down)
}

//export _gotk4_gtksource5_View_ConnectMoveToMatchingBracket
func _gotk4_gtksource5_View_ConnectMoveToMatchingBracket(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(extendSelection bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(extendSelection bool))
	}

	var _extendSelection bool // out

	if arg1 != 0 {
		_extendSelection = true
	}

	f(_extendSelection)
}

//export _gotk4_gtksource5_View_ConnectMoveWords
func _gotk4_gtksource5_View_ConnectMoveWords(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(count int))
	}

	var _count int // out

	_count = int(arg1)

	f(_count)
}

//export _gotk4_gtksource5_View_ConnectShowCompletion
func _gotk4_gtksource5_View_ConnectShowCompletion(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtksource5_View_ConnectSmartHomeEnd
func _gotk4_gtksource5_View_ConnectSmartHomeEnd(arg0 C.gpointer, arg1 *C.GtkTextIter, arg2 C.gint, arg3 C.guintptr) {
	var f func(iter *gtk.TextIter, count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(iter *gtk.TextIter, count int))
	}

	var _iter *gtk.TextIter // out
	var _count int          // out

	_iter = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_count = int(arg2)

	f(_iter, _count)
}