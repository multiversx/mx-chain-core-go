package storage

import (
	"errors"
)

// ErrCacheSizeInvalid signals that size of cache is less than 1
var ErrCacheSizeInvalid = errors.New("cache size is less than 1")

// ErrCacheCapacityInvalid signals that capacity of cache is less than 1
var ErrCacheCapacityInvalid = errors.New("cache capacity is less than 1")
