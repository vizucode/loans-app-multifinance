package formatting

func OptimizeUUIDBytes(b []byte) []byte {
	if len(b) != 16 {
		panic("UUID must be 16 bytes")
	}
	return append(
		append(
			append(b[6:8], b[4:6]...),
			b[0:4]...,
		),
		b[8:]...,
	)
}
