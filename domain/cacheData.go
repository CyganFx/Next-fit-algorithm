package domain

type CacheData struct {
	MemoryBlocks []*MemoryBlock
}

type MemoryBlock struct {
	Id             int
	FreeMemoryLeft int
	Description    string
}

type LRUCacheData struct {
	CurrentMemoryData string
	PageFaults        int
	PageHits          int
}
