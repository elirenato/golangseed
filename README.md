# Golang seed project created using the following features:

* [Revel framework] (https://revel.github.io/) is a high-productivity web framework for the [Go language](http://www.golang.org/).
* [JSON Web Token] (https://github.com/elirenato/jwt) forked originaly from [https://github.com/benjic/jwt] (https://github.com/benjic/jwt).
* Postgresql as the database.
* pgmgr (https://github.com/rnubel/pgmgr) for migration tool.

PS 1: This project is still in progress, be aware that it is unfinshed and maybe with bugs.
PS 2: Import the Golang.postman_collection.json file of the repository to the Postman Chrome Extension, it has consumers examples of the Golang`s endpoints.

### Database 

This project is using Postgresql and https://github.com/rnubel/pgmgr to control the migration.

To initialize the database:
* Create an rule called golangseed with 123456 password.
* Create a database and set the owner as the new golangseed rule. We could use pgmgr db create command but it does not use the rule from the .pgmgr.json file as the owner of the new database.
* Install pgcrypto extension. Run "create extension pgcrypto" from pgadmin or pgsql into the current server.
* run "pgmgr db migrate" to create the basic tables.

### Start the web server:

   revel run github.com/elirenato/golangseed

### Go to http://localhost:9000/ and you'll see:

{
  "id": "unauthorized",
  "message": "Invalid or token is not provided"
}

All routes are protected for know using Json web token

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

