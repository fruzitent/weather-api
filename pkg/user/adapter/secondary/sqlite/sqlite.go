package sqlite

import _ "embed"

//go:embed bobgen/schema.sql
var Schema []byte
