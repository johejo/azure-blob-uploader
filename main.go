// Thanks to Toru Makabe. http://torumakabe.github.io/post/azblob_golang/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

  "github.com/Azure/azure-storage-blob-go/azblob"
)

var (
	accountName    string
	accountKey     string
	containerName  string
	fileName       string
	blockSize      int64
	blockSizeBytes int64
	parallelism    int64
)

func init() {
	flag.StringVar(&accountName, "account-name", "", "(Required) Storage Account Name")
	flag.StringVar(&accountKey, "account-key", "", "(Required) Storage Account Key")
	flag.StringVar(&containerName, "c", "", "(Required - short option) Blob Container Name")
	flag.StringVar(&containerName, "container-name", "", "(Required) Blob Container Name")
	flag.StringVar(&fileName, "f", "", "(Required - short option) Upload filename")
	flag.StringVar(&fileName, "file", "", "(Required) Upload filename")
	flag.Int64Var(&blockSize, "b", 4, "(Optional - short option) Blob Blocksize (MB) - From 1 to 100. Max filesize depends on this value. Max filesize = Blocksize * 50,000 blocks")
	flag.Int64Var(&blockSize, "blocksize", 4, "(Optional) Blob Blocksize (MB) - From 1 to 100. Max filesize depends on this value. Max filesize = Blocksize * 50,000 blocks")
	flag.Int64Var(&parallelism, "p", 5, "(Optional - short option) Parallelism - From 0 to 32. Default 5.")
	flag.Int64Var(&parallelism, "parallelism", 5, "(Optional) Parallelism - From 0 to 32. Default 5.")
	flag.Parse()

	checkOpt("Blocksize", blockSize, 1, 100)
	blockSizeBytes = blockSize * 1024 * 1024
	checkOpt("Parallelism", parallelism, 0, 32)
}

func checkOpt(optName string, opt int64, min int64, max int64) {
	if opt < min || max < opt {
		log.Fatalf("%v must be from %v to %v", optName, min, max)
	}
}

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok {
			switch serr.ServiceCode() {
			case azblob.ServiceCodeContainerAlreadyExists:
				log.Println("Received 409. Container already exists")
				return
			}
		}
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open(fileName)
	handleErrors(err)
	defer file.Close()

	handleErrors(err)
	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, fileName))
  credential, _ := azblob.NewSharedKeyCredential(accountName, accountKey)
	blockBlobURL := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background()

	log.Printf("Uploading %v to %v", fileName, containerName)
	response, err := azblob.UploadFileToBlockBlob(ctx, file, blockBlobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   blockSizeBytes,
			Parallelism: uint16(parallelism),
		})
	handleErrors(err)
	_ = response

	log.Println("Done")
}
