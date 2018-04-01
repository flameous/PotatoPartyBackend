package db

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/flameous/PotatoPartyBackend/types"
	"log"
	"math/rand"
)

type Storage struct {
	db *sql.DB
}

func ConnectToDatabase(host string) (*Storage, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://potato:potato@%s:5432/potato?sslmode=disable", host))
	if err != nil {
		return nil, err
	}
	s := &Storage{db}
	s.randomizeData()
	return s, db.Ping()
}

func (s *Storage) randomizeData() {
	var needUpd bool
	err := s.db.QueryRow(`SELECT is_updated FROM debug`).Scan(&needUpd)
	log.Println(err)
	if needUpd {
		log.Println("db is ready")
		return
	}

	var maxUserID int64
	err = s.db.QueryRow(`SELECT max(id) FROM users`).Scan(&maxUserID)
	log.Println(err, maxUserID)

	var maxInterestsID int64
	err = s.db.QueryRow(`SELECT max(id) FROM interests`).Scan(&maxInterestsID)
	log.Println(err, maxInterestsID)

	// add random users and interests to all events
	var maxEventsID int64
	err = s.db.QueryRow(`SELECT max(id) FROM events`).Scan(&maxEventsID)
	log.Println(err, maxEventsID)

	for eventID := int64(1); eventID <= maxEventsID; eventID++ {
		for uid := int64(2); uid <= maxUserID-rand.Int63n(15); uid++ {
			s.AttendToEvent(eventID, uid)
		}

		for uid := int64(0); uid < int64(1)+rand.Int63n(5); uid++ {
			err = s.createEventInterests(eventID, rand.Int63n(maxInterestsID)+1)
		}
	}

	_, err = s.db.Exec(`UPDATE debug SET is_updated = TRUE`)
	log.Println(err)
}

func (s *Storage) UpsertUser(u *types.User) (int64, error) {
	panic("implement me, please!")
}

func (s *Storage) GetUserByID(id int64) (*types.User, error) {
	u, err := s.readUserBy(&id, nil, nil)
	return u, err
}

func (s *Storage) GetUserByNickname(nickname string) (*types.User, error) {
	u, err := s.readUserBy(nil, &nickname, nil)
	return u, err
}

func (s *Storage) GetUserByToken(token string) (*types.User, error) {
	u, err := s.readUserBy(nil, nil, &token)
	return u, err
}

func (s *Storage) readUserBy(id *int64, nickname, token *string) (*types.User, error) {
	var u types.User
	var val interface{}

	var q string
	if id != nil {
		q = `SELECT id, nickname, password, email, last_seen, lat, lon, picture_url FROM users WHERE id = $1`
		val = *id
	} else if nickname != nil {
		q = `SELECT id, nickname, password, email, last_seen, lat, lon, picture_url FROM users WHERE nickname = $1`
		val = *nickname
	} else if token != nil {
		q = `SELECT id, nickname, password, email, last_seen, lat, lon, picture_url FROM users WHERE token = $1`
		val = *token
	} else {
		return nil, errors.New("no id or nickname provided")
	}

	err := s.db.QueryRow(q, val).Scan(&u.ID, &u.Nickname, &u.Password, &u.Email, &u.LastSeen, &u.Lat, &u.Lon, &u.PictureURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "failed to get user. id = %d", id)
	}

	i, err := s.GetUserInterests(u.ID)
	if err != nil {
		return nil, err
	}
	u.Interests = i
	return &u, nil
}

func (s *Storage) GetUserInterests(id int64) (*types.AllInterests, error) {
	const q = `SELECT id, category_id, name FROM interests WHERE id = ANY(SELECT interest_id FROM user_interests WHERE user_id = $1)`

	rows, err := s.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var is []*types.Interests
	for rows.Next() {
		var i types.Interests
		err = rows.Scan(&i.ID, &i.CategoryID, &i.Name)
		if err != nil {
			return nil, err
		}
		is = append(is, &i)
	}

	return types.Categorize(is), nil
}

func (s *Storage) UpdateUserToken(u *types.User) error {
	_, err := s.db.Exec(`UPDATE  users SET token = $1 WHERE id = $2`, u.Token, u.ID)
	return err
}

func (s *Storage) GetAllUsers() ([]*types.User, error) {
	const q = `SELECT id, nickname, email, last_seen, lat, lon, picture_url FROM users`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	var users []*types.User
	for rows.Next() {
		var u types.User
		err = rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.LastSeen, &u.Lat, &u.Lon, &u.PictureURL)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (s *Storage) GetAllEvents(u *types.User) (*types.EventsList, error) {
	const q = `SELECT id, name, description, date, owner_id, lat, lon, is_private FROM events 
	WHERE is_private = FALSE OR id IN (SELECT event_id FROM event_attendees WHERE user_id = $1);`

	rows, err := s.db.Query(q, u.ID)
	if err != nil {
		return nil, err
	}

	var events []*types.Event
	for rows.Next() {
		var e types.Event
		e.Owner = &types.User{}
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.Date, &e.Owner.ID, &e.Lat, &e.Lon, &e.IsPrivate)
		if err != nil {
			return nil, err
		}
		events = append(events, &e)
	}

	m := u.Interests.ToMapByID()
	el := types.NewEventsList()
	for _, e := range events {
		err = s.GetExtendedData(e, m, u.ID)
		if err != nil {
			return nil, err
		}

		if e.Owner.ID == u.ID {
			el.OwnEvents = append(el.OwnEvents, e)
			continue
		}

		if e.IsPrivate {
			el.Private = append(el.Private, e)
			continue
		}
		if e.IsUserAttended {
			el.Attended = append(el.Attended, e)
			continue
		}

		if len(e.SameInterests) == 0 {
			continue
		}
		el.Suggested = append(el.Suggested, e)
	}
	return el, nil
}

