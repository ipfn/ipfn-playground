/**
 * @license Apache-2.0
 */
import 'jest';
import { readJSON, promisify } from './../src/index';

/**
 * Utilities test
 */
describe('Utilities', () => {
  it('readJSON', async () => {
    const file = await readJSON('test/fixtures/test.json');
    expect(file.ok).toBe('yes');
  })
  it('promisify ok', async () => {
    const func = (some, args, cb) => {
      cb(null, some + args);
    }
    const pfunc = promisify(func);
    expect(await pfunc(1, 2)).toBe(3);
  })
  it('promisify err', async () => {
    const func = (some, args, cb) => {
      cb('err');
    }
    const pfunc = promisify(func);
    expect(pfunc(1, 2)).rejects.toBe('err');
  })
})
