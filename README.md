# code review tool

想做一个可以将 code review comment 直接存储在本地代码库中的一个工具

## 初步计划

1. cli: 作为一个 CLI 工具提供将 code review comment 直接记录在本地文件中的能力
2. IDE plugin: 图形界面，调用 CLI 的能力，目前首先集成 IntellJ

## CLI 设计

### 命令列表

#### add

* 描述: 将 review comment 记录在文件中
* 参数: 
    * commit object hash 
        * 类型: string
        * 描述: commit 的 hash 值
    * file path
        * 类型: string
        * 描述: 数据对象的 hash 值，用于在 git 中快速定位文件的内容
    * lines:
        * 类型：string
        * 描述: 文件中的代码行数，如果是对多行代码记录评审意见格式为： x,y,z 例如 12,13,14
    * comment
        * 类型: string
        * 描述: 代码评审意见
* flag:
    * labels
        * 类型: \[string\]string
        * 描述: 代码评审意见标签，键值对形势，例如 type: style; level: error
* 返回值
    * comment record id
        * 类型: string
        * 描述: 内部记录的代码评审意见对应的唯一的 id
