#!/bin/bash


dst=gdaemon
source_dir=`pwd`
mkdir -p build

case "$1" in

mini)
    go build -o build/$dst -ldflags "-s -w"
;;

rpm)
 version=$(cat ${dst}.spec  | grep Version | awk '{print $2}')
 echo "build rpm package version $version"
 rm -rf build/*
 rpmbuild_dir=/root/rpmbuild
 
 go build -o $dst -ldflags "-s -w"

 rpmdev-setuptree

 mkdir -p ${rpmbuild_dir}/SOURCES/${dst}-${version}
 cp ${dst} ${rpmbuild_dir}/SOURCES/${dst}-${version}
 cd ${rpmbuild_dir}/SOURCES/
 tar czvf ${dst}-${version}.tar.gz ${dst}-${version}

 cd ${source_dir}
 cp ${dst}.spec ${rpmbuild_dir}/SPECS/${dst}.spec
 rpmbuild -bb ${rpmbuild_dir}/SPECS/${dst}.spec
 cp ${rpmbuild_dir}/RPMS/x86_64/*.rpm build/

 rm -rf ${rpmbuild_dir}

 go clean 
;;

*)
    echo "default go build"
    go build -o ${dst}
esac

