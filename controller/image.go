package controller

import (
	"path"
	"strings"

	"github.com/google/uuid"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/log"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-gonic/gin"
)

func UploadImageController(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" {
		log.Error("file is not allowed")
		util.NotAllowedResp(context)
		return
	}

	fileName := uuid.New().String() + fileExt
	filePath := path.Join("./data/image/" + fileName)
	if err := context.SaveUploadedFile(file, filePath); err != nil {
		log.Error(err)
		util.InternalErrResp(context)
		return
	}

	util.SuccessResp(context, gin.H{
		"url": global.Config.Serv.Host + "/image/show/" + fileName,
	})
}
