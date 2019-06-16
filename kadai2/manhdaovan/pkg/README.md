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

// Example of converting PNG to Bitmap (bmp),
// that Bitmap is not supported by default.
//
// 1. Define new type that implement Encode interface
type bmpEncoder struct{}

// 2. Implement Encode method of Encoder interface
func (bmpEncoder) Encode(w io.Writer, img image.Image) error {
	return bmp.Encode(w, img)
}

func main() {
	// 3. Init new encoder
	var be bmpEncoder
	var pd imgconv.PNGDecoder
	// 4. Init Converter
	conv := &imgconv.Converter{
		DestImgExt: "bmp",
		Dec:        ipd,
		Enc:        be,
		// Picker is optional. If you not set,
		// imgconv.DefaultImgPicker will be set automatically.
		// Picker: imgconv.DefaultImgPicker{}
		SkipErr:    true, // skip and continue on error while converting
		KeepSrcImg: true, // delete source files after converted
	}
	// Convert all bmp files in directory recursively
	if err := conv.ConvertDir("/path/to/dir", "png"); err != nil {
		panic(err) // dont panic! Handle error yourself.
	}
	// Or convert single image only
	if err := conv.ConvertImg("/path/to/image.png"); err != nil {
		panic(err) // dont panic! Handle error yourself.
	}
}
```