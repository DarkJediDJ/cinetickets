package session

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/darkjedidj/cinema-service/internal"
)

var session = &Resource{
	ID:        15,
	Hall_id:   0,
	Movie_id:  0,
	Starts_at: "13:25",
	VIP:       true,
	Name:      "Matrix",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	testCreateCases := []struct {
		name           string
		expectedError  error
		expectedResult internal.Identifiable
		prepare        func(sqlm2 sqlmock.Sqlmock)
		object         internal.Identifiable
	}{
		{
			name:           "failed, database error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("INSERT INTO sessions (.*)").
					WillReturnError(internal.ErrInternalFailure)
			},
			object: session,
		},
		{
			name:           "success",
			expectedError:  nil,
			expectedResult: session,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("INSERT INTO sessions (.*)").
					WillReturnRows(sqlm2.
						NewRows([]string{"id"}).
						AddRow(session.ID))
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id WHERE sessions.id = \\$1").
					WithArgs(session.ID).
					WillReturnRows(sqlm2.
						NewRows([]string{"id", "vip", "name", "starts_at"}).
						AddRow(session.ID, session.VIP, session.Name, session.Starts_at))
			},
			object: session,
		},
		{
			name:           "failed, retrieve error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("INSERT INTO sessions (.*)").
					WillReturnError(fmt.Errorf("unable to retrieve Resource"))
			},
			object: session,
		},
		{
			name:           "failed, assertion error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec(regexp.QuoteMeta("INSERT INTO sessions (.*)")).
					WillReturnError(internal.ErrInternalFailure)
			},
			object: nil,
		},
	}

	for _, tc := range testCreateCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			repo := &Repository{DB: db, Log: logger}
			ctx := context.Background()

			tc.prepare(mock)
			res, err := repo.Create(tc.object, ctx)

			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestRetrieve(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	testRetrieveCases := []struct {
		name           string
		expectedError  error
		expectedResult internal.Identifiable
		prepare        func(sqlm2 sqlmock.Sqlmock)
	}{
		{
			name:           "success",
			expectedError:  nil,
			expectedResult: session,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id WHERE sessions.id = \\$1").
					WithArgs(session.ID).
					WillReturnRows(sqlm2.
						NewRows([]string{"id", "vip", "name", "starts_at"}).
						AddRow(session.ID, session.VIP, session.Name, session.Starts_at))
			},
		},
		{
			name:           "failed, database error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id WHERE sessions.id = \\$1").
					WillReturnError(fmt.Errorf("unable to perform your request, please try again later"))
			},
		},
		{
			name:           "failed, sql no rows error",
			expectedError:  nil,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id WHERE sessions.id = \\$1").
					WillReturnRows(sqlm2.
						NewRows(nil))
			},
		},
	}

	for _, tc := range testRetrieveCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			repo := &Repository{DB: db, Log: logger}
			ctx := context.Background()

			tc.prepare(mock)
			res, err := repo.Retrieve(session.ID, ctx)

			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestRetrieveAll(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	testRetrieveAllCases := []struct {
		name           string
		expectedError  error
		expectedResult []internal.Identifiable
		prepare        func(sqlm2 sqlmock.Sqlmock)
	}{
		{
			name:           "success",
			expectedError:  nil,
			expectedResult: []internal.Identifiable{session},
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id").
					WillReturnRows(sqlm2.
						NewRows([]string{"id", "vip", "name", "starts_at"}).
						AddRow(session.ID, session.VIP, session.Name, session.Starts_at))
			},
		},
		{
			name:           "failed, database error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id").
					WillReturnError(internal.ErrInternalFailure)
			},
		},
		{
			name:           "failed, sql no rows error",
			expectedError:  nil,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectQuery("SELECT sessions.id, halls.vip, movies.name, starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id JOIN halls ON sessions.hall_id = halls.id").
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, tc := range testRetrieveAllCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			repo := &Repository{DB: db, Log: logger}
			ctx := context.Background()

			tc.prepare(mock)
			res, err := repo.RetrieveAll(ctx)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	testDeleteCases := []struct {
		name           string
		expectedError  error
		expectedResult internal.Identifiable
		prepare        func(sqlm2 sqlmock.Sqlmock)
		id             int64
	}{
		{
			name:           "success",
			expectedError:  nil,
			expectedResult: session,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec("DELETE FROM sessions WHERE id = \\$1").
					WithArgs(session.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			id: int64(session.ID),
		},
		{
			name:           "failed, database error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: nil,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec("DELETE FROM sessions WHERE id = \\$1").
					WithArgs(session.ID).
					WillReturnError(fmt.Errorf("unable to perform your request, please try again later"))
			},
			id: int64(session.ID),
		},
	}

	for _, tc := range testDeleteCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			repo := &Repository{DB: db, Log: logger}
			ctx := context.Background()

			tc.prepare(mock)
			err = repo.Delete(tc.id, ctx)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTimeValid(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	testTimeValidCases := []struct {
		name           string
		expectedError  error
		expectedResult bool
		prepare        func(sqlm2 sqlmock.Sqlmock)
		object         internal.Identifiable
	}{
		{
			name:           "success",
			expectedError:  nil,
			expectedResult: true,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec(regexp.QuoteMeta("SELECT movies.duration, sessions.starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id WHERE ($1, movies.duration) OVERLAPS (sessions.starts_at , movies.duration) AND sessions.hall_id = $2 AND sessions.movie_id = $3")).
					WithArgs(session.Starts_at, session.Hall_id, session.Movie_id).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			object: session,
		},
		{
			name:           "failed, validation error",
			expectedError:  internal.ErrInternalFailure,
			expectedResult: false,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec(regexp.QuoteMeta("SELECT movies.duration, sessions.starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id WHERE ($1, movies.duration) OVERLAPS (sessions.starts_at , movies.duration) AND sessions.hall_id = $2 AND sessions.movie_id = $3")).
					WithArgs(session.Starts_at, session.Hall_id, session.Movie_id).
					WillReturnError(internal.ErrInternalFailure)
			},
			object: session,
		},
		{
			name:           "failed, assertion error",
			expectedError:  internal.ErrValidationFailed,
			expectedResult: false,
			prepare: func(sqlm2 sqlmock.Sqlmock) {
				sqlm2.ExpectExec(regexp.QuoteMeta("SELECT movies.duration, sessions.starts_at FROM sessions JOIN movies ON sessions.movie_id = movies.id WHERE ($1, movies.duration) OVERLAPS (sessions.starts_at , movies.duration) AND sessions.hall_id = $2 AND sessions.movie_id = $3")).
					WillReturnError(internal.ErrValidationFailed)
			},
			object: nil,
		},
	}

	for _, tc := range testTimeValidCases {
		t.Run(tc.name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			if err != nil {
				log.Fatalf("can't initialize zap logger: %v", err)
			}

			defer func() {
				if err := logger.Sync(); err != nil {
					fmt.Println(err)
				}
			}()

			repo := &Repository{DB: db, Log: logger}
			ctx := context.Background()

			tc.prepare(mock)
			valid, err := repo.TimeValid(tc.object, ctx)
			assert.Equal(t, tc.expectedResult, valid)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGID(t *testing.T) {
	res := &Resource{ID: session.ID}
	assert.Equal(t, session.ID, res.GID())
}
