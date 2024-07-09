#!/usr/bin/env bash

rg "draft = true" --files-with-matches ./web/content | xargs sed -i 's/draft = true/draft = false/g'
