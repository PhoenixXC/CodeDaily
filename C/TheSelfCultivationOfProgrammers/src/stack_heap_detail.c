#include <elf.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
  int *p = (int *)argv;
  int i;
  Elf32_auxv_t *aux;

  printf("Arguement count: %d\n", *(p - 1));

  for (i = 0; i < *(p - 1); i++) {
    printf("Arguement %d: %s\n", i, *(p + i));
  }

  p += i;
  p++;

  printf("Environment:\n");
  while (*p) {
    printf("%s\n", *p);
    p++;
  }

  p++;

  printf("Auxiliary Vectors: \n");
  aux = (Elf32_auxv_t *)p;
  while (aux->a_type != AT_NULL) {
    printf("Type: %02d Value: %x\n", aux->a_type, aux->a_un.a_val);
    aux++;
  }

  return 0;
}