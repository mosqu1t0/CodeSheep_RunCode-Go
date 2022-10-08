package controllers

import (
	"CodeSheep-runcode/configs"
	"CodeSheep-runcode/middles"
	"CodeSheep-runcode/models"
	"CodeSheep-runcode/utils"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type CodeController struct {
}

var id string

func (controller CodeController) HandleRunCode(c *gin.Context) {
	// get code and id
	var code models.Code
	bindCodeErr := c.ShouldBindJSON(&code)
	if bindCodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400})
	}
	id = middles.GetSessionId(c)

	//build the path and codefile
	pathName := configs.WorkPath + utils.LangMapPath(code.Language)
	prefix := utils.LangMapPath(code.Language) + id

	if _, mkdirErr := utils.FoldExists(pathName); mkdirErr != nil {
		dealErr("Error when mkdir: ", mkdirErr, c)
	}
	if chdirErr := os.Chdir(pathName); chdirErr != nil {
		dealErr("Error when chdir: ", chdirErr, c)
	}

	comCode, saveCodeErr := os.Create(prefix + utils.LangMapSuffix(code.Language))
	if saveCodeErr != nil {
		dealErr("Error when save code: ", saveCodeErr, c)
	}
	comCode.WriteString(code.Code)
	comCode.Close()

	//first time run sh, compile and run
	var cmd *exec.Cmd
	sh := configs.WorkPath + utils.LangMapFirSh(code.Language)
	if !utils.IsNeedsCompile(code.Language) {
		cmd = exec.Command(sh, code.Input, id)
	} else {
		cmd = exec.Command(sh, id)
	}
	cmd.Run()

	//check if compile err
	comErrPath := pathName + "/" + prefix + ".err"
	ok, errErr := utils.FileExists(comErrPath)
	if errErr != nil {
		dealErr("Error when read comErr: ", errErr, c)
		return
	}
	if ok {
		dealWrongCode(pathName, prefix, code.Language, c)
		return
	} else if !utils.IsNeedsCompile(code.Language) {
		dealGoodCode(pathName, prefix, code.Language, "", c)
		return
	}

	// second time run sh, run o
	sh = configs.WorkPath + utils.LangMapSecSh(code.Language)
	cmd = exec.Command(sh, code.Input, id)
	cmd.Run()

	exePath := pathName + "/" + prefix
	if ok, exeErr := utils.FileExists(exePath); ok {
		dealGoodCode(pathName, prefix, code.Language, "", c)
	} else {
		dealErr("Error when exec bin: ", exeErr, c)
	}
}
