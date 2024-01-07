{
	fetchFromGitHub ? (import <nixpkgs> {}).fetchFromGitHub,

	gotk4-nix ? fetchFromGitHub {
		owner  = "diamondburned";
		repo   = "gotk4-nix";
		rev    = "ad91dabf706946c4380d0a105f0937e4e8ffd75f";
		sha256 = "0rkw9k98qy7ifwypkh2fqhdn7y2qphy2f8xjisj0cyp5pjja62im";
	},

	nixpkgs ? fetchFromGitHub {
		owner = "NixOS";
		repo  = "nixpkgs";
		rev   = "70bdadeb94ffc8806c0570eb5c2695ad29f0e421"; # nixos-23.05
		hash  = "sha256-LWvKHp7kGxk/GEtlrGYV68qIvPHkU9iToomNFGagixU=";
	},
}:

import "${gotk4-nix}/shell.nix" {
	base = {
		pname = "gotk4-sourceview";
		version = "dev";
		buildInputs = pkgs: with pkgs; [
			gtksourceview3
			gtksourceview4
			gtksourceview5
		];
	};
	pkgs = import "${gotk4-nix}/pkgs.nix" {
		sourceNixpkgs = nixpkgs;
		useFetched = true;
	};
}
