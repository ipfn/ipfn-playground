/**
 * @license Apache-2.0
 */

export namespace includeDeep {

  namespace nonexport { }

  export namespace includeDeeper {
    function noexpo() { }

    export function deepSimpleFunction(firstParam?: string | number) {
      return {
        test: 'ok'
      };
    }
  }

}
