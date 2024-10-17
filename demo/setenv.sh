# go env -w GOPROXY=https://goproxy.cn,direct &&
# go install golang.org/x/tools/gopls@latest &&
# go install github.com/go-delve/delve/cmd/dlv@latest &&
# go install -v golang.org/x/tools/cmd/goimports@latest &&
# go install github.com/daixiang0/gci@latest &&
pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
pip3 install resp-benchmark