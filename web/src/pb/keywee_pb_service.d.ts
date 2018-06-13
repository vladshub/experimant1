// package: pb
// file: keywee.proto

import * as keywee_pb from "./keywee_pb";
import {grpc} from "grpc-web-client";

type APIAdd = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof keywee_pb.Item;
  readonly responseType: typeof keywee_pb.Status;
};

type APIRead = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof keywee_pb.Item;
  readonly responseType: typeof keywee_pb.Item;
};

type APIIndex = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof keywee_pb.Empty;
  readonly responseType: typeof keywee_pb.Item;
};

export class API {
  static readonly serviceName: string;
  static readonly Add: APIAdd;
  static readonly Read: APIRead;
  static readonly Index: APIIndex;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class APIClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  add(
    requestMessage: keywee_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: keywee_pb.Status|null) => void
  ): void;
  add(
    requestMessage: keywee_pb.Item,
    callback: (error: ServiceError, responseMessage: keywee_pb.Status|null) => void
  ): void;
  read(
    requestMessage: keywee_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: keywee_pb.Item|null) => void
  ): void;
  read(
    requestMessage: keywee_pb.Item,
    callback: (error: ServiceError, responseMessage: keywee_pb.Item|null) => void
  ): void;
  index(requestMessage: keywee_pb.Empty, metadata?: grpc.Metadata): ResponseStream<keywee_pb.Item>;
}

