package minecraft

import (
	"fmt"
	"image/color"

	"github.com/EvanJGunn/YAMSG/lib/emath"
)

type Render string

const (
	RenderForce  = Render("force")  // 512 blocks
	RenderNormal = Render("normal") // less than 512 blocks, I forget :(

	particleDustFormat = "/particle minecraft:dust %v %v %v %f %s %f %f %f %f %d %s"
)

type ParticleColoredCommand interface {
	Generate(x, y, z float32, rgb color.Color) string
}

type particleDustCommand struct {
	size        float32
	relativePos bool
	boxDir      emath.Vector3F // boxdir can be multiple things... I think its the box the particle can spawn in
	speed       float32
	count       int32
	render      Render
}

func NewParticleDustCommand(size float32, relativePos bool, boxDir emath.Vector3F, speed float32, count int32, render Render) ParticleColoredCommand {
	return particleDustCommand{
		size:        size,
		relativePos: relativePos,
		boxDir:      boxDir,
		speed:       speed,
		count:       count,
		render:      render,
	}
}

func (pc particleDustCommand) Generate(x, y, z float32, rgb color.Color) string {
	r, g, b, _ := rgb.RGBA()

	if pc.relativePos {
		return fmt.Sprintf(particleDustFormat, r, g, b, pc.size, fmt.Sprintf("~%f ~%f ~%f", x, y, z), pc.boxDir.X, pc.boxDir.Y, pc.boxDir.Z, pc.speed, pc.count, string(pc.render))
	}
	return fmt.Sprintf(particleDustFormat, r, g, b, pc.size, fmt.Sprintf("%f %f %f", x, y, z), pc.boxDir.X, pc.boxDir.Y, pc.boxDir.Z, pc.speed, pc.count, string(pc.render))
}
