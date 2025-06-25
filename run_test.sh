#!/bin/bash
docker run --rm -t -v "$(pwd)":/workspace -w /workspace dart:2 /bin/bash -c "\
  cd tool && dart pub get && cd .. && \
  dart tool/bin/test.dart chap06_parsing --interpreter go/lox
  # ./go/lox
"  
