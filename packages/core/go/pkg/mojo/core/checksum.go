package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

const ChecksumTypeName = "Checksum"
const ChecksumTypeFullName = "mojo.core.Checksum"

var checksumLength = map[Checksum_Algorithm]int{
	Checksum_ALGORITHM_MD5:    32,
	Checksum_ALGORITHM_SHA1:   40,
	Checksum_ALGORITHM_SHA256: 64,
	Checksum_ALGORITHM_SHA512: 128,
}

func NewChecksum(algorithm Checksum_Algorithm, data []byte) *Checksum {
	return (&Checksum{Algorithm: algorithm}).Sum(data)
}

func (x *Checksum) Check(data []byte) bool {
	if x != nil && x.Algorithm > 0 && len(x.Val) > 0 {
		switch x.Algorithm {
		case Checksum_ALGORITHM_MD5:
			return x.Val == fmt.Sprintf("%x", md5.Sum(data))
		case Checksum_ALGORITHM_SHA1:
			return x.Val == fmt.Sprintf("%x", sha1.Sum(data))
		case Checksum_ALGORITHM_SHA256:
			return x.Val == fmt.Sprintf("%x", sha256.Sum256(data))
		case Checksum_ALGORITHM_SHA512:
			return x.Val == fmt.Sprintf("%x", sha512.Sum512(data))
		}
	}
	return false
}

func (x *Checksum) Sum(data []byte) *Checksum {
	if x != nil {
		if x.Algorithm == 0 {
			x.Algorithm = Checksum_ALGORITHM_SHA256
		}
		switch x.Algorithm {
		case Checksum_ALGORITHM_MD5:
			x.Val = fmt.Sprintf("%x", md5.Sum(data))
		case Checksum_ALGORITHM_SHA1:
			x.Val = fmt.Sprintf("%x", sha1.Sum(data))
		case Checksum_ALGORITHM_SHA256:
			x.Val = fmt.Sprintf("%x", sha256.Sum256(data))
		case Checksum_ALGORITHM_SHA512:
			x.Val = fmt.Sprintf("%x", sha512.Sum512(data))
		}
	}
	return x
}

func (x *Checksum) IsValid() bool {
	if !x.IsEmpty() {
		return len(x.Val) == checksumLength[x.Algorithm]
	}
	return false
}

func (x *Checksum) IsEmpty() bool {
	return x == nil || x.Algorithm == 0 || len(x.Val) == 0
}

func (x *Checksum) GetValue() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *Checksum) ToString() string {
	return x.Format()
}

func (x *Checksum) ToScalar() interface{} {
	return x.ToString()
}
