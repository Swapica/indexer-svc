package postgres

import (
	"math/big"

	"github.com/Masterminds/squirrel"
	"github.com/Swapica/order-indexer-svc/internal/data"
	"github.com/Swapica/order-indexer-svc/internal/gobind"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const ordersTable = "orders"

type orders struct {
	db       *pgdb.DB
	srcChain string
}

func NewOrders(db *pgdb.DB, chainName string) data.Orders {
	return orders{db: db, srcChain: chainName}
}

func (q orders) Insert(order data.Order) error {
	// todo: test if it's added, or add a struct field
	stmt := squirrel.Insert(ordersTable).SetMap(structs.Map(order)).
		Columns("src_chain").Values(q.srcChain)
	err := q.db.Exec(stmt)
	return errors.Wrap(err, "failed to insert order")
}

func (q orders) Update(id *big.Int, st gobind.SwapicaStatus) error {
	stmt := squirrel.Update(ordersTable).
		SetMap(map[string]interface{}{"state": st.State, "executed_by": st.ExecutedBy, "match_swapica": st.MatchSwapica}).
		Where(squirrel.Eq{"id": id})
	err := q.db.Exec(stmt)
	return errors.Wrap(err, "failed to insert order")
}
