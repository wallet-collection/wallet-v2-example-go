package biz

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
	"wallet/model"
)

type MemberCoinBiz struct {
}

type MemberCoinInfo struct {
	MemberId       int64           // 用户ID
	CoinSymbol     string          // 币种符号
	CoinName       string          // 币种名称
	CoinIcon       string          // 币种图标
	UsdtPrice      decimal.Decimal // 币种价格
	Precision      int             // 精度
	IsTransfer     int             // 是否可划转（0：否，1：手动，2：自动）
	TransferRate   decimal.Decimal // 划转费率
	MinTransferFee decimal.Decimal // 最低划转费用
	MinTransfer    decimal.Decimal // 最低划转
	MaxTransfer    decimal.Decimal // 最大划转
	Balance        decimal.Decimal // 可用余额
	FrozenBalance  decimal.Decimal // 冻结余额
	VirtualBalance decimal.Decimal // 虚拟余额
	CreateTime     time.Time       // 创建时间
	ModifiedTime   time.Time       // 更新时间
}

func (m *MemberCoinBiz) ListAll(memberId int64, offset, limit int, db *gorm.DB) (int64, []*MemberCoinInfo, error) {

	memberCoinModel := model.NewMemberCoinModel(db)
	coinModel := model.NewCoinModel(db)

	coins, err := coinModel.ListAll()
	if err != nil {
		return 0, nil, err
	}

	total, memberCoins, err := memberCoinModel.ListAllByMemberId(memberId, offset, limit)
	if err != nil {
		return total, nil, err
	}
	var info []*MemberCoinInfo
	if memberId == 0 {
		coinMap := make(map[string]*model.Coin)
		for _, coin := range coins {
			coinMap[coin.Symbol] = coin
		}
		for _, memberCoin := range memberCoins {
			coin, ok := coinMap[memberCoin.CoinSymbol]
			if ok {
				info = append(info, &MemberCoinInfo{
					MemberId:       memberCoin.MemberId,
					CoinSymbol:     coin.Symbol,
					CoinName:       coin.Name,
					CoinIcon:       coin.Icon,
					UsdtPrice:      coin.UsdtPrice,
					Precision:      coin.Precision,
					IsTransfer:     coin.IsTransfer,
					TransferRate:   coin.TransferRate,
					MinTransferFee: coin.MinTransferFee,
					MinTransfer:    coin.MinTransfer,
					MaxTransfer:    coin.MaxTransfer,
					Balance:        memberCoin.Balance,
					FrozenBalance:  memberCoin.FrozenBalance,
					VirtualBalance: memberCoin.VirtualBalance,
					CreateTime:     time.Now(),
					ModifiedTime:   time.Now(),
				})
			}
		}
		return total, info, nil
	}
	coinMap := make(map[string]*model.MemberCoin)
	for _, coin := range memberCoins {
		coinMap[coin.CoinSymbol] = coin
	}
	decimal0 := decimal.NewFromInt(0)
	var addData []*model.MemberCoin
	for _, coin := range coins {
		memberCoin, ok := coinMap[coin.Symbol]
		if !ok && memberId > 0 {
			tempCoin := &model.MemberCoin{
				MemberId:      memberId,
				CoinSymbol:    coin.Symbol,
				Balance:       decimal0,
				FrozenBalance: decimal0,
				CreateTime:    time.Now(),
				ModifiedTime:  time.Now(),
			}
			addData = append(addData, tempCoin)
			info = append(info, &MemberCoinInfo{
				MemberId:       memberId,
				CoinSymbol:     coin.Symbol,
				CoinName:       coin.Name,
				CoinIcon:       coin.Icon,
				UsdtPrice:      coin.UsdtPrice,
				Precision:      coin.Precision,
				IsTransfer:     coin.IsTransfer,
				TransferRate:   coin.TransferRate,
				MinTransferFee: coin.MinTransferFee,
				MinTransfer:    coin.MinTransfer,
				MaxTransfer:    coin.MaxTransfer,
				Balance:        decimal0,
				FrozenBalance:  decimal0,
				VirtualBalance: decimal0,
				CreateTime:     time.Now(),
				ModifiedTime:   time.Now(),
			})
		} else {
			info = append(info, &MemberCoinInfo{
				MemberId:       memberCoin.MemberId,
				CoinSymbol:     coin.Symbol,
				CoinName:       coin.Name,
				CoinIcon:       coin.Icon,
				UsdtPrice:      coin.UsdtPrice,
				Precision:      coin.Precision,
				IsTransfer:     coin.IsTransfer,
				TransferRate:   coin.TransferRate,
				MinTransferFee: coin.MinTransferFee,
				MinTransfer:    coin.MinTransfer,
				MaxTransfer:    coin.MaxTransfer,
				Balance:        memberCoin.Balance,
				FrozenBalance:  memberCoin.FrozenBalance,
				VirtualBalance: memberCoin.VirtualBalance,
				CreateTime:     time.Now(),
				ModifiedTime:   time.Now(),
			})
		}
	}

	if len(addData) > 0 {
		_ = memberCoinModel.InsertAll(addData)
	}

	return total, info, nil
}

