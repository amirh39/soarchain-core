/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Challenger } from "./challenger";
import { Runner } from "./runner";

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

function createBaseVrfData(): VrfData {
  return {
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
    selectedChallenger: undefined,
    selectedRunner: undefined,
  };
}

export const VrfData = {
  encode(message: VrfData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      Challenger.encode(message.selectedChallenger, writer.uint32(98).fork()).ldelim();
    }
    if (message.selectedRunner !== undefined) {
      Runner.encode(message.selectedRunner, writer.uint32(106).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VrfData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVrfData();
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
          message.selectedChallenger = Challenger.decode(reader, reader.uint32());
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
    return {
      index: isSet(object.index) ? String(object.index) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      vrv: isSet(object.vrv) ? String(object.vrv) : "",
      multiplier: isSet(object.multiplier) ? String(object.multiplier) : "",
      proof: isSet(object.proof) ? String(object.proof) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
      message: isSet(object.message) ? String(object.message) : "",
      parsedVrv: isSet(object.parsedVrv) ? String(object.parsedVrv) : "",
      floatVrv: isSet(object.floatVrv) ? String(object.floatVrv) : "",
      finalVrv: isSet(object.finalVrv) ? String(object.finalVrv) : "",
      finalVrvFloat: isSet(object.finalVrvFloat) ? String(object.finalVrvFloat) : "",
      selectedChallenger: isSet(object.selectedChallenger) ? Challenger.fromJSON(object.selectedChallenger) : undefined,
      selectedRunner: isSet(object.selectedRunner) ? Runner.fromJSON(object.selectedRunner) : undefined,
    };
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
    message.finalVrvFloat !== undefined && (obj.finalVrvFloat = message.finalVrvFloat);
    message.selectedChallenger !== undefined && (obj.selectedChallenger = message.selectedChallenger
      ? Challenger.toJSON(message.selectedChallenger)
      : undefined);
    message.selectedRunner !== undefined
      && (obj.selectedRunner = message.selectedRunner ? Runner.toJSON(message.selectedRunner) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VrfData>, I>>(object: I): VrfData {
    const message = createBaseVrfData();
    message.index = object.index ?? "";
    message.creator = object.creator ?? "";
    message.vrv = object.vrv ?? "";
    message.multiplier = object.multiplier ?? "";
    message.proof = object.proof ?? "";
    message.pubkey = object.pubkey ?? "";
    message.message = object.message ?? "";
    message.parsedVrv = object.parsedVrv ?? "";
    message.floatVrv = object.floatVrv ?? "";
    message.finalVrv = object.finalVrv ?? "";
    message.finalVrvFloat = object.finalVrvFloat ?? "";
    message.selectedChallenger = (object.selectedChallenger !== undefined && object.selectedChallenger !== null)
      ? Challenger.fromPartial(object.selectedChallenger)
      : undefined;
    message.selectedRunner = (object.selectedRunner !== undefined && object.selectedRunner !== null)
      ? Runner.fromPartial(object.selectedRunner)
      : undefined;
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
