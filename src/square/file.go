package square

type File int8

const InvalidFile = File(-1)

func FileFromRune(r rune) File {
	if r == '-' {
		return InvalidFile
	}

	return File(r - 'A')
}

func (f File) Rune() rune {
	if f == InvalidFile {
		return '-'
	}
	return 'A' + rune(f)
}
