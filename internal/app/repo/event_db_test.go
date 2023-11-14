package repo

import (
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/execut/ozon-keyword-api/internal/model"
    "github.com/jmoiron/sqlx"
    "gotest.tools/v3/assert"
    "regexp"
    "testing"
)

func newSut(t *testing.T) (EventRepo, sqlmock.Sqlmock, *sql.DB) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }

    sqlxDB := sqlx.NewDb(db, "sqlmock")

    repo := NewEventRepo(sqlxDB)

    return repo, mock, db
}

func TestEventDbRepo_LockUnlock(t *testing.T) {
    t.Parallel()
    sut, mock, db := newSut(t)
    defer db.Close()
    t.Run("Lock_success", func(t *testing.T) {
        mock.ExpectBegin()
        expectedId := uint64(123)
        expectedKeywordId := uint64(345)
        expectedType := model.Created
        expectedSecondId := 555
        mock.ExpectQuery(regexp.QuoteMeta("SELECT id, keyword_id, type FROM keyword_events WHERE status is null ORDER BY id LIMIT 678")).
            WillReturnRows(mock.
                NewRows([]string{"id", "keyword_id", "type"}).
                AddRow(expectedId, expectedKeywordId, expectedType).
                AddRow(expectedSecondId, 3, model.Updated))
        mock.ExpectExec(regexp.QuoteMeta("UPDATE keyword_events SET status = $1 WHERE id IN ($2,$3)")).
            WithArgs(model.Deferred, expectedId, expectedSecondId).
            WillReturnResult(sqlmock.NewResult(1, 2))
        mock.ExpectCommit()

        keywords, err := sut.Lock(678)

        assert.NilError(t, err)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("there were unfulfilled expectations: %s", err)
        }
        assert.Equal(t, 2, len(keywords))
        keyword := keywords[0]
        assert.Equal(t, expectedId, keyword.ID)
        assert.Equal(t, true, keyword.Entity != nil)
        assert.Equal(t, expectedKeywordId, keyword.Entity.ID)
        assert.Equal(t, expectedType, keyword.Type)
    })
}
