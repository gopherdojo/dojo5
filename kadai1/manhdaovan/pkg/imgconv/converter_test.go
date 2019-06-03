package imgconv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSupportSrcImgTypes_IsSupport(t *testing.T) {
	allSupportImgTypes := GetSupportSrcImgTypes()

	type args struct {
		imgType string
	}
	tests := []struct {
		name  string
		ssits SupportSrcImgTypes
		args  args
		want  bool
	}{
		{
			"no supported img type",
			allSupportImgTypes,
			args{"aaa"},
			false,
		},
		{
			"no supported img type",
			allSupportImgTypes,
			args{"png"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ssits.IsSupport(ImgType(tt.args.imgType)); got != tt.want {
				t.Errorf("SupportSrcImgTypes.IsSupport() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testdataDir = "./testdata"
	// orgFilesDir contains:
	// - textFile.txt
	// - textFileRenameToPNG.png
	// - validGIF.gif
	// - validJPEG.jpg
	// - validPNG.png
	orgFilesDir = testdataDir + "/orgFiles"
	orgFiles    = paths{
		orgFilesDir + "/textFile.txt",
		orgFilesDir + "/textFileRenameToPNG.png",
		orgFilesDir + "/validGIF.gif",
		orgFilesDir + "/validJPEG.jpg",
		orgFilesDir + "/validPNG.png",
	}
	rootForTestDir = testdataDir + "/rootForTest"
	subDir         = rootForTestDir + "/subdir"
)

func Test_converter_convert(t *testing.T) {
	testDirs := paths{rootForTestDir, subDir}
	noAffectedFiles := paths{
		rootForTestDir + "/textFile.txt",
		rootForTestDir + "/textFileRenameToPNG.png",

		subDir + "/textFile.txt",
		subDir + "/textFileRenameToPNG.png",
	}

	type fields struct {
		dirPath     string
		srcImgType  ImgType
		destImgType ImgType
		destImgExt  ImgExt
		skipErr     bool
		decoder     Decoder
		encoder     Encoder
	}
	tests := []struct {
		name   string
		fields fields
		// ensure target images were deleted after converted
		expectDeletedImgs []string
		// ensure imgs were convert correctly
		// and no file was deleted after converted
		expectConvertedImgs []string
		wantErr             bool
	}{
		// convert without error
		{
			name: "jpg to png",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypeJPEG,
				destImgType: ImgTypePNG,
				destImgExt:  "png",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypeJPEG),
				encoder:     GetEncoder(ImgTypePNG),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.jpg",
				subDir + "/validJPEG.jpg",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.png",
				subDir + "/validJPEG.png",
			},
			wantErr: false,
		},
		{
			name: "jpg to gif",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypeJPEG,
				destImgType: ImgTypeGIF,
				destImgExt:  "gif",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypeJPEG),
				encoder:     GetEncoder(ImgTypeGIF),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.jpg",
				subDir + "/validJPEG.jpg",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.gif",
				subDir + "/validJPEG.gif",
			},
			wantErr: false,
		},
		{
			name: "gif to jpg",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypeGIF,
				destImgType: ImgTypeJPEG,
				destImgExt:  "jpg",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypeGIF),
				encoder:     GetEncoder(ImgTypeJPEG),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.gif",
				subDir + "/validGIF.gif",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.jpg",
				subDir + "/validGIF.jpg",
			},
			wantErr: false,
		},
		{
			name: "gif to png",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypeGIF,
				destImgType: ImgTypePNG,
				destImgExt:  "png",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypeGIF),
				encoder:     GetEncoder(ImgTypePNG),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.gif",
				subDir + "/validGIF.gif",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.png",
				subDir + "/validGIF.png",
			},
			wantErr: false,
		},
		{
			name: "png to gif",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypePNG,
				destImgType: ImgTypeGIF,
				destImgExt:  "gif",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypePNG),
				encoder:     GetEncoder(ImgTypeGIF),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.png",
				subDir + "/validPNG.png",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.gif",
				subDir + "/validPNG.gif",
			},
			wantErr: false,
		},
		{
			name: "png to jpg",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypePNG,
				destImgType: ImgTypeJPEG,
				destImgExt:  "jpg",
				skipErr:     false,
				decoder:     GetDecoder(ImgTypePNG),
				encoder:     GetEncoder(ImgTypeJPEG),
			},
			expectDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.gif",
				subDir + "/validPNG.gif",
			},
			expectConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.jpg",
				subDir + "/validPNG.jpg",
			},
			wantErr: false,
		},

		// convert with error
		{
			name: "wrong decoder",
			fields: fields{
				dirPath:     rootForTestDir,
				srcImgType:  ImgTypePNG,
				destImgType: ImgTypeJPEG,
				skipErr:     false,
				decoder:     GetDecoder(ImgTypeGIF), // wrong decoder
				encoder:     GetEncoder(ImgTypeJPEG),
			},
			// no new files were added
			expectDeletedImgs: paths{},
			// no new imgs were converted
			expectConvertedImgs: paths{},
			wantErr:             true,
		},
		// in case of wrong encoder,
		// the image still be converted successfully,
		// so we need to ensure no wrong encoder is taken
		// by test cases of getEncoder method
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyTestFilesToDir(orgFiles, testDirs); err != nil {
				t.Errorf("error on copy test files to: %v, err: %+v", testDirs, err)
				return
			}

			defer func() {
				// clean all files in test dirs after each test
				if err := deleteAllFiles(testDirs); err != nil {
					t.Errorf("error on cleaning files in: %v, err: %+v", testDirs, err)
				}
			}()

			conv := converter{
				dirPath:     tt.fields.dirPath,
				srcImgType:  tt.fields.srcImgType,
				destImgType: tt.fields.destImgType,
				destImgExt:  tt.fields.destImgExt,
				skipErr:     tt.fields.skipErr,
				encoder:     tt.fields.encoder,
				decoder:     tt.fields.decoder,
			}
			conv.errOnConvImg = errBuilder(conv.skipErr)

			// ensure about error
			if err := conv.convert(); (err != nil) != tt.wantErr {
				t.Errorf("converter.convert() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ensure about files not be converted are unaffected
			if err := verifyFiles(noAffectedFiles, true); err != nil {
				t.Errorf("file not existed after test: %+v", err)
			}

			// ensure about files were deleted after converted
			if err := verifyFiles(tt.expectDeletedImgs, false); err != nil {
				t.Errorf("file not deleted after test: %+v", err)
			}

			// ensure new images were added after converted
			if err := verifyImgs(tt.expectConvertedImgs, tt.fields.destImgType); err != nil {
				t.Errorf("img not correct after test, err: %+v", err)
			}
		})
	}
}

