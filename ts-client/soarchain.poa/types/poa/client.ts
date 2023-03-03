/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Client {
  index: string;
  address: string;
  score: string;
  rewardMultiplier: string;
  netEarnings: string;
  lastTimeChallenged: string;
  coolDownTolerance: string;
}

function createBaseClient(): Client {
  return {
    index: "",
    address: "",
    score: "",
    rewardMultiplier: "",
    netEarnings: "",
    lastTimeChallenged: "",
    coolDownTolerance: "",
  };
}

export const Client = {
  encode(message: Client, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.score !== "") {
      writer.uint32(34).string(message.score);
    }
    if (message.rewardMultiplier !== "") {
      writer.uint32(42).string(message.rewardMultiplier);
    }
    if (message.netEarnings !== "") {
      writer.uint32(50).string(message.netEarnings);
    }
    if (message.lastTimeChallenged !== "") {
      writer.uint32(58).string(message.lastTimeChallenged);
    }
    if (message.coolDownTolerance !== "") {
      writer.uint32(66).string(message.coolDownTolerance);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Client {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClient();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 4:
          message.score = reader.string();
          break;
        case 5:
          message.rewardMultiplier = reader.string();
          break;
        case 6:
          message.netEarnings = reader.string();
          break;
        case 7:
          message.lastTimeChallenged = reader.string();
          break;
        case 8:
          message.coolDownTolerance = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Client {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? String(object.score) : "",
      rewardMultiplier: isSet(object.rewardMultiplier) ? String(object.rewardMultiplier) : "",
      netEarnings: isSet(object.netEarnings) ? String(object.netEarnings) : "",
      lastTimeChallenged: isSet(object.lastTimeChallenged) ? String(object.lastTimeChallenged) : "",
      coolDownTolerance: isSet(object.coolDownTolerance) ? String(object.coolDownTolerance) : "",
    };
  },

  toJSON(message: Client): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    message.rewardMultiplier !== undefined && (obj.rewardMultiplier = message.rewardMultiplier);
    message.netEarnings !== undefined && (obj.netEarnings = message.netEarnings);
    message.lastTimeChallenged !== undefined && (obj.lastTimeChallenged = message.lastTimeChallenged);
    message.coolDownTolerance !== undefined && (obj.coolDownTolerance = message.coolDownTolerance);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Client>, I>>(object: I): Client {
    const message = createBaseClient();
    message.index = object.index ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? "";
    message.rewardMultiplier = object.rewardMultiplier ?? "";
    message.netEarnings = object.netEarnings ?? "";
    message.lastTimeChallenged = object.lastTimeChallenged ?? "";
    message.coolDownTolerance = object.coolDownTolerance ?? "";
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
