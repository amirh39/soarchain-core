/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "soarchain.poa";

export interface MsgGenClient {
  creator: string;
  address: string;
  fee: string;
}

export interface MsgGenClientResponse {}

export interface MsgGenChallenger {
  creator: string;
  stakeAmount: string;
}

export interface MsgGenChallengerResponse {}

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
  fee: string;
}

export interface MsgUnregisterChallengerResponse {}

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

const baseMsgGenChallenger: object = { creator: "", stakeAmount: "" };

export const MsgGenChallenger = {
  encode(message: MsgGenChallenger, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.stakeAmount !== "") {
      writer.uint32(18).string(message.stakeAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgGenChallenger {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgGenChallenger } as MsgGenChallenger;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.stakeAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgGenChallenger {
    const message = { ...baseMsgGenChallenger } as MsgGenChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = String(object.stakeAmount);
    } else {
      message.stakeAmount = "";
    }
    return message;
  },

  toJSON(message: MsgGenChallenger): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgGenChallenger>): MsgGenChallenger {
    const message = { ...baseMsgGenChallenger } as MsgGenChallenger;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = object.stakeAmount;
    } else {
      message.stakeAmount = "";
    }
    return message;
  },
};

const baseMsgGenChallengerResponse: object = {};

export const MsgGenChallengerResponse = {
  encode(
    _: MsgGenChallengerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgGenChallengerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgGenChallengerResponse,
    } as MsgGenChallengerResponse;
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

  fromJSON(_: any): MsgGenChallengerResponse {
    const message = {
      ...baseMsgGenChallengerResponse,
    } as MsgGenChallengerResponse;
    return message;
  },

  toJSON(_: MsgGenChallengerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgGenChallengerResponse>
  ): MsgGenChallengerResponse {
    const message = {
      ...baseMsgGenChallengerResponse,
    } as MsgGenChallengerResponse;
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

const baseMsgUnregisterChallenger: object = { creator: "", fee: "" };

export const MsgUnregisterChallenger = {
  encode(
    message: MsgUnregisterChallenger,
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

/** Msg defines the Msg service. */
export interface Msg {
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse>;
  GenChallenger(request: MsgGenChallenger): Promise<MsgGenChallengerResponse>;
  ChallengeService(
    request: MsgChallengeService
  ): Promise<MsgChallengeServiceResponse>;
  UnregisterClient(
    request: MsgUnregisterClient
  ): Promise<MsgUnregisterClientResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UnregisterChallenger(
    request: MsgUnregisterChallenger
  ): Promise<MsgUnregisterChallengerResponse>;
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

  GenChallenger(request: MsgGenChallenger): Promise<MsgGenChallengerResponse> {
    const data = MsgGenChallenger.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Msg",
      "GenChallenger",
      data
    );
    return promise.then((data) =>
      MsgGenChallengerResponse.decode(new Reader(data))
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
