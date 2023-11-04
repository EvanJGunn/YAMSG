package data

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/EvanJGunn/YAMSG/lib/emath"
)

type ObjVertice struct {
	Pos   emath.Vector3F
	Color emath.Vector3F
}

func ObjFileToVerts(file *os.File) (vertices []ObjVertice, err error) {
	// Manually doing color for now until I learn about texture mappings - its Saturday and I am too lazy rn
	megamindBlue := emath.Vector3F{X: 0.211764, Y: 0.501960, Z: 0.968627}

	rdr := bufio.NewScanner(file)
	for rdr.Scan() {
		line := rdr.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "v" && len(parts) == 4 {
			x, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, err
			}
			z, err := strconv.ParseFloat(parts[3], 64)
			if err != nil {
				return nil, err
			}
			vertices = append(vertices, ObjVertice{Pos: emath.Vector3F{X: x, Y: y, Z: z}, Color: megamindBlue})
		}
	}

	if err = rdr.Err(); err != nil {
		return nil, err
	}
	return vertices, nil
}
