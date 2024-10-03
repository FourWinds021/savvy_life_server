package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"runtime/debug"
	. "savvy_life/common"
	. "savvy_life/proto/module"
	"time"
	"xorm.io/xorm"
)

type validate interface {
	Validate() error
}

func Auth() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		if !CheckAuthEnable {
			Logger.Info("check auth not enable.")
			return handler(ctx, req)
		}

		session := Engine.NewSession()
		defer func(session *xorm.Session) {
			var err error
			err = session.Close()
			if err != nil {
				Logger.Error("session.Close failed. error:", err)
			}
		}(session)

		/*
			fullMethod := info.FullMethod
			if FullMethodAuthGet == fullMethod {
				Logger.Info("no need to check auth. method: ", fullMethod)
				return handler(ctx, req)
			}
		*/

		metaData, _ := metadata.FromIncomingContext(ctx)
		userId := metaData["x-user-id"]
		token := metaData["x-token"]

		if 0 == len(userId) || 0 == len(token) {
			Logger.Error("header X-User-Id and X-Token need.")
			err = errors.New("header X-User-Id and X-Token need")
			return resp, status.Error(codes.PermissionDenied, err.Error())
		}
		Logger.Info("header X-User-Id: ", userId[0], " X-Token: ", token[0])

		user := new(User)
		session = session.Where("user_id=?", userId[0])
		if has, err := session.Get(user); err != nil {
			Logger.Error("session.Get user failed. error:", err)
			return resp, status.Error(codes.PermissionDenied, err.Error())
		} else if !has {
			Logger.Error("no user record.")
			err = errors.New("no user record")
			return resp, status.Error(codes.PermissionDenied, err.Error())
		}
		/*
			current := time.Now().Format("2006-01-02 15:04:05")

			if user.AccessToken != token[0] || current >= user.TokenExpireTime {
				Logger.Error("check auth failed."+
					" sys token: ", user.AccessToken,
					" req token: ", token[0],
					" expire time: ", user.TokenExpireTime)
				err = errors.New("check auth failed")
				return resp, status.Error(codes.PermissionDenied, err.Error())
			}

			if FullMethodUserGet == fullMethod {
				if 0 == user.Role&TUserFRoleEAdminBit {
					Logger.Error("check auth failed, need admin user."+
						" method: ", fullMethod,
						" role: ", user.Role)
					err = errors.New("check auth failed")
					return resp, status.Error(codes.PermissionDenied, err.Error())
				}
			} else {
				if 0 == user.Role {
					Logger.Error("check auth failed."+
						" method: ", fullMethod,
						" role: ", user.Role)
					err = errors.New("check auth failed")
					return resp, status.Error(codes.PermissionDenied, err.Error())
				}
			}
		*/
		return handler(ctx, req)
	}
}

func Trace() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		start := time.Now().UnixMilli()
		reqId := ctx.Value("X-Request-Id")
		if reqId == nil {
			ctx = context.WithValue(ctx, "X-Request-Id", uuid.NewV4().String())
		}

		Logger.AddHook(NewRequestIdHook(ctx.Value("X-Request-Id").(string)))
		reqBuf, _ := json.Marshal(req)
		Logger.Info("FullMethod: ", info.FullMethod, " Request: ", string(reqBuf))
		if r, ok := req.(validate); ok {
			if err := r.Validate(); err != nil {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		resp, err = handler(ctx, req)
		end := time.Now().UnixMilli()
		spend := end - start
		respBuf, _ := json.Marshal(resp)
		Logger.Info("Response: ", string(respBuf), " Spend: ", spend, "ms")
		return resp, err
	}
}

func Recovery() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := recover(); e != nil {
				Logger.Error("Recovery FullMethod: ", info.FullMethod,
					" \nRequest: ", req,
					" \nMessage: ", e,
					" \nStack: ", string(debug.Stack()[:]))
				err = grpc.Errorf(codes.Internal, "System internal error")
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-User-Id":
		return key, true
	case "X-Token":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
