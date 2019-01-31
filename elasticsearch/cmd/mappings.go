package main

const defaultInfoMapping string = `
{
  "settings": {
    "number_of_shards": 5,
    "number_of_replicas": 1
  },
  "mappings": {
    "_doc": {
      "dynamic": "strict",
      "properties": {
        "ip_address": {
          "type": "ip_range",
          "coerce": false,
          "index": true
        },
        "country_code": {
          "type": "keyword",
          "index": false
        },
        "region": {
          "type": "keyword",
          "index": false
        },
        "region_code": {
          "type": "integer",
          "index": false
        },
        "city": {
          "type": "keyword",
          "index": false
        },
        "city_code": {
          "type": "integer",
          "index": false
        },
        "conn_speed": {
          "type": "keyword",
          "index": false
        },
        "mobile_isp": {
          "type": "keyword",
          "index": false
        },
        "mobile_isp_code": {
          "type": "integer",
          "index": false
        },
        "proxy_type": {
          "type": "keyword",
          "index": false
        }
      }
    }
  }
}
`
const defaultCountryMapping string = `
{
  "settings": {
    "number_of_shards": 5,
    "number_of_replicas": 1
  },
  "mappings": {
    "_doc": {
      "dynamic": "strict",
      "properties": {
        "title": {
          "type": "text",
          "index": true,
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "iso": {
          "type": "keyword",
          "index": true
        }
      }
    }
  }
}
`

const defaultRegionsMapping string = `
{
  "settings": {
    "number_of_shards": 5,
    "number_of_replicas": 1
  },
  "mappings": {
    "_doc": {
      "dynamic": "strict",
      "properties": {
        "country_iso": {
          "type": "keyword",
          "index": true
        },
        "title": {
          "type": "text",
          "index": true,
          "fields": {
            "keyword": {
              "type": "keyword"
            }
          }
        },
        "code": {
          "type": "integer",
          "index": true
        }
      }
    }
  }
}
`
