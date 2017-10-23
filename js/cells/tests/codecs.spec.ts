/**
 * @license Apache-2.0
 */
import 'jest';
import {encodeCell, decodeCell} from '../src/codecs';


describe('codecs', () => {
  it('should decode json cell', async () => {
    const cell = {name: 'test'};
    const buf = JSON.stringify(cell)
    const res = decodeCell(buf);
    expect(res).toEqual({name: 'test'});
  });
  it('should encode and decode protobuf cell', async () => {
    const cell = {name: 'test'};
    const buf = encodeCell(cell)
    const res = decodeCell(buf);
    expect(res).toEqual({name: 'test'});
  });
});
