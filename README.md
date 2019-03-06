# Bailey - gRPC CSV uploader

## Requirements

* Go 1.12
* Go Modules
* Docker
* Docker Compose
* `make`

# Running the Application

Clone the repository to any location on your system, Go 1.12 is required to
build as this use go modules for dependency management.

If you have `make` installed you can build and run the server component of the
application by running:

```
make run
```

This will spawn a database and gRPC server. To stop theses service run:

```
make down
```

If this is the first time running the application you will need to also run
database migrations:

```
make migrate_up
```

You will also need to build the CLI tool component, again you can use `make`:

```
make csv.consumer.cmd
```

A binary named `csv.consumer.cmd` will be placed into a `bin` directory located
in the project root.

You can now call this binary by passing it a path to a CSV file to consume, for
example:

```
bin/csv.consumer.cmd -csv ~/path/to/data.csv
```

You should see the following output:

```
2019/03/06 21:49:54 quit
2019/03/06 21:49:54 Processed: 103
```

If you wish you can also build the server binary and also run that locally:

```
make client.store.svc
bin/client.store.svc
```
