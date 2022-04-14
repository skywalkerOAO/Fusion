package Utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ChangeExcelValue(filepath string, Sheet string, axis string, value interface{}) (bool, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	// 获取工作表中指定单元格的值
	f.SetCellValue(Sheet, axis, value)
	f.UpdateLinkedValue() // 更新Excel链接值
	err = f.SaveAs(filepath)
	if err != nil {
		return false, err
	}
	return true, nil
}
