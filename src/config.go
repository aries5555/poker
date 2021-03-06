package src

var MatchSamplesPaths = []string{
	"./match_samples/seven_cards_with_ghost.result.json",
	//"./match_samples/seven_cards_with_ghost.json",
	//"./match_samples/five_cards_with_ghost.json",
	//"./match_samples/match.json",
}

// 牌色对应编号
var Suits = map[string]int{
	"s": 3,
	"h": 2,
	"d": 1,
	"c": 0,
}

// 牌面对应编号（对应bit位置）
var Faces = map[string]uint64{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

const (
	StraightFlush = 8 // 皇家同花顺&同花顺
	FourOfAKind   = 7 // 四条
	FullHouse     = 6 // 葫芦
	Flush         = 5 // 同花
	Straight      = 4 // 顺子
	ThreeOfAKind  = 3 // 三条
	TwoPair       = 2 // 两对
	OnePair       = 1 // 一对
	HighCard      = 0 // 散牌
)

const (
	// 特殊值        AKQJT98765432
	A2345 = 4111 // 1000000001111
	AKQJT = 7936 // 1111100000000
	A     = 4096 // 1000000000000
)
