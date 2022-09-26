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

/** Msg defines the Msg service. */
export interface Msg {
  GenClient(request: MsgGenClient): Promise<MsgGenClientResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  GenChallenger(request: MsgGenChallenger): Promise<MsgGenChallengerResponse>;
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
