package reporters

var (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Purple        = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
	Reverse       = "\u001b[7m"
	Bold          = "\u001b[1m"
	Underline     = "\u001b[4m"
)

// Colorize wraps a given message in a given color.
func Colorize(color, message string) string {
	return color + message + Reset
}
