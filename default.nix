{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
      import (fetchTree nixpkgs.locked) {
        overlays = [
          (import "${fetchTree gomod2nix.locked}/overlay.nix")
        ];
      }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
}:
buildGoApplication rec {
  pname = "cypher";
  version = "0.0.1";
  src = ./.;
  modules = ./gomod2nix.toml;
  installPhase = ''
    mkdir -p $out/bin
    cp $GOPATH/bin/src $out/bin/${pname}
  '';
}
