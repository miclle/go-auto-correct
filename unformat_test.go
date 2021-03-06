package autocorrect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnformat(t *testing.T) {
	assert.Equal(t, " Hello world ", Unformat(" Hello world "))
	assert.Equal(t, "100中文", Unformat("100 中文"))
	assert.Equal(t, "中文100", Unformat("中文 100"))

	raw := "据港交所最新权益披露资料显示，2019 年 12 月 27 日，三生制药获 JP Morgan Chase & Co.每股均价 9.582 港元，增持 270.3 万股，总价约 2590 万港元。"
	expected := "据港交所最新权益披露资料显示，2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股，总价约2590万港元。"
	assert.Equal(t, expected, Unformat(raw))
}
