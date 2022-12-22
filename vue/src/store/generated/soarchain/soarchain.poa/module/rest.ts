/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface PoaChallenger {
  index?: string;
  address?: string;
  score?: string;
  stakedAmount?: string;
  netEarnings?: string;
  type?: string;
  ipAddr?: string;
}

export interface PoaChallengerByIndex {
  index?: string;
  challenger?: PoaChallenger;
}

export interface PoaClient {
  index?: string;
  address?: string;
  score?: string;
  netEarnings?: string;
  lastTimeChallenged?: string;
}

export interface PoaGuard {
  index?: string;
  guardId?: string;
  runner?: PoaRunner;
  v2XChallenger?: PoaChallenger;
  v2NChallenger?: PoaChallenger;
}

export type PoaMsgChallengeServiceResponse = object;

export type PoaMsgCreateTotalClientsResponse = object;

export type PoaMsgDeleteTotalClientsResponse = object;

export type PoaMsgGenClientResponse = object;

export type PoaMsgGenGuardResponse = object;

export type PoaMsgRunnerChallengeResponse = object;

export type PoaMsgUnregisterChallengerResponse = object;

export type PoaMsgUnregisterClientResponse = object;

export type PoaMsgUnregisterGuardResponse = object;

export type PoaMsgUnregisterRunnerResponse = object;

export type PoaMsgUpdateTotalClientsResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type PoaParams = object;

