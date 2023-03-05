#ifndef JSCW_H
#define JSCW_H
#ifdef __cplusplus

#include "JavaScriptCore/JavaScript.h"

extern "C" {
#else

#endif
const char* Version();
#ifdef __cplusplus
}  // extern "C"
#endif
#endif  // JSCW_H
