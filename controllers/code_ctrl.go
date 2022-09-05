package controllers

import (
	"CodeSheep-runcode/configs"
	"CodeSheep-runcode/middles"
	"CodeSheep-runcode/models"
	"CodeSheep-runcode/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type CodeController struct {
}

func (controller CodeController) HandleRunCode(c *gin.Context) {
	// get code and id
	var code models.Code
	bindCodeErr := c.ShouldBindJSON(&code)
	if bindCodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400})
	}
	id := middles.GetSessionId(c)

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
		log.Println(sh, "cmd")
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
	} else if !utils.IsNeedsCompile(code.Code) {
		dealGoodCode(pathName, prefix, code.Language, "", c)
		return
	}

}

// deal with io err
func dealErr(tip string, err error, c *gin.Context) {
	log.Println(tip, err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": configs.WrongCode,
		"msg":  configs.WrongMsg,
		"res":  "",
	})
}

// deal with compile err
func dealWrongCode(path, prefix, lang string, c *gin.Context) {
	errMsg, errErr := ioutil.ReadFile(path + "/" + prefix + ".err")
	if errErr != nil {
		dealErr("Error when read err file: ", errErr, c)
	}
	// long time err happen, only when script language
	if string(errMsg) == configs.LongInfo {
		dealGoodCode(path, prefix, lang, string(errMsg), c)
	}

	//TODO delete file
	c.JSON(http.StatusOK, gin.H{
		"code": configs.NopCode,
		"msg":  configs.NopMsg,
		"res":  string(errMsg),
	})
}

// deal with compile right
func dealGoodCode(path, prefix, lang, errinfo string, c *gin.Context) {
	comErrPath := path + "/" + prefix + ".err"
	comOutPath := path + "/" + prefix + ".out"
	comInfoPath := path + "/" + prefix + ".info"
	comCodePath := path + "/" + prefix + utils.LangMapSuffix(lang)

	log.Println(comErrPath)
	log.Println(comOutPath)
	log.Println(comInfoPath)
	log.Println(comCodePath)

	var outContent string
	if ok, _ := utils.FileExists(comOutPath); ok {
		outFile, openErr := os.Open(comOutPath)
		if openErr != nil {
			dealErr("Error when read out: ", openErr, c)
		}
		outInfo, _ := outFile.Stat()
		out, _ := utils.FileReadN(outFile,
			utils.Min(outInfo.Size(), configs.ContentMaxSize))

		outContent = string(out)
	}

	var result gin.H
	result = gin.H{
		"code": configs.GoodCode,
		"msg":  configs.GoodMsg,
		"res":  outContent,
	}
	if ok, _ := utils.FileExists(comInfoPath); ok {
		info, infoErr := ioutil.ReadFile(comInfoPath)
		if infoErr != nil {
			dealErr("Error when read out: ", infoErr, c)
		}
		outContent += string(info)

		if errinfo == configs.LongInfo {
			//TODO email
			result = gin.H{
				"code": configs.WrongCode,
				"msg":  configs.WrongMsg,
				"res":  outContent,
			}
		} else {
			result = gin.H{
				"code": configs.LongCode,
				"msg":  configs.LongMsg,
				"res":  outContent,
			}
		}
		// TODO rm code shit

	}
	c.JSON(http.StatusOK, result)

}
