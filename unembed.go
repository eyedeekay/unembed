package unembed

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Unembed unpacks an embed.FS to a directory recursively.
func Unembed(efs embed.FS, dir string) error {
	return fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			err := os.MkdirAll(dir+"/"+path, 0755)
			if err != nil {
				return err
			}
		} else {
			file, err := efs.Open(path)
			if err != nil {
				return err
			}
			info, err := file.Stat()
			if err != nil {
				return err
			}
			mode := info.Mode()
			data, err := efs.ReadFile(path)
			if err != nil {
				return err
			}
			unpackpath := filepath.Join(dir, path)
			if err := ioutil.WriteFile(unpackpath, data, mode); err != nil {
				return err
			}
		}
		return nil
	})
}
