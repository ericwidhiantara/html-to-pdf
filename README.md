## CLONE THIS REPO
```bash
git clone https://github.com/thxrhmn/html-to-pdf.git
```

## COPY / FILL .ENV VARIABLE
```bash
cp .env.example .env
nano .env
# change with your ip
GOTENBERG_API=http://YOUR-IP:4000 # http://127.0.0.1:4000
```

## RUN
```bash
docker compose up -d
```

## USAGE:
```bash
curl -X POST -d 'html=<html><head><title>Test</title></head><body><h1>Hello, World!</h1></body></html>' http://0.0.0.0:5000/generate-html-pdf -o index.pdf
```