package test

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"testing"
)

func Test_Write_Excel(t *testing.T) {
	f := excelize.NewFile()
	f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "知识标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "知识类型")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "一级分类")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "二级分类")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "阅读下载权限")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "积分设置")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "知识内容")
	f.SetCellValue("sheet1", "A"+strconv.Itoa(2), "测试标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(2), "普通")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(2), "招标采购")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(2), "采购制度")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(2), "所有")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(2), "50")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(2), "知识内容")
	// 根据指定路径保存文件
	if err := f.SaveAs("template.xlsx"); err != nil {
		fmt.Println(err)
	}
}

//excelize.OpenFile()

func Test_Parse_Excel(t *testing.T) {
	f, err := excelize.OpenFile("template.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(rows) > 1 {
		for _, row := range rows[1:] {
			// 标题
			fmt.Print(row[0], "\t")
			// 知识类型
			fmt.Print(row[1], "\t")
			// 一级分类
			fmt.Print(row[2], "\t")
			// 二级分类
			fmt.Print(row[3], "\t")
			// 阅读下载权限
			fmt.Print(row[4], "\t")
			// 积分设置
			fmt.Print(row[5], "\t")
			// 知识内容
			fmt.Print(row[6], "\t")
			break
			fmt.Println()
		}
	}
}
