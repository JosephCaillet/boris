package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	startTime             time.Time
	progress, progressMax int
	errorColor            = color.New(color.BgRed, color.FgBlack)
	warnColor             = color.New(color.BgYellow, color.FgBlack)
)

func Mode(mode string) {
	fmt.Printf("⏺  %s ⏺\n", mode)
}

func Error(message string, err error) {
	fmt.Printf("❌ %s: %v", message, err)
}

func StartOperation(stepNumber int) {
	progress = 0
	progressMax = stepNumber
	startTime = time.Now()
}

func ProgressOperation() {
	progress++
}

func getPrefix() string {
	color.Unset()
	return fmt.Sprintf("[ %d%% ][ %s ]",
		int(float32(progress)/float32(progressMax)*100.0),
		time.Since(startTime).Round(time.Second),
	)
}

func EnteringFolder(folder string) {
	p := getPrefix()
	fmt.Printf("%s\n%s\t↳ %s/\n", p, p, folder)
}

func MoveFile(oldPath, newPath string, musicFile bool) {
	p := getPrefix()
	icon := map[bool]string{true: "♫", false: "🖺"}[musicFile]
	fmt.Printf("%s\t\t%s %s\t➜\t%s\n", p, icon, oldPath, newPath)
}

func ErrorTag(err error) {
	p := errorColor.Sprint(getPrefix())
	fmt.Printf("%s\t\t❌ error: %v\n", p, err)
}

func WarnWrongMove() {
	p := warnColor.Sprint(getPrefix())
	fmt.Printf("%s\t\t⚠ No tagged music file found, moving file(s) below to unrecognised music directory.\n", p)
}
