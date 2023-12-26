package camera

// Factory returns an array of Cameras from an array of Camera Configs
func Factory(configs []*Config) ([]*Camera, error) {
	cameraMap := []*Camera{}
	for _, config := range configs {
		camera, err := NewCamera(config.Position, config.Ratio)
		if err != nil {
			return nil, err
		}
		camera.Name = config.Name
		cameraMap = append(cameraMap, camera)
	}

	return cameraMap, nil
}
