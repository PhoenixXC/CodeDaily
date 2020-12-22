/**
 * Created by xuanc on 2020/12/22.
 */

#include "log_lib.h"

extern int printf(const char * str, ...);

int LogExtHandle(const char * log) {
  printf(" --> ext handle, the log is { %s }, mock write log to file\n", log);
  return 0;
}