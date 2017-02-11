package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type FS struct {
	root string
}

type List struct {
	Parent Dir
	Dirs   []Dir
	Files  []File
}

type Dir struct {
	Name    string
	Path    string
	LastMod time.Time
}

type File struct {
	Name    string
	Path    string
	LastMod time.Time
	Size    int64
}

func New(root string) *FS {
	return &FS{
		root: filepath.Clean(root),
	}
}

func (fs *FS) Abs(rel string) string {
	rel = path.Join("/", rel)
	return filepath.Join(fs.root, rel)
}

func (fs *FS) IsFile(rel string) bool {
	stat, err := os.Stat(fs.Abs(rel))
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func (fs *FS) List(rel string) (*List, error) {
	rel = path.Join("/", rel)
	abs := filepath.Join(fs.root, rel)

	stat, err := os.Lstat(abs)
	if err != nil {
		return nil, err
	}

	list := &List{
		Dirs:  []Dir{},
		Files: []File{},
	}

	if stat.IsDir() {
		list.Parent = Dir{
			Name:    stat.Name(),
			Path:    rel,
			LastMod: stat.ModTime(),
		}

		fis, err := ioutil.ReadDir(abs)
		if err != nil {
			return nil, err
		}

		for _, fi := range fis {
			name := fi.Name()
			if strings.HasPrefix(name, ".") {
				continue
			}
			if fi.IsDir() {
				list.Dirs = append(list.Dirs, Dir{
					Name:    name,
					Path:    path.Join(rel, name),
					LastMod: fi.ModTime(),
				})
			} else {
				list.Files = append(list.Files, File{
					Name:    name,
					Path:    path.Join(rel, name),
					LastMod: fi.ModTime(),
					Size:    fi.Size(),
				})
			}
		}

		return list, nil
	}

	parentPath := path.Dir(rel)
	parentStat, err := os.Stat(filepath.Join(fs.root, parentPath))
	if err != nil {
		return nil, err
	}
	if !parentStat.IsDir() {
		return nil, fmt.Errorf("filesystem: expected parent %s to be dir", parentPath)
	}
	list.Parent = Dir{
		Name:    parentStat.Name(),
		Path:    parentPath,
		LastMod: parentStat.ModTime(),
	}
	list.Files = append(list.Files, File{
		Name:    stat.Name(),
		Path:    rel,
		LastMod: stat.ModTime(),
		Size:    stat.Size(),
	})
	return list, nil
}