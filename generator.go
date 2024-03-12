package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4-sourceview/gensourceview"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

func main() {
	genmain.Run(genmain.Overlay(
		gendata.Main,
		gensourceview.Data,
	))
}
