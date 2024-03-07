{ pkgs ? import <nixpkgs> {} }:
with pkgs;

pkgs.mkShell {

  nativeBuildInputs = with pkgs.buildPackages; [
    pkgs.go_1_21 # 1.21.6  # nix-env -qP --available go
    pkgs.gopls # 0.14.2
  ];
}
