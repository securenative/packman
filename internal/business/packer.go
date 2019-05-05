package business

import "github.com/securenative/packman/internal/data"

type PackmanPacker struct {
	backend data.Backend
}

func (this *PackmanPacker) Pack(name string, sourcePath string) error {
	return this.backend.Push(name, sourcePath)
}