func Test_isImgWithType(t *testing.T) {
	pngFile := rootForTestDir + "/validPNG.png"
	txtFile := rootForTestDir + "/textFile.txt"
	txtFileRenamedToPNGFile := rootForTestDir + "/textFileRenameToPNG.png"

	type args struct {
		path    string
		imgType ImgType
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"correct img and type",
			args{pngFile, ImgTypePNG},
			true,
			false,
		},
		{
			"correct img and wrong type",
			args{pngFile, ImgTypeJPEG},
			false,
			false,
		},
		{
			"wrong img",
			args{txtFileRenamedToPNGFile, ImgTypeJPEG},
			false,
			false,
		},
		{
			"wrong file",
			args{txtFile, ImgTypeJPEG},
			false,
			false,
		},
		{
			"img not exist",
			args{"not_exit_img.jpg", ImgTypeJPEG},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyTestFilesToDir(orgFiles, paths{rootForTestDir}); err != nil {
				t.Errorf("error on copy test files to test root dirs: %+v", err)
				return
			}

			defer func() {
				// clean all files in test dirs after each test
				if err := deleteAllFiles(paths{rootForTestDir}); err != nil {
					t.Errorf("error on cleaning files in test root dir: %+v", err)
				}
			}()

			got, err := isImgWithType(tt.args.path, tt.args.imgType)
			if (err != nil) != tt.wantErr {
				t.Errorf("isImgWithType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isImgWithType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_converter_convertImg(t *testing.T) {
	tests := []struct {
		name    string
		skipErr bool
		img     image
		wantErr bool
	}{
		{
			"skip error is true",
			true,
			image{"/no/exist/file.png"},
			false,
		},
		{
			"skip error is false",
			false,
			image{"/no/exist/file.png"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := converter{skipErr: tt.skipErr}
			conv.errOnConvImg = errBuilder(conv.skipErr)
			errReturn := conv.errOnConvImg(conv.convertImg(tt.img))

			if (errReturn != nil) != tt.wantErr {
				t.Errorf("converter.convertImg() error = %v, wantErr %v", errReturn, tt.wantErr)
			}
		})
	}
}

type paths []string

// verifyImgs verifies that images are existing or not,
// and corresponding to it type
func verifyImgs(imgs paths, imgType ImgType) error {
	for _, img := range imgs {
		ok, err := isImgWithType(img, imgType)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("converted image type is wrong, img: %s, expect type: %s", img, imgType)
		}
	}

	return nil
}

// verifyFiles verifies files are existing or not
func verifyFiles(files paths, checkExisting bool) error {
	for _, f := range files {
		file, err := os.Open(f)
		defer file.Close()
		switch checkExisting {
		case true:
			if err == nil {
				continue
			}
			return fmt.Errorf("need file %v existing, got err: %+v", f, err)
		case false:
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("need file %v removed, got err: %v", f, err)
		}
	}

	return nil
}

func copyTestFilesToDir(files paths, dirs paths) error {
	for _, dir := range dirs {
		for _, f := range files {
			srcFile, err := ioutil.ReadFile(f)
			if err != nil {
				return err
			}

			fileNameIdx := strings.LastIndex(f, "/")
			destPath := dir + string(f[fileNameIdx:])
			err = ioutil.WriteFile(destPath, srcFile, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func deleteAllFiles(dirs paths) error {
	for _, dir := range dirs {
		outerErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() { // file
				return os.Remove(path)
			}

			return nil
		})

		if outerErr != nil {
			return outerErr
		}
	}

	return nil
}
