// 查看 yarn 全局安装过的包
yarn global list --depth=0
// 全局安装
yarn global add <package>
// 安装所有依赖
yarn
// 安装指定版本
yarn add <package>@<version>
// 安装最新版本
yarn add <package>
// 删除包,会更新package.json和yarn.lock
yarn remove <package>
// 更新包
yarn upgrade
// 更新指定的包
yarn upgrade <package>
// 获取可更新的包列表选择更新
yarn upgrade-interactive --latest
// 更新全局依赖
yarn global upgrade --latest
// 更新全局依赖，有交互
yarn global upgrade-interactive --latest
// 列出已缓存的包
yarn cache list
// 查找缓存包的路径
yarn cache dir
// 清除缓存的包
yarn cache clean


// 查看 npm 全局安装过的包
npm list -g --depth 0
// 全局安装
npm install <package> -g
// 安装项目所有依赖
npm install
// 安装指定版本
npm install <package>@1.2.3
// 安装最新版本
npm i <package>@latest
// 删除全局的包
npm uninstall -g <package>
// 删除 node_modules 目录下面的包
npm uninstall <package>
// 更新指定包
npm update <package>
// 更新指定全局包
npm update -g <package>
// 更新本地安装的包
// 在 package.json 文件所在的目录中执行 npm update 命令
// 执行 npm outdated 命令。不应该有任何输出。

