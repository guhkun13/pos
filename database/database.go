package database

import "guhkun13/pizza-api/config"

type Database struct {
	PgxPool *PgxPoolConn
}

func NewDatabase(env *config.EnvironmentVariables) *Database {
	return &Database{
		PgxPool: NewPgxPoolConn(env),
	}
}

func (d *Database) Close() {
	d.PgxPool.Conn.Close()
}
