// made by recanman
package opnsense

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type ClientConfig struct {
	Host       string
	CertPath   string
	Key        string
	Secret     string
	SkipVerify bool
}

type Client struct {
	host   string
	key    string
	secret string
	c      *http.Client
}

func loadCert(certPath string) (*x509.CertPool, error) {
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %v", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certData) {
		return nil, fmt.Errorf("failed to append cert to pool")
	}

	return certPool, nil
}

func New(c ClientConfig) (*Client, error) {
	var err error

	tlsConfig := &tls.Config{
		InsecureSkipVerify: c.SkipVerify,
	}
	if !c.SkipVerify {
		var certPool *x509.CertPool
		if certPool, err = loadCert(c.CertPath); err != nil {
			return nil, err
		}
		tlsConfig.RootCAs = certPool
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	client := &Client{
		host:   c.Host,
		key:    c.Key,
		secret: c.Secret,
		c:      httpClient,
	}
	return client, nil
}

func (c *Client) Post(path string, data any) (*http.Response, error) {
	var jsonData []byte
	var err error

	if jsonData, err = json.Marshal(data); err != nil {
		return nil, err
	}
	var url *url.URL
	if url, err = url.Parse(fmt.Sprintf("https://%s%s", c.host, path)); err != nil {
		return nil, err
	}

	var req *http.Request
	if req, err = http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonData)); err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.key, c.secret)
	req.Header.Set("Content-Type", "application/json")
	return c.c.Do(req)
}
