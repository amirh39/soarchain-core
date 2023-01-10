/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Client {
  index: string;
  address: string;
  registrant: string;
  score: string;
  rewardMultiplier: string;
  netEarnings: string;
  lastTimeChallenged: string;
  coolDownTolerance: string;
}

const baseClient: object = {
  index: "",
  address: "",
  registrant: "",
  score: "",
  rewardMultiplier: "",
  netEarnings: "",
  lastTimeChallenged: "",
  coolDownTolerance: "",
};

export const Client = {
  encode(message: Client, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.registrant !== "") {
      writer.uint32(26).string(message.registrant);
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

  decode(input: Reader | Uint8Array, length?: number): Client {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClient } as Client;
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
          message.registrant = reader.string();
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
    const message = { ...baseClient } as Client;
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
    if (object.registrant !== undefined && object.registrant !== null) {
      message.registrant = String(object.registrant);
    } else {
      message.registrant = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = String(object.score);
    } else {
      message.score = "";
    }
    if (
      object.rewardMultiplier !== undefined &&
      object.rewardMultiplier !== null
    ) {
      message.rewardMultiplier = String(object.rewardMultiplier);
    } else {
      message.rewardMultiplier = "";
    }
    if (object.netEarnings !== undefined && object.netEarnings !== null) {
      message.netEarnings = String(object.netEarnings);
    } else {
      message.netEarnings = "";
    }
    if (
      object.lastTimeChallenged !== undefined &&
      object.lastTimeChallenged !== null
    ) {
      message.lastTimeChallenged = String(object.lastTimeChallenged);
    } else {
      message.lastTimeChallenged = "";
    }
    if (
      object.coolDownTolerance !== undefined &&
      object.coolDownTolerance !== null
    ) {
      message.coolDownTolerance = String(object.coolDownTolerance);
    } else {
      message.coolDownTolerance = "";
    }
    return message;
  },

  toJSON(message: Client): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.registrant !== undefined && (obj.registrant = message.registrant);
    message.score !== undefined && (obj.score = message.score);
    message.rewardMultiplier !== undefined &&
      (obj.rewardMultiplier = message.rewardMultiplier);
    message.netEarnings !== undefined &&
      (obj.netEarnings = message.netEarnings);
    message.lastTimeChallenged !== undefined &&
      (obj.lastTimeChallenged = message.lastTimeChallenged);
    message.coolDownTolerance !== undefined &&
      (obj.coolDownTolerance = message.coolDownTolerance);
    return obj;
  },

  fromPartial(object: DeepPartial<Client>): Client {
    const message = { ...baseClient } as Client;
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
    if (object.registrant !== undefined && object.registrant !== null) {
      message.registrant = object.registrant;
    } else {
      message.registrant = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = "";
    }
    if (
      object.rewardMultiplier !== undefined &&
      object.rewardMultiplier !== null
    ) {
      message.rewardMultiplier = object.rewardMultiplier;
    } else {
      message.rewardMultiplier = "";
    }
    if (object.netEarnings !== undefined && object.netEarnings !== null) {
      message.netEarnings = object.netEarnings;
    } else {
      message.netEarnings = "";
    }
    if (
      object.lastTimeChallenged !== undefined &&
      object.lastTimeChallenged !== null
    ) {
      message.lastTimeChallenged = object.lastTimeChallenged;
    } else {
      message.lastTimeChallenged = "";
    }
    if (
      object.coolDownTolerance !== undefined &&
      object.coolDownTolerance !== null
    ) {
      message.coolDownTolerance = object.coolDownTolerance;
    } else {
      message.coolDownTolerance = "";
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