export interface PoaQueryAllChallengerByIndexResponse {
  challengerByIndex?: PoaChallengerByIndex[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryAllChallengerResponse {
  challenger?: PoaChallenger[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryAllClientResponse {
  client?: PoaClient[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryAllGuardResponse {
  guard?: PoaGuard[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryAllRunnerByIndexResponse {
  runnerByIndex?: PoaRunnerByIndex[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryAllRunnerResponse {
  runner?: PoaRunner[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface PoaQueryGetChallengerByAddressResponse {
  challenger?: PoaChallenger;
}

export interface PoaQueryGetChallengerByIndexResponse {
  challengerByIndex?: PoaChallengerByIndex;
}

export interface PoaQueryGetChallengerResponse {
  challenger?: PoaChallenger;
}

export interface PoaQueryGetClientByAddressResponse {
  client?: PoaClient;
}

export interface PoaQueryGetClientResponse {
  client?: PoaClient;
}

export interface PoaQueryGetGuardResponse {
  guard?: PoaGuard;
}

export interface PoaQueryGetRandomChallengerResponse {
  challenger?: PoaChallenger;
}

export interface PoaQueryGetRandomRunnerResponse {
  runner?: PoaRunner;
}

export interface PoaQueryGetRunnerByIndexResponse {
  runnerByIndex?: PoaRunnerByIndex;
}

export interface PoaQueryGetRunnerResponse {
  runner?: PoaRunner;
}

export interface PoaQueryGetTotalChallengersResponse {
  TotalChallengers?: PoaTotalChallengers;
}

export interface PoaQueryGetTotalClientsResponse {
  TotalClients?: PoaTotalClients;
}

export interface PoaQueryGetTotalRunnersResponse {
  TotalRunners?: PoaTotalRunners;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface PoaQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: PoaParams;
}

export interface PoaRunner {
  index?: string;
  address?: string;
  score?: string;
  stakedAmount?: string;
  netEarnings?: string;
  ipAddr?: string;
  lastTimeChallenged?: string;
}

export interface PoaRunnerByIndex {
  index?: string;
  runner?: PoaRunner;
}

export interface PoaTotalChallengers {
  /** @format uint64 */
  count?: string;
}

export interface PoaTotalClients {
  /** @format uint64 */
  count?: string;
  creator?: string;
}

export interface PoaTotalRunners {
  /** @format uint64 */
  count?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title poa/challenger.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryChallengerAll
   * @summary Queries a list of Challenger items.
   * @request GET:/soarchain/poa/challenger
   */
  queryChallengerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllChallengerResponse, RpcStatus>({
      path: `/soarchain/poa/challenger`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChallenger
   * @summary Queries a Challenger by index.
   * @request GET:/soarchain/poa/challenger/{index}
   */
  queryChallenger = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetChallengerResponse, RpcStatus>({
      path: `/soarchain/poa/challenger/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChallengerByIndexAll
   * @summary Queries a list of ChallengerByIndex items.
   * @request GET:/soarchain/poa/challenger_by_index
   */
  queryChallengerByIndexAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllChallengerByIndexResponse, RpcStatus>({
      path: `/soarchain/poa/challenger_by_index`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChallengerByIndex
   * @summary Queries a ChallengerByIndex by index.
   * @request GET:/soarchain/poa/challenger_by_index/{index}
   */
  queryChallengerByIndex = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetChallengerByIndexResponse, RpcStatus>({
      path: `/soarchain/poa/challenger_by_index/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClientAll
   * @summary Queries a list of Client items.
   * @request GET:/soarchain/poa/client
   */
  queryClientAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllClientResponse, RpcStatus>({
      path: `/soarchain/poa/client`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClient
   * @summary Queries a Client by index.
   * @request GET:/soarchain/poa/client/{index}
   */
  queryClient = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetClientResponse, RpcStatus>({
      path: `/soarchain/poa/client/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetChallengerByAddress
   * @summary Queries a list of GetChallengerByAddress items.
   * @request GET:/soarchain/poa/get_challenger_by_address/{address}
   */
  queryGetChallengerByAddress = (address: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetChallengerByAddressResponse, RpcStatus>({
      path: `/soarchain/poa/get_challenger_by_address/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetClientByAddress
   * @summary Queries a list of GetClientByAddress items.
   * @request GET:/soarchain/poa/get_client_by_address/{address}
   */
  queryGetClientByAddress = (address: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetClientByAddressResponse, RpcStatus>({
      path: `/soarchain/poa/get_client_by_address/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetRandomChallenger
   * @summary Queries a list of GetRandomChallenger items.
   * @request GET:/soarchain/poa/get_random_challenger
   */
  queryGetRandomChallenger = (params: RequestParams = {}) =>
    this.request<PoaQueryGetRandomChallengerResponse, RpcStatus>({
      path: `/soarchain/poa/get_random_challenger`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetRandomRunner
   * @summary Queries a list of GetRandomRunner items.
   * @request GET:/soarchain/poa/get_random_runner
   */
  queryGetRandomRunner = (params: RequestParams = {}) =>
    this.request<PoaQueryGetRandomRunnerResponse, RpcStatus>({
      path: `/soarchain/poa/get_random_runner`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGuardAll
   * @summary Queries a list of Guard items.
   * @request GET:/soarchain/poa/guard
   */
  queryGuardAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllGuardResponse, RpcStatus>({
      path: `/soarchain/poa/guard`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGuard
   * @summary Queries a Guard by index.
   * @request GET:/soarchain/poa/guard/{index}
   */
  queryGuard = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetGuardResponse, RpcStatus>({
      path: `/soarchain/poa/guard/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/soarchain/poa/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<PoaQueryParamsResponse, RpcStatus>({
      path: `/soarchain/poa/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRunnerAll
   * @summary Queries a list of Runner items.
   * @request GET:/soarchain/poa/runner
   */
  queryRunnerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllRunnerResponse, RpcStatus>({
      path: `/soarchain/poa/runner`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRunner
   * @summary Queries a Runner by index.
   * @request GET:/soarchain/poa/runner/{index}
   */
  queryRunner = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetRunnerResponse, RpcStatus>({
      path: `/soarchain/poa/runner/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRunnerByIndexAll
   * @summary Queries a list of RunnerByIndex items.
   * @request GET:/soarchain/poa/runner_by_index
   */
  queryRunnerByIndexAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllRunnerByIndexResponse, RpcStatus>({
      path: `/soarchain/poa/runner_by_index`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRunnerByIndex
   * @summary Queries a RunnerByIndex by index.
   * @request GET:/soarchain/poa/runner_by_index/{index}
   */
  queryRunnerByIndex = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetRunnerByIndexResponse, RpcStatus>({
      path: `/soarchain/poa/runner_by_index/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalChallengers
   * @summary Queries a TotalChallengers by index.
   * @request GET:/soarchain/poa/total_challengers
   */
  queryTotalChallengers = (params: RequestParams = {}) =>
    this.request<PoaQueryGetTotalChallengersResponse, RpcStatus>({
      path: `/soarchain/poa/total_challengers`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalClients
   * @summary Queries a TotalClients by index.
   * @request GET:/soarchain/poa/total_clients
   */
  queryTotalClients = (params: RequestParams = {}) =>
    this.request<PoaQueryGetTotalClientsResponse, RpcStatus>({
      path: `/soarchain/poa/total_clients`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalRunners
   * @summary Queries a TotalRunners by index.
   * @request GET:/soarchain/poa/total_runners
   */
  queryTotalRunners = (params: RequestParams = {}) =>
    this.request<PoaQueryGetTotalRunnersResponse, RpcStatus>({
      path: `/soarchain/poa/total_runners`,
      method: "GET",
      format: "json",
      ...params,
    });
}
