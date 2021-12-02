package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	envMap = map[string]string{
		"test": "测试",
		"prev": "预发",
		"prod": "生产",
	}
	docsDir = "./docs"
)

type swaggerJson struct {
	Info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
}

func main() {
	// 扫描docs中的文件
	fis, err := ioutil.ReadDir(docsDir)
	if err != nil {
		fmt.Println("当前目录不存在docs目录")
	}

	var data []map[string]string
	for _, fi := range fis {
		fn := fi.Name()
		if !fi.IsDir() && strings.HasSuffix(fn, ".json") {
			np := "未知"
			for env, name := range envMap {
				if strings.Contains(fn, env) {
					np = name
					break
				}
			}
			dir := fmt.Sprintf("%s/%s", docsDir, fn)
			file, err := os.Open(dir)
			if err != nil {
				break
			}

			all, err := ioutil.ReadAll(file)
			if err != nil {
				_ = file.Close()
				break
			} else {
				_ = file.Close()
			}

			var sj swaggerJson
			err = json.Unmarshal(all, &sj)
			if err != nil {
				break
			}

			data = append(data, map[string]string{
				"name": fmt.Sprintf("[%s]%s-%s", np, sj.Info.Title, sj.Info.Version),
				"url":  dir,
			})

		}
	}
	config, _ := json.Marshal(data)
	// 依次读取docs文件中的json，解析出title和文件名，生成map
	template, err := os.Open("./index-template.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer template.Close()
	all, err := ioutil.ReadAll(template)
	content := strings.Replace(string(all), "{config-placeholder}", string(config), 1)
	err = ioutil.WriteFile("./page/index.html", []byte(content), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s] 🐶 init success.\n", time.Now().Format("2006-01-02 15:04:05"))
	os.Exit(0)
}
