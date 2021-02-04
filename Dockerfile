FROM golang
RUN mkdir -p /usr/src/GoWebProject/
WORKDIR /usr/src/GoWebProject/
COPY . /usr/src/GoWebProject/
CMD ["go", "run", "./cmd/web"]