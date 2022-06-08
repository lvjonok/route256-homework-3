#!/bin/bash

DBSTRING="host=localhost user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up