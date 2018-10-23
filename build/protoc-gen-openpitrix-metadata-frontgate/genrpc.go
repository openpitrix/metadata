// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Protoc plugin for OpenPitrix metadata/frontgate.
package main

import (
	"bytes"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

const genrpcPluginName = "openpitrix-metadata-frontgate"

// genrpcPlugin produce the Service interface.
type genrpcPlugin struct {
	*generator.Generator
}

// Name returns the name of the plugin.
func (p *genrpcPlugin) Name() string { return genrpcPluginName }

// Init is called once after data structures are built but before
// code generation begins.
func (p *genrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

// Generate produces the code generated by the plugin for this file.
func (p *genrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.P(`import "bufio"`)
		p.P(`import "crypto/tls"`)
		p.P(`import "errors"`)
		p.P(`import "io"`)
		p.P(`import "log"`)
		p.P(`import "net"`)
		p.P(`import "net/http"`)
		p.P(`import "net/rpc"`)
		p.P(`import "time"`)
	}
}

// Generate generates the Service interface.
// rpc service can't handle other proto message!!!
func (p *genrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceInterface(file, svc)
		p.genServiceServer(file, svc)
		p.genServiceClient(file, svc)
	}
}

func (p *genrpcPlugin) genServiceInterface(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const serviceInterfaceTmpl = `
type {{.ServiceName}} interface {
	{{.CallMethodList}}
}
`
	const callMethodTmpl = `
{{.MethodName}}(in *{{.ArgsType}}, out *{{.ReplyType}}) error`

	// gen call method list
	var callMethodList string
	for _, m := range svc.Method {
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(callMethodTmpl))
		t.Execute(out, &struct{ ServiceName, MethodName, ArgsType, ReplyType string }{
			ServiceName: generator.CamelCase(svc.GetName()),
			MethodName:  generator.CamelCase(m.GetName()),
			ArgsType:    p.TypeName(p.ObjectNamed(m.GetInputType())),
			ReplyType:   p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
		callMethodList += out.String()

		p.RecordTypeUse(m.GetInputType())
		p.RecordTypeUse(m.GetOutputType())
	}

	// gen all interface code
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(serviceInterfaceTmpl))
		t.Execute(out, &struct{ ServiceName, CallMethodList string }{
			ServiceName:    generator.CamelCase(svc.GetName()),
			CallMethodList: callMethodList,
		})
		p.P(out.String())
	}
}

func (p *genrpcPlugin) genServiceServer(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const serviceHelperFunTmpl = `
// Accept{{.ServiceName}}Client accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func Accept{{.ServiceName}}Client(lis net.Listener, x {{.ServiceName}}) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("{{.ServiceRegisterName}}", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// Register{{.ServiceName}} publish the given {{.ServiceName}} implementation on the server.
func Register{{.ServiceName}}(srv *rpc.Server, x {{.ServiceName}}) error {
	if err := srv.RegisterName("{{.ServiceRegisterName}}", x); err != nil {
		return err
	}
	return nil
}

// New{{.ServiceName}}Server returns a new {{.ServiceName}} Server.
func New{{.ServiceName}}Server(x {{.ServiceName}}) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("{{.ServiceRegisterName}}", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServe{{.ServiceName}} listen announces on the local network address laddr
// and serves the given {{.ServiceName}} implementation.
func ListenAndServe{{.ServiceName}}(network, addr string, x {{.ServiceName}}) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("{{.ServiceRegisterName}}", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// Serve{{.ServiceName}} serves the given {{.ServiceName}} implementation.
func Serve{{.ServiceName}}(conn io.ReadWriteCloser, x {{.ServiceName}}) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("{{.ServiceRegisterName}}", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeConn(conn)
}
`
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(serviceHelperFunTmpl))
		t.Execute(out, &struct{ PackageName, ServiceName, ServiceRegisterName string }{
			PackageName: file.GetPackage(),
			ServiceName: generator.CamelCase(svc.GetName()),
			ServiceRegisterName: p.makeServiceRegisterName(
				file, file.GetPackage(), generator.CamelCase(svc.GetName()),
			),
		})
		p.P(out.String())
	}
}

