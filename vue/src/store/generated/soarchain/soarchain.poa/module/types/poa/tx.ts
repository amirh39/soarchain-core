/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface MsgGenClient {
  creator: string;
  address: string;
  fee: string;
}

export interface MsgGenClientResponse {}

export interface MsgChallengeService {
  creator: string;
  challengeeAddress: string;
  challengeResult: string;
}

export interface MsgChallengeServiceResponse {}

export interface MsgUnregisterClient {
  creator: string;
  address: string;
  fee: string;
}

export interface MsgUnregisterClientResponse {}

export interface MsgUnregisterChallenger {
  creator: string;
  challengerAddress: string;
  fee: string;
}

export interface MsgUnregisterChallengerResponse {}

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

export interface MsgGenGuardResponse {}

export interface MsgUnregisterRunner {
  creator: string;
  runnerAddress: string;
  fee: string;
}

export interface MsgUnregisterRunnerResponse {}

export interface MsgRunnerChallenge {
  creator: string;
  runnerAddress: string;
  challengeResult: string;
}

export interface MsgRunnerChallengeResponse {}

export interface MsgUnregisterGuard {
  creator: string;
  fee: string;
}

export interface MsgUnregisterGuardResponse {}

export interface MsgSelectRandomChallenger {
  creator: string;
  multiplier: string;
}

export interface MsgSelectRandomChallengerResponse {
  randomChallenger: string;
}

const baseMsgGenClient: object = { creator: "", address: "", fee: "" };

export const MsgGenClient = {
  encode(message: MsgGenClient, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgGenClient {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgGenClient } as MsgGenClient;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgGenClient {
    const message = { ...baseMsgGenClient } as MsgGenClient;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgGenClient): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgGenClient>): MsgGenClient {
    const message = { ...baseMsgGenClient } as MsgGenClient;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgGenClientResponse: object = {};

export const MsgGenClientResponse = {
  encode(_: MsgGenClientResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgGenClientResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgGenClientResponse } as MsgGenClientResponse;
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
    const message = { ...baseMsgGenClientResponse } as MsgGenClientResponse;
    return message;
  },

  toJSON(_: MsgGenClientResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgGenClientResponse>): MsgGenClientResponse {
    const message = { ...baseMsgGenClientResponse } as MsgGenClientResponse;
    return message;
  },
};

const baseMsgChallengeService: object = {
  creator: "",
  challengeeAddress: "",
  challengeResult: "",
};

export const MsgChallengeService = {
  encode(
    message: MsgChallengeService,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.challengeeAddress !== "") {
      writer.uint32(18).string(message.challengeeAddress);
    }
    if (message.challengeResult !== "") {
      writer.uint32(26).string(message.challengeResult);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChallengeService {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChallengeService } as MsgChallengeService;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.challengeeAddress = reader.string();
          break;
        case 3:
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
    const message = { ...baseMsgChallengeService } as MsgChallengeService;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.challengeeAddress !== undefined &&
      object.challengeeAddress !== null
    ) {
      message.challengeeAddress = String(object.challengeeAddress);
    } else {
      message.challengeeAddress = "";
    }
    if (
      object.challengeResult !== undefined &&
      object.challengeResult !== null
    ) {
      message.challengeResult = String(object.challengeResult);
    } else {
      message.challengeResult = "";
    }
    return message;
  },

  toJSON(message: MsgChallengeService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.challengeeAddress !== undefined &&
      (obj.challengeeAddress = message.challengeeAddress);
    message.challengeResult !== undefined &&
      (obj.challengeResult = message.challengeResult);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgChallengeService>): MsgChallengeService {
    const message = { ...baseMsgChallengeService } as MsgChallengeService;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.challengeeAddress !== undefined &&
      object.challengeeAddress !== null
    ) {
      message.challengeeAddress = object.challengeeAddress;
    } else {
      message.challengeeAddress = "";
    }
    if (
      object.challengeResult !== undefined &&
      object.challengeResult !== null
    ) {
      message.challengeResult = object.challengeResult;
    } else {
      message.challengeResult = "";
    }
    return message;
  },
};

