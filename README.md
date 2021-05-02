# Simple deep.ai client

# Introduction

Simple api client for https://deepai.org/machine-learning-model/colorizer image colorizer. Register
on https://deepai.org and fetch accessKey

## How to run in shell
1. Fetch DEEP_AI_KEY from https://deepai.org and store this key as DEEP_AI_KEY env
2. Go to cmd folder
3. go build
4. Run as 
```shell
make build
DEEP_AI_KEY=123 ./main {PATH_TO_FILE_TO_COLORIZE}
```

