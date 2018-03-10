/**
 * @license Apache-2.0
 */
import protons from '@ipfn/protons';
import { Cell, Bond } from '@ipfn/cell';

/**
 * Protobuf coder interface.
 */
export interface ProtobufCodec<T> {
  /**
   * Encodes a cell to protocol buffers.
   */
  encode(cell: T): Buffer;

  /**
   * Decodes protocol buffers into a cell.
   */
  decode(cell: Buffer): T;
}


/**
 * Protobuf typings.
 */
export interface Protobuf {
  Cell: ProtobufCodec<Cell>;
  Bond: ProtobufCodec<Bond>;
}

/**
 * Protocol buffers.
 * Also: /ipfs/QmeX5H9x2qNdGC1R5uhyX2HuG5izxR2SGi71jSWyEQjV6Q
 */
export const protobuf: Protobuf = protons(`
  syntax = "proto3";

  package ipfn;

  message Cell {
    string name         = 1;
    string soul         = 2;
    bytes  value        = 3;
    repeated Bond bonds = 4;
    repeated Cell body  = 5;
  }

  message Bond {
    string name  = 1;
    string kind  = 2;
    string from  = 3;
    string to    = 4;
  }
`);

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

/**
 * Codec information.
 * Usage: https://github.com/ipfn/js-ipfn-cell-codecs
 */
export const codec = {
  name: 'cell-pb-v1',
  hex: '70bc',
  encode,
  decode,
};
