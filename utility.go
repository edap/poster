package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func isImage(filename string) bool {
	is_img, _ := regexp.MatchString(".png$|.jpg$|.jpeg$|.gif$", filename)
	return is_img
}

// implementare log, o almeno, avere una politica coerente sugli errori
func listFiles(source_dir string) (int, []string) {
	dirname := source_dir
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

func randStr(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
