VERSION=$(git describe --tags)
gox -output "dist/{{.OS}}_{{.Arch}}_{{.Dir}}" -os="windows darwin linux" -ldflags="-s -w -X main.version=${VERSION}"
cd dist
mv  darwin_386_docker-selector        docker-selector
zip darwin_386_docker-selector        docker-selector     -qm
mv  darwin_amd64_docker-selector      docker-selector
zip darwin_amd64_docker-selector      docker-selector     -qm
mv  linux_386_docker-selector         docker-selector
zip linux_386_docker-selector         docker-selector     -qm
mv  linux_amd64_docker-selector       docker-selector
zip linux_amd64_docker-selector       docker-selector     -qm
mv  windows_386_docker-selector.exe   docker-selector.exe
zip windows_386_docker-selector       docker-selector.exe -qm
mv  windows_amd64_docker-selector.exe docker-selector.exe
zip windows_amd64_docker-selector     docker-selector.exe -qm
cd -
