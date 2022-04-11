module get.porter.sh/example-bundles

go 1.18

// ensuremixins
replace get.porter.sh/magefiles => github.com/carolynvs/magefiles v0.1.3-0.20220411141833-0c2a892e6ff1

require (
	get.porter.sh/magefiles v0.1.3
	github.com/carolynvs/magex v0.8.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/magefile/mage v1.11.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/andybalholm/brotli v1.0.0 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/golang/snappy v0.0.4-0.20210608040537-544b4180ac70 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/klauspost/pgzip v1.2.4 // indirect
	github.com/mholt/archiver/v3 v3.5.0 // indirect
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/pierrec/lz4/v4 v4.0.3 // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
)
