# **kafrest**

<!-- BADGES -->

<div align="center">

[![BADGE_DOCKER_PULLS](https://img.shields.io/docker/pulls/mauricioromagnollo/kafrest.svg)](https://hub.docker.com/r/mauricioromagnollo/kafrest)
![BADGE_DOCKER_IMAGE_SIZE](https://img.shields.io/docker/image-size/mauricioromagnollo/kafrest/latest)
![BADGE_GOLANG](https://img.shields.io/badge/Golang-grey?style=flat&logo=go)
![BADGE_GO_VERSION](https://img.shields.io/github/go-mod/go-version/mauricioromagnollo/kafrest?logo=golang)
![BADGE_LICENSE](https://img.shields.io/github/license/mauricioromagnollo/kafrest)
![BADGE_OPEN_SOURCE](https://badges.frapsoft.com/os/v1/open-source.png?v=103)
![BADGE_ISSUES_OPEN](https://img.shields.io/github/issues/mauricioromagnollo/kafrest?color=green)
![BADGE_ISSUES_CLOSED](https://img.shields.io/github/issues-closed/mauricioromagnollo/kafrest?color=red)

</div>

<!-- BRIEF DESCRIPTION -->
 
> A simple REST API developed with Golang to communicate with a Kafka broker, provided as a Docker image. The motivation is to provide a simple way to produce and consume messages from a Kafka broker, without the need to install any other tool. That way, it is possible to use tools like Insomnia or Postman to test your worker and your api at the same time.

<div align="center">
  <a href="https://hub.docker.com/r/mauricioromagnollo/kafrest">
    <img 
      alt="Dockerhub" 
      src="https://img.shields.io/badge/Dockerhub-cyan"
      width="100"
    >
  </a>
</div>

<br>

## **Installation**

```sh
docker pull mauricioromagnollo/kafrest:latest
```

```sh
docker run -d -p 8888:8080 -e KAFKA_BROKERCONNECT=kafka:29092 mauricioromagnollo/kafrest:latest
```

Or, `docker-compose.yml` example:

```yaml
services:
  kafrest:
    image: mauricioromagnollo/kafrest:latest
    ports:
      - "8888:8080"
    environment:
      KAFKA_BROKERCONNECT: kafka:29092
```

## **Contributing**

Contributions are welcome! Feel free to open an issue if you have a way to improve this project.

<!-- 
https://mauricioromagnollo.github.io/kafrest/contributing
-->

See [CONTRIBUTING](./CONTRIBUTING.md) for more information.

## **Documentation**

The documentation is available at [mauricioromagnollo.github.io/kafrest](https://mauricioromagnollo.github.io/kafrest/).

If you are running the project locally, start the documentation using the command:

```sh
make docs
```

So, the documentation will be available at [localhost:8000](http://localhost:8000).

## **License**

Distributed under the MIT License. See [LICENSE](./LICENSE) file for more information. &copy; Mauricio Romagnollo
