/**
 * IPFN Protocol Buffer Cell Version 1 encoder and decoder functions.
 *
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cell';
import { protobuf } from './protobuf';

/**
 * Multicodec information.
 */
export namespace multicodec {
  export const name: string = 'cell-pb-v1';
  export const hex: string = '70bc'; // (28860)
}

/**
 * Encodes cell to multicodec prefixed protobuf message.
 */
export function encode(cell: Cell): Buffer {
  return protobuf.Cell.encode(cell);
}


/**
 * Decodes cell from a string or a buffer.
 * Cell can be a either plain JSON or with multicodec prefix.
 */
export function decode(body: Buffer): Cell {
  return protobuf.Cell.decode(body);
}
