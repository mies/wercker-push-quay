#!/bin/bash

# Download kubernetes to ./kubernetes.tar.gz and extract it to ./kubernetes.
main() {
  local destination="./kubernetes.tar.gz"
  local version="0.15.0"

  curl -L https://github.com/GoogleCloudPlatform/kubernetes/releases/download/v${version}/kubernetes.tar.gz > "${destination}";

  tar xvfz "${destination}" -C ./
}

main $@;
