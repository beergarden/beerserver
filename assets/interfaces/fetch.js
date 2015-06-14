declare class Response {
  ok: boolean;
  json(): Promise<Object>;
  text(): Promise<string>;
}

declare function fetch(input: string, init?: Object): Promise<Response>;
