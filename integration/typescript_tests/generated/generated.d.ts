// Code generated by scale-signature 0.4.5, DO NOT EDIT.
// output: generated

import { Encoder, Decoder, Kind } from "@loopholelabs/polyglot"

export declare class Context {
  a: number;
  b: number;
  c: number;

  /**
  * @throws {Error}
  */
  constructor (decoder?: Decoder);

  /**
  * @throws {Error}
  */
  encode (encoder: Encoder);

  /**
  * @throws {Error}
  */
  static decode (decoder: Decoder): Context | undefined;

  /**
  * @throws {Error}
  */
  static encode_undefined (encoder: Encoder);
}

