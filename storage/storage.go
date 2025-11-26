package storage

import "3-bin/bins"

type Storage struct {
	bins []bins.Bin
}

func NewStorage() *Storage {
	return &Storage{}
}

func (storage *Storage) PutBin(bin *bins.Bin) {
	storage.bins = append(storage.bins, *bin)
}

func (storage *Storage) GetBins() *[]bins.Bin {
	return &storage.bins
}
