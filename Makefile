ifeq ($(OS),Windows_NT)
    detected_OS := Windows
else
    detected_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

LIB_PATH = embedder/cmake-build-debug
LIB_PATH_OSX = embedder/libs/osx
LIB_PATH_LINUX = embedder/libs/linux

PYTORCH_URL_OSX = https://download.pytorch.org/libtorch/cpu/libtorch-macos-1.5.0.zip
PYTORCH_URL_LINUX = https://download.pytorch.org/libtorch/cpu/libtorch-cxx11-abi-shared-with-deps-1.5.0%2Bcpu.zip

MKL_URL_OSX = https://github.com/intel/mkl-dnn/releases/download/v0.21/mklml_mac_2019.0.5.20190502.tgz
MKL_URL_LINUX = https://github.com/intel/mkl-dnn/releases/download/v0.21/mklml_lnx_2019.0.5.20190502.tgz

BUILD_DIR = build
BIN ?= kawaii_search
IMAGES_PATH = images
MODEL = scripts/model.pt
CONFIG = scripts/config.json

ifeq ($(detected_OS),Darwin)
	LIBRARY_PATH = $(LIB_PATH_OSX)/libtorch/lib:$(LIB_PATH_OSX)/mklml/lib:$(LIB_PATH)
endif
ifeq ($(detected_OS),Linux)
	LIBRARY_PATH = $(LIB_PATH_LINUX)/libtorch/lib:$(LIB_PATH_LINUX)/mklml/lib:$(LIB_PATH)
endif

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
 	# go test ./...
	go build -o $(BUILD_DIR)/$(BIN)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/*

.PHONY: run_create_db
run_create_db:
	LD_LIBRARY_PATH=$(LIBRARY_PATH) MODEL=$(MODEL) CONFIG=$(CONFIG) ./$(BUILD_DIR)/$(BIN) create_db $(IMAGES_PATH)

.PHONY: download_libs_osx
download_libs_osx:
	mkdir -p $(LIB_PATH_OSX)
	cd $(LIB_PATH_OSX);curl -L -O $(PYTORCH_URL_OSX)
	cd $(LIB_PATH_OSX);unzip libtorch-macos-1.5.0.zip
	# cd $(LIB_PATH_OSX);curl -L -O $(MKL_URL_OSX)
	# cd $(LIB_PATH_OSX);tar zxvf mklml_mac_2019.0.5.20190502.tgz
	# cd $(LIB_PATH_OSX);mv mklml_mac_2019.0.5.20190502 mklml

.PHONY: download_libs_linux
download_libs_linux:
	mkdir -p $(LIB_PATH_LINUX)
	cd $(LIB_PATH_LINUX);curl -L -O $(PYTORCH_URL_LINUX)
	cd $(LIB_PATH_LINUX);unzip libtorch-cxx11-abi-shared-with-deps-1.5.0%2Bcpu.zip
	# cd $(LIB_PATH_LINUX);curl -L -O $(MKL_URL_LINUX)
	# cd $(LIB_PATH_LINUX);tar zxvf mklml_lnx_2019.0.5.20190502.tgz
	# cd $(LIB_PATH_LINUX);mv mklml_lnx_2019.0.5.20190502 mklml
