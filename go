#!/bin/bash

[ $# = 1 ] || { echo "Usage: $0 version" >&2; exit 1; }

for x in \
    nvidia-installer \
    nvidia-modprobe \
    nvidia-persistenced \
    nvidia-settings \
    nvidia-xconfig;
  do
      (
        set -e
        cd "$x"
        rm -r *
        tar --strip-components=1 -xvf /mnt/release/builds/release/display/x86_64/$1/$x-$1.tar.bz2
        git add -A
        git commit -m $1
        git tag $1
      )
done
