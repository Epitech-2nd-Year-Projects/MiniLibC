NAME	= 	libasm.so

RM		?=	rm -f

all: build

build:
	@$(MAKE) -C lib

test: build
	@export LD_LIBRARY_PATH=$(pwd)
	go test -v ./...

clean:
	@$(MAKE) -C lib clean

fclean: clean
	$(RM) $(NAME) 

re: fclean all

.PHONY: all build clean fclean re