func (p *genrpcPlugin) genServiceClient(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const clientHelperFuncTmpl = `
type {{.ServiceName}}Client struct {
	*rpc.Client
}

// New{{.ServiceName}}Client returns a {{.ServiceName}} stub to handle
// requests to the set of {{.ServiceName}} at the other end of the connection.
func New{{.ServiceName}}Client(conn io.ReadWriteCloser) (*{{.ServiceName}}Client) {
	c := rpc.NewClient(conn)
	return &{{.ServiceName}}Client{c}
}

{{.MethodList}}

// Dial{{.ServiceName}} connects to an {{.ServiceName}} at the specified network address.
func Dial{{.ServiceName}}(network, addr string) (*{{.ServiceName}}Client, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{c}, nil
}

// Dial{{.ServiceName}}Timeout connects to an {{.ServiceName}} at the specified network address.
func Dial{{.ServiceName}}Timeout(network, addr string, timeout time.Duration) (*{{.ServiceName}}Client, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{rpc.NewClient(conn)}, nil
}

// Dial{{.ServiceName}}HTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
func Dial{{.ServiceName}}HTTP(network, address string) (*{{.ServiceName}}Client, error) {
	return Dial{{.ServiceName}}HTTPPath(network, address, rpc.DefaultRPCPath)
}

// Dial{{.ServiceName}}HTTPPath connects to an HTTP RPC server
// at the specified network address and path.
func Dial{{.ServiceName}}HTTPPath(network, address, path string) (*{{.ServiceName}}Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return dial{{.ServiceName}}Path(network, address, path, conn)
}

// Dial{{.ServiceName}}HTTPS connects to an HTTPS RPC server at the specified network address
// listening on the default HTTP RPC path.
func Dial{{.ServiceName}}HTTPS(network, address string, tlsConfig *tls.Config) (*{{.ServiceName}}Client, error) {
	return Dial{{.ServiceName}}HTTPSPath(network, address, rpc.DefaultRPCPath, tlsConfig)
}

// Dial{{.ServiceName}}HTTPSPath connects to an HTTPS RPC server
// at the specified network address and path.
func Dial{{.ServiceName}}HTTPSPath(network, address, path string, tlsConfig *tls.Config) (*{{.ServiceName}}Client, error) {
	conn, err := tls.Dial(network, address, tlsConfig)
	if err != nil {
		return nil, err
	}
	return dial{{.ServiceName}}Path(network, address, path, conn)
}

func dial{{.ServiceName}}Path(network, address, path string, conn net.Conn) (*{{.ServiceName}}Client, error) {
	const net_rpc_connected = "200 Connected to Go RPC"

	io.WriteString(conn, "CONNECT "+path+" HTTP/1.0\n\n")

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
	if err == nil && resp.Status == net_rpc_connected {
		return &{{.ServiceName}}Client{rpc.NewClient(conn)}, nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	conn.Close()
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}
`
	const clientMethodTmpl = `
func (c *{{.ServiceName}}Client) {{.MethodName}}(in *{{.ArgsType}}) (out *{{.ReplyType}}, err error) {
	if in == nil {
		in = new({{.ArgsType}})
	}
	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}
	out = new({{.ReplyType}})
	if err = c.Call("{{.ServiceRegisterName}}.{{.MethodName}}", in, out); err != nil {
		return nil, err
	}
	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}
	return out, nil
}

func (c *{{.ServiceName}}Client) Async{{.MethodName}}(in *{{.ArgsType}}, out *{{.ReplyType}}, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new({{.ArgsType}})
	}
	return c.Go(
		"{{.ServiceRegisterName}}.{{.MethodName}}",
		in, out,
		done,
	)
}
`

	// gen client method list
	var methodList string
	for _, m := range svc.Method {
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(clientMethodTmpl))
		t.Execute(out, &struct{ ServiceName, ServiceRegisterName, MethodName, ArgsType, ReplyType string }{
			ServiceName: generator.CamelCase(svc.GetName()),
			ServiceRegisterName: p.makeServiceRegisterName(
				file, file.GetPackage(), generator.CamelCase(svc.GetName()),
			),
			MethodName: generator.CamelCase(m.GetName()),
			ArgsType:   p.TypeName(p.ObjectNamed(m.GetInputType())),
			ReplyType:  p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
		methodList += out.String()
	}

	// gen all client code
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(clientHelperFuncTmpl))
		t.Execute(out, &struct{ PackageName, ServiceName, MethodList string }{
			PackageName: file.GetPackage(),
			ServiceName: generator.CamelCase(svc.GetName()),
			MethodList:  methodList,
		})
		p.P(out.String())
	}
}

func (p *genrpcPlugin) makeServiceRegisterName(
	file *generator.FileDescriptor,
	packageName, serviceName string,
) string {
	return packageName + "." + serviceName
}

func init() {
	generator.RegisterPlugin(new(genrpcPlugin))
}