const baseMsgChallengeServiceResponse: object = {};

export const MsgChallengeServiceResponse = {
  encode(
    _: MsgChallengeServiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgChallengeServiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgChallengeServiceResponse,
    } as MsgChallengeServiceResponse;
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
    const message = {
      ...baseMsgChallengeServiceResponse,
    } as MsgChallengeServiceResponse;
    return message;
  },

  toJSON(_: MsgChallengeServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgChallengeServiceResponse>
  ): MsgChallengeServiceResponse {
    const message = {
      ...baseMsgChallengeServiceResponse,
    } as MsgChallengeServiceResponse;
    return message;
  },
};

const baseMsgUnregisterClient: object = { creator: "", address: "", fee: "" };

export const MsgUnregisterClient = {
  encode(
    message: MsgUnregisterClient,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUnregisterClient {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUnregisterClient } as MsgUnregisterClient;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterClient {
    const message = { ...baseMsgUnregisterClient } as MsgUnregisterClient;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgUnregisterClient): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUnregisterClient>): MsgUnregisterClient {
    const message = { ...baseMsgUnregisterClient } as MsgUnregisterClient;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgUnregisterClientResponse: object = {};

export const MsgUnregisterClientResponse = {
  encode(
    _: MsgUnregisterClientResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUnregisterClientResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnregisterClientResponse,
    } as MsgUnregisterClientResponse;
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
    const message = {
      ...baseMsgUnregisterClientResponse,
    } as MsgUnregisterClientResponse;
    return message;
  },

  toJSON(_: MsgUnregisterClientResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUnregisterClientResponse>
  ): MsgUnregisterClientResponse {
    const message = {
      ...baseMsgUnregisterClientResponse,
    } as MsgUnregisterClientResponse;
    return message;
  },
};

const baseMsgUnregisterChallenger: object = {
  creator: "",
  challengerAddress: "",
  fee: "",
};

export const MsgUnregisterChallenger = {
  encode(
    message: MsgUnregisterChallenger,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.challengerAddress !== "") {
      writer.uint32(18).string(message.challengerAddress);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUnregisterChallenger {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnregisterChallenger,
    } as MsgUnregisterChallenger;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.challengerAddress = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterChallenger {
    const message = {
      ...baseMsgUnregisterChallenger,
    } as MsgUnregisterChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.challengerAddress !== undefined &&
      object.challengerAddress !== null
    ) {
      message.challengerAddress = String(object.challengerAddress);
    } else {
      message.challengerAddress = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgUnregisterChallenger): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.challengerAddress !== undefined &&
      (obj.challengerAddress = message.challengerAddress);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUnregisterChallenger>
  ): MsgUnregisterChallenger {
    const message = {
      ...baseMsgUnregisterChallenger,
    } as MsgUnregisterChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.challengerAddress !== undefined &&
      object.challengerAddress !== null
    ) {
      message.challengerAddress = object.challengerAddress;
    } else {
      message.challengerAddress = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgUnregisterChallengerResponse: object = {};

export const MsgUnregisterChallengerResponse = {
  encode(
    _: MsgUnregisterChallengerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUnregisterChallengerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnregisterChallengerResponse,
    } as MsgUnregisterChallengerResponse;
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
    const message = {
      ...baseMsgUnregisterChallengerResponse,
    } as MsgUnregisterChallengerResponse;
    return message;
  },

  toJSON(_: MsgUnregisterChallengerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUnregisterChallengerResponse>
  ): MsgUnregisterChallengerResponse {
    const message = {
      ...baseMsgUnregisterChallengerResponse,
    } as MsgUnregisterChallengerResponse;
    return message;
  },
};

const baseMsgGenGuard: object = {
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

export const MsgGenGuard = {
  encode(message: MsgGenGuard, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgGenGuard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgGenGuard } as MsgGenGuard;
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
    const message = { ...baseMsgGenGuard } as MsgGenGuard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.guardPubKey !== undefined && object.guardPubKey !== null) {
      message.guardPubKey = String(object.guardPubKey);
    } else {
      message.guardPubKey = "";
    }
    if (object.v2XAddr !== undefined && object.v2XAddr !== null) {
      message.v2XAddr = String(object.v2XAddr);
    } else {
      message.v2XAddr = "";
    }
    if (object.v2XStake !== undefined && object.v2XStake !== null) {
      message.v2XStake = String(object.v2XStake);
    } else {
      message.v2XStake = "";
    }
    if (object.v2XIp !== undefined && object.v2XIp !== null) {
      message.v2XIp = String(object.v2XIp);
    } else {
      message.v2XIp = "";
    }
    if (object.v2NAddr !== undefined && object.v2NAddr !== null) {
      message.v2NAddr = String(object.v2NAddr);
    } else {
      message.v2NAddr = "";
    }
    if (object.v2NStake !== undefined && object.v2NStake !== null) {
      message.v2NStake = String(object.v2NStake);
    } else {
      message.v2NStake = "";
    }
    if (object.v2NIp !== undefined && object.v2NIp !== null) {
      message.v2NIp = String(object.v2NIp);
    } else {
      message.v2NIp = "";
    }
    if (object.runnerAddr !== undefined && object.runnerAddr !== null) {
      message.runnerAddr = String(object.runnerAddr);
    } else {
      message.runnerAddr = "";
    }
    if (object.runnerStake !== undefined && object.runnerStake !== null) {
      message.runnerStake = String(object.runnerStake);
    } else {
      message.runnerStake = "";
    }
    if (object.runnerIp !== undefined && object.runnerIp !== null) {
      message.runnerIp = String(object.runnerIp);
    } else {
      message.runnerIp = "";
    }
    return message;
  },

  toJSON(message: MsgGenGuard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.guardPubKey !== undefined &&
      (obj.guardPubKey = message.guardPubKey);
    message.v2XAddr !== undefined && (obj.v2XAddr = message.v2XAddr);
    message.v2XStake !== undefined && (obj.v2XStake = message.v2XStake);
    message.v2XIp !== undefined && (obj.v2XIp = message.v2XIp);
    message.v2NAddr !== undefined && (obj.v2NAddr = message.v2NAddr);
    message.v2NStake !== undefined && (obj.v2NStake = message.v2NStake);
    message.v2NIp !== undefined && (obj.v2NIp = message.v2NIp);
    message.runnerAddr !== undefined && (obj.runnerAddr = message.runnerAddr);
    message.runnerStake !== undefined &&
      (obj.runnerStake = message.runnerStake);
    message.runnerIp !== undefined && (obj.runnerIp = message.runnerIp);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgGenGuard>): MsgGenGuard {
    const message = { ...baseMsgGenGuard } as MsgGenGuard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.guardPubKey !== undefined && object.guardPubKey !== null) {
      message.guardPubKey = object.guardPubKey;
    } else {
      message.guardPubKey = "";
    }
    if (object.v2XAddr !== undefined && object.v2XAddr !== null) {
      message.v2XAddr = object.v2XAddr;
    } else {
      message.v2XAddr = "";
    }
    if (object.v2XStake !== undefined && object.v2XStake !== null) {
      message.v2XStake = object.v2XStake;
    } else {
      message.v2XStake = "";
    }
    if (object.v2XIp !== undefined && object.v2XIp !== null) {
      message.v2XIp = object.v2XIp;
    } else {
      message.v2XIp = "";
    }
    if (object.v2NAddr !== undefined && object.v2NAddr !== null) {
      message.v2NAddr = object.v2NAddr;
    } else {
      message.v2NAddr = "";
    }
    if (object.v2NStake !== undefined && object.v2NStake !== null) {
      message.v2NStake = object.v2NStake;
    } else {
      message.v2NStake = "";
    }
    if (object.v2NIp !== undefined && object.v2NIp !== null) {
      message.v2NIp = object.v2NIp;
    } else {
      message.v2NIp = "";
    }
    if (object.runnerAddr !== undefined && object.runnerAddr !== null) {
      message.runnerAddr = object.runnerAddr;
    } else {
      message.runnerAddr = "";
    }
    if (object.runnerStake !== undefined && object.runnerStake !== null) {
      message.runnerStake = object.runnerStake;
    } else {
      message.runnerStake = "";
    }
    if (object.runnerIp !== undefined && object.runnerIp !== null) {
      message.runnerIp = object.runnerIp;
    } else {
      message.runnerIp = "";
    }
    return message;
  },
};

