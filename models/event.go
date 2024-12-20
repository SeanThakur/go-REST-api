package models

import (
	"seanThakur/go-restapi/db"
	"time"
)

type Event struct {
	Id          int64
	UserId      int64
	Description string    `binding:"required"`
	Name        string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, date_time, location, user_id) VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvent() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.DateTime, &event.Location, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.DateTime, &event.Location, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
		UPDATE events
		SET
			name = ?,
			description = ?,
			location = ?,
			date_time = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)

	return err

}

func (event *Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id)

	return err
}

func (event Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations(user_id, event_id) VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, event.Id)

	return err
}
