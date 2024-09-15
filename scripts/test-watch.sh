#! /usr/bin/env nix-shell
#! nix-shell -i bash -p bash

watchexec -e go go test ./...