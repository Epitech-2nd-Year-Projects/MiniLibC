NAME	= 	libasm.so

RM		?=	rm -f

all: build

build:
	@$(MAKE) -C lib

clean:
	@$(MAKE) -C lib clean

fclean: clean
	$(RM) $(NAME) 

re: fclean all

.PHONY: all build clean fclean re
