package filesystem

import (
	"sync"
	"time"
)

type FileSystem struct {
	root      *Directory
	mutex     sync.RWMutex
	currentWD string
}

func NewFileSystem() *FileSystem {
	root := &Directory{
		Node: Node{
			name:    "/",
			path:    "/",
			modTime: time.Now(),
		},
		children: make(map[string]INode),
	}
	return &FileSystem{
		root:      root,
		currentWD: "/",
	}
}

func (fs *FileSystem) CreateFile(filepath string, content []byte) error {
	return nil
}

func (fs *FileSystem) CreateDirectory(dirPath string) (*Directory, error) {
	return nil, nil
}

func (fs *FileSystem) resolveDirectory(dirpath string) (*Directory, error) {
	return nil, nil
}

func (fs *FileSystem) ListDirectory(dirpath string) ([]string, error) {
	return []string{}, nil
}
func (fs *FileSystem) DeleteFile(filepath string) error {
	return nil
}
