package main

import (
	"flag"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
	"github.com/titusjaka/geoip-fts-testing/elasticsearch"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
)

var (
	// Geo-Info index
	fillInfo  = flag.Bool("info.fill", false, "Fill Info index")
	infoIndex = flag.String("info.index", "geo-info-index", "Elasticsearch infoIndex name")
	infoFile  = flag.String("info.file", "", "Path to SCV with GEO-info")
	// Country info index
	fillCountry  = flag.Bool("country.fill", false, "Fill Info index")
	countryIndex = flag.String("country.index", "country-info-index", "Elasticsearch countryIndex name")
	countryFile  = flag.String("country.file", "", "Path to SCV with country info")
	// Region info index
	fillRegion  = flag.Bool("region.fill", false, "Fill Info index")
	regionIndex = flag.String("region.index", "region-info-index", "Elasticsearch regionIndex name")
	regionFile  = flag.String("region.file", "", "Path to SCV with region info")
	// Other params
	url      = flag.String("url", "http://localhost:9200", "Elasticsearch URL")
	typ      = flag.String("type", "_doc", "Elasticsearch type name")
	bulkSize = flag.Uint64("bulksize", 100000, "Number of documents to collect before committing")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	if !*fillInfo && !*fillCountry && !*fillRegion {
		log.Printf("no jobs")
		os.Exit(0)
	}

	if *url == "" {
		log.Fatal("missing url parameter")
	}
	if *typ == "" {
		log.Fatal("missing type parameter")
	}
	if *bulkSize <= 0 {
		log.Fatal("bulk-size must be a positive number")
	}

	// Geo-Info index
	if *fillInfo && *infoIndex == "" {
		log.Fatal("missing infoIndex parameter")
	}
	if *fillInfo && *infoFile == "" {
		log.Fatal("missing PATH to Geo-Info CSV-file")
	}

	// Country info index
	if *fillCountry && *countryIndex == "" {
		log.Fatal("missing countryIndex parameter")
	}
	if *fillCountry && *countryFile == "" {
		log.Fatal("missing PATH to country CSV-file")
	}

	// Country info index
	if *fillRegion && *regionIndex == "" {
		log.Fatal("missing regionIndex parameter")
	}
	if *fillRegion && *regionFile == "" {
		log.Fatal("missing PATH to region CSV-file")
	}

	client, err := elasticsearch.NewElasticClient(*url)
	if err != nil {
		log.Fatal(err)
	}

	if *fillInfo {
		fillInfoIndex(client)
	}

	if *fillCountry {
		fillCountryIndex(client)
	}
	if *fillRegion {
		fillRegionIndex(client)
	}
}

func fillRegionIndex(client *elasticsearch.ElasticClient) {
	gr, ctx := errgroup.WithContext(context.Background())
	csvRegionChan := make(chan csv_helpers.RegionLine)
	esRegionChan := make(chan interface{})

	if err := client.CreateIndexMapping(ctx, *regionIndex, defaultRegionsMapping); err != nil {
		log.Fatalf("Cannot create %s mapping", *regionIndex)
	}

	gr.Go(
		func() error {
			return client.Update(ctx, *regionIndex, *bulkSize, esRegionChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] CSV region channel is closed")
			defer close(csvRegionChan)
			return csv_helpers.ReadRegionInfoFromCSV(*regionFile, ctx, csvRegionChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] Elastic region channel is closed")
			defer close(esRegionChan)
			for line := range csvRegionChan {
				eo := elasticsearch.RegionInfoToElasticObject(&line)
				esRegionChan <- *eo
			}
			return nil
		},
	)
	err := gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func fillCountryIndex(client *elasticsearch.ElasticClient) {
	gr, ctx := errgroup.WithContext(context.Background())
	csvCountryChan := make(chan csv_helpers.CountryLine)
	esCountryChan := make(chan interface{})

	if err := client.CreateIndexMapping(ctx, *countryIndex, defaultCountryMapping); err != nil {
		log.Fatalf("Cannot create %s mapping", *countryIndex)
	}

	gr.Go(
		func() error {
			return client.Update(ctx, *countryIndex, *bulkSize, esCountryChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] CSV Country channel is closed")
			defer close(csvCountryChan)
			return csv_helpers.ReadCountryInfoFromCSV(*countryFile, ctx, csvCountryChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] Elastic Country channel is closed")
			defer close(esCountryChan)
			for line := range csvCountryChan {
				eo := elasticsearch.CountryInfoToElasticObject(&line)
				esCountryChan <- *eo
			}
			return nil
		},
	)
	err := gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func fillInfoIndex(client *elasticsearch.ElasticClient) {
	gr, ctx := errgroup.WithContext(context.Background())
	csvInfoChan := make(chan csv_helpers.GeoInfoLine)
	esInfoChan := make(chan interface{})

	if err := client.CreateIndexMapping(ctx, *infoIndex, defaultInfoMapping); err != nil {
		log.Fatalf("Cannot create %s mapping", *infoIndex)
	}

	gr.Go(
		func() error {
			return client.Update(ctx, *infoIndex, *bulkSize, esInfoChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] CSV Info channel is closed")
			defer close(csvInfoChan)
			return csv_helpers.ReadGeoInfoFromCSV(*infoFile, ctx, csvInfoChan)
		},
	)
	gr.Go(
		func() error {
			defer log.Println("[DEBUG] Elastic Info channel is closed")
			defer close(esInfoChan)
			for line := range csvInfoChan {
				eo := elasticsearch.GeoInfoLineToElasticObject(&line)
				esInfoChan <- *eo
			}
			return nil
		},
	)
	err := gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
