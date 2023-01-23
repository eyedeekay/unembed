package unembed

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Unembed unpacks an embed.FS to a directory recursively.
func Unembed(efs embed.FS, dir string, prefix ...string) error {
	concatenatedPrefix := filepath.Join(prefix...)
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
			unpackpath = strings.Replace(unpackpath, concatenatedPrefix, "", 1)
			if err := ioutil.WriteFile(unpackpath, data, mode); err != nil {
				return err
			}
			if mode.Perm() == 0444 {
				log.Println("chmod", unpackpath, 0644)
				if err := os.Chmod(unpackpath, 0644); err != nil {
					return err
				}
			}
			//add write permission to the file

		}
		return nil
	})
}
