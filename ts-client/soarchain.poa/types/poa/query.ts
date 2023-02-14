/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Challenger } from "./challenger";
import { Client } from "./client";
import { Guard } from "./guard";
import { Params } from "./params";
import { Runner } from "./runner";
import { VrfData } from "./vrf_data";
import { VrfUser } from "./vrf_user";

export const protobufPackage = "soarchain.poa";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetClientRequest {
  index: string;
}

export interface QueryGetClientResponse {
  client: Client | undefined;
}

export interface QueryAllClientRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClientResponse {
  client: Client[];
  pagination: PageResponse | undefined;
}

export interface QueryGetChallengerRequest {
  index: string;
}

export interface QueryGetChallengerResponse {
  challenger: Challenger | undefined;
}

export interface QueryAllChallengerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllChallengerResponse {
  challenger: Challenger[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRunnerRequest {
  index: string;
}

export interface QueryGetRunnerResponse {
  runner: Runner | undefined;
}

export interface QueryAllRunnerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRunnerResponse {
  runner: Runner[];
  pagination: PageResponse | undefined;
}

export interface QueryGetGuardRequest {
  index: string;
}

export interface QueryGetGuardResponse {
  guard: Guard | undefined;
}

export interface QueryAllGuardRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllGuardResponse {
  guard: Guard[];
  pagination: PageResponse | undefined;
}

export interface QueryGetClientByAddressRequest {
  address: string;
}

export interface QueryGetClientByAddressResponse {
  client: Client | undefined;
}

export interface QueryGetChallengerByAddressRequest {
  address: string;
}

export interface QueryGetChallengerByAddressResponse {
  challenger: Challenger | undefined;
}

export interface QueryGetVrfDataRequest {
  index: string;
}

export interface QueryGetVrfDataResponse {
  vrfData: VrfData | undefined;
}

export interface QueryAllVrfDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVrfDataResponse {
  vrfData: VrfData[];
  pagination: PageResponse | undefined;
}

export interface QueryGetVrfUserRequest {
  index: string;
}

export interface QueryGetVrfUserResponse {
  vrfUser: VrfUser | undefined;
}

export interface QueryAllVrfUserRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVrfUserResponse {
  vrfUser: VrfUser[];
  pagination: PageResponse | undefined;
}

export interface QueryVerifyRandomNumberRequest {
  pubkey: string;
  message: string;
  vrv: string;
  proof: string;
}

export interface QueryVerifyRandomNumberResponse {
  result: boolean;
}

export interface QueryIsChallengeableRequest {
  clientAddr: string;
}

export interface QueryIsChallengeableResponse {
  resultBool: string;
  challengeabilityScore: string;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetClientRequest(): QueryGetClientRequest {
  return { index: "" };
}

export const QueryGetClientRequest = {
  encode(message: QueryGetClientRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetClientRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetClientRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetClientRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetClientRequest>, I>>(object: I): QueryGetClientRequest {
    const message = createBaseQueryGetClientRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetClientResponse(): QueryGetClientResponse {
  return { client: undefined };
}

export const QueryGetClientResponse = {
  encode(message: QueryGetClientResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.client !== undefined) {
      Client.encode(message.client, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetClientResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetClientResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.client = Client.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientResponse {
    return { client: isSet(object.client) ? Client.fromJSON(object.client) : undefined };
  },

  toJSON(message: QueryGetClientResponse): unknown {
    const obj: any = {};
    message.client !== undefined && (obj.client = message.client ? Client.toJSON(message.client) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetClientResponse>, I>>(object: I): QueryGetClientResponse {
    const message = createBaseQueryGetClientResponse();
    message.client = (object.client !== undefined && object.client !== null)
      ? Client.fromPartial(object.client)
      : undefined;
    return message;
  },
};

function createBaseQueryAllClientRequest(): QueryAllClientRequest {
  return { pagination: undefined };
}

export const QueryAllClientRequest = {
  encode(message: QueryAllClientRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllClientRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllClientRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClientRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllClientRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllClientRequest>, I>>(object: I): QueryAllClientRequest {
    const message = createBaseQueryAllClientRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllClientResponse(): QueryAllClientResponse {
  return { client: [], pagination: undefined };
}

export const QueryAllClientResponse = {
  encode(message: QueryAllClientResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.client) {
      Client.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllClientResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllClientResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.client.push(Client.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClientResponse {
    return {
      client: Array.isArray(object?.client) ? object.client.map((e: any) => Client.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllClientResponse): unknown {
    const obj: any = {};
    if (message.client) {
      obj.client = message.client.map((e) => e ? Client.toJSON(e) : undefined);
    } else {
      obj.client = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllClientResponse>, I>>(object: I): QueryAllClientResponse {
    const message = createBaseQueryAllClientResponse();
    message.client = object.client?.map((e) => Client.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetChallengerRequest(): QueryGetChallengerRequest {
  return { index: "" };
}

export const QueryGetChallengerRequest = {
  encode(message: QueryGetChallengerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChallengerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChallengerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChallengerRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetChallengerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChallengerRequest>, I>>(object: I): QueryGetChallengerRequest {
    const message = createBaseQueryGetChallengerRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetChallengerResponse(): QueryGetChallengerResponse {
  return { challenger: undefined };
}

export const QueryGetChallengerResponse = {
  encode(message: QueryGetChallengerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.challenger !== undefined) {
      Challenger.encode(message.challenger, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChallengerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChallengerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.challenger = Challenger.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChallengerResponse {
    return { challenger: isSet(object.challenger) ? Challenger.fromJSON(object.challenger) : undefined };
  },

  toJSON(message: QueryGetChallengerResponse): unknown {
    const obj: any = {};
    message.challenger !== undefined
      && (obj.challenger = message.challenger ? Challenger.toJSON(message.challenger) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChallengerResponse>, I>>(object: I): QueryGetChallengerResponse {
    const message = createBaseQueryGetChallengerResponse();
    message.challenger = (object.challenger !== undefined && object.challenger !== null)
      ? Challenger.fromPartial(object.challenger)
      : undefined;
    return message;
  },
};

function createBaseQueryAllChallengerRequest(): QueryAllChallengerRequest {
  return { pagination: undefined };
}

export const QueryAllChallengerRequest = {
  encode(message: QueryAllChallengerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllChallengerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllChallengerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllChallengerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllChallengerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllChallengerRequest>, I>>(object: I): QueryAllChallengerRequest {
    const message = createBaseQueryAllChallengerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllChallengerResponse(): QueryAllChallengerResponse {
  return { challenger: [], pagination: undefined };
}

export const QueryAllChallengerResponse = {
  encode(message: QueryAllChallengerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.challenger) {
      Challenger.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllChallengerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllChallengerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.challenger.push(Challenger.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllChallengerResponse {
    return {
      challenger: Array.isArray(object?.challenger) ? object.challenger.map((e: any) => Challenger.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllChallengerResponse): unknown {
    const obj: any = {};
    if (message.challenger) {
      obj.challenger = message.challenger.map((e) => e ? Challenger.toJSON(e) : undefined);
    } else {
      obj.challenger = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllChallengerResponse>, I>>(object: I): QueryAllChallengerResponse {
    const message = createBaseQueryAllChallengerResponse();
    message.challenger = object.challenger?.map((e) => Challenger.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRunnerRequest(): QueryGetRunnerRequest {
  return { index: "" };
}

export const QueryGetRunnerRequest = {
  encode(message: QueryGetRunnerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRunnerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRunnerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRunnerRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetRunnerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRunnerRequest>, I>>(object: I): QueryGetRunnerRequest {
    const message = createBaseQueryGetRunnerRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetRunnerResponse(): QueryGetRunnerResponse {
  return { runner: undefined };
}

export const QueryGetRunnerResponse = {
  encode(message: QueryGetRunnerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.runner !== undefined) {
      Runner.encode(message.runner, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRunnerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRunnerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.runner = Runner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRunnerResponse {
    return { runner: isSet(object.runner) ? Runner.fromJSON(object.runner) : undefined };
  },

  toJSON(message: QueryGetRunnerResponse): unknown {
    const obj: any = {};
    message.runner !== undefined && (obj.runner = message.runner ? Runner.toJSON(message.runner) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRunnerResponse>, I>>(object: I): QueryGetRunnerResponse {
    const message = createBaseQueryGetRunnerResponse();
    message.runner = (object.runner !== undefined && object.runner !== null)
      ? Runner.fromPartial(object.runner)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRunnerRequest(): QueryAllRunnerRequest {
  return { pagination: undefined };
}

export const QueryAllRunnerRequest = {
  encode(message: QueryAllRunnerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRunnerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRunnerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRunnerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRunnerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRunnerRequest>, I>>(object: I): QueryAllRunnerRequest {
    const message = createBaseQueryAllRunnerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRunnerResponse(): QueryAllRunnerResponse {
  return { runner: [], pagination: undefined };
}

export const QueryAllRunnerResponse = {
  encode(message: QueryAllRunnerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.runner) {
      Runner.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRunnerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRunnerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.runner.push(Runner.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRunnerResponse {
    return {
      runner: Array.isArray(object?.runner) ? object.runner.map((e: any) => Runner.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRunnerResponse): unknown {
    const obj: any = {};
    if (message.runner) {
      obj.runner = message.runner.map((e) => e ? Runner.toJSON(e) : undefined);
    } else {
      obj.runner = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRunnerResponse>, I>>(object: I): QueryAllRunnerResponse {
    const message = createBaseQueryAllRunnerResponse();
    message.runner = object.runner?.map((e) => Runner.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetGuardRequest(): QueryGetGuardRequest {
  return { index: "" };
}

export const QueryGetGuardRequest = {
  encode(message: QueryGetGuardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetGuardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetGuardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGuardRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetGuardRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetGuardRequest>, I>>(object: I): QueryGetGuardRequest {
    const message = createBaseQueryGetGuardRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetGuardResponse(): QueryGetGuardResponse {
  return { guard: undefined };
}

export const QueryGetGuardResponse = {
  encode(message: QueryGetGuardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.guard !== undefined) {
      Guard.encode(message.guard, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetGuardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetGuardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guard = Guard.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGuardResponse {
    return { guard: isSet(object.guard) ? Guard.fromJSON(object.guard) : undefined };
  },

  toJSON(message: QueryGetGuardResponse): unknown {
    const obj: any = {};
    message.guard !== undefined && (obj.guard = message.guard ? Guard.toJSON(message.guard) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetGuardResponse>, I>>(object: I): QueryGetGuardResponse {
    const message = createBaseQueryGetGuardResponse();
    message.guard = (object.guard !== undefined && object.guard !== null) ? Guard.fromPartial(object.guard) : undefined;
    return message;
  },
};

function createBaseQueryAllGuardRequest(): QueryAllGuardRequest {
  return { pagination: undefined };
}

export const QueryAllGuardRequest = {
  encode(message: QueryAllGuardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllGuardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllGuardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGuardRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllGuardRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllGuardRequest>, I>>(object: I): QueryAllGuardRequest {
    const message = createBaseQueryAllGuardRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllGuardResponse(): QueryAllGuardResponse {
  return { guard: [], pagination: undefined };
}

export const QueryAllGuardResponse = {
  encode(message: QueryAllGuardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.guard) {
      Guard.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllGuardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllGuardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.guard.push(Guard.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGuardResponse {
    return {
      guard: Array.isArray(object?.guard) ? object.guard.map((e: any) => Guard.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllGuardResponse): unknown {
    const obj: any = {};
    if (message.guard) {
      obj.guard = message.guard.map((e) => e ? Guard.toJSON(e) : undefined);
    } else {
      obj.guard = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllGuardResponse>, I>>(object: I): QueryAllGuardResponse {
    const message = createBaseQueryAllGuardResponse();
    message.guard = object.guard?.map((e) => Guard.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetClientByAddressRequest(): QueryGetClientByAddressRequest {
  return { address: "" };
}

export const QueryGetClientByAddressRequest = {
  encode(message: QueryGetClientByAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetClientByAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetClientByAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientByAddressRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetClientByAddressRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetClientByAddressRequest>, I>>(
    object: I,
  ): QueryGetClientByAddressRequest {
    const message = createBaseQueryGetClientByAddressRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetClientByAddressResponse(): QueryGetClientByAddressResponse {
  return { client: undefined };
}

export const QueryGetClientByAddressResponse = {
  encode(message: QueryGetClientByAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.client !== undefined) {
      Client.encode(message.client, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetClientByAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetClientByAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.client = Client.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientByAddressResponse {
    return { client: isSet(object.client) ? Client.fromJSON(object.client) : undefined };
  },

  toJSON(message: QueryGetClientByAddressResponse): unknown {
    const obj: any = {};
    message.client !== undefined && (obj.client = message.client ? Client.toJSON(message.client) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetClientByAddressResponse>, I>>(
    object: I,
  ): QueryGetClientByAddressResponse {
    const message = createBaseQueryGetClientByAddressResponse();
    message.client = (object.client !== undefined && object.client !== null)
      ? Client.fromPartial(object.client)
      : undefined;
    return message;
  },
};

function createBaseQueryGetChallengerByAddressRequest(): QueryGetChallengerByAddressRequest {
  return { address: "" };
}

export const QueryGetChallengerByAddressRequest = {
  encode(message: QueryGetChallengerByAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChallengerByAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChallengerByAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChallengerByAddressRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetChallengerByAddressRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChallengerByAddressRequest>, I>>(
    object: I,
  ): QueryGetChallengerByAddressRequest {
    const message = createBaseQueryGetChallengerByAddressRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetChallengerByAddressResponse(): QueryGetChallengerByAddressResponse {
  return { challenger: undefined };
}

export const QueryGetChallengerByAddressResponse = {
  encode(message: QueryGetChallengerByAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.challenger !== undefined) {
      Challenger.encode(message.challenger, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChallengerByAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChallengerByAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.challenger = Challenger.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChallengerByAddressResponse {
    return { challenger: isSet(object.challenger) ? Challenger.fromJSON(object.challenger) : undefined };
  },

  toJSON(message: QueryGetChallengerByAddressResponse): unknown {
    const obj: any = {};
    message.challenger !== undefined
      && (obj.challenger = message.challenger ? Challenger.toJSON(message.challenger) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChallengerByAddressResponse>, I>>(
    object: I,
  ): QueryGetChallengerByAddressResponse {
    const message = createBaseQueryGetChallengerByAddressResponse();
    message.challenger = (object.challenger !== undefined && object.challenger !== null)
      ? Challenger.fromPartial(object.challenger)
      : undefined;
    return message;
  },
};

function createBaseQueryGetVrfDataRequest(): QueryGetVrfDataRequest {
  return { index: "" };
}

export const QueryGetVrfDataRequest = {
  encode(message: QueryGetVrfDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVrfDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVrfDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVrfDataRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetVrfDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVrfDataRequest>, I>>(object: I): QueryGetVrfDataRequest {
    const message = createBaseQueryGetVrfDataRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetVrfDataResponse(): QueryGetVrfDataResponse {
  return { vrfData: undefined };
}

export const QueryGetVrfDataResponse = {
  encode(message: QueryGetVrfDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vrfData !== undefined) {
      VrfData.encode(message.vrfData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVrfDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVrfDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vrfData = VrfData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVrfDataResponse {
    return { vrfData: isSet(object.vrfData) ? VrfData.fromJSON(object.vrfData) : undefined };
  },

  toJSON(message: QueryGetVrfDataResponse): unknown {
    const obj: any = {};
    message.vrfData !== undefined && (obj.vrfData = message.vrfData ? VrfData.toJSON(message.vrfData) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVrfDataResponse>, I>>(object: I): QueryGetVrfDataResponse {
    const message = createBaseQueryGetVrfDataResponse();
    message.vrfData = (object.vrfData !== undefined && object.vrfData !== null)
      ? VrfData.fromPartial(object.vrfData)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVrfDataRequest(): QueryAllVrfDataRequest {
  return { pagination: undefined };
}

export const QueryAllVrfDataRequest = {
  encode(message: QueryAllVrfDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVrfDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVrfDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVrfDataRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllVrfDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVrfDataRequest>, I>>(object: I): QueryAllVrfDataRequest {
    const message = createBaseQueryAllVrfDataRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVrfDataResponse(): QueryAllVrfDataResponse {
  return { vrfData: [], pagination: undefined };
}

export const QueryAllVrfDataResponse = {
  encode(message: QueryAllVrfDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vrfData) {
      VrfData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVrfDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVrfDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vrfData.push(VrfData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVrfDataResponse {
    return {
      vrfData: Array.isArray(object?.vrfData) ? object.vrfData.map((e: any) => VrfData.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllVrfDataResponse): unknown {
    const obj: any = {};
    if (message.vrfData) {
      obj.vrfData = message.vrfData.map((e) => e ? VrfData.toJSON(e) : undefined);
    } else {
      obj.vrfData = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVrfDataResponse>, I>>(object: I): QueryAllVrfDataResponse {
    const message = createBaseQueryAllVrfDataResponse();
    message.vrfData = object.vrfData?.map((e) => VrfData.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetVrfUserRequest(): QueryGetVrfUserRequest {
  return { index: "" };
}

export const QueryGetVrfUserRequest = {
  encode(message: QueryGetVrfUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVrfUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVrfUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVrfUserRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetVrfUserRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVrfUserRequest>, I>>(object: I): QueryGetVrfUserRequest {
    const message = createBaseQueryGetVrfUserRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetVrfUserResponse(): QueryGetVrfUserResponse {
  return { vrfUser: undefined };
}

export const QueryGetVrfUserResponse = {
  encode(message: QueryGetVrfUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vrfUser !== undefined) {
      VrfUser.encode(message.vrfUser, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVrfUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVrfUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vrfUser = VrfUser.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVrfUserResponse {
    return { vrfUser: isSet(object.vrfUser) ? VrfUser.fromJSON(object.vrfUser) : undefined };
  },

  toJSON(message: QueryGetVrfUserResponse): unknown {
    const obj: any = {};
    message.vrfUser !== undefined && (obj.vrfUser = message.vrfUser ? VrfUser.toJSON(message.vrfUser) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVrfUserResponse>, I>>(object: I): QueryGetVrfUserResponse {
    const message = createBaseQueryGetVrfUserResponse();
    message.vrfUser = (object.vrfUser !== undefined && object.vrfUser !== null)
      ? VrfUser.fromPartial(object.vrfUser)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVrfUserRequest(): QueryAllVrfUserRequest {
  return { pagination: undefined };
}

export const QueryAllVrfUserRequest = {
  encode(message: QueryAllVrfUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVrfUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVrfUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVrfUserRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllVrfUserRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVrfUserRequest>, I>>(object: I): QueryAllVrfUserRequest {
    const message = createBaseQueryAllVrfUserRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVrfUserResponse(): QueryAllVrfUserResponse {
  return { vrfUser: [], pagination: undefined };
}

export const QueryAllVrfUserResponse = {
  encode(message: QueryAllVrfUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vrfUser) {
      VrfUser.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVrfUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVrfUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vrfUser.push(VrfUser.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVrfUserResponse {
    return {
      vrfUser: Array.isArray(object?.vrfUser) ? object.vrfUser.map((e: any) => VrfUser.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllVrfUserResponse): unknown {
    const obj: any = {};
    if (message.vrfUser) {
      obj.vrfUser = message.vrfUser.map((e) => e ? VrfUser.toJSON(e) : undefined);
    } else {
      obj.vrfUser = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVrfUserResponse>, I>>(object: I): QueryAllVrfUserResponse {
    const message = createBaseQueryAllVrfUserResponse();
    message.vrfUser = object.vrfUser?.map((e) => VrfUser.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryVerifyRandomNumberRequest(): QueryVerifyRandomNumberRequest {
  return { pubkey: "", message: "", vrv: "", proof: "" };
}

export const QueryVerifyRandomNumberRequest = {
  encode(message: QueryVerifyRandomNumberRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pubkey !== "") {
      writer.uint32(10).string(message.pubkey);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.vrv !== "") {
      writer.uint32(26).string(message.vrv);
    }
    if (message.proof !== "") {
      writer.uint32(34).string(message.proof);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifyRandomNumberRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifyRandomNumberRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pubkey = reader.string();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.vrv = reader.string();
          break;
        case 4:
          message.proof = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifyRandomNumberRequest {
    return {
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
      message: isSet(object.message) ? String(object.message) : "",
      vrv: isSet(object.vrv) ? String(object.vrv) : "",
      proof: isSet(object.proof) ? String(object.proof) : "",
    };
  },

  toJSON(message: QueryVerifyRandomNumberRequest): unknown {
    const obj: any = {};
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    message.message !== undefined && (obj.message = message.message);
    message.vrv !== undefined && (obj.vrv = message.vrv);
    message.proof !== undefined && (obj.proof = message.proof);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifyRandomNumberRequest>, I>>(
    object: I,
  ): QueryVerifyRandomNumberRequest {
    const message = createBaseQueryVerifyRandomNumberRequest();
    message.pubkey = object.pubkey ?? "";
    message.message = object.message ?? "";
    message.vrv = object.vrv ?? "";
    message.proof = object.proof ?? "";
    return message;
  },
};

function createBaseQueryVerifyRandomNumberResponse(): QueryVerifyRandomNumberResponse {
  return { result: false };
}

export const QueryVerifyRandomNumberResponse = {
  encode(message: QueryVerifyRandomNumberResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result === true) {
      writer.uint32(8).bool(message.result);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifyRandomNumberResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifyRandomNumberResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.result = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifyRandomNumberResponse {
    return { result: isSet(object.result) ? Boolean(object.result) : false };
  },

  toJSON(message: QueryVerifyRandomNumberResponse): unknown {
    const obj: any = {};
    message.result !== undefined && (obj.result = message.result);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifyRandomNumberResponse>, I>>(
    object: I,
  ): QueryVerifyRandomNumberResponse {
    const message = createBaseQueryVerifyRandomNumberResponse();
    message.result = object.result ?? false;
    return message;
  },
};

function createBaseQueryIsChallengeableRequest(): QueryIsChallengeableRequest {
  return { clientAddr: "" };
}

export const QueryIsChallengeableRequest = {
  encode(message: QueryIsChallengeableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.clientAddr !== "") {
      writer.uint32(10).string(message.clientAddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIsChallengeableRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIsChallengeableRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.clientAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIsChallengeableRequest {
    return { clientAddr: isSet(object.clientAddr) ? String(object.clientAddr) : "" };
  },

  toJSON(message: QueryIsChallengeableRequest): unknown {
    const obj: any = {};
    message.clientAddr !== undefined && (obj.clientAddr = message.clientAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIsChallengeableRequest>, I>>(object: I): QueryIsChallengeableRequest {
    const message = createBaseQueryIsChallengeableRequest();
    message.clientAddr = object.clientAddr ?? "";
    return message;
  },
};

function createBaseQueryIsChallengeableResponse(): QueryIsChallengeableResponse {
  return { resultBool: "", challengeabilityScore: "" };
}

export const QueryIsChallengeableResponse = {
  encode(message: QueryIsChallengeableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resultBool !== "") {
      writer.uint32(10).string(message.resultBool);
    }
    if (message.challengeabilityScore !== "") {
      writer.uint32(18).string(message.challengeabilityScore);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryIsChallengeableResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryIsChallengeableResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.resultBool = reader.string();
          break;
        case 2:
          message.challengeabilityScore = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryIsChallengeableResponse {
    return {
      resultBool: isSet(object.resultBool) ? String(object.resultBool) : "",
      challengeabilityScore: isSet(object.challengeabilityScore) ? String(object.challengeabilityScore) : "",
    };
  },

  toJSON(message: QueryIsChallengeableResponse): unknown {
    const obj: any = {};
    message.resultBool !== undefined && (obj.resultBool = message.resultBool);
    message.challengeabilityScore !== undefined && (obj.challengeabilityScore = message.challengeabilityScore);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryIsChallengeableResponse>, I>>(object: I): QueryIsChallengeableResponse {
    const message = createBaseQueryIsChallengeableResponse();
    message.resultBool = object.resultBool ?? "";
    message.challengeabilityScore = object.challengeabilityScore ?? "";
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Client by index. */
  Client(request: QueryGetClientRequest): Promise<QueryGetClientResponse>;
  /** Queries a list of Client items. */
  ClientAll(request: QueryAllClientRequest): Promise<QueryAllClientResponse>;
  /** Queries a Challenger by index. */
  Challenger(request: QueryGetChallengerRequest): Promise<QueryGetChallengerResponse>;
  /** Queries a list of Challenger items. */
  ChallengerAll(request: QueryAllChallengerRequest): Promise<QueryAllChallengerResponse>;
  /** Queries a Runner by index. */
  Runner(request: QueryGetRunnerRequest): Promise<QueryGetRunnerResponse>;
  /** Queries a list of Runner items. */
  RunnerAll(request: QueryAllRunnerRequest): Promise<QueryAllRunnerResponse>;
  /** Queries a Guard by index. */
  Guard(request: QueryGetGuardRequest): Promise<QueryGetGuardResponse>;
  /** Queries a list of Guard items. */
  GuardAll(request: QueryAllGuardRequest): Promise<QueryAllGuardResponse>;
  /** Queries a list of GetClientByAddress items. */
  GetClientByAddress(request: QueryGetClientByAddressRequest): Promise<QueryGetClientByAddressResponse>;
  /** Queries a list of GetChallengerByAddress items. */
  GetChallengerByAddress(request: QueryGetChallengerByAddressRequest): Promise<QueryGetChallengerByAddressResponse>;
  /** Queries a VrfData by index. */
  VrfData(request: QueryGetVrfDataRequest): Promise<QueryGetVrfDataResponse>;
  /** Queries a list of VrfData items. */
  VrfDataAll(request: QueryAllVrfDataRequest): Promise<QueryAllVrfDataResponse>;
  /** Queries a VrfUser by index. */
  VrfUser(request: QueryGetVrfUserRequest): Promise<QueryGetVrfUserResponse>;
  /** Queries a list of VrfUser items. */
  VrfUserAll(request: QueryAllVrfUserRequest): Promise<QueryAllVrfUserResponse>;
  /** Queries a list of VerifyRandomNumber items. */
  VerifyRandomNumber(request: QueryVerifyRandomNumberRequest): Promise<QueryVerifyRandomNumberResponse>;
  /** Queries a list of IsChallengeable items. */
  IsChallengeable(request: QueryIsChallengeableRequest): Promise<QueryIsChallengeableResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Client = this.Client.bind(this);
    this.ClientAll = this.ClientAll.bind(this);
    this.Challenger = this.Challenger.bind(this);
    this.ChallengerAll = this.ChallengerAll.bind(this);
    this.Runner = this.Runner.bind(this);
    this.RunnerAll = this.RunnerAll.bind(this);
    this.Guard = this.Guard.bind(this);
    this.GuardAll = this.GuardAll.bind(this);
    this.GetClientByAddress = this.GetClientByAddress.bind(this);
    this.GetChallengerByAddress = this.GetChallengerByAddress.bind(this);
    this.VrfData = this.VrfData.bind(this);
    this.VrfDataAll = this.VrfDataAll.bind(this);
    this.VrfUser = this.VrfUser.bind(this);
    this.VrfUserAll = this.VrfUserAll.bind(this);
    this.VerifyRandomNumber = this.VerifyRandomNumber.bind(this);
    this.IsChallengeable = this.IsChallengeable.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Client(request: QueryGetClientRequest): Promise<QueryGetClientResponse> {
    const data = QueryGetClientRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Client", data);
    return promise.then((data) => QueryGetClientResponse.decode(new _m0.Reader(data)));
  }

  ClientAll(request: QueryAllClientRequest): Promise<QueryAllClientResponse> {
    const data = QueryAllClientRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "ClientAll", data);
    return promise.then((data) => QueryAllClientResponse.decode(new _m0.Reader(data)));
  }

  Challenger(request: QueryGetChallengerRequest): Promise<QueryGetChallengerResponse> {
    const data = QueryGetChallengerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Challenger", data);
    return promise.then((data) => QueryGetChallengerResponse.decode(new _m0.Reader(data)));
  }

  ChallengerAll(request: QueryAllChallengerRequest): Promise<QueryAllChallengerResponse> {
    const data = QueryAllChallengerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "ChallengerAll", data);
    return promise.then((data) => QueryAllChallengerResponse.decode(new _m0.Reader(data)));
  }

  Runner(request: QueryGetRunnerRequest): Promise<QueryGetRunnerResponse> {
    const data = QueryGetRunnerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Runner", data);
    return promise.then((data) => QueryGetRunnerResponse.decode(new _m0.Reader(data)));
  }

  RunnerAll(request: QueryAllRunnerRequest): Promise<QueryAllRunnerResponse> {
    const data = QueryAllRunnerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "RunnerAll", data);
    return promise.then((data) => QueryAllRunnerResponse.decode(new _m0.Reader(data)));
  }

  Guard(request: QueryGetGuardRequest): Promise<QueryGetGuardResponse> {
    const data = QueryGetGuardRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Guard", data);
    return promise.then((data) => QueryGetGuardResponse.decode(new _m0.Reader(data)));
  }

  GuardAll(request: QueryAllGuardRequest): Promise<QueryAllGuardResponse> {
    const data = QueryAllGuardRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "GuardAll", data);
    return promise.then((data) => QueryAllGuardResponse.decode(new _m0.Reader(data)));
  }

  GetClientByAddress(request: QueryGetClientByAddressRequest): Promise<QueryGetClientByAddressResponse> {
    const data = QueryGetClientByAddressRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "GetClientByAddress", data);
    return promise.then((data) => QueryGetClientByAddressResponse.decode(new _m0.Reader(data)));
  }

  GetChallengerByAddress(request: QueryGetChallengerByAddressRequest): Promise<QueryGetChallengerByAddressResponse> {
    const data = QueryGetChallengerByAddressRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "GetChallengerByAddress", data);
    return promise.then((data) => QueryGetChallengerByAddressResponse.decode(new _m0.Reader(data)));
  }

  VrfData(request: QueryGetVrfDataRequest): Promise<QueryGetVrfDataResponse> {
    const data = QueryGetVrfDataRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "VrfData", data);
    return promise.then((data) => QueryGetVrfDataResponse.decode(new _m0.Reader(data)));
  }

  VrfDataAll(request: QueryAllVrfDataRequest): Promise<QueryAllVrfDataResponse> {
    const data = QueryAllVrfDataRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "VrfDataAll", data);
    return promise.then((data) => QueryAllVrfDataResponse.decode(new _m0.Reader(data)));
  }

  VrfUser(request: QueryGetVrfUserRequest): Promise<QueryGetVrfUserResponse> {
    const data = QueryGetVrfUserRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "VrfUser", data);
    return promise.then((data) => QueryGetVrfUserResponse.decode(new _m0.Reader(data)));
  }

  VrfUserAll(request: QueryAllVrfUserRequest): Promise<QueryAllVrfUserResponse> {
    const data = QueryAllVrfUserRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "VrfUserAll", data);
    return promise.then((data) => QueryAllVrfUserResponse.decode(new _m0.Reader(data)));
  }

  VerifyRandomNumber(request: QueryVerifyRandomNumberRequest): Promise<QueryVerifyRandomNumberResponse> {
    const data = QueryVerifyRandomNumberRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "VerifyRandomNumber", data);
    return promise.then((data) => QueryVerifyRandomNumberResponse.decode(new _m0.Reader(data)));
  }

  IsChallengeable(request: QueryIsChallengeableRequest): Promise<QueryIsChallengeableResponse> {
    const data = QueryIsChallengeableRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "IsChallengeable", data);
    return promise.then((data) => QueryIsChallengeableResponse.decode(new _m0.Reader(data)));
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
