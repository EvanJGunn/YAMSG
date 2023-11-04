package minecraft

import (
	"image"
	"log"

	"github.com/EvanJGunn/YAMSG/lib/data"
	"github.com/EvanJGunn/YAMSG/lib/emath"
)

// assuming this is correct, atleast from what I can find online it is
const maxPixels = 16384 //65536

// TODO add scaling, rotation, etc
// TODO loc unused rn
// TODO convert this to using verts to commands maybe
func ImageToCommands(pc ParticleColoredCommand, loc emath.Vector3F, img image.Image) (commands []string) {
	x := img.Bounds().Dx()
	y := img.Bounds().Dy()

	// sample data to make sure it is within maxPixels
	samples := data.Sample2d(x, y, maxPixels)

	for i := 0; i < len(samples); i++ {
		pos := samples[i]
		rgb := img.At(pos.X, pos.Y)

		r, g, b, alpha := rgb.RGBA()
		if alpha == 0 {
			// skip 0 alpha samples
			continue
		}

		// the mc dust command deals in percentages for colors
		rPercentage := float32(float32(r) / maxColorVal)
		gPercentage := float32(float32(g) / maxColorVal)
		bPercentage := float32(float32(b) / maxColorVal)

		// Y = y(max)-pos.y because up is down in png land, and up is up in minecraft land
		cmd := pc.Generate(float32(pos.X), float32(y-pos.Y), 0, emath.Vector3F{X: float64(rPercentage), Y: float64(gPercentage), Z: float64(bPercentage)})
		commands = append(commands, cmd)
	}
	return commands
}

// TODO Loc unused
func VertsToCommands(pc ParticleColoredCommand, loc emath.Vector3F, vertices []data.ObjVertice) (commands []string) {
	if len(vertices) > maxPixels {
		log.Printf("Too many vertices, found %v, max is %v\n", len(vertices), maxPixels)
		return commands
	}

	for i := 0; i < len(vertices); i++ {
		vert := vertices[i]
		cmd := pc.Generate(float32(vert.Pos.X), float32(vert.Pos.Y), float32(vert.Pos.Z), vert.Color)
		commands = append(commands, cmd)
	}

	return commands
}
