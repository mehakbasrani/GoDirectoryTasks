package main

import(
	"fmt"
	"os"
	"path/filepath"
)
func main(){

	dir := "F:/go"
    size,  err := totlaSizeOfDir(dir)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Size:", size)
}

func totlaSizeOfDir(dir string) (int64, error) {
        var	dirSize int64

   err := filepath.Walk(dir,func(path string, info os.FileInfo, err error) error {
		
	if err != nil{
			return err
		}
		
		if !info.IsDir(){
			dirSize += info.Size()
		}
		return nil
	})

	if err != nil{
		return  0, err
	}

return dirSize, nil
}