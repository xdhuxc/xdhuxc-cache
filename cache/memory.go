package cache

import "sync"

type inMemoryCache struct {
	store map[string][]byte
	mutex sync.RWMutex
	Status
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		store:  make(map[string][]byte),
		mutex:  sync.RWMutex{},
		Status: Status{},
	}
}

func (imc *inMemoryCache) Set(k string, v []byte) error {
	imc.mutex.Lock()
	defer imc.mutex.Unlock()

	// 预先处理长度问题
	t, ok := imc.store[k]
	if ok {
		imc.del(k, t)
	}

	imc.store[k] = v
	imc.add(k, v)

	return nil
}

func (imc *inMemoryCache) Get(k string) ([]byte, error) {
	imc.mutex.RLock()
	defer imc.mutex.RUnlock()

	return imc.store[k], nil
}

func (imc *inMemoryCache) Del(k string) error {
	imc.mutex.Lock()
	defer imc.mutex.Unlock()

	v, ok := imc.store[k]
	if ok {
		delete(imc.store, k)
		imc.del(k, v)
	}

	return nil
}

func (imc *inMemoryCache) GetStatus() Status {
	return imc.Status
}
