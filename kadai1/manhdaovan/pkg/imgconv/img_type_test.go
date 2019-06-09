package imgconv

import "testing"

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
