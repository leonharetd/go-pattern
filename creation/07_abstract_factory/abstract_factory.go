package abstractfactory

import "fmt"

type IDataFactory interface {
	CreateExport() ExportAPI
	CreateImport() ImportAPI
}

type MYSQLDataFactory struct {}
func (sql *MYSQLDataFactory) CreateExport() ExportAPI {
    return &ExportDB{}
}
func (sql *MYSQLDataFactory) CreateImport() ImportAPI {
	return &ImportDB{}
}

type FileDataFactory struct {}
func (sql *FileDataFactory) CreateExport() ExportAPI {
    return &ExportFile{}
}
func (sql *FileDataFactory) CreateImport() ImportAPI {
	return &ImportFile{}
}

type ExportAPI interface {
	Export(data string) bool
}

type ImportAPI interface {
	Import(data string) bool
}

type ExportDB struct {}
func (db *ExportDB) Export(data string) bool {
	fmt.Println("Export DB", data)
	return true
}

type ExportFile struct {}
func (db *ExportFile) Export(data string) bool {
	fmt.Println("Export File", data)
	return true
}

type ImportFile struct {
	Data string
}
func (db *ImportFile) Import(data string) bool {
	db.Data = data
	fmt.Println("Import File", data)
	return true
}

type ImportDB struct {
	Data string
}
func (db *ImportDB) Import(data string) bool {
	db.Data = data
	fmt.Println("Import DB", data)
	return true
}
