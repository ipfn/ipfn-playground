/**
 * @license Apache-2.0
 */
import 'jest';

import { Cell } from '@ipfn/cell';
import { cell, synaptic, fromValue, fromObject, fromCell } from '../src';

describe('convert', () => {
  it('should create named cell', async () => {
    const test = cell('test').build();
    expect(test).toEqual(fromValue(undefined, 'test'));
  });

  it('should create cell from string', async () => {
    const test = synaptic.string('test').value('test').build();
    expect(test).toEqual(fromValue('test', 'test'));
  });

  it('should create cell from bool', async () => {
    const test = synaptic.bool('test').value(true).build();
    expect(test).toEqual(fromValue(true, 'test'));
  });

  it('should build cell from object', async () => {
    const test = synaptic.string('test').value('test').build();
    const num = synaptic.number('num').value(10000).build();
    const obj = { test: 'test', num: 10000 };
    const res = { body: [test, num] };
    expect(fromObject(obj)).toEqual(res);
    expect(fromObject({ obj })).toEqual({
      body: [{
        ...res,
        name: 'obj'
      }]
    });
  });

  it('should create string from cell', async () => {
    const test = synaptic.string().value('test').build();
    expect(fromCell(test)).toEqual('test');
  });

  it('should create bool from cell', async () => {
    const test = synaptic.bool().value(true).build();
    expect(fromCell(test)).toEqual(true);
  });

  it('should create bool from string cell', async () => {
    const test = synaptic.bool().value('1').build();
    expect(fromCell(test)).toEqual(true);
  });

  it('should create number from cell', async () => {
    const test = synaptic.number().value(10000).build();
    expect(fromCell(test)).toEqual(10000);
  });

  it('should ignore empty cells', async () => {
    const test = synaptic.string('test').build();
    expect(fromCell(test)).toEqual(undefined);
  });

  it('should create object from cell in body', async () => {
    const test = synaptic.string('test').value('test').build();
    expect(fromCell({ body: [test] })).toEqual({ test: 'test' });
  });

  it('should create object from cell in named body', async () => {
    const test = synaptic.string('test').value('test').build();
    expect(fromCell({ name: 'parent', body: [test] })).toEqual({ parent: { test: 'test' } });
  });
});
