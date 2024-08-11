package types

// 运行模式
type RunMode int32

const (
	// 回测
	BACKTEST RunMode = iota
	// 模拟盘
	SANDBOX
	// 实盘
	REAL
)

func (v RunMode) Valid() bool {
	return v >= BACKTEST && v <= REAL
}

// 交易类型
type TradeType int32

const (
	// 限价交易
	LIMIT TradeType = iota
	// 市价交易
	MARKET
	// 止损单
	STOP
	// 价单止损市
	STOP_MARKET
	// 止盈单
	TAKE_PROFIT
	// 市价止盈单
	TAKE_PROFIT_MARKET
	// 跟踪止损市价单
	TRAILING_STOP_MARKET
)

func (v TradeType) Valid() bool {
	return v >= LIMIT && v <= TRAILING_STOP_MARKET
}

// 交易方向
type TradeSide int32

const (
	// 做多
	BUY TradeSide = iota
	// 做空
	SELL
)

func (v TradeSide) Valid() bool {
	return v >= BUY && v <= SELL
}
