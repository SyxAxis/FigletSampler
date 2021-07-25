package main

import (
	"embed"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	figletlib "github.com/lukesampson/figlet/figletlib"
)

var SrcFolder string = "fonts"
var OutputMessage string = "Sample message"

//go:embed fonts/*.flf
var embededFiles embed.FS

func ShowAllFontsWithSamples(showSamples bool) {

	fileList, _ := embededFiles.ReadDir(SrcFolder)
	for id, fileName := range fileList {
		if showSamples {
			err := OpenFontFile(id, fileName.Name(), true)
			if err != nil {
				fmt.Println("Could not find that font!")
			}
		} else {
			fmt.Printf(" %v\n", fileName.Name()[:len(fileName.Name())-4])
		}
	}

}

func GetRandomFont() string {

	minRnd := 1
	getFontFileList, _ := embededFiles.ReadDir(SrcFolder)
	maxRnd := len(getFontFileList)
	rndFontNum := rand.Intn(maxRnd-minRnd+1) + minRnd
	return getFontFileList[rndFontNum].Name()
}

func OpenFontFile(id int, fontFileName string, showBanner bool) error {

	srcFontLocnRef := SrcFolder + "/" + fontFileName
	fontFileName = fontFileName[:len(fontFileName)-4]

	if showBanner {
		fmt.Println("=================================================================================================================")
		fmt.Printf("   LOADING : ( %v ) %v\n", id, fontFileName)
		fmt.Println("=================================================================================================================")
	}

	fontBytes, err := embededFiles.ReadFile(srcFontLocnRef)
	if err != nil {
		fmt.Printf("Cannot find font : %v\n", fontFileName)
	}

	fontData, err := figletlib.ReadFontFromBytes(fontBytes)
	if err != nil {
		return err
	}
	DisplayFont(fontData)

	return nil
}

func DisplayFont(f *figletlib.Font) {
	figletlib.PrintMsg(OutputMessage, f, 180, f.Settings(), "left")
}

func main() {

	flgShowAllFonts := flag.Bool("show_all_fonts", false, "show all fonts")
	flgShowSamples := flag.Bool("show_samples", false, "show samples")
	flgShowOne := flag.String("showone", "larry3d.flf", "show one font")
	flgShowRandomFont := flag.Bool("show_random_font", false, "show any font at random")
	flgNoBanner := flag.Bool("show_no_banner", true, "show the font info banner")
	flgOptMessage := flag.String("message", "Wiggy Pig", "optional sample message")
	flag.Parse()

	OutputMessage = *flgOptMessage

	if *flgShowAllFonts {
		if *flgShowSamples {
			ShowAllFontsWithSamples(*flgShowSamples)
		} else {
			ShowAllFontsWithSamples(*flgShowSamples)
			fmt.Println("\nAdd the flag -show_samples to see a sample of each font.\n")
		}
	} else {
		var selectedFontName string
		if *flgShowRandomFont {
			rand.Seed(time.Now().UnixNano())
			selectedFontName = GetRandomFont()
		} else {
			selectedFontName = *flgShowOne
		}

		// every filename into OpenFontFile must be fully formed
		if !strings.HasSuffix(selectedFontName, ".flf") {
			selectedFontName = selectedFontName + ".flf"
		}
		OpenFontFile(1, selectedFontName, *flgNoBanner)

	}

}
