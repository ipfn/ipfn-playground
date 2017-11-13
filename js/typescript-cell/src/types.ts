/**
 * @license Apache-2.0
 */
import { Type, PropertySignature } from 'ts-simple-ast';
import { Cell, cell } from '@ipfn/cells';
import { reduceCells } from './util';

/**
 * Converts `Type`s to descriptor `Cell`s.
 */
// TODO: Promise<>
export function typeToCells(type: Type, name?: string): Cell[] {
  if (type.isUnionType()) {
    const types = type
      .getUnionTypes()
      .map(type => typeToCells(type, name));
    return reduceCells(types);
  }
  if (type.isInterfaceType()) {
    const props = type.getProperties()
      .map(prop => {
        const propName = prop.getName();
        const cells = prop.getDeclarations()
          .map((decl: PropertySignature) => {
            // if (propName === 'header') {
            //   console.log(`${propName}`, decl.getType().getStringIndexType())
            // }

            return typeToCells(decl.getType(), propName);
          });
        return reduceCells(cells);
      })

    const indexType = type.getStringIndexType();
    if (indexType) {
      console.log(indexType);
    }

    const result = cell(name)
      .body(reduceCells(props))
      .build();
    return [result];
  }
  if (type.isObjectType()) {
    // TODO
  }
  const result = cell(name)
    .soul(`/synaptic/${type.getText()}`)
    .build();
  return [result];
}
