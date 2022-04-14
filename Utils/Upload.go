package Utils

import (
	"FusionGO/Conf"
	"io"
	"net/http"
	"os"
)

func Receivefile(r *http.Request) map[int]map[string]interface{} {
	//设置内存缓冲区大小
	r.ParseMultipartForm(2048 << 20)
	if len(r.MultipartForm.Value) != 0 {
		//获取上传的文件组，值为前端formdata.append的组名
		files := r.MultipartForm.File["file[]"]
		values := r.MultipartForm.Value["post"][0]
		lens := len(files)
		filelist := make(map[int]map[string]interface{})
		for i := 0; i < lens; i++ {
			//创建文件
			file, err := files[i].Open()
			targetPath := Conf.FILE_TEMP_DIR + "/" + Time() + files[i].Filename
			filelist[i] = make(map[string]interface{})
			filelist[i]["temp_target_path"] = targetPath
			filelist[i]["file_Args"] = values
			filelist[i]["file_name"] = files[i].Filename
			filepath, err := os.Create(targetPath)
			if err != nil {
				panic(err)
			}
			//写入二进制数据
			_, err = io.Copy(filepath, file)
			if err != nil {
				panic(err)
			}
			// 结束打开的文件释放占用状态（内存）
			file.Close()
			// 结束创建的新文件的占用状态）（硬盘）
			filepath.Close()
		}

		return filelist
	}
	files := r.MultipartForm.File["file[]"]
	lens := len(files)
	filelist := make(map[int]map[string]interface{})
	for i := 0; i < lens; i++ {
		//创建文件
		file, err := files[i].Open()
		targetPath := Conf.FILE_TEMP_DIR + "/" + Time() + files[i].Filename
		filelist[i] = make(map[string]interface{})
		filelist[i]["temp_target_path"] = targetPath
		filepath, err := os.Create(targetPath)
		if err != nil {
			panic(err)
		}
		//写入二进制数据
		_, err = io.Copy(filepath, file)
		if err != nil {
			panic(err)
		}
		// 结束打开的文件释放占用状态（内存）
		file.Close()
		// 结束创建的新文件的占用状态）（硬盘）
		filepath.Close()
	}
	return filelist
}
