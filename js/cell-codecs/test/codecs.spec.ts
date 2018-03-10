/**
 * @license Apache-2.0
 */
import 'jest';
import * as multicodec from 'multicodec';
import { encode, decode } from '../src';
import { Cell } from '@ipfn/cell';


describe('codecs', () => {
  it('should encode and decode cell', async () => {
    const cell = { name: 'test' };
    const test = (enc: any) => expect(decode(enc(cell))).toEqual(cell);
    test((body: Cell) => encode(body, 'cell-json-v1'));
    test(JSON.stringify);
  });

  it('should throw on unknown prefix', async () => {
    const cell = { name: 'test' };
    const buf = Buffer.from(JSON.stringify(cell));
    const enc = multicodec.addPrefix('cbor', buf);
    expect(() => decode(enc))
      .toThrow('cell codec not found: "51"');
  });
});
