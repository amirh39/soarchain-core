/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Challenger } from "./challenger";
import { Runner } from "./runner";

export const protobufPackage = "soarchain.poa";

export interface MsgGenClient {
  creator: string;
  pubkey: string;
}

export interface MsgGenClientResponse {
}

export interface MsgChallengeService {
  creator: string;
  clientPubkey: string;
  clientCommunicationMode: string;
  challengeResult: string;
}

export interface MsgChallengeServiceResponse {
}

export interface MsgUnregisterClient {
  creator: string;
  pubkey: string;
}

export interface MsgUnregisterClientResponse {
}

export interface MsgUnregisterChallenger {
  creator: string;
  challengerAddress: string;
}

export interface MsgUnregisterChallengerResponse {
}

export interface MsgGenGuard {
  creator: string;
  guardPubKey: string;
  v2XAddr: string;
  v2XStake: string;
  v2XIp: string;
  v2NAddr: string;
  v2NStake: string;
  v2NIp: string;
  runnerAddr: string;
  runnerStake: string;
  runnerIp: string;
}

export interface MsgGenGuardResponse {
}

export interface MsgUnregisterRunner {
  creator: string;
  runnerAddress: string;
}

export interface MsgUnregisterRunnerResponse {
}

export interface MsgRunnerChallenge {
  creator: string;
  runnerAddress: string;
  clientPubkeys: string[];
}

export interface MsgRunnerChallengeResponse {
}

export interface MsgUnregisterGuard {
  creator: string;
}

export interface MsgUnregisterGuardResponse {
}

export interface MsgSelectRandomChallenger {
  creator: string;
}

export interface MsgSelectRandomChallengerResponse {
  randomChallenger: Challenger | undefined;
}

export interface MsgSelectRandomRunner {
  creator: string;
}

export interface MsgSelectRandomRunnerResponse {
  randomRunner: Runner | undefined;
}

export interface MsgUpdateGuard {
  creator: string;
  v2XAddr: string;
  v2XStake: string;
  v2XIp: string;
  v2NAddr: string;
  v2NStake: string;
  v2NIp: string;
  runnerAddr: string;
  runnerStake: string;
  runnerIp: string;
}

export interface MsgUpdateGuardResponse {
}

export interface MsgClaimMotusRewards {
  creator: string;
  amount: string;
}

export interface MsgClaimMotusRewardsResponse {
}

export interface MsgClaimRunnerRewards {
  creator: string;
  amount: string;
}

export interface MsgClaimRunnerRewardsResponse {
}

function createBaseMsgGenClient(): MsgGenClient {
  return { creator: "", pubkey: "" };
}

