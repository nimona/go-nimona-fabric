
```sh
LOG_LEVEL=INFO \
NIMONA_BIND_ADDRESS=0.0.0.0:18000 \
NIMONA_DEBUG_METRICS_PORT=6060 \
NIMONA_PEER_PRIVATE_KEY=ed25519.prv.nzNG38rXKFGTPqfNxBNGUyte2hpGgJP77br9GmUQiQ3e9HpqUMFuSavRfz5K5MWhwZskHr48uDD9X8Y2hw3Yg1q \
go run ./main.go serve 10m.data
```

```sh
LOG_LEVEL=INFO \
NIMONA_DEBUG_METRICS_PORT=6061 \
NIMONA_BIND_ADDRESS=0.0.0.0:18001 \
NIMONA_PEER_PRIVATE_KEY=ed25519.prv.574sFU7bQUpCsYUnsw9RF4fveUBEYAfbu1DbmVnZ4ieuFkRyPWudHYeHeesrYRQJf2qFh2V5b98AmMseT7VGGEcm \
go run main.go get oh1.8xRwACh8PichMTozpo8UsttQVWjzE9RMNASkDCJjYY4L
```