Go Angular Starter
==================

Demo
----

Demo site available at: http://go-angular-starter.buz.co:3000/

Features
--------

Go Angular Starter is a full stack opinionated template for advanced web applications. It includes:

- Go HTTP Server with Go templates
- Go JSON-RPC API Server
- Angular based Admin Panel single page app with Browserify, ES6 and SASS (using Gulp build tasks)
- Authorisation (sign in, sign up, API key generation)
- DB migrations with Goose

Quick Start
-----------

Install the project to your GOPATH:

```bash
go get github.com/cjdell/go_angular_starter
```

Install Go dependencies:

```bash
go get bitbucket.org/liamstask/goose/cmd/goose
go get github.com/jmoiron/sqlx
go get github.com/gorilla/rpc
go get github.com/gorilla/context
go get github.com/kylelemons/go-gypsy/yaml
```

Set up the database. Out of the box a Postgres SQL database is expected to exist called "go_angular_starter" on the localhost, accessible by the current user. Change the configuration located at:

```bash
vim db/dbconf.yml
```

NOTE: You may wish to alter your pg_hba.conf to allow access via localhost without a password. If so add this line:

```
host    all   all   127.0.0.1/32    trust
```

Run the initial migration to create tables:

```bash
goose up  # Assuming $GOPATH/bin is in your PATH
```

Install NPM / Bower dependencies

```bash
sudo npm -g install gulp

cd $GOPATH/src/github.com/cjdell/go_angular_starter

npm install
bower install
```

Run the Gulp tasks (for Browserify / SASS compilation)

```bash
gulp admin-watch  # Leave this running and open a new terminal
```

Run the HTTP server:

```bash
go run server.go
```

Database
--------

To create a new migration, run:

```bash
goose create MIGRATION_NAME [sql]
```

i.e.

```bash
goose create NewTables sql
```

Then run:

```bash
goose up
```

To undo this migration:

```bash
goose down
```
