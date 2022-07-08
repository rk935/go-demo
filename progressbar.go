package main

import (
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	//bar := progressbar.NewOptions(1000,
	//	//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
	//	progressbar.OptionEnableColorCodes(true),
	//	progressbar.OptionShowBytes(true),
	//	progressbar.OptionSetWidth(15),
	//	progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
	//	progressbar.OptionSetTheme(progressbar.Theme{
	//		Saucer:        "[green]=[reset]",
	//		SaucerHead:    "[green]>[reset]",
	//		SaucerPadding: " ",
	//		BarStart:      "[",
	//		BarEnd:        "]",
	//	}))
	//for i := 0; i < 1000; i++ {
	//	bar.Add(1)
	//	time.Sleep(5 * time.Millisecond)
	//}
}
