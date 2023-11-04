package minecraft

import (
	"fmt"

	"github.com/EvanJGunn/YAMSG/lib/emath"
)

type Render string

const (
	RenderForce  = Render("force")  // 512 blocks
	RenderNormal = Render("normal") // less than 512 blocks, I forget :(

	particleDustFormat = "particle minecraft:dust %v %v %v %f %s %f %f %f %f %d %s"

	maxColorVal = 65535.0
)

type ParticleColoredCommand interface {
	Generate(x, y, z float32, rgb emath.Vector3F) string
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

func (pc particleDustCommand) Generate(x, y, z float32, rgb emath.Vector3F) string {
	if pc.relativePos {
		return fmt.Sprintf(particleDustFormat, rgb.X, rgb.Y, rgb.Z, pc.size, fmt.Sprintf("~%f ~%f ~%f", x, y, z), pc.boxDir.X, pc.boxDir.Y, pc.boxDir.Z, pc.speed, pc.count, string(pc.render))
	}
	return fmt.Sprintf(particleDustFormat, rgb.X, rgb.Y, rgb.Z, pc.size, fmt.Sprintf("%f %f %f", x, y, z), pc.boxDir.X, pc.boxDir.Y, pc.boxDir.Z, pc.speed, pc.count, string(pc.render))
}
