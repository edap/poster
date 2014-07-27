// Copyright (c) 2014 Davide Prati

// MIT License

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

/*

poster is a command line tool that allows you to compose one images using more images located in a folder.

Given a source folder containing the images, the program resizes all the images
at the same dimension (default is 120x90) and calculates the disposition of the images in a rectangle.
If the number of the images is a prime number, and one of the image can not fit into a rectangle, the tool will skip one images, until the total number of the images will fit into a  rectangle.

Simply run 'poster' in the folder containing your images. If you run 'poster -h' the default options will be displayed

	Usage of poster:
	  -dest_dir=".": the destination directory that will contain the grid
	  -log_file="stdout": specify a log file, as default it will print on stdout
	  -source_dir=".": the origin directory that contains the images to compose the grid
	  -thumb_height=90: the height of a single thumb
	  -thumb_width=120: the width of a single thumb

To specify a different source directory as the current direcotry, a different destination directory as the current one, and a logfile, do as follow.

	poster dest_dir=/home/username/dest source_dir=/home/username/source log_file=/home/username/my.log

*/
package main
