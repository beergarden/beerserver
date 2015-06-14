declare class Response {
  ok: boolean;
  json(): Promise<any>;
  text(): Promise<string>;
}

declare function fetch(input: string, init?: Object): Promise<Response>;
