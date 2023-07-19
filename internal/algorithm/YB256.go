package algorithm

import (
	"encoding/binary"
	"fmt"
)

func (H *YB128Hash) StringHash(data []byte) string {
	return fmt.Sprintf("%x", H.HashBytes(data))
}

func (H *YB128Hash) Compress(state *State, block []byte) {
	a, b, c, d, e, f, g, h := state.A, state.B, state.C, state.D, state.E, state.F, state.G, state.H

	var words [blockSize / 4]uint32
	for i := 0; i < blockSize/4; i++ {
		words[i] = binary.LittleEndian.Uint32(block[i*4 : (i+1)*4])
	}

	Rounds := [8]uint32{
		round1Const, round2Const, round3Const, round4Const,
		round5Const, round6Const, round7Const, round8Const,
	}

	Values := [8]uint{7, 12, 17, 22, 7, 12, 17, 22}

	for i := 0; i < 8; i++ {
		var k, v uint32

		switch i {
		case 0:
			v = (b & c) | (^b & d)
			k = Rounds[i]
		case 1:
			v = (d & b) | (^d & c)
			k = Rounds[i]
		case 2:
			v = b ^ c ^ d
			k = Rounds[i]
		case 3:
			v = c ^ (b | ^d)
			k = Rounds[i]
		default:
			v = 0
			k = Rounds[i]
		}

		d = c
		c = b
		b += (a + v + words[i] + k)
		b = (b << Values[i]) | (b >> (32 - Values[i]))
		a = d

	}
	state.A += a
	state.B += b
	state.C += c
	state.D += d
	state.E += e
	state.F += f
	state.G += g
	state.H += h
}

func (H *YB128Hash) HashBytes(data []byte) []byte {
	state := State{
		A: initialA, B: initialB, C: initialC, D: initialD,
		E: initialE, F: initialF, G: initialG, H: initialH,
	}

	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize
		if end > len(data) {
			end = len(data)
		}
		block := data[i:end]
		if len(block) < blockSize {
			padding := make([]byte, blockSize-len(block))
			block = append(block, padding...)
		}

		H.Compress(&state, block)
	}

	hash := make([]byte, 32)

	for i, s := range []*uint32{
		&state.A, &state.B, &state.C, &state.D,
		&state.E, &state.F, &state.G, &state.H,
	} {
		binary.LittleEndian.PutUint32(hash[i*4:i*4+4], *s)
	}
	return hash
}
