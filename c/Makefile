
CC=gcc
CFLAGS=-O2 -std=c99

%.o: %.c 
	$(CC) -c -o $@ $< $(CFLAGS)

all: cwrap.o
	ar rcs libcwrap.a $<


.PHONY: clean

clean:
	rm -f *.o *.a
