package frames

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames)-1)]
	}
}

// Create a function that returns frame length
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
	}
}

var FrameMap = map[string]FrameType{
	"forrest":         Forrest,
	"parrot":          Parrot,
	"clock":           Clock,
	"nyan":            Nyan,
	"rick":            Rick,
	"can-you-hear-me": Rick,
	"donut":           Donut,
	"coin":            Coin,
	"popcat":          Popcat,
	"batman":          Batman,
	"fuck":            Fuck,
}

const colorRed = "\033[1;31m"
const colorGreen = "\033[1;32m"
const colorYellow = "\033[1;33m"
const colorBlue = "\033[1;34m"
const colorMagenta = "\033[1;35m"
const colorCyan = "\033[1;36m"
const colorWhite = "\033[1;37m"
const colorReset = "\033[0m"
const colorOrange = "\033[38;5;208m"
const colorPink = "\033[38;5;206m"
const colorPurple = "\033[38;5;141m"
const colorBrown = "\033[38;5;130m"