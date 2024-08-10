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

// 交易品种
type TradeKind int32

const (
	// 现货
	SPOT TradeKind = iota
	// 永续合约
	FUTURES
)

func (v TradeKind) Valid() bool {
	return v >= SPOT && v <= FUTURES
}

// 交易类型
type TradeType int32

const (
	// 限价交易
	LIMIT TradeType = iota
	// 市价交易
	MARKET
)

func (v TradeType) Valid() bool {
	return v >= LIMIT && v <= MARKET
}

// 交易方式
type TradeSide int32

const (
	// 做多
	LONG TradeSide = iota
	// 做空格
	SHORT
)

func (v TradeSide) Valid() bool {
	return v >= LONG && v <= SHORT
}

// 交易动作
type TradeAction int32

const (
	// 买入
	BUY TradeAction = iota
	// 卖出
	SELL
)

func (v TradeAction) Valid() bool {
	return v >= BUY && v <= SELL
}
