package tool

import (
	"fmt"
	"time"

	"github.com/machsix/alist/v3/internal/conf"
	"github.com/machsix/alist/v3/internal/errs"
	"github.com/machsix/alist/v3/internal/setting"
	"github.com/machsix/alist/v3/internal/task"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/xhofe/tache"
)

type DownloadTask struct {
	task.TaskWithCreator
	Url               string       `json:"url"`
	DstDirPath        string       `json:"dst_dir_path"`
	TempDir           string       `json:"temp_dir"`
	DeletePolicy      DeletePolicy `json:"delete_policy"`
	Toolname          string       `json:"toolname"`
	Status            string       `json:"-"`
	Signal            chan int     `json:"-"`
	GID               string       `json:"-"`
	tool              Tool
	callStatusRetried int
}

func (t *DownloadTask) Run() error {
	if t.tool == nil {
		tool, err := Tools.Get(t.Toolname)
		if err != nil {
			return errors.WithMessage(err, "failed get tool")
		}
		t.tool = tool
	}
	if err := t.tool.Run(t); !errs.IsNotSupportError(err) {
		if err == nil {
			return t.Complete()
		}
		return err
	}
	t.Signal = make(chan int)
	defer func() {
		t.Signal = nil
	}()
	gid, err := t.tool.AddURL(&AddUrlArgs{
		Url:     t.Url,
		UID:     t.ID,
		TempDir: t.TempDir,
		Signal:  t.Signal,
	})
	if err != nil {
		return err
	}
	t.GID = gid
	var ok bool
outer:
	for {
		select {
		case <-t.CtxDone():
			err := t.tool.Remove(t)
			return err
		case <-t.Signal:
			ok, err = t.Update()
			if ok {
				break outer
			}
		case <-time.After(time.Second * 3):
			ok, err = t.Update()
			if ok {
				break outer
			}
		}
	}
	if err != nil {
		return err
	}
	if t.tool.Name() == "pikpak" {
		return nil
	}
	if t.tool.Name() == "115 Cloud" {
		// hack for 115
		<-time.After(time.Second * 1)
		err := t.tool.Remove(t)
		if err != nil {
			log.Errorln(err.Error())
		}
		return nil
	}
	t.Status = "offline download completed, maybe transferring"
	// hack for qBittorrent
	if t.tool.Name() == "qBittorrent" {
		seedTime := setting.GetInt(conf.QbittorrentSeedtime, 0)
		if seedTime >= 0 {
			t.Status = "offline download completed, waiting for seeding"
			<-time.After(time.Minute * time.Duration(seedTime))
			err := t.tool.Remove(t)
			if err != nil {
				log.Errorln(err.Error())
			}
		}
	}

	if t.tool.Name() == "transmission" {
		// hack for transmission
		seedTime := setting.GetInt(conf.TransmissionSeedtime, 0)
		if seedTime >= 0 {
			t.Status = "offline download completed, waiting for seeding"
			<-time.After(time.Minute * time.Duration(seedTime))
			err := t.tool.Remove(t)
			if err != nil {
				log.Errorln(err.Error())
			}
		}
	}
	return nil
}

// Update download status, return true if download completed
func (t *DownloadTask) Update() (bool, error) {
	info, err := t.tool.Status(t)
	if err != nil {
		t.callStatusRetried++
		log.Errorf("failed to get status of %s, retried %d times", t.ID, t.callStatusRetried)
		return false, nil
	}
	if t.callStatusRetried > 5 {
		return true, errors.Errorf("failed to get status of %s, retried %d times", t.ID, t.callStatusRetried)
	}
	t.callStatusRetried = 0
	t.SetProgress(info.Progress)
	t.Status = fmt.Sprintf("[%s]: %s", t.tool.Name(), info.Status)
	if info.NewGID != "" {
		log.Debugf("followen by: %+v", info.NewGID)
		t.GID = info.NewGID
		return false, nil
	}
	// if download completed
	if info.Completed {
		err := t.Complete()
		return true, errors.WithMessage(err, "failed to transfer file")
	}
	// if download failed
	if info.Err != nil {
		return true, errors.Errorf("failed to download %s, error: %s", t.ID, info.Err.Error())
	}
	return false, nil
}

func (t *DownloadTask) Complete() error {
	var (
		files []File
		err   error
	)
	if t.tool.Name() == "pikpak" {
		return nil
	}
	if t.tool.Name() == "115 Cloud" {
		return nil
	}
	if getFileser, ok := t.tool.(GetFileser); ok {
		files = getFileser.GetFiles(t)
	} else {
		files, err = GetFiles(t.TempDir)
		if err != nil {
			return errors.Wrapf(err, "failed to get files")
		}
	}
	// upload files
	for i := range files {
		file := files[i]
		TransferTaskManager.Add(&TransferTask{
			TaskWithCreator: task.TaskWithCreator{
				Creator: t.Creator,
			},
			file:         file,
			DstDirPath:   t.DstDirPath,
			TempDir:      t.TempDir,
			DeletePolicy: t.DeletePolicy,
			FileDir:      file.Path,
		})
	}
	return nil
}

func (t *DownloadTask) GetName() string {
	return fmt.Sprintf("download %s to (%s)", t.Url, t.DstDirPath)
}

func (t *DownloadTask) GetStatus() string {
	return t.Status
}

var DownloadTaskManager *tache.Manager[*DownloadTask]
