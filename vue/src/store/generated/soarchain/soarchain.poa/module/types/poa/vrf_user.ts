/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface VrfUser {
  index: string;
  address: string;
  count: string;
}

const baseVrfUser: object = { index: "", address: "", count: "" };

export const VrfUser = {
  encode(message: VrfUser, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.count !== "") {
      writer.uint32(26).string(message.count);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VrfUser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVrfUser } as VrfUser;
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
          message.count = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VrfUser {
    const message = { ...baseVrfUser } as VrfUser;
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
    if (object.count !== undefined && object.count !== null) {
      message.count = String(object.count);
    } else {
      message.count = "";
    }
    return message;
  },

  toJSON(message: VrfUser): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(object: DeepPartial<VrfUser>): VrfUser {
    const message = { ...baseVrfUser } as VrfUser;
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
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = "";
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
