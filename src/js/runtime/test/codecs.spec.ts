/**
 * @license Apache-2.0
 */
import 'jest';
import { encodeCell, decodeCell } from '../src/codecs';


describe('codecs', () => {
  it('should encode and decode cell', async () => {
    const cell = { name: 'test' };
    const test = (enc) => expect(decodeCell(enc(cell))).toEqual(cell);
    test(encodeCell);
    test(JSON.stringify);
  });
});
