package filesystem

import "time"

type INode interface {
	Name() string
	Path() string
	Size() int64
	IsDirectory() bool
	LastModified() time.Time
}
