package server

import (
	"github.com/riversy/file-storage-server/service"
	"golang.org/x/net/context"
)

type fileService struct {
	storageFolder      string
	maxSize            int
	concurrentUploads  int
	concurrentRequests int
	service.FilesServiceServer
}

// PutFile function implements a file receiving functionality.
// It validates the file contents, mime, etc., and stores the file into the storage folder.
// As a result, it returns the file upload status.
func (s *fileService) PutFile(context.Context, *service.File) (*service.FileUploadStatus, error) {
	return nil, nil
}

// GetAllFilesMeta function opens a Stream that sends the client the meta-information about all
// previously uploaded files. Later it sends the information about all newly uploaded files.
// This stream is expected to be closed by the client or if the server exits.
func (s *fileService) GetAllFilesMeta(*service.FilesRequest, service.FilesService_GetAllFilesMetaServer) error {
	return nil
}

// GetAllFilesMetaSync function sends to the client the meta-information about all
// previously uploaded files. After the information is sent, the connection will be closed by Server.
func (s *fileService) GetAllFilesMetaSync(context.Context, *service.FilesRequest) (*service.FileMetaList, error) {
	return nil, nil
}
