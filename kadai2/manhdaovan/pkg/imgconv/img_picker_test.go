package imgconv

import (
	"testing"
)

func TestDefaultImgPicker_Pick(t *testing.T) {
	type args struct {
		path    string
		imgType ImgType
	}
	tests := []struct {
		name         string
		dp           DefaultImgPicker
		args         args
		wantImgPaths []string
		wantErr      bool
	}{
		{
			name: "pick imgs in dir",
			dp:   DefaultImgPicker{},
			args: args{
				path:    rootForTestDir,
				imgType: "png",
			},
			wantImgPaths: []string{
				rootForTestDir + "/validPNG.png",
				subDir + "/validPNG.png",
			},
			wantErr: false,
		},
		{
			name: "pick single img",
			dp:   DefaultImgPicker{},
			args: args{
				path:    rootForTestDir + "/validPNG.png",
				imgType: "png",
			},
			wantImgPaths: []string{
				rootForTestDir + "/validPNG.png",
			},
			wantErr: false,
		},
		{
			name: "no existing path",
			dp:   DefaultImgPicker{},
			args: args{
				path:    "no/existing/path",
				imgType: "png",
			},
			wantImgPaths: []string{},
			wantErr:      true,
		},
	}
	testDirs := paths{rootForTestDir, subDir}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyTestFilesToDir(t, orgFiles, testDirs)
			defer func() {
				// clean all files in test dirs after each test
				deleteAllFiles(t, testDirs)
			}()

			dp := DefaultImgPicker{}
			gotImgPaths, err := dp.Pick(tt.args.path, tt.args.imgType)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultImgPicker.Pick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isSameFiles(gotImgPaths, tt.wantImgPaths) {
				t.Errorf("DefaultImgPicker.Pick() = %v, want %v", gotImgPaths, tt.wantImgPaths)
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
			copyTestFilesToDir(t, orgFiles, paths{rootForTestDir})
			defer func() {
				// clean all files in test dirs after each test
				deleteAllFiles(t, paths{rootForTestDir})
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
