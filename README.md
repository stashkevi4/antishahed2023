# Anti-Shahed sound detecting system

## Local run

1. Make sure that [Go](https://www.docker.com/products/docker-desktop/) with version 1.17 or higher is **installed**
2. Make sure that [Docker](https://www.docker.com/products/docker-desktop/)  with version 4.5.1 or higher is **installed** and **running**
3. Run `go mod tidy` (only for the first time after repo downloaded)
4. Launch Kafka from console: `docker-compose up zookeeper kafka`
5. Launch detector (in another console): `go run ./cmd/detector`
6. Open `web/index.html` file in browser, there should be a message in browser's console: `WebSocket connection established`
7. Launch the generator (in another console): `go run ./cmd/generator`
8. Go back to browser and watch how shahed is flying :D
9. After run, stop all consoles
10. Stop docker-compose with clean-up volumes: `docker-compose down -v`

To change shahed's path, you can go to `configs/shahed-vectors.json` and add/change vectors.

Also, you can try to add more config files with shaheds and run more generators simultaneously (but change file name in `cmd/generator/generator.go`)