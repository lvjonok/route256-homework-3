#!/bin/bash

DBSTRING="host=db-marketplace user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up