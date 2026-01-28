# Makefile

# Tools
GO := go
QEMU := qemu-system-i386

# Directories
KERNEL_DIR := kernel

# Output
KERNEL_BIN := kernel.bin

# Flags
# -T sets text segment address (0x100000 = 1MB)
# -E sets entry symbol (_start)
LDFLAGS := -T 0x100000 -E _start
GO_ENV := GOOS=linux GOARCH=386 CGO_ENABLED=0

.PHONY: all clean run

all: $(KERNEL_BIN)

$(KERNEL_BIN):
	$(GO_ENV) $(GO) build -ldflags "$(LDFLAGS)" -o $@ ./$(KERNEL_DIR)

run: $(KERNEL_BIN)
	$(QEMU) -kernel $(KERNEL_BIN)

clean:
	rm -f $(KERNEL_BIN)
