# Use whatever method you prefer to get access to nixpkgs.
{ pkgs ? import <nixpkgs> {}}:
let
  inherit (pkgs)
    buildFHSEnv;

  fhsEnv = 
    buildFHSEnv {
      name = "geos";
      targetPkgs = pkgs:
        builtins.attrValues {
          inherit (pkgs)
            geos
          ;
        };

      # Ensure /usr/lib is part of the library search path.
      profile = ''
		export LD_LIBRARY_PATH="${pkgs.geos}/lib:$LD_LIBRARY_PATH"
      export CPATH="${pkgs.geos}/include:$CPATH"   # Add the path to geos_c.h
      '';
    };
in

fhsEnv.env
