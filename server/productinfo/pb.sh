#!/bin/bash
out_path=`pwd`
#echo $out_path
protoc -I$out_path/ --go_out=plugins=grpc:$out_path/ $out_path/ecommerce/product_info.proto