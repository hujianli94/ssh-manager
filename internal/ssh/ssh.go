package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"ssh-manager/internal/config"
)

// AddAlias 添加 SSH 别名
func AddAlias(cfg config.SSHConfig) error {
	// 构建 SSH 配置内容
	configStr := fmt.Sprintf(`
Host %s
    HostName %s
    User %s
    Port %d
`, cfg.Host, cfg.HostName, cfg.User, cfg.Port)

	// 将配置写入 SSH 配置文件 (~/.ssh/config)
	sshConfigPath := os.Getenv("HOME") + "/.ssh/config"
	f, err := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open SSH config file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(configStr); err != nil {
		return fmt.Errorf("failed to write to SSH config file: %w", err)
	}

	return nil
}

// ListAliases 列出所有 SSH 别名
func ListAliases() ([]string, error) {
	sshConfigPath := os.Getenv("HOME") + "/.ssh/config"
	content, err := ioutil.ReadFile(sshConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SSH config file: %w", err)
	}

	// 使用正则表达式匹配 Host 字段
	re := regexp.MustCompile(`Host\s+([^\s]+)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	var aliases []string
	for _, match := range matches {
		aliases = append(aliases, match[1]) // 提取 Host 字段的值
	}

	return aliases, nil
}

// RemoveAlias 删除 SSH 别名
func RemoveAlias(alias string) error {
	sshConfigPath := os.Getenv("HOME") + "/.ssh/config"
	content, err := os.ReadFile(sshConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read SSH config file: %w", err)
	}

	// Use a more robust regular expression to match the Host block
	re := regexp.MustCompile(fmt.Sprintf(`(?m)^\s*Host\s+%s(\s+.*)?\s*$([^\n]*\n)*`, regexp.QuoteMeta(alias)))
	newContent := re.ReplaceAllString(string(content), "")

	if string(content) == newContent {
		return fmt.Errorf("alias '%s' not found", alias)
	}

	err = os.WriteFile(sshConfigPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to SSH config file: %w", err)
	}

	return nil
}
