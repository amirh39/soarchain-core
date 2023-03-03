/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Challenger {
  index: string;
  address: string;
  score: string;
  stakedAmount: string;
  netEarnings: string;
  type: string;
  ipAddr: string;
}

function createBaseChallenger(): Challenger {
  return { index: "", address: "", score: "", stakedAmount: "", netEarnings: "", type: "", ipAddr: "" };
}

export const Challenger = {
  encode(message: Challenger, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.score !== "") {
      writer.uint32(26).string(message.score);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(34).string(message.stakedAmount);
    }
    if (message.netEarnings !== "") {
      writer.uint32(42).string(message.netEarnings);
    }
    if (message.type !== "") {
      writer.uint32(50).string(message.type);
    }
    if (message.ipAddr !== "") {
      writer.uint32(58).string(message.ipAddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Challenger {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChallenger();
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
          message.score = reader.string();
          break;
        case 4:
          message.stakedAmount = reader.string();
          break;
        case 5:
          message.netEarnings = reader.string();
          break;
        case 6:
          message.type = reader.string();
          break;
        case 7:
          message.ipAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Challenger {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? String(object.score) : "",
      stakedAmount: isSet(object.stakedAmount) ? String(object.stakedAmount) : "",
      netEarnings: isSet(object.netEarnings) ? String(object.netEarnings) : "",
      type: isSet(object.type) ? String(object.type) : "",
      ipAddr: isSet(object.ipAddr) ? String(object.ipAddr) : "",
    };
  },

  toJSON(message: Challenger): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    message.stakedAmount !== undefined && (obj.stakedAmount = message.stakedAmount);
    message.netEarnings !== undefined && (obj.netEarnings = message.netEarnings);
    message.type !== undefined && (obj.type = message.type);
    message.ipAddr !== undefined && (obj.ipAddr = message.ipAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Challenger>, I>>(object: I): Challenger {
    const message = createBaseChallenger();
    message.index = object.index ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? "";
    message.stakedAmount = object.stakedAmount ?? "";
    message.netEarnings = object.netEarnings ?? "";
    message.type = object.type ?? "";
    message.ipAddr = object.ipAddr ?? "";
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
