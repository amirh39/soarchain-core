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

export interface PoaClient {
  index?: string;
  address?: string;
  score?: string;
  rewardMultiplier?: string;
  netEarnings?: string;
  lastTimeChallenged?: string;
  coolDownTolerance?: string;
}

export interface PoaEpochData {
  /** @format uint64 */
  totalEpochs?: string;
  epochV2VRX?: string;
  epochV2VBX?: string;
  epochV2NBX?: string;
  epochRunner?: string;
}

export interface PoaFactoryKeys {
  /** @format uint64 */
  id?: string;
  factoryCert?: string;
}

export interface PoaGuard {
  index?: string;
  guardId?: string;
  runner?: PoaRunner;
  v2XChallenger?: PoaChallenger;
  v2NChallenger?: PoaChallenger;
}

export interface PoaMasterKey {
  masterCertificate?: string;
  masterAccount?: string;
}

export interface PoaMotusWallet {
  index?: string;
  client?: PoaClient;
}

export type PoaMsgChallengeServiceResponse = object;

export type PoaMsgClaimMotusRewardsResponse = object;

export type PoaMsgClaimRunnerRewardsResponse = object;

export type PoaMsgGenClientResponse = object;

export type PoaMsgGenGuardResponse = object;

export type PoaMsgRegisterFactoryKeyResponse = object;

export type PoaMsgRunnerChallengeResponse = object;

export interface PoaMsgSelectRandomChallengerResponse {
  randomChallenger?: PoaChallenger;
}

export interface PoaMsgSelectRandomRunnerResponse {
  randomRunner?: PoaRunner;
}

export type PoaMsgUnregisterChallengerResponse = object;

export type PoaMsgUnregisterClientResponse = object;

export type PoaMsgUnregisterGuardResponse = object;

export type PoaMsgUnregisterRunnerResponse = object;

export type PoaMsgUpdateGuardResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type PoaParams = object;

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

