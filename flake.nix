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
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go development
            go_1_24
            templ
            air
            golangci-lint
            go-swag
            k6
            imagemagick
            nodePackages.prettier

            # DB (optional, but good to have CLI tools)
            postgresql_16
            sqlite
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
