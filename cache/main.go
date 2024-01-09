package main

import (
	"fmt"
	"time"
)

const (
	MaxcacheSize int = 1000
	CacheDBBook      = "cache_book"
	CacheDBCD        = "cache_cd"
)

type CacheItem struct {
	Value string
	TTL   time.Time
}

var (
	cacheData map[string]CacheItem
)

func cacheGet(key string) (string, bool) {
	item, found := cacheData[key]
	if found && time.Now().Before(item.TTL) {
		return item.Value, true
	}
	return "", false
}

func cacheSet(key, val string, ttl time.Duration) error {
	if len(cacheData)+1 >= MaxcacheSize {
		return fmt.Errorf("out size cache")
	}
	cacheData[key] = CacheItem{
		Value: val,
		TTL:   time.Now().Add(ttl),
	}
	// goroutine để xóa các key hết hạn
	go func() {
		time.Sleep(ttl)
		delete(cacheData, key)
	}()
	return nil
}

func cacheDel(key string) {
	delete(cacheData, key)
}

func SetBook(key, val string, ttl time.Duration) error {
	err := cacheSet(CacheDBBook+":"+key, val, ttl)
	return err
}

func GetBook(key string) (string, bool) {
	val, ok := cacheGet(CacheDBBook + ":" + key)
	return val, ok
}

func DelBook(key string) {
	cacheDel(CacheDBBook + ":" + key)
}

func SetCD(key, val string, ttl time.Duration) error {
	err := cacheSet(CacheDBCD+":"+key, val, ttl)
	return err
}

func GetCD(key string) (string, bool) {
	val, ok := cacheGet(CacheDBCD + ":" + key)
	return val, ok
}

func DelCD(key string) {
	cacheDel(CacheDBCD + ":" + key)
}

func main() {
	// mô phỏng một hệ thống caching đơn giản để lưu trữ thông tin
	cacheData = make(map[string]CacheItem)
	SetBook("sach kinh te", "100", 300)
	SetCD("CD hoat hinh", "20", 300)
	fmt.Println(GetBook("sach kinh te")) // {"100
	fmt.Println(GetCD("CD hoat hinh"))
}
