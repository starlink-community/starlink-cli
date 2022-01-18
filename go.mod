module github.com/starlink-community/starlink-cli

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/starlink-community/starlink-grpc-go v0.0.0-20210211202449-2e89f3d7e309
	google.golang.org/grpc v1.43.0
)

// TODO(@cpu): Once the starlink-grpc-go PR updating the proto definitions[0]
//             lands this hack can be removed..
//             [0] https://github.com/starlink-community/starlink-grpc-go/pull/1
replace github.com/starlink-community/starlink-grpc-go v0.0.0-20210211202449-2e89f3d7e309 => github.com/cpu/starlink-grpc-go v0.0.0-20220118002722-5b7e82ec264a
