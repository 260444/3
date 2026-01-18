package utils

import (
	"github.com/mojocn/base64Captcha"
)

// CaptchaStore 自定义验证码存储
type CaptchaStore struct {
	// 这里可以集成Redis或其他存储方式
	// 为了演示，我们使用内存存储
	store map[string]string
}

func NewCaptchaStore() *CaptchaStore {
	return &CaptchaStore{
		store: make(map[string]string),
	}
}

// Set 设置验证码
func (c *CaptchaStore) Set(id string, value string) {
	c.store[id] = value
}

// Get 获取验证码
func (c *CaptchaStore) Get(id string, clear bool) string {
	value := c.store[id]
	if clear {
		delete(c.store, id)
	}
	return value
}

// Verify 验证验证码
func (c *CaptchaStore) Verify(id, answer string, clear bool) bool {
	stored := c.Get(id, clear)
	return stored == answer
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha() (id, b64s string, err error) {
	// 配置验证码参数 - 使用 v1.3.8 版本的 ConfigCharacter 结构体
	var config = base64Captcha.DriverDigit{
		Height:   80,  // 画布高
		Width:    240, // 画布宽
		Length:   5,   // 字符个数
		MaxSkew:  0.7, // 最大倾斜因子
		DotCount: 80,  // 噪点数量
	}

	// 创建验证码对象
	captcha := base64Captcha.NewCaptcha(
		&config,
		base64Captcha.DefaultMemStore,
	)

	// 生成验证码
	id, b64s, answer, err := captcha.Generate()

	// 将答案存储到自定义存储中
	NewCaptchaStore().Set(id, answer)

	return id, b64s, nil
}
