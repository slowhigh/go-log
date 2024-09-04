package memcpy

func memcpy(data []byte, src, dest, size int64) {
	if src < 0 || dest < 0 || src+size > int64(len(data)) || dest+size > int64(len(data)) {
		panic("out of bounds memory access")
	}

	if dest < src {
		for i := int64(0); i < size; i++ {
			data[dest+i] = data[src+i]
		}
	} else {
		for i := int64(size); i >= 0; i-- {
			data[dest+i] = data[src+i]
		}
	}
}