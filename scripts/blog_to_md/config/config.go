package config

const ImageDir = "tmp/images"

var mdDir string

func SetMdDir(dir string) {
	mdDir = dir
}

func GetMdDir() string {
	return mdDir
}
