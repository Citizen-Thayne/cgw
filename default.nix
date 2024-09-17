{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.air
    pkgs.watchexec
    pkgs.flyio
  ];
}
