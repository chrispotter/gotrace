package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/chrispotter/trace/internal/scene"
)

func main() {
	var sceneFile, outputFile string
	flag.StringVar(&sceneFile, "scene", "scene/test.yml", "This is the scene file to be loaded for rendering")
	flag.StringVar(&outputFile, "output", "test.png", "This file will be what images will be called in the output folder")

	flag.Parse()

	content, err := ioutil.ReadFile(sceneFile)
	if err != nil {
		log.Fatal(err)
	}

	s := scene.Scene{}

	err = s.FromYaml(content)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Render(outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
