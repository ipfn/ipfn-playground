/**
 * @license Apache-2.0
 */
import * as multicodec from 'multicodec';

import {Cell} from './cell';
import {protobuf} from './protobuf';


/**
 * Encodes cell to multicodec prefixed protobuf message.
 */
export function encodeCell(cell: Cell) {
  const buf = protobuf.Cell.encode(cell);
  return multicodec.addPrefix('protobuf', buf);
}

/**
 * Decodes cell from a multicodec message.
 * Cell can be a serialized JSON or protobuf.
 */
export function decodeCell(body: any): Cell {
  if (body.length != 0 && body[0] === '{') {
    return JSON.parse(body);
  }
  const prefix = multicodec.getCodec(body);
  switch (prefix) {
    case 'protobuf':
    // TODO(crackcomm): Protobuf `Any` and `Value` parser.
    return protobuf.Cell.decode(multicodec.rmPrefix(body));
  default:
    throw `unknown cell codec ${prefix}`;  
  }
}
