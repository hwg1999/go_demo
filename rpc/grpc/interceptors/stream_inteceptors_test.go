package interceptors_test

import (
	"context"
	"io"
	"log"
	"net"
	"testing"
	"time"

	pb "github.com/hwg1999/go_demo/rpc/grpc/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_streaming_interceptors_server(t *testing.T) {
	s := grpc.NewServer(grpc.ChainStreamInterceptor(orderStreamServerInterceptor1, orderStreamServerInterceptor2))

	pb.RegisterOrderManagement4Server(s, &OrderManagement4Impl{})

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func Test_streaming_interceptors_client(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(
			orderStreamClientInterceptor1,
			orderStreamClientInterceptor2,
		))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewOrderManagement4Client(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	stream, err := c.ProcessOrders(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := stream.Send(&wrapperspb.StringValue{Value: "101"}); err != nil {
			panic(err)
		}

		if err := stream.Send(&wrapperspb.StringValue{Value: "102"}); err != nil {
			panic(err)
		}

		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
	}()

	for {
		combinedShipment, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Combined shipment : ", combinedShipment.OrderList)
	}
}

func orderStreamServerInterceptor1(srv interface{},
	ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	s := time.Now()
	err := handler(srv, ss)
	log.Printf("Method: %s, latency: %s\n", info.FullMethod, time.Since(s))

	return err
}

type wrappedServerStream struct {
	Recv []interface{}
	Send []interface{}
	grpc.ServerStream
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)
	w.Recv = append(w.Recv, w)
	return err
}

func (w *wrappedServerStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)
	w.Send = append(w.Send, m)
	return err
}

func newWrappedServerStream(s grpc.ServerStream) *wrappedServerStream {
	return &wrappedServerStream{
		make([]interface{}, 0),
		make([]interface{}, 0),
		s,
	}
}

func orderStreamServerInterceptor2(srv interface{},
	ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	s := time.Now()
	nss := newWrappedServerStream(ss)
	err := handler(srv, nss)
	log.Printf("Method: %s, req: %+v, resp: %+v, latency: %s\n",
		info.FullMethod, nss.Recv, nss.Send, time.Since(s))
	return err
}

func orderStreamClientInterceptor1(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s := time.Now()
	cs, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("method: %s, latency: %s\n", method, time.Since(s))
	return cs, err
}

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

func newWrappedStream(method string, s grpc.ClientStream) *wrappedClientStream {
	return &wrappedClientStream{
		method,
		s,
	}
}

func orderStreamClientInterceptor2(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {

	s := time.Now()
	cs, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("method: %s, latency: %s\n", method, time.Since(s))

	return newWrappedStream(method, cs), err
}

const (
	orderBatchSize = 3
)

var _ pb.OrderManagement4Server = &OrderManagement4Impl{}

type OrderManagement4Impl struct {
	pb.UnimplementedOrderManagement4Server
}

var orders4 = make(map[string]pb.Order4)

func (s *OrderManagement4Impl) ProcessOrders(stream pb.OrderManagement4_ProcessOrdersServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)
		if err == io.EOF {
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		destination := orders4[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := orders4[orderId.GetValue()]
			shipment.OrderList = append(shipment.OrderList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orders4[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := orders4[orderId.GetValue()]
			comShip.OrderList = append(shipment.OrderList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrderList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrderList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}

func init() {
	orders4["101"] = pb.Order4{
		Id:          "101",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}

	orders4["102"] = pb.Order4{
		Id:          "102",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}
}
