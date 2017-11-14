/**
 * @license Apache-2.0
 */
import * as multicodec from 'multicodec';

import { Cell } from './cell';
import { protobuf } from './protobuf';

/**
 * Encodes cell to multicodec prefixed protobuf message.
 */
export function encodeCell(cell: Cell): Buffer {
  const buf = protobuf.Cell.encode(cell);
  return multicodec.addPrefix('protobuf', buf);
}

/**
 * Decodes cell from a multicodec message.
 * Cell can be a serialized JSON or protobuf.
 */
export function decodeCell(body: string | Buffer): Cell {
  if (!(body instanceof Buffer) && body.length !== 0 && body[0] === '{') {
    return JSON.parse(body);
  }
  const buff = body instanceof Buffer ? body : Buffer.from(body);
  const prefix = multicodec.getCodec(buff);
  switch (prefix) {
    case 'protobuf':
      // TODO(crackcomm): memory
      return protobuf.Cell.decode(multicodec.rmPrefix(buff));
    default:
      throw new Error(`multicodec not recognized: "${prefix}"`);
  }
}
