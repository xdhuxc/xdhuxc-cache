package cache

import (
	log "github.com/sirupsen/logrus"
)

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStatus() Status
}

type Status struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Status) add(k string, v []byte) {
	s.Count = s.Count + 1
	s.KeySize = s.KeySize + int64(len(k))
	s.ValueSize = s.ValueSize + int64(len(v))
}

func (s *Status) del(k string, v []byte) {
	s.Count = s.Count - 1
	s.KeySize = s.KeySize - int64(len(k))
	s.ValueSize = s.ValueSize - int64(len(v))
}

func New(kind string) Cache {
	var c Cache

	if kind == "inmemory" {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknown cache kind " + kind)
	}

	log.Println(kind, "ready to serve")

	return c
}
