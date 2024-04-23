git pull origin main
docker stop html-pdf
docker rm html-pdf
docker rmi html-pdf
docker build -t html-pdf:latest .
docker run -d -p 5000:5000 --name html-pdf html-pdf:latest