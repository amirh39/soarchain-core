/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Challenger } from "./challenger";
import { Runner } from "./runner";

export const protobufPackage = "soarchain.poa";

export interface Guard {
  index: string;
  guardId: string;
  runner: Runner | undefined;
  v2XChallenger: Challenger | undefined;
  v2NChallenger: Challenger | undefined;
}

function createBaseGuard(): Guard {
  return { index: "", guardId: "", runner: undefined, v2XChallenger: undefined, v2NChallenger: undefined };
}

export const Guard = {
  encode(message: Guard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      Challenger.encode(message.v2XChallenger, writer.uint32(34).fork()).ldelim();
    }
    if (message.v2NChallenger !== undefined) {
      Challenger.encode(message.v2NChallenger, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Guard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGuard();
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
    return {
      index: isSet(object.index) ? String(object.index) : "",
      guardId: isSet(object.guardId) ? String(object.guardId) : "",
      runner: isSet(object.runner) ? Runner.fromJSON(object.runner) : undefined,
      v2XChallenger: isSet(object.v2XChallenger) ? Challenger.fromJSON(object.v2XChallenger) : undefined,
      v2NChallenger: isSet(object.v2NChallenger) ? Challenger.fromJSON(object.v2NChallenger) : undefined,
    };
  },

  toJSON(message: Guard): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.guardId !== undefined && (obj.guardId = message.guardId);
    message.runner !== undefined && (obj.runner = message.runner ? Runner.toJSON(message.runner) : undefined);
    message.v2XChallenger !== undefined
      && (obj.v2XChallenger = message.v2XChallenger ? Challenger.toJSON(message.v2XChallenger) : undefined);
    message.v2NChallenger !== undefined
      && (obj.v2NChallenger = message.v2NChallenger ? Challenger.toJSON(message.v2NChallenger) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Guard>, I>>(object: I): Guard {
    const message = createBaseGuard();
    message.index = object.index ?? "";
    message.guardId = object.guardId ?? "";
    message.runner = (object.runner !== undefined && object.runner !== null)
      ? Runner.fromPartial(object.runner)
      : undefined;
    message.v2XChallenger = (object.v2XChallenger !== undefined && object.v2XChallenger !== null)
      ? Challenger.fromPartial(object.v2XChallenger)
      : undefined;
    message.v2NChallenger = (object.v2NChallenger !== undefined && object.v2NChallenger !== null)
      ? Challenger.fromPartial(object.v2NChallenger)
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
