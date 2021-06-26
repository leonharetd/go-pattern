package factory

import "fmt"

type ExportApi interface {
	Export(data string) bool
}

type ExportTxtFile struct{}

func (e *ExportTxtFile) Export(data string) bool {
	fmt.Println("export txt file", data)
	return true
}

type ExportDB struct{}

func (db *ExportDB) Export(data string) bool {
	fmt.Println("export db", data)
	return true
}

type IExportFactory interface {
	CreateExport() ExportApi
}

type ExportTxtFileFactory struct{}
func (e *ExportTxtFileFactory) CreateExport() ExportApi {
	return &ExportTxtFile{}
}

type ExportDBFactory struct{}
func (e *ExportDBFactory) CreateExport() ExportApi {
	return &ExportDB{}
}
