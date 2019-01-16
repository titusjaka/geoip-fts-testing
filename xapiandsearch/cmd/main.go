package main

import (
	"flag"
	"github.com/titusjaka/geoip-fts-testing"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
	"github.com/titusjaka/geoip-fts-testing/xapiandsearch"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	url      = flag.String("url", "http://localhost:8880", "Xapiand URL")
	index    = flag.String("index", "geoip_index", "Xapiand index name")
	filename = flag.String("filename", "", "Path to SCV with GEO-info")
	bulkSize = flag.Int("bulksize", 10000, "Number of documents to collect before committing")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	if *url == "" {
		log.Fatal("missing url parameter")
	}
	if *index == "" {
		log.Fatal("missing index parameter")
	}
	if *filename == "" {
		log.Fatal("missing PATH to CSV-file")
	}
	if *bulkSize <= 0 {
		log.Fatal("bulk-size must be a positive number")
	}

	client := xapiandsearch.NewXapiandClient(*url, *index)

	csvChan := make(chan csv_helpers.DataLine)
	xapiandChan := make(chan xapiandsearch.XapiandObject)

	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(
		func() error {
			return client.Update(*bulkSize, xapiandChan)
		},
	)

	gr.Go(
		func() error {
			defer log.Println("[DEBUG] CSV channel is closed")
			defer close(csvChan)
			return csv_helpers.ReadDataFromCSV(*filename, ctx, csvChan)
		},
	)

	gr.Go(
		func() error {
			defer log.Println("[DEBUG] Elastic channel is closed")
			defer close(xapiandChan)
			for line := range csvChan {
				xo := xapiandsearch.DataLineToXapiandObject(&line)
				id := geoip_fts_testing.GetIdFromIpRange(line.StartIP, line.EndIP)
				xo.ID = id
				xapiandChan <- *xo
			}
			return nil
		},
	)

	err := gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
