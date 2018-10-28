/**
 * @license Apache-2.0
 */

export namespace http {

  /**
   * HTTP header
   */
  export interface Header {
    [key: string]: string[];
  }

  /**
   * URL
   */
  export interface URL {
    scheme: string;
    host: string;
    path?: string;
    query?: string;
  }

  /**
   * HTTP Request
   */
  export interface Request {
    url: URL;
    method?: string;
    header?: Header;
    body?: string;
  }

  /**
   * HTTP Response
   */
  export interface Response {
    code: number;
    header: Header;
    body: string;
  }

  /**
   * Requests HTTP entity
   */
  export function request(request: Request): Promise<Response> {
    // **test only generator**
    return null;
  }
}
