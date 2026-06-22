package filesystem

type File struct {
	Node
	size    int64
	content []byte
}

func (f *File) Size() int64       { return f.size }
func (f *File) IsDirectory() bool { return false }
