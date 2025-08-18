package metadata

import (
	"GameEngine/internal/service"
	"sync"
)

var (
	metadataOnce     sync.Once
	metadataInstance *metadata
)

type metadata struct{}

func NewMetadata() service.IMetadata {
	metadataOnce.Do(func() {
		metadataInstance = &metadata{}
	})
	return metadataInstance
}

// 确保metadata实现了IMetadata接口
var _ service.IMetadata = (*metadata)(nil)
