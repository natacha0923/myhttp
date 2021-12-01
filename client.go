package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/url"
)

type Requester interface {
	SendRequest(urlPath string) ([]byte, error)
}

type Client struct {
	Requester Requester
}

func (c *Client) GetRequestWithMD5(urlPath string) (Response, error) {
	urlPath, err := checkProtocolSchema(urlPath)
	if err != nil {
		return Response{}, err
	}

	res, err := c.Requester.SendRequest(urlPath)
	if err != nil {
		return Response{}, err
	}

	result := Response{
		Url:  urlPath,
		Hash: getMD5Hash(res),
	}

	return result, nil
}

type Response struct {
	Url  string
	Hash string
}

// checkProtocolSchema checks that urlPath includes protocol scheme. If not, adds it to url string.
func checkProtocolSchema(urlPath string) (string, error) {
	if urlPath == "" {
		return "", errors.New("url is empty")
	}

	const httpSchema, httpsSchema = "http", "https"
	result, err := url.Parse(urlPath)
	if err != nil {
		return "", err
	}

	if result.Scheme == httpSchema || result.Scheme == httpsSchema {
		return urlPath, nil
	}

	return httpSchema + "://" + urlPath, nil
}

// getMD5Hash returns the encoding MD5 checksum of the str.
func getMD5Hash(str []byte) string {
	hash := md5.Sum(str)
	return hex.EncodeToString(hash[:])
}
