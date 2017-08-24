Find Tribonacci Number
======================

## Overview

What is tribonacci number? See http://mathworld.wolfram.com/TribonacciNumber.html

This is example of Golang program to find tribonacci sequence value for given index (N).

The algorithm uses linear O(n) implementation and works with all natural numbers.

For N > 1,000,000 it takes some time to calculate tribonacci number.

Program runs as service available on URL http://localhost:8081/v1/trib/{n} where {n} is placeholder for N-index of tribonacci sequence.

## Run

Use **Docker** and **docker-compose** to run program.

```sh
docker-compose build
docker-compose up
```
Before running service unit tests and acceptance tests executed.

## Example

http://localhost:8081/v1/trib/1000

```sh
{"result":"443382579490226307661986241584270009256355236429858450381499235934108943134478901646797270328593836893366107162717822510963842586116043942479088674053663996392411782672993524690287662511197858910187264664163782145563472265666010074477859199789932765503984125240893"}
```

## License

MIT licensed. See the LICENSE file for details.