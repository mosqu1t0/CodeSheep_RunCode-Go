package utils

var (
	path   = map[string]string{}
	suffix = map[string]string{}
	firSh  = map[string]string{}
	secSh  = map[string]string{}
)

func LangMapPath(lang string) string {
	return path[lang]
}

func LangMapSuffix(lang string) string {
	return suffix[lang]
}

func LangMapFirSh(lang string) string {
	return firSh[lang]
}

func LangMapSecSh(lang string) string {
	return secSh[lang]
}

func IsNeedsCompile(lang string) bool {
	_, ok := secSh[lang]
	return ok
}

func init() {

	path["c"] = "C"
	path["cpp"] = "Cpp"
	path["rust"] = "Rust"
	path["golang"] = "Go"
	path["python"] = "Py"
	path["javascript"] = "Js"

	suffix["c"] = ".c"
	suffix["cpp"] = ".cpp"
	suffix["rust"] = ".rs"
	suffix["golang"] = ".go"
	suffix["python"] = ".py"
	suffix["javascript"] = ".js"

	firSh["c"] = "compileC.sh"
	firSh["cpp"] = "compileCpp.sh"
	firSh["rust"] = "compileRust.sh"
	firSh["golang"] = "compileGo.sh"
	firSh["python"] = "runPy.sh"
	firSh["javascript"] = "runJs.sh"

	secSh["c"] = "runC.sh"
	secSh["cpp"] = "runCpp.sh"
	secSh["rust"] = "runRust.sh"
	secSh["golang"] = "runGo.sh"
}
