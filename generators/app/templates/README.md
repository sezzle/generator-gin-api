# Go Generator Gin API

## Package Overviews

* **config**: environment based configuration items
* **gin**: core REST API package with all route handlers
* **gorm**: persistence layer with connection setup for mysql
* **gin_suite_test**: ginkgo packaged test suite for running integration style tests on API endpoints

## Configuration

The `config/localConfig.go` is used as the local configuration file.  This file is gitignored and should stay up to date with the `config/localConfig_sample.go` file.

## Adding Routes

Routes are defined in `gin/routes.go` are initialized in the router on startup of the application.  

## Adding Models

All models are defined in the `gorm` package.  Each model should have its own file and will be auto-migrated on initialization of the database.
