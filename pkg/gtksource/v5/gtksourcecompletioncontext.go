// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

// glib.Type values for gtksourcecompletioncontext.go.
var GTypeCompletionContext = externglib.Type(C.gtk_source_completion_context_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCompletionContext, F: marshalCompletionContext},
	})
}

// CompletionContextOverrider contains methods that are overridable.
type CompletionContextOverrider interface {
}

type CompletionContext struct {
	_ [0]func() // equal guard
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*CompletionContext)(nil)
)

func classInitCompletionContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCompletionContext(obj *externglib.Object) *CompletionContext {
	return &CompletionContext{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalCompletionContext(p uintptr) (interface{}, error) {
	return wrapCompletionContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Activation gets the mode for which the context was activated.
//
// The function returns the following values:
//
func (self *CompletionContext) Activation() CompletionActivation {
	var _arg0 *C.GtkSourceCompletionContext   // out
	var _cret C.GtkSourceCompletionActivation // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_activation(_arg0)
	runtime.KeepAlive(self)

	var _completionActivation CompletionActivation // out

	_completionActivation = CompletionActivation(_cret)

	return _completionActivation
}

// Bounds gets the bounds for the completion, which is the beginning of the
// current word (taking break characters into account) to the current insertion
// cursor.
//
// If begin is non-NULL, it will be set to the start position of the current
// word being completed.
//
// If end is non-NULL, it will be set to the insertion cursor for the current
// word being completed.
//
// The function returns the following values:
//
//    - begin (optional): TextIter.
//    - end (optional): TextIter.
//    - ok: TRUE if the marks are still valid and begin or end was set.
//
func (self *CompletionContext) Bounds() (begin *gtk.TextIter, end *gtk.TextIter, ok bool) {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _arg1 C.GtkTextIter                 // in
	var _arg2 C.GtkTextIter                 // in
	var _cret C.gboolean                    // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_bounds(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(self)

	var _begin *gtk.TextIter // out
	var _end *gtk.TextIter   // out
	var _ok bool             // out

	_begin = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_end = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _begin, _end, _ok
}

// Buffer gets the underlying buffer used by the context.
//
// This is a convenience function to get the buffer via the SourceCompletion
// property.
//
// The function returns the following values:
//
//    - buffer (optional) or NULL.
//
func (self *CompletionContext) Buffer() *Buffer {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret *C.GtkSourceBuffer            // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_buffer(_arg0)
	runtime.KeepAlive(self)

	var _buffer *Buffer // out

	if _cret != nil {
		_buffer = wrapBuffer(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _buffer
}

// Busy gets the "busy" property. This is set to TRUE while the completion
// context is actively fetching proposals from registered
// SourceCompletionProvider's.
//
// The function returns the following values:
//
//    - ok: TRUE if the context is busy.
//
func (self *CompletionContext) Busy() bool {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_busy(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Completion gets the SourceCompletion that created the context.
//
// The function returns the following values:
//
//    - completion (optional) or NULL.
//
func (self *CompletionContext) Completion() *Completion {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret *C.GtkSourceCompletion        // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_completion(_arg0)
	runtime.KeepAlive(self)

	var _completion *Completion // out

	if _cret != nil {
		_completion = wrapCompletion(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _completion
}

// Empty checks if any proposals have been provided to the context.
//
// Out of convenience, this function will return TRUE if self is NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if there are no proposals in the context.
//
func (self *CompletionContext) Empty() bool {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_empty(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Language gets the language of the underlying buffer, if any.
//
// The function returns the following values:
//
//    - language (optional) or NULL.
//
func (self *CompletionContext) Language() *Language {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret *C.GtkSourceLanguage          // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_language(_arg0)
	runtime.KeepAlive(self)

	var _language *Language // out

	if _cret != nil {
		_language = wrapLanguage(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _language
}

// View gets the text view for the context.
//
// The function returns the following values:
//
//    - view (optional) or NULL.
//
func (self *CompletionContext) View() *View {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret *C.GtkSourceView              // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_view(_arg0)
	runtime.KeepAlive(self)

	var _view *View // out

	if _cret != nil {
		_view = wrapView(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _view
}

// Word gets the word that is being completed up to the position of the insert
// mark.
//
// The function returns the following values:
//
//    - utf8: string containing the current word.
//
func (self *CompletionContext) Word() string {
	var _arg0 *C.GtkSourceCompletionContext // out
	var _cret *C.char                       // in

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_source_completion_context_get_word(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetProposalsForProvider: this function allows providers to update their
// results for a context outside of a call to
// gtk_source_completion_provider_populate_async(). This can be used to
// immediately return results for a provider while it does additional
// asynchronous work. Doing so will allow the completions to update while the
// operation is in progress.
//
// The function takes the following parameters:
//
//    - provider: SourceCompletionProvider.
//    - results (optional) or NULL.
//
func (self *CompletionContext) SetProposalsForProvider(provider CompletionProviderer, results gio.ListModeller) {
	var _arg0 *C.GtkSourceCompletionContext  // out
	var _arg1 *C.GtkSourceCompletionProvider // out
	var _arg2 *C.GListModel                  // out

	_arg0 = (*C.GtkSourceCompletionContext)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkSourceCompletionProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	if results != nil {
		_arg2 = (*C.GListModel)(unsafe.Pointer(externglib.InternObject(results).Native()))
	}

	C.gtk_source_completion_context_set_proposals_for_provider(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(results)
}