const baseMsgGenGuardResponse: object = {};

export const MsgGenGuardResponse = {
  encode(_: MsgGenGuardResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgGenGuardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgGenGuardResponse } as MsgGenGuardResponse;
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
    const message = { ...baseMsgGenGuardResponse } as MsgGenGuardResponse;
    return message;
  },

  toJSON(_: MsgGenGuardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgGenGuardResponse>): MsgGenGuardResponse {
    const message = { ...baseMsgGenGuardResponse } as MsgGenGuardResponse;
    return message;
  },
};

const baseMsgUnregisterRunner: object = {
  creator: "",
  runnerAddress: "",
  fee: "",
};

export const MsgUnregisterRunner = {
  encode(
    message: MsgUnregisterRunner,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.runnerAddress !== "") {
      writer.uint32(18).string(message.runnerAddress);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUnregisterRunner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUnregisterRunner } as MsgUnregisterRunner;
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
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterRunner {
    const message = { ...baseMsgUnregisterRunner } as MsgUnregisterRunner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.runnerAddress !== undefined && object.runnerAddress !== null) {
      message.runnerAddress = String(object.runnerAddress);
    } else {
      message.runnerAddress = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgUnregisterRunner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.runnerAddress !== undefined &&
      (obj.runnerAddress = message.runnerAddress);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUnregisterRunner>): MsgUnregisterRunner {
    const message = { ...baseMsgUnregisterRunner } as MsgUnregisterRunner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.runnerAddress !== undefined && object.runnerAddress !== null) {
      message.runnerAddress = object.runnerAddress;
    } else {
      message.runnerAddress = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgUnregisterRunnerResponse: object = {};

export const MsgUnregisterRunnerResponse = {
  encode(
    _: MsgUnregisterRunnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUnregisterRunnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnregisterRunnerResponse,
    } as MsgUnregisterRunnerResponse;
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
    const message = {
      ...baseMsgUnregisterRunnerResponse,
    } as MsgUnregisterRunnerResponse;
    return message;
  },

  toJSON(_: MsgUnregisterRunnerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUnregisterRunnerResponse>
  ): MsgUnregisterRunnerResponse {
    const message = {
      ...baseMsgUnregisterRunnerResponse,
    } as MsgUnregisterRunnerResponse;
    return message;
  },
};

const baseMsgRunnerChallenge: object = {
  creator: "",
  runnerAddress: "",
  challengeResult: "",
};

export const MsgRunnerChallenge = {
  encode(
    message: MsgRunnerChallenge,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.runnerAddress !== "") {
      writer.uint32(18).string(message.runnerAddress);
    }
    if (message.challengeResult !== "") {
      writer.uint32(26).string(message.challengeResult);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRunnerChallenge {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRunnerChallenge } as MsgRunnerChallenge;
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
          message.challengeResult = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRunnerChallenge {
    const message = { ...baseMsgRunnerChallenge } as MsgRunnerChallenge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.runnerAddress !== undefined && object.runnerAddress !== null) {
      message.runnerAddress = String(object.runnerAddress);
    } else {
      message.runnerAddress = "";
    }
    if (
      object.challengeResult !== undefined &&
      object.challengeResult !== null
    ) {
      message.challengeResult = String(object.challengeResult);
    } else {
      message.challengeResult = "";
    }
    return message;
  },

  toJSON(message: MsgRunnerChallenge): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.runnerAddress !== undefined &&
      (obj.runnerAddress = message.runnerAddress);
    message.challengeResult !== undefined &&
      (obj.challengeResult = message.challengeResult);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRunnerChallenge>): MsgRunnerChallenge {
    const message = { ...baseMsgRunnerChallenge } as MsgRunnerChallenge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.runnerAddress !== undefined && object.runnerAddress !== null) {
      message.runnerAddress = object.runnerAddress;
    } else {
      message.runnerAddress = "";
    }
    if (
      object.challengeResult !== undefined &&
      object.challengeResult !== null
    ) {
      message.challengeResult = object.challengeResult;
    } else {
      message.challengeResult = "";
    }
    return message;
  },
};

