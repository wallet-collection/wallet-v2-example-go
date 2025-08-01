package client

import (
	"errors"
	"github.com/shopspring/decimal"
	"net/http"
	"net/url"
)

type CreateWalletRes struct {
	Res
	Data CreateWalletDataRes `json:"data"` // 创建钱包的地址
}

type CreateWalletDataRes struct {
	Address string `json:"address"` // 钱包地址
}

type CreateWithdrawRes struct {
	Res
	Data CreateWithdrawDataRes `json:"data"` // 创建钱包的地址
}

type CreateWithdrawDataRes struct {
	From string `json:"from"` // 发起转账的地址
	Hash string `json:"hash"` // 哈希
}

type RechargeAddressRes struct {
	Res
	Data RechargeAddressDataRes `json:"data"` //
}

type RechargeAddressDataRes struct {
}

type WalletClient struct {
	appid     string
	secretKey string
	client    *Client
}

var WalletClientGlobal *WalletClient

// NewWalletClient 创建
func NewWalletClient(appid string, secretKey string, url string) *WalletClient {

	client := NewClient(url, 10*1000)

	walletClient := &WalletClient{
		appid:     appid,
		secretKey: secretKey,
		client:    client,
	}
	WalletClientGlobal = walletClient
	return walletClient
}

// CreateWallet 创建钱包
func (w *WalletClient) CreateWallet(memberId string, networkName string, coinSymbol string, callUrl string) (*CreateWalletDataRes, error) {

	data := make(map[string]interface{})
	data["network_name"] = networkName
	data["coin_symbol"] = coinSymbol
	data["member_id"] = memberId
	data["call_url"] = callUrl

	var res CreateWalletRes

	err := w.signRequest("/createWallet", http.MethodPost, nil, nil, data, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return &res.Data, nil
}

// CreateWithdraw 生成提现
func (w *WalletClient) CreateWithdraw(networkName string, coinSymbol string, address string, amount decimal.Decimal, businessId string, privateKey string, callUrl string) (*CreateWithdrawDataRes, error) {

	data := make(map[string]interface{})
	data["network_name"] = networkName
	data["coin_symbol"] = coinSymbol
	data["address"] = address
	data["amount"] = amount.String()
	data["business_id"] = businessId
	data["private_key"] = privateKey
	data["call_url"] = callUrl

	var res CreateWithdrawRes

	err := w.signRequest("/createWithdraw", http.MethodPost, nil, nil, data, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return &res.Data, nil
}

// RechargeAddress 修改充值地址
func (w *WalletClient) RechargeAddress(networkName string, coinSymbol string, address string) (*RechargeAddressDataRes, error) {

	data := make(map[string]interface{})
	data["network_name"] = networkName
	data["coin_symbol"] = coinSymbol
	data["address"] = address

	var res RechargeAddressRes

	err := w.signRequest("/rechargeAddress", http.MethodPost, nil, nil, data, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return &res.Data, nil
}

// post 请求
func (w *WalletClient) signRequest(urlStr string, method string, header http.Header, params url.Values, data map[string]interface{}, res interface{}) error {

	if data != nil {
		data["appid"] = w.appid
		data["secret_key"] = w.secretKey
		// 签名
		sign := Sign(data)
		data["sign"] = sign

		delete(data, "secret_key")
	}

	return w.client.request(urlStr, method, header, params, data, res)
}
