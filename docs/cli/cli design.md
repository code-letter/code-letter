# CLI 设计

## 命令列表

### add

* 描述: 将 review comment 记录在文件中
* 参数:
    * commit object hash
        * 类型: string
        * 描述: commit 的 hash 值
    * file path
        * 类型: string
        * 描述: 文件在 repo 根文件夹下的相对路径
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
