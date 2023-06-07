package server

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *fileTransferServer) UploadFile(stream filetransfer.FileTransferService_UploadFileServer) error {
	var content []byte
	var filename string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 文件传输完成
			return stream.SendAndClose(&empty.Empty{})
		}
		if err != nil {
			return err
		}
		content = append(content, req.Content...)
		filename = req.Filename
	}
	// 保存文件
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
