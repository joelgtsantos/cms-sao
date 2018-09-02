# CMS Sao

A microservice that orbits around [CMS](https://github.com/cms-dev/cms) exposing
the Entries, Results and Scores entities as [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)
resources.

## Up and running

CMS Sao can be deployed and run as a Docker container or a standalone binary; either way is recommend to run this
application as the former one.

### Prerequisites

CMS Sao heavily relies on [CMS](https://github.com/cms-dev/cms) including its database, so
in order to have this application up and running you will need:

1. CMS 1.3.x or greater(the current Sao version was designed against the last CMS revision in Jan 2018)
2. CMS PostreSQL DB schema access (it could be the same credentials that CMS uses but is not recommended)
3. Docker engine 17.x or greater

### Deployment

**TODO**: Package as Docker image 

### Configuration

All the intrinsic configurations can be overriden via the `config.properties` file in Sao's working directory (you could 
use `config.properties.example` as guide) or as environment variables with the `SAO_` prefix. As for example in order 
to override the `datasource.url` value, you could start the Docker container with the following syntax:

```shell
docker run -p 8080:8080 -e 'SAO_DATASOURCE_URL=postgresql://10.10.37.10:5432/cmsdb' jossemargt/cmssao
```

Some of the properties of interest could be:

Property name | Default value | Description
--- | --- | ---
server.port | 8080 | The port where the application will listen for incoming requests
server.log.request | false | When set as true the application will log each incoming request to the STDOUT
server.error.tracedump | false | When set as true the application will include the error trace in failure responses
cms.url | http://localhost/ | The URL where Sao can stablish communication with CMS
datasource.username | cmsuser | PostgreSQL datasource username
datasource.password | notsecure | PostgreSQL datasource password
datasource.url | postgresql://localhost:5432 | PostgreSQL datasource URL

## License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE)
file for details.
