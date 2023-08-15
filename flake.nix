{
  description = "L1 development flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs =
      import nixpkgs {inherit system;};

    gofumpt = pkgs.buildGoModule rec {
      pname = "gofumpt";
      version = "0.5.0";
      src = pkgs.fetchFromGitHub {
        owner = "mvdan";
        repo = pname;
        rev = "v${version}";
        sha256 = "sha256-3buGLgxAaAIwLXWLpX+K7VRx47DuvUI4W8vw4TuXSts=";
      };
      vendorSha256 = "sha256-W0WKEQgOIFloWsB4E1RTICVKVlj9ChGSpo92X+bjNEk=";
      doCheck = false; # tests require network
    };

    scripts = builtins.attrValues (builtins.mapAttrs pkgs.writeShellScriptBin {
      run-tests = ''
        pushd src
        go test ./... -coverprofile=cover.out
        go tool cover -html=cover.out -o cover.html
        rm cover.out
        popd
      '';
      preview-readme-md = ''
        ${pkgs.python311Packages.grip}/bin/grip --browser --quiet --norefresh
      '';
    });
    back = (with pkgs; [go_1_20 gopls delve go-tools gotools]) ++ [gofumpt];
  in {
    devShells.${system}.default = pkgs.mkShell {
      name = "WB-L1";
      buildInputs = back ++ scripts;
      CGO_ENABLED = 0; # delve from nixpkgs refuses to work otherwise
    };
  };
}
