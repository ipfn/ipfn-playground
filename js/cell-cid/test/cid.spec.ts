/**
 * @license Apache-2.0
 */
import 'jest';
import { cellCID } from '../src';


describe('CID', () => {
  it('should create valid CID', async () => {
    const cid = await cellCID({
      soul: '/synaptic/string',
      value: 'test'
    });
    expect(cid.toBaseEncodedString()).toEqual('zFunckyav7J6aWMDYJMMTcQXh1hJowx3GD8RwvZjmGBoXg5HsCgM');
  });
});
