package upload

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path/filepath"
	"time"
)

/*
* 文件类型验证
* @params mimieType string
* @params typeArr []string
* @return bool
*/
func FileType(mimeType string,typeArr []string) bool{
	for _,value := range typeArr{
		if mimeType == value{
			return true
		}
	}
	return false
}

/**
*  文件大小原则
 */
func FileSize(size,maxSize int64) bool{
	if size > maxSize {
		return false
	}
	return true
}
/*
* 生成文件名称，路径
*/
func GetFilePath(fileName string) (fpath string,urlPath string){
	path := time.Now().Format("20060102")
	appPath, _:= filepath.Abs(filepath.Dir(filepath.Join(".", string(filepath.Separator))))
	appPath = filepath.Dir(filepath.Dir(appPath))
	pt := appPath + "/upload/"+path
	b, err := PathExists(pt)
	if err != nil{
		logx.Error("path-err:",err)
	}
	if b == false{
		err = os.Mkdir(pt,os.ModePerm)
		logx.Error("mkdir-err:",err)
	}
	fpath = fmt.Sprintf("%s/%s",pt,fileName)
	urlPath = fmt.Sprintf("/upload/%s/%s",path,fileName)
	return fpath,urlPath
	
}
/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {
	
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
/*
* 保存文件
*/
func SaveFile(fileBuffer []byte, filename string){
	
	file,err := os.Create(filename)
	if err != nil{
		logx.Error("fileError:",err)
	}
	defer file.Close()
	file.Write(fileBuffer)
}