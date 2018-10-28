/**
 * @license Apache-2.0
 */
import { Cell, Bond } from '@ipfn/cell';


/**
 * Cell or a builder.
 */
export type CellOrBuilder = Cell | CellBuilder;

/**
 * Converts builders to cells.
 */
export function toCells(body: CellOrBuilder[]): Cell[] {
  return body.map(c => c instanceof CellBuilder ? c.build() : c);
}

/**
 * Creates a new instance of Cell Builder.
 */
export function cell(name?: string): CellBuilder {
  return new CellBuilder(name);
}

/**
 * Builder of a cell.
 */
export class CellBuilder {
  private result: Cell = {};

  /**
   * Constructs a cell builder.
   * @param name Cell name
   */
  constructor(name?: string) {
    if (name) {
      this.result.name = name;
    }
  }

  /**
   * Sets cell soul.
   */
  soul(soul: string): CellBuilder {
    this.result.soul = soul;
    return this;
  }

  /**
   * Sets cell body.
   */
  body(body: CellOrBuilder | CellOrBuilder[]): CellBuilder {
    if (body instanceof Array) {
      this.result.body = toCells(body);
    } else if (this.result.body) {
      this.result.body.push(body instanceof CellBuilder ? body.build() : body);
    } else {
      this.result.body = toCells([body]);
    }
    return this;
  }

  /**
   * Sets cell bonds.
   */
  bonds(bonds: Bond[]): CellBuilder {
    this.result.bonds = bonds;
    return this;
  }

  /**
   * Sets cell bond.
   */
  bond(bond: Bond): CellBuilder {
    if (bond) {
      if (this.result.bonds) {
        this.result.bonds.push(bond);
      } else {
        this.result.bonds = [bond];
      }
    }
    return this;
  }

  /**
   * Sets cell value.
   */
  value(value: any): CellBuilder {
    this.result.value = value;
    return this;
  }

  /**
   * Returns underlying cell which was built.
   */
  build(): Cell {
    return this.result;
  }
}
