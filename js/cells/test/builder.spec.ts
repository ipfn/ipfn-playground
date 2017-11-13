/**
 * @license Apache-2.0
 */
import 'jest';

import { Cell, cell, input, output, method, number } from '../src/index';

describe('builder', () => {
  it('should build cell with name and soul', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
    };
    expect(cell(json.name, json.soul).build()).toEqual(json);
  })
  it('should set cell bonds', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
      "bonds": [
        "input:test"
      ]
    };
    const res = cell(json.name, json.soul)
      .bonds(json.bonds)
      .build();
    expect(res).toEqual(json);
  })
  it('should set cell memory', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
      "memory": "test"
    };
    const res = cell(json.name, json.soul)
      .memory(json.memory)
      .build();
    expect(res).toEqual(json);
  })
  it('should build counter', async () => {
    const builder = cell('Counter')
      .soul('/ipfn/Qm...xyz')
      .body([
        number('count'),
        method('increment')
          .body([
            input([number('count')]),
            output([number('count')])
          ])
      ])
      .build();

    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
      "body": [
        {
          "name": "count",
          "soul": "/synaptic/number"
        },
        {
          "name": "increment",
          "soul": "/core/method",
          "body": [
            {
              "soul": "/core/input",
              "body": [
                {
                  "name": "count",
                  "soul": "/synaptic/number"
                }
              ]
            },
            {
              "soul": "/core/output",
              "body": [
                {
                  "name": "count",
                  "soul": "/synaptic/number"
                }
              ]
            }
          ]
        }
      ]
    };
    expect(builder).toEqual(json);
  });
});
