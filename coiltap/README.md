# Coiltap

Coiltap uses [libpcap]() to capture HTTP traffic to and from a local port,
storing the results in ElasticSearch.

## Usage

    coiltap -p 80 -i eth0 http://localhost:9200/requests

## Development

Coiltap doesn't currently work on Vagrant. I'm not sure, but something to do
with the type of packets that are returned on `localhost`. 
