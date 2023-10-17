package env

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type env struct {
	vp *viper.Viper
}

// LoadingEnv 创建一个读取配置的对象
//
// filename 文件名，包括扩展名, 如：config.toml;
//
// paths 多配所在的目录集合，如：configs/aaa/config.toml ...;
//
// target 将结果映射到target， 结构体指针;
func LoadingEnv(filename string, target interface{}, paths ...string) (*env, error) {
	vp := viper.New()
	for _, path := range paths {
		if path != "" {
			vp.AddConfigPath(path)
		}
	}

	ext := filepath.Ext(filename)
	parts := strings.Split(ext, ".")
	if len(parts) != 2 {
		return nil, errors.New("filename invalid")
	}

	// 设置配置文件名
	vp.SetConfigName(filepath.Base(filename))
	// 设置配置扩展名
	vp.SetConfigType(parts[1])
	// 加载使用环境变量
	vp.AutomaticEnv()

	// 读取配置进来
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &env{
		vp: vp,
	}

	err = cfg.vp.Unmarshal(target)
	if err != nil {
		return nil, err
	}

	// 监听
	cfg.Watch(target)

	return cfg, nil
}

// Watch 监听 配置是否发生变化，更改了就立马重载
func (cfg *env) Watch(target interface{}) {
	go func() {
		cfg.vp.WatchConfig()
		cfg.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = cfg.vp.Unmarshal(target)
		})
	}()
}
