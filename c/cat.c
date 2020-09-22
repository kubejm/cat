#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void cat(FILE *fp) {
  char buffer[8192];
  size_t n;
  ssize_t nw;

  while ((n = fread(buffer, 1, sizeof(buffer), fp)) > 0) {
    if ((nw = write(STDOUT_FILENO, buffer, n)) < 0) {
      printf("Error writing to stdout");
      exit(EXIT_FAILURE);
    }
  }
}

int main(int argc, char *argv[]) {
  for (int i = 1; i < argc; i++) {
    FILE *fp = fopen(argv[i], "r");

    if (fp == NULL) {
      printf("Error opening: %s\n", strerror(errno));
      exit(EXIT_FAILURE);
    }

    cat(fp);
    fclose(fp);
  }

  return 0;
};
