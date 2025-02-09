{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {

  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.delve
  ];

  shellHook = ''
    export GOPATH=$(pwd)/.go
    export GOBIN=$GOPATH/bin
    export PATH=$GOBIN:$PATH
    mkdir -p $GOPATH
  '';
}