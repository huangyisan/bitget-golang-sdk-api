package v2

import (
	"github.com/huangyisan/bitget-golang-sdk-api/internal"
	"github.com/huangyisan/bitget-golang-sdk-api/internal/common"
)

type SpotMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotMarketClient) Init() *SpotMarketClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *SpotMarketClient) Coins() (string, error) {
	params := internal.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/coins", params)
	return resp, err
}

func (p *SpotMarketClient) Symbols(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
	return resp, err
}

func (p *SpotMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/fills", params)
	return resp, err
}

func (p *SpotMarketClient) Orderbook(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/orderbook", params)
	return resp, err
}

func (p *SpotMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/tickers", params)
	return resp, err
}

func (p *SpotMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/candles", params)
	return resp, err
}

func (p *SpotMarketClient) HistoryCandles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/history-candles", params)
	return resp, err
}

func (p *SpotMarketClient) CustomGet(params map[string]string, uri string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(uri, params)
	return resp, err
}
func (p *SpotMarketClient) CustomPost(params map[string]string, uri string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(uri, postBody)
	return resp, err
}
