package v2

import (
	"github.com/huangyisan/bitget-golang-sdk-api/internal"
	"github.com/huangyisan/bitget-golang-sdk-api/internal/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init() *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := internal.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/assets", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}

func (p *SpotAccountClient) CustomGet(params map[string]string, uri string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(uri, params)
	return resp, err
}

func (p *SpotAccountClient) CustomPost(params map[string]string, uri string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(uri, postBody)
	return resp, err
}
