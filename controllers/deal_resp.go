package controllers

import (
	"CodeSheep-runcode/configs"
	"CodeSheep-runcode/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	errPath := path + "/" + prefix + ".err"
	outPath := path + "/" + prefix + ".out"
	codePath := path + "/" + prefix + utils.LangMapSuffix(lang)
	errByte, errErr := ioutil.ReadFile(errPath)
	errMsg := string(errByte)
	if errErr != nil {
		dealErr("Error when read err file: ", errErr, c)
	}

	// remove id in the err info
	errMsg = strings.ReplaceAll(errMsg, id, "")

	// long time err happen, only when script language
	if errMsg == configs.LongInfo {
		dealGoodCode(path, prefix, lang, errMsg, c)
	}

	go utils.DeleteCodes(errPath, codePath, outPath)

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
	var comExePath string
	if utils.IsNeedsCompile(lang) {
		comExePath = path + "/" + prefix
	}

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
			go utils.SendBugEmail(prefix)
			result = gin.H{
				"code": configs.WrongCode,
				"msg":  configs.WrongMsg,
				"res":  outContent,
			}
			comCodePath = ""
		} else {
			result = gin.H{
				"code": configs.LongCode,
				"msg":  configs.LongMsg,
				"res":  outContent,
			}
		}
	} else {
		comInfoPath = ""
	}
	go utils.DeleteCodes(comOutPath, comCodePath, comInfoPath, comErrPath, comExePath)

	c.JSON(http.StatusOK, result)

}
