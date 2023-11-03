package data

import (
	"log"
	"math"

	"github.com/EvanJGunn/YAMSG/lib/emath"
)

// preserves ratio as best it can, but kind of results in scan lines, probably better sampling methods out there
// "proportional constraint" problem
func Sample2d(x, y, max int) (samples []emath.Vector2I) {
	if x*y < max {
		// sample everything if its already under the max
		for i := 0; i <= y; i++ {
			for j := 0; j <= x; j++ {
				samples = append(samples, emath.Vector2I{X: j, Y: i})
			}
		}
		return samples
	}

	log.Println("Image over max pixels, sampling to a lower resolution - TODO this is garbo that doesn't work yet, need to fix.")

	// determine the new x and y sizes by scaling down to within the max
	// while preserving the aspect ratio
	newX := math.Floor(math.Sqrt((float64(x) * float64(max)) / float64(y)))
	newY := math.Floor(math.Sqrt((float64(y) * float64(max)) / float64(x)))

	log.Println(newX)
	log.Println(newY)

	periodX := float64(x) / newX
	periodY := float64(y) / newY

	log.Println(periodX)
	log.Println(periodY)

	totalY := float64(0)
	for int(math.Ceil(totalY)) < int(newY) {
		totalX := float64(0)
		for int(math.Ceil(totalX)) < int(newX) {
			samples = append(samples, emath.Vector2I{X: int(math.Ceil(totalX)), Y: int(math.Ceil(totalY))})
			totalX += periodX
		}
		totalY += periodY
	}

	return samples
}
