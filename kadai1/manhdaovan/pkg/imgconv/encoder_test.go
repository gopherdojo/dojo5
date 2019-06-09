package imgconv

import (
	"reflect"
	"testing"
)

func Test_getEncoder(t *testing.T) {
	type args struct {
		imgType ImgType
	}
	tests := []struct {
		name string
		args args
		want Encoder
	}{
		{
			"get jpg decoder",
			args{ImgTypeJPEG},
			jpegEncoder{},
		},
		{
			"get png decoder",
			args{ImgTypePNG},
			pngEncoder{},
		},
		{
			"get gif decoder",
			args{ImgTypeGIF},
			gifEncoder{},
		},
		{
			"get wrong decoder",
			args{ImgType("notsupport")},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEncoder(tt.args.imgType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
