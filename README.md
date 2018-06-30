# azure-blob-uploader
[![Travis CI](https://travis-ci.org/johejo/azure-blob-uploader.svg?branch=master)](https://travis-ci.org/johejo/azure-blob-uploader)
[![Appveyor](https://ci.appveyor.com/api/projects/status/xbl1prn96p9kltfc/branch/master?svg=true
)](https://ci.appveyor.com/project/johejo/azure-blob-uploader)
[![CircleCI](https://circleci.com/gh/johejo/azure-blob-uploader.svg?style=svg)](https://circleci.com/gh/johejo/azure-blob-uploader)

A small CLI tool written in golang for uploading files to Microsoft Azure Blob Storage

## Feature

### Single binary
It does not depend on other components such as openssl and python like azure-cli.

### Small binary
For linux/amd64, the size is about 1.7 MB. (We're looking for ideas to further reduce the size.)

## Install

### Binally
Binally can be download from [release page](https://github.com/johejo/azure-blob-uploader/releases).

### go get
```
$ go get github.com/johejo/azure-blob-uploader
```

### Build from source
```
$ git clone https://github.com/johejo/azure-blob-uploader.git
$ cd azure-blob-uploader/
$ go build
```

## Usage
```
$ ./azure-blob-uploader -accout-name <ACCOUNT_NAME> -account-key <ACCOUNT_KEY> -c <CONTAINER_NAME> -f <FILE>
```

## Story
I was ordered to upload files from a very old machine to Azure Blob Storage at a certain legacy project.  
It was impossible to install azure-cli on that machine.  
As a result I found an article by Toru Makabe and made an upload tool.  
Because there may be people who are in trouble with other similar things, leave my deliverables here.  
The executable file was compressed with GNU strip and UPX.

## Great appreciation
Toru Makabe http://torumakabe.github.io/post/azblob_golang/
