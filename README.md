# build the app
docker run \
    -v /var/run/docker.sock.raw:/var/run/docker.sock \
    -v $PWD:/workspace -w /workspace \
    -u501 \
    buildpacksio/pack build go-sample \
        --builder=gcr.io/buildpacks/builder:v1

# run the app
docker run -p8080:8080 go-sample

#build with custom buildpacks
pack config default-builder cnbs/sample-builder:bionic
pack build go-sample --path . --buildpack ./go-buildpack

# build with environment variable
pack build go-sample --path . --env="PORT:8081" --buildpack ./go-buildpack
