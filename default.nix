{ lib, buildGoModule }:
buildGoModule {
  pname = "coffetime";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-i2FG/Dlw0r5PVHak+37VBeRwG7Vf7qWNlYzNyJUIURg=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "coffetime";
    homepage = "https://github.com/comfysage/coffetime";
    maintainers = with lib.maintainers; [ comfysage ];
    mainPackage = "coffetime";
  };
}
