SRC	= strlen.asm \
	  strchr.asm \
	  strrchr.asm \
	  memset.asm \
	  strcmp.asm \
	  strncmp.asm \
	  strcasecmp.asm \
	  strcspn.asm \
	  strpbrk.asm \
	  strstr.asm  \
	  memcpy.asm  \
	  memmove.asm \
	  index.asm

RM	?=	rf -f	

OBJ	= $(SRC:.asm=.o)

NAME	= libasm.so

ASM			= nasm
ASM_FLAGS	= -f elf64

LD		= ld -shared

all: $(NAME)

$(NAME): $(OBJ)
	$(LD) -o $@ $^
	@mv $(NAME) ..

%.o: %.asm
	$(ASM) $(ASM_FLAGS) $< -o $@

clean:
	$(RM) $(OBJ)

.PHONY: all clean
