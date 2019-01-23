package elasticsearch

const defaultInfoMapping string = `
{
  "settings": {
    "number_of_shards": 3,
    "number_of_replicas": 0
  },
  "mappings": {
    "_doc": {
      "properties": {
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
        "country": {
          "type": "keyword",
          "index": false
        },
        "ip_address": {
          "type": "ip_range",
          "coerce": false
        },
        "isp": {
          "type": "keyword",
          "index": false
        },
        "mobile_carrier": {
          "type": "keyword",
          "index": false
        },
        "mobile_carrier_code": {
          "type": "integer",
          "index": false
        },
        "region": {
          "type": "keyword",
          "index": false
        },
        "region_code": {
          "type": "integer",
          "index": false
        }
      }
    }
  }
}
`
