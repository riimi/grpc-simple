package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonClient struct {
	addr string
}

func NewJsonClient(addr string, nWorker int) *JsonClient {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2 * nWorker
	return &JsonClient{
		addr: addr,
	}
}

func (j *JsonClient) Data() interface{} {
	return struct{}{}
}

func (j *JsonClient) LoadTest() func(context.Context, interface{}) {
	return func(ctx context.Context, x interface{}) {
		req := IncrRequest{
			Api: "Incr",
			Sid: "nadana",
			Uid: "123132",
			Key: "myspoon",
		}
		body, err := json.Marshal(req)
		if err != nil {
			log.Print(err)
			return
		}
		addr := "http://" + j.addr + "/v1/count/incr"
		res, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Print(err)
			return
		}
		//bs, _ := ioutil.ReadAll(res.Body)
		//log.Print(string(bs))
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}
}
