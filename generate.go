package sdk

//go:generate swagger/swagger-codegen generate --config swagger/config.json -Dmodels -l go -i swagger/databricks-2.0.yaml -o models
//go:generate gofmt -s -w models
