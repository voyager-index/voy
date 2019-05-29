#!/usr/bin/env bash
# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}
build_dir="build"

if [ ! -d $build_dir ]
then
    mkdir $build_dir
fi

platforms=(
    "windows/amd64"
    "windows/386"
    "darwin/amd64"
    "linux/amd64"
    "linux/386"
    "linux/arm"
    "linux/arm64"
)

i=1
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building $i/${#platforms[@]}: $output_name"
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $build_dir/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    ((i+=1))
done
