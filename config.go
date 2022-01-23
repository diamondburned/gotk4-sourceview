package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module = "github.com/diamondburned/gotk4/pkg"
	thisModule  = "github.com/diamondburned/gotk4-sourceview/pkg"
)

var packages = []gendata.Package{
	{PkgName: "gtksourceview-3.0", Namespaces: nil},
	{PkgName: "gtksourceview-4", Namespaces: nil},
	{PkgName: "gtksourceview-5", Namespaces: nil},
}

var pkgGenerated = []string{
	"gtksource",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var preprocessors = []types.Preprocessor{}

var filters = []types.FilterMatcher{
	// /nix/store/kmqs0wll31ylwbqkpmlgbjrn6ny3myik-binutils-2.35.1/bin/ld: $WORK/b069/_x006.o: in function `_cgo_f791a2727c11_Cfunc_gtk_source_completion_context_get_start_iter':
	// gtksourcecompletioncontext.cgo2.c:(.text+0x140): undefined reference to `gtk_source_completion_context_get_start_iter'
	// /nix/store/kmqs0wll31ylwbqkpmlgbjrn6ny3myik-binutils-2.35.1/bin/ld: $WORK/b069/_x016.o: in function `_cgo_f791a2727c11_Cfunc_gtk_source_gutter_lines_get_yrange':
	// gtksourcegutterlines.cgo2.c:(.text+0x14d): undefined reference to `gtk_source_gutter_lines_get_yrange'
	types.AbsoluteFilter("C.gtk_source_completion_context_get_start_iter"),
	types.AbsoluteFilter("C.gtk_source_gutter_lines_get_yrange"),
}
