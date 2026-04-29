#!/bin/bash
# ============================================================================
# check_deps.sh - kit/go.mod 禁止依赖检查脚本
#
# 用途：验证 kit/go.mod 不包含禁止依赖（kratos、gorm、redis 等重型框架）
# 用法：可从项目根目录或 kit/ 目录运行
#   - 从项目根目录：bash kit/scripts/check_deps.sh
#   - 从 kit 目录：bash scripts/check_deps.sh
# ============================================================================

set -e

# 禁止依赖列表
FORBIDDEN_DEPS=(
    "github.com/go-kratos/kratos"
    "gorm.io/gorm"
    "github.com/redis/go-redis"
    "go.mongodb.org/mongo-driver"
    "github.com/ThreeDotsLabs/watermill"
    "github.com/hashicorp/consul"
    "go.etcd.io/etcd"
    "go.opentelemetry.io/otel"
)

# 确定 kit/go.mod 的路径（支持从项目根目录或 kit/ 目录运行）
GOMOD_PATH=""
if [ -f "kit/go.mod" ]; then
    GOMOD_PATH="kit/go.mod"
elif [ -f "go.mod" ] && grep -q "module github.com/ikaiguang/go-srv-kit/kit" go.mod 2>/dev/null; then
    GOMOD_PATH="go.mod"
else
    echo "❌ 错误：找不到 kit/go.mod 文件"
    echo "   请从项目根目录或 kit/ 目录运行此脚本"
    exit 1
fi

echo "🔍 检查 ${GOMOD_PATH} 中的禁止依赖..."
echo ""

# 记录是否发现禁止依赖
FOUND_FORBIDDEN=0

for dep in "${FORBIDDEN_DEPS[@]}"; do
    if grep -q "${dep}" "${GOMOD_PATH}"; then
        echo "❌ 发现禁止依赖：${dep}"
        FOUND_FORBIDDEN=1
    fi
done

echo ""

# 输出检查结果
if [ ${FOUND_FORBIDDEN} -eq 1 ]; then
    echo "❌ 检查失败：kit/go.mod 包含禁止依赖"
    echo "   kit/ 模块不应依赖 kratos、gorm、redis 等重型框架"
    exit 1
else
    echo "✅ 检查通过：kit/go.mod 不包含任何禁止依赖"
    exit 0
fi
