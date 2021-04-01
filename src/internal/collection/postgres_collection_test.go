package collection_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"

	col "github.com/pachyderm/pachyderm/v2/src/internal/collection"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"
	"github.com/pachyderm/pachyderm/v2/src/internal/testutil"
)

func TestPostgresCollections(suite *testing.T) {
	suite.Parallel()
	postgres := testutil.NewPostgresDeployment(suite)

	newCollection := func(t *testing.T) (col.ReadOnlyCollection, WriteCallback) {
		db, listener := postgres.NewDatabase(t)
		testCol, err := col.NewPostgresCollection(context.Background(), db, listener, &col.TestItem{}, []*col.Index{TestSecondaryIndex})
		require.NoError(t, err)

		writeCallback := func(f func(col.ReadWriteCollection) error) error {
			return col.NewSQLTx(context.Background(), db, func(tx *sqlx.Tx) error {
				return f(testCol.ReadWrite(tx))
			})
		}

		return testCol.ReadOnly(context.Background()), writeCallback
	}

	collectionTests(suite, newCollection)

	// TODO: postgres-specific collection tests - With()
}
