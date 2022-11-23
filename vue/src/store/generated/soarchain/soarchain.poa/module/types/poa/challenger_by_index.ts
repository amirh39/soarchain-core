/* eslint-disable */
import { Challenger } from "../poa/challenger";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface ChallengerByIndex {
  index: string;
  challenger: Challenger | undefined;
}

const baseChallengerByIndex: object = { index: "" };

export const ChallengerByIndex = {
  encode(message: ChallengerByIndex, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.challenger !== undefined) {
      Challenger.encode(message.challenger, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChallengerByIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChallengerByIndex } as ChallengerByIndex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.challenger = Challenger.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChallengerByIndex {
    const message = { ...baseChallengerByIndex } as ChallengerByIndex;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromJSON(object.challenger);
    } else {
      message.challenger = undefined;
    }
    return message;
  },

  toJSON(message: ChallengerByIndex): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.challenger !== undefined &&
      (obj.challenger = message.challenger
        ? Challenger.toJSON(message.challenger)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<ChallengerByIndex>): ChallengerByIndex {
    const message = { ...baseChallengerByIndex } as ChallengerByIndex;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromPartial(object.challenger);
    } else {
      message.challenger = undefined;
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
