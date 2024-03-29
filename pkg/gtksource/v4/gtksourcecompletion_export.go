// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

//export _gotk4_gtksource4_Completion_ConnectActivateProposal
func _gotk4_gtksource4_Completion_ConnectActivateProposal(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtksource4_Completion_ConnectHide
func _gotk4_gtksource4_Completion_ConnectHide(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtksource4_Completion_ConnectMoveCursor
func _gotk4_gtksource4_Completion_ConnectMoveCursor(arg0 C.gpointer, arg1 C.GtkScrollStep, arg2 C.gint, arg3 C.guintptr) {
	var f func(step gtk.ScrollStep, num int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step gtk.ScrollStep, num int))
	}

	var _step gtk.ScrollStep // out
	var _num int             // out

	_step = gtk.ScrollStep(arg1)
	_num = int(arg2)

	f(_step, _num)
}

//export _gotk4_gtksource4_Completion_ConnectMovePage
func _gotk4_gtksource4_Completion_ConnectMovePage(arg0 C.gpointer, arg1 C.GtkScrollStep, arg2 C.gint, arg3 C.guintptr) {
	var f func(step gtk.ScrollStep, num int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step gtk.ScrollStep, num int))
	}

	var _step gtk.ScrollStep // out
	var _num int             // out

	_step = gtk.ScrollStep(arg1)
	_num = int(arg2)

	f(_step, _num)
}

//export _gotk4_gtksource4_Completion_ConnectPopulateContext
func _gotk4_gtksource4_Completion_ConnectPopulateContext(arg0 C.gpointer, arg1 *C.GtkSourceCompletionContext, arg2 C.guintptr) {
	var f func(context *CompletionContext)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *CompletionContext))
	}

	var _context *CompletionContext // out

	_context = wrapCompletionContext(coreglib.Take(unsafe.Pointer(arg1)))

	f(_context)
}

//export _gotk4_gtksource4_Completion_ConnectShow
func _gotk4_gtksource4_Completion_ConnectShow(arg0 C.gpointer, arg1 C.guintptr) {
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
