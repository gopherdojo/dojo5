package imgconv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConverter_ConvertDir(t *testing.T) {
	testDirs := paths{rootForTestDir, subDir}
	noAffectedFiles := paths{
		rootForTestDir + "/textFile.txt",
		rootForTestDir + "/textFileRenameToPNG.png",

		subDir + "/textFile.txt",
		subDir + "/textFileRenameToPNG.png",
	}

	type fields struct {
		DestImgExt ImgExt
		SkipErr    bool
		KeepSrcImg bool
		Dec        Decoder
		Enc        Encoder
	}

	type args struct {
		srcImgType ImgType
		dirPath    string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		// ensure correct image type after converted
		expectImgType ImgType
		// ensure target images were deleted after converted
		wantDeletedImgs []string
		// ensure imgs were convert correctly
		// and no file was deleted after converted
		wantConvertedImgs []string
		wantErr           bool
	}{
		// convert without error, no keep org img
		{
			name: "jpg to png",
			fields: fields{
				DestImgExt: "png",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        JPEGDecoder{},
				Enc:        PNGEncoder{},
			},
			args: args{
				srcImgType: "jpeg",
				dirPath:    rootForTestDir,
			},
			expectImgType: "png",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.jpg",
				subDir + "/validJPEG.jpg",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.png",
				subDir + "/validJPEG.png",
			},
			wantErr: false,
		},
		{
			name: "jpg to gif",
			fields: fields{
				DestImgExt: "gif",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        JPEGDecoder{},
				Enc:        GIFEncoder{},
			},
			args: args{
				srcImgType: "jpeg",
				dirPath:    rootForTestDir,
			},
			expectImgType: "gif",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.jpg",
				subDir + "/validJPEG.jpg",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validJPEG.gif",
				subDir + "/validJPEG.gif",
			},
			wantErr: false,
		},
		{
			name: "gif to jpg",
			fields: fields{
				DestImgExt: "jpg",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        GIFDecoder{},
				Enc:        JPEGEncoder{},
			},
			args: args{
				srcImgType: "gif",
				dirPath:    rootForTestDir,
			},
			expectImgType: "jpeg",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.gif",
				subDir + "/validGIF.gif",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.jpg",
				subDir + "/validGIF.jpg",
			},
			wantErr: false,
		},
		{
			name: "gif to png",
			fields: fields{
				DestImgExt: "png",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        GIFDecoder{},
				Enc:        PNGEncoder{},
			},
			args: args{
				srcImgType: "gif",
				dirPath:    rootForTestDir,
			},
			expectImgType: "png",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.gif",
				subDir + "/validGIF.gif",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validGIF.png",
				subDir + "/validGIF.png",
			},
			wantErr: false,
		},
		{
			name: "png to gif",
			fields: fields{
				DestImgExt: "gif",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        PNGDecoder{},
				Enc:        GIFEncoder{},
			},
			args: args{
				srcImgType: "png",
				dirPath:    rootForTestDir,
			},
			expectImgType: "gif",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.png",
				subDir + "/validPNG.png",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.gif",
				subDir + "/validPNG.gif",
			},
			wantErr: false,
		},
		{
			name: "png to jpg",
			fields: fields{
				DestImgExt: "jpg",
				SkipErr:    false,
				KeepSrcImg: false,
				Dec:        PNGDecoder{},
				Enc:        JPEGEncoder{},
			},
			args: args{
				srcImgType: "png",
				dirPath:    rootForTestDir,
			},
			expectImgType: "jpeg",
			wantDeletedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.gif",
				subDir + "/validPNG.gif",
			},
			wantConvertedImgs: paths{
				// converted image
				rootForTestDir + "/validPNG.jpg",
				subDir + "/validPNG.jpg",
			},
			wantErr: false,
		},
		// convert without error, keep org img
		{
			name: "png to jpg",
			fields: fields{
				DestImgExt: "jpg",
				SkipErr:    false,
				KeepSrcImg: true,
				Dec:        PNGDecoder{},
				Enc:        JPEGEncoder{},
			},
			args: args{
				srcImgType: "png",
				dirPath:    rootForTestDir,
			},
			expectImgType: "jpeg",
			// no img deleted after converted
			wantDeletedImgs: paths{},
			wantConvertedImgs: paths{
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
				SkipErr: false,
				Dec:     GIFDecoder{}, // wrong decoder
				Enc:     JPEGEncoder{},
			},
			// no new files were added
			wantDeletedImgs: paths{},
			// no new imgs were converted
			wantConvertedImgs: paths{},
			wantErr:           true,
		},
		// in case of wrong encoder,
		// the image still be converted successfully.
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

			conv := Converter{
				DestImgExt: tt.fields.DestImgExt,
				SkipErr:    tt.fields.SkipErr,
				KeepSrcImg: tt.fields.KeepSrcImg,
				Enc:        tt.fields.Enc,
				Dec:        tt.fields.Dec,
			}
			conv.errOnConvImg = errBuilder(conv.SkipErr)

			// ensure about error
			if err := conv.ConvertDir(tt.args.dirPath, tt.args.srcImgType); (err != nil) != tt.wantErr {
				t.Errorf("Converter.convertDir() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ensure about files not be converted are unaffected
			if err := verifyFiles(noAffectedFiles, true); err != nil {
				t.Errorf("file not existed after test: %+v", err)
			}

			// ensure about files were deleted after converted
			if err := verifyFiles(tt.wantDeletedImgs, false); err != nil {
				t.Errorf("file not deleted after test: %+v", err)
			}

			// ensure new images were added after converted
			if err := verifyImgs(tt.wantConvertedImgs, tt.expectImgType); err != nil {
				t.Errorf("img not correct after test, err: %+v", err)
			}
		})
	}
}

