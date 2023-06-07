package client

import (
	"context"
	"io"
	"os"
)

func uploadFile(filename string, client filetransfer.FileTransferServiceClient) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		req := &filetransfer.FileRequest{
			Filename: filename,
			Content:  buf[:n],
		}
		err = stream.Send(req)
		if err != nil {
			return err
		}
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}
