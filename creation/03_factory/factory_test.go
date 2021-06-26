package factory

import "testing"

func TestFactory(t *testing.T) {
	dbFactory := ExportDBFactory{}
	db := dbFactory.CreateExport()
	db.Export("db")

	fileFactory := ExportTxtFileFactory{}
	file := fileFactory.CreateExport()
	file.Export("file")
	
}