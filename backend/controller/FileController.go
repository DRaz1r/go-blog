package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "格式错误",
		})
		return
	}
	fileName := header.Filename
	ext := path.Ext(fileName)
	newFileName := "image_" + time.Now().Format("20060102150405") + ext
	out, err := os.Create("static/images/" + newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建文件失败",
		})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "保存文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": gin.H{
			"filePath": "/images/" + newFileName,
		},
	})
	return
}

// RichEditorUpload 上传富文本编辑器中的图像
func RichEditorUpload(c *gin.Context) {
	fromData, _ := c.MultipartForm()
	files := fromData.File["wangeditor-uploaded-image"]
	var url []string
	for _, file := range files {
		ext := path.Ext(file.Filename)
		name := "image_" + time.Now().Format("20060102150405")
		newFilename := name + ext
		dst := path.Join("./static/images", newFilename)
		fileurl := "/images/" + newFilename
		url = append(url, fileurl)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno":   1,
				"message": "上传失败",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errno": 0,
		"data": gin.H{
			"url": url[0],
		},
	})
}
