# Go Ray Tracer 

This is a simple ray tracer implemented in Go that showcases the use of a factory pattern for object creation and YAML files for scene configurations.

## Features

- Ray tracing implementation in Go.
- Factory pattern for creating scene objects.
- Scene configurations defined in YAML files.

## Prerequisites

Make sure you have Go installed on your machine. If not, you can download it [here](https://golang.org/dl/).

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/chrispotter/gotrace.git
    ```

2. Change into the project directory:

    ```bash
    cd gotrace
    ```

3. Build and run the ray tracer:

    ```bash
    go run cmd/trace/main.go -scene test_scenes/material.yaml -output material
    ```

You should find png files prefixed with material in a folder called `output` in
the main directory

## Configuration

Modify the scene configuration by editing the `scene.yaml` file. This YAML file defines the camera, lights, and objects in the scene.

Example `scene.yaml`:

```yaml
cameras:  
  camera1:
    position: 
      - 0.0
      - 0.0
      - 35.0
    ratio: 
      - 1280.0
      - 720.0
colors:
  lakersPurple:
    color:
      - 253.0
      - 185.0
      - 39.0
  lakersYellow:
    color:
      - 85.0
      - 37.0
      - 130.0
  lightWhite:
    color:
      - 255.0
      - 255.0
      - 255.0
  outlineBlue:
    color:
      - 0.0
      - 0.0
      - 255.0
materials:
  lambert1:
    type: lambert
    color: 
      - lakersPurple 
      - lakersYellow
  cartoon1:
    type: cartoon
    segments: 10
    color: 
      - lakersPurple 
      - lakersYellow
      - lightWhite
      - outlineBlue
shapes:
  sphere1:
    type: sphere
    position:
      - 0.0
      - 0.0
      - 0.0
    radius: 2.0
    material: lambert1
  sphere2:
    type: sphere
    position:
      - -6.0
      - 0.0
      - 0.0
    radius: 2.0
    material: cartoon1
lights:
  dir1:
    type: directional
    view:
      - -1.0
      - -1.5
      - 0.0
    color: lightWhite
```

Feel free to experiment with different camera settings, lights, and objects to create your custom scenes.

## Factory Pattern
The factory pattern is used to create objects dynamically based on their types. Each scene object (e.g., sphere, plane) is created using a factory method, allowing for easy extension with new object types.

Acknowledgments
The `smallfish/simpleyaml` package is used for parsing YAML files.
