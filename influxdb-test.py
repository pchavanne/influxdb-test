## pip install influxdb-client

import influxdb_client

from influxdb_client.client.write_api import SYNCHRONOUS

bucket = "a_bucket"
org = "scm"
token = "aOToMeQiEtEMduD6fLFrUfIPI4__5RcEMGBGmWfL8DQMvd4VYt-iRakFZgYi5xjy6J3IpzSysdqu348MpCIb4A=="
# Store the URL of your InfluxDB instance
url="http://192.168.1.205:8086"

client = influxdb_client.InfluxDBClient(
  url=url,
  token=token,
  org=org
)

## Write
write_api = client.write_api(write_options=SYNCHRONOUS)

p = influxdb_client.Point("stats").tag("unit", "temperature").field("avg", 25.3).field("max",36)

write_api.write(bucket=bucket, org=org, record=p)

## Read

query_api = client.query_api()
query = ‘ from(bucket:"my-bucket")\
|> range(start: -10m)\
|> filter(fn:(r) => r._measurement == "stats")\
|> filter(fn: (r) => r.unit == "temperatture")‘
## |> filter(fn:(r) => r._field == "temperature" )‘
result = client.query_api().query(org=org, query=query)
results = []
for table in result:
    for record in table.records:
        results.append((record.get_field(), record.get_value()))

print(results)
[(temperature, 25.3)]
