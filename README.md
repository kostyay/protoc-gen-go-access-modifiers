protoc-gen-go-private
======================

This is a proof of concept plugin for protoc that generates Go code.
It allows defining message fields as private and generates a `AsPublic()` method that 
empties the private fields into a public struct.

It is provided as a part of my blog post on gRPC API gateways.

This is not the final version of the plugin, but it is a good starting point for getting the idea how to use it.

In addition to the plugin you will find a gRPC client middleware strips private fields off incoming client responses.


# Usage
```bash
    go install github.com/kostyay/protoc-gen-go-private@latest
```

Add to `buf.gen.yaml`:
```yaml
    plugins:
      - name: protoc-gen-go-private
        out: .
        opt: paths=source_relative
```