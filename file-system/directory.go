package filesystem

import "sync"

type Directory struct {
	Node
	children map[string]INode
	mutex    sync.RWMutex
}

func (d *Directory) Size() int64 {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	var size int64
	for _, child := range d.children {
		size += child.Size()
	}
	return size
}
func (d *Directory) IsDirectory() bool { return true }
