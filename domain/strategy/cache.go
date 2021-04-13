package strategy

type Cache struct {
	Storage      map[string]string
	EvictionAlgo EvictionAlgo
	Capacity     int
	MaxCapacity  int
}

