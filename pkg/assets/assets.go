package assets

import "embed"

//go:embed seeders/* templates/*
var fs embed.FS

// FS return a filesystem with all files in assets folder
func FS() embed.FS {
	return fs
}
