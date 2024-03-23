package dbintegration

import (
	"database/sql"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DatabaseName string = "db/db-test.sqlite3"
var DatabaseInitializer string = "db/db-setup.sql"

func nullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func Init() {
	os.Remove(DatabaseName)
	db, err := sqlx.Connect("sqlite", DatabaseName)
	if err != nil {
		log.Fatalln(err)
	}

	content, err := os.ReadFile(DatabaseInitializer)
	if err != nil {
		log.Fatalln(err)
	}
	var queryInitializer string = string(content)
	db.MustExec(queryInitializer)

	insertGenres(db)
	insertType(db)
	insertStudios(db)
	insertStatuses(db)
	insertSeasons(db)
	insertImages(db)
	insertCharacters(db)
	insertDescriptions(db)
	insertAnime(db)
	insertEpisode(db)

	insertJoinAnimeCharacters(db)
	insertJoinAnimeDescriptions(db)
	insertJoinAnimeGenres(db)
	insertJoinAnimeImages(db)
	insertJoinAnimeStudios(db)
	insertJoinCharacterDescriptions(db)
	insertJoinCharacterImages(db)
	insertJoinEpisodeDescriptions(db)

	db.Close()
}

// https://jmoiron.github.io/sqlx/
// https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme
// https://gitlab.com/cznic/sqlite/-/blob/v1.29.5/examples/example1/main.go
