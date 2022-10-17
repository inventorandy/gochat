#!/bin/bash
mockgen -source=./platform/services/account/interfaces/repository.go \
        -destination=./platform/services/account/interfaces/mocks/repository.mock.go -package=mocks

mockgen -source=./platform/services/conversation/interfaces/repository.go \
        -destination=./platform/services/conversation/interfaces/mocks/repository.mock.go -package=mocks