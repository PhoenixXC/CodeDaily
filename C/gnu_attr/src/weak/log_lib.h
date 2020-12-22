/**
 * Created by xuanc on 2020/12/22.
 */

#ifndef GNU_ATTR_SRC_WEAK_LOG_LIB_H_
#define GNU_ATTR_SRC_WEAK_LOG_LIB_H_

enum LOG_LEVEL{
  TRACE,
  INFO,
  WARN,
  ERROR,
  DEBUG
};

/**
 * optional, handle log, like write to disk or other things
 */
int LogExtHandle(const char *);

void SetLogFormat(char * format_str, enum LOG_LEVEL log_level);

const char* LogLevelEnumToStr(enum LOG_LEVEL log_level);

void Log(enum LOG_LEVEL log_level, const char * msg);

#endif //GNU_ATTR_SRC_WEAK_LOG_LIB_H_
