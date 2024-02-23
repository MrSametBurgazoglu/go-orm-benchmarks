# Go ORM Benchmarks

This repository contains benchmarks for various Go ORM libraries. The benchmarks are run against a Postgres database and measure the time it takes to perform various operations.

## Running the benchmarks

First, you need to install the dependencies:

```bash
go install
```

Then, you can run the docker-compose to up databases:

```bash
docker-compose up -d
```

Finally, you can run the benchmarks:

```bash
go test -benchmem -bench . .\benchmarks
```
