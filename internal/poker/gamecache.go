package poker

import (
	"github.com/hashicorp/golang-lru"
)

type GameCache interface {
	Get(string) Game
	Set(string, Game)
}

func NewGameCache() GameCache {
	c, _ := lru.New(1e4)
	return &gamecache{
		cache: c,
	}
}

type gamecache struct {
	cache *lru.Cache
}

func (c *gamecache) Get(hash string) (res Game) {
	obj, success := c.cache.Get(hash)
	if success {
		res = obj.(*game)
	}
	return res
}

func (c *gamecache) Set(hash string, game Game) {
	c.cache.Add(hash, game)
}
