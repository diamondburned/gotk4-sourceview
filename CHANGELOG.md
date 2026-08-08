# Changelog

All notable changes to `gotk4-sourceview` will be documented in this file.

## [0.4.0] - 2026-08-08

### Version Upgrades

- Updated GtkSourceView versions: `gtksourceview-5` to `5.20.0`, `gtksourceview-4` to `4.8.4`, and `gtksourceview-3` to `3.24.11`.
- Updated Go toolchain to `1.25.0`.
- Updated Nixpkgs channel ref to `nixos-26.05`.
- Upgraded `github.com/diamondburned/gotk4` dependency to `v0.4.0`.

### GIR Code Generation & Generator Fixes

- Re-generated Go bindings against GtkSourceView 3, 4, and 5 GIR definitions using `girgen`.

### New APIs Overview

#### GtkSourceView 5.18

- `Annotation` (`GtkSourceAnnotation`): Represents an annotation added to view displayed at the end of a line (`NewAnnotation`, `Description`, `Icon`, `Line`, `Style`).
- `AnnotationProvider` (`GtkSourceAnnotationProvider`): Provides annotations and populates hover displays (`NewAnnotationProvider`, `AddAnnotation`, `RemoveAnnotation`, `RemoveAll`, `PopulateHoverAsync`, `PopulateHoverFinish`, `ConnectChanged`).
- `Annotations` (`GtkSourceAnnotations`): Manager for view annotations (`AddProvider`, `RemoveProvider`, `ConnectChanged`).
- `View` (`GtkSourceView`): Added `Annotations` method (`Annotations`).
