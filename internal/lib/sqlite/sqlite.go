package sqlite

import (
	"context"
	"database/sql"
	"flag"
	"io/fs"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type Config struct {
	Source string
}

func NewConfig(fs *flag.FlagSet) *Config {
	config := &Config{}
	fs.StringVar(&config.Source, "sqlite.source", "db.sqlite3", "")
	return config
}

func Open(ctx context.Context, source string, schema []byte) (*sql.DB, error) {
	db, err := open(ctx, source)
	if err != nil {
		return nil, err
	}

	if err := migration(ctx, source, schema); err != nil {
		return nil, err
	}

	return db, nil
}

func open(ctx context.Context, source string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func migration(ctx context.Context, source string, schema []byte) error {
	oldData, err := os.ReadFile(source)
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

	log.Printf("[sqlite] Applied %d migrations\n", len(res.Changes.Applied))
	if len(res.Changes.Applied) > 0 {
		newData, err := fs.ReadFile(dir.DirFS(), "db.sqlite3")
		if err != nil {
			return err
		}
		info, err := os.Stat(source)
		if err != nil {
			return err
		}
		if err := os.WriteFile(source, newData, info.Mode().Perm()); err != nil {
			return err
		}
	}

	return nil
}
