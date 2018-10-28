/**
 * This file defines codec interface.
 *
 * @license Apache-2.0
 */
import { varintBufferDecode } from 'multicodec/src/util';
import { Cell } from '@ipfn/cell';
import { Codec } from './codec';

/**
 * Codec bytes type.
 */
export type Bytes = Buffer | string;

/**
 * Codec interface.
 */
export interface Codec {
  // Name of codec
  name: string;
  // Multicodec hex
  hex?: string;
  // Encodes cell
  encode(cell: Cell): Bytes;
  // Decodes cell
  decode(body: Bytes): Cell;
}
