package dbintegration

import (
	"database/sql"
	"log"
	"os"

	db "AniGo/db"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

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
	os.Remove(db.DatabaseName)
	conn, err := sqlx.Connect(db.DatabaseDriver, db.DatabaseName)
	if err != nil {
		log.Fatalln(err)
	}

	content, err := os.ReadFile(db.DatabaseInitializer)
	if err != nil {
		log.Fatalln(err)
	}
	var queryInitializer string = string(content)
	conn.MustExec(queryInitializer)

	insertGenres(conn)
	insertType(conn)
	insertStudios(conn)
	insertStatuses(conn)
	insertSeasons(conn)
	insertImages(conn)
	insertCharacters(conn)
	insertDescriptions(conn)
	insertAnime(conn)
	insertEpisode(conn)

	insertJoinAnimeCharacters(conn)
	insertJoinAnimeDescriptions(conn)
	insertJoinAnimeGenres(conn)
	insertJoinAnimeImages(conn)
	insertJoinAnimeStudios(conn)
	insertJoinCharacterDescriptions(conn)
	insertJoinCharacterImages(conn)
	insertJoinEpisodeDescriptions(conn)

	conn.Close()
}

// https://jmoiron.github.io/sqlx/
// https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme
// https://gitlab.com/cznic/sqlite/-/blob/v1.29.5/examples/example1/main.go
