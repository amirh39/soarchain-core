/* eslint-disable */
import { Params } from "../poa/params";
import { Client } from "../poa/client";
import { Challenger } from "../poa/challenger";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

/** GenesisState defines the poa module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  clientList: Client[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  challengerList: Challenger[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.clientList) {
      Client.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.challengerList) {
      Challenger.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.clientList = [];
    message.challengerList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.clientList.push(Client.decode(reader, reader.uint32()));
          break;
        case 3:
          message.challengerList.push(
            Challenger.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.clientList = [];
    message.challengerList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.clientList !== undefined && object.clientList !== null) {
      for (const e of object.clientList) {
        message.clientList.push(Client.fromJSON(e));
      }
    }
    if (object.challengerList !== undefined && object.challengerList !== null) {
      for (const e of object.challengerList) {
        message.challengerList.push(Challenger.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.clientList) {
      obj.clientList = message.clientList.map((e) =>
        e ? Client.toJSON(e) : undefined
      );
    } else {
      obj.clientList = [];
    }
    if (message.challengerList) {
      obj.challengerList = message.challengerList.map((e) =>
        e ? Challenger.toJSON(e) : undefined
      );
    } else {
      obj.challengerList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.clientList = [];
    message.challengerList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.clientList !== undefined && object.clientList !== null) {
      for (const e of object.clientList) {
        message.clientList.push(Client.fromPartial(e));
      }
    }
    if (object.challengerList !== undefined && object.challengerList !== null) {
      for (const e of object.challengerList) {
        message.challengerList.push(Challenger.fromPartial(e));
      }
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
