[![GoDoc](https://godoc.org/github.com/edap/poster?status.png)](https://godoc.org/github.com/edap/poster)
[![Build Status](https://drone.io/github.com/edap/poster/status.png)](https://drone.io/github.com/edap/poster/latest) 

# poster
poster is a command line tool that allows you to merge more images in one.
Given a source folder containing the images, the program resizes all the images at the same dimension (default is 120x90) and calculates the disposition of the images in a rectangle. If the number of the images is a prime number, and one of the image can not fit into a rectangle, the tool will skip one images, until the total number of the images will fit into a  rectangle.

## Usage
Simply run `poster` in the folder containing your images. If you run `poster -h` the default options will be displayed

```go
Usage of poster:
  -dest_dir=".": the destination directory that will contain the grid
  -log_file="stdout": specify a log file, as default it will print on stdout
  -source_dir=".": the origin directory that contains the images to compose the grid
  -thumb_height=90: the height of a single thumb
  -thumb_width=120: the width of a single thumb
```

To specify a different source directory as the current direcotry, a different destination directory as the current one, and a logfile, do as follow.

`poster -dest_dir=/home/username/dest -source_dir=/home/username/source -log_file=/home/username/my.log` 

##Installation
Assuming that you have the go toolchain installed, download the package with `go get github.com/edap/poster` and install it moving in the downloaded folder and running `go install`.

## TODO

* ~~provide log option~~ _done_
* Support multiple image formats, .gif, .png, and not only jpeg
* Use goroutine during the canvas creation
* Provide batch resize command, without merge
* Aspect ratio option, 4/3 and 16/9

