package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Please specify a database file and an output directory.")
	}
	dbfile := os.Args[1]
	outdir := os.Args[2]

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Play it cool with fields that might be empty
	// But we won't write an m3u entry if either the path or file is null.
	//
	type Video struct {
		name     sql.NullString
		artist   sql.NullString
		path     sql.NullString
		file     sql.NullString
		duration sql.NullString
	}

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Write a new .m3u file for each genre
	genres, err := db.Query("select genre_id, name from genre")
	if err != nil {
		log.Fatal(err)
	}
	defer genres.Close()

	err = os.Mkdir(path.Join(cwd, "m3u"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	for genres.Next() {
		var gid int
		var genreName string
		err = genres.Scan(&gid, &genreName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Found genre " + genreName)

		f, err := os.Create(path.Join(cwd, "m3u", genreName+".m3u"))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString("#EXTM3U\n")

		// Get music videos
		videos, err := db.Query("select musicvideo_view.c00 as name, musicvideo_view.c10 as artist, musicvideo_view.strPath as path, musicvideo_view.strFileName as file, musicvideo_view.totalTimeInSeconds as duration from genre_link, musicvideo_view where (genre_link.media_id = musicvideo_view.idMVideo) and media_type = 'musicvideo' and genre_id = ?", gid)

		if err != nil {
			log.Fatal(err)
		}
		defer videos.Close()
		for videos.Next() {

			v := Video{}
			err = videos.Scan(&v.name, &v.artist, &v.path, &v.file, &v.duration)
			if err != nil {
				log.Fatal(err)
			}

			artist := "Unknown"
			if v.artist.Valid {
				artist = v.artist.String
			}

			name := "Unknown"
			if v.name.Valid {
				name = v.name.String
			}

			fmt.Printf("  %s - %s\n", artist, name)

			if v.path.Valid && v.file.Valid {
				duration := "-1"
				if v.duration.Valid {
					duration = v.duration.String
				}

				// #EXTINF:123, Sample artist - Sample title
				// C:\Documents and Settings\I\My Music\Sample.mp3
				f.WriteString(fmt.Sprintf("#EXTINF:%s,%s\n", duration, name))
				f.WriteString(fmt.Sprintf("%s%s\n", v.path.String, v.file.String))
			}
		}
	}

	err = genres.Err()
	if err != nil {
		log.Fatal(err)
	}
}
