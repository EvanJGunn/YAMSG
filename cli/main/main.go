package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/EvanJGunn/YAMSG/lib/emath"
	"github.com/EvanJGunn/YAMSG/lib/minecraft"
)

func main() {
	// TODO flags n shtuff

	numArgs := len(os.Args)

	if numArgs < 2 {
		log.Fatal("must provide the file path to an image as an argument")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("failed to close file, err: %v\n", err)
		}
	}()

	parts := strings.Split(file.Name(), ".")
	extension := parts[len(parts)-1]

	switch extension {
	case "png":
		img, err := png.Decode(file)
		if err != nil {
			log.Fatalf("error decoding png file: %v", err)
		}
		pdc := minecraft.NewParticleDustCommand(1, true, emath.Vector3F{X: 0.0001, Y: 0.0001, Z: 0.0001}, 0, 1, minecraft.RenderForce)
		commands := minecraft.ImageToCommands(pdc, emath.Vector3F{X: 0, Y: 1, Z: 0}, img)
		err = exportCommandsToFile(commands, "./bin/pngframe.mcfunction")
		if err != nil {
			log.Fatalf("error exporting commands to file: %v", err)
		}
		break
	case "jpg", "jpeg":
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatalf("error decoding png file: %v", err)
		}
		pdc := minecraft.NewParticleDustCommand(1, true, emath.Vector3F{X: 0.0001, Y: 0.0001, Z: 0.0001}, 0, 1, minecraft.RenderForce)
		commands := minecraft.ImageToCommands(pdc, emath.Vector3F{X: 0, Y: 1, Z: 0}, img)
		err = exportCommandsToFile(commands, "./bin/jpgframe.mcfunction")
		if err != nil {
			log.Fatalf("error exporting commands to file: %v", err)
		}
		break
	default:
		log.Fatalf("file extension currently unsupported: %s", extension)
	}
}

func exportCommandsToFile(commands []string, filename string) error {
	var fileData []byte
	for i := 0; i < len(commands); i++ {
		fileData = append(fileData, []byte(commands[i]+"\n")...)
	}
	err := os.WriteFile(filename, fileData, 0644)
	if err == nil {
		log.Printf("Successfully exported file to %s\n", filename)
	}
	return err
}
