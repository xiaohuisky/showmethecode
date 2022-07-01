#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, int len) {
    int i;
	for(i=0; i<len; i++) {
		printf("%.2x ", start[i]);
	}
	printf("\n");
}

void show_int(int x) {
	show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x) {
	show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x) {
    show_bytes((byte_pointer) &x, sizeof(void *));
}

int main(int val,  char** argc) {
//	int ival = val;
//	float fval = (float) ival;
//	int *pval = &ival;
//	show_int(ival);
//	show_float(fval);
//	show_pointer(pval);
    const char *s1 = "123456";
	show_bytes((byte_pointer) s1, strlen(s1));
	const char *s = "abcdef";
	show_bytes((byte_pointer) s, strlen(s));
	return 0;
}