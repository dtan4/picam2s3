# picam2s3

Send [mjpeg-streamer](https://github.com/jacksonliam/mjpg-streamer) snapshot to S3

## How to build and run

```bash
$ go get -d github.com/dtan4/picam2s3
$ cd $GOPATH/src/github.com/dtan4/picam2s3
$ make
$ bin/picam2s3

# To build binary for Raspberry Pi 1 (ARMv6)
$ make build-raspi
$ bin/picam2s3-raspi
```

## Usage

`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` must be set. Currently AWS region is fixed to `ap-northeast-1`.

```bash
$ picam2s3 <bucket> http://<IP address of Raspberry Pi>:8080/?action=snapshot
```

Snapshot via webcam will be uploaded as `s3://<bucket>/<UNIX time>.jpg`.

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
