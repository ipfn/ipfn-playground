/**
 * @license Apache-2.0
 */
import 'jest';

import { Cell } from '@ipfn/cell';
import { cell, synaptic, method, input, output } from '../src';

describe('builder', () => {
  it('should build cell with name and soul', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
    };
    expect(cell(json.name).soul(json.soul).build()).toEqual(json);
  })

  it('should set cell bonds', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
      "bonds": [
        { "kind": "input", "from": "test" }
      ]
    };

    const builder = cell(json.name)
      .soul(json.soul)
      .bonds(json.bonds);

    expect(builder.build()).toEqual(json);
  })

  it('should set cell value', async () => {
    const json = {
      "name": "Counter",
      "soul": "/ipfn/Qm...xyz",
      "value": "test"
    };

    const builder = cell(json.name)
      .soul(json.soul)
      .value(json.value);

    expect(builder.build()).toEqual(json);
  })

  it('should build counter', async () => {
    const builder = cell('Counter')
      .soul('/ipfn/Qm...xyz')
      .body([
        synaptic.number('count'),
        method('increment')
          .body([
            input([synaptic.number('count')]),
            output([synaptic.number('count')])
          ])
      ]);

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

    expect(builder.build()).toEqual(json);
  });
});
