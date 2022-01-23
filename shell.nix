let systemPkgs = import <nixpkgs> {};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "2b51f7b62";
		hash  = "sha256:0lcwqkdqqkqsihzashyn3vw8mk6my0f3msxsjw7xsdk9zdv6frpf";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {
		src = systemPkgs.fetchFromGitHub {
			owner = "NixOS";
			repo  = "nixpkgs";
			rev   = "3fdd780";
			hash  = "sha256:0df9v2snlk9ag7jnmxiv31pzhd0rqx2h3kzpsxpj07xns8k8dghz";
		};
	}),
}:

let shell = import "${gotk4}/.nix/shell.nix" {
	inherit pkgs;
};

in shell.overrideAttrs (old: {
	buildInputs = old.buildInputs ++ (with pkgs; [
		gtksourceview3
		gtksourceview4
		gtksourceview5
	]);
})
