/**
 * @license Apache-2.0
 */
import * as multicodec from 'multicodec';
import varintTable from 'multicodec/src/varint-table';
import { varintBufferEncode, varintBufferDecode } from 'multicodec/src/util';

import { Cell } from '@ipfn/cell';
import { protobuf } from './protobuf';
import { BasicCell } from './types';

/**
 * Encodes cell to multicodec prefixed protobuf message.
 */
export function encodeCell(cell: CellOrBuilder, withoutPrefix?: boolean): Buffer {
  // Encode to protocol buffers
  const buf = protobuf.Cell.encode(cell instanceof CellBuilder ? cell.cell : cell);
  if (!withoutPrefix) {
    // Create buffer with multicodec prefix
    return Buffer.concat([cellPbV1.varint, buf]);
  }
  return buf;
}


/**
 * Decodes cell from a string or a buffer.
 * Cell can be a either plain JSON or with multicodec prefix.
 */
export function decodeCell(body: string | Buffer, prefix?: string): BasicCell {
  // Check if message is a JSON string
  if (!(body instanceof Buffer) && body.length !== 0 && body[0] === '{') {
    return JSON.parse(body);
  }
  // Converts string into a buffer always
  const buff = body instanceof Buffer ? body : Buffer.from(body);
  // Read multicodec prefix from message body
  if (!prefix) {
    prefix = varintBufferDecode(buff).toString('hex');
  }
  // Deserialize message depending on prefix
  switch (prefix) {
    case cellPbV1.hex:
      return protobuf.Cell.decode(multicodec.rmPrefix(buff));
    default:
      throw new Error(`cell codec not recognized: "${prefix}"`);
  }
}
