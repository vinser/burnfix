#!/bin/bash

# set semantic version
[[ $(cat FyneApp.toml) =~ ([0-9]+\.[0-9]+\.[0-9]+) ]] && SEMVER="${BASH_REMATCH[0]}"
sed -i -r 's/burnfix\/v[0-9]+\.[0-9]+\.[0-9]+/burnfix\/v'${SEMVER}'/g' io.github.vinser.burnfix.metainfo.xml
sed -i -r 's/tags\/v[0-9]+\.[0-9]+\.[0-9]+/tags\/v'${SEMVER}'/g' io.github.vinser.burnfix.yml
sed -i -r 's/tag\/v[0-9]+\.[0-9]+\.[0-9]+/tag\/v'${SEMVER}'/g' README.md
sed -i -r 's/tag\/v[0-9]+\.[0-9]+\.[0-9]+/tag\/v'${SEMVER}'/g' docs/README.md
sed -i -r 's/^  Build = [0-9]+//g' FyneApp.toml
echo $SEMVER