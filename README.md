## BUILD
```bash
docker build -t html-pdf:latest .
```

## RUN

```bash
docker run -d -p 5000:5000 --name html-pdf html-pdf:latest
```

## USAGE:
```bash
curl -X POST -d 'html=<html><head><title>Test</title></head><body><h1>Hello, World!</h1></body></html>' http://localhost:5000/generate-html-pdf -o index.pdf
```