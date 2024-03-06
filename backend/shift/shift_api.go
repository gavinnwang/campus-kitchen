package shift

import (
	"backend/auth"
	"backend/db"
	"backend/endpoint"
	"time"

	server_error "backend/error"
	"backend/validator"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// API represents a struct for the user API
type API struct {
	q         *db.Queries
	validator *validator.Validate
}

// NewAPI initializes an API struct
func NewAPI(conn *db.DB, v *validator.Validate) API {
	q := db.New(conn)
	return API{
		q:         q,
		validator: v,
	}
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

// package db

// import (
// 	"context"
// 	"database/sql"
// 	"time"
// )

// const createShift = `-- name: CreateShift :execresult
// INSERT INTO shifts (start_time, end_time, type)
// VALUES (?, ?, ?)
// `

// type CreateShiftParams struct {
// 	StartTime time.Time
// 	EndTime   time.Time
// 	Type      ShiftsType
// }

// func (q *Queries) CreateShift(ctx context.Context, arg CreateShiftParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, createShift, arg.StartTime, arg.EndTime, arg.Type)
// }

// const createShiftVolunteer = `-- name: CreateShiftVolunteer :execresult
// INSERT INTO shift_volunteers (user_id, shift_id)
// VALUES (?, ?)
// `

// type CreateShiftVolunteerParams struct {
// 	UserID  int32
// 	ShiftID int32
// }

// func (q *Queries) CreateShiftVolunteer(ctx context.Context, arg CreateShiftVolunteerParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, createShiftVolunteer, arg.UserID, arg.ShiftID)
// }

// const createUser = `-- name: CreateUser :execresult
// INSERT INTO users (name, email, type)
// VALUES (?, ?, ?)
// `

// type CreateUserParams struct {
// 	Name  string
// 	Email string
// 	Type  UsersType
// }

// func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, createUser, arg.Name, arg.Email, arg.Type)
// }

// const deleteShift = `-- name: DeleteShift :execresult
// DELETE FROM shifts
// WHERE shifts.id = ?
// `

// func (q *Queries) DeleteShift(ctx context.Context, id int32) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, deleteShift, id)
// }

// const deleteShiftVolunteer = `-- name: DeleteShiftVolunteer :execresult
// DELETE FROM shift_volunteers
// WHERE shift_volunteers.user_id = ? AND shift_volunteers.shift_id = ?
// `

// type DeleteShiftVolunteerParams struct {
// 	UserID  int32
// 	ShiftID int32
// }

// func (q *Queries) DeleteShiftVolunteer(ctx context.Context, arg DeleteShiftVolunteerParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, deleteShiftVolunteer, arg.UserID, arg.ShiftID)
// }

// const deleteUser = `-- name: DeleteUser :execresult
// DELETE FROM users
// WHERE users.id = ?
// `

// func (q *Queries) DeleteUser(ctx context.Context, id int32) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, deleteUser, id)
// }

// const getShift = `-- name: GetShift :one
// SELECT id, start_time, end_time, type FROM shifts
// WHERE shifts.id = ?
// `

// func (q *Queries) GetShift(ctx context.Context, id int32) (Shift, error) {
// 	row := q.db.QueryRowContext(ctx, getShift, id)
// 	var i Shift
// 	err := row.Scan(
// 		&i.ID,
// 		&i.StartTime,
// 		&i.EndTime,
// 		&i.Type,
// 	)
// 	return i, err
// }

// const getShiftLeadersForShift = `-- name: GetShiftLeadersForShift :many
// SELECT users.id, users.name, users.email, users.phone, users.type
// FROM users
// JOIN shift_leaders ON users.id = shift_leaders.user_id
// WHERE shift_leaders.shift_id = ?
// `

// func (q *Queries) GetShiftLeadersForShift(ctx context.Context, shiftID int32) ([]User, error) {
// 	rows, err := q.db.QueryContext(ctx, getShiftLeadersForShift, shiftID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []User
// 	for rows.Next() {
// 		var i User
// 		if err := rows.Scan(
// 			&i.ID,
// 			&i.Name,
// 			&i.Email,
// 			&i.Phone,
// 			&i.Type,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const getShiftVolunteersForShift = `-- name: GetShiftVolunteersForShift :many
// SELECT users.id, users.name, users.email, users.phone, users.type
// FROM users
// JOIN shift_volunteers ON users.id = shift_volunteers.user_id
// WHERE shift_volunteers.shift_id = ?
// `

// func (q *Queries) GetShiftVolunteersForShift(ctx context.Context, shiftID int32) ([]User, error) {
// 	rows, err := q.db.QueryContext(ctx, getShiftVolunteersForShift, shiftID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []User
// 	for rows.Next() {
// 		var i User
// 		if err := rows.Scan(
// 			&i.ID,
// 			&i.Name,
// 			&i.Email,
// 			&i.Phone,
// 			&i.Type,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const getShifts = `-- name: GetShifts :many
// SELECT id, start_time, end_time, type FROM shifts
// `

// func (q *Queries) GetShifts(ctx context.Context) ([]Shift, error) {
// 	rows, err := q.db.QueryContext(ctx, getShifts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []Shift
// 	for rows.Next() {
// 		var i Shift
// 		if err := rows.Scan(
// 			&i.ID,
// 			&i.StartTime,
// 			&i.EndTime,
// 			&i.Type,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const getUser = `-- name: GetUser :one

// SELECT id, name, email, phone, type FROM users
// WHERE users.id = ?
// `

// CREATE TABLE IF NOT EXISTS users (
//
//	id INT AUTO_INCREMENT PRIMARY KEY,
//	name VARCHAR(255) NOT NULL,
//	email VARCHAR(255) UNIQUE NOT NULL,
//	phone VARCHAR(255),
//	type ENUM('volunteer', 'shift_lead') NOT NULL
//
// );
// CREATE TABLE IF NOT EXISTS shifts (
//
//	id INT AUTO_INCREMENT PRIMARY KEY,
//	start_time DATETIME NOT NULL,
//	end_time DATETIME NOT NULL,
//	type ENUM('recovery', 'resourcing', 'meal_prep') NOT NULL
//
// );
// CREATE TABLE IF NOT EXISTS shift_leaders (
//
//	id INT AUTO_INCREMENT PRIMARY KEY,
//	user_id INT NOT NULL,
//	shift_id INT NOT NULL,
//	FOREIGN KEY (user_id) REFERENCES users(id),
//	FOREIGN KEY (shift_id) REFERENCES shifts(id)
//
// );
// CREATE TABLE IF NOT EXISTS shift_volunteers (
//
//	id INT AUTO_INCREMENT PRIMARY KEY,
//	user_id INT NOT NULL,
//	shift_id INT NOT NULL,
//	FOREIGN KEY (user_id) REFERENCES users(id)
//
// );
// CREATE TABLE IF NOT EXISTS volunteer_completed_shifts (
//
//	id INT AUTO_INCREMENT PRIMARY KEY,
//	user_id INT NOT NULL,
//	shift_id INT NOT NULL,
//	FOREIGN KEY (user_id) REFERENCES users(id),
//	FOREIGN KEY (shift_id) REFERENCES shifts(id)
//
// );
// func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
// 	row := q.db.QueryRowContext(ctx, getUser, id)
// 	var i User
// 	err := row.Scan(
// 		&i.ID,
// 		&i.Name,
// 		&i.Email,
// 		&i.Phone,
// 		&i.Type,
// 	)
// 	return i, err
// }

// const getUserByEmail = `-- name: GetUserByEmail :one
// SELECT id, name, email, phone, type FROM users
// WHERE users.email = ?
// `

// func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
// 	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
// 	var i User
// 	err := row.Scan(
// 		&i.ID,
// 		&i.Name,
// 		&i.Email,
// 		&i.Phone,
// 		&i.Type,
// 	)
// 	return i, err
// }

// const getUsers = `-- name: GetUsers :many
// SELECT id, name, email, phone, type FROM users
// `

// func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
// 	rows, err := q.db.QueryContext(ctx, getUsers)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []User
// 	for rows.Next() {
// 		var i User
// 		if err := rows.Scan(
// 			&i.ID,
// 			&i.Name,
// 			&i.Email,
// 			&i.Phone,
// 			&i.Type,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const updateShift = `-- name: UpdateShift :execresult
// UPDATE shifts
// SET start_time = ?, end_time = ?, type = ?
// WHERE shifts.id = ?
// `

// type UpdateShiftParams struct {
// 	StartTime time.Time
// 	EndTime   time.Time
// 	Type      ShiftsType
// 	ID        int32
// }

// func (q *Queries) UpdateShift(ctx context.Context, arg UpdateShiftParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, updateShift,
// 		arg.StartTime,
// 		arg.EndTime,
// 		arg.Type,
// 		arg.ID,
// 	)
// }

// const updateUserName = `-- name: UpdateUserName :execresult
// UPDATE users
// SET name = ?
// WHERE users.id = ?
// `

// type UpdateUserNameParams struct {
// 	Name string
// 	ID   int32
// }

// func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (sql.Result, error) {
// 	return q.db.ExecContext(ctx, updateUserName, arg.Name, arg.ID)
// }

// func (api *API) HandleGetUser(w http.ResponseWriter, r *http.Request) {

// 	// // var input GetUserInput
// 	// // err := endpoint.DecodeAndValidateJson(w, r, api.validator, &input)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	id := chi.URLParam(r, "id")

// 	result, err := strconv.ParseInt(id, 10, 32)
// 	var id1 int32

// 	if err != nil {
// 		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	} else {
// 		id1 = int32(result)
// 	}

// 	user, err := api.q.GetUser(r.Context(), id1)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			endpoint.WriteWithError(w, http.StatusNotFound, server_error.ErrUserNotFound)
// 			return
// 		}
// 		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	endpoint.WriteWithStatus(w, http.StatusOK, user)
// }

// func (api *API) HandleGetUsers(w http.ResponseWriter, r *http.Request) {

// 	user, err := api.q.GetUsers(r.Context())
// 	if err != nil {
// 		log.Printf("Error get user from database: %v", err)
// 		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	endpoint.WriteWithStatus(w, http.StatusOK, user)
// }

// func (api *API) HandleGetMe(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	userID, err := auth.UserIDFromContext(ctx)
// 	if err != nil {
// 		endpoint.WriteWithError(w, http.StatusUnauthorized, server_error.ErrUnauthorized)
// 		return
// 	}

// 	user, err := api.q.GetUser(r.Context(), int32(userID))
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			endpoint.WriteWithError(w, http.StatusNotFound, server_error.ErrUserNotFound)
// 			return
// 		}
// 		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	endpoint.WriteWithStatus(w, http.StatusOK, user)
// }

// type UpdateUserNameRequest struct {
// 	Name string `json:"name" validate:"required"`
// }

// func (api *API) HandleUpdateUserName(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	userID, err := auth.UserIDFromContext(ctx)
// 	if err != nil {
// 		endpoint.WriteWithError(w, http.StatusUnauthorized, server_error.ErrUnauthorized)
// 		return
// 	}

// 	var input UpdateUserNameRequest

// 	err = endpoint.DecodeAndValidateJson(w, r, api.validator, &input)
// 	if err != nil {
// 		return
// 	}

// 	_, err = api.q.UpdateUserName(r.Context(), db.UpdateUserNameParams{
// 		ID:   int32(userID),
// 		Name: input.Name,
// 	})

// 	if err != nil {
// 		log.Printf("Error update user name: %v", err)
// 		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	endpoint.WriteWithStatus(w, http.StatusOK, "update success")
// }

// func (api *API) RegisterHandlers(r chi.Router, auth_guard func(http.Handler) http.Handler) {
// 	r.Route("/users", func(r chi.Router) {
// 		r.Group(func(r chi.Router) {
// 			r.Use(auth_guard)
// 			r.Get("/me", api.HandleGetMe)
// 			r.Post("/me/name", api.HandleUpdateUserName)
// 		})
// 		r.Get("/", api.HandleGetUsers)
// 		r.Get("/{id}", api.HandleGetUser)
// 	})
// }

type CreateShiftRequest struct {
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
	Type      string    `json:"type" validate:"required"`
}

func (api *API) HandleCreateShift(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated and if user is shift leader
	ctx := r.Context()
	userID, err := auth.UserIDFromContext(ctx)
	if err != nil {
		endpoint.WriteWithError(w, http.StatusUnauthorized, server_error.ErrUnauthorized)
		return
	}

	// get user using user id
	user, err := api.q.GetUser(ctx, int32(userID))
	if err != nil {
		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// check if user is shift leader
	if user.Type != db.UsersTypeShiftLead {
		endpoint.WriteWithError(w, http.StatusUnauthorized, server_error.ErrUnauthorized)
		return
	}

	// decode and validate input
	var input CreateShiftRequest
	err = endpoint.DecodeAndValidateJson(w, r, api.validator, &input)
	if err != nil {
		return
	}

	// create shift
	_, err = api.q.CreateShift(ctx, db.CreateShiftParams{
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Type:      db.ShiftsType(input.Type),
	})

	if err != nil {
		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	endpoint.WriteWithStatus(w, http.StatusCreated, "shift created")
}

func (api *API) HandleGetShifts(w http.ResponseWriter, r *http.Request) {

	shifts, err := api.q.GetShifts(r.Context())
	if err != nil {
		endpoint.WriteWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	endpoint.WriteWithStatus(w, http.StatusOK, shifts)
}

func (api *API) RegisterHandlers(r chi.Router, auth_guard func(http.Handler) http.Handler) {
	r.Route("/shifts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth_guard)
			r.Post("/", api.HandleCreateShift)
		})
		r.Get("/", api.HandleGetShifts)
	})
}
