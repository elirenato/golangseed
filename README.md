# Golang seed project created with the following features/capabilities:

* [Revel framework] (https://revel.github.io/) is a high-productivity web framework for the [Go language](http://www.golang.org/). Feature of Revel framework that have been used by this project: Hot Code Reload, Routing, Internationalization and Unitest for now.
* Authentication using [JSON Web Token] (https://github.com/elirenato/jwt) forked originaly from [https://github.com/benjic/jwt] (https://github.com/benjic/jwt).
* Migration tool: [pgmgr] (https://github.com/rnubel/pgmgr) for database updates and changelog.

* PS 1: This project is still in progress, be aware that it is unfinshed and maybe with bugs.
* PS 2: Import the Golang.postman_collection.json file of the repository to the Postman Chrome Extension, it has consumers examples of the Golang`s endpoints.

### Database

The seed project uses PostgreSQL and it will connect to the database using the username golangseed and password 123456, so, create a new rule with these credentials, after that run the commands below to initialize the database:

* run "pgmgr db create" to database.

* run "pgmgr db migrate" to update the database with the necessary tables.

### Start the web server:

   ./run.sh or revel run github.com/elirenato/golangseed

### Go to http://localhost:9000/ and you'll see:

{
  "id": "unauthorized",
  "message": "Invalid or token is not provided"
}

All routes are protected using Json web token

### Run unit tests:

   ./unittest.sh

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

# TODO
* The frontend (https://github.com/elirenato/login-flow) that have been used to work with this backend was using bcrypt (https://github.com/dcodeIO/bcrypt.js) to send encrypted password. Even that we are going to use https, its nice to have send password encrypted.
