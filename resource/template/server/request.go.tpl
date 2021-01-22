package request

import "github.com/siuvlqnm/bookmark/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}