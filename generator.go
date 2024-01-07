package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const sourceviewModule = "github.com/diamondburned/gotk4-sourceview/pkg"

var data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: sourceviewModule,
		Packages: []genmain.Package{
			{Name: "gtksourceview-3.0", Namespaces: nil},
			{Name: "gtksourceview-4", Namespaces: nil},
			{Name: "gtksourceview-5", Namespaces: nil},
		},
		PkgGenerated: []string{
			"gtksource",
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
		},
		Filters: []types.FilterMatcher{
			// /nix/store/kmqs0wll31ylwbqkpmlgbjrn6ny3myik-binutils-2.35.1/bin/ld: $WORK/b069/_x006.o: in function `_cgo_f791a2727c11_Cfunc_gtk_source_completion_context_get_start_iter':
			// gtksourcecompletioncontext.cgo2.c:(.text+0x140): undefined reference to `gtk_source_completion_context_get_start_iter'
			// /nix/store/kmqs0wll31ylwbqkpmlgbjrn6ny3myik-binutils-2.35.1/bin/ld: $WORK/b069/_x016.o: in function `_cgo_f791a2727c11_Cfunc_gtk_source_gutter_lines_get_yrange':
			// gtksourcegutterlines.cgo2.c:(.text+0x14d): undefined reference to `gtk_source_gutter_lines_get_yrange'
			types.AbsoluteFilter("C.gtk_source_completion_context_get_start_iter"),
			types.AbsoluteFilter("C.gtk_source_gutter_lines_get_yrange"),
		},
	},
)

func main() {
	genmain.Run(data)
}