func TestConverter_ConvertImg(t *testing.T) {
	type fields struct {
		DestImgExt ImgExt
		Dec        Decoder
		Enc        Encoder
		Picker     ImgPicker
		SkipErr    bool
		KeepSrcImg bool
	}
	type args struct {
		imgPath string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantConvertedImgs paths
		wantDeletedImgs   paths
		expectImgType     ImgType
		wantErr           bool
	}{
		{
			name: "convert ok, without keep org file",
			fields: fields{
				DestImgExt: "jpg",
				Dec:        PNGDecoder{},
				Enc:        JPEGEncoder{},
				Picker:     DefaultImgPicker{},
				SkipErr:    false,
				KeepSrcImg: false,
			},
			args: args{imgPath: rootForTestDir + "/validPNG.png"},
			wantConvertedImgs: paths{
				rootForTestDir + "/validPNG.jpg",
			},
			wantDeletedImgs: paths{
				rootForTestDir + "/validPNG.png",
			},
			expectImgType: "jpeg",
			wantErr:       false,
		},
		{
			name: "convert ok, with keep org file",
			fields: fields{
				DestImgExt: "png",
				Dec:        GIFDecoder{},
				Enc:        PNGEncoder{},
				Picker:     DefaultImgPicker{},
				SkipErr:    false,
				KeepSrcImg: true,
			},
			args: args{imgPath: rootForTestDir + "/validGIF.gif"},
			wantConvertedImgs: paths{
				rootForTestDir + "/validGIF.png",
			},
			wantDeletedImgs: paths{},
			expectImgType:   "png",
			wantErr:         false,
		},
		{
			name: "convert error, no skip error",
			fields: fields{
				DestImgExt: "png",
				Dec:        GIFDecoder{},
				Enc:        PNGEncoder{},
				Picker:     DefaultImgPicker{},
				SkipErr:    false,
				KeepSrcImg: false,
			},
			args:              args{imgPath: rootForTestDir + "/noExistingImg.gif"},
			wantConvertedImgs: paths{},
			wantDeletedImgs:   paths{},
			expectImgType:     "png",
			wantErr:           true,
		},
		{
			name: "convert error on verify file, skip error",
			fields: fields{
				DestImgExt: "png",
				Dec:        GIFDecoder{},
				Enc:        PNGEncoder{},
				Picker:     DefaultImgPicker{},
				SkipErr:    true,
				KeepSrcImg: false,
			},
			args:              args{imgPath: "/no/existing/img.gif"},
			wantConvertedImgs: paths{},
			wantDeletedImgs:   paths{},
			expectImgType:     "png",
			wantErr:           true,
		},
		{
			name: "convert error on decoding file, skip error",
			fields: fields{
				DestImgExt: "png",
				Dec:        errDecoder{},
				Enc:        PNGEncoder{},
				Picker:     DefaultImgPicker{},
				SkipErr:    true,
				KeepSrcImg: false,
			},
			args:              args{imgPath: rootForTestDir + "/validPNG.png"},
			wantConvertedImgs: paths{},
			wantDeletedImgs:   paths{},
			expectImgType:     "png",
			wantErr:           false,
		},
	}

	testDirs := paths{rootForTestDir}
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

			conv := &Converter{
				DestImgExt: tt.fields.DestImgExt,
				Dec:        tt.fields.Dec,
				Enc:        tt.fields.Enc,
				Picker:     tt.fields.Picker,
				SkipErr:    tt.fields.SkipErr,
				KeepSrcImg: tt.fields.KeepSrcImg,
			}
			conv.errOnConvImg = errBuilder(conv.SkipErr)

			if err := conv.ConvertImg(tt.args.imgPath); (err != nil) != tt.wantErr {
				t.Errorf("Converter.ConvertImg() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ensure about files were deleted after converted
			if err := verifyFiles(tt.wantDeletedImgs, false); err != nil {
				t.Errorf("file not deleted after test: %+v", err)
			}
			// ensure new images were added after converted
			if err := verifyImgs(tt.wantConvertedImgs, tt.expectImgType); err != nil {
				t.Errorf("img not correct after test, err: %+v", err)
			}
		})
	}
}

