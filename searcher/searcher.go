package searcher

import (
	"image"
	"kawaii_search/storage"
)

type Searcher struct {
	hashes *storage.Storage
}

func NewSearcher(hashes *storage.Storage) *Searcher {
	return &Searcher{
		hashes: hashes,
	}
}

func (s* Searcher) Find(img image.Image) ([]string, error) {
	// TODO: do actual searching
	return s.hashes.GetFiles()
}
