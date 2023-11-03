package minecraft

import (
	"image"

	"github.com/EvanJGunn/YAMSG/lib/data"
	"github.com/EvanJGunn/YAMSG/lib/emath"
)

// assuming this is correct, atleast from what I can find online it is
const maxPixels = 65536

// TODO add scaling, rotation, etc
func ImageToCommands(pc ParticleColoredCommand, loc emath.Vector3F, img image.Image) (commands []string) {
	x := img.Bounds().Dx()
	y := img.Bounds().Dy()

	// sample data to make sure it is within maxPixels
	samples := data.Sample2d(x, y, maxPixels)

	for i := 0; i < len(samples); i++ {
		pos := samples[i]
		rgb := img.Bounds().RGBA64At(pos.X, pos.Y)
		if rgb.A == 0 {
			// skip 0 alpha samples
			continue
		}

		cmd := pc.Generate(float32(pos.X), float32(pos.Y), 0, emath.Vector3F{X: float64(rgb.R), Y: float64(rgb.G), Z: float64(rgb.B)})
		commands = append(commands, cmd)
	}
	return commands
}
