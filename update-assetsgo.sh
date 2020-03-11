#!/bin/bash

GEN_GO="gen.go"
trap "{ rm -f ${GEN_GO}; }" EXIT

cat > ${GEN_GO} << EOF
package main

import "github.com/zserge/lorca"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("main", "assets.go", "resources")
}
EOF

go run "${GEN_GO}"
