package main

import(
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	dir := "F:/go"
	name, size, path, err := findLargestFile(dir)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Name:", name, " ","Size:", size, " ","Path:", path)
}

func findLargestFile(dir string) (string, int64, string, error) {
        var	largestFileSize int64
        var largestFileName string
        var	largestFilePath string

   err := filepath.Walk(dir,func(path string, info os.FileInfo, err error) error {
		
	if err != nil{
			return err
		}

		if !info.IsDir(){
			if info.Size() > largestFileSize{
				largestFileSize = info.Size()
				largestFileName = info.Name()
				largestFilePath = path
			}
		}
		return nil
	})

	if err != nil{
		return "", 0, "", err
	}

return largestFileName, largestFileSize, largestFilePath, nil
}