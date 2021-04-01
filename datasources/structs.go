package datasources

func NewDoc() DatasourceDoc {
	return DatasourceDoc{Datasources: map[string]Datasource{}}
}

type DatasourceDoc struct {
	Datasources map[string]Datasource
}

type Datasource struct {
	Name     string
	Type     string
	Host     string
	Port     uint16
	Sid      string
	Username string
	Password string
}
