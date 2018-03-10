/**
 * This file implements JSON cell codec.
 * 
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cell';
import { Bytes } from './codec';

/**
 * Encodes cell to multicodec prefixed protobuf message.
 */
function encode(cell: Cell): Bytes {
  return JSON.stringify(cell);
}


/**
 * Decodes cell from a string or a buffer.
 * Cell can be a either plain JSON or with multicodec prefix.
 */
function decode(body: Bytes): Cell {
  return JSON.parse(body instanceof Buffer ? body.toString() : body);
}

/**
 * Codec information.
 * Usage: https://github.com/ipfn/ipfn/tree/master/js/cell-codecs
 */
export const codec = {
  name: 'cell-json-v1',
  encode,
  decode,
};
