package postgres

import "time"

type Option func(pg *Postgres)

func WithTimeOut(timeOut time.Duration) Option {
	return func(pg *Postgres) {
		pg.timeout = timeOut
	}
}

func WithPoolSize(size int) Option {
	return func(pg *Postgres) {
		pg.poolSize = size
	}
}

func WithAttempts(count int) Option {
	return func(pg *Postgres) {
		pg.attempts = count
	}
}
