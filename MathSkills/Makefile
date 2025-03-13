CC = gcc
CFLAGS = -Wall -Wextra -Werror -lm
TARGET = stats
SRC = main.c

all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC)

run: $(TARGET)
	./$(TARGET) data.txt

clean:
	rm -f $(TARGET)
