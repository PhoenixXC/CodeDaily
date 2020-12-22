/**
 * Created by xuanc on 2020/12/22.
 */

#include <time.h>
#include <stdio.h>
#include "log_lib.h"

#define LOG_FORMAT_LEN 50
#define TIME_STR_LEN 20
#define LOG_STR_LEN 500

__attribute__ ((weak))
void SetLogFormat(char * format_str, enum LOG_LEVEL log_level) {
  sprintf(format_str, "%s", "%s - %s > %s");
}

static __attribute__ ((weakref ("LogExtHandle")))
int vLogExtHandle(const char *);

static void GetTime(char *time_str) {
  time_t now;
  struct tm *tm_now;
  time(&now);
  tm_now = localtime(&now);
  sprintf(time_str, "%02d.%02d.%02d %02d:%02d:%02d",
          tm_now->tm_year + 1900, tm_now->tm_mon + 1, tm_now->tm_mday,
          tm_now->tm_hour, tm_now->tm_min, tm_now->tm_sec);
}

void Log(enum LOG_LEVEL log_level, const char *msg) {
  char log_format[LOG_FORMAT_LEN];
  SetLogFormat(log_format, log_level);
  char time_str[TIME_STR_LEN];
  GetTime(time_str);

  char log_str[LOG_STR_LEN];
  sprintf(log_str, log_format, time_str, LogLevelEnumToStr(log_level), msg);

  printf("%s\n", log_str);
  if (vLogExtHandle) {
    vLogExtHandle(log_str);
  }
}

const char *LogLevelEnumToStr(enum LOG_LEVEL log_level) {
  static const char* level_to_str[] = {
      "TRACE",
      "INFO",
      "WARN",
      "ERROR",
      "DEBUG"
  };
  if (log_level >= sizeof(level_to_str) / sizeof(level_to_str[0])) {
    return "UNKNOWN";
  }
  return level_to_str[log_level];
}
