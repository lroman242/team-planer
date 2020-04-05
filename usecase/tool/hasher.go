package tool

type Hasher interface {
	Hash(value string) ([]byte, error)
	Compare(value string, hash []byte) bool
}
