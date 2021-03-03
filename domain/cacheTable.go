package domain

type CacheTable struct {
	MemoryBlocks []*MemoryBlock
}

type MemoryBlock struct {
	Id             int
	FreeMemoryLeft int
	Description    string
}
