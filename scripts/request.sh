#!/usr/bin/env bash

BaseURL='http://localhost:8080'

curl "$BaseURL/event/core/create" -XPOST

curl "$BaseURL/event/core/fetch/1" -XGET

curl "$BaseURL/event/core/list" -XGET

curl "$BaseURL/event/core/delete/2" -XDELETE

curl "$BaseURL/event/core/update/3" -XPATCH
