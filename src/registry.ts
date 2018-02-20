/**
 * @license Apache-2.0
 */
import varintTable from 'multicodec/src/varint-table';
import { varintBufferEncode } from 'multicodec/src/util';
import { Cell } from '@ipfn/cell';

/** 
 * Codec interface.
 */
export interface Codec {
  // Name of codec
  name: string;
  // Multicodec hex
  hex: string;
  encode(cell: Cell): Buffer;
  decode(body: Buffer): Cell;
}

interface CodecMap {
  [key: string]: Codec;
}

const CODEC_BY_HEX: CodecMap = {};
const CODEC_BY_NAME: CodecMap = {};

/**
 * Registers codec in multicodec table.
 */
export function register(codec: Codec) {
  const base = Buffer.from(codec.hex, 'hex');
  const varint = varintBufferEncode(base);
  // Register in varint table
  varintTable[codec.name] = varint;
  // Register in local codecs indexes
  CODEC_BY_HEX[codec.hex] = codec;
  CODEC_BY_NAME[codec.name] = codec;
}

/**
 * Retrieves a codec from codec table by name.
 */
export function codecByName(name: string): Codec | undefined {
  return CODEC_BY_NAME[name];
}

/**
 * Retrieves a codec from codec table by hex prefix.
 */
export function codecByHex(hex: string): Codec | undefined {
  return CODEC_BY_HEX[hex];
}
