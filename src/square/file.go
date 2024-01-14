package square

type File int8
type Files []File

const (
	InvalidFile = File(-1)
	ZeroFile    = File(0)
	LastFile    = File(7)
)

var ValidFiles = Files{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}
var AllFiles = Files{InvalidFile, FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

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
