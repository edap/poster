package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func isImage(filename string) bool {
	is_img, _ := regexp.MatchString("(?i)\\.(png|jpg|jpeg|gif)$", filename)
	return is_img
}

// implementare log, o almeno, avere una politica coerente sugli errori
func listFiles(source_dir string) (int, []string) {
	dirname := source_dir
	d, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	tot := 0
	files := []string{}
	for _, fi := range fi {
		if fi.Mode().IsRegular() && isImage(fi.Name()) {
			tot += 1
			files = append(files, filepath.Join(source_dir, fi.Name()))
			//fmt.Println(fi.Name(), fi.Size(), "bytes")
		}
	}
	return tot, files
}

func isWritableByTheUser(fi os.FileInfo, path string) error {
	perm := fi.Mode().String()
	if (strings.IndexAny(perm, "w")) == 2 {
		return nil
	} else {
		return fmt.Errorf("Folder %s not writable by the user with id %d", path, os.Getuid())
	}
}

func createDirectory(path string) error {
	finfo, err := os.Stat(path)
	if err != nil {
		return os.Mkdir(path, 0777)
	}
	if finfo.Mode().IsRegular() {
		return fmt.Errorf("there is already a file called %g ", path)
	}
	if finfo.IsDir() {
		err := isWritableByTheUser(finfo, path)
		if err != nil {
			return err
		}
	}
	return err
}

func randStr(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
