/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface EpochData {
  totalEpochs: number;
  epochV2VRX: string;
  epochV2VBX: string;
  epochV2NBX: string;
  epochRunner: string;
}

function createBaseEpochData(): EpochData {
  return { totalEpochs: 0, epochV2VRX: "", epochV2VBX: "", epochV2NBX: "", epochRunner: "" };
}

export const EpochData = {
  encode(message: EpochData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.totalEpochs !== 0) {
      writer.uint32(8).uint64(message.totalEpochs);
    }
    if (message.epochV2VRX !== "") {
      writer.uint32(18).string(message.epochV2VRX);
    }
    if (message.epochV2VBX !== "") {
      writer.uint32(26).string(message.epochV2VBX);
    }
    if (message.epochV2NBX !== "") {
      writer.uint32(34).string(message.epochV2NBX);
    }
    if (message.epochRunner !== "") {
      writer.uint32(42).string(message.epochRunner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EpochData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEpochData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalEpochs = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.epochV2VRX = reader.string();
          break;
        case 3:
          message.epochV2VBX = reader.string();
          break;
        case 4:
          message.epochV2NBX = reader.string();
          break;
        case 5:
          message.epochRunner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EpochData {
    return {
      totalEpochs: isSet(object.totalEpochs) ? Number(object.totalEpochs) : 0,
      epochV2VRX: isSet(object.epochV2VRX) ? String(object.epochV2VRX) : "",
      epochV2VBX: isSet(object.epochV2VBX) ? String(object.epochV2VBX) : "",
      epochV2NBX: isSet(object.epochV2NBX) ? String(object.epochV2NBX) : "",
      epochRunner: isSet(object.epochRunner) ? String(object.epochRunner) : "",
    };
  },

  toJSON(message: EpochData): unknown {
    const obj: any = {};
    message.totalEpochs !== undefined && (obj.totalEpochs = Math.round(message.totalEpochs));
    message.epochV2VRX !== undefined && (obj.epochV2VRX = message.epochV2VRX);
    message.epochV2VBX !== undefined && (obj.epochV2VBX = message.epochV2VBX);
    message.epochV2NBX !== undefined && (obj.epochV2NBX = message.epochV2NBX);
    message.epochRunner !== undefined && (obj.epochRunner = message.epochRunner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EpochData>, I>>(object: I): EpochData {
    const message = createBaseEpochData();
    message.totalEpochs = object.totalEpochs ?? 0;
    message.epochV2VRX = object.epochV2VRX ?? "";
    message.epochV2VBX = object.epochV2VBX ?? "";
    message.epochV2NBX = object.epochV2NBX ?? "";
    message.epochRunner = object.epochRunner ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
