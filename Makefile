BIN_NAME=mydocker
SRC=.

.PHONY: all build run

all:build

build:
	go build -o $(BIN_NAME) $(SRC)

run:
	./$(BIN_NAME) run -it $(filter-out $@,$(MAKECMDGOALS))

# 防止将 run 作为文件名处理
%:
	@: