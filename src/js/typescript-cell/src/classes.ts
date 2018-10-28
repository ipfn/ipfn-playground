/**
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cells';
import {
  ClassDeclaration,
  PropertyDeclaration,
  GetAccessorDeclaration,
  SetAccessorDeclaration
} from 'ts-simple-ast';

import { typeToCells } from './types';
import { functionToCell } from './functions';
import { reduceCells } from './util';

/**
 * Creates a cell from class declaration.
 */
// TODO: class constructors
// TODO: static class functions
export function classToCell(decl: ClassDeclaration): Cell {
  const functions: Cell[] = decl.getInstanceMethods()
    .filter(method => method.isImplementation())
    .map(method => functionToCell(method, 'method'));
  const properties: Cell[] = reduceCells(decl.getInstanceProperties()
    .map(prop => {
      if (prop instanceof PropertyDeclaration) {
        return typeToCells(prop.getType(), prop.getName());
      }
      if (prop instanceof GetAccessorDeclaration) {
        return typeToCells(prop.getReturnType(), prop.getName());
      }
      if (prop instanceof SetAccessorDeclaration) {
        return [functionToCell(prop, 'method')];
      }
      // TODO: ParameterDeclaration
      return [{ name: prop.getName() }];
    }));
  // class cell name
  const name = decl.getName();
  // cell body
  const body: Cell[] = functions.concat(properties);
  return { name, soul: '/core/class', body };
}
