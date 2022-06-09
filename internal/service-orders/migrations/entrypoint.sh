#!/bin/bash

DBSTRING="host=db-orders user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up