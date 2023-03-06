/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface MasterKey {
  masterCertificate: string;
  masterAccount: string;
}

function createBaseMasterKey(): MasterKey {
  return { masterCertificate: "", masterAccount: "" };
}

export const MasterKey = {
  encode(message: MasterKey, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.masterCertificate !== "") {
      writer.uint32(10).string(message.masterCertificate);
    }
    if (message.masterAccount !== "") {
      writer.uint32(18).string(message.masterAccount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MasterKey {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMasterKey();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.masterCertificate = reader.string();
          break;
        case 2:
          message.masterAccount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MasterKey {
    return {
      masterCertificate: isSet(object.masterCertificate) ? String(object.masterCertificate) : "",
      masterAccount: isSet(object.masterAccount) ? String(object.masterAccount) : "",
    };
  },

  toJSON(message: MasterKey): unknown {
    const obj: any = {};
    message.masterCertificate !== undefined && (obj.masterCertificate = message.masterCertificate);
    message.masterAccount !== undefined && (obj.masterAccount = message.masterAccount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MasterKey>, I>>(object: I): MasterKey {
    const message = createBaseMasterKey();
    message.masterCertificate = object.masterCertificate ?? "";
    message.masterAccount = object.masterAccount ?? "";
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