const baseMsgRunnerChallengeResponse: object = {};

export const MsgRunnerChallengeResponse = {
  encode(
    _: MsgRunnerChallengeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRunnerChallengeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRunnerChallengeResponse,
    } as MsgRunnerChallengeResponse;
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
    const message = {
      ...baseMsgRunnerChallengeResponse,
    } as MsgRunnerChallengeResponse;
    return message;
  },

  toJSON(_: MsgRunnerChallengeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRunnerChallengeResponse>
  ): MsgRunnerChallengeResponse {
    const message = {
      ...baseMsgRunnerChallengeResponse,
    } as MsgRunnerChallengeResponse;
    return message;
  },
};

const baseMsgUnregisterGuard: object = { creator: "", fee: "" };

export const MsgUnregisterGuard = {
  encode(
    message: MsgUnregisterGuard,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUnregisterGuard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUnregisterGuard } as MsgUnregisterGuard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnregisterGuard {
    const message = { ...baseMsgUnregisterGuard } as MsgUnregisterGuard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    return message;
  },

  toJSON(message: MsgUnregisterGuard): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fee !== undefined && (obj.fee = message.fee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUnregisterGuard>): MsgUnregisterGuard {
    const message = { ...baseMsgUnregisterGuard } as MsgUnregisterGuard;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    return message;
  },
};

