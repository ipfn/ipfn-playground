/**
 * @license Apache-2.0
 */
import 'jest';
import { encode, decode } from '../src';


describe('base32', () => {
  const decoded = 'test';
  const encoded = 'w3jhxa0';

  it('should encode', async () => {
    expect(encode(Buffer.from(decoded))).toEqual(encoded);
  });

  it('should decode', async () => {
    expect(decode(encoded)).toEqual(Buffer.from(decoded));
  });
});
