package sinks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opsgenie/kubernetes-event-exporter/pkg/kube"
)

type T4Config struct {
	Path string `yaml:"path"`
}

type T4 struct {
	db *sql.DB
}

func NewT4Sink() (*T4, error) {
	db, err := sql.Open("mysql", "local:local@/test")
	if err != nil {
		return nil, err
	}

	t := &T4{
		db: db,
	}
	return t, nil
}

func (t *T4) Send(ctx context.Context, ev *kube.EnhancedEvent) error {
	v, err := json.Marshal(ev)
	if err != nil {
		return fmt.Errorf("json error on send: %s", err)
	}

	_, err = t.db.Exec("insert into events (event) values (?)", v)
	return err
}

func (t *T4) Close() {
	t.db.Close()
}
