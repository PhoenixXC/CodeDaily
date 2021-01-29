# this is a make file for a kernel object
# see online for more information

# will build "hello.ko"
obj-m += hello.o

hello-y := \
	src/main.o \

all:
	make -C /lib/modules/$(shell uname -r)/build/ M=$(PWD) modules

clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
