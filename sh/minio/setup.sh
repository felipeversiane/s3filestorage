#!/bin/sh

sleep 10

mc alias set myminio http://localhost:9000 minioadmin minioadmin

mc mb myminio/filestorage

mc policy set-json /policy.json myminio/filestorage
