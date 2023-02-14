/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Runner {
  index: string;
  address: string;
  score: string;
  rewardMultiplier: string;
  stakedAmount: string;
  netEarnings: string;
  ipAddr: string;
  lastTimeChallenged: string;
  coolDownTolerance: string;
}

function createBaseRunner(): Runner {
  return {
    index: "",
    address: "",
    score: "",
    rewardMultiplier: "",
    stakedAmount: "",
    netEarnings: "",
    ipAddr: "",
    lastTimeChallenged: "",
    coolDownTolerance: "",
  };
}

export const Runner = {
  encode(message: Runner, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.score !== "") {
      writer.uint32(26).string(message.score);
    }
    if (message.rewardMultiplier !== "") {
      writer.uint32(34).string(message.rewardMultiplier);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(42).string(message.stakedAmount);
    }
    if (message.netEarnings !== "") {
      writer.uint32(50).string(message.netEarnings);
    }
    if (message.ipAddr !== "") {
      writer.uint32(58).string(message.ipAddr);
    }
    if (message.lastTimeChallenged !== "") {
      writer.uint32(66).string(message.lastTimeChallenged);
    }
    if (message.coolDownTolerance !== "") {
      writer.uint32(74).string(message.coolDownTolerance);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Runner {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRunner();
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
          message.rewardMultiplier = reader.string();
          break;
        case 5:
          message.stakedAmount = reader.string();
          break;
        case 6:
          message.netEarnings = reader.string();
          break;
        case 7:
          message.ipAddr = reader.string();
          break;
        case 8:
          message.lastTimeChallenged = reader.string();
          break;
        case 9:
          message.coolDownTolerance = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Runner {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      address: isSet(object.address) ? String(object.address) : "",
      score: isSet(object.score) ? String(object.score) : "",
      rewardMultiplier: isSet(object.rewardMultiplier) ? String(object.rewardMultiplier) : "",
      stakedAmount: isSet(object.stakedAmount) ? String(object.stakedAmount) : "",
      netEarnings: isSet(object.netEarnings) ? String(object.netEarnings) : "",
      ipAddr: isSet(object.ipAddr) ? String(object.ipAddr) : "",
      lastTimeChallenged: isSet(object.lastTimeChallenged) ? String(object.lastTimeChallenged) : "",
      coolDownTolerance: isSet(object.coolDownTolerance) ? String(object.coolDownTolerance) : "",
    };
  },

  toJSON(message: Runner): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    message.rewardMultiplier !== undefined && (obj.rewardMultiplier = message.rewardMultiplier);
    message.stakedAmount !== undefined && (obj.stakedAmount = message.stakedAmount);
    message.netEarnings !== undefined && (obj.netEarnings = message.netEarnings);
    message.ipAddr !== undefined && (obj.ipAddr = message.ipAddr);
    message.lastTimeChallenged !== undefined && (obj.lastTimeChallenged = message.lastTimeChallenged);
    message.coolDownTolerance !== undefined && (obj.coolDownTolerance = message.coolDownTolerance);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Runner>, I>>(object: I): Runner {
    const message = createBaseRunner();
    message.index = object.index ?? "";
    message.address = object.address ?? "";
    message.score = object.score ?? "";
    message.rewardMultiplier = object.rewardMultiplier ?? "";
    message.stakedAmount = object.stakedAmount ?? "";
    message.netEarnings = object.netEarnings ?? "";
    message.ipAddr = object.ipAddr ?? "";
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
