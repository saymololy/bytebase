// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: store/instance.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { Duration } from "../google/protobuf/duration";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "bytebase.store";

/** InstanceMetadata is the metadata for instances. */
export interface InstanceMetadata {
  /**
   * The lower_case_table_names config for MySQL instances.
   * It is used to determine whether the table names and database names are case sensitive.
   */
  mysqlLowerCaseTableNames: number;
  lastSyncTime: Timestamp | undefined;
  roles: InstanceRole[];
  /** How often the instance is synced. */
  syncInterval:
    | Duration
    | undefined;
  /**
   * The maximum number of connections.
   * The default is 10 if the value is unset or zero.
   */
  maximumConnections: number;
  /**
   * Enable sync for following databases.
   * Default empty, means sync all schemas & databases.
   */
  syncDatabases: string[];
}

/** InstanceRole is the API message for instance role. */
export interface InstanceRole {
  /** The role name. */
  name: string;
  /** The connection count limit for this role. */
  connectionLimit?:
    | number
    | undefined;
  /** The expiration for the role's password. */
  validUntil?:
    | string
    | undefined;
  /**
   * The role attribute.
   * For PostgreSQL, it containt super_user, no_inherit, create_role, create_db, can_login, replication and bypass_rls. Docs: https://www.postgresql.org/docs/current/role-attributes.html
   * For MySQL, it's the global privileges as GRANT statements, which means it only contains "GRANT ... ON *.* TO ...". Docs: https://dev.mysql.com/doc/refman/8.0/en/grant.html
   */
  attribute?: string | undefined;
}

function createBaseInstanceMetadata(): InstanceMetadata {
  return {
    mysqlLowerCaseTableNames: 0,
    lastSyncTime: undefined,
    roles: [],
    syncInterval: undefined,
    maximumConnections: 0,
    syncDatabases: [],
  };
}

export const InstanceMetadata: MessageFns<InstanceMetadata> = {
  encode(message: InstanceMetadata, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.mysqlLowerCaseTableNames !== 0) {
      writer.uint32(8).int32(message.mysqlLowerCaseTableNames);
    }
    if (message.lastSyncTime !== undefined) {
      Timestamp.encode(message.lastSyncTime, writer.uint32(18).fork()).join();
    }
    for (const v of message.roles) {
      InstanceRole.encode(v!, writer.uint32(26).fork()).join();
    }
    if (message.syncInterval !== undefined) {
      Duration.encode(message.syncInterval, writer.uint32(34).fork()).join();
    }
    if (message.maximumConnections !== 0) {
      writer.uint32(40).int32(message.maximumConnections);
    }
    for (const v of message.syncDatabases) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): InstanceMetadata {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInstanceMetadata();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.mysqlLowerCaseTableNames = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.lastSyncTime = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.roles.push(InstanceRole.decode(reader, reader.uint32()));
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.syncInterval = Duration.decode(reader, reader.uint32());
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.maximumConnections = reader.int32();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.syncDatabases.push(reader.string());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InstanceMetadata {
    return {
      mysqlLowerCaseTableNames: isSet(object.mysqlLowerCaseTableNames)
        ? globalThis.Number(object.mysqlLowerCaseTableNames)
        : 0,
      lastSyncTime: isSet(object.lastSyncTime) ? fromJsonTimestamp(object.lastSyncTime) : undefined,
      roles: globalThis.Array.isArray(object?.roles) ? object.roles.map((e: any) => InstanceRole.fromJSON(e)) : [],
      syncInterval: isSet(object.syncInterval) ? Duration.fromJSON(object.syncInterval) : undefined,
      maximumConnections: isSet(object.maximumConnections) ? globalThis.Number(object.maximumConnections) : 0,
      syncDatabases: globalThis.Array.isArray(object?.syncDatabases)
        ? object.syncDatabases.map((e: any) => globalThis.String(e))
        : [],
    };
  },

  toJSON(message: InstanceMetadata): unknown {
    const obj: any = {};
    if (message.mysqlLowerCaseTableNames !== 0) {
      obj.mysqlLowerCaseTableNames = Math.round(message.mysqlLowerCaseTableNames);
    }
    if (message.lastSyncTime !== undefined) {
      obj.lastSyncTime = fromTimestamp(message.lastSyncTime).toISOString();
    }
    if (message.roles?.length) {
      obj.roles = message.roles.map((e) => InstanceRole.toJSON(e));
    }
    if (message.syncInterval !== undefined) {
      obj.syncInterval = Duration.toJSON(message.syncInterval);
    }
    if (message.maximumConnections !== 0) {
      obj.maximumConnections = Math.round(message.maximumConnections);
    }
    if (message.syncDatabases?.length) {
      obj.syncDatabases = message.syncDatabases;
    }
    return obj;
  },

  create(base?: DeepPartial<InstanceMetadata>): InstanceMetadata {
    return InstanceMetadata.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<InstanceMetadata>): InstanceMetadata {
    const message = createBaseInstanceMetadata();
    message.mysqlLowerCaseTableNames = object.mysqlLowerCaseTableNames ?? 0;
    message.lastSyncTime = (object.lastSyncTime !== undefined && object.lastSyncTime !== null)
      ? Timestamp.fromPartial(object.lastSyncTime)
      : undefined;
    message.roles = object.roles?.map((e) => InstanceRole.fromPartial(e)) || [];
    message.syncInterval = (object.syncInterval !== undefined && object.syncInterval !== null)
      ? Duration.fromPartial(object.syncInterval)
      : undefined;
    message.maximumConnections = object.maximumConnections ?? 0;
    message.syncDatabases = object.syncDatabases?.map((e) => e) || [];
    return message;
  },
};

function createBaseInstanceRole(): InstanceRole {
  return { name: "", connectionLimit: undefined, validUntil: undefined, attribute: undefined };
}

export const InstanceRole: MessageFns<InstanceRole> = {
  encode(message: InstanceRole, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.connectionLimit !== undefined) {
      writer.uint32(16).int32(message.connectionLimit);
    }
    if (message.validUntil !== undefined) {
      writer.uint32(26).string(message.validUntil);
    }
    if (message.attribute !== undefined) {
      writer.uint32(34).string(message.attribute);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): InstanceRole {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInstanceRole();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.connectionLimit = reader.int32();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.validUntil = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.attribute = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InstanceRole {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      connectionLimit: isSet(object.connectionLimit) ? globalThis.Number(object.connectionLimit) : undefined,
      validUntil: isSet(object.validUntil) ? globalThis.String(object.validUntil) : undefined,
      attribute: isSet(object.attribute) ? globalThis.String(object.attribute) : undefined,
    };
  },

  toJSON(message: InstanceRole): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.connectionLimit !== undefined) {
      obj.connectionLimit = Math.round(message.connectionLimit);
    }
    if (message.validUntil !== undefined) {
      obj.validUntil = message.validUntil;
    }
    if (message.attribute !== undefined) {
      obj.attribute = message.attribute;
    }
    return obj;
  },

  create(base?: DeepPartial<InstanceRole>): InstanceRole {
    return InstanceRole.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<InstanceRole>): InstanceRole {
    const message = createBaseInstanceRole();
    message.name = object.name ?? "";
    message.connectionLimit = object.connectionLimit ?? undefined;
    message.validUntil = object.validUntil ?? undefined;
    message.attribute = object.attribute ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = numberToLong(Math.trunc(date.getTime() / 1_000));
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds.toNumber() || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function numberToLong(number: number) {
  return Long.fromNumber(number);
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
