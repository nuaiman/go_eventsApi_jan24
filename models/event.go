package models

import (
	"time"

	"example.com/events_api/db"
)

type Event struct {
	Id          int64
	UserId      int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (userId, name, description, location, dateTime)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&e.UserId, &e.Name, &e.Description, &e.Location, &e.DateTime)
	return err
}

func GetEvents() (*[]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.UserId, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return &events, nil
}

func GetEvent(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(eventId)
	var event Event
	err = row.Scan(&event.Id, &event.UserId, &event.Name, &event.Description, &event.Location, &event.DateTime)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET userId = ? , name = ? , description = ? , location = ? , dateTime = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&e.UserId, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.Id)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&e.Id)
	return err
}

func (e *Event) Register(userId int64) error {
	query := "INSERT INTO registrations (userId, eventId) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, &e.Id)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE userId = ? AND eventId = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, &e.Id)
	return err
}
