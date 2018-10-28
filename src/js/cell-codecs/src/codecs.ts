/**
 * This file contains simple codecs registry.
 * 
 * @license Apache-2.0
 */
import * as varintTable from 'multicodec/src/varint-table';
import { varintBufferEncode } from 'multicodec/src/util';
import { Cell } from '@ipfn/cell';
import { Codec } from './codec';

export interface CodecMap {
  [key: string]: Codec;
}

export const CODEC_BY_HEX: CodecMap = {};
export const CODEC_BY_NAME: CodecMap = {};

/**
 * Registers codec in multicodec table.
 */
export function register(codec: Codec) {
  // Register in local codecs indexes
  // and in multicodec varint table
  // Hex can be empty in case of codecs
  // like JSON or non-multicodec codecs
  if (codec.hex) {
    const base = Buffer.from(codec.hex, 'hex');
    const varint = varintBufferEncode(base);
    varintTable[codec.name] = varint;
    CODEC_BY_HEX[codec.hex] = codec;
  }
  CODEC_BY_NAME[codec.name] = codec;
}

/**
 * Retrieves a codec from codec table by name.
 * Throws when codec does not exist in local table.
 */
export function codecByName(name: string): Codec {
  const codec = CODEC_BY_NAME[name];
  if (!codec) {
    throw new Error(`cell codec not found: "${name}"`);
  }
  return codec;
}

/**
 * Retrieves a codec from codec table by hex prefix.
 * Throws when codec does not exist in local table.
 */
export function codecByHex(hex: string): Codec {
  const codec = CODEC_BY_HEX[hex];
  if (!codec) {
    throw new Error(`cell codec not found: "${hex}"`);
  }
  return codec;
}
