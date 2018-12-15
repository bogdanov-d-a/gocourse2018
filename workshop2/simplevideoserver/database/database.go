package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type VideoData struct {
	Id       string
	Name     string
	Duration int
}

func connectAndHandle(handler func(db *sql.DB)) {
	db, err := sql.Open("mysql", "root:root@/gocourse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	handler(db)
}

func exec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	connectAndHandle(func(db *sql.DB) {
		exec(db, "DROP TABLE IF EXISTS video;")

		exec(db,
			`CREATE TABLE video
(
	id            INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,
	video_key     VARCHAR(255) UNIQUE,
	title         VARCHAR(255) NOT NULL,
	duration      INT UNSIGNED DEFAULT 0,
	PRIMARY KEY (id)
);`)
	})
}

func GetVideoList() ([]VideoData, error) {
	elements := make([]VideoData, 0)
	var err_result error = nil

	connectAndHandle(func(db *sql.DB) {
		rows, err := db.Query("SELECT video_key, title, duration FROM video;")
		if err != nil {
			err_result = err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var element VideoData
			err := rows.Scan(&element.Id, &element.Name, &element.Duration)
			if err != nil {
				err_result = err
				return
			}
			elements = append(elements, element)
		}
	})

	return elements, err_result
}

func GetVideoListDataById(id string) (VideoData, error) {
	result := VideoData{}
	var err_result error = nil

	connectAndHandle(func(db *sql.DB) {
		rows, err := db.Query("SELECT video_key, title, duration FROM video WHERE video_key=?;", id)
		if err != nil {
			err_result = err
			return
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&result.Id, &result.Name, &result.Duration)
			if err != nil {
				err_result = err
				return
			}
		}
	})

	return result, err_result
}

func AddVideo(data VideoData) {
	connectAndHandle(func(db *sql.DB) {
		_, err := db.Exec("INSERT INTO video (video_key, title, duration) VALUES (?, ?, ?);", data.Id, data.Name, data.Duration)
		if err != nil {
			log.Fatal(err)
		}
	})
}
