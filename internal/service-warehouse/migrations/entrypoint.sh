#!/bin/bash

DBSTRING="host=db-warehouse user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up