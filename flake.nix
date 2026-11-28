{
  description = "Development environment for workout-tracker";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };

        prettier_3_0_0 = pkgs.stdenv.mkDerivation {
          name = "prettier";
          version = "3.0.0";
          src = pkgs.fetchurl {
            url = "https://registry.npmjs.org/prettier/-/prettier-3.0.0.tgz";
            hash = "sha256-3aOqO+7Hyc76KgcWYo1kQPMeXhCssaXrZOpDN601gDA=";
          };
          nativeBuildInputs = [ pkgs.makeWrapper ];
          unpackPhase = "tar -xf $src";
          installPhase = ''
            mkdir -p $out/bin $out/lib
            cp -r package/* $out/lib/
            makeWrapper ${pkgs.nodejs}/bin/node $out/bin/prettier \
              --add-flags $out/lib/bin/prettier.cjs
          '';
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go development
            go
            templ
            air
            golangci-lint
            go-swag
            k6
            imagemagick
            prettier_3_0_0

            # DB (optional, but good to have CLI tools)
            postgresql_16
            sqlite

            # Helpers for scripts/reset-testdata.sh
            curl
            jq
            lsof
          ];

          shellHook = ''
            echo "🏋️ Workout Tracker development environment loaded!"
            echo "Go: $(go version)"
            echo "Node: $(node --version)"
            
            # Ensure templ and other tools are in PATH if installed via go install
            export PATH=$PATH:$(go env GOPATH)/bin
          '';
        };
      });
}