func (s *Storage) GetExtendedData(event *types.Event, uis map[int64]bool, uid int64) error {
	const q = `SELECT id, category_id, name FROM interests 
		WHERE id = ANY(SELECT interest_id FROM event_interests WHERE event_id = $1)`

	rows, err := s.db.Query(q, event.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var is []*types.Interests
	same := make([]*types.Interests, 0)
	for rows.Next() {
		var i types.Interests
		err = rows.Scan(&i.ID, &i.CategoryID, &i.Name)
		if err != nil {
			return err
		}
		is = append(is, &i)
		if uis[i.ID] {
			same = append(same, &i)
		}
	}

	// if event is NOT private and we haven't same interests, we'll drop it later
	if len(same) == 0 && !event.IsPrivate && event.Owner.ID != uid {
		return nil
	}

	event.Owner, err = s.GetUserByID(event.Owner.ID)
	if err != nil {
		return errors.Wrap(err, "get event owner user")
	}

	event.Interests = types.Categorize(is)
	event.SameInterests = same
	s.getEventAttendees(event, uid)
	return nil
}

func (s *Storage) getEventAttendees(event *types.Event, uid int64) error {
	const q = `SELECT id, nickname, email, last_seen, lat, lon, picture_url 
		FROM users WHERE id = ANY(SELECT user_id FROM event_attendees WHERE event_id = $1)`

	rows, err := s.db.Query(q, event.ID)
	for rows.Next() {
		var u types.User
		err = rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.LastSeen, &u.Lat, &u.Lon, &u.PictureURL)
		if err != nil {
			return errors.Wrap(err, "failed to get event attendee")
		}
		if u.ID == uid {
			event.IsUserAttended = true
		}
		event.Attendees = append(event.Attendees, &u)
	}
	return nil
}

func (s *Storage) CreateNewEvent(event *types.Event) (int64, error) {
	var id int64
	err := s.db.QueryRow(
		`INSERT INTO 
			events (name, description, date, lat, lon, is_private, owner_id) 
		VALUES
			($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		event.Name, event.Description, event.Date, event.Lat, event.Lon, event.IsPrivate, event.Owner.ID).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	for _, v := range event.Interests.ToArray() {
		if err = s.createEventInterests(id, v.ID); err != nil {
			return 0, errors.Wrap(err, "insert event's interests")
		}
	}

	// fixme: to escape null vals in JSON
	event.Interests.Validate()
	for _, v := range event.Attendees {
		if err = s.createEventAttendees(id, v.ID); err != nil {
			return 0, errors.Wrap(err, "insert event's interests")
		}
	}
	return id, nil
}

func (s *Storage) createEventInterests(eventID, interestID int64) error {
	_, err := s.db.Exec(`INSERT INTO event_interests (event_id, interest_id)
 	VALUES ($1, $2) ON CONFLICT DO NOTHING`, eventID, interestID)
	return err
}

func (s *Storage) createEventAttendees(eventID int64, userID int64) error {
	_, err := s.db.Exec(`INSERT INTO event_attendees (event_id, user_id)
 	VALUES ($1, $2) ON CONFLICT DO NOTHING`, eventID, userID)
	return err
}

func (s *Storage) GetEventByID(id int64, u *types.User) (*types.Event, error) {
	const q = `SELECT id, name, description, date, owner_id, lat, lon, is_private FROM events WHERE id = $1`
	var e types.Event
	e.Owner = &types.User{}
	err := s.db.QueryRow(q, id).Scan(&e.ID, &e.Name, &e.Description, &e.Date, &e.Owner.ID, &e.Lat, &e.Lon, &e.IsPrivate)
	if err != nil {
		log.Println("get event by id, id = ", id, err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if err = s.GetExtendedData(&e, u.Interests.ToMapByID(), u.ID); err != nil {
		return nil, err
	}
	return &e, nil
}

func (s *Storage) AttendToEvent(eventID, userID int64) error {
	_, err := s.db.Exec(`INSERT INTO event_attendees VALUES ($1, $2) ON CONFLICT DO NOTHING`, eventID, userID)
	return err
}

func (s *Storage) RefuseAttendance(eventID, userID int64) error {
	_, err := s.db.Exec(`DELETE FROM event_attendees WHERE event_id = $1 AND user_id = $2;`, eventID, userID)
	return err
}