const baseMsgUnregisterGuardResponse: object = {};

export const MsgUnregisterGuardResponse = {
  encode(
    _: MsgUnregisterGuardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUnregisterGuardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnregisterGuardResponse,
    } as MsgUnregisterGuardResponse;
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
    const message = {
      ...baseMsgUnregisterGuardResponse,
    } as MsgUnregisterGuardResponse;
    return message;
  },

  toJSON(_: MsgUnregisterGuardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUnregisterGuardResponse>
  ): MsgUnregisterGuardResponse {
    const message = {
      ...baseMsgUnregisterGuardResponse,
    } as MsgUnregisterGuardResponse;
    return message;
  },
};

const baseMsgSelectRandomChallenger: object = { creator: "", multiplier: "" };

export const MsgSelectRandomChallenger = {
  encode(
    message: MsgSelectRandomChallenger,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.multiplier !== "") {
      writer.uint32(18).string(message.multiplier);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSelectRandomChallenger {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSelectRandomChallenger,
    } as MsgSelectRandomChallenger;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.multiplier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomChallenger {
    const message = {
      ...baseMsgSelectRandomChallenger,
    } as MsgSelectRandomChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = String(object.multiplier);
    } else {
      message.multiplier = "";
    }
    return message;
  },

  toJSON(message: MsgSelectRandomChallenger): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.multiplier !== undefined && (obj.multiplier = message.multiplier);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSelectRandomChallenger>
  ): MsgSelectRandomChallenger {
    const message = {
      ...baseMsgSelectRandomChallenger,
    } as MsgSelectRandomChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = object.multiplier;
    } else {
      message.multiplier = "";
    }
    return message;
  },
};

const baseMsgSelectRandomChallengerResponse: object = { randomChallenger: "" };

