#!/bin/bash

DBSTRING="host=db-marketplace-even user=root password=root dbname=root sslmode=disable"

goose postgres "$DBSTRING" up