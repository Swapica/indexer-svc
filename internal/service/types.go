package service

import "database/sql"

type Order struct {
	// ID surrogate key is strongly preferred against PRIMARY KEY (OrderID, SrcChain)
	ID         int64  `structs:"-" db:"id"`
	OrderID    int64  `structs:"order_id" db:"order_id"`
	SrcChain   int64  `structs:"src_chain" db:"src_chain"`
	Creator    string `structs:"creator" db:"creator"`
	SellToken  int64  `structs:"sell_token" db:"sell_token"`
	BuyToken   int64  `structs:"buy_token" db:"buy_token"`
	SellAmount string `structs:"sell_amount" db:"sell_amount"`
	BuyAmount  string `structs:"buy_amount" db:"buy_amount"`
	DestChain  int64  `structs:"dest_chain" db:"dest_chain"`
	State      uint8  `structs:"state" db:"state"`
	UseRelayer bool   `structs:"use_relayer" db:"use_relayer"`

	// ExecutedByMatch foreign key for match_orders(ID)
	ExecutedByMatch sql.NullInt64  `structs:"executed_by_match,omitempty,omitnested" db:"executed_by_match"`
	MatchID         sql.NullInt64  `structs:"match_id,omitempty,omitnested" db:"match_id"`
	MatchSwapica    sql.NullString `structs:"match_swapica,omitempty,omitnested" db:"match_swapica"`
}

type Match struct {
	// ID surrogate key is strongly preferred against PRIMARY KEY (MatchID, SrcChain)
	ID       int64 `structs:"-" db:"id"`
	MatchID  int64 `structs:"match_id" db:"match_id"`
	SrcChain int64 `structs:"src_chain" db:"src_chain"`
	// OriginOrder foreign key for orders(ID)
	OriginOrder int64  `structs:"origin_order" db:"origin_order"`
	OrderID     int64  `structs:"order_id" db:"order_id"`
	OrderChain  int64  `structs:"order_chain" db:"order_chain"`
	Creator     string `structs:"creator" db:"creator"`
	SellToken   int64  `structs:"sell_token" db:"sell_token"`
	SellAmount  string `structs:"sell_amount" db:"sell_amount"`
	State       uint8  `structs:"state" db:"state"`
	UseRelayer  bool   `structs:"use_relayer" db:"use_relayer"`
}
