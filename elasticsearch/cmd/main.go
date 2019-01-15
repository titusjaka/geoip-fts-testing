package main

import (
	"flag"
	"github.com/titusjaka/geoip-fts-testing"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
	"github.com/titusjaka/geoip-fts-testing/elasticsearch"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	url      = flag.String("url", "http://localhost:9200", "Elasticsearch URL")
	index    = flag.String("index", "geoip_index", "Elasticsearch index name")
	typ      = flag.String("type", "_doc", "Elasticsearch type name")
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
	if *typ == "" {
		log.Fatal("missing type parameter")
	}
	if *filename == "" {
		log.Fatal("missing PATH to CSV-file")
	}
	if *bulkSize <= 0 {
		log.Fatal("bulk-size must be a positive number")
	}

	client, err := elasticsearch.NewElasticClient(*url)
	if err != nil {
		log.Fatal(err)
	}

	csvChan := make(chan csv_helpers.DataLine)
	elasticChan := make(chan elasticsearch.ElasticObject)

	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(
		func() error {
			return client.Update(ctx, *index, *bulkSize, elasticChan)
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
			defer close(elasticChan)
			for line := range csvChan {
				eo := elasticsearch.DataLineToElasticObject(&line)
				id := geoip_fts_testing.GetIdFromIpRange(eo.IPAddress.StartIP, eo.IPAddress.EndIP)
				eo.ID = id
				elasticChan <- *eo
			}
			return nil
		},
	)

	err = gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