export interface PoaQueryAllFactoryKeysResponse {
  FactoryKeys?: PoaFactoryKeys[];

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

export interface PoaQueryAllMotusWalletResponse {
  motusWallet?: PoaMotusWallet[];

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

export interface PoaQueryAllVrfDataResponse {
  vrfData?: PoaVrfData[];

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

export interface PoaQueryAllVrfUserResponse {
  vrfUser?: PoaVrfUser[];

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

export interface PoaQueryGetChallengerResponse {
  challenger?: PoaChallenger;
}

export interface PoaQueryGetClientByAddressResponse {
  client?: PoaClient;
}

export interface PoaQueryGetClientResponse {
  client?: PoaClient;
}

export interface PoaQueryGetEpochDataResponse {
  EpochData?: PoaEpochData;
}

export interface PoaQueryGetFactoryKeysResponse {
  FactoryKeys?: PoaFactoryKeys;
}

export interface PoaQueryGetGuardResponse {
  guard?: PoaGuard;
}

export interface PoaQueryGetMasterKeyResponse {
  MasterKey?: PoaMasterKey;
}

export interface PoaQueryGetMotusWalletResponse {
  motusWallet?: PoaMotusWallet;
}

export interface PoaQueryGetRunnerResponse {
  runner?: PoaRunner;
}

export interface PoaQueryGetVrfDataResponse {
  vrfData?: PoaVrfData;
}

export interface PoaQueryGetVrfUserResponse {
  vrfUser?: PoaVrfUser;
}

export interface PoaQueryIsChallengeableResponse {
  resultBool?: string;
  challengeabilityScore?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface PoaQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: PoaParams;
}

export interface PoaQueryVerifyRandomNumberResponse {
  result?: boolean;
}

export interface PoaRunner {
  index?: string;
  address?: string;
  score?: string;
  rewardMultiplier?: string;
  stakedAmount?: string;
  netEarnings?: string;
  ipAddr?: string;
  lastTimeChallenged?: string;
  coolDownTolerance?: string;
  guardAddress?: string;
}

export interface PoaVrfData {
  index?: string;
  creator?: string;
  vrv?: string;
  multiplier?: string;
  proof?: string;
  pubkey?: string;
  message?: string;
  parsedVrv?: string;
  floatVrv?: string;
  finalVrv?: string;
  finalVrvFloat?: string;
  selectedChallenger?: PoaChallenger;
  selectedRunner?: PoaRunner;
}

export interface PoaVrfUser {
  index?: string;
  address?: string;
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
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
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
   * @name QueryEpochData
   * @summary Queries a EpochData by index.
   * @request GET:/soarchain/poa/epoch_data
   */
  queryEpochData = (params: RequestParams = {}) =>
    this.request<PoaQueryGetEpochDataResponse, RpcStatus>({
      path: `/soarchain/poa/epoch_data`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFactoryKeysAll
   * @summary Queries a list of FactoryKeys items.
   * @request GET:/soarchain/poa/factory_keys
   */
  queryFactoryKeysAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllFactoryKeysResponse, RpcStatus>({
      path: `/soarchain/poa/factory_keys`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFactoryKeys
   * @summary Queries a FactoryKeys by id.
   * @request GET:/soarchain/poa/factory_keys/{id}
   */
  queryFactoryKeys = (id: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetFactoryKeysResponse, RpcStatus>({
      path: `/soarchain/poa/factory_keys/${id}`,
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
   * @name QueryIsChallengeable
   * @summary Queries a list of IsChallengeable items.
   * @request GET:/soarchain/poa/is_challengeable/{clientAddr}
   */
  queryIsChallengeable = (clientAddr: string, params: RequestParams = {}) =>
    this.request<PoaQueryIsChallengeableResponse, RpcStatus>({
      path: `/soarchain/poa/is_challengeable/${clientAddr}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMasterKey
   * @summary Queries a MasterKey by index.
   * @request GET:/soarchain/poa/master_key
   */
  queryMasterKey = (params: RequestParams = {}) =>
    this.request<PoaQueryGetMasterKeyResponse, RpcStatus>({
      path: `/soarchain/poa/master_key`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMotusWalletAll
   * @summary Queries a list of MotusWallet items.
   * @request GET:/soarchain/poa/motus_wallet
   */
  queryMotusWalletAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllMotusWalletResponse, RpcStatus>({
      path: `/soarchain/poa/motus_wallet`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMotusWallet
   * @summary Queries a MotusWallet by index.
   * @request GET:/soarchain/poa/motus_wallet/{index}
   */
  queryMotusWallet = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetMotusWalletResponse, RpcStatus>({
      path: `/soarchain/poa/motus_wallet/${index}`,
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
   * @name QueryVerifyRandomNumber
   * @summary Queries a list of VerifyRandomNumber items.
   * @request GET:/soarchain/poa/verify_random_number/{pubkey}/{message}/{vrv}/{proof}
   */
  queryVerifyRandomNumber = (pubkey: string, message: string, vrv: string, proof: string, params: RequestParams = {}) =>
    this.request<PoaQueryVerifyRandomNumberResponse, RpcStatus>({
      path: `/soarchain/poa/verify_random_number/${pubkey}/${message}/${vrv}/${proof}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVrfDataAll
   * @summary Queries a list of VrfData items.
   * @request GET:/soarchain/poa/vrf_data
   */
  queryVrfDataAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllVrfDataResponse, RpcStatus>({
      path: `/soarchain/poa/vrf_data`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVrfData
   * @summary Queries a VrfData by index.
   * @request GET:/soarchain/poa/vrf_data/{index}
   */
  queryVrfData = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetVrfDataResponse, RpcStatus>({
      path: `/soarchain/poa/vrf_data/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVrfUserAll
   * @summary Queries a list of VrfUser items.
   * @request GET:/soarchain/poa/vrf_user
   */
  queryVrfUserAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PoaQueryAllVrfUserResponse, RpcStatus>({
      path: `/soarchain/poa/vrf_user`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVrfUser
   * @summary Queries a VrfUser by index.
   * @request GET:/soarchain/poa/vrf_user/{index}
   */
  queryVrfUser = (index: string, params: RequestParams = {}) =>
    this.request<PoaQueryGetVrfUserResponse, RpcStatus>({
      path: `/soarchain/poa/vrf_user/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
