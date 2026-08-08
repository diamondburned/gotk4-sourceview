{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";
    nixpkgs-gotk4.url = "github:NixOS/nixpkgs?ref=nixos-26.05";
    gotk4-nix.url = "github:diamondburned/gotk4-nix/main";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      nixpkgs-gotk4,
      gotk4-nix,
      flake-utils,
    }:

    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = gotk4-nix.lib.mkShell {
          base = {
            pname = "gotk4-sourceview";
            buildInputs =
              p: with p; [
                gtksourceview3
                gtksourceview4
                gtksourceview5
                pcre2 # propagated from gtksourceview5
              ];
          };
          pkgs = import nixpkgs-gotk4 {
            inherit system;
            overlays = [ gotk4-nix.overlays.patchelf ];
          };
          packages = with pkgs; [
            go # use nixpkgs'
            self.formatter.${system}
          ];
        };

        formatter = pkgs.nixfmt-rfc-style;
      }
    );
}
