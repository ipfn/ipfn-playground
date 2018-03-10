/**
 * @license Apache-2.0
 */
import 'jest';
import * as multicodec from 'multicodec';
import { Cell } from '@ipfn/cell';
import { codec as protobuf } from '@ipfn/cell-pb';
import { encode, codecByName, register } from '../src';

register(protobuf);

describe('protobuf', () => {
  it('should find codec by name', async () => {
    const codec = codecByName(protobuf.name);
    expect(codec).toBe(protobuf);
  });
  // it('should encode and decode cell', async () => {
  //   const cell = { name: 'test' };
  //   const body = encode(cell, protobuf.name);
  //   const res = decode(cell, protobuf.name);
  //   const test = (enc: any) => expect(decode(enc(cell))).toEqual(cell);
  //   test((body: Cell) => encode(body, 'cell-json-v1'));
  //   test(JSON.stringify);
  // });
});
