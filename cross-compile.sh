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
    "darwin/amd64"
    "freebsd/amd64"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
    "netbsd/amd64"
    "openbsd/amd64"
    "windows/386"
    "windows/amd64"
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

    printf "%-11s %-6s %-30s\n" "Building" "$i/${#platforms[@]}:" "$output_name"
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $build_dir/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    ((i+=1))
done

rm -rf ./"$build_dir"/*.zip
rm -rf ./"$build_dir"/*.tar.gz

i=1
for file in "$build_dir"/*
do
    if echo $file | grep -q "sum"
    then
        continue
    fi

    printf "%-11s %-6s %-30s\n" "Compressing" "$i/${#platforms[@]}:" "$(basename $file)"
    if echo "$file" | grep -q "windows"
    then
        zip --quiet "$file.zip" "$file"
    else
        tar -zcf "$file.tar.gz" "$file"
    fi
    ((i+=1))
done

sums="$build_dir"/sha256sums.txt
sha256sum "$build_dir"/* > $sums
gpg --passphrase $(pass show gpgpass) --batch --yes --detach-sign -a $sums
