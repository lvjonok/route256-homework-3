#!/bin/bash

DBSTRING="host=db user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up