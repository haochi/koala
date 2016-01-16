if [[ "$PWD" != *scripts ]]
then
    cd scripts
    go build
    ./scripts
    rm ./scripts
fi
