/**
 * @license Apache-2.0
 */
import * as path from 'path';
import Ast, { SourceFile } from 'ts-simple-ast';
import { Cell } from '@ipfn/cells';
import { ModuleParser } from './modules';

// TODO: async read package.json
// export function parsePackage()

/**
 * Package parser.
 */
export class PackageParser {
  private dir: string;
  private ast: Ast;

  /**
   * Constructs parser.
   */
  constructor(dir: string) {
    this.dir = dir;
    this.ast = new Ast({
      compilerOptions: { rootDir: dir },
      tsConfigFilePath: `${dir}/tsconfig.json`,
    });
    this.ast.addSourceFiles(`${dir}/**/*.ts`);
  }

  /**
   * Gets source file by file name.
   */
  getSourceFile(name: string): SourceFile | undefined {
    return this.ast.getSourceFile(path.resolve(this.dir, `${name}.ts`));
  }

  /**
   * Gets module parser by file name.
   */
  getModule(name: string): ModuleParser {
    return new ModuleParser(this, this.getSourceFile(name)!);
  }

  /**
   * Returns entrypoint of a package.
   * TODO: proper index finding
   */
  parse(): Cell | undefined {
    const body = this.getMain().parse();
    if (body.length === 1) {
      return body[0];
    }
    return { body };
  }

  /**
   * Returns entrypoint of a package.
   * TODO: proper index finding
   */
  private getMain(): ModuleParser {
    return this.getModule('index');
  }
}
