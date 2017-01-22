package repo

import ()

type (
	Repo interface {
		Start() bool
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		Commit() error
		Rollback() error
	}

	Data struct {
		Value interface{}
	}
)