export const MsgGenClient = {
  encode(message: MsgGenClient, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.pubkey !== "") {
      writer.uint32(18).string(message.pubkey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGenClient {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGenClient();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgGenClient {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
    };
  },

  toJSON(message: MsgGenClient): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgGenClient>, I>>(object: I): MsgGenClient {
    const message = createBaseMsgGenClient();
    message.creator = object.creator ?? "";
    message.pubkey = object.pubkey ?? "";
    return message;
  },
};

function createBaseMsgGenClientResponse(): MsgGenClientResponse {
  return {};
}

export const MsgGenClientResponse = {
  encode(_: MsgGenClientResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGenClientResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGenClientResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgGenClientResponse {
    return {};
  },

  toJSON(_: MsgGenClientResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgGenClientResponse>, I>>(_: I): MsgGenClientResponse {
    const message = createBaseMsgGenClientResponse();
    return message;
  },
};

function createBaseMsgChallengeService(): MsgChallengeService {
  return { creator: "", clientPubkey: "", clientCommunicationMode: "", challengeResult: "" };
}

export const MsgChallengeService = {
  encode(message: MsgChallengeService, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.clientPubkey !== "") {
      writer.uint32(18).string(message.clientPubkey);
    }
    if (message.clientCommunicationMode !== "") {
      writer.uint32(26).string(message.clientCommunicationMode);
    }
    if (message.challengeResult !== "") {
      writer.uint32(34).string(message.challengeResult);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChallengeService {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChallengeService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.clientPubkey = reader.string();
          break;
        case 3:
          message.clientCommunicationMode = reader.string();
          break;
        case 4:
          message.challengeResult = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChallengeService {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      clientPubkey: isSet(object.clientPubkey) ? String(object.clientPubkey) : "",
      clientCommunicationMode: isSet(object.clientCommunicationMode) ? String(object.clientCommunicationMode) : "",
      challengeResult: isSet(object.challengeResult) ? String(object.challengeResult) : "",
    };
  },

  toJSON(message: MsgChallengeService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.clientPubkey !== undefined && (obj.clientPubkey = message.clientPubkey);
    message.clientCommunicationMode !== undefined && (obj.clientCommunicationMode = message.clientCommunicationMode);
    message.challengeResult !== undefined && (obj.challengeResult = message.challengeResult);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChallengeService>, I>>(object: I): MsgChallengeService {
    const message = createBaseMsgChallengeService();
    message.creator = object.creator ?? "";
    message.clientPubkey = object.clientPubkey ?? "";
    message.clientCommunicationMode = object.clientCommunicationMode ?? "";
    message.challengeResult = object.challengeResult ?? "";
    return message;
  },
};

function createBaseMsgChallengeServiceResponse(): MsgChallengeServiceResponse {
  return {};
}

export const MsgChallengeServiceResponse = {
  encode(_: MsgChallengeServiceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChallengeServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChallengeServiceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgChallengeServiceResponse {
    return {};
  },

  toJSON(_: MsgChallengeServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChallengeServiceResponse>, I>>(_: I): MsgChallengeServiceResponse {
    const message = createBaseMsgChallengeServiceResponse();
    return message;
  },
};

function createBaseMsgUnregisterClient(): MsgUnregisterClient {
  return { creator: "", pubkey: "" };
}

export const MsgUnregisterClient = {
  encode(message: MsgUnregisterClient, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.pubkey !== "") {
      writer.uint32(18).string(message.pubkey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterClient {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterClient();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterClient {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
    };
  },

  toJSON(message: MsgUnregisterClient): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterClient>, I>>(object: I): MsgUnregisterClient {
    const message = createBaseMsgUnregisterClient();
    message.creator = object.creator ?? "";
    message.pubkey = object.pubkey ?? "";
    return message;
  },
};

function createBaseMsgUnregisterClientResponse(): MsgUnregisterClientResponse {
  return {};
}

export const MsgUnregisterClientResponse = {
  encode(_: MsgUnregisterClientResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterClientResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterClientResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUnregisterClientResponse {
    return {};
  },

  toJSON(_: MsgUnregisterClientResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterClientResponse>, I>>(_: I): MsgUnregisterClientResponse {
    const message = createBaseMsgUnregisterClientResponse();
    return message;
  },
};

function createBaseMsgUnregisterChallenger(): MsgUnregisterChallenger {
  return { creator: "", challengerAddress: "" };
}

export const MsgUnregisterChallenger = {
  encode(message: MsgUnregisterChallenger, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.challengerAddress !== "") {
      writer.uint32(18).string(message.challengerAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterChallenger {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterChallenger();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.challengerAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterChallenger {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      challengerAddress: isSet(object.challengerAddress) ? String(object.challengerAddress) : "",
    };
  },

  toJSON(message: MsgUnregisterChallenger): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.challengerAddress !== undefined && (obj.challengerAddress = message.challengerAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterChallenger>, I>>(object: I): MsgUnregisterChallenger {
    const message = createBaseMsgUnregisterChallenger();
    message.creator = object.creator ?? "";
    message.challengerAddress = object.challengerAddress ?? "";
    return message;
  },
};

function createBaseMsgUnregisterChallengerResponse(): MsgUnregisterChallengerResponse {
  return {};
}

export const MsgUnregisterChallengerResponse = {
  encode(_: MsgUnregisterChallengerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterChallengerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterChallengerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUnregisterChallengerResponse {
    return {};
  },

  toJSON(_: MsgUnregisterChallengerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterChallengerResponse>, I>>(_: I): MsgUnregisterChallengerResponse {
    const message = createBaseMsgUnregisterChallengerResponse();
    return message;
  },
};

function createBaseMsgGenGuard(): MsgGenGuard {
  return {
    creator: "",
    guardPubKey: "",
    v2XAddr: "",
    v2XStake: "",
    v2XIp: "",
    v2NAddr: "",
    v2NStake: "",
    v2NIp: "",
    runnerAddr: "",
    runnerStake: "",
    runnerIp: "",
  };
}

export const MsgGenGuard = {
  encode(message: MsgGenGuard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.guardPubKey !== "") {
      writer.uint32(18).string(message.guardPubKey);
    }
    if (message.v2XAddr !== "") {
      writer.uint32(26).string(message.v2XAddr);
    }
    if (message.v2XStake !== "") {
      writer.uint32(34).string(message.v2XStake);
    }
    if (message.v2XIp !== "") {
      writer.uint32(42).string(message.v2XIp);
    }
    if (message.v2NAddr !== "") {
      writer.uint32(50).string(message.v2NAddr);
    }
    if (message.v2NStake !== "") {
      writer.uint32(58).string(message.v2NStake);
    }
    if (message.v2NIp !== "") {
      writer.uint32(66).string(message.v2NIp);
    }
    if (message.runnerAddr !== "") {
      writer.uint32(74).string(message.runnerAddr);
    }
    if (message.runnerStake !== "") {
      writer.uint32(82).string(message.runnerStake);
    }
    if (message.runnerIp !== "") {
      writer.uint32(90).string(message.runnerIp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGenGuard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGenGuard();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.guardPubKey = reader.string();
          break;
        case 3:
          message.v2XAddr = reader.string();
          break;
        case 4:
          message.v2XStake = reader.string();
          break;
        case 5:
          message.v2XIp = reader.string();
          break;
        case 6:
          message.v2NAddr = reader.string();
          break;
        case 7:
          message.v2NStake = reader.string();
          break;
        case 8:
          message.v2NIp = reader.string();
          break;
        case 9:
          message.runnerAddr = reader.string();
          break;
        case 10:
          message.runnerStake = reader.string();
          break;
        case 11:
          message.runnerIp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgGenGuard {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      guardPubKey: isSet(object.guardPubKey) ? String(object.guardPubKey) : "",
      v2XAddr: isSet(object.v2XAddr) ? String(object.v2XAddr) : "",
      v2XStake: isSet(object.v2XStake) ? String(object.v2XStake) : "",
      v2XIp: isSet(object.v2XIp) ? String(object.v2XIp) : "",
      v2NAddr: isSet(object.v2NAddr) ? String(object.v2NAddr) : "",
      v2NStake: isSet(object.v2NStake) ? String(object.v2NStake) : "",
      v2NIp: isSet(object.v2NIp) ? String(object.v2NIp) : "",
      runnerAddr: isSet(object.runnerAddr) ? String(object.runnerAddr) : "",
      runnerStake: isSet(object.runnerStake) ? String(object.runnerStake) : "",
      runnerIp: isSet(object.runnerIp) ? String(object.runnerIp) : "",
    };
  },

  toJSON(message: MsgGenGuard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.guardPubKey !== undefined && (obj.guardPubKey = message.guardPubKey);
    message.v2XAddr !== undefined && (obj.v2XAddr = message.v2XAddr);
    message.v2XStake !== undefined && (obj.v2XStake = message.v2XStake);
    message.v2XIp !== undefined && (obj.v2XIp = message.v2XIp);
    message.v2NAddr !== undefined && (obj.v2NAddr = message.v2NAddr);
    message.v2NStake !== undefined && (obj.v2NStake = message.v2NStake);
    message.v2NIp !== undefined && (obj.v2NIp = message.v2NIp);
    message.runnerAddr !== undefined && (obj.runnerAddr = message.runnerAddr);
    message.runnerStake !== undefined && (obj.runnerStake = message.runnerStake);
    message.runnerIp !== undefined && (obj.runnerIp = message.runnerIp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgGenGuard>, I>>(object: I): MsgGenGuard {
    const message = createBaseMsgGenGuard();
    message.creator = object.creator ?? "";
    message.guardPubKey = object.guardPubKey ?? "";
    message.v2XAddr = object.v2XAddr ?? "";
    message.v2XStake = object.v2XStake ?? "";
    message.v2XIp = object.v2XIp ?? "";
    message.v2NAddr = object.v2NAddr ?? "";
    message.v2NStake = object.v2NStake ?? "";
    message.v2NIp = object.v2NIp ?? "";
    message.runnerAddr = object.runnerAddr ?? "";
    message.runnerStake = object.runnerStake ?? "";
    message.runnerIp = object.runnerIp ?? "";
    return message;
  },
};

function createBaseMsgGenGuardResponse(): MsgGenGuardResponse {
  return {};
}

export const MsgGenGuardResponse = {
  encode(_: MsgGenGuardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGenGuardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGenGuardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgGenGuardResponse {
    return {};
  },

  toJSON(_: MsgGenGuardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgGenGuardResponse>, I>>(_: I): MsgGenGuardResponse {
    const message = createBaseMsgGenGuardResponse();
    return message;
  },
};

function createBaseMsgUnregisterRunner(): MsgUnregisterRunner {
  return { creator: "", runnerAddress: "" };
}

export const MsgUnregisterRunner = {
  encode(message: MsgUnregisterRunner, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.runnerAddress !== "") {
      writer.uint32(18).string(message.runnerAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterRunner {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterRunner();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.runnerAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterRunner {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      runnerAddress: isSet(object.runnerAddress) ? String(object.runnerAddress) : "",
    };
  },

  toJSON(message: MsgUnregisterRunner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.runnerAddress !== undefined && (obj.runnerAddress = message.runnerAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterRunner>, I>>(object: I): MsgUnregisterRunner {
    const message = createBaseMsgUnregisterRunner();
    message.creator = object.creator ?? "";
    message.runnerAddress = object.runnerAddress ?? "";
    return message;
  },
};

function createBaseMsgUnregisterRunnerResponse(): MsgUnregisterRunnerResponse {
  return {};
}

export const MsgUnregisterRunnerResponse = {
  encode(_: MsgUnregisterRunnerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterRunnerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterRunnerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUnregisterRunnerResponse {
    return {};
  },

  toJSON(_: MsgUnregisterRunnerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterRunnerResponse>, I>>(_: I): MsgUnregisterRunnerResponse {
    const message = createBaseMsgUnregisterRunnerResponse();
    return message;
  },
};

function createBaseMsgRunnerChallenge(): MsgRunnerChallenge {
  return { creator: "", runnerAddress: "", clientPubkeys: [] };
}

export const MsgRunnerChallenge = {
  encode(message: MsgRunnerChallenge, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.runnerAddress !== "") {
      writer.uint32(18).string(message.runnerAddress);
    }
    for (const v of message.clientPubkeys) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRunnerChallenge {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRunnerChallenge();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.runnerAddress = reader.string();
          break;
        case 3:
          message.clientPubkeys.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRunnerChallenge {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      runnerAddress: isSet(object.runnerAddress) ? String(object.runnerAddress) : "",
      clientPubkeys: Array.isArray(object?.clientPubkeys) ? object.clientPubkeys.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgRunnerChallenge): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.runnerAddress !== undefined && (obj.runnerAddress = message.runnerAddress);
    if (message.clientPubkeys) {
      obj.clientPubkeys = message.clientPubkeys.map((e) => e);
    } else {
      obj.clientPubkeys = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRunnerChallenge>, I>>(object: I): MsgRunnerChallenge {
    const message = createBaseMsgRunnerChallenge();
    message.creator = object.creator ?? "";
    message.runnerAddress = object.runnerAddress ?? "";
    message.clientPubkeys = object.clientPubkeys?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgRunnerChallengeResponse(): MsgRunnerChallengeResponse {
  return {};
}

export const MsgRunnerChallengeResponse = {
  encode(_: MsgRunnerChallengeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRunnerChallengeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRunnerChallengeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRunnerChallengeResponse {
    return {};
  },

  toJSON(_: MsgRunnerChallengeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRunnerChallengeResponse>, I>>(_: I): MsgRunnerChallengeResponse {
    const message = createBaseMsgRunnerChallengeResponse();
    return message;
  },
};

function createBaseMsgUnregisterGuard(): MsgUnregisterGuard {
  return { creator: "" };
}

export const MsgUnregisterGuard = {
  encode(message: MsgUnregisterGuard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterGuard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterGuard();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterGuard {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgUnregisterGuard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterGuard>, I>>(object: I): MsgUnregisterGuard {
    const message = createBaseMsgUnregisterGuard();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgUnregisterGuardResponse(): MsgUnregisterGuardResponse {
  return {};
}

export const MsgUnregisterGuardResponse = {
  encode(_: MsgUnregisterGuardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUnregisterGuardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUnregisterGuardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUnregisterGuardResponse {
    return {};
  },

  toJSON(_: MsgUnregisterGuardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUnregisterGuardResponse>, I>>(_: I): MsgUnregisterGuardResponse {
    const message = createBaseMsgUnregisterGuardResponse();
    return message;
  },
};

function createBaseMsgSelectRandomChallenger(): MsgSelectRandomChallenger {
  return { creator: "" };
}

export const MsgSelectRandomChallenger = {
  encode(message: MsgSelectRandomChallenger, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSelectRandomChallenger {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSelectRandomChallenger();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomChallenger {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgSelectRandomChallenger): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSelectRandomChallenger>, I>>(object: I): MsgSelectRandomChallenger {
    const message = createBaseMsgSelectRandomChallenger();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgSelectRandomChallengerResponse(): MsgSelectRandomChallengerResponse {
  return { randomChallenger: undefined };
}

export const MsgSelectRandomChallengerResponse = {
  encode(message: MsgSelectRandomChallengerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.randomChallenger !== undefined) {
      Challenger.encode(message.randomChallenger, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSelectRandomChallengerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSelectRandomChallengerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.randomChallenger = Challenger.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomChallengerResponse {
    return {
      randomChallenger: isSet(object.randomChallenger) ? Challenger.fromJSON(object.randomChallenger) : undefined,
    };
  },

  toJSON(message: MsgSelectRandomChallengerResponse): unknown {
    const obj: any = {};
    message.randomChallenger !== undefined
      && (obj.randomChallenger = message.randomChallenger ? Challenger.toJSON(message.randomChallenger) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSelectRandomChallengerResponse>, I>>(
    object: I,
  ): MsgSelectRandomChallengerResponse {
    const message = createBaseMsgSelectRandomChallengerResponse();
    message.randomChallenger = (object.randomChallenger !== undefined && object.randomChallenger !== null)
      ? Challenger.fromPartial(object.randomChallenger)
      : undefined;
    return message;
  },
};

function createBaseMsgSelectRandomRunner(): MsgSelectRandomRunner {
  return { creator: "" };
}

export const MsgSelectRandomRunner = {
  encode(message: MsgSelectRandomRunner, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSelectRandomRunner {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSelectRandomRunner();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomRunner {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgSelectRandomRunner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSelectRandomRunner>, I>>(object: I): MsgSelectRandomRunner {
    const message = createBaseMsgSelectRandomRunner();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgSelectRandomRunnerResponse(): MsgSelectRandomRunnerResponse {
  return { randomRunner: undefined };
}

export const MsgSelectRandomRunnerResponse = {
  encode(message: MsgSelectRandomRunnerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.randomRunner !== undefined) {
      Runner.encode(message.randomRunner, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSelectRandomRunnerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSelectRandomRunnerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.randomRunner = Runner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomRunnerResponse {
    return { randomRunner: isSet(object.randomRunner) ? Runner.fromJSON(object.randomRunner) : undefined };
  },

  toJSON(message: MsgSelectRandomRunnerResponse): unknown {
    const obj: any = {};
    message.randomRunner !== undefined
      && (obj.randomRunner = message.randomRunner ? Runner.toJSON(message.randomRunner) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSelectRandomRunnerResponse>, I>>(
    object: I,
  ): MsgSelectRandomRunnerResponse {
    const message = createBaseMsgSelectRandomRunnerResponse();
    message.randomRunner = (object.randomRunner !== undefined && object.randomRunner !== null)
      ? Runner.fromPartial(object.randomRunner)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateGuard(): MsgUpdateGuard {
  return {
    creator: "",
    v2XAddr: "",
    v2XStake: "",
    v2XIp: "",
    v2NAddr: "",
    v2NStake: "",
    v2NIp: "",
    runnerAddr: "",
    runnerStake: "",
    runnerIp: "",
  };
}

export const MsgUpdateGuard = {
  encode(message: MsgUpdateGuard, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.v2XAddr !== "") {
      writer.uint32(26).string(message.v2XAddr);
    }
    if (message.v2XStake !== "") {
      writer.uint32(34).string(message.v2XStake);
    }
    if (message.v2XIp !== "") {
      writer.uint32(42).string(message.v2XIp);
    }
    if (message.v2NAddr !== "") {
      writer.uint32(50).string(message.v2NAddr);
    }
    if (message.v2NStake !== "") {
      writer.uint32(58).string(message.v2NStake);
    }
    if (message.v2NIp !== "") {
      writer.uint32(66).string(message.v2NIp);
    }
    if (message.runnerAddr !== "") {
      writer.uint32(74).string(message.runnerAddr);
    }
    if (message.runnerStake !== "") {
      writer.uint32(82).string(message.runnerStake);
    }
    if (message.runnerIp !== "") {
      writer.uint32(90).string(message.runnerIp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateGuard {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateGuard();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 3:
          message.v2XAddr = reader.string();
          break;
        case 4:
          message.v2XStake = reader.string();
          break;
        case 5:
          message.v2XIp = reader.string();
          break;
        case 6:
          message.v2NAddr = reader.string();
          break;
        case 7:
          message.v2NStake = reader.string();
          break;
        case 8:
          message.v2NIp = reader.string();
          break;
        case 9:
          message.runnerAddr = reader.string();
          break;
        case 10:
          message.runnerStake = reader.string();
          break;
        case 11:
          message.runnerIp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateGuard {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      v2XAddr: isSet(object.v2XAddr) ? String(object.v2XAddr) : "",
      v2XStake: isSet(object.v2XStake) ? String(object.v2XStake) : "",
      v2XIp: isSet(object.v2XIp) ? String(object.v2XIp) : "",
      v2NAddr: isSet(object.v2NAddr) ? String(object.v2NAddr) : "",
      v2NStake: isSet(object.v2NStake) ? String(object.v2NStake) : "",
      v2NIp: isSet(object.v2NIp) ? String(object.v2NIp) : "",
      runnerAddr: isSet(object.runnerAddr) ? String(object.runnerAddr) : "",
      runnerStake: isSet(object.runnerStake) ? String(object.runnerStake) : "",
      runnerIp: isSet(object.runnerIp) ? String(object.runnerIp) : "",
    };
  },

  toJSON(message: MsgUpdateGuard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.v2XAddr !== undefined && (obj.v2XAddr = message.v2XAddr);
    message.v2XStake !== undefined && (obj.v2XStake = message.v2XStake);
    message.v2XIp !== undefined && (obj.v2XIp = message.v2XIp);
    message.v2NAddr !== undefined && (obj.v2NAddr = message.v2NAddr);
    message.v2NStake !== undefined && (obj.v2NStake = message.v2NStake);
    message.v2NIp !== undefined && (obj.v2NIp = message.v2NIp);
    message.runnerAddr !== undefined && (obj.runnerAddr = message.runnerAddr);
    message.runnerStake !== undefined && (obj.runnerStake = message.runnerStake);
    message.runnerIp !== undefined && (obj.runnerIp = message.runnerIp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateGuard>, I>>(object: I): MsgUpdateGuard {
    const message = createBaseMsgUpdateGuard();
    message.creator = object.creator ?? "";
    message.v2XAddr = object.v2XAddr ?? "";
    message.v2XStake = object.v2XStake ?? "";
    message.v2XIp = object.v2XIp ?? "";
    message.v2NAddr = object.v2NAddr ?? "";
    message.v2NStake = object.v2NStake ?? "";
    message.v2NIp = object.v2NIp ?? "";
    message.runnerAddr = object.runnerAddr ?? "";
    message.runnerStake = object.runnerStake ?? "";
    message.runnerIp = object.runnerIp ?? "";
    return message;
  },
};

function createBaseMsgUpdateGuardResponse(): MsgUpdateGuardResponse {
  return {};
}

export const MsgUpdateGuardResponse = {
  encode(_: MsgUpdateGuardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateGuardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateGuardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateGuardResponse {
    return {};
  },

  toJSON(_: MsgUpdateGuardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateGuardResponse>, I>>(_: I): MsgUpdateGuardResponse {
    const message = createBaseMsgUpdateGuardResponse();
    return message;
  },
};

function createBaseMsgClaimMotusRewards(): MsgClaimMotusRewards {
  return { creator: "", amount: "" };
}

export const MsgClaimMotusRewards = {
  encode(message: MsgClaimMotusRewards, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimMotusRewards {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimMotusRewards();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimMotusRewards {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: MsgClaimMotusRewards): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimMotusRewards>, I>>(object: I): MsgClaimMotusRewards {
    const message = createBaseMsgClaimMotusRewards();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseMsgClaimMotusRewardsResponse(): MsgClaimMotusRewardsResponse {
  return {};
}

export const MsgClaimMotusRewardsResponse = {
  encode(_: MsgClaimMotusRewardsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimMotusRewardsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimMotusRewardsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimMotusRewardsResponse {
    return {};
  },

  toJSON(_: MsgClaimMotusRewardsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimMotusRewardsResponse>, I>>(_: I): MsgClaimMotusRewardsResponse {
    const message = createBaseMsgClaimMotusRewardsResponse();
    return message;
  },
};

function createBaseMsgClaimRunnerRewards(): MsgClaimRunnerRewards {
  return { creator: "", amount: "" };
}

export const MsgClaimRunnerRewards = {
  encode(message: MsgClaimRunnerRewards, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimRunnerRewards {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimRunnerRewards();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimRunnerRewards {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: MsgClaimRunnerRewards): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimRunnerRewards>, I>>(object: I): MsgClaimRunnerRewards {
    const message = createBaseMsgClaimRunnerRewards();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseMsgClaimRunnerRewardsResponse(): MsgClaimRunnerRewardsResponse {
  return {};
}

export const MsgClaimRunnerRewardsResponse = {
  encode(_: MsgClaimRunnerRewardsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimRunnerRewardsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimRunnerRewardsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimRunnerRewardsResponse {
    return {};
  },

  toJSON(_: MsgClaimRunnerRewardsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimRunnerRewardsResponse>, I>>(_: I): MsgClaimRunnerRewardsResponse {
    const message = createBaseMsgClaimRunnerRewardsResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse>;
  ChallengeService(request: MsgChallengeService): Promise<MsgChallengeServiceResponse>;
  UnregisterClient(request: MsgUnregisterClient): Promise<MsgUnregisterClientResponse>;
  UnregisterChallenger(request: MsgUnregisterChallenger): Promise<MsgUnregisterChallengerResponse>;
  GenGuard(request: MsgGenGuard): Promise<MsgGenGuardResponse>;
  UnregisterRunner(request: MsgUnregisterRunner): Promise<MsgUnregisterRunnerResponse>;
  RunnerChallenge(request: MsgRunnerChallenge): Promise<MsgRunnerChallengeResponse>;
  UnregisterGuard(request: MsgUnregisterGuard): Promise<MsgUnregisterGuardResponse>;
  SelectRandomChallenger(request: MsgSelectRandomChallenger): Promise<MsgSelectRandomChallengerResponse>;
  SelectRandomRunner(request: MsgSelectRandomRunner): Promise<MsgSelectRandomRunnerResponse>;
  UpdateGuard(request: MsgUpdateGuard): Promise<MsgUpdateGuardResponse>;
  ClaimMotusRewards(request: MsgClaimMotusRewards): Promise<MsgClaimMotusRewardsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClaimRunnerRewards(request: MsgClaimRunnerRewards): Promise<MsgClaimRunnerRewardsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GenClient = this.GenClient.bind(this);
    this.ChallengeService = this.ChallengeService.bind(this);
    this.UnregisterClient = this.UnregisterClient.bind(this);
    this.UnregisterChallenger = this.UnregisterChallenger.bind(this);
    this.GenGuard = this.GenGuard.bind(this);
    this.UnregisterRunner = this.UnregisterRunner.bind(this);
    this.RunnerChallenge = this.RunnerChallenge.bind(this);
    this.UnregisterGuard = this.UnregisterGuard.bind(this);
    this.SelectRandomChallenger = this.SelectRandomChallenger.bind(this);
    this.SelectRandomRunner = this.SelectRandomRunner.bind(this);
    this.UpdateGuard = this.UpdateGuard.bind(this);
    this.ClaimMotusRewards = this.ClaimMotusRewards.bind(this);
    this.ClaimRunnerRewards = this.ClaimRunnerRewards.bind(this);
  }
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse> {
    const data = MsgGenClient.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "GenClient", data);
    return promise.then((data) => MsgGenClientResponse.decode(new _m0.Reader(data)));
  }

  ChallengeService(request: MsgChallengeService): Promise<MsgChallengeServiceResponse> {
    const data = MsgChallengeService.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "ChallengeService", data);
    return promise.then((data) => MsgChallengeServiceResponse.decode(new _m0.Reader(data)));
  }

  UnregisterClient(request: MsgUnregisterClient): Promise<MsgUnregisterClientResponse> {
    const data = MsgUnregisterClient.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "UnregisterClient", data);
    return promise.then((data) => MsgUnregisterClientResponse.decode(new _m0.Reader(data)));
  }

  UnregisterChallenger(request: MsgUnregisterChallenger): Promise<MsgUnregisterChallengerResponse> {
    const data = MsgUnregisterChallenger.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "UnregisterChallenger", data);
    return promise.then((data) => MsgUnregisterChallengerResponse.decode(new _m0.Reader(data)));
  }

  GenGuard(request: MsgGenGuard): Promise<MsgGenGuardResponse> {
    const data = MsgGenGuard.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "GenGuard", data);
    return promise.then((data) => MsgGenGuardResponse.decode(new _m0.Reader(data)));
  }

  UnregisterRunner(request: MsgUnregisterRunner): Promise<MsgUnregisterRunnerResponse> {
    const data = MsgUnregisterRunner.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "UnregisterRunner", data);
    return promise.then((data) => MsgUnregisterRunnerResponse.decode(new _m0.Reader(data)));
  }

  RunnerChallenge(request: MsgRunnerChallenge): Promise<MsgRunnerChallengeResponse> {
    const data = MsgRunnerChallenge.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "RunnerChallenge", data);
    return promise.then((data) => MsgRunnerChallengeResponse.decode(new _m0.Reader(data)));
  }

  UnregisterGuard(request: MsgUnregisterGuard): Promise<MsgUnregisterGuardResponse> {
    const data = MsgUnregisterGuard.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "UnregisterGuard", data);
    return promise.then((data) => MsgUnregisterGuardResponse.decode(new _m0.Reader(data)));
  }

  SelectRandomChallenger(request: MsgSelectRandomChallenger): Promise<MsgSelectRandomChallengerResponse> {
    const data = MsgSelectRandomChallenger.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "SelectRandomChallenger", data);
    return promise.then((data) => MsgSelectRandomChallengerResponse.decode(new _m0.Reader(data)));
  }

  SelectRandomRunner(request: MsgSelectRandomRunner): Promise<MsgSelectRandomRunnerResponse> {
    const data = MsgSelectRandomRunner.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "SelectRandomRunner", data);
    return promise.then((data) => MsgSelectRandomRunnerResponse.decode(new _m0.Reader(data)));
  }

  UpdateGuard(request: MsgUpdateGuard): Promise<MsgUpdateGuardResponse> {
    const data = MsgUpdateGuard.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "UpdateGuard", data);
    return promise.then((data) => MsgUpdateGuardResponse.decode(new _m0.Reader(data)));
  }

  ClaimMotusRewards(request: MsgClaimMotusRewards): Promise<MsgClaimMotusRewardsResponse> {
    const data = MsgClaimMotusRewards.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "ClaimMotusRewards", data);
    return promise.then((data) => MsgClaimMotusRewardsResponse.decode(new _m0.Reader(data)));
  }

  ClaimRunnerRewards(request: MsgClaimRunnerRewards): Promise<MsgClaimRunnerRewardsResponse> {
    const data = MsgClaimRunnerRewards.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "ClaimRunnerRewards", data);
    return promise.then((data) => MsgClaimRunnerRewardsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
