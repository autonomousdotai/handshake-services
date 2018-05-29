package service

import (
	"github.com/autonomousdotai/handshake-services/comment-service/dao"
	"github.com/autonomousdotai/handshake-services/comment-service/utils"
)

var fileUploadService = utils.GSService{}
// service
var commentDao = dao.CommentDao{}
// template
var netUtil = utils.NetUtil{}
