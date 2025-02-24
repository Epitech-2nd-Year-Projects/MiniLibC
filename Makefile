SRC	= src/strlen.asm \
	  src/strchr.asm \
	  src/strrchr.asm

OBJ	= $(SRC:.asm=.o)

NAME	= libasm.so

ASM			= nasm
ASM_FLAGS	= -f elf64

LD		= ld -shared

all: $(NAME)

$(NAME): $(OBJ)
	$(LD) -o $@ $^

%.o: %.asm
	$(ASM) $(ASM_FLAGS) $< -o $@

clean:
	rm -f $(OBJ)

fclean: clean
	rm -f $(NAME)

re: fclean all

.PHONY: all clean fclean re
