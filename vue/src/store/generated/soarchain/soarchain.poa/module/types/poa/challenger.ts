/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Challenger {
  index: string;
  address: string;
  challengerId: string;
  score: string;
  stakedAmount: string;
  netEarnings: string;
}

const baseChallenger: object = {
  index: "",
  address: "",
  challengerId: "",
  score: "",
  stakedAmount: "",
  netEarnings: "",
};

export const Challenger = {
  encode(message: Challenger, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.challengerId !== "") {
      writer.uint32(26).string(message.challengerId);
    }
    if (message.score !== "") {
      writer.uint32(34).string(message.score);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(42).string(message.stakedAmount);
    }
    if (message.netEarnings !== "") {
      writer.uint32(50).string(message.netEarnings);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Challenger {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChallenger } as Challenger;
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
          message.challengerId = reader.string();
          break;
        case 4:
          message.score = reader.string();
          break;
        case 5:
          message.stakedAmount = reader.string();
          break;
        case 6:
          message.netEarnings = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Challenger {
    const message = { ...baseChallenger } as Challenger;
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
    if (object.challengerId !== undefined && object.challengerId !== null) {
      message.challengerId = String(object.challengerId);
    } else {
      message.challengerId = "";
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
    return message;
  },

  toJSON(message: Challenger): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.challengerId !== undefined &&
      (obj.challengerId = message.challengerId);
    message.score !== undefined && (obj.score = message.score);
    message.stakedAmount !== undefined &&
      (obj.stakedAmount = message.stakedAmount);
    message.netEarnings !== undefined &&
      (obj.netEarnings = message.netEarnings);
    return obj;
  },

  fromPartial(object: DeepPartial<Challenger>): Challenger {
    const message = { ...baseChallenger } as Challenger;
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
    if (object.challengerId !== undefined && object.challengerId !== null) {
      message.challengerId = object.challengerId;
    } else {
      message.challengerId = "";
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
