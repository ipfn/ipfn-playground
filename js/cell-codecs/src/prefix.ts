/**
 * @license Apache-2.0
 */
import * as multicodec from 'multicodec';
import { varintBufferDecode } from 'multicodec/src/util';
import { Cell } from '@ipfn/cell';
import { Bytes, Codec } from './codec';
import { codecByHex, codecByName } from './codecs';

/**
 * Prepends multicodec prefix to a message.
 * It works only for cell codecs not others.
 */
export function prepend(body: Buffer, name: string): Buffer {
  const codec = codecByName(name);
  if (!codec.hex) {
    return body;
  }
  return multicodec.addPrefix(name, body);
}

/**
 * Extracts multicodec prefix from message.
 * Returns codec name and message buffer.
 * Throws if codec prefix was not found.
 */
export function splitPrefix(body: Bytes): [Codec, Bytes] {
  if (isJSONObject(body)) {
    const codec = codecByName('cell-json-v1');
    return [codec, body];
  }
  const buffer = body instanceof Buffer ? body : Buffer.from(body);
  const prefix = varintBufferDecode(buffer).toString('hex');
  const codec = codecByHex(prefix);
  return [codec, multicodec.rmPrefix(buffer)];
}

/**
 * Checks if string or a buffer is a JSON object.
 */
function isJSONObject(body: Buffer | string): boolean {
  body = body.toString();
  const length = body.length;
  return length >= 2 && body[0] === '{' && body[length - 1] === '}';
}
