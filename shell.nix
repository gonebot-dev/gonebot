{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {

  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.delve
    pkgs.go-tools
  ];

  shellHook = ''
    go env -w GO111MODULE=on
  '';
}