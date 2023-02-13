package database

import (
	"fmt"
	"log"
)

func PrintAllCourses() {
	db := GetDb()

	defer db.Close()

	rows, err := db.Query("SELECT * FROM courses")

	if err != nil {
		log.Fatalln("Couldn't query courses")
	}

	for rows.Next() {
		var course_id int
		var name string

		err = rows.Scan(&course_id, &name)

		if err != nil {
			log.Fatalln("Couldn't scan course rows")
		}

		fmt.Printf("%d: %s\n", course_id, name)
	}
}
