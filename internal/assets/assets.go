package assets

import (
	"crypto/sha256"
	"embed"
	"fmt"
	"sync"
)

var StaticFiles embed.FS
var cache = make(map[string]string)
var mu sync.RWMutex

func Hash(path string) string {
	mu.RLock()
	if h, ok := cache[path]; ok {
		mu.RUnlock()
		return fmt.Sprintf("/%s?v=%s", path, h)
	}
	mu.RUnlock()

	data, err := StaticFiles.ReadFile(path)
	if err != nil {
		return "/" + path
	}

	h := fmt.Sprintf("%x", sha256.Sum256(data))[:8]

	mu.Lock()
	cache[path] = h
	mu.Unlock()

	return fmt.Sprintf("/%s?v=%s", path, h)
}
