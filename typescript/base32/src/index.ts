/**
 * @license Apache-2.0
 */

import { _encode, _decode } from 'bs32';

const CHARSET = '0pzqy9x8bf2tvrwds3jn54khce6mua7l';
const TABLE = [
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  0, -1, 10, 17, 21, 20, 26, 30, 7, 5, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
  29, 8, 24, 15, 25, 9, -1, 23, -1, 18, 22, 31, 27, 19, -1, 1, 3,
  13, 16, 11, 28, 12, 14, 6, 4, 2, -1, -1, -1, -1, -1
];

/**
 * Encodes data with IPFN base32 character table.
 */
export function encode(data: Buffer, pad: boolean = false, charset: string = CHARSET): string {
  return _encode(data, charset, pad);
}

/**
 * Decodes data with IPFN base32 character table.
 */
export function decode(data: string, unpad: boolean = false, table: number[] = TABLE): Buffer {
  return _decode(data, table, unpad);
}
