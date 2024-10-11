# Use whatever method you prefer to get access to nixpkgs.
{ pkgs ? import <nixpkgs> {}}:
let
  inherit (pkgs)
    buildFHSEnv;

  fhsEnv = 
    buildFHSEnv {
      name = "sqlite-with-libspatialite";
      targetPkgs = pkgs:
        builtins.attrValues {
          inherit (pkgs)
            sqlite libspatialite
          ;
        };

      # Ensure /usr/lib is part of the library search path.
      profile = ''
        export LD_LIBRARY_PATH="/usr/lib"
      '';
    };
in

fhsEnv.env
