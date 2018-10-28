/**
 * @license Apache-2.0
 */
import { Symbol, Type, ParameterDeclaration } from 'ts-simple-ast';
import { Cell, input, output } from '@ipfn/cells';
import { typeToCells } from './types';
import { reduceCells } from './util';

export interface IFunctionDeclaration {
  getName(): string;
  getParameters(): ParameterDeclaration[];
  getReturnType(): Type;
}

export function functionToCell(func: IFunctionDeclaration, named?: 'function' | 'method'): Cell {
  // Get function name
  const name = func.getName();
  // Create function cell body (input, output)
  const body: Cell[] = [];
  // Get function input parameters
  const params = func.getParameters();
  // Check if function has any parameters
  if (params && params.length > 0) {
    body.push(input(paramsToCells(params)).build());
  }
  // Get function return type
  const returnType = func.getReturnType();
  if (!returnType.isUndefinedType() && returnType.getText() !== 'void') {
    body.push(output(typeToCells(returnType)).build());
  }
  // Create result cell
  const result: Cell = { name, soul: `/core/${named || 'function'}` };
  if (body.length > 0) {
    result.body = body;
  }
  return result;
}

/**
 * Converts function input parameters to cells.
 */
export function paramsToCells(params: ParameterDeclaration[]): Cell[] {
  return reduceCells(params.map(paramToCells));
}

/**
 * Converts function input parameter to cell.
 * TODO: only works on union type
 */
function paramToCells(param: ParameterDeclaration): Cell[] {
  return typeToCells(param.getType(), param.getName());
}
