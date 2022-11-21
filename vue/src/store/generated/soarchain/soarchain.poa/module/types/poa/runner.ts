/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Runner {
  index: string;
  address: string;
  score: string;
  stakedAmount: string;
  netEarnings: string;
  ipAddr: string;
}

const baseRunner: object = {
  index: "",
  address: "",
  score: "",
  stakedAmount: "",
  netEarnings: "",
  ipAddr: "",
};

export const Runner = {
  encode(message: Runner, writer: Writer = Writer.create()): Writer {
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
    if (message.ipAddr !== "") {
      writer.uint32(50).string(message.ipAddr);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Runner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRunner } as Runner;
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
          message.ipAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Runner {
    const message = { ...baseRunner } as Runner;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = String(object.score);
    } else {
      message.score = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = String(object.stakedAmount);
    } else {
      message.stakedAmount = "";
    }
    if (object.netEarnings !== undefined && object.netEarnings !== null) {
      message.netEarnings = String(object.netEarnings);
    } else {
      message.netEarnings = "";
    }
    if (object.ipAddr !== undefined && object.ipAddr !== null) {
      message.ipAddr = String(object.ipAddr);
    } else {
      message.ipAddr = "";
    }
    return message;
  },

  toJSON(message: Runner): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    message.stakedAmount !== undefined &&
      (obj.stakedAmount = message.stakedAmount);
    message.netEarnings !== undefined &&
      (obj.netEarnings = message.netEarnings);
    message.ipAddr !== undefined && (obj.ipAddr = message.ipAddr);
    return obj;
  },

  fromPartial(object: DeepPartial<Runner>): Runner {
    const message = { ...baseRunner } as Runner;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = object.stakedAmount;
    } else {
      message.stakedAmount = "";
    }
    if (object.netEarnings !== undefined && object.netEarnings !== null) {
      message.netEarnings = object.netEarnings;
    } else {
      message.netEarnings = "";
    }
    if (object.ipAddr !== undefined && object.ipAddr !== null) {
      message.ipAddr = object.ipAddr;
    } else {
      message.ipAddr = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
