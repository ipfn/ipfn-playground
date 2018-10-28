/**
 * @license Apache-2.0
 */
import 'jest';
import { readJSON } from '@ipfn/util';
import { PackageParser } from '../src/package';

describe('Fixtures', () => {
  const testFixture = async (name) => {
    const dir = `tests/fixtures/${name}`;
    const parser = new PackageParser(dir);
    const cell = parser.parse();
    expect(cell).toEqual(await readJSON(`${dir}/cell.json`));
  }

  it('should convert `default-class` typedocs to neurons', async () => await testFixture('default-class'));
  it('should convert `default-func` typedocs to neurons', async () => await testFixture('default-func'));
  it('should convert `export-glob` typedocs to neurons', async () => await testFixture('export-glob'));
  it('should convert `http-request` typedocs to neurons', async () => await testFixture('http-request'));
  it('should convert `simple-funcs` typedocs to neurons', async () => await testFixture('simple-funcs'));
  it('should convert `simple-namespace` typedocs to neurons', async () => await testFixture('simple-namespace'));
});
