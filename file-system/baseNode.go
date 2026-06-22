package filesystem

import "time"

type Node struct {
	name    string
	path    string
	modTime time.Time
}

func (n *Node) Name() string            { return n.name }
func (n *Node) Path() string            { return n.path }
func (n *Node) LastModified() time.Time { return n.modTime }
