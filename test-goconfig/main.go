package main

import (
	"github.com/Unknwon/goconfig"
	"log"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %s", err)
	}

	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s", "key_default", nil)
	}

	log.Printf("%s > %s: %s", goconfig.DEFAULT_SECTION, "key_default", value)

	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key_default", "这是新的值")
	if err != nil {
		log.Fatalf("无法设置键值(%s): %s", "key_default", err)
	}
	log.Printf("设置键值 %s 为插入操作: %v", "key_default", isInsert)

	comment := cfg.GetSectionComments("super")
	log.Printf("分区 %s 的注释: %s", "super", comment)

	v := cfg.SetSectionComments("super", "# 这是新的分区注释")
	log.Printf("分区 %s 的注释被插入或删除: %v", "super", v)

	vInt, err := cfg.Int("must", "int")
	if err != nil {
		log.Fatalf("无法获取键值(%s): %s)", "int", err)
	}

	log.Printf("%s > %s: %v", "must", "int", vInt)

	vBool := cfg.MustBool("must", "bool")
	log.Printf("%s > %s: %v", "must", "bool", vBool)

	vBool = cfg.MustBool("must", "bool404")
	log.Printf("%s > %s: %v", "must", "bool", vBool)

	ok := cfg.DeleteKey("must", "string")
	log.Printf("删除键值 %s 是否成功: %v", "string", ok)

	err = goconfig.SaveConfigFile(cfg, "conf_save.ini")
	if err != nil {
		log.Fatalf("无法保存配置文件: %s", err)
	}
}
