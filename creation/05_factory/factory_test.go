package factory

import "testing"

func TestFactory(t *testing.T) {
	dbFactory := ExportDBFactory{}
	db := dbFactory.Build()
	db.Export("db")

	fileFactory := ExportTxtFileFactory{}
	file := fileFactory.Build()
	file.Export("file")
	
}