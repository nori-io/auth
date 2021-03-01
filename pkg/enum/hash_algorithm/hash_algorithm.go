package hash_algorithm

type HashAlgorithm uint8

const Bcrypt HashAlgorithm = iota

func (h HashAlgorithm) Value() uint8 {
	return uint8(h)
}

func New(algorithm uint8) HashAlgorithm {
	return HashAlgorithm(algorithm)
}
