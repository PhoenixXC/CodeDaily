/**
 * Created by xuanc on 2020/12/22.
 */

#include <stdio.h>
#include "log_lib.h"

/**
 * set custom log format, optional
 */
void SetLogFormat(char * format_str, enum LOG_LEVEL log_level) {
  sprintf(format_str, "%s", "[custom log format] %s - %s > %s");
}

int main(void) {

  Log(INFO, "info log");
  Log(WARN, "warn log");
  Log(ERROR, "error log");
  Log(DEBUG, "debug log");

  return 0;
}