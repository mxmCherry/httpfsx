package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/mxmCherry/httpfsx/internal/mime"
)

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
	Mime    string
}

func Abs(root, rel string) string {
	rel = path.Join("/", rel)
	return filepath.Join(root, rel)
}

func Ls(root, rel string) (*List, error) {
	rel = path.Join("/", rel)
	abs := filepath.Join(root, rel)

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
				mime, err := mime.Detect(filepath.Join(abs, name))
				if err != nil {
					return nil, err
				}
				list.Files = append(list.Files, File{
					Name:    name,
					Path:    path.Join(rel, name),
					LastMod: fi.ModTime(),
					Size:    fi.Size(),
					Mime:    mime,
				})
			}
		}

		return list, nil
	}

	parentPath := path.Dir(rel)
	parentStat, err := os.Stat(filepath.Join(root, parentPath))
	if err != nil {
		return nil, err
	}
	if !parentStat.IsDir() {
		return nil, fmt.Errorf("filesystem: expected parent %s to be dir", parentPath)
	}

	mime, err := mime.Detect(abs)
	if err != nil {
		return nil, err
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
		Mime:    mime,
	})
	return list, nil
}
