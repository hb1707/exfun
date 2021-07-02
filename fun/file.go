package fun

import (
	"archive/zip"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func FileMD5(path string) (string, error) {
	var err error
	fileTemp, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer fileTemp.Close()
	md5h := md5.New()
	if _, err := io.Copy(md5h, fileTemp); err != nil {
		fmt.Println("Copy", err)
		return "", err
	}
	str := fmt.Sprintf("%x", md5h.Sum(nil))
	return str, err
}
func Read(path string) string {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("ioutil open err : %v\n", err)
	}
	return string(byt)
}
func Write(path, str string) {
	err := ioutil.WriteFile(path, []byte(str), os.ModePerm)
	if err != nil {
		fmt.Printf("ioutil write err : %v\n", err)
	}
}
func Download(url, local string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dir, _ := path.Split(local)
	if err = os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	out, err := os.Create(local)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func Unzip(zipFile, dest string, exclude []string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	toDist := func(f *zip.File) error {
		//log.Println(f.Name)
		if InSliceString(f.Name, exclude) {
			return nil
		}
		fPath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(fPath, 0755); err != nil {
				return err
			}
			return nil
		}
		if err = os.MkdirAll(filepath.Dir(fPath), 0755); err != nil {
			return err
		}
		w, err := os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			return err
		}
		defer w.Close()
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		_, err = io.Copy(w, rc)
		return err
	}
	for _, f := range reader.File {
		err := toDist(f)
		if err != nil {
			return err
		}
	}
	return nil
}

//判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
