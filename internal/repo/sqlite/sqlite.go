package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func Open(ctx context.Context, dataSourceName string) (*sql.DB, error) {
	db, err := open(ctx, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := migration(ctx, dataSourceName); err != nil {
		return nil, err
	}

	return db, nil
}

func open(ctx context.Context, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

//go:embed schema.sql
var embedded embed.FS

func migration(ctx context.Context, dataSourceName string) error {
	oldData, err := os.ReadFile(dataSourceName)
	if err != nil {
		return err
	}
	schema, err := embedded.ReadFile("schema.sql")
	if err != nil {
		return err
	}

	dir, err := atlasexec.NewWorkingDir(func(ce *atlasexec.WorkingDir) error {
		if _, err := ce.WriteFile("db.sqlite3", oldData); err != nil {
			return err
		}
		if _, err := ce.WriteFile("schema.sql", schema); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	client, err := atlasexec.NewClient(dir.Path(), "atlas")
	if err != nil {
		return err
	}

	res, err := client.SchemaApply(ctx, &atlasexec.SchemaApplyParams{
		AutoApprove: true,
		DevURL:      "sqlite://dev.sqlite3?mode=memory",
		To:          "file://schema.sql",
		URL:         "sqlite://db.sqlite3",
	})
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations\n", len(res.Changes.Applied))
	if len(res.Changes.Applied) > 0 {
		newData, err := fs.ReadFile(dir.DirFS(), "db.sqlite3")
		if err != nil {
			return err
		}
		info, err := os.Stat(dataSourceName)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dataSourceName, newData, info.Mode().Perm()); err != nil {
			return err
		}
	}

	return nil
}
