# Anti-Shahed sound detecting system

## Local run

1. Run `go mod tidy` (only for the first time after repo downloaded)
2. Make sure that [Docker](https://www.docker.com/products/docker-desktop/) is **installed** and **running**
3. Launch Kafka from console: `docker-compose up zookeeper kafka`
   3a. After run, stop and clean-up volumes: `docker-compose down -v`
4. Launch detector (in another console): `go run ./cmd/detector`
5. Open `web/index.html` file in browser, there should be a message in browser's console: `WebSocket connection established`
6. Launch the generator (in another console): `go run ./cmd/generator`
7. Go back to browser and watch how shahed is flying :D

To change shahed's path, you can go to `configs/shahed-vectors.json` and add/change vectors.

Also, you can try to add more config files with shaheds and run more generators simultaneously (but change file name in `cmd/generator/generator.go`)