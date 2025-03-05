CC = gcc
CFLAGS = -Wall -Wextra -std=c11 -O2 -lm
TARGET = stats
SRC = main.c

all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC)

run: $(TARGET)
	./$(TARGET) data.txt

clean:
	rm -f $(TARGET)
