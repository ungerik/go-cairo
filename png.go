package cairo

// #cgo CFLAGS: -Wall -O2
// #cgo pkg-config: cairo
// #include <stdio.h>
// #include <stdlib.h>
// #include <strings.h>
// #include <cairo.h>
//
// #define GO_CAIRO_VECTOR_INIT_SIZE 64
//
// typedef struct go_cairo_vector {
// 	char *buf;
// 	size_t cap;
// 	size_t len;
// } go_cairo_vector;
//
// void go_cairo_vector_free(go_cairo_vector *v) {
// 	if (v->buf) {
// 		free(v->buf);
// 	}
// 	free(v);
// }
//
// void go_cairo_vector_init(go_cairo_vector *v) {
// 	v->len = 0;
// 	v->cap = 0;
// 	size_t cap = GO_CAIRO_VECTOR_INIT_SIZE;
// 	char *buf = malloc(sizeof(char) * cap);
// 	if (buf) {
// 		v->buf = buf;
// 		v->cap = cap;
// 	}
// }
//
// go_cairo_vector *go_cairo_vector_new() {
// 	go_cairo_vector *v = malloc(sizeof(go_cairo_vector));
// 	go_cairo_vector_init(v);
// 	return v;
// }
//
// static int go_cairo_vector_grow(go_cairo_vector *v, size_t len) {
// 	int status = 0;
// 	if (v->cap <= v->len+len) {
// 		unsigned int cap;
// 		char* buf;
// 		if (v->cap == 0) {
// 			cap = GO_CAIRO_VECTOR_INIT_SIZE + len;
// 			buf = malloc(sizeof(char) * cap);
// 		} else {
// 			cap = (v->cap*2) + len;
// 			buf = realloc(v->buf, sizeof(char) * cap);
// 		}
// 		if (buf) {
// 			v->buf = buf;
// 			v->cap = cap;
// 		} else {
// 			status = 1;
// 		}
// 	}
// 	return status;
// }
//
// static int go_cairo_vector_append(go_cairo_vector *v, const unsigned char *data, size_t len) {
// 	int status = go_cairo_vector_grow(v, len);
// 	if (status == 0) {
// 		memcpy(&v->buf[v->len], data, len);
// 		v->len += len;
// 	}
// 	return status;
// }
//
// cairo_status_t go_cairo_writer(void *closure, const unsigned char *data, unsigned int length) {
// 	cairo_status_t status;
// 	go_cairo_vector *v = (go_cairo_vector*)closure;
// 	switch (go_cairo_vector_append(v, data, (size_t) length)) {
// 	case 0:
// 		status = CAIRO_STATUS_SUCCESS;
// 		break;
// 	case 1:
// 		status = CAIRO_STATUS_NO_MEMORY;
// 		break;
// 	default:
// 		status = CAIRO_STATUS_WRITE_ERROR;
// 		break;
// 	}
// 	return status;
// }
//
// cairo_status_t go_cairo_write_surface_to_vector(cairo_surface_t *surface, go_cairo_vector *vec)
// {
// 	cairo_status_t status = cairo_surface_write_to_png_stream(surface, go_cairo_writer, vec);
// 	cairo_surface_destroy(surface);
// 	return status;
// }
import "C"

import "unsafe"

func (self *Surface) WriteToPNGStream() ([]byte, Status) {
	vec := C.go_cairo_vector_new()
	status := Status(C.go_cairo_write_surface_to_vector(self.surface, vec))
	buf := C.GoBytes(unsafe.Pointer(vec.buf), C.int(vec.len))
	C.go_cairo_vector_free(vec)
	return buf, status
}
