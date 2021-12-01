# HTTP tool

This tool makes http requests and prints the address of the request along with the MD5 hash of the response.

## Installation

You must have Go installed and configured properly for your computer. Please see
https://golang.org/doc/install

## Usage

First you need to build the app
```
go build
```

Then you can run it
```
./myhttp adjust.com google.com facebook.com 

// example output:
http://adjust.com dbc37763cc539c958bf7ab245d7681c9
http://google.com eebc25991108806f1ce47fb1ec4e3d18
http://facebook.com d1676fdcb2c4c46e02aa92257dc93803
```

The tool performs the requests in parallel so that the tool can complete sooner. \
You can limit the number of parallel requests using the flag ```-parallel```.
Default is 10 if the flag is not provided.

```
./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny

// example output:
http://adjust.com dbc37763cc539c958bf7ab245d7681c9
http://google.com da870d02c4c5f70b72c69a135faab422
http://facebook.com 0e848994e56422b06416612a49614d5b
http://yandex.com f216b23997092b1586be4fb729bf3cd3
http://twitter.com 15d79fe790a07a9e229c6d2b2aea437c
http://yahoo.com e3280c6a61ec405a2da0a13469401573
http://reddit.com/r/notfunny dd314050edd6adcf83d2ea89b1df1521
http://reddit.com/r/funny fa435083bb6a5fc7d164f410f039334b
```