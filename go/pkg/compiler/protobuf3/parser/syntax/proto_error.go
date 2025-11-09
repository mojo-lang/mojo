package syntax

type ProtoError struct {
}

func (e ProtoError) Error() string {
	return "the proto version not matched"
}
