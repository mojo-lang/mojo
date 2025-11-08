package core

const ColorTypeName = "Color"
const ColorTypeFullName = "mojo.core.Color"

func NewColor(r uint32, g uint32, b uint32) *Color {
	return &Color{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

func Rgb(r uint32, g uint32, b uint32) *Color {
	return NewColor(r, g, b)
}

func Rgba(r uint32, g uint32, b uint32, alpha float32) *Color {
	return NewColor(r, g, b).SetAlphaValue(alpha)
}

func (x *Color) SetAlphaValue(value float32) *Color {
	if x != nil {
		x.Alpha = &Float32Value{Val: value}
	}
	return x
}

func (x *Color) GetAlphaValue() float32 {
	if x != nil && x.Alpha != nil {
		return x.Alpha.Val
	}
	return 1
}
