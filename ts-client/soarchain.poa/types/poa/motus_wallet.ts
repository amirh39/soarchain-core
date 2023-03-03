/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Client } from "./client";

export const protobufPackage = "soarchain.poa";

export interface MotusWallet {
  index: string;
  client: Client | undefined;
}

function createBaseMotusWallet(): MotusWallet {
  return { index: "", client: undefined };
}

export const MotusWallet = {
  encode(message: MotusWallet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.client !== undefined) {
      Client.encode(message.client, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MotusWallet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMotusWallet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.client = Client.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MotusWallet {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      client: isSet(object.client) ? Client.fromJSON(object.client) : undefined,
    };
  },

  toJSON(message: MotusWallet): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.client !== undefined && (obj.client = message.client ? Client.toJSON(message.client) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MotusWallet>, I>>(object: I): MotusWallet {
    const message = createBaseMotusWallet();
    message.index = object.index ?? "";
    message.client = (object.client !== undefined && object.client !== null)
      ? Client.fromPartial(object.client)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
