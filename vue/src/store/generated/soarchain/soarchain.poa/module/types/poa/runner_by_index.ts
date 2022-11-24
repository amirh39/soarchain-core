/* eslint-disable */
import { Runner } from "../poa/runner";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface RunnerByIndex {
  index: string;
  runner: Runner | undefined;
}

const baseRunnerByIndex: object = { index: "" };

export const RunnerByIndex = {
  encode(message: RunnerByIndex, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.runner !== undefined) {
      Runner.encode(message.runner, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RunnerByIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRunnerByIndex } as RunnerByIndex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.runner = Runner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RunnerByIndex {
    const message = { ...baseRunnerByIndex } as RunnerByIndex;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromJSON(object.runner);
    } else {
      message.runner = undefined;
    }
    return message;
  },

  toJSON(message: RunnerByIndex): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.runner !== undefined &&
      (obj.runner = message.runner ? Runner.toJSON(message.runner) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<RunnerByIndex>): RunnerByIndex {
    const message = { ...baseRunnerByIndex } as RunnerByIndex;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromPartial(object.runner);
    } else {
      message.runner = undefined;
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
