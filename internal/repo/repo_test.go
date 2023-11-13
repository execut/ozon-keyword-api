package repo

import (
    "context"
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/jmoiron/sqlx"
    "gotest.tools/v3/assert"
    "regexp"
    "testing"
)

func newSut(t *testing.T) (Repo, sqlmock.Sqlmock, *sql.DB) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }

    sqlxDB := sqlx.NewDb(db, "sqlmock")

    repo := NewRepo(sqlxDB, 100)

    return repo, mock, db
}

func Test_repo(t *testing.T) {
    t.Parallel()
    t.Run("GetExistKeyword_success", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        expectedName := "test name"
        expectedId := uint64(1)
        rows := sqlmock.NewRows([]string{"id", "name"}).
            AddRow(expectedId, expectedName)
        mock.ExpectQuery("SELECT id, name FROM keywords").WillReturnRows(rows)

        keyword, err := sut.Get(context.Background(), expectedId)

        assert.NilError(t, err)
        assert.Equal(t, expectedId, keyword.ID)
        assert.Equal(t, expectedName, keyword.Name)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
    })
    t.Run("DeleteExistKeyword_success", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        expectedId := uint64(1)
        mock.ExpectExec(regexp.QuoteMeta("DELETE FROM keywords")).WithArgs(expectedId).WillReturnResult(sqlmock.NewResult(1, 1))

        err := sut.Remove(context.Background(), expectedId)

        assert.NilError(t, err)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
    })
    t.Run("DeleteNonExistentKeyword_failed", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        expectedId := uint64(1)
        mock.ExpectExec(regexp.QuoteMeta("DELETE FROM keywords")).WithArgs(expectedId).WillReturnResult(sqlmock.NewResult(1, 0))

        err := sut.Remove(context.Background(), expectedId)

        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }

        assert.ErrorIs(t, ErrKeywordNotFound, err)
    })
}
