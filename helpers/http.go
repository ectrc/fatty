package helpers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MustParse(_url string) *url.URL {
	u, err := url.Parse(_url)
	if err != nil {
		panic(err)
	}
	return u
}

type ProxiedClient struct {
	client *http.Client
	proxy *http.Transport
}

func NewProxiedClient() *ProxiedClient {
	return &ProxiedClient{
		client: &http.Client{
		},
	}
}

func (p *ProxiedClient) Client() *http.Client {
	return p.client
}

func (p *ProxiedClient) SetProxy(proxyUrl string) {
	p.proxy = &http.Transport{
		Proxy: http.ProxyURL(MustParse(proxyUrl)),
	}
	p.client.Transport = p.proxy
}

type ProxiedClientResponse struct {
	Body []byte
	Err error
}

func (p *ProxiedClient) Get(url string, headers ...JSON) ProxiedClientResponse {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")

	for _, header := range headers {
		for key, value := range header {
			request.Header.Set(key, value.(string))
		}
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}

	return ProxiedClientResponse{
		Body: bytes,
	}
}

func (p *ProxiedClient) Post(url string, body JSON, headers ...JSON) ProxiedClientResponse {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
	request.Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	for _, header := range headers {
		for key, value := range header {
			request.Header.Set(key, value.(string))
		}
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}

	return ProxiedClientResponse{
		Body: bytes,
	}
}

func (p *ProxiedClient) PostForm(url string, body url.Values, headers ...JSON) ProxiedClientResponse {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
	request.Body = io.NopCloser(strings.NewReader(body.Encode()))

	for _, header := range headers {
		for key, value := range header {
			request.Header.Set(key, value.(string))
		}
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}

	return ProxiedClientResponse{
		Body: bytes,
	}
}

func (p *ProxiedClient) Put(url string, body JSON, headers ...JSON) ProxiedClientResponse {
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
	request.Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	for _, header := range headers {
		for key, value := range header {
			request.Header.Set(key, value.(string))
		}
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProxiedClientResponse{Err: err}
	}

	return ProxiedClientResponse{
		Body: bytes,
	}
}