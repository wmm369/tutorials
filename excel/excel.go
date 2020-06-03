package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	createExcel()
	readExcel()
}

func createExcel()  {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")

	//隐藏名称为 Sheet1 的工作表中的 D 至 F 列：
	f.SetColVisible("Sheet1", "D:F", false)

	//根据给定的工作表名称（大小写敏感）、列范围和宽度值设置单个或多个列的宽度。
	//例如设置名为 Sheet1 工作表上 A 到 D 列的宽度为 20：
	f.SetColWidth("Sheet1", "A", "D", 20)

	//根据给定的工作表名称（大小写敏感）、行号和高度值设置单行高度。
	//例如设置名为 Sheet1 工作表第二行行的高度为 50：
	f.SetRowHeight("Sheet1", 2, 50)

	// 设置单元格的值
	f.SetCellValue("Sheet1", "A2", 100)

	//根据给定的工作表名（大小写敏感）和单元格坐标区域合并单元格。例如，合并名为 Sheet1 的工作表上 D3:E9 区域内的单元格：
	f.MergeCell("Sheet1", "D3", "D4")


	f.SetCellValue("Sheet1", "D3", 1000)
	f.SetCellValue("Sheet1", "D5", "hello")

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}

func readExcel()  {
	f, err := excelize.OpenFile("book.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell := f.GetCellValue("Sheet1", "B2")

	println(cell)

	// 获取 Sheet1 上所有单元格
	rows := f.GetRows("Sheet1")

	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
}
