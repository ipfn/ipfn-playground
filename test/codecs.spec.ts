/**
 * @license Apache-2.0
 */
import 'jest';
import { encode, decode } from '../src';


describe('codecs', () => {
  it('should encode and decode cell', async () => {
    const cell = { name: 'test' };
    const enc = encode(cell);
    const dec = decode(enc);
    expect(dec).toEqual(cell);
  });
});
