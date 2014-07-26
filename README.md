# gomerger
Gomerger is a command line tool that allows you to merge more images in one.
Given a source folder containing the images, the program resizes all the images at the same dimension (default is 120x90) and calculates the disposition of the images in a rectangle. If the number of the images is a prime number, and one of the image can not fit into a rectangle, the tool will skip it, until the total number of the images will fit into the rectangle.

## Usage
Simply runs gomerger in the folder containing your images.



## TODO

* ~~allow only jpg~~ _done_
* Support multiple image formats, .gif, .png, and not only jpeg
* Use goroutine during the canvas creation
* Provide batch resize command, without merge
* Aspect ratio option, 4/3 and 16/9