func TestConverter_validate(t *testing.T) {
	type fields struct {
		DestImgExt   ImgExt
		Dec          Decoder
		Enc          Encoder
		Picker       ImgPicker
		SkipErr      bool
		KeepSrcImg   bool
		errOnConvImg func(err error) error
		path         string
		srcImgType   ImgType
	}
	type args struct {
		isConvertDir bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:    "using ConvertDir for file",
			fields:  fields{path: rootForTestDir},
			args:    args{false},
			wantErr: fmt.Errorf("%s is dir, use ConvertDir method instead", rootForTestDir),
		},
		{
			name:    "using ConvertImg for dir",
			fields:  fields{path: orgFilesDir + "/validPNG.png"},
			args:    args{true},
			wantErr: fmt.Errorf("%s is file, use ConvertImg method instead", orgFilesDir+"/validPNG.png"),
		},
		{
			name:    "decoder is not set",
			fields:  fields{path: orgFilesDir + "/validPNG.png"},
			args:    args{false},
			wantErr: fmt.Errorf("decoder is not set"),
		},
		{
			name: "encoder is not set",
			fields: fields{
				Dec:  PNGDecoder{},
				path: orgFilesDir + "/validPNG.png",
			},
			args:    args{false},
			wantErr: fmt.Errorf("encoder is not set"),
		},
		{
			name: "picker is not set",
			fields: fields{
				Dec:  PNGDecoder{},
				Enc:  JPEGEncoder{},
				path: orgFilesDir + "/validPNG.png",
			},
			args:    args{false},
			wantErr: fmt.Errorf("image file picker is not set"),
		},
		{
			name: "DestImgExt is not set",
			fields: fields{
				Dec:    PNGDecoder{},
				Enc:    JPEGEncoder{},
				Picker: DefaultImgPicker{},
				path:   orgFilesDir + "/validPNG.png",
			},
			args:    args{false},
			wantErr: fmt.Errorf("destination extension is not set"),
		},
		{
			name: "valid converter",
			fields: fields{
				Dec:        PNGDecoder{},
				Enc:        JPEGEncoder{},
				Picker:     DefaultImgPicker{},
				DestImgExt: "jpg",
				path:       orgFilesDir + "/validPNG.png",
			},
			args:    args{false},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conv := &Converter{
				DestImgExt:   tt.fields.DestImgExt,
				Dec:          tt.fields.Dec,
				Enc:          tt.fields.Enc,
				Picker:       tt.fields.Picker,
				SkipErr:      tt.fields.SkipErr,
				KeepSrcImg:   tt.fields.KeepSrcImg,
				errOnConvImg: tt.fields.errOnConvImg,
				path:         tt.fields.path,
				srcImgType:   tt.fields.srcImgType,
			}
			if err := conv.validate(tt.args.isConvertDir); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Converter.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
