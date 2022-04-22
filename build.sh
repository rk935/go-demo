#!/usr/bin/env bash
echo 'build done.'
#set -x

# 获取git tag版本号
GitTag=`git tag --sort=version:refname | tail -n 1`
# 获取源码最近一次git commit log，包含commit sha 值，以及commit message
GitCommitLog=`git log --pretty=oneline -n 1`
# 将log原始字符串中的单引号替换成双引号
GitCommitLog=${GitCommitLog//\'/\"}
# 检查源码在git commit基础上，是否有本地修改，且未提交的内容
GitStatus=`git status -s`
# 获取当前时间
BuildTime=`date +'%Y.%m.%d.%H%M%S'`
# 获取Go的版本
BuildGoVersion=`go version`

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'github.com/q191201771/naza/pkg/bininfo.GitTag=${GitTag}' \
    -X 'github.com/q191201771/naza/pkg/bininfo.GitCommitLog=${GitCommitLog}' \
    -X 'github.com/q191201771/naza/pkg/bininfo.GitStatus=${GitStatus}' \
    -X 'github.com/q191201771/naza/pkg/bininfo.BuildTime=${BuildTime}' \
    -X 'github.com/q191201771/naza/pkg/bininfo.BuildGoVersion=${BuildGoVersion}' \
"

ROOT_DIR=`pwd`

# 如果可执行程序输出目录不存在，则创建
if [ ! -d ${ROOT_DIR}/bin ]; then
  mkdir bin
fi

# 编译多个可执行程序
for file in `ls ${ROOT_DIR}/demo`
do
  if [ -d ${ROOT_DIR}/demo/${file} ]; then
    echo "build" ${ROOT_DIR}/demo/${file} "..."
    cd ${ROOT_DIR}/demo/${file} && go build -ldflags "$LDFlags" -o ${ROOT_DIR}/bin/${file}
  fi
done