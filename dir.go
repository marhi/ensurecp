package ensurecp

import (
	"os"
	"path"
	"io/ioutil"
)

func RCopy(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		srcCopyPath := CopyPath(src)
		_, err := srcCopyPath.CopyTo(dst, true)

		if err != nil {
			return err
		}

		return nil
	} else {
		os.MkdirAll(dst, fi.Mode())
	}

	files, err := ioutil.ReadDir(src)
	if  err != nil {
		return err
	}

	for _, v := range files {
		srcPath := path.Join(src, v.Name())
		dstPath := path.Join(dst, v.Name())
		RCopy(srcPath, dstPath)
	}

	return nil
}
