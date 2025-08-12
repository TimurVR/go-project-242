package goproject242
import ("os"
"path/filepath"
"io/fs"
"strconv")

func GetPathSize(path string, recursive, human, all bool) (string, error){
	fileInfo, err := os.Lstat(path)

	var totalSize int64
	if fileInfo.Mode().IsRegular(){
		totalSize=fileInfo.Size()
	}else{	
		err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			totalSize += info.Size()
		}
		return nil
	})
	}
	s:=strconv.FormatInt(totalSize,10)
	return s+"\t"+path,err
}