import {grpc,} from "grpc-web-client";
import {API} from "./pb/keywee_pb_service"
import {Empty, Item} from "./pb/keywee_pb";

const host = '';
const emptry = new Empty();
const SitesAPI = {
    _all: function (success, fail) {
        const msgs = {};
        const client = grpc.client(API.Index, {
            host: host,
        });
        client.onHeaders((headers: grpc.Metadata) => {
            console.log("index.onHeaders", headers);
        });
        client.onMessage((message: Item) => {
            const msg = message.toObject();
            console.log("index.onMessage", message.toObject());
            msgs[msg.id] = msg;
        });
        client.onEnd((code: grpc.Code, msg: string, trailers: grpc.Metadata) => {
            console.log("index.onEnd", code, msg, trailers);
            if (code === grpc.Code.OK && success !== undefined) {
                success(msgs);
            } else if (code !== grpc.Code.OK && fail !== undefined) {
                fail();
            }

        });
        client.start();
        client.send(emptry);
    },
    _get: function (id, success, fail) {
        let item = new Item();
        item.setId(id);
        grpc.unary(API.Read, {
            request: item,
            host: host,
            onEnd: ({status, statusMessage, headers, message, trailers}) => {
                console.log("Get.onEnd.status", status, statusMessage);
                console.log("Get.onEnd.headers", headers);
                if (status === grpc.Code.OK && message) {
                    const msg = message.toObject();
                    console.log("Get.onEnd.message", msg);
                    if (success !== undefined) {
                        success(msg);
                    }
                } else {
                    if (fail !== undefined) {
                        fail(status);
                    }
                }
                console.log("Get.onEnd.trailers", trailers);
            }
        });
    },
    _add: function (url, success, fail) {
        let item = new Item();
        item.setUrl(url);
        grpc.unary(API.Add, {
            request: item,
            host: host,
            onEnd: ({status, statusMessage, headers, message, trailers}) => {
                console.log("Add.onEnd.status", status, statusMessage);
                console.log("Add.onEnd.headers", headers);
                if (status === grpc.Code.OK && message) {
                    const msg = message.toObject();
                    console.log("Add.onEnd.message", msg);
                    if (success !== undefined) {
                        success(msg);
                    }
                } else {
                    if (fail !== undefined) {
                        fail(status);
                    }
                }
                console.log("Add.onEnd.trailers", trailers);
            }
        });
    },
    all: function () {
        return new Promise((resolve, reject) => {
            this._all(resolve, reject);
        });
    },
    get: function (siteId) {
        return new Promise((resolve, reject) => {
            this._get(siteId, resolve, reject);
        });
    },
    add: function (url) {
        return new Promise((resolve, reject) => {
            this._add(url, resolve, reject);
        });
    }
};

export default SitesAPI
