/**
 * Created by xuanc on 2020/12/23.
 */

int shared = 1;

void swap(int* a, int* b) {
  *a ^= *b ^= *a ^= *b;
}