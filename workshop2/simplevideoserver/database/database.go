package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type VideoData struct {
	ID       string
	Name     string
	Duration int
}

type PendingVideoData struct {
	DbID int
	ID   string
}

const statusPending = 1
const statusInProgress = 2
const statusComplete = 3

type Database struct {
	db *sql.DB
}

func Open() Database {
	db, err := sql.Open("mysql", "root:root@/gocourse")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return Database{db}
}

func (db Database) Close() {
	db.db.Close()
}

func (db Database) exec(query string) {
	if _, err := db.db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func (db Database) Init() {
	db.exec("DROP TABLE IF EXISTS video;")

	db.exec(
		`CREATE TABLE video
(
	id            INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,
	video_key     VARCHAR(255) UNIQUE,
	title         VARCHAR(255) NOT NULL,
	status        TINYINT DEFAULT 1,
	duration      INT UNSIGNED DEFAULT 0,
	PRIMARY KEY (id)
);`)
}

func (db Database) GetVideoList(skip int, take int) ([]VideoData, error) {
	elements := make([]VideoData, 0)

	rows, err := db.db.Query("SELECT video_key, title, duration FROM video LIMIT ?, ?;", skip, take)
	if err != nil {
		return elements, err
	}
	defer rows.Close()

	for rows.Next() {
		var element VideoData
		if err := rows.Scan(&element.ID, &element.Name, &element.Duration); err != nil {
			return elements, err
		}
		elements = append(elements, element)
	}

	return elements, nil
}

func (db Database) GetVideoDataByID(id string) (VideoData, error) {
	result := VideoData{}

	rows, err := db.db.Query("SELECT video_key, title, duration FROM video WHERE video_key=?;", id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Name, &result.Duration); err != nil {
			return result, err
		}
	}

	return result, nil
}

func (db Database) AddVideo(data VideoData) {
	if _, err := db.db.Exec("INSERT INTO video (video_key, title, duration) VALUES (?, ?, ?);", data.ID, data.Name, data.Duration); err != nil {
		log.Fatal(err)
	}
}

func (db Database) GetPendingVideoList() ([]PendingVideoData, error) {
	elements := make([]PendingVideoData, 0)

	rows, err := db.db.Query("SELECT id, video_key FROM video WHERE status=?;", statusPending)
	if err != nil {
		return elements, err
	}
	defer rows.Close()

	for rows.Next() {
		var element PendingVideoData
		if err := rows.Scan(&element.DbID, &element.ID); err != nil {
			return elements, err
		}
		elements = append(elements, element)
	}

	return elements, nil
}

func (db Database) MarkVideoAsInProgress(dbID int) {
	if _, err := db.db.Exec("UPDATE video SET status = ? WHERE id = ?;", statusInProgress, dbID); err != nil {
		log.Fatal(err)
	}
}

func (db Database) MarkInProgressVideosAsPending() {
	if _, err := db.db.Exec("UPDATE video SET status = ? WHERE status = ?;", statusPending, statusInProgress); err != nil {
		log.Fatal(err)
	}
}

func (db Database) MarkVideoAsComplete(dbID int, duration int) {
	if _, err := db.db.Exec("UPDATE video SET status = ?, duration = ? WHERE id = ?;", statusComplete, duration, dbID); err != nil {
		log.Fatal(err)
	}
}
