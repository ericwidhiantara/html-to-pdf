## CLONE THIS REPO
```bash
git clone https://github.com/thxrhmn/html-to-pdf.git
```

## RUN GOTENBERG
```bash
docker run --rm -p 4000:3000 -d gotenberg/gotenberg:8
```

## COPY / FILL .ENV VARIABLE
```bash
cp .env.example .env
```

## BUILD
```bash
docker build -t html-to-pdf:latest .
```

## RUN
```bash
docker run -d -p 5000:5000 --name html-to-pdf html-to-pdf:latest
```

## USAGE:
```bash
curl -X POST -d 'html=<html><head><title>Test</title></head><body><h1>Hello, World!</h1></body></html>' http://localhost:5000/generate-html-pdf -o index.pdf
```