#include "wol.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  char *mac = getenv("WOLMAC");
  if (!mac) {
    puts("WOLMAC environment variable is not set");
    return 1;
  }

  char *magic = makeMagic(mac);
  sendMagic(magic);
  return 0;
}
