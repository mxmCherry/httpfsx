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
	Parent  Dir
	Dirs    []Dir
	Files   []File
	LastMod time.Time
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

	stats, err := os.Lstat(abs)
	if err != nil {
		return nil, err
	}

	list := &List{
		Dirs:    []Dir{},
		Files:   []File{},
		LastMod: stats.ModTime(),
	}

	if stats.IsDir() {
		list.Parent = Dir{
			Name:    stats.Name(),
			Path:    rel,
			LastMod: stats.ModTime(),
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

			lm := fi.ModTime()

			if fi.IsDir() {
				list.Dirs = append(list.Dirs, Dir{
					Name:    name,
					Path:    path.Join(rel, name),
					LastMod: lm,
				})
			} else if fi.Mode().IsRegular() {
				mime, err := mime.Detect(filepath.Join(abs, name))
				if err != nil {
					return nil, err
				}
				list.Files = append(list.Files, File{
					Name:    name,
					Path:    path.Join(rel, name),
					LastMod: lm,
					Size:    fi.Size(),
					Mime:    mime,
				})
			} else {
				continue
			}

			if lm.After(list.LastMod) {
				list.LastMod = lm
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
		Name:    stats.Name(),
		Path:    rel,
		LastMod: stats.ModTime(),
		Size:    stats.Size(),
		Mime:    mime,
	})
	return list, nil
}
