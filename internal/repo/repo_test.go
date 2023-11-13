package repo

import (
    "context"
    "database/sql"
    "errors"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/execut/ozon-keyword-api/internal/model"
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
    expectedId := uint64(123)
    expectedName := "test name"
    t.Run("GetExistKeyword_success", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        rows := sqlmock.NewRows([]string{"id", "name"}).
            AddRow(expectedId, expectedName)
        mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM keywords WHERE id=$1 AND !removed")).
            WithArgs(expectedId).
            WillReturnRows(rows)

        keyword, err := sut.Get(context.Background(), expectedId)

        assert.NilError(t, err)
        assert.Equal(t, expectedId, keyword.ID)
        assert.Equal(t, expectedName, keyword.Name)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
    })
    t.Run("RemoveExistKeyword_success", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        mock.ExpectExec(regexp.QuoteMeta("UPDATE keywords SET removed = $1 WHERE id = $2 AND !removed")).
            WithArgs(true, expectedId).
            WillReturnResult(sqlmock.NewResult(1, 1))

        err := sut.Remove(context.Background(), expectedId)

        assert.NilError(t, err)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
    })
    t.Run("RemoveNonExistentKeyword_failed", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        mock.ExpectExec("UPDATE").
            WillReturnResult(sqlmock.NewResult(1, 0))

        err := sut.Remove(context.Background(), expectedId)

        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }

        assert.ErrorIs(t, ErrKeywordNotFound, err)
    })

    t.Run("AddKeyword_success", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        mock.ExpectExec("INSERT INTO keywords").
            WithArgs(expectedName).
            WillReturnResult(sqlmock.NewResult(int64(expectedId), 1))

        keyword := &model.Keyword{
            Name: expectedName,
        }
        actualId, err := sut.Add(context.Background(), keyword)

        assert.NilError(t, err)
        assert.Equal(t, expectedId, keyword.ID)
        assert.Equal(t, expectedId, actualId)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
    })

    t.Run("AddKeywordWithoutName_fail", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        expectedError := errors.New("empty name")
        mock.ExpectExec("INSERT INTO keywords").
            WillReturnResult(sqlmock.NewErrorResult(expectedError))

        keyword := &model.Keyword{}
        _, err := sut.Add(context.Background(), keyword)

        assert.ErrorIs(t, err, expectedError)
    })

    t.Run("ListWithOffset0_returnTwoKeywords", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM keywords WHERE !removed LIMIT 50 OFFSET 100`)).
            WillReturnRows(sqlmock.
                NewRows([]string{"id", "name"}).
                AddRow(expectedId, expectedName).
                AddRow(456, "test name2"))

        keywords, err := sut.List(context.Background(), 50, 100)

        assert.NilError(t, err)
        assert.Equal(t, 2, len(keywords))
        keyword := keywords[0]
        assert.Equal(t, expectedId, keyword.ID)
        assert.Equal(t, expectedName, keyword.Name)
    })

    t.Run("ListWithError_returnError", func(t *testing.T) {
        t.Parallel()
        sut, mock, db := newSut(t)
        defer db.Close()
        expectedError := errors.New("failed select")
        mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
            WillReturnError(expectedError)

        _, err := sut.List(context.Background(), 50, 100)

        assert.ErrorIs(t, err, expectedError)
    })
}
