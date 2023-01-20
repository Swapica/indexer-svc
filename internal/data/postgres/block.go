package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/Swapica/order-indexer-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const blockTable = "last_blocks"
const chainCol = "src_chain"

type block struct {
	db    *pgdb.DB
	chain string
}

func NewLastBlock(db *pgdb.DB, chainName string) (data.LastBlock, error) {
	q := block{db: db, chain: chainName}
	if err := q.init(chainName); err != nil {
		return block{}, errors.Wrap(err, "failed to initialize last block storage")
	}
	return q, nil
}

func (q block) init(chain string) error {
	b, err := q.Get()
	if err != nil {
		return errors.Wrap(err, "failed to check block existence")
	}
	if b != nil {
		return nil
	}

	stmt := squirrel.Insert(blockTable).Columns("id", chainCol).Values(0, chain)
	err = q.db.Exec(stmt)
	return errors.Wrap(err, "failed to insert last block")
}

func (q block) Set(id uint64) error {
	stmt := squirrel.Update(blockTable).Set("id", id).Where(squirrel.Eq{chainCol: q.chain})
	err := q.db.Exec(stmt)
	return errors.Wrap(err, "failed to update last block")
}

func (q block) Get() (*uint64, error) {
	var result struct {
		ID uint64 `db:"id"`
	}
	stmt := squirrel.Select("id").From(blockTable).Where(squirrel.Eq{chainCol: q.chain})

	if err := q.db.Get(&result, stmt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to select last block")
	}

	return &result.ID, nil
}
