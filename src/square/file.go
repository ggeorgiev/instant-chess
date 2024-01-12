package square

type File int8
type Files []File

const (
	InvalidFile = File(-1)
	ZeroFile    = File(0)
	LastFile    = File(7)
)

var ValidFiles = Files{ZeroFile, File(1), File(2), File(3), File(4), File(5), File(6), LastFile}
var AllFiles = Files{InvalidFile, ZeroFile, File(1), File(2), File(3), File(4), File(5), File(6), LastFile}

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
