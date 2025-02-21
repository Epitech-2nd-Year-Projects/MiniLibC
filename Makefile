SRC	=

OBJ	= $(SRC:.asm=.o)

NAME	= libasm.so

NASM	= nasm -f elf64

LD		= ld -shared

all: $(NAME)

$(NAME): $(OBJ)
	$(LD) -o $@ $^

%.o: %.asm
	$(NASM) $< -o $@

clean:
	rm -f $(OBJ)

fclean: clean
	rm -f $(NAME)

re: fclean all

.PHONY: all clean fclean re
