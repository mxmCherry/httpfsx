package httpfsx

import (
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
)

func FileServer(fs http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := path.Join("/", r.URL.Path)

		f, err := fs.Open(name)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !stat.IsDir() {
			http.ServeContent(w, r, name, stat.ModTime(), f)
			return
		}

		fis, err := f.Readdir(0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sort.Slice(fis, func(i, j int) bool {
			return fis[i].Name() < fis[j].Name()
		})

		data := struct {
			Base  string
			Self  os.FileInfo
			Dirs  []os.FileInfo
			Files []os.FileInfo
		}{
			Base: name,
			Self: stat,
		}
		for _, fi := range fis {
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			if fi.IsDir() {
				data.Dirs = append(data.Dirs, fi)
			} else if fi.Mode().IsRegular() {
				data.Files = append(data.Files, fi)
			}
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err.Error())
		}
	})
}
