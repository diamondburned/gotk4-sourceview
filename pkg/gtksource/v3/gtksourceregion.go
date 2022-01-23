// Code generated by girgen. DO NOT EDIT.

package gtksource

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_source_region_get_type()), F: marshalRegioner},
	})
}

type Region struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Region)(nil)
)

func wrapRegion(obj *externglib.Object) *Region {
	return &Region{
		Object: obj,
	}
}

func marshalRegioner(p uintptr) (interface{}, error) {
	return wrapRegion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
//    - buffer: TextBuffer.
//
// The function returns the following values:
//
//    - region: new SourceRegion object for buffer.
//
func NewRegion(buffer *gtk.TextBuffer) *Region {
	var _arg1 *C.GtkTextBuffer   // out
	var _cret *C.GtkSourceRegion // in

	_arg1 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_source_region_new(_arg1)
	runtime.KeepAlive(buffer)

	var _region *Region // out

	_region = wrapRegion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _region
}

// AddRegion adds region_to_add to region. region_to_add is not modified.
//
// The function takes the following parameters:
//
//    - regionToAdd (optional) to add to region, or NULL.
//
func (region *Region) AddRegion(regionToAdd *Region) {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkSourceRegion // out

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	if regionToAdd != nil {
		_arg1 = (*C.GtkSourceRegion)(unsafe.Pointer(regionToAdd.Native()))
	}

	C.gtk_source_region_add_region(_arg0, _arg1)
	runtime.KeepAlive(region)
	runtime.KeepAlive(regionToAdd)
}

// AddSubregion adds the subregion delimited by _start and _end to region.
//
// The function takes the following parameters:
//
//    - Start: start of the subregion.
//    - End: end of the subregion.
//
func (region *Region) AddSubregion(Start, End *gtk.TextIter) {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(Start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(End)))

	C.gtk_source_region_add_subregion(_arg0, _arg1, _arg2)
	runtime.KeepAlive(region)
	runtime.KeepAlive(Start)
	runtime.KeepAlive(End)
}

// Bounds gets the start and end bounds of the region.
//
// The function returns the following values:
//
//    - start (optional): iterator to initialize with the start of region, or
//      NULL.
//    - end (optional): iterator to initialize with the end of region, or NULL.
//    - ok: TRUE if start and end have been set successfully (if non-NULL), or
//      FALSE if the region is empty.
//
func (region *Region) Bounds() (start *gtk.TextIter, end *gtk.TextIter, ok bool) {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 C.GtkTextIter      // in
	var _arg2 C.GtkTextIter      // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))

	_cret = C.gtk_source_region_get_bounds(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(region)

	var _start *gtk.TextIter // out
	var _end *gtk.TextIter   // out
	var _ok bool             // out

	_start = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_end = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _start, _end, _ok
}

// The function returns the following values:
//
//    - textBuffer (optional): TextBuffer.
//
func (region *Region) Buffer() *gtk.TextBuffer {
	var _arg0 *C.GtkSourceRegion // out
	var _cret *C.GtkTextBuffer   // in

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))

	_cret = C.gtk_source_region_get_buffer(_arg0)
	runtime.KeepAlive(region)

	var _textBuffer *gtk.TextBuffer // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_textBuffer = &gtk.TextBuffer{
				Object: obj,
			}
		}
	}

	return _textBuffer
}

// StartRegionIter initializes a SourceRegionIter to the first subregion of
// region. If region is empty, iter will be initialized to the end iterator.
//
// The function returns the following values:
//
//    - iter: iterator to initialize to the first subregion.
//
func (region *Region) StartRegionIter() *RegionIter {
	var _arg0 *C.GtkSourceRegion    // out
	var _arg1 C.GtkSourceRegionIter // in

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))

	C.gtk_source_region_get_start_region_iter(_arg0, &_arg1)
	runtime.KeepAlive(region)

	var _iter *RegionIter // out

	_iter = (*RegionIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// IntersectRegion returns the intersection between region1 and region2. region1
// and region2 are not modified.
//
// The function takes the following parameters:
//
//    - region2 (optional) or NULL.
//
// The function returns the following values:
//
//    - region (optional): intersection as a SourceRegion object.
//
func (region1 *Region) IntersectRegion(region2 *Region) *Region {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkSourceRegion // out
	var _cret *C.GtkSourceRegion // in

	if region1 != nil {
		_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region1.Native()))
	}
	if region2 != nil {
		_arg1 = (*C.GtkSourceRegion)(unsafe.Pointer(region2.Native()))
	}

	_cret = C.gtk_source_region_intersect_region(_arg0, _arg1)
	runtime.KeepAlive(region1)
	runtime.KeepAlive(region2)

	var _region *Region // out

	if _cret != nil {
		_region = wrapRegion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _region
}

