package base62

import ()

const encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type Encoding struct {
	decodeMap [256]byte
}

func New() *Encoding {
	enc := &Encoding{}

	for i := 0; i < len(encoding); i++ {
		enc.decodeMap[encoding[i]] = byte(i)
	}

	return enc
}

func (enc *Encoding) EncodeToString(src []byte) string {
	dst := make([]byte, len(src)*2)
	n := enc.Encode(dst, src)
	return string(dst[:n])
}

func (enc *Encoding) Encode(dst, src []byte) int {
	rem := 0
	var r uint
	j := 0

	for i := 0; i < len(src); i++ {
		r = (r << 8) | uint(src[i])
		rem += 8

		for rem >= 6 {
			cur := (r >> (rem - 6)) & 63

			consumed := 6
			if cur >= 61 {
				consumed = 4
				cur = 61
			}

			dst[j] = encoding[cur]
			j++

			rem -= consumed
		}
	}

	if rem > 0 {
		cur := r & (63 >> (6 - rem)) //isolate remainder
		cur = cur << (6 - rem)       //left align
		dst[j] = encoding[cur]
		j++
	}

	dst = dst[:j]
	return j
}

func (enc *Encoding) DecodeString(s string) []byte {
	dst := make([]byte, len(s))
	n := enc.Decode(dst, []byte(s))
	return dst[:n]
}

func (enc *Encoding) Decode(dst, src []byte) int {
	rem := 0
	var r uint
	j := 0

	for i := 0; i < len(src); i++ {
		used := 6
		read := uint(enc.decodeMap[src[i]])
		if read >= 61 {
			used = 4
			read = read >> 2
		}

		r = (r << used) | read
		rem += used

		for rem >= 8 {
			cur := (r >> (rem - 8)) & 255
			dst[j] = byte(cur)
			j++

			rem -= 8
		}
	}

	if rem > 0 {
		v := r & (255 >> (8 - rem))
		//check if its just padding
		if v > 0 {
			dst[j] = byte(v)
			j++
		}
	}

	return j
}
