/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../poa/params";
import { Client } from "../poa/client";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Challenger } from "../poa/challenger";
import { Runner } from "../poa/runner";
import { Guard } from "../poa/guard";
import { TotalClients } from "../poa/total_clients";
import { TotalChallengers } from "../poa/total_challengers";
import { TotalRunners } from "../poa/total_runners";

export const protobufPackage = "soarchain.poa";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

export interface QueryGetTotalClientsRequest {}

export interface QueryGetTotalClientsResponse {
  TotalClients: TotalClients | undefined;
}

export interface QueryGetTotalChallengersRequest {}

export interface QueryGetTotalChallengersResponse {
  TotalChallengers: TotalChallengers | undefined;
}

export interface QueryGetTotalRunnersRequest {}

export interface QueryGetTotalRunnersResponse {
  TotalRunners: TotalRunners | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetClientRequest: object = { index: "" };

export const QueryGetClientRequest = {
  encode(
    message: QueryGetClientRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetClientRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetClientRequest } as QueryGetClientRequest;
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
    const message = { ...baseQueryGetClientRequest } as QueryGetClientRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetClientRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientRequest>
  ): QueryGetClientRequest {
    const message = { ...baseQueryGetClientRequest } as QueryGetClientRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetClientResponse: object = {};

export const QueryGetClientResponse = {
  encode(
    message: QueryGetClientResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.client !== undefined) {
      Client.encode(message.client, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetClientResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetClientResponse } as QueryGetClientResponse;
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
    const message = { ...baseQueryGetClientResponse } as QueryGetClientResponse;
    if (object.client !== undefined && object.client !== null) {
      message.client = Client.fromJSON(object.client);
    } else {
      message.client = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClientResponse): unknown {
    const obj: any = {};
    message.client !== undefined &&
      (obj.client = message.client ? Client.toJSON(message.client) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientResponse>
  ): QueryGetClientResponse {
    const message = { ...baseQueryGetClientResponse } as QueryGetClientResponse;
    if (object.client !== undefined && object.client !== null) {
      message.client = Client.fromPartial(object.client);
    } else {
      message.client = undefined;
    }
    return message;
  },
};

const baseQueryAllClientRequest: object = {};

export const QueryAllClientRequest = {
  encode(
    message: QueryAllClientRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllClientRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllClientRequest } as QueryAllClientRequest;
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
    const message = { ...baseQueryAllClientRequest } as QueryAllClientRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClientRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClientRequest>
  ): QueryAllClientRequest {
    const message = { ...baseQueryAllClientRequest } as QueryAllClientRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClientResponse: object = {};

export const QueryAllClientResponse = {
  encode(
    message: QueryAllClientResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.client) {
      Client.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllClientResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllClientResponse } as QueryAllClientResponse;
    message.client = [];
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
    const message = { ...baseQueryAllClientResponse } as QueryAllClientResponse;
    message.client = [];
    if (object.client !== undefined && object.client !== null) {
      for (const e of object.client) {
        message.client.push(Client.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClientResponse): unknown {
    const obj: any = {};
    if (message.client) {
      obj.client = message.client.map((e) =>
        e ? Client.toJSON(e) : undefined
      );
    } else {
      obj.client = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClientResponse>
  ): QueryAllClientResponse {
    const message = { ...baseQueryAllClientResponse } as QueryAllClientResponse;
    message.client = [];
    if (object.client !== undefined && object.client !== null) {
      for (const e of object.client) {
        message.client.push(Client.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetChallengerRequest: object = { index: "" };

export const QueryGetChallengerRequest = {
  encode(
    message: QueryGetChallengerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChallengerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChallengerRequest,
    } as QueryGetChallengerRequest;
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
    const message = {
      ...baseQueryGetChallengerRequest,
    } as QueryGetChallengerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetChallengerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetChallengerRequest>
  ): QueryGetChallengerRequest {
    const message = {
      ...baseQueryGetChallengerRequest,
    } as QueryGetChallengerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetChallengerResponse: object = {};

export const QueryGetChallengerResponse = {
  encode(
    message: QueryGetChallengerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.challenger !== undefined) {
      Challenger.encode(message.challenger, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChallengerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChallengerResponse,
    } as QueryGetChallengerResponse;
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
    const message = {
      ...baseQueryGetChallengerResponse,
    } as QueryGetChallengerResponse;
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromJSON(object.challenger);
    } else {
      message.challenger = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetChallengerResponse): unknown {
    const obj: any = {};
    message.challenger !== undefined &&
      (obj.challenger = message.challenger
        ? Challenger.toJSON(message.challenger)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetChallengerResponse>
  ): QueryGetChallengerResponse {
    const message = {
      ...baseQueryGetChallengerResponse,
    } as QueryGetChallengerResponse;
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromPartial(object.challenger);
    } else {
      message.challenger = undefined;
    }
    return message;
  },
};

const baseQueryAllChallengerRequest: object = {};

export const QueryAllChallengerRequest = {
  encode(
    message: QueryAllChallengerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllChallengerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllChallengerRequest,
    } as QueryAllChallengerRequest;
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
    const message = {
      ...baseQueryAllChallengerRequest,
    } as QueryAllChallengerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllChallengerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllChallengerRequest>
  ): QueryAllChallengerRequest {
    const message = {
      ...baseQueryAllChallengerRequest,
    } as QueryAllChallengerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllChallengerResponse: object = {};

export const QueryAllChallengerResponse = {
  encode(
    message: QueryAllChallengerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.challenger) {
      Challenger.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllChallengerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllChallengerResponse,
    } as QueryAllChallengerResponse;
    message.challenger = [];
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
    const message = {
      ...baseQueryAllChallengerResponse,
    } as QueryAllChallengerResponse;
    message.challenger = [];
    if (object.challenger !== undefined && object.challenger !== null) {
      for (const e of object.challenger) {
        message.challenger.push(Challenger.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllChallengerResponse): unknown {
    const obj: any = {};
    if (message.challenger) {
      obj.challenger = message.challenger.map((e) =>
        e ? Challenger.toJSON(e) : undefined
      );
    } else {
      obj.challenger = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllChallengerResponse>
  ): QueryAllChallengerResponse {
    const message = {
      ...baseQueryAllChallengerResponse,
    } as QueryAllChallengerResponse;
    message.challenger = [];
    if (object.challenger !== undefined && object.challenger !== null) {
      for (const e of object.challenger) {
        message.challenger.push(Challenger.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetRunnerRequest: object = { index: "" };

export const QueryGetRunnerRequest = {
  encode(
    message: QueryGetRunnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRunnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetRunnerRequest } as QueryGetRunnerRequest;
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
    const message = { ...baseQueryGetRunnerRequest } as QueryGetRunnerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetRunnerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRunnerRequest>
  ): QueryGetRunnerRequest {
    const message = { ...baseQueryGetRunnerRequest } as QueryGetRunnerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetRunnerResponse: object = {};

export const QueryGetRunnerResponse = {
  encode(
    message: QueryGetRunnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.runner !== undefined) {
      Runner.encode(message.runner, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRunnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetRunnerResponse } as QueryGetRunnerResponse;
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
    const message = { ...baseQueryGetRunnerResponse } as QueryGetRunnerResponse;
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromJSON(object.runner);
    } else {
      message.runner = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRunnerResponse): unknown {
    const obj: any = {};
    message.runner !== undefined &&
      (obj.runner = message.runner ? Runner.toJSON(message.runner) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRunnerResponse>
  ): QueryGetRunnerResponse {
    const message = { ...baseQueryGetRunnerResponse } as QueryGetRunnerResponse;
    if (object.runner !== undefined && object.runner !== null) {
      message.runner = Runner.fromPartial(object.runner);
    } else {
      message.runner = undefined;
    }
    return message;
  },
};

const baseQueryAllRunnerRequest: object = {};

export const QueryAllRunnerRequest = {
  encode(
    message: QueryAllRunnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRunnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllRunnerRequest } as QueryAllRunnerRequest;
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
    const message = { ...baseQueryAllRunnerRequest } as QueryAllRunnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRunnerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRunnerRequest>
  ): QueryAllRunnerRequest {
    const message = { ...baseQueryAllRunnerRequest } as QueryAllRunnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRunnerResponse: object = {};

export const QueryAllRunnerResponse = {
  encode(
    message: QueryAllRunnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.runner) {
      Runner.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRunnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllRunnerResponse } as QueryAllRunnerResponse;
    message.runner = [];
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
    const message = { ...baseQueryAllRunnerResponse } as QueryAllRunnerResponse;
    message.runner = [];
    if (object.runner !== undefined && object.runner !== null) {
      for (const e of object.runner) {
        message.runner.push(Runner.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRunnerResponse): unknown {
    const obj: any = {};
    if (message.runner) {
      obj.runner = message.runner.map((e) =>
        e ? Runner.toJSON(e) : undefined
      );
    } else {
      obj.runner = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRunnerResponse>
  ): QueryAllRunnerResponse {
    const message = { ...baseQueryAllRunnerResponse } as QueryAllRunnerResponse;
    message.runner = [];
    if (object.runner !== undefined && object.runner !== null) {
      for (const e of object.runner) {
        message.runner.push(Runner.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetGuardRequest: object = { index: "" };

export const QueryGetGuardRequest = {
  encode(
    message: QueryGetGuardRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGuardRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGuardRequest } as QueryGetGuardRequest;
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
    const message = { ...baseQueryGetGuardRequest } as QueryGetGuardRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetGuardRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetGuardRequest>): QueryGetGuardRequest {
    const message = { ...baseQueryGetGuardRequest } as QueryGetGuardRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetGuardResponse: object = {};

export const QueryGetGuardResponse = {
  encode(
    message: QueryGetGuardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.guard !== undefined) {
      Guard.encode(message.guard, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGuardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGuardResponse } as QueryGetGuardResponse;
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
    const message = { ...baseQueryGetGuardResponse } as QueryGetGuardResponse;
    if (object.guard !== undefined && object.guard !== null) {
      message.guard = Guard.fromJSON(object.guard);
    } else {
      message.guard = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGuardResponse): unknown {
    const obj: any = {};
    message.guard !== undefined &&
      (obj.guard = message.guard ? Guard.toJSON(message.guard) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGuardResponse>
  ): QueryGetGuardResponse {
    const message = { ...baseQueryGetGuardResponse } as QueryGetGuardResponse;
    if (object.guard !== undefined && object.guard !== null) {
      message.guard = Guard.fromPartial(object.guard);
    } else {
      message.guard = undefined;
    }
    return message;
  },
};

const baseQueryAllGuardRequest: object = {};

export const QueryAllGuardRequest = {
  encode(
    message: QueryAllGuardRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGuardRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllGuardRequest } as QueryAllGuardRequest;
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
    const message = { ...baseQueryAllGuardRequest } as QueryAllGuardRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGuardRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllGuardRequest>): QueryAllGuardRequest {
    const message = { ...baseQueryAllGuardRequest } as QueryAllGuardRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllGuardResponse: object = {};

export const QueryAllGuardResponse = {
  encode(
    message: QueryAllGuardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.guard) {
      Guard.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGuardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllGuardResponse } as QueryAllGuardResponse;
    message.guard = [];
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
    const message = { ...baseQueryAllGuardResponse } as QueryAllGuardResponse;
    message.guard = [];
    if (object.guard !== undefined && object.guard !== null) {
      for (const e of object.guard) {
        message.guard.push(Guard.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGuardResponse): unknown {
    const obj: any = {};
    if (message.guard) {
      obj.guard = message.guard.map((e) => (e ? Guard.toJSON(e) : undefined));
    } else {
      obj.guard = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGuardResponse>
  ): QueryAllGuardResponse {
    const message = { ...baseQueryAllGuardResponse } as QueryAllGuardResponse;
    message.guard = [];
    if (object.guard !== undefined && object.guard !== null) {
      for (const e of object.guard) {
        message.guard.push(Guard.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetClientByAddressRequest: object = { address: "" };

export const QueryGetClientByAddressRequest = {
  encode(
    message: QueryGetClientByAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientByAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientByAddressRequest,
    } as QueryGetClientByAddressRequest;
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
    const message = {
      ...baseQueryGetClientByAddressRequest,
    } as QueryGetClientByAddressRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetClientByAddressRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientByAddressRequest>
  ): QueryGetClientByAddressRequest {
    const message = {
      ...baseQueryGetClientByAddressRequest,
    } as QueryGetClientByAddressRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetClientByAddressResponse: object = {};

export const QueryGetClientByAddressResponse = {
  encode(
    message: QueryGetClientByAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.client !== undefined) {
      Client.encode(message.client, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientByAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientByAddressResponse,
    } as QueryGetClientByAddressResponse;
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
    const message = {
      ...baseQueryGetClientByAddressResponse,
    } as QueryGetClientByAddressResponse;
    if (object.client !== undefined && object.client !== null) {
      message.client = Client.fromJSON(object.client);
    } else {
      message.client = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClientByAddressResponse): unknown {
    const obj: any = {};
    message.client !== undefined &&
      (obj.client = message.client ? Client.toJSON(message.client) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientByAddressResponse>
  ): QueryGetClientByAddressResponse {
    const message = {
      ...baseQueryGetClientByAddressResponse,
    } as QueryGetClientByAddressResponse;
    if (object.client !== undefined && object.client !== null) {
      message.client = Client.fromPartial(object.client);
    } else {
      message.client = undefined;
    }
    return message;
  },
};

const baseQueryGetChallengerByAddressRequest: object = { address: "" };

export const QueryGetChallengerByAddressRequest = {
  encode(
    message: QueryGetChallengerByAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChallengerByAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChallengerByAddressRequest,
    } as QueryGetChallengerByAddressRequest;
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
    const message = {
      ...baseQueryGetChallengerByAddressRequest,
    } as QueryGetChallengerByAddressRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetChallengerByAddressRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetChallengerByAddressRequest>
  ): QueryGetChallengerByAddressRequest {
    const message = {
      ...baseQueryGetChallengerByAddressRequest,
    } as QueryGetChallengerByAddressRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetChallengerByAddressResponse: object = {};

export const QueryGetChallengerByAddressResponse = {
  encode(
    message: QueryGetChallengerByAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.challenger !== undefined) {
      Challenger.encode(message.challenger, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetChallengerByAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetChallengerByAddressResponse,
    } as QueryGetChallengerByAddressResponse;
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
    const message = {
      ...baseQueryGetChallengerByAddressResponse,
    } as QueryGetChallengerByAddressResponse;
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromJSON(object.challenger);
    } else {
      message.challenger = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetChallengerByAddressResponse): unknown {
    const obj: any = {};
    message.challenger !== undefined &&
      (obj.challenger = message.challenger
        ? Challenger.toJSON(message.challenger)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetChallengerByAddressResponse>
  ): QueryGetChallengerByAddressResponse {
    const message = {
      ...baseQueryGetChallengerByAddressResponse,
    } as QueryGetChallengerByAddressResponse;
    if (object.challenger !== undefined && object.challenger !== null) {
      message.challenger = Challenger.fromPartial(object.challenger);
    } else {
      message.challenger = undefined;
    }
    return message;
  },
};

const baseQueryGetTotalClientsRequest: object = {};

export const QueryGetTotalClientsRequest = {
  encode(
    _: QueryGetTotalClientsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalClientsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalClientsRequest,
    } as QueryGetTotalClientsRequest;
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

  fromJSON(_: any): QueryGetTotalClientsRequest {
    const message = {
      ...baseQueryGetTotalClientsRequest,
    } as QueryGetTotalClientsRequest;
    return message;
  },

  toJSON(_: QueryGetTotalClientsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTotalClientsRequest>
  ): QueryGetTotalClientsRequest {
    const message = {
      ...baseQueryGetTotalClientsRequest,
    } as QueryGetTotalClientsRequest;
    return message;
  },
};

const baseQueryGetTotalClientsResponse: object = {};

export const QueryGetTotalClientsResponse = {
  encode(
    message: QueryGetTotalClientsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TotalClients !== undefined) {
      TotalClients.encode(
        message.TotalClients,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalClientsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalClientsResponse,
    } as QueryGetTotalClientsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TotalClients = TotalClients.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTotalClientsResponse {
    const message = {
      ...baseQueryGetTotalClientsResponse,
    } as QueryGetTotalClientsResponse;
    if (object.TotalClients !== undefined && object.TotalClients !== null) {
      message.TotalClients = TotalClients.fromJSON(object.TotalClients);
    } else {
      message.TotalClients = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTotalClientsResponse): unknown {
    const obj: any = {};
    message.TotalClients !== undefined &&
      (obj.TotalClients = message.TotalClients
        ? TotalClients.toJSON(message.TotalClients)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTotalClientsResponse>
  ): QueryGetTotalClientsResponse {
    const message = {
      ...baseQueryGetTotalClientsResponse,
    } as QueryGetTotalClientsResponse;
    if (object.TotalClients !== undefined && object.TotalClients !== null) {
      message.TotalClients = TotalClients.fromPartial(object.TotalClients);
    } else {
      message.TotalClients = undefined;
    }
    return message;
  },
};

const baseQueryGetTotalChallengersRequest: object = {};

export const QueryGetTotalChallengersRequest = {
  encode(
    _: QueryGetTotalChallengersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalChallengersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalChallengersRequest,
    } as QueryGetTotalChallengersRequest;
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

  fromJSON(_: any): QueryGetTotalChallengersRequest {
    const message = {
      ...baseQueryGetTotalChallengersRequest,
    } as QueryGetTotalChallengersRequest;
    return message;
  },

  toJSON(_: QueryGetTotalChallengersRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTotalChallengersRequest>
  ): QueryGetTotalChallengersRequest {
    const message = {
      ...baseQueryGetTotalChallengersRequest,
    } as QueryGetTotalChallengersRequest;
    return message;
  },
};

const baseQueryGetTotalChallengersResponse: object = {};

export const QueryGetTotalChallengersResponse = {
  encode(
    message: QueryGetTotalChallengersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TotalChallengers !== undefined) {
      TotalChallengers.encode(
        message.TotalChallengers,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalChallengersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalChallengersResponse,
    } as QueryGetTotalChallengersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TotalChallengers = TotalChallengers.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTotalChallengersResponse {
    const message = {
      ...baseQueryGetTotalChallengersResponse,
    } as QueryGetTotalChallengersResponse;
    if (
      object.TotalChallengers !== undefined &&
      object.TotalChallengers !== null
    ) {
      message.TotalChallengers = TotalChallengers.fromJSON(
        object.TotalChallengers
      );
    } else {
      message.TotalChallengers = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTotalChallengersResponse): unknown {
    const obj: any = {};
    message.TotalChallengers !== undefined &&
      (obj.TotalChallengers = message.TotalChallengers
        ? TotalChallengers.toJSON(message.TotalChallengers)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTotalChallengersResponse>
  ): QueryGetTotalChallengersResponse {
    const message = {
      ...baseQueryGetTotalChallengersResponse,
    } as QueryGetTotalChallengersResponse;
    if (
      object.TotalChallengers !== undefined &&
      object.TotalChallengers !== null
    ) {
      message.TotalChallengers = TotalChallengers.fromPartial(
        object.TotalChallengers
      );
    } else {
      message.TotalChallengers = undefined;
    }
    return message;
  },
};

const baseQueryGetTotalRunnersRequest: object = {};

export const QueryGetTotalRunnersRequest = {
  encode(
    _: QueryGetTotalRunnersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalRunnersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalRunnersRequest,
    } as QueryGetTotalRunnersRequest;
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

  fromJSON(_: any): QueryGetTotalRunnersRequest {
    const message = {
      ...baseQueryGetTotalRunnersRequest,
    } as QueryGetTotalRunnersRequest;
    return message;
  },

  toJSON(_: QueryGetTotalRunnersRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetTotalRunnersRequest>
  ): QueryGetTotalRunnersRequest {
    const message = {
      ...baseQueryGetTotalRunnersRequest,
    } as QueryGetTotalRunnersRequest;
    return message;
  },
};

const baseQueryGetTotalRunnersResponse: object = {};

export const QueryGetTotalRunnersResponse = {
  encode(
    message: QueryGetTotalRunnersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TotalRunners !== undefined) {
      TotalRunners.encode(
        message.TotalRunners,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTotalRunnersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTotalRunnersResponse,
    } as QueryGetTotalRunnersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TotalRunners = TotalRunners.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTotalRunnersResponse {
    const message = {
      ...baseQueryGetTotalRunnersResponse,
    } as QueryGetTotalRunnersResponse;
    if (object.TotalRunners !== undefined && object.TotalRunners !== null) {
      message.TotalRunners = TotalRunners.fromJSON(object.TotalRunners);
    } else {
      message.TotalRunners = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTotalRunnersResponse): unknown {
    const obj: any = {};
    message.TotalRunners !== undefined &&
      (obj.TotalRunners = message.TotalRunners
        ? TotalRunners.toJSON(message.TotalRunners)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTotalRunnersResponse>
  ): QueryGetTotalRunnersResponse {
    const message = {
      ...baseQueryGetTotalRunnersResponse,
    } as QueryGetTotalRunnersResponse;
    if (object.TotalRunners !== undefined && object.TotalRunners !== null) {
      message.TotalRunners = TotalRunners.fromPartial(object.TotalRunners);
    } else {
      message.TotalRunners = undefined;
    }
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
  Challenger(
    request: QueryGetChallengerRequest
  ): Promise<QueryGetChallengerResponse>;
  /** Queries a list of Challenger items. */
  ChallengerAll(
    request: QueryAllChallengerRequest
  ): Promise<QueryAllChallengerResponse>;
  /** Queries a Runner by index. */
  Runner(request: QueryGetRunnerRequest): Promise<QueryGetRunnerResponse>;
  /** Queries a list of Runner items. */
  RunnerAll(request: QueryAllRunnerRequest): Promise<QueryAllRunnerResponse>;
  /** Queries a Guard by index. */
  Guard(request: QueryGetGuardRequest): Promise<QueryGetGuardResponse>;
  /** Queries a list of Guard items. */
  GuardAll(request: QueryAllGuardRequest): Promise<QueryAllGuardResponse>;
  /** Queries a list of GetClientByAddress items. */
  GetClientByAddress(
    request: QueryGetClientByAddressRequest
  ): Promise<QueryGetClientByAddressResponse>;
  /** Queries a list of GetChallengerByAddress items. */
  GetChallengerByAddress(
    request: QueryGetChallengerByAddressRequest
  ): Promise<QueryGetChallengerByAddressResponse>;
  /** Queries a TotalClients by index. */
  TotalClients(
    request: QueryGetTotalClientsRequest
  ): Promise<QueryGetTotalClientsResponse>;
  /** Queries a TotalChallengers by index. */
  TotalChallengers(
    request: QueryGetTotalChallengersRequest
  ): Promise<QueryGetTotalChallengersResponse>;
  /** Queries a TotalRunners by index. */
  TotalRunners(
    request: QueryGetTotalRunnersRequest
  ): Promise<QueryGetTotalRunnersResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Client(request: QueryGetClientRequest): Promise<QueryGetClientResponse> {
    const data = QueryGetClientRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Client", data);
    return promise.then((data) =>
      QueryGetClientResponse.decode(new Reader(data))
    );
  }

  ClientAll(request: QueryAllClientRequest): Promise<QueryAllClientResponse> {
    const data = QueryAllClientRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "ClientAll", data);
    return promise.then((data) =>
      QueryAllClientResponse.decode(new Reader(data))
    );
  }

  Challenger(
    request: QueryGetChallengerRequest
  ): Promise<QueryGetChallengerResponse> {
    const data = QueryGetChallengerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Challenger", data);
    return promise.then((data) =>
      QueryGetChallengerResponse.decode(new Reader(data))
    );
  }

  ChallengerAll(
    request: QueryAllChallengerRequest
  ): Promise<QueryAllChallengerResponse> {
    const data = QueryAllChallengerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "ChallengerAll",
      data
    );
    return promise.then((data) =>
      QueryAllChallengerResponse.decode(new Reader(data))
    );
  }

  Runner(request: QueryGetRunnerRequest): Promise<QueryGetRunnerResponse> {
    const data = QueryGetRunnerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Runner", data);
    return promise.then((data) =>
      QueryGetRunnerResponse.decode(new Reader(data))
    );
  }

  RunnerAll(request: QueryAllRunnerRequest): Promise<QueryAllRunnerResponse> {
    const data = QueryAllRunnerRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "RunnerAll", data);
    return promise.then((data) =>
      QueryAllRunnerResponse.decode(new Reader(data))
    );
  }

  Guard(request: QueryGetGuardRequest): Promise<QueryGetGuardResponse> {
    const data = QueryGetGuardRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "Guard", data);
    return promise.then((data) =>
      QueryGetGuardResponse.decode(new Reader(data))
    );
  }

  GuardAll(request: QueryAllGuardRequest): Promise<QueryAllGuardResponse> {
    const data = QueryAllGuardRequest.encode(request).finish();
    const promise = this.rpc.request("soarchain.poa.Query", "GuardAll", data);
    return promise.then((data) =>
      QueryAllGuardResponse.decode(new Reader(data))
    );
  }

  GetClientByAddress(
    request: QueryGetClientByAddressRequest
  ): Promise<QueryGetClientByAddressResponse> {
    const data = QueryGetClientByAddressRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "GetClientByAddress",
      data
    );
    return promise.then((data) =>
      QueryGetClientByAddressResponse.decode(new Reader(data))
    );
  }

  GetChallengerByAddress(
    request: QueryGetChallengerByAddressRequest
  ): Promise<QueryGetChallengerByAddressResponse> {
    const data = QueryGetChallengerByAddressRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "GetChallengerByAddress",
      data
    );
    return promise.then((data) =>
      QueryGetChallengerByAddressResponse.decode(new Reader(data))
    );
  }

  TotalClients(
    request: QueryGetTotalClientsRequest
  ): Promise<QueryGetTotalClientsResponse> {
    const data = QueryGetTotalClientsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "TotalClients",
      data
    );
    return promise.then((data) =>
      QueryGetTotalClientsResponse.decode(new Reader(data))
    );
  }

  TotalChallengers(
    request: QueryGetTotalChallengersRequest
  ): Promise<QueryGetTotalChallengersResponse> {
    const data = QueryGetTotalChallengersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "TotalChallengers",
      data
    );
    return promise.then((data) =>
      QueryGetTotalChallengersResponse.decode(new Reader(data))
    );
  }

  TotalRunners(
    request: QueryGetTotalRunnersRequest
  ): Promise<QueryGetTotalRunnersResponse> {
    const data = QueryGetTotalRunnersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soarchain.poa.Query",
      "TotalRunners",
      data
    );
    return promise.then((data) =>
      QueryGetTotalRunnersResponse.decode(new Reader(data))
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
