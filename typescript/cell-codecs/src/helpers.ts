/**
 * This file contains multicodec helpers.
 * 
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cell';
import { Bytes } from './codec';
import { splitPrefix } from './prefix';
import { codecByName } from './codecs';

/**
 * Encodes cell using codec found by name.
 * Throws if multicodec is not found in local table.
 */
export function encode(cell: Cell, name: string): Bytes {
  const codec = codecByName(name);
  return codec.encode(cell);
}

/**
 * Splits body multicodec prefix and decodes cell.
 * Throws if multicodec is not found in local table.
 */
export function decode(body: Bytes): Cell {
  const [codec, content] = splitPrefix(body);
  return codec.decode(content);
}
