package sqlite

import _ "embed"

//go:embed bob/schema.sql
var Schema []byte
