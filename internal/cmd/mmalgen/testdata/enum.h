// +build ignore

#ifndef MMAL_PORT_H
#define MMAL_PORT_H

#ifdef __cplusplus
extern "C" {
#endif

typedef enum
{
   MMAL_PORT_TYPE_UNKNOWN = 0,          /**< Unknown port type */
   MMAL_PORT_TYPE_CONTROL,              /**< Control port */
   MMAL_PORT_TYPE_INPUT,                /**< Input port */
   MMAL_PORT_TYPE_OUTPUT,               /**< Output port */
   MMAL_PORT_TYPE_CLOCK,                /**< Clock port */
   MMAL_PORT_TYPE_INVALID = 0xffffffff  /**< Dummy value to force 32bit enum */

} MMAL_PORT_TYPE_T;

#ifdef __cplusplus
}
#endif

#endif /* MMAL_PORT_H */
