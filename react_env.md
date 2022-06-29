# React前端环境搭建
React框架的生态丰富，有React-boostrap, ant ui等，对前端开发很友好，让Web开发成为Flutter一样。
然而React的环境搭建比较复杂，下面介绍两种方式。

## 第一种create-react-app
### 创建项目
先安装react项目快速构建工具：
```bash
npm install create-react-app --location=global
```
然后创建项目，
```bash
npx create-react-app AppName
```
### 整理项目
这个会创建一个名为AppName的文件夹，里面会有一些文件。
```txt
AppName
    node_modules
    public
    src
    .gitignore
    package.json
    package-lock.json
```
其中，public和src下面的文件全部都可以删除。
然后在public下面新建：
```txt
css
js
pages
index.html
```
public文件夹的意思是发布的意思，上面的结构就和传统web结构相同。
在src下面新建：
```txt
index.jsx
```
src里面是js源代码，会打包加密好，放在public/js下面。

### 编译，测试项目
在package.json下面可以看到相关的命令
```
npm run start
npm run build
```

## 手动搭建react环境
react环境需要以下包：
1. react：react核心
2. react-dom：react-dom相关操作
3. webpack：js打包工具
4. webpack-dev-server：热更新调试工具
5. @babel/core：babel核心，用来增强js功能的
6. @babel/preset-env：用于把js es6代码转换成es5
7. @babel/preset-react：用于把jsx语法，转换成es5语法
8. babel-loader：用于webpack和babel的配合
9. style-loader：用于css文件的嵌入style语块
10. css-loader：用于打包css文件
11. antd：ant ui，可以选择其他的

可见，需要安装非常多的东西。

### webpack配置
上面的包安装好后，最重要的是配置webpack
在项目新建webpack.config.js
按照webpack官方手册去配置，这里给出配置样例。
```js
const config = {
    mode: "development",
    entry: {
        index: "./src/index.jsx"
    },
    devServer: {
        static: {
            directory: __dirname + "/public/"
        },
        port: 8080,
        open: true,
        hot: true
    },
    module: {
        rules: [{
            test: /.jsx$/,
            exclude: "/node_modules/",
            use: {
                loader: "babel-loader",
                options: {
                    presets: [
                        "@babel/preset-env",
                        "@babel/preset-react"
                    ]
                }
            }
        }]
    },
    output: {
        filename: "[name].bundle.js",
        path: __dirname + "/public/js/"
    }
}

module.exports = config;
```

## 热更新
webpack一般就用于打包js，所以这些js还是需要自己在html里面手动引用，并且以type=module的方式引入。

如果需要支持热更新，在引用的时候js只需要填写文件名，而不用填写具体的路径，否则热更新无效。
如下给出样例。
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="index.bundle.js" type="module"></script>
    <!-- <script src="js/index.bundle.js" type="module"></script> -->
    <link href="css/antd.css" rel="stylesheet">
    <link href="css/index.css" rel="stylesheet">
    <title>Rainy Panel</title>
</head>

<body>
    <div id="root" class="root">

    </div>
</body>

</html>
```

## React的开发方式
React拥抱了jsx，Vue拥抱了template。
jsx允许在js文件里面写html，
template是分开的html模板

这就使得，react的开发方式十分类似Flutter，界面的设计基本就是在js里面完成的。
如下图所示，
```js
import React from "react"
import ReactDOM from "react-dom/client"
import { Menu, Col, Row } from "antd"

const items = [{
    label: "菜单111",
    key: "key-1"
}, {
    label: "菜单222",
    key: "key-2"
}, {
    label: "菜单33",
    key: "key-3"
}, {
    label: "菜单4",
    key: "key-4"
}, {
    label: "菜单5",
    key: "key-5"
}];

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <div className="root">
        <Row align="middle">
            <Col span={8} className="header"><img></img></Col>
            <Col span={8} className="header"><a>Rainy Control Panel</a></Col>
            <Col span={8} className="header"></Col>
        </Row>
        <Row align="middle">
            <Col span={4}>
                <Menu mode="vertical" items={items} className="menu_class"></Menu>
            </Col>
            <Col span={20}>

            </Col>
        </Row>
    </div>
);
```