export const MsgSelectRandomChallengerResponse = {
  encode(
    message: MsgSelectRandomChallengerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.randomChallenger !== "") {
      writer.uint32(10).string(message.randomChallenger);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSelectRandomChallengerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSelectRandomChallengerResponse,
    } as MsgSelectRandomChallengerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.randomChallenger = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSelectRandomChallengerResponse {
    const message = {
      ...baseMsgSelectRandomChallengerResponse,
    } as MsgSelectRandomChallengerResponse;
    if (
      object.randomChallenger !== undefined &&
      object.randomChallenger !== null
    ) {
      message.randomChallenger = String(object.randomChallenger);
    } else {
      message.randomChallenger = "";
    }
    return message;
  },

  toJSON(message: MsgSelectRandomChallengerResponse): unknown {
    const obj: any = {};
    message.randomChallenger !== undefined &&
      (obj.randomChallenger = message.randomChallenger);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSelectRandomChallengerResponse>
  ): MsgSelectRandomChallengerResponse {
    const message = {
      ...baseMsgSelectRandomChallengerResponse,
    } as MsgSelectRandomChallengerResponse;
    if (
      object.randomChallenger !== undefined &&
      object.randomChallenger !== null
    ) {
      message.randomChallenger = object.randomChallenger;
    } else {
      message.randomChallenger = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse>;
  ChallengeService(
    request: MsgChallengeService
  ): Promise<MsgChallengeServiceResponse>;
  UnregisterClient(
    request: MsgUnregisterClient
  ): Promise<MsgUnregisterClientResponse>;
  UnregisterChallenger(
    request: MsgUnregisterChallenger
  ): Promise<MsgUnregisterChallengerResponse>;
  GenGuard(request: MsgGenGuard): Promise<MsgGenGuardResponse>;
  UnregisterRunner(
    request: MsgUnregisterRunner
  ): Promise<MsgUnregisterRunnerResponse>;
  RunnerChallenge(
    request: MsgRunnerChallenge
  ): Promise<MsgRunnerChallengeResponse>;
  UnregisterGuard(
    request: MsgUnregisterGuard
  ): Promise<MsgUnregisterGuardResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SelectRandomChallenger(
    request: MsgSelectRandomChallenger
  ): Promise<MsgSelectRandomChallengerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse> {
    const data = MsgGenClient.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "GenClient", data);
    return promise.then((data) =>
      MsgGenClientResponse.decode(new Reader(data))
    );
  }

  ChallengeService(
    request: MsgChallengeService
  ): Promise<MsgChallengeServiceResponse> {
    const data = MsgChallengeService.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "ChallengeService",
      data
    );
    return promise.then((data) =>
      MsgChallengeServiceResponse.decode(new Reader(data))
    );
  }

  UnregisterClient(
    request: MsgUnregisterClient
  ): Promise<MsgUnregisterClientResponse> {
    const data = MsgUnregisterClient.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "UnregisterClient",
      data
    );
    return promise.then((data) =>
      MsgUnregisterClientResponse.decode(new Reader(data))
    );
  }

  UnregisterChallenger(
    request: MsgUnregisterChallenger
  ): Promise<MsgUnregisterChallengerResponse> {
    const data = MsgUnregisterChallenger.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "UnregisterChallenger",
      data
    );
    return promise.then((data) =>
      MsgUnregisterChallengerResponse.decode(new Reader(data))
    );
  }

  GenGuard(request: MsgGenGuard): Promise<MsgGenGuardResponse> {
    const data = MsgGenGuard.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Msg", "GenGuard", data);
    return promise.then((data) => MsgGenGuardResponse.decode(new Reader(data)));
  }

  UnregisterRunner(
    request: MsgUnregisterRunner
  ): Promise<MsgUnregisterRunnerResponse> {
    const data = MsgUnregisterRunner.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "UnregisterRunner",
      data
    );
    return promise.then((data) =>
      MsgUnregisterRunnerResponse.decode(new Reader(data))
    );
  }

  RunnerChallenge(
    request: MsgRunnerChallenge
  ): Promise<MsgRunnerChallengeResponse> {
    const data = MsgRunnerChallenge.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "RunnerChallenge",
      data
    );
    return promise.then((data) =>
      MsgRunnerChallengeResponse.decode(new Reader(data))
    );
  }

  UnregisterGuard(
    request: MsgUnregisterGuard
  ): Promise<MsgUnregisterGuardResponse> {
    const data = MsgUnregisterGuard.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "UnregisterGuard",
      data
    );
    return promise.then((data) =>
      MsgUnregisterGuardResponse.decode(new Reader(data))
    );
  }

  SelectRandomChallenger(
    request: MsgSelectRandomChallenger
  ): Promise<MsgSelectRandomChallengerResponse> {
    const data = MsgSelectRandomChallenger.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "SelectRandomChallenger",
      data
    );
    return promise.then((data) =>
      MsgSelectRandomChallengerResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
