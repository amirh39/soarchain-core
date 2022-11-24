/* eslint-disable */
import { Params } from "../poa/params";
import { Client } from "../poa/client";
import { Challenger } from "../poa/challenger";
import { Runner } from "../poa/runner";
import { Guard } from "../poa/guard";
import { TotalClients } from "../poa/total_clients";
import { TotalChallengers } from "../poa/total_challengers";
import { TotalRunners } from "../poa/total_runners";
import { ChallengerByIndex } from "../poa/challenger_by_index";
import { RunnerByIndex } from "../poa/runner_by_index";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

/** GenesisState defines the poa module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  clientList: Client[];
  challengerList: Challenger[];
  runnerList: Runner[];
  guardList: Guard[];
  totalClients: TotalClients | undefined;
  totalChallengers: TotalChallengers | undefined;
  totalRunners: TotalRunners | undefined;
  challengerByIndexList: ChallengerByIndex[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  runnerByIndexList: RunnerByIndex[];
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
    for (const v of message.runnerList) {
      Runner.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.guardList) {
      Guard.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.totalClients !== undefined) {
      TotalClients.encode(
        message.totalClients,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.totalChallengers !== undefined) {
      TotalChallengers.encode(
        message.totalChallengers,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.totalRunners !== undefined) {
      TotalRunners.encode(
        message.totalRunners,
        writer.uint32(66).fork()
      ).ldelim();
    }
    for (const v of message.challengerByIndexList) {
      ChallengerByIndex.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    for (const v of message.runnerByIndexList) {
      RunnerByIndex.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.clientList = [];
    message.challengerList = [];
    message.runnerList = [];
    message.guardList = [];
    message.challengerByIndexList = [];
    message.runnerByIndexList = [];
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
        case 4:
          message.runnerList.push(Runner.decode(reader, reader.uint32()));
          break;
        case 5:
          message.guardList.push(Guard.decode(reader, reader.uint32()));
          break;
        case 6:
          message.totalClients = TotalClients.decode(reader, reader.uint32());
          break;
        case 7:
          message.totalChallengers = TotalChallengers.decode(
            reader,
            reader.uint32()
          );
          break;
        case 8:
          message.totalRunners = TotalRunners.decode(reader, reader.uint32());
          break;
        case 9:
          message.challengerByIndexList.push(
            ChallengerByIndex.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.runnerByIndexList.push(
            RunnerByIndex.decode(reader, reader.uint32())
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
    message.runnerList = [];
    message.guardList = [];
    message.challengerByIndexList = [];
    message.runnerByIndexList = [];
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
    if (object.runnerList !== undefined && object.runnerList !== null) {
      for (const e of object.runnerList) {
        message.runnerList.push(Runner.fromJSON(e));
      }
    }
    if (object.guardList !== undefined && object.guardList !== null) {
      for (const e of object.guardList) {
        message.guardList.push(Guard.fromJSON(e));
      }
    }
    if (object.totalClients !== undefined && object.totalClients !== null) {
      message.totalClients = TotalClients.fromJSON(object.totalClients);
    } else {
      message.totalClients = undefined;
    }
    if (
      object.totalChallengers !== undefined &&
      object.totalChallengers !== null
    ) {
      message.totalChallengers = TotalChallengers.fromJSON(
        object.totalChallengers
      );
    } else {
      message.totalChallengers = undefined;
    }
    if (object.totalRunners !== undefined && object.totalRunners !== null) {
      message.totalRunners = TotalRunners.fromJSON(object.totalRunners);
    } else {
      message.totalRunners = undefined;
    }
    if (
      object.challengerByIndexList !== undefined &&
      object.challengerByIndexList !== null
    ) {
      for (const e of object.challengerByIndexList) {
        message.challengerByIndexList.push(ChallengerByIndex.fromJSON(e));
      }
    }
    if (
      object.runnerByIndexList !== undefined &&
      object.runnerByIndexList !== null
    ) {
      for (const e of object.runnerByIndexList) {
        message.runnerByIndexList.push(RunnerByIndex.fromJSON(e));
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
    if (message.runnerList) {
      obj.runnerList = message.runnerList.map((e) =>
        e ? Runner.toJSON(e) : undefined
      );
    } else {
      obj.runnerList = [];
    }
    if (message.guardList) {
      obj.guardList = message.guardList.map((e) =>
        e ? Guard.toJSON(e) : undefined
      );
    } else {
      obj.guardList = [];
    }
    message.totalClients !== undefined &&
      (obj.totalClients = message.totalClients
        ? TotalClients.toJSON(message.totalClients)
        : undefined);
    message.totalChallengers !== undefined &&
      (obj.totalChallengers = message.totalChallengers
        ? TotalChallengers.toJSON(message.totalChallengers)
        : undefined);
    message.totalRunners !== undefined &&
      (obj.totalRunners = message.totalRunners
        ? TotalRunners.toJSON(message.totalRunners)
        : undefined);
    if (message.challengerByIndexList) {
      obj.challengerByIndexList = message.challengerByIndexList.map((e) =>
        e ? ChallengerByIndex.toJSON(e) : undefined
      );
    } else {
      obj.challengerByIndexList = [];
    }
    if (message.runnerByIndexList) {
      obj.runnerByIndexList = message.runnerByIndexList.map((e) =>
        e ? RunnerByIndex.toJSON(e) : undefined
      );
    } else {
      obj.runnerByIndexList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.clientList = [];
    message.challengerList = [];
    message.runnerList = [];
    message.guardList = [];
    message.challengerByIndexList = [];
    message.runnerByIndexList = [];
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
    if (object.runnerList !== undefined && object.runnerList !== null) {
      for (const e of object.runnerList) {
        message.runnerList.push(Runner.fromPartial(e));
      }
    }
    if (object.guardList !== undefined && object.guardList !== null) {
      for (const e of object.guardList) {
        message.guardList.push(Guard.fromPartial(e));
      }
    }
    if (object.totalClients !== undefined && object.totalClients !== null) {
      message.totalClients = TotalClients.fromPartial(object.totalClients);
    } else {
      message.totalClients = undefined;
    }
    if (
      object.totalChallengers !== undefined &&
      object.totalChallengers !== null
    ) {
      message.totalChallengers = TotalChallengers.fromPartial(
        object.totalChallengers
      );
    } else {
      message.totalChallengers = undefined;
    }
    if (object.totalRunners !== undefined && object.totalRunners !== null) {
      message.totalRunners = TotalRunners.fromPartial(object.totalRunners);
    } else {
      message.totalRunners = undefined;
    }
    if (
      object.challengerByIndexList !== undefined &&
      object.challengerByIndexList !== null
    ) {
      for (const e of object.challengerByIndexList) {
        message.challengerByIndexList.push(ChallengerByIndex.fromPartial(e));
      }
    }
    if (
      object.runnerByIndexList !== undefined &&
      object.runnerByIndexList !== null
    ) {
      for (const e of object.runnerByIndexList) {
        message.runnerByIndexList.push(RunnerByIndex.fromPartial(e));
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
