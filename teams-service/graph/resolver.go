package graph

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Resolver struct {
    db      *sqlx.DB
    Logger  *zap.Logger
}

func NewResolver(db *sqlx.DB, logger *zap.Logger) *Resolver {
    return &Resolver{
        db:     db,
        Logger: logger,
    }
}
