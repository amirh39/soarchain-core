/* eslint-disable */
import { Challenger } from "../poa/challenger";
import { Runner } from "../poa/runner";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface VrfData {
  index: string;
  creator: string;
  vrv: string;
  multiplier: string;
  proof: string;
  pubkey: string;
  message: string;
  parsedVrv: string;
  floatVrv: string;
  finalVrv: string;
  finalVrvFloat: string;
  selectedChallenger: Challenger | undefined;
  selectedRunner: Runner | undefined;
}

const baseVrfData: object = {
  index: "",
  creator: "",
  vrv: "",
  multiplier: "",
  proof: "",
  pubkey: "",
  message: "",
  parsedVrv: "",
  floatVrv: "",
  finalVrv: "",
  finalVrvFloat: "",
};

export const VrfData = {
  encode(message: VrfData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.vrv !== "") {
      writer.uint32(26).string(message.vrv);
    }
    if (message.multiplier !== "") {
      writer.uint32(34).string(message.multiplier);
    }
    if (message.proof !== "") {
      writer.uint32(42).string(message.proof);
    }
    if (message.pubkey !== "") {
      writer.uint32(50).string(message.pubkey);
    }
    if (message.message !== "") {
      writer.uint32(58).string(message.message);
    }
    if (message.parsedVrv !== "") {
      writer.uint32(66).string(message.parsedVrv);
    }
    if (message.floatVrv !== "") {
      writer.uint32(74).string(message.floatVrv);
    }
    if (message.finalVrv !== "") {
      writer.uint32(82).string(message.finalVrv);
    }
    if (message.finalVrvFloat !== "") {
      writer.uint32(90).string(message.finalVrvFloat);
    }
    if (message.selectedChallenger !== undefined) {
      Challenger.encode(
        message.selectedChallenger,
        writer.uint32(98).fork()
      ).ldelim();
    }
    if (message.selectedRunner !== undefined) {
      Runner.encode(message.selectedRunner, writer.uint32(106).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VrfData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVrfData } as VrfData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.vrv = reader.string();
          break;
        case 4:
          message.multiplier = reader.string();
          break;
        case 5:
          message.proof = reader.string();
          break;
        case 6:
          message.pubkey = reader.string();
          break;
        case 7:
          message.message = reader.string();
          break;
        case 8:
          message.parsedVrv = reader.string();
          break;
        case 9:
          message.floatVrv = reader.string();
          break;
        case 10:
          message.finalVrv = reader.string();
          break;
        case 11:
          message.finalVrvFloat = reader.string();
          break;
        case 12:
          message.selectedChallenger = Challenger.decode(
            reader,
            reader.uint32()
          );
          break;
        case 13:
          message.selectedRunner = Runner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VrfData {
    const message = { ...baseVrfData } as VrfData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.vrv !== undefined && object.vrv !== null) {
      message.vrv = String(object.vrv);
    } else {
      message.vrv = "";
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = String(object.multiplier);
    } else {
      message.multiplier = "";
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = String(object.proof);
    } else {
      message.proof = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = String(object.pubkey);
    } else {
      message.pubkey = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.parsedVrv !== undefined && object.parsedVrv !== null) {
      message.parsedVrv = String(object.parsedVrv);
    } else {
      message.parsedVrv = "";
    }
    if (object.floatVrv !== undefined && object.floatVrv !== null) {
      message.floatVrv = String(object.floatVrv);
    } else {
      message.floatVrv = "";
    }
    if (object.finalVrv !== undefined && object.finalVrv !== null) {
      message.finalVrv = String(object.finalVrv);
    } else {
      message.finalVrv = "";
    }
    if (object.finalVrvFloat !== undefined && object.finalVrvFloat !== null) {
      message.finalVrvFloat = String(object.finalVrvFloat);
    } else {
      message.finalVrvFloat = "";
    }
    if (
      object.selectedChallenger !== undefined &&
      object.selectedChallenger !== null
    ) {
      message.selectedChallenger = Challenger.fromJSON(
        object.selectedChallenger
      );
    } else {
      message.selectedChallenger = undefined;
    }
    if (object.selectedRunner !== undefined && object.selectedRunner !== null) {
      message.selectedRunner = Runner.fromJSON(object.selectedRunner);
    } else {
      message.selectedRunner = undefined;
    }
    return message;
  },

  toJSON(message: VrfData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.vrv !== undefined && (obj.vrv = message.vrv);
    message.multiplier !== undefined && (obj.multiplier = message.multiplier);
    message.proof !== undefined && (obj.proof = message.proof);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    message.message !== undefined && (obj.message = message.message);
    message.parsedVrv !== undefined && (obj.parsedVrv = message.parsedVrv);
    message.floatVrv !== undefined && (obj.floatVrv = message.floatVrv);
    message.finalVrv !== undefined && (obj.finalVrv = message.finalVrv);
    message.finalVrvFloat !== undefined &&
      (obj.finalVrvFloat = message.finalVrvFloat);
    message.selectedChallenger !== undefined &&
      (obj.selectedChallenger = message.selectedChallenger
        ? Challenger.toJSON(message.selectedChallenger)
        : undefined);
    message.selectedRunner !== undefined &&
      (obj.selectedRunner = message.selectedRunner
        ? Runner.toJSON(message.selectedRunner)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<VrfData>): VrfData {
    const message = { ...baseVrfData } as VrfData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.vrv !== undefined && object.vrv !== null) {
      message.vrv = object.vrv;
    } else {
      message.vrv = "";
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = object.multiplier;
    } else {
      message.multiplier = "";
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = object.proof;
    } else {
      message.proof = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = object.pubkey;
    } else {
      message.pubkey = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.parsedVrv !== undefined && object.parsedVrv !== null) {
      message.parsedVrv = object.parsedVrv;
    } else {
      message.parsedVrv = "";
    }
    if (object.floatVrv !== undefined && object.floatVrv !== null) {
      message.floatVrv = object.floatVrv;
    } else {
      message.floatVrv = "";
    }
    if (object.finalVrv !== undefined && object.finalVrv !== null) {
      message.finalVrv = object.finalVrv;
    } else {
      message.finalVrv = "";
    }
    if (object.finalVrvFloat !== undefined && object.finalVrvFloat !== null) {
      message.finalVrvFloat = object.finalVrvFloat;
    } else {
      message.finalVrvFloat = "";
    }
    if (
      object.selectedChallenger !== undefined &&
      object.selectedChallenger !== null
    ) {
      message.selectedChallenger = Challenger.fromPartial(
        object.selectedChallenger
      );
    } else {
      message.selectedChallenger = undefined;
    }
    if (object.selectedRunner !== undefined && object.selectedRunner !== null) {
      message.selectedRunner = Runner.fromPartial(object.selectedRunner);
    } else {
      message.selectedRunner = undefined;
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
