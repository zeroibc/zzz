package static

import (
	"errors"
)

// AssetDirectory is a collection of file groups
type AssetDirectory struct {
	FileGroups map[string]*FileGroup
}

// NewAssetDirectory creates a new asset directory
func NewAssetDirectory() *AssetDirectory {
	return &AssetDirectory{
		FileGroups: make(map[string]*FileGroup),
	}
}

// NewFileGroup creates a new file group
func (a *AssetDirectory) NewFileGroup(baseDirectory string) (*FileGroup, error) {
	if _, exists := a.FileGroups[baseDirectory];exists {
		return nil, errors.New("FileGroup '"+baseDirectory+"' already registered")
	}
	result := &FileGroup{
		baseDirectory:  baseDirectory,
		assetDirectory: make(map[string]string),
	}
	a.FileGroups[baseDirectory] = result
	return result, nil
}

// GetGroup gets a filegroup by name. Returns nil if not found
func (a *AssetDirectory) GetGroup(groupName string) *FileGroup {
	return a.FileGroups[groupName]
}
