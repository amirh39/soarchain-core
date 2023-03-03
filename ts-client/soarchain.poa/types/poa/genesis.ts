/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Challenger } from "./challenger";
import { Client } from "./client";
import { EpochData } from "./epoch_data";
import { Guard } from "./guard";
import { MotusWallet } from "./motus_wallet";
import { Params } from "./params";
import { Runner } from "./runner";
import { VrfData } from "./vrf_data";
import { VrfUser } from "./vrf_user";

export const protobufPackage = "soarchain.poa";

/** GenesisState defines the poa module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  clientList: Client[];
  challengerList: Challenger[];
  runnerList: Runner[];
  guardList: Guard[];
  vrfDataList: VrfData[];
  vrfUserList: VrfUser[];
  epochData:
    | EpochData
    | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  motusWalletList: MotusWallet[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    clientList: [],
    challengerList: [],
    runnerList: [],
    guardList: [],
    vrfDataList: [],
    vrfUserList: [],
    epochData: undefined,
    motusWalletList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    for (const v of message.vrfDataList) {
      VrfData.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.vrfUserList) {
      VrfUser.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.epochData !== undefined) {
      EpochData.encode(message.epochData, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.motusWalletList) {
      MotusWallet.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
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
          message.challengerList.push(Challenger.decode(reader, reader.uint32()));
          break;
        case 4:
          message.runnerList.push(Runner.decode(reader, reader.uint32()));
          break;
        case 5:
          message.guardList.push(Guard.decode(reader, reader.uint32()));
          break;
        case 6:
          message.vrfDataList.push(VrfData.decode(reader, reader.uint32()));
          break;
        case 7:
          message.vrfUserList.push(VrfUser.decode(reader, reader.uint32()));
          break;
        case 8:
          message.epochData = EpochData.decode(reader, reader.uint32());
          break;
        case 9:
          message.motusWalletList.push(MotusWallet.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      clientList: Array.isArray(object?.clientList) ? object.clientList.map((e: any) => Client.fromJSON(e)) : [],
      challengerList: Array.isArray(object?.challengerList)
        ? object.challengerList.map((e: any) => Challenger.fromJSON(e))
        : [],
      runnerList: Array.isArray(object?.runnerList) ? object.runnerList.map((e: any) => Runner.fromJSON(e)) : [],
      guardList: Array.isArray(object?.guardList) ? object.guardList.map((e: any) => Guard.fromJSON(e)) : [],
      vrfDataList: Array.isArray(object?.vrfDataList) ? object.vrfDataList.map((e: any) => VrfData.fromJSON(e)) : [],
      vrfUserList: Array.isArray(object?.vrfUserList) ? object.vrfUserList.map((e: any) => VrfUser.fromJSON(e)) : [],
      epochData: isSet(object.epochData) ? EpochData.fromJSON(object.epochData) : undefined,
      motusWalletList: Array.isArray(object?.motusWalletList)
        ? object.motusWalletList.map((e: any) => MotusWallet.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.clientList) {
      obj.clientList = message.clientList.map((e) => e ? Client.toJSON(e) : undefined);
    } else {
      obj.clientList = [];
    }
    if (message.challengerList) {
      obj.challengerList = message.challengerList.map((e) => e ? Challenger.toJSON(e) : undefined);
    } else {
      obj.challengerList = [];
    }
    if (message.runnerList) {
      obj.runnerList = message.runnerList.map((e) => e ? Runner.toJSON(e) : undefined);
    } else {
      obj.runnerList = [];
    }
    if (message.guardList) {
      obj.guardList = message.guardList.map((e) => e ? Guard.toJSON(e) : undefined);
    } else {
      obj.guardList = [];
    }
    if (message.vrfDataList) {
      obj.vrfDataList = message.vrfDataList.map((e) => e ? VrfData.toJSON(e) : undefined);
    } else {
      obj.vrfDataList = [];
    }
    if (message.vrfUserList) {
      obj.vrfUserList = message.vrfUserList.map((e) => e ? VrfUser.toJSON(e) : undefined);
    } else {
      obj.vrfUserList = [];
    }
    message.epochData !== undefined
      && (obj.epochData = message.epochData ? EpochData.toJSON(message.epochData) : undefined);
    if (message.motusWalletList) {
      obj.motusWalletList = message.motusWalletList.map((e) => e ? MotusWallet.toJSON(e) : undefined);
    } else {
      obj.motusWalletList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.clientList = object.clientList?.map((e) => Client.fromPartial(e)) || [];
    message.challengerList = object.challengerList?.map((e) => Challenger.fromPartial(e)) || [];
    message.runnerList = object.runnerList?.map((e) => Runner.fromPartial(e)) || [];
    message.guardList = object.guardList?.map((e) => Guard.fromPartial(e)) || [];
    message.vrfDataList = object.vrfDataList?.map((e) => VrfData.fromPartial(e)) || [];
    message.vrfUserList = object.vrfUserList?.map((e) => VrfUser.fromPartial(e)) || [];
    message.epochData = (object.epochData !== undefined && object.epochData !== null)
      ? EpochData.fromPartial(object.epochData)
      : undefined;
    message.motusWalletList = object.motusWalletList?.map((e) => MotusWallet.fromPartial(e)) || [];
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
