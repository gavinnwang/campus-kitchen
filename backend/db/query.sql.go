// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT volunteer_name, volunteer_type, total_hours_volunteered, total_pounds_food_recovered, total_number_meals_given FROM users
WHERE users.id = $1
`

func (q *Queries) GetUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser)
	var i User
	err := row.Scan(
		&i.VolunteerName,
		&i.VolunteerType,
		&i.TotalHoursVolunteered,
		&i.TotalPoundsFoodRecovered,
		&i.TotalNumberMealsGiven,
	)
	return i, err
}
