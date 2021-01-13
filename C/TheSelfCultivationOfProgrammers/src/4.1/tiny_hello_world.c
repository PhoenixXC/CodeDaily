/**
 * Created by xuanc on 2020/12/23.
 */

/*
 * gcc error:
 * @see https://blog.csdn.net/neuq_jtxw007/article/details/78112672
 *
 * the compile and link command:
 * - gcc -c -fno-builtin -m32 tiny_hello_world.c
 * - ld -static -e nomain -o tiny_hello_world.out -m elf_i386 tiny_hello_world.o
 * - ./tiny_hello_world.out
 */

const char *kStr = "Hello World!\n";

void print() {
  asm("movl $13, %%edx \n\t"
      "movl %0,  %%ecx \n\t"
      "movl $0,  %%ebx \n\t"
      "movl $4,  %%eax \n\t"
      "int $0x80 \n\t"
  ::"r"(kStr):"edx", "ecx", "ebx");
}

void exit() {
  asm("movl $42, %ebx \n\t"
      "movl $1,  %eax \n\t"
      "int  $0x80     \n\t");
}

void nomain() {
  print();
  exit();
}