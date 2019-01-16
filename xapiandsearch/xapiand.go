package xapiandsearch

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"log"
	"sync/atomic"
	"time"
)

type XapiandClient struct {
	URL       string
	IndexName string
}

func NewXapiandClient(url, indexName string) *XapiandClient {
	return &XapiandClient{
		URL:       url,
		IndexName: indexName,
	}
}

func (client *XapiandClient) Update(bulkSize int, objects chan XapiandObject) error {
	var total, index uint64
	begin := time.Now()

	batch := make([]XapiandObject, bulkSize)

	indexUrl := fmt.Sprintf("%s/%s/:restore", client.URL, client.IndexName)

	for obj := range objects {
		current := atomic.AddUint64(&total, 1)
		duration := time.Since(begin).Seconds()
		seconds := int(duration)
		pps := int64(float64(current) / duration)
		fmt.Printf("%10d | %6d req/s | %02d:%02d\r", current, pps, seconds/60, seconds%60)

		batch[index] = obj
		index++

		if int(index) >= bulkSize {
			jsonData, err := json.Marshal(batch)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(jsonData).Post(indexUrl)
			if err != nil {
				log.Fatal(err)
				return err
			}
			if !resp.IsSuccess() {
				log.Fatal("Wrong response status: ", resp.Result())
				return fmt.Errorf("wrong response status")
			}
			batch = make([]XapiandObject, bulkSize)
			index = 0
		}
	}

	if len(batch) > 0 {
		jsonData, err := json.Marshal(batch)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(jsonData).Post(indexUrl)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if !resp.IsSuccess() {
			log.Fatal("Wrong response status: ", resp.Result())
			return fmt.Errorf("wrong response status")
		}
		batch = make([]XapiandObject, bulkSize)
		index = 0
	}

	dur := time.Since(begin).Seconds()
	sec := int(dur)
	pps := int64(float64(total) / dur)
	fmt.Printf("%10d | %6d req/s | %02d:%02d\n", total, pps, sec/60, sec%60)
	return nil
}
