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

// isImage check if a file is in jpg, png or gif format reading the extension of the filename
func isImage(filename string) bool {
	is_img, _ := regexp.MatchString("(?i)\\.(png|jpg|jpeg|gif)$", filename)
	return is_img
}

// isJpeg check if an image is in jpg format or not reading the extension of the filename
func isJpeg(filename string) bool {
	is_jpeg, _ := regexp.MatchString("(?i)\\.(jpg|jpeg)$", filename)
	return is_jpeg
}

// listFiles read the jpgs contained in a folder, it returns the total number of the images
// and an array containing the paths
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
		if fi.Mode().IsRegular() && isJpeg(fi.Name()) {
			// image conversion still to implement
			//if fi.Mode().IsRegular() && isImage(fi.Name()) {
			tot += 1
			files = append(files, filepath.Join(source_dir, fi.Name()))
			//fmt.Println(fi.Name(), fi.Size(), "bytes")
		}
	}
	return tot, files
}

// isWritableByTheUser take a os.FileInfo and a path as parameter. Check the write permission
// for the given path for the current user
func isWritableByTheUser(fi os.FileInfo, path string) error {
	perm := fi.Mode().String()
	if (strings.IndexAny(perm, "w")) == 2 {
		return nil
	} else {
		return fmt.Errorf("Folder %s not writable by the user with id %d", path, os.Getuid())
	}
}

// createDirectory try to create a directory in the given path. Return an error if it fails
func createDirectory(path string) error {
	finfo, err := os.Stat(path)
	if err != nil {
		return os.Mkdir(path, 0777)
	}
	if finfo.Mode().IsRegular() {
		return fmt.Errorf("there is already a file called %s ", path)
	}
	if finfo.IsDir() {
		err := isWritableByTheUser(finfo, path)
		if err != nil {
			return err
		}
	}
	return err
}

// randStr generate an alfanuperich random string
func randStr(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
