# Test project

A http server with s3 upload function.

Requirement:
- Config folder

```bash
mkdir conf
touch conf/config.toml
```

- Mongo:

```bash
docker pull mongo
docker run --name mgserver -d mongo:latest -p 127.0.0.1:27017:27017
```

Create database and collection:
- DB: `testdb`
- Collection: `file_upload`

Build Docker image

```bash
docker build -t testprj:latest -f Dockerfile .
```

Run HTTP server

```bash
docker run -d -p 8000:8000 --name testprj testprj:latest
```

Endpoint:

```
curl -X POST \
  http://localhost:8000/v1/file/s3-upload \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F file=@/Users/ducnt114/Pictures/wallpaper2you_78963.jpg
```