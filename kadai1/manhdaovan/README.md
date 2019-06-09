## About this tool

```
imgconv -- convert image to other image type
usage: $./bin/imgconv [-s] [-d] [-k] dir
options:
  -s string
	Source image type. jpeg, png and gif are supported
  -d string
  	Destination image type. jpeg, png and gif are supported
  -k
	Skip error and continue while converting
```

## Build & test
```bash
$cd /path/this/repo
# for build
$make build
# for test 
$make test
```

## Extend and usage the imgconv package source code

JPEG, PNG and GIF are supported by default.
But sometime you want to convert another image type.

For example, you want to convert from PNG to Bitmap image,
so you can do this by few lines of code as bellow:

(And you can do the similar steps to new decoder)

```go
package main

import (
	"image"
	"io"

	"github.com/gopherdojo/dojo5/kadai1/manhdaovan/pkg/imgconv"
	"golang.org/x/image/bmp"
)

// Define new type that implement Encode interface
type bmpEncoder struct{}

func (be bmpEncoder) Encode(w io.Writer, img image.Image) error {
	return bmp.Encode(w, img)
}

func main() {
	// init new encoder
	be := bmpEncoder{}
	// define new dest type and extension
	destImgType := imgconv.ImgType("bmp")
	destImgExt := imgconv.ImgExt("bmp")

	// Register new dest img type
	if err := imgconv.RegisterDestImgType(destImgType, be, destImgExt); err != nil {
		panic(err)
	}

	// And call ConvertDir with new dest img type
	if err := imgconv.ConvertDir("/path/to/dir", imgconv.ImgTypePNG, destImgType, true); err != nil {
		panic(err)
	}
}
```