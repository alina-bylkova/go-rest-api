FROM redis:6-alpine
WORKDIR /usr/src/app
COPY go.mod ./
RUN go get
COPY . .
EXPOSE 8080
CMD ["make", "dev"]