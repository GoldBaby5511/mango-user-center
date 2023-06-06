package service

import (
	"encoding/json"
	"errors"
	"mango-user-center/config"
	"mango-user-center/pkg/util"
)

type Rpc struct{}

type rpcCommonResp struct {
	Code int
	Msg  string
	Data json.RawMessage
}

type postAddressResp struct {
	PublicKey     string `json:"public_key"`
	PrivateKey    string `json:"private_key"`
	WalletAddress string `json:"wallet_address"`
	Mnemonic      string `json:"mnemonic"`
}

func (Rpc) post(api string, bodyData interface{}, to interface{}) error {
	var resp rpcCommonResp
	err := util.Http.PostJson(api, bodyData).DoTo(&resp)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}
	return json.Unmarshal(resp.Data, to)
}

// 生成钱包
func (r Rpc) GenerateAddress() (postAddressResp, error) {
	var data postAddressResp
	err := r.post(config.RPC.CommonLogic+"/address", nil, &data)
	return data, err
}
