/**
 * @license Apache-2.0
 */
import 'jest';

import { synaptic } from '../src';

describe('synaptic', () => {
  it('should build string cell', async () => {
    const test = synaptic.string('test').value('test').build();
    expect(test.name).toEqual('test');
    expect(test.soul).toEqual('/synaptic/string');
    expect(test.value).toEqual('test');
  });
});
