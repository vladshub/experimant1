// package: pb
// file: keywee.proto

import * as jspb from "google-protobuf";

export class Item extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getEstimateReady(): boolean;
  setEstimateReady(value: boolean): void;

  getAudianceSize(): number;
  setAudianceSize(value: number): void;

  getContent(): string;
  setContent(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getIntroduction(): string;
  setIntroduction(value: string): void;

  clearGeoList(): void;
  getGeoList(): Array<string>;
  setGeoList(value: Array<string>): void;
  addGeo(value: string, index?: number): string;

  clearTopicsList(): void;
  getTopicsList(): Array<string>;
  setTopicsList(value: Array<string>): void;
  addTopics(value: string, index?: number): string;

  clearFacebookIntrestsList(): void;
  getFacebookIntrestsList(): Array<FacebookIntrest>;
  setFacebookIntrestsList(value: Array<FacebookIntrest>): void;
  addFacebookIntrests(value?: FacebookIntrest, index?: number): FacebookIntrest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Item.AsObject;
  static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Item;
  static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
}

export namespace Item {
  export type AsObject = {
    id: string,
    url: string,
    estimateReady: boolean,
    audianceSize: number,
    content: string,
    title: string,
    introduction: string,
    geoList: Array<string>,
    topicsList: Array<string>,
    facebookIntrestsList: Array<FacebookIntrest.AsObject>,
  }
}

export class Status extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Status.AsObject;
  static toObject(includeInstance: boolean, msg: Status): Status.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Status, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Status;
  static deserializeBinaryFromReader(message: Status, reader: jspb.BinaryReader): Status;
}

export namespace Status {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class FacebookIntrest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FacebookIntrest.AsObject;
  static toObject(includeInstance: boolean, msg: FacebookIntrest): FacebookIntrest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FacebookIntrest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FacebookIntrest;
  static deserializeBinaryFromReader(message: FacebookIntrest, reader: jspb.BinaryReader): FacebookIntrest;
}

export namespace FacebookIntrest {
  export type AsObject = {
    name: string,
    id: string,
  }
}

