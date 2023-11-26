package sharded

import (
	"crypto/sha1"
	"sync"
)

type shard struct {
	sync.RWMutex
	m map[string]interface{}
}
type ShardMap []*shard

func NewShardedMap(numberOfShard int) ShardMap {
	shards := make([]*shard, numberOfShard)

	for i := 0; i < numberOfShard; i++ {
		shards[i] = &shard{
			m: make(map[string]interface{}),
		}
	}

	return shards
}

func (m ShardMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[17])
	return hash % len(m)
}

func (m ShardMap) getShard(key string) *shard {
	index := m.getShardIndex(key)
	return m[index]
}

func (m ShardMap) Get(key string) interface{} {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	return shard.m[key]
}

func (m ShardMap) Set(key string, val interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.m[key] = val
}
