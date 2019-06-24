# CMS Sao

A microservice that orbits around [CMS](https://github.com/cms-dev/cms) exposing
the Entry, Result, Score and Scoreboard entities as [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)
resources.

## Up and running

CMS Sao can be deployed and run as a Docker container or a standalone binary;
either way is recommend to run this application as the former one.

### Prerequisites

CMS Sao heavily relies on [CMS](https://github.com/cms-dev/cms) including its database, so
in order to have this application up and running you will need:

1. CMS 1.3.x or greater (the current Sao version was designed against the last CMS revision in Jan 2018)
2. CMS PostreSQL DB schema access (it could be the same credentials that CMS uses but is not recommended)
3. Docker engine 17.x or greater

### Deployment

CMS Sao can be deployed and run as a Docker container, it can be done directly
using `docker container run` command:

```shell
docker container run -p 8000:8000 jossemargt/cms-sao
```

Or it can be run using `docker-compose up` with a `docker-compose.yml` file
similar to the one on the project root.

### Configuration

All the intrinsic configurations can be overridden via the `config.properties`
file in Sao's working directory (you could use `config.properties.example` as
guide) or as environment variables with the `SAO_` prefix. As for example in
order to override the `datasource.host` value, you could start the Docker
container with the following syntax:

```shell
docker container run -p 8000:8000 -e 'SAO_CMS_DATASOURCE_HOST=10.10.37.10' jossemargt/cms-sao
```

When CMS Sao is executed using `docker-compose` the override values can be provided directly
in the `docker-compose.yml` or `docker-compose.override.yml` files (within the `environment`
section).

The most relevant properties are:

Property name | Default value | Description
--- | --- | ---
server.port | 8000 | The port where the application will listen for incoming requests
server.log.request | false | When set as true the application will log each incoming request to the STDOUT
server.error.tracedump | false | When set as true the application will include the error trace in failure responses
cms.url | http://localhost/ | The URL where Sao can stablish communication with CMS
cms.datasource.name | cmsdb | CMS PostgreSQL schema name
cms.datasource.username | cmsuser | CMS PostgreSQL datasource username
cms.datasource.password | | CMS PostgreSQL datasource password
cms.datasource.host | 127.0.0.1 | CMS PostgreSQL host network address
cms.datasource.port | 5432 | CMS PostgreSQL instance port
cms.datasource.sslmode | disable | CMS PostgreSQL [ssl mode](https://www.postgresql.org/docs/9.1/libpq-ssl.html) (valid values: require, disable, verify-ca, verify-full)
documentsource.name | cmsdb | MongoDB database name
documentsource.usernaname | cmsuser | MongoDB datasource username
documentsource.password | | MongoDB datasource password
documentsource.host | 127.0.0.1 | MongoDB host network address
documentsource.port | 27017 | MongoDB port

## License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE)
file for details.
