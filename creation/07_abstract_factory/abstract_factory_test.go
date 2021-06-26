package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
    sql := MYSQLDataFactory{}
	sqlEx :=sql.CreateExport()
	sqlIMport := sql.CreateImport()
	sqlEx.Export("sql test")
	sqlIMport.Import("sql test")


	file := FileDataFactory{}
	fileEx :=file.CreateExport()
	fileIMport := file.CreateImport()
	fileEx.Export("file test")
	fileIMport.Import("file dsd")

}