// IntersectSubregion returns the intersection between region and the subregion
// delimited by _start and _end. region is not modified.
//
// The function takes the following parameters:
//
//    - Start: start of the subregion.
//    - End: end of the subregion.
//
// The function returns the following values:
//
//    - ret (optional): intersection as a new SourceRegion.
//
func (region *Region) IntersectSubregion(Start, End *gtk.TextIter) *Region {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out
	var _cret *C.GtkSourceRegion // in

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(Start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(End)))

	_cret = C.gtk_source_region_intersect_subregion(_arg0, _arg1, _arg2)
	runtime.KeepAlive(region)
	runtime.KeepAlive(Start)
	runtime.KeepAlive(End)

	var _ret *Region // out

	if _cret != nil {
		_ret = wrapRegion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _ret
}

// IsEmpty returns whether the region is empty. A NULL region is considered
// empty.
//
// The function returns the following values:
//
//    - ok: whether the region is empty.
//
func (region *Region) IsEmpty() bool {
	var _arg0 *C.GtkSourceRegion // out
	var _cret C.gboolean         // in

	if region != nil {
		_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	}

	_cret = C.gtk_source_region_is_empty(_arg0)
	runtime.KeepAlive(region)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SubtractRegion subtracts region_to_subtract from region. region_to_subtract
// is not modified.
//
// The function takes the following parameters:
//
//    - regionToSubtract (optional) to subtract from region, or NULL.
//
func (region *Region) SubtractRegion(regionToSubtract *Region) {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkSourceRegion // out

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	if regionToSubtract != nil {
		_arg1 = (*C.GtkSourceRegion)(unsafe.Pointer(regionToSubtract.Native()))
	}

	C.gtk_source_region_subtract_region(_arg0, _arg1)
	runtime.KeepAlive(region)
	runtime.KeepAlive(regionToSubtract)
}

// SubtractSubregion subtracts the subregion delimited by _start and _end from
// region.
//
// The function takes the following parameters:
//
//    - Start: start of the subregion.
//    - End: end of the subregion.
//
func (region *Region) SubtractSubregion(Start, End *gtk.TextIter) {
	var _arg0 *C.GtkSourceRegion // out
	var _arg1 *C.GtkTextIter     // out
	var _arg2 *C.GtkTextIter     // out

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))
	_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(Start)))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(End)))

	C.gtk_source_region_subtract_subregion(_arg0, _arg1, _arg2)
	runtime.KeepAlive(region)
	runtime.KeepAlive(Start)
	runtime.KeepAlive(End)
}

// String gets a string represention of region, for debugging purposes.
//
// The returned string contains the character offsets of the subregions. It
// doesn't include a newline character at the end of the string.
//
// The function returns the following values:
//
//    - utf8 (optional): string represention of region. Free with g_free() when
//      no longer needed.
//
func (region *Region) String() string {
	var _arg0 *C.GtkSourceRegion // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GtkSourceRegion)(unsafe.Pointer(region.Native()))

	_cret = C.gtk_source_region_to_string(_arg0)
	runtime.KeepAlive(region)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// RegionIter is an opaque datatype; ignore all its fields. Initialize the iter
// with gtk_source_region_get_start_region_iter().
//
// An instance of this type is always passed by reference.
type RegionIter struct {
	*regionIter
}

// regionIter is the struct that's finalized.
type regionIter struct {
	native *C.GtkSourceRegionIter
}

// Subregion gets the subregion at this iterator.
//
// The function returns the following values:
//
//    - start (optional): iterator to initialize with the subregion start, or
//      NULL.
//    - end (optional): iterator to initialize with the subregion end, or NULL.
//    - ok: TRUE if start and end have been set successfully (if non-NULL), or
//      FALSE if iter is the end iterator or if the region is empty.
//
func (iter *RegionIter) Subregion() (start *gtk.TextIter, end *gtk.TextIter, ok bool) {
	var _arg0 *C.GtkSourceRegionIter // out
	var _arg1 C.GtkTextIter          // in
	var _arg2 C.GtkTextIter          // in
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkSourceRegionIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_region_iter_get_subregion(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(iter)

	var _start *gtk.TextIter // out
	var _end *gtk.TextIter   // out
	var _ok bool             // out

	_start = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_end = (*gtk.TextIter)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _start, _end, _ok
}

// The function returns the following values:
//
//    - ok: whether iter is the end iterator.
//
func (iter *RegionIter) IsEnd() bool {
	var _arg0 *C.GtkSourceRegionIter // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkSourceRegionIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_region_iter_is_end(_arg0)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Next moves iter to the next subregion.
//
// The function returns the following values:
//
//    - ok: TRUE if iter moved and is dereferenceable, or FALSE if iter has been
//      set to the end iterator.
//
func (iter *RegionIter) Next() bool {
	var _arg0 *C.GtkSourceRegionIter // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkSourceRegionIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_source_region_iter_next(_arg0)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}