package minecraft

import (
	"image"

	"github.com/EvanJGunn/YAMSG/lib/data"
	"github.com/EvanJGunn/YAMSG/lib/emath"
)

// assuming this is correct, atleast from what I can find online it is
const maxPixels = 10000 //65536

// TODO add scaling, rotation, etc
func ImageToCommands(pc ParticleColoredCommand, loc emath.Vector3F, img image.Image) (commands []string) {
	x := img.Bounds().Dx()
	y := img.Bounds().Dy()

	// sample data to make sure it is within maxPixels
	samples := data.Sample2d(x, y, maxPixels)

	for i := 0; i < len(samples); i++ {
		pos := samples[i]
		rgb := img.At(pos.X, pos.Y)

		_, _, _, alpha := rgb.RGBA()
		if alpha == 0 {
			// skip 0 alpha samples
			continue
		}

		// Y = y(max)-pos.y because up is down in png land, and up is up in minecraft land
		cmd := pc.Generate(float32(pos.X), float32(y-pos.Y), 0, rgb)
		commands = append(commands, cmd)
	}
	return commands
}
