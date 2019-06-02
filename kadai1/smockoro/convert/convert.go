package convert

type Converter interface {
	ImageConvert(string) error
}
