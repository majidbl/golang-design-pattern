package strategy

import "design_pattern/pattern/strategy"

func main() {

	lfu := &strategy.LFU{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &strategy.LRU{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &strategy.FIFO{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}