/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.soarmint";

export interface Minter {
  inflation: string;
  phase: number;
  startPhaseBlock: number;
  annualProvisions: string;
  targetSupply: string;
}

function createBaseMinter(): Minter {
  return { inflation: "", phase: 0, startPhaseBlock: 0, annualProvisions: "", targetSupply: "" };
}

export const Minter = {
  encode(message: Minter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.inflation !== "") {
      writer.uint32(10).string(message.inflation);
    }
    if (message.phase !== 0) {
      writer.uint32(16).uint64(message.phase);
    }
    if (message.startPhaseBlock !== 0) {
      writer.uint32(24).uint64(message.startPhaseBlock);
    }
    if (message.annualProvisions !== "") {
      writer.uint32(34).string(message.annualProvisions);
    }
    if (message.targetSupply !== "") {
      writer.uint32(42).string(message.targetSupply);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inflation = reader.string();
          break;
        case 2:
          message.phase = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.startPhaseBlock = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.annualProvisions = reader.string();
          break;
        case 5:
          message.targetSupply = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Minter {
    return {
      inflation: isSet(object.inflation) ? String(object.inflation) : "",
      phase: isSet(object.phase) ? Number(object.phase) : 0,
      startPhaseBlock: isSet(object.startPhaseBlock) ? Number(object.startPhaseBlock) : 0,
      annualProvisions: isSet(object.annualProvisions) ? String(object.annualProvisions) : "",
      targetSupply: isSet(object.targetSupply) ? String(object.targetSupply) : "",
    };
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.inflation !== undefined && (obj.inflation = message.inflation);
    message.phase !== undefined && (obj.phase = Math.round(message.phase));
    message.startPhaseBlock !== undefined && (obj.startPhaseBlock = Math.round(message.startPhaseBlock));
    message.annualProvisions !== undefined && (obj.annualProvisions = message.annualProvisions);
    message.targetSupply !== undefined && (obj.targetSupply = message.targetSupply);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Minter>, I>>(object: I): Minter {
    const message = createBaseMinter();
    message.inflation = object.inflation ?? "";
    message.phase = object.phase ?? 0;
    message.startPhaseBlock = object.startPhaseBlock ?? 0;
    message.annualProvisions = object.annualProvisions ?? "";
    message.targetSupply = object.targetSupply ?? "";
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
