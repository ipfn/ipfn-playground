/**
 * @license Apache-2.0
 */
import * as child_process from 'child_process';
import * as fs from 'fs';
import * as fsextra from 'fs-extra';
import * as globBase from 'glob';
import * as os from 'os';
import * as path from 'path';

/**
 * Promisifies a function
 */
export function promisify(fn: Function) {
  return function (...args: any[]): Promise<any> {
    return new Promise((resolve, reject) => fn(...args, (err: any, res: any) => {
      if (err) {
        reject(err);
      } else {
        resolve(res);
      }
    }))
  };
}

/**
 * Creates a directory with random name in `os.tmpdir()`.
 */
export async function tmpdir(): Promise<string> {
  const dirname = Math.random().toString(36).substring(7);
  const dir = path.join(os.tmpdir(), dirname);
  await ensureDir(dir);
  return dir;
};

export const copyDir = promisify(fsextra.copy);
export const exec = promisify(child_process.exec);
export const glob = promisify(globBase);
export const ensureDir = promisify(fsextra.ensureDir);
export const readFile = promisify(fs.readFile);
export const readJSON = promisify(fsextra.readJson);
export const writeFile = promisify(fs.writeFile);
