/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface VrfUser {
  index: string;
  address: string;
  count: string;
}

function createBaseVrfUser(): VrfUser {
  return { index: "", address: "", count: "" };
}

export const VrfUser = {
  encode(message: VrfUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.count !== "") {
      writer.uint32(26).string(message.count);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VrfUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVrfUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.count = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VrfUser {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      address: isSet(object.address) ? String(object.address) : "",
      count: isSet(object.count) ? String(object.count) : "",
    };
  },

  toJSON(message: VrfUser): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VrfUser>, I>>(object: I): VrfUser {
    const message = createBaseVrfUser();
    message.index = object.index ?? "";
    message.address = object.address ?? "";
    message.count = object.count ?? "";
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
