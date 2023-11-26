# This is just a Makefile to simplify running some of the commands mentioned in the README.

build:
	flatpak-builder --user --force-clean build-dir com.github.vinser.burnfix.yml

install:
	flatpak-builder --user --install --force-clean build-dir com.github.vinser.burnfix.yml

run:
	flatpak run --user com.github.vinser.burnfix
