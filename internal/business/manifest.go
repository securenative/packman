package business

import "path/filepath"

// Encapsulates all logic required for packing a project
type Packer interface {
	Pack(name string, sourcePath string) error
}

// Encapsulates all logic required for unpacking a project
type Unpacker interface {
	Unpack(name string, destPath string, args []string) error
}

// Encapsulates all logic required for initializing new package
type ProjectInit interface {
	Init(destPath string) error
}

func packmanPath(destPath string) string {
	return filepath.Join(destPath, "packman")
}
