/* eslint-disable */
import { Runner } from "../poa/runner";
import { Challenger } from "../poa/challenger";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface Guard {
  index: string;
  guardId: string;
  runner: Runner | undefined;
  v2XChallenger: Challenger | undefined;
  v2NChallenger: Challenger | undefined;
}

const baseGuard: object = { index: "", guardId: "" };

export const Guard = {
  encode(message: Guard, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.guardId !== "") {
      writer.uint32(18).string(message.guardId);
    }
    if (message.runner !== undefined) {
      Runner.encode(message.runner, writer.uint32(26).fork()).ldelim();
    }
    if (message.v2XChallenger !== undefined) {
      Challenger.encode(
        message.v2XChallenger,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.v2NChallenger !== undefined) {
      Challenger.encode(
        message.v2NChallenger,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Guard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGuard } as Guard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.guardId = reader.string();
          break;
        case 3:
          message.runner = Runner.decode(reader, reader.uint32());
          break;
        case 4:
          message.v2XChallenger = Challenger.decode(reader, reader.uint32());
          break;
        case 5:
          message.v2NChallenger = Challenger.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Guard {
    const message = { ...baseGuard } as Guard;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.guardId !== undefined && object.guardId !== null) {
      message.guardId = String(object.guardId);
    } else {
      message.guardId = "";
    }
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromJSON(object.runner);
    } else {
      message.runner = undefined;
    }
    if (object.v2XChallenger !== undefined && object.v2XChallenger !== null) {
      message.v2XChallenger = Challenger.fromJSON(object.v2XChallenger);
    } else {
      message.v2XChallenger = undefined;
    }
    if (object.v2NChallenger !== undefined && object.v2NChallenger !== null) {
      message.v2NChallenger = Challenger.fromJSON(object.v2NChallenger);
    } else {
      message.v2NChallenger = undefined;
    }
    return message;
  },

  toJSON(message: Guard): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.guardId !== undefined && (obj.guardId = message.guardId);
    message.runner !== undefined &&
      (obj.runner = message.runner ? Runner.toJSON(message.runner) : undefined);
    message.v2XChallenger !== undefined &&
      (obj.v2XChallenger = message.v2XChallenger
        ? Challenger.toJSON(message.v2XChallenger)
        : undefined);
    message.v2NChallenger !== undefined &&
      (obj.v2NChallenger = message.v2NChallenger
        ? Challenger.toJSON(message.v2NChallenger)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Guard>): Guard {
    const message = { ...baseGuard } as Guard;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.guardId !== undefined && object.guardId !== null) {
      message.guardId = object.guardId;
    } else {
      message.guardId = "";
    }
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromPartial(object.runner);
    } else {
      message.runner = undefined;
    }
    if (object.v2XChallenger !== undefined && object.v2XChallenger !== null) {
      message.v2XChallenger = Challenger.fromPartial(object.v2XChallenger);
    } else {
      message.v2XChallenger = undefined;
    }
    if (object.v2NChallenger !== undefined && object.v2NChallenger !== null) {
      message.v2NChallenger = Challenger.fromPartial(object.v2NChallenger);
    } else {
      message.v2NChallenger = undefined;
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
