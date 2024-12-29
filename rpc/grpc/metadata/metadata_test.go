package metadata_test

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_metadata_server(t *testing.T) {
	s := grpc.NewServer()

	pb.RegisterOrderManagement6Server(s, &OrderManagement6Impl{})

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_metadata_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(orderStreamClientInterceptor))
	if err != nil {
		panic(err)
	}

	c := pb.NewOrderManagement6Client(conn)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("k1", "v1", "k2", "v2"))
	ctx = metadata.AppendToOutgoingContext(ctx, "time",
		"raw"+strconv.FormatInt(time.Now().UnixNano(), 10))

	// RPC using the context with new metadata.
	var header, trailer metadata.MD

	// Add Order
	order := pb.Order6{
		Id:          "101",
		Items:       []string{"iPhone XS", "Mac Book Pro"},
		Destination: wrapperspb.String("San Jose, CA"),
		Price:       2300.00,
	}
	res, err := c.AddOrder(ctx, &order, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		panic(err)
	}

	log.Printf("#AddOrder## header: %v. trailer: %v", header, trailer)

	//////////

	stream, err := c.UpdateOrders(ctx)
	if err != nil {
		panic(err)
	}
	// retrieve header
	header, _ = stream.Header()
	// retrieve trailer
	trailer = stream.Trailer()

	if err := stream.Send(&pb.Order6{
		Id:          "00",
		Items:       []string{"A", "B"},
		Description: "A with B",
		Price:       0.11,
		Destination: wrapperspb.String("ABC"),
	}); err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order6{
		Id:          "01",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: wrapperspb.String("ABCDEFG"),
	}); err != nil {
		panic(err)
	}

	res, err = stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	// retrieve trailer
	trailer = stream.Trailer()

	log.Printf("##UpdateOrders## header: %v. trailer: %v", header, trailer)

	log.Printf("Update Orders Res : %s", res)
}

func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	var s string

	// 获取要发送给服务端的`metadata`
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok && len(md.Get("time")) > 0 {
		s = md.Get("time")[0]
	} else {
		// 如果没有则补充这个时间戳字段
		s = "inter" + strconv.FormatInt(time.Now().UnixNano(), 10)
		ctx = metadata.AppendToOutgoingContext(ctx, "time", s)
	}

	log.Printf("call timestamp: %s", s)

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}

// SendMsg method call.
type wrappedClientStream struct {
	method string
	grpc.ClientStream
}

func (w *wrappedClientStream) RecvMsg(m interface{}) error {
	err := w.ClientStream.RecvMsg(m)

	log.Printf("method: %s, res: %s\n", w.method, m)

	return err
}

func (w *wrappedClientStream) SendMsg(m interface{}) error {
	err := w.ClientStream.SendMsg(m)

	log.Printf("method: %s, req: %s\n", w.method, m)

	return err
}

func newWrappedClientStream(method string, s grpc.ClientStream) *wrappedClientStream {
	return &wrappedClientStream{
		method,
		s,
	}
}

func orderStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s := time.Now()
	cs, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("method: %s, latency: %s\n", method, time.Since(s))
	return newWrappedClientStream(method, cs), err
}

var _ pb.OrderManagement6Server = &OrderManagement6Impl{}

var orders6 = make(map[string]pb.Order6, 0)

type OrderManagement6Impl struct {
	pb.UnimplementedOrderManagement6Server
}

// Simple RPC
func (s *OrderManagement6Impl) AddOrder(ctx context.Context, orderReq *pb.Order6) (*wrapperspb.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)

	md, ok := metadata.FromIncomingContext(ctx)
	log.Printf("has: %t. md: %v", ok, md)

	orders6[orderReq.Id] = *orderReq

	grpc.SetHeader(ctx, metadata.Pairs("header-key1", "val1"))

	// create and send header
	header := metadata.Pairs("header-key", "val")
	grpc.SendHeader(ctx, header)

	// create and set trailer
	trailer := metadata.Pairs("trailer-key", "val")
	grpc.SetTrailer(ctx, trailer)

	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

func (s *OrderManagement6Impl) UpdateOrders(stream pb.OrderManagement6_UpdateOrdersServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	log.Printf("has: %t. md: %v", ok, md)

	// create and send header
	header := metadata.Pairs("header-key", "val")
	stream.SetHeader(header)

	// create and set trailer
	trailer := metadata.Pairs("trailer-key", "val")
	stream.SetTrailer(trailer)

	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(
				&wrapperspb.StringValue{Value: "Orders processed " + ordersStr})
		}

		// Update order
		orders6[order.Id] = *order

		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}
