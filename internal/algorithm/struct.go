package algorithm

const blockSize = 64

type YB128Hash struct{}

type State struct {
	A, B, C, D, E, F, G, H uint32
}
