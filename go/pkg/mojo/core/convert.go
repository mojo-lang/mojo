package core

type FromInt32Converter interface {
	FromInt32(value int32) error
}

type ToInt32Converter interface {
	ToInt32() int32
}

type FromUInt32Converter interface {
	FromUInt32(value uint32) error
}

type ToUInt32Converter interface {
	ToUInt32() uint32
}

type FromInt64Converter interface {
	FromInt64(value int64) error
}

type ToInt64Converter interface {
	ToInt64() int64
}

type FromUInt64Converter interface {
	FromUInt64(value uint64) error
}

type ToUInt64Converter interface {
	ToUInt64() uint64
}

type FromFloat32Converter interface {
	FromFloat32(value float32) error
}

type ToFloat32Converter interface {
	ToFloat32() float32
}

type FromFloat64Converter interface {
	FromFloat64(value float64) error
}

type ToFloat64Converter interface {
	ToFloat64() float64
}

type FromStringConverter interface {
	FromString(value string) error
}

type ToStringConverter interface {
	ToString() string
}