func (m *MemberCoinBiz) Balance(memberId int64, coinSymbol string, db *gorm.DB) (*MemberCoinInfo, error) {

	memberCoinModel := model.NewMemberCoinModel(db)
	coinModel := model.NewCoinModel(db)

	memberCoin, err := memberCoinModel.FindByMemberIdAndCoinSymbol(memberId, coinSymbol)
	var info *MemberCoinInfo
	decimal0 := decimal.NewFromInt(0)

	coin, err := coinModel.FindBySymbol(coinSymbol)
	if err != nil {
		return nil, err
	}

	if err != nil {

		memberCoin = &model.MemberCoin{
			MemberId:       memberId,
			CoinSymbol:     coinSymbol,
			Balance:        decimal.NewFromInt(0),
			FrozenBalance:  decimal.NewFromInt(0),
			VirtualBalance: decimal.NewFromInt(0),
			CreateTime:     time.Now(),
			ModifiedTime:   time.Now(),
		}
		err = memberCoinModel.Insert(memberCoin)
		if err != nil {
			return nil, err
		}
		info = &MemberCoinInfo{
			MemberId:       memberId,
			CoinSymbol:     coin.Symbol,
			CoinName:       coin.Name,
			CoinIcon:       coin.Icon,
			UsdtPrice:      coin.UsdtPrice,
			Precision:      coin.Precision,
			IsTransfer:     coin.IsTransfer,
			TransferRate:   coin.TransferRate,
			MinTransferFee: coin.MinTransferFee,
			MinTransfer:    coin.MinTransfer,
			MaxTransfer:    coin.MaxTransfer,
			Balance:        decimal0,
			FrozenBalance:  decimal0,
			VirtualBalance: decimal0,
			CreateTime:     time.Now(),
			ModifiedTime:   time.Now(),
		}
	} else {
		info = &MemberCoinInfo{
			MemberId:       memberId,
			CoinSymbol:     coin.Symbol,
			CoinName:       coin.Name,
			CoinIcon:       coin.Icon,
			UsdtPrice:      coin.UsdtPrice,
			Precision:      coin.Precision,
			IsTransfer:     coin.IsTransfer,
			TransferRate:   coin.TransferRate,
			MinTransferFee: coin.MinTransferFee,
			MinTransfer:    coin.MinTransfer,
			MaxTransfer:    coin.MaxTransfer,
			Balance:        memberCoin.Balance,
			FrozenBalance:  memberCoin.FrozenBalance,
			VirtualBalance: memberCoin.VirtualBalance,
			CreateTime:     time.Now(),
			ModifiedTime:   time.Now(),
		}
	}

	return info, nil
}
