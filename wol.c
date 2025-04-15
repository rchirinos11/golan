#include "wol.h"
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

#define PACKET_SIZE 103

char *makeMagic(char *mac) {
  puts("Creating magic...");
  int *macBytes = parseMacAddr(mac);

  char *bytes = malloc(PACKET_SIZE);
  memset(bytes, -1, 6);

  for (int i = 6; i < PACKET_SIZE; i++) {
    bytes[i] = macBytes[i % 6];
  }

  return bytes;
}

void sendMagic(char *magic) {
  puts("Sending magic...");
  int socketFd;

  struct addrinfo hints;
  struct addrinfo *res;

  memset(&hints, 0, sizeof hints);
  hints.ai_family = AF_INET;
  hints.ai_socktype = SOCK_DGRAM;

  int status;
  if ((status = getaddrinfo("255.255.255.255", "9", &hints, &res)) != 0) {
    fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(status));
    exit(2);
  }

  socketFd = socket(res->ai_family, res->ai_socktype, res->ai_protocol);
  if (socketFd == -1) {
    perror("Error creating socket\n");
    exit(2);
  }

  int broadcast = 1;
  if (setsockopt(socketFd, SOL_SOCKET, SO_BROADCAST, &broadcast,
                 sizeof broadcast) == -1) {
    perror("setsockopt (SO_BROADCAST)");
    exit(1);
  }

  if (sendto(socketFd, magic, PACKET_SIZE, 0, res[0].ai_addr,
             res[0].ai_addrlen) == -1) {
    perror("Error sending udp");
    exit(2);
  }

  printf("Magic packet sent.\n");
  freeaddrinfo(res);
  close(socketFd);
}

void printMacAddr(char *mac) {
  while (*mac) {
    printf("%02hhX:", *mac);
    mac++;
  }
  puts("");
}

int *parseMacAddr(char *mac) {
  int *values = malloc(6);
  if (6 != sscanf(mac, "%x:%x:%x:%x:%x:%x", &values[0], &values[1], &values[2],
                  &values[3], &values[4], &values[5])) {
    puts("Error parsing mac");
    exit(1);
  }
  return values;
}
