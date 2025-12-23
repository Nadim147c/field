{ pkgs, ... }:
pkgs.mkShell {
  name = "field";
  inputsFrom = [ (pkgs.callPackage ./package.nix { }) ];
  buildInputs = with pkgs; [
    gnumake
    go
    gofumpt
    golines
    gopls
    revive
  ];
}
