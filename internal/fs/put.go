package fs

import (
	"context"
	"fmt"
	"github.com/machsix/alist/v3/internal/driver"
	"github.com/machsix/alist/v3/internal/errs"
	"github.com/machsix/alist/v3/internal/model"
	"github.com/machsix/alist/v3/internal/op"
	"github.com/machsix/alist/v3/internal/task"
	"github.com/pkg/errors"
	"github.com/xhofe/tache"
)

type UploadTask struct {
	task.TaskWithCreator
	storage          driver.Driver
	dstDirActualPath string
	file             model.FileStreamer
}

func (t *UploadTask) GetName() string {
	return fmt.Sprintf("upload %s to [%s](%s)", t.file.GetName(), t.storage.GetStorage().MountPath, t.dstDirActualPath)
}

func (t *UploadTask) GetStatus() string {
	return "uploading"
}

func (t *UploadTask) Run() error {
	return op.Put(t.Ctx(), t.storage, t.dstDirActualPath, t.file, t.SetProgress, true)
}

var UploadTaskManager *tache.Manager[*UploadTask]

// putAsTask add as a put task and return immediately
func putAsTask(ctx context.Context, dstDirPath string, file model.FileStreamer) (task.TaskInfoWithCreator, error) {
	storage, dstDirActualPath, err := op.GetStorageAndActualPath(dstDirPath)
	if err != nil {
		return nil, errors.WithMessage(err, "failed get storage")
	}
	if storage.Config().NoUpload {
		return nil, errors.WithStack(errs.UploadNotSupported)
	}
	if file.NeedStore() {
		_, err := file.CacheFullInTempFile()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create temp file")
		}
		//file.SetReader(tempFile)
		//file.SetTmpFile(tempFile)
	}
	taskCreator, _ := ctx.Value("user").(*model.User) // taskCreator is nil when convert failed
	t := &UploadTask{
		TaskWithCreator: task.TaskWithCreator{
			Creator: taskCreator,
		},
		storage:          storage,
		dstDirActualPath: dstDirActualPath,
		file:             file,
	}
	UploadTaskManager.Add(t)
	return t, nil
}

// putDirect put the file and return after finish
func putDirectly(ctx context.Context, dstDirPath string, file model.FileStreamer, lazyCache ...bool) error {
	storage, dstDirActualPath, err := op.GetStorageAndActualPath(dstDirPath)
	if err != nil {
		return errors.WithMessage(err, "failed get storage")
	}
	if storage.Config().NoUpload {
		return errors.WithStack(errs.UploadNotSupported)
	}
	return op.Put(ctx, storage, dstDirActualPath, file, nil, lazyCache...)
}
