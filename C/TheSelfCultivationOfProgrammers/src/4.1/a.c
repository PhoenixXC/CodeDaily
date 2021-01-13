/**
 * Created by xuanc on 2020/12/23.
 */

extern int shared;
extern void swap(int *, int *);

/**
 * 使用 ld 链接目标文件的时候可能会出现一些问题：
 * @see https://zhuanlan.zhihu.com/p/150793679
 */
void SysExit() {
  asm( "movq $66,%rdi \n\t"
       "movq $60,%rax \n\t"
       "syscall \n\t");
}

int main() {
  int a = 100;
  swap(&a, &shared);
  SysExit();
  return 0;
}