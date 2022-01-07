# Short URL

To verify the setup, run the project tests suite: 

```bash
go test ./... -count=1
```

In case of failures, review carefully the setup steps listed below (GCP or local) .

Then, build and start Short URL service:

```bash
go build . && ./short-url
```

## Setup

### Local 

```bash
brew install redis
brew services restart redis
redis-server
```

To ping redis:

```bash
redis-cli ping
```

To stop redis:
```bash
brew services stop redis
```

### GCP

## Reference
* [Let's build a URL shortener in Go - with Gin & Redis](https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-iv-forwarding/)
* [Session based authentication in Go](https://github.com/sohamkamani/go-session-auth-example)


docker run -it docker.io/library/shorturl:latest /bin/bash