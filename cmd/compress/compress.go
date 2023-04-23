package compress

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/h2non/bimg"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		return
	}
	filePath := args[0]
	fileOut := ""
	format := cmd.Flag("format").Value.String()
	if len(args) == 2 {
		fileOut = args[1]
	}
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if fileOut == "" {
		n := strings.Split(f.Name(), "/")
		fileOut = strings.Split(n[len(n)-1], ".")[0] + "compressed." + format
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	img := bimg.NewImage(buf)
	fType := bimg.WEBP
	switch format {
	case "png":
		fType = bimg.PNG
	}
	buf, err = img.Convert(fType)
	if err != nil {
		panic(err)
	}
	if fileOut == "" {
		n := strings.Split(f.Name(), "/")
		fileOut = n[len(n)-1] + "compressed." + format
	}
	fmt.Println(fileOut)
	ioutil.WriteFile(fileOut, buf, 0755)
}
