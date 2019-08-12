<!-- Autogenerated by Typical-Go. DO NOT EDIT. -->

# Typical-RESTful-Server

Example of typical and scalable RESTful API Server for Go

## Getting Started

This is intruction to start working with the project:
0. Install [Go](https://golang.org/doc/install) or using homebrew if you're using macOS `brew install go`


## Usage

There is no specific requirement to run the application. 

## Build Tool

Use `./typicalw` to execute development task

## Make a release

Use `./typicalw release` to make the release. More information check [here](https://typical-go.github.io/release.html)

## Configurations

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|	
|APP_ADDRESS|String|:8089|true||	

Echo Server with Logrus

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|	
|SERVER_DEBUG|True or False|false|||	

Postgres Database

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|	
|PG_DBNAME|String||true||	
|PG_USER|String|postgres|true||	
|PG_PASSWORD|String|changeme|true||	
|PG_HOST|String|localhost|||	
|PG_PORT|Integer|5432|||	
|PG_MIGRATIONSRC|String|scripts/migration|||	


