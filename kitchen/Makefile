BIN_DIR := bin
KITCHEN_BINARY := $(BIN_DIR)/kitchen
ORDERS_BINARY := $(BIN_DIR)/orders

.PHONY: build

build:
	@echo "Building kitchen..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(KITCHEN_BINARY) ./services/kitchen/main.go
	@echo "Build completed: $(KITCHEN_BINARY)"

	@echo "Building orders..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(ORDERS_BINARY) ./services/orders/main.go
	@echo "Build completed: $(ORDERS_BINARY)"