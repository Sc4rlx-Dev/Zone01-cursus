objects = src/main.o src/system.o src/auth.o src/myfunct.o

atm : $(objects)
	cc -o atm $(objects)

main.o : src/header.h
kbd.o : src/header.h
command.o : src/header.h
display.o : src/header.h
insert.o : src/header.h
search.o : src/header.h
files.o : src/header.h
utils.o : src/header.h

fclean : clean
	rm -f atm

clean :
	rm -f $(objects)

help:
	@echo "Available targets:"
	@echo "  atm      - Build the ATM executable."
	@echo "  clean    - Remove object files."
	@echo "  fclean   - Remove object files and the ATM executable."
	@echo "  help     - Display this help message."