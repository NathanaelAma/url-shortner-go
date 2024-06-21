package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/NathanaelAma/url-shortener/model"
	"sync"
)

var (
	ErrInvalidLongUrl = errors.New("valid long URL is required")
)

var urlStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func ShortenUrl(url model.Url) (model.Url, error) {
	if url.Long == "" {
		return model.Url{}, ErrInvalidLongUrl
	}

	if url.Short == "" {
		url.Short = GenerateShortUrl(url.Long)
	}

	urlStore.Lock()
	urlStore.m[url.Short] = url.Long
	urlStore.Unlock()

	return url, nil
}

func GenerateShortUrl(longUrl string) string {
	hash := sha256.New()
	hash.Write([]byte(longUrl))
	hashSum := hash.Sum(nil)

	base64Encoded := base64.URLEncoding.EncodeToString(hashSum)

	shortUrl := base64Encoded[:8]

	return shortUrl
}

func GetLongUrl(shortUrl string) (string, error) {
	urlStore.RLock()
	defer urlStore.RUnlock()

	longUrl, exists := urlStore.m[shortUrl]
	if !exists {
		return "", errors.New("short URL not found")
	}

	return longUrl, nil
}
