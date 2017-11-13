/**
 * @license Apache-2.0
 */
import * as path from 'path';
import { Cell } from '@ipfn/cells';
import Ast, {
  SourceFile,
  ClassDeclaration,
  FunctionDeclaration,
  NamespaceDeclaration
} from 'ts-simple-ast';

import { classToCell } from './classes';
import { functionToCell, paramsToCells } from './functions';
import { PackageParser } from './package';

import * as compactJSON from 'json-stringify-pretty-compact';
import { reduceCells } from './util';

/**
 * Module parser.
 */
export class ModuleParser {
  private pkg: PackageParser;
  private source: SourceFile;

  constructor(pkg: PackageParser, source: SourceFile) {
    this.pkg = pkg;
    this.source = source;
  }

  /**
   * Parses source file to cell.
   */
  parse(): Cell[] {
    let results: Cell[][] = [];
    const defaults = this.parseDefault();
    if (defaults) {
      results.push([defaults]);
    }
    results.push(this.parseFunctions());
    results.push(this.parseNamespaces());
    results.push(this.parseExported());
    return reduceCells(results);
  }

  /**
   * Parses exported in source.
   * Example: `export * from '...';`.
   */
  private parseExported(): Cell[] {
    const cells: Cell[][] = this.source.getExports()
      .filter(exp => exp.isNamespaceExport())
      .map(exp => exp.getModuleSpecifier()!)
      .map(mod => this.pkg.getModule(mod)!.parse()!);
    return reduceCells(cells);
  }

  /**
   * Parse functions in source.
   * Example: `export function ...`.
   */
  private parseFunctions(): Cell[] {
    return this.source.getFunctions()
      .filter(fn => fn.isNamedExport())
      .map(fn => functionToCell(fn));
  }

  /**
   * Parses namespaces in source.
   * Example: `export namespace { ... }`.
   */
  private parseNamespaces(): Cell[] {
    return this.source.getNamespaces()
      .filter(ns => ns.isNamedExport())
      .map(parseNamespace);
  }

  /**
   * Parses default exports in source.
   * Example: `export default ...`.
   */
  private parseDefault(): Cell | undefined {
    // Get module default export
    const symbolDefault = this.source.getDefaultExportSymbol();
    // If module has a default export
    if (symbolDefault) {
      // Filter function declarations
      const decls = symbolDefault.getDeclarations()
        .filter(decl => !(decl instanceof FunctionDeclaration) || decl.isImplementation());
      // Get first declaration
      if (decls.length >= 1) {
        const decl = decls[0];
        // Return function cell
        if (decl instanceof FunctionDeclaration) {
          return functionToCell(decl);
        }
        // Return class cell
        if (decl instanceof ClassDeclaration) {
          return classToCell(decl);
        }
      }
    }
  }
}

function parseNamespace(ns: NamespaceDeclaration): Cell {
  const funcs: Cell[] = ns.getFunctions()
    .filter(fn => fn.hasExportKeyword())
    .map(fn => functionToCell(fn));
  const classes: Cell[] = ns.getClasses()
    .filter(fn => fn.hasExportKeyword())
    .map(fn => classToCell(fn));
  const namespaces: Cell[] = ns.getNamespaces()
    .filter(ns => ns.hasExportKeyword())
    .map(ns => parseNamespace(ns));
  const body = reduceCells([funcs, classes, namespaces]);
  return {
    name: ns.getName(),
    soul: '/core/namespace',
    body
  };
}
