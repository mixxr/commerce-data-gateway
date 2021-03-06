# Data Gateway
Intelligent API Gateway to create a data commerce platform in which vendors (owners) can publish and sell data.

## mission
The idea is to create a platform for a marketplace in which data owners and data shoppers can meet. There are 2 actors: 
- **owners**: provide data, create premium plans if needed
- **shoppers**: search for data, consume data, subscribe to premium plans if needed

Each data owner can decide what kind of data provide, the data is always expressed via a table format (columns and rows) and when ingested the owner can decide what part of the table is open to free access and what part, if any, requires a premium. 

The architecture has to be RESTful, light and easy to setup. 

## project modules
The project is composed by several modules in order to make it scalable, elastic and agile.

- **Tag Manager** module: it is in charge to manage the tags that the owners can assign to the data and that the shoppers can use to make searches
- **Ingestor** module: it is used by the owners to ingest data, it can be used via CLI or Web or API of course. The Ingestor uses the Tag Manager for data categorization. The data can be ingested by humans (eg. via web interface) or by application (eg. via batch operations) and can be provided in several formats like xml, csv, json.
- **Storefront** module: it is used by the shoppers, they can search, retrieve data. The data is formatted for humans (eg. html5, csv) or for applications (eg. json, xml)
- **Identity Manager** module: it is responsible for accounts, plans and subscriptions.


## to Configure
The application uses viper with YAML configuration file. The environemnt variables DCGW_RUNMODE and DCGW_CONFIGPATH are used for:
- DCGW_RUNMODE: it is the runmode for the application. Eg. dev, prod. The runmode is used to build the config filename to search:
  - For example, DCGW_RUNMODE=dev means a file named config.dev.yaml is needed to load the configurations
  - Default is 'dev'
- DCGW_CONFIGPATH: it is used to search the YAML config file. 
  - Default is the current directory.

## to RUN
### Pre-conditions
- MariaDB docker is needed (or a local installation). The project provides a Dockerfile at location: scripts/mariadb/Dockerfile
- when the docker is up then execute the gateway:
### Run via maven
`cd gateway`

`mvn clean install`
or
`DCGW_RUNMODE=prod mvn clean install`
to configure the runmode for the specific command

mvn runs tests verbose also

### Runs with go cli
`cd gateway/src`

`go build`
`go run main`

or
`DCGW_RUNMODE=prod go run main`

to run all tests:
`go test ./... -cover -count=1`

## 3rd party library usage
- JS: it uses ES6 fetch() 

## URL examples

```
curl -d "@table.json" -X POST https://localhost:8443/services/v1/samurl/bicycleurl.csv -ik
curl -d "@tablecols.en.json" -X POST https://localhost:8443/services/v1/samurl/bicycleurl/colnames/en -ik
curl -d "@tablevalues.json" -X POST https://localhost:8443/services/v1/samurl/bicycleurl/values -ik
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/colnames/it
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/colnames/en
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/values/0/10
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/values/1/10
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/values/0/1
curl -ik -X GET https://localhost:8443/services/v1/samurl/bicycleurl/values/0/0
```
## Service Status
A service can have 3 status:
- Deleted: 0
- Draft: 1, it is available to its owner only 
- Enabled: 2, it is available to shoppers 

