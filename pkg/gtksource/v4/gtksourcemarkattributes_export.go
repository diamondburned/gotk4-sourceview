// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtksourceview/gtksource.h>
import "C"

//export _gotk4_gtksource4_MarkAttributes_ConnectQueryTooltipMarkup
func _gotk4_gtksource4_MarkAttributes_ConnectQueryTooltipMarkup(arg0 C.gpointer, arg1 *C.GtkSourceMark, arg2 C.guintptr) (cret *C.gchar) {
	var f func(mark *Mark) (utf8 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(mark *Mark) (utf8 string))
	}

	var _mark *Mark // out

	_mark = wrapMark(coreglib.Take(unsafe.Pointer(arg1)))

	utf8 := f(_mark)

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_gtksource4_MarkAttributes_ConnectQueryTooltipText
func _gotk4_gtksource4_MarkAttributes_ConnectQueryTooltipText(arg0 C.gpointer, arg1 *C.GtkSourceMark, arg2 C.guintptr) (cret *C.gchar) {
	var f func(mark *Mark) (utf8 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(mark *Mark) (utf8 string))
	}

	var _mark *Mark // out

	_mark = wrapMark(coreglib.Take(unsafe.Pointer(arg1)))

	utf8 := f(_mark)

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}
