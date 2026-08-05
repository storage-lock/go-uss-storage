package uss_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// USSStorage UCloud对象存储实现 (S3兼容)
type USSStorage struct {
	*s3_storage.S3Storage
}

func NewUSSStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*USSStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &USSStorage{S3Storage: s3Storage}, nil
}

const StorageName = "uss-storage"

func (x *USSStorage) GetName() string {
	return StorageName
}

func (x *USSStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
