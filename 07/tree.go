package day07

type directory struct {
	name        string
	size        int
	parent      *directory
	directories []*directory
}

var root = &directory{
	name:        "/",
	size:        0,
	parent:      nil,
	directories: []*directory{},
}

func (d *directory) isRoot() bool {
	return d.name == "/" && d.parent == nil
}

func (d *directory) addFile(name string, size int) int {
	d.size += size
	return d.size
}

func (d *directory) pushDirectory(name string) *directory {
	n := &directory{
		name:        name,
		size:        0,
		parent:      d,
		directories: []*directory{},
	}
	d.directories = append(d.directories, n)
	return n
}

func (d *directory) popDirectory() (*directory, int) {
	p := d.parent
	p.size += d.size
	return p, d.size
}
