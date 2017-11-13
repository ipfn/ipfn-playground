/**
 * @license Apache-2.0
 */
import { Cell } from '../cell';

/**
 * Cell or a builder.
 */
export type CellOrBuilder = Cell | Builder;

/**
 * Converts builders to cells.
 */
export function toCells(body: CellOrBuilder[]): Cell[] {
  return body.map(cell => cell instanceof Builder ? cell.build() : cell);
}

/**
 * Builder of a mutable cell.
 */
export class Builder {
  private cell: Cell = {};

  /**
   * Constructs a cell builder.
   * @param name Cell name
   */
  constructor(name?: string) {
    if (name) {
      this.cell.name = name;
    }
  }

  /**
   * Returns cell.
   */
  build(): Cell {
    return this.cell;
  }

  /**
   * Sets cell soul.
   */
  soul(soul: string): Builder {
    this.cell.soul = soul;
    return this;
  }

  /**
   * Sets cell body.
   */
  body(body: (Cell | Builder)[]): Builder {
    this.cell.body = toCells(body);
    return this;
  }

  /**
   * Sets cell bonds.
   */
  bonds(bonds: string[]): Builder {
    this.cell.bonds = bonds;
    return this;
  }

  /**
   * Sets cell memory.
   */
  memory(memory: any): Builder {
    this.cell.memory = memory;
    return this;
  }
}
