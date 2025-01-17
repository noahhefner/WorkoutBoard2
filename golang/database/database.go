package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	// open database
	db, err := sql.Open("sqlite3", "workouts.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// create workouts table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS workouts (
			workout_id INTEGER PRIMARY KEY AUTOINCREMENT,
			workout_name TEXT NOT NULL,
			workout_description TEXT
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create sections table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sections (
			section_id INTEGER PRIMARY KEY AUTOINCREMENT,
			section_name TEXT NOT NULL
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create movements table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS movements (
			movement_id INTEGER PRIMARY KEY AUTOINCREMENT,
			movement_name TEXT NOT NULL,
			movement_description TEXT
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create workouts_movements table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS workouts_movements (
			workouts_movements_id INTEGER PRIMARY KEY AUTOINCREMENT,
			workout_id INTEGER NOT NULL,
			section_id INTEGER NOT NULL,
			section_order INTEGER NOT NULL,
			movement_id INTEGER NOT NULL,
			movement_order INTEGER NOT NULL,
			movement_duration INTEGER NOT NULL,
			FOREIGN KEY (workout_id) REFERENCES workouts(workout_id) ON DELETE CASCADE,
			FOREIGN KEY (section_id) REFERENCES sections(section_id) ON DELETE CASCADE,
			FOREIGN KEY (movement_id) REFERENCES movements(movement_id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// insert some workouts
	stmt, err := db.Prepare(`
		INSERT INTO workouts (workout_name, workout_description) 
        VALUES (?, ?)
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("Full Body", "A comprehensive full body workout.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*_, err = stmt.Exec("Upper Body Strength", "Focuses on upper body strength and endurance.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt.Exec("Cardio", "High intensity interval training.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

	// insert some sections
	stmt2, err := db.Prepare(`
		INSERT INTO sections (section_name)
        VALUES (?)
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt2.Close()

	_, err = stmt2.Exec("Warm-Up")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt2.Exec("Strength Training")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*_, err = stmt2.Exec("Conditioning")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt2.Exec("Speed and Agility")
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/
	_, err = stmt2.Exec("Cooldown")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// insert some movements
	stmt3, err := db.Prepare(`
		INSERT INTO movements (movement_name, movement_description)
        VALUES (?, ?)
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt3.Close()

	/*_, err = stmt3.Exec("Squat", "A lower body strength exercise focusing on thighs and glutes.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Deadlift", "A weightlifting movement that targets the back and legs.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/
	_, err = stmt3.Exec("Plank", "A core strengthening exercise involving a static hold.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Pushup", "A chest exercise.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Treadmill", "Run on the treadmill.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*_, err = stmt3.Exec("Rower Machine", "Row at a moderate pace.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Sit-ups", "Do standard sit-ups.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Lunges", "Alternating lunges.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Step-Ups", "Fast-paced lower body mobility.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt3.Exec("Ladder In-and-Outs", "Fast footwork.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

	// insert some movements_movements
	stmt4, err := db.Prepare(`
		INSERT INTO workouts_movements (workout_id, section_id, section_order, movement_id, movement_order, movement_duration)
        VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(1, 1, 1, 3, 1, 600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt4.Exec(1, 2, 1, 2, 1, 600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt4.Exec(1, 3, 1, 1, 1, 600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func GetAllWorkouts() {

	// open database
	db, err := sql.Open("sqlite3", "workouts.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
        	workouts.workout_id, workouts.workout_name, workouts.workout_description,
            sections.section_id, sections.section_name, workouts_movements.section_order,
            movements.movement_id, movements.movement_name, movements.movement_description, workouts_movements.movement_order, workouts_movements.movement_duration
		FROM workouts
		LEFT JOIN workouts_movements ON workouts.workout_id = workouts_movements.workout_id
		LEFT JOIN sections ON sections.section_id = workouts_movements.section_id
		LEFT JOIN movements ON movements.movement_id = workouts_movements.movement_id
		ORDER BY workouts.workout_id, workouts_movements.section_order, workouts_movements.movement_order
	`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {

		fmt.Println(rows)

		var workoutID int
		var workoutName string
		var workoutDescription string
		
		var sectionID int
		var sectionName string
		var sectionOrder int

		var movementID int
		var movementName string
		var movementDescription string
		var movementOrder int
		var movementDuration int



		err := rows.Scan(
			&workoutID,
			&workoutName,
			&workoutDescription,
			&sectionID,
			&sectionName,
			&sectionOrder,
			&movementID,
			&movementName,
			&movementDescription,
			&movementOrder,
			&movementDuration,
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(workoutName)

	}

}