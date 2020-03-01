package main

type DownloadStrategy int

const (
	LocalFileDownloadStrategy DownloadStrategy = iota
	AWSDownloadStrategy       DownloadStrategy = iota
)

// Environment variables declaration
const FileDownloadStrategy = 0
const DefaultFilePath = "downloads"
const AWSBucketPath = ""
