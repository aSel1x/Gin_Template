package adapters

type Adapters struct {
	*Postgres
}

func NewAdapters(postgresDsn string) *Adapters {
	postgres, _ := NewPostgres(postgresDsn)
	return &Adapters{
		postgres,
	}
}
