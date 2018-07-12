#!/usr/bin/env python

import sys, grpc, grpcbasic0_pb2, grpcbasic0_pb2_grpc

if __name__ == '__main__':
    strt = int((sys.argv[1:2]+[0])[0])
    chan = grpc.insecure_channel('localhost:3323')
    stub = grpcbasic0_pb2_grpc.UserServiceStub(chan)
    resp = stub.GetUsers(grpcbasic0_pb2.GetUsersReq(start=strt, count=2))
    tmpl = "ID: {}, Name: {}, Fortune: {}"
    for u in resp.users:
        print(tmpl.format(u.id, u.name, u.fortune))
