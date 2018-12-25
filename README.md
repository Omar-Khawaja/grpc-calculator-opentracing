Simple calculator API implemented using gRPC for testing purposes. The server
and client have been written in Go, but you may generate client and server
interfaces from calculator.proto file in the calculator directory for any of the supported languages. See https://grpc.io/docs/tutorials/ for more details.

This code has been instrumented with OpenTracing. See https://opentracing.io/
for more details. Original code without instrumentation can be found here: https://github.com/Omar-Khawaja/grpc-calculator

Run **make** to install dependencies and set up the Jaeger trace collector as a docker container. You can then go to localhost:16686 to see your trace after you run the server and client code.
