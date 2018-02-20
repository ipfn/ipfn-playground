/**
 * @license Apache-2.0
 */
import 'jest';
import * as multicodec from 'multicodec';
import { encodeCell, decodeCell } from '../src/codecs';


describe('codecs', () => {
  it('should encode and decode cell', async () => {
    const cell = { name: 'test' };
    const test = (enc: any) => expect(decodeCell(enc(cell))).toEqual(cell);
    test(encodeCell);
    test(JSON.stringify);
  });

  it('should throw on unknown prefix', async () => {
    const cell = { name: 'test' };
    const buf = Buffer.from(JSON.stringify(cell));
    const enc = multicodec.addPrefix('cbor', buf);
    expect(() => decodeCell(enc))
      .toThrow('cell codec not recognized: "51"');
  });
});
