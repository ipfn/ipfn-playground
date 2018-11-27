#define _GNU_SOURCE

#include <math.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include "x25519.h"

extern "C" {
extern uint64_t
tsc_read();
}

int
main(int argc, char **argv) {
  uint8_t private_key[32], pubkey[32], peer1[32], peer2[32], output[32];
  static const uint8_t basepoint[32] = {9};
  unsigned i;
  uint64_t sum = 0, sum_squares = 0, skipped = 0, mean;
  static const unsigned count = 200000;

  memset(private_key, 42, sizeof(private_key));

  private_key[0] &= 248;
  private_key[31] &= 127;
  private_key[31] |= 64;

  x25519(pubkey, private_key, basepoint);
  memset(peer1, 0, sizeof(peer1));
  memset(peer2, 255, sizeof(peer2));

  for (i = 0; i < count; ++i) {
    const uint64_t start = tsc_read();
    x25519(output, peer1, pubkey);
    const uint64_t end = tsc_read();
    const uint64_t delta = end - start;
    if (delta > 650000) {
      // something terrible happened (task switch etc)
      skipped++;
      continue;
    }
    sum += delta;
    sum_squares += (delta * delta);
  }

  mean = sum / ((uint64_t)count);
  printf("all 0: mean:%lu sd:%f skipped:%lu\n", mean,
         sqrt((double)(sum_squares / ((uint64_t)count) - mean * mean)),
         skipped);

  sum = sum_squares = skipped = 0;

  for (i = 0; i < count; ++i) {
    const uint64_t start = tsc_read();
    x25519(output, peer2, pubkey);
    const uint64_t end = tsc_read();
    const uint64_t delta = end - start;
    if (delta > 650000) {
      // something terrible happened (task switch etc)
      skipped++;
      continue;
    }
    sum += delta;
    sum_squares += (delta * delta);
  }

  mean = sum / ((uint64_t)count);
  printf("all 1: mean:%lu sd:%f skipped:%lu\n", mean,
         sqrt((double)(sum_squares / ((uint64_t)count) - mean * mean)),
         skipped);

  return 0;
}
