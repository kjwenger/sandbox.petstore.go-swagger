#!/usr/bin/env bash

# latest=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r .tag_name)
# sudo curl -o /usr/local/bin/swagger \
#     -L'#' https://github.com/go-swagger/go-swagger/releases/download/$latest/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
# sudo chmod +x /usr/local/bin/swagger

swagger generate server \
    --spec /com.u14n/development/projects/sandbox/sandbox.petstore/swagger.json \
    --name petstore
swagger generate client \
    --spec /com.u14n/development/projects/sandbox/sandbox.petstore/swagger.json \
    --name petstore